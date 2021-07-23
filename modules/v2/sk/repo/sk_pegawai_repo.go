package repo

import (
	"context"
	"errors"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v2/sk/model"

	"gorm.io/gorm"
)

// TODO: add error handling and return it
func GetAllSkPegawai(a *app.App, ctx context.Context) []model.SkPegawai {
	var ssp []model.SkPegawai
	a.GormDB.
		WithContext(ctx).
		Where(&model.SkPegawai{FlagAktif: 1}).
		Find(&ssp)
	return ssp
}

// TODO: add error handling and return it
func GetSkPegawai(a *app.App, ctx context.Context, uuid string) *model.SkPegawai {
	var skp model.SkPegawai
	err := a.GormDB.
		WithContext(ctx).
		Joins("Pegawai").
		// Where(&model.SkPegawai{Uuid: uuid}).
		First(&skp).
		// First(&skp, "sk_pengangkatan_tendik.uuid = ?", uuid).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	return &skp
}

func GetAllJenisIjazah(a *app.App, ctx context.Context) ([]model.JenisIjazah, error) {
	var jj []model.JenisIjazah
	err := a.GormDB.
		WithContext(ctx).
		// Where(&model.JenisIjazah{FlagAktif: 1}). //?
		Find(&jj).
		Error

	if err != nil {
		return nil, err
	}

	return jj, nil
}

func GetJenisIjazah(a *app.App, ctx context.Context, uuid string) (*model.JenisIjazah, error) {
	var j model.JenisIjazah
	err := a.GormDB.
		WithContext(ctx).
		Where(&model.JenisIjazah{
			FlagAktif: 1,
			Uuid:      uuid,
		}).
		Find(&j).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("error get jenis ijazah: %w", err)
	}

	return &j, nil
}
