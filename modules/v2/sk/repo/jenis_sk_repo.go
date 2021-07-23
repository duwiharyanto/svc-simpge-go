package repo

import (
	"context"
	"errors"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v2/sk/model"

	"gorm.io/gorm"
)

func GetAllJenisSk(a *app.App, ctx context.Context) ([]model.JenisSk, error) {
	var jjs []model.JenisSk
	err := a.GormDB.
		WithContext(ctx).
		Where(&model.JenisSk{FlagAktif: 1}).
		Find(&jjs).
		Error

	if err != nil {
		return nil, fmt.Errorf("error get all jenis sk: %w", err)
	}

	return jjs, nil
}

func GetJenisSk(a *app.App, ctx context.Context, code string) (*model.JenisSk, error) {
	var js model.JenisSk
	err := a.GormDB.
		WithContext(ctx).
		Where(&model.JenisSk{
			KdJenisSk: code,
			FlagAktif: 1,
		}).
		First(&js).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("error get jenis sk by code: %w", err)
	}

	return &js, nil
}
