package repo

import (
	"context"
	"errors"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v2/kepegawaian/model"

	"gorm.io/gorm"
)

func GetPegawai(a app.App, ctx context.Context, uuid string) (*model.Pegawai, error) {
	var pgw model.Pegawai
	err := a.GormDB.
		WithContext(ctx).
		Joins("JenisPegawai").
		Joins("Unit2").
		Where(&model.Pegawai{
			FlagAktif: 1, // untuk sk pengangkatan tampilkan semua flag?
			Uuid:      uuid,
		}).
		First(&pgw).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("error get pegawai by uuid: %w", err)
	}

	return &pgw, nil
}
