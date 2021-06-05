package repo

import (
	"context"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/personal/model"
)

func SearchPersonal(a app.App, ctx context.Context, nama string, nik_pegawai string) (*model.PersonalDataPribadi, error) {

	var personal model.PersonalDataPribadi
	tx := a.GormDB.WithContext(ctx)

	if nama != "" {
		res := tx.Where("nama_lengkap LIKE ?", "%"+nama+"%").
			First(&personal)
		if res.Error != nil {
			return nil, res.Error
		}
	}

	if nama == "" && nik_pegawai != "" {
		res := tx.Where("nik_ktp LIKE ?", "%"+nik_pegawai+"%").
			First(&personal)

		if res.Error != nil {
			return nil, res.Error
		}
	}

	return &personal, nil
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
