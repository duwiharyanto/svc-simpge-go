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
		return nil, fmt.Errorf("error get all kelompok sk pengangkatan")
	}

	return ksk, nil
}

func GetKelompokSkPengangkatan(a app.App, ctx context.Context, uuid string) (*model.KelompokSkPengangkatan, error) {
	var k model.KelompokSkPengangkatan
	err := a.GormDB.
		WithContext(ctx).
		Where(&model.KelompokSkPengangkatan{
			FlagAktif: 1,
			Uuid:      uuid,
		}).
		First(&k).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("error get kelompok sk pengangkatan: %w", err)
	}

	return &k, nil
}

func CreateSkPengangkatanTendik(a app.App, ctx context.Context, skpt *model.SkPengangkatanTendik) error {
	tx := a.GormDB.Session(&gorm.Session{
		Context:              ctx,
		FullSaveAssociations: true,
	})

	result := tx.Create(&skpt)
	if result.Error != nil {
		return fmt.Errorf("error creating sk pengangkatan tendik: %w", result.Error)
	}

	return nil
}

func UpdateSkPengangkatanTendik(a app.App, ctx context.Context, skpt *model.SkPengangkatanTendik) error {
	tx := a.GormDB.Session(&gorm.Session{
		Context:              ctx,
		FullSaveAssociations: true,
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
		Preload("SkPegawai.Pegawai.Unit2").
		Preload("SkPegawai.Pegawai.JenisPegawai").
		Preload("SkPegawai.Pegawai").
		Joins("JabatanPenetap").
		Joins("PejabatPenetap").
		Joins("JabatanFungsional").
		Joins("JenisIjazah").
		Joins("KelompokSkPengangkatan").
		Joins("PangkatGolonganRuang").
		Joins("StatusPengangkatan").
		Joins("UnitKerja").
		Joins("UnitPengangkat").
		Where(&model.SkPengangkatanTendik{
			Uuid:      uuid,
			FlagAktif: 1,
		}).
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

func GetAllStatusPengangkatan(a app.App, ctx context.Context) ([]model.StatusPengangkatan, error) {
	var ss []model.StatusPengangkatan
	err := a.GormDB.
		WithContext(ctx).
		Where(&model.StatusPengangkatan{FlagAktif: 1}).
		Find(&ss).
		Error

	if err != nil {
		return nil, err
	}

	return ss, nil
}

func GetStatusPengangkatan(a app.App, ctx context.Context, uuid string) (*model.StatusPengangkatan, error) {
	var s model.StatusPengangkatan
	err := a.GormDB.
		WithContext(ctx).
		Where(&model.StatusPengangkatan{
			FlagAktif: 1,
			Uuid:      uuid,
		}).
		First(&s).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("error get status pengangkatan: %w", err)
	}

	return &s, nil
}
