package repo

import (
	"context"
	"errors"
	"svc-insani-go/app"
	"svc-insani-go/modules/v3/model"

	"gorm.io/gorm"
)

func GetPersonalByUuid(a *app.App, ctx context.Context, uuid string) (*model.PersonalDataPribadi, error) {
	var personal model.PersonalDataPribadi
	err := a.GormDB.
		WithContext(ctx).
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
