package repo

import (
	"context"
	"errors"
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

	return &personal, nil
}
