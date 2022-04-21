package repo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"svc-insani-go/app"
	"svc-insani-go/app/helper"
	"svc-insani-go/modules/v1/personal/model"

	"gorm.io/gorm"
)

func SearchPersonal(a *app.App, ctx context.Context, cari string) ([]model.PersonalDataPribadi, error) {

	personals := []model.PersonalDataPribadi{}
	tx := a.GormDB.WithContext(ctx)

	if cari != "" {
		q := `SELECT x.* FROM (
			SELECT a.id, a.nama_lengkap, a.gelar_depan, a.gelar_belakang, a.nik_ktp, a.nik_pegawai, a.uuid FROM personal_data_pribadi a
			WHERE a.flag_aktif = 1 AND a.nama_lengkap LIKE ?
			UNION
			SELECT b.id, b.nama_lengkap, b.gelar_depan, b.gelar_belakang, b.nik_ktp, b.nik_pegawai, b.uuid FROM personal_data_pribadi b
			WHERE b.flag_aktif = 1 AND b.nik_ktp LIKE ?
		) x LEFT JOIN pegawai p ON x.id = p.id_personal_data_pribadi WHERE p.id IS NULL`
		res := tx.Raw(helper.FlatQuery(q),
			"%"+cari+"%", "%"+cari+"%",
		).Find(&personals)
		if res.Error != nil {
			return nil, res.Error
		}
	}

	return personals, nil
}

func GetPersonalByUuid(a *app.App, ctx context.Context, uuid string) (*model.PersonalDataPribadiId, error) {
	var personal model.PersonalDataPribadiId
	err := a.GormDB.
		WithContext(ctx).
		Preload("Pegawai.PegawaiFungsional.StatusPegawaiAktif").
		Joins("Agama").
		Joins("GolonganDarah").
		Joins("StatusPernikahan").
		Joins("Pegawai").
		Where("personal_data_pribadi.flag_aktif = 1 AND personal_data_pribadi.uuid = ?", uuid).
		First(&personal).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	if strings.ToLower(personal.GolonganDarah.GolonganDarah) == model.UnknownBloodType {
		personal.GolonganDarah = model.GolonganDarah{}
	}

	return &personal, nil
}

func PersonalActivation(uuidPersonal string) error {
	var client = &http.Client{}
	var data model.PersonalActivationResponse

	baseURL := os.Getenv("URL_ACTIVATION_PERSONAL")
	destinationURL := baseURL + "/public/api/v1/" + uuidPersonal
	request, err := http.NewRequest("PUT", destinationURL, nil)
	if err != nil {
		fmt.Printf("[ERROR] error created http request - %s - at modules/v1/personal/repo/personal_repo.go - PersonalActivation()\n", err)
		return err
	}
	fmt.Println("[DEBUG] send request to personal, ", destinationURL)
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("[ERROR] error sending http request - %s - at modules/v1/personal/repo/personal_repo.go - PersonalActivation()\n", err)
		return err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Printf("[ERROR] %s - at modules/v1/personal/repo/personal_repo.go - PersonalActivation()\n", err)
		return err
	}

	statusCodeOK := 200
	if response.StatusCode != statusCodeOK {
		fmt.Println("error status not OK ", data)
		return fmt.Errorf("error status not ok: %s", response.Body)
	}
	fmt.Println("[DEBUG] berhasil aktivasi personal: ", response.Body)
	return nil
}
