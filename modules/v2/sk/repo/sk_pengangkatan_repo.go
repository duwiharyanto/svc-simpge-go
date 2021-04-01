package repo

import (
	"context"
	"errors"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v2/sk/model"

	"gorm.io/gorm"
)

func GetAllKelompokSkPengangkatan(a app.App, ctx context.Context) ([]model.KelompokSkPengangkatan, error) {
	var ksk []model.KelompokSkPengangkatan
	err := a.GormDB.
		WithContext(ctx).
		Where(&model.KelompokSkPengangkatan{FlagAktif: 1}).
		Find(&ksk).
		Error

	if err != nil {
		return nil, err
	}

	return ksk, nil
}

func UpdateSkPengangkatanTendik(a app.App, ctx context.Context, skpt *model.SkPengangkatanTendik) error {
	tx := a.GormDB.Session(&gorm.Session{
		Context: ctx,
		// FullSaveAssociations: true,
	})

	result := tx.Save(&skpt)
	if result.Error != nil {
		return fmt.Errorf("error updating sk pengangkatan tendik: %w", result.Error)
	}

	return nil
}

func GetSkPengangkatanTendik(a app.App, ctx context.Context, uuid string) (*model.SkPengangkatanTendik, error) {
	var skpt model.SkPengangkatanTendik
	err := a.GormDB.
		WithContext(ctx).
		Preload("SkPegawai.Pegawai").
		Joins("KelompokSkPengangkatan").
		Joins("UnitPengangkat").
		Joins("UnitKerja").
		Joins("JabatanFungsional").
		Joins("PangkatGolonganRuang").
		Joins("StatusPengangkatan").
		Joins("JenisIjazah").
		Where(&model.SkPengangkatanTendik{Uuid: uuid}).
		First(&skpt).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &skpt, nil
}
