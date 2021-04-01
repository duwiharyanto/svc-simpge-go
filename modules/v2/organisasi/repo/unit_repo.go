package repo

import (
	"context"
	"errors"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v2/organisasi/model"

	"gorm.io/gorm"
)

func GetAllUnit2(a app.App, ctx context.Context) []model.Unit2 {
	var uu []model.Unit2
	a.GormDB.
		WithContext(ctx).
		// Where(&model.Unit2{FlagAktif: 1}). // untuk sk pengangkatan tampilkan semua unit
		Find(&uu)
	return uu
}

func GetUnit2(a app.App, ctx context.Context, uuid string) (*model.Unit2, error) {
	var u model.Unit2
	err := a.GormDB.
		WithContext(ctx).
		Where(&model.Unit2{
			// FlagAktif: 1, // untuk sk pengangkatan tampilkan semua unit
			Uuid: uuid,
		}).
		First(&u).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("error get unit2: %w", err)
	}

	return &u, nil
}
