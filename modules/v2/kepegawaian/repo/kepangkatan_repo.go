package repo

import (
	"context"
	"errors"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v2/kepegawaian/model"

	"gorm.io/gorm"
)

// TODO: tambah handle dan return error
func GetAllPangkatGolonganRuang(a app.App, ctx context.Context) []model.PangkatGolonganRuang {
	var pp []model.PangkatGolonganRuang
	a.GormDB.
		WithContext(ctx).
		// Where(&model.PangkatGolonganRuang{FlagAktif: 1}). // untuk sk pengangkatan tampilkan semua flag
		Find(&pp)
	return pp
}

func GetPangkatGolonganRuang(a app.App, ctx context.Context, uuid string) (*model.PangkatGolonganRuang, error) {
	var p model.PangkatGolonganRuang
	err := a.GormDB.
		WithContext(ctx).
		Where(&model.PangkatGolonganRuang{
			FlagAktif: 1, // untuk sk pengangkatan tampilkan semua flag?
			Uuid:      uuid,
		}).
		First(&p).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("error get pangkat golongan ruang: %w", err)
	}

	return &p, nil
}

func GetAllJabatanFungsional(a app.App, ctx context.Context) ([]model.JabatanFungsional, error) {
	var jj []model.JabatanFungsional
	err := a.GormDB.
		WithContext(ctx).
		// Where(&model.JabatanFungsional{FlagAktif: 1}). // untuk sk pengangkatan tampilkan semua flag
		Find(&jj).
		Error

	if err != nil {
		return nil, err
	}

	return jj, nil
}

func GetJabatanFungsional(a app.App, ctx context.Context, uuid string) (*model.JabatanFungsional, error) {
	var j model.JabatanFungsional
	err := a.GormDB.
		WithContext(ctx).
		Where(&model.JabatanFungsional{
			FlagAktif: 1, // untuk sk pengangkatan tampilkan semua flag?
			Uuid:      uuid,
		}).
		First(&j).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &j, nil
}
