package repo

import (
	"context"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/personal/model"
)

func SearchPersonal(a app.App, ctx context.Context, cari string) ([]model.PersonalDataPribadi, error) {

	personals := []model.PersonalDataPribadi{}
	tx := a.GormDB.WithContext(ctx)

	if cari != "" {
		res := tx.Where("(nama_lengkap LIKE ? OR nik_ktp LIKE ? ) AND (id NOT IN(SELECT id_personal_data_pribadi FROM pegawai))", "%"+cari+"%", "%"+cari+"%").
			Find(&personals)
		if res.Error != nil {
			return nil, res.Error
		}
	}

	return personals, nil
}

func AllPersonal(a app.App, ctx context.Context) ([]model.PersonalDataPribadi, error) {

	personals := []model.PersonalDataPribadi{}
	tx := a.GormDB.WithContext(ctx)

	res := tx.Where("id NOT IN(SELECT id_personal_data_pribadi FROM pegawai)").
		Find(&personals)
	if res.Error != nil {
		return nil, res.Error
	}

	return personals, nil
}

func GetPersonalByUuid(a app.App, ctx context.Context, uuid string) (*model.PersonalDataPribadiId, error) {

	var personal model.PersonalDataPribadiId
	tx := a.GormDB.WithContext(ctx)

	res := tx.Where("uuid = ?", uuid).
		First(&personal)

	if res.Error != nil {
		return nil, res.Error
	}

	return &personal, nil
}
