package repo

import (
	"context"
	"errors"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v2/organisasi/model"

	"gorm.io/gorm"
)

func GetAllJabatanStruktural(a *app.App, ctx context.Context) ([]model.JabatanStruktural, error) {
	var jj []model.JabatanStruktural
	err := a.GormDB.
		WithContext(ctx).
		Where(
			// &model.JabatanStruktural{FlagAktif: 1}, // untuk sk pengangkatan tampilkan semua?
			"kd_unit IN ? AND nama_jenis_jabatan IN ?", []string{"000", "100"}, []string{"Ketua", "Rektor"},
		).
		Find(&jj).
		Error

	if err != nil {
		return nil, fmt.Errorf("error get all jabatan struktural: %w", err)
	}

	return jj, nil
}

func GetJabatanStruktural(a *app.App, ctx context.Context, uuid string) (*model.JabatanStruktural, error) {
	var js model.JabatanStruktural
	err := a.GormDB.
		WithContext(ctx).
		Where(&model.JabatanStruktural{
			// FlagAktif: 1, // untuk sk pengangkatan tampilkan semua
			Uuid: uuid,
		}).
		First(&js).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("error get jabatan struktural: %w", err)
	}

	return &js, nil
}

func GetPejabatStruktural(a *app.App, ctx context.Context, uuidJabatanStruktural string) ([]model.PejabatStruktural, error) {
	var pp []model.PejabatStruktural
	err := a.GormDB.
		WithContext(ctx).
		Joins("JabatanStruktural").
		Where("JabatanStruktural.uuid = ?", uuidJabatanStruktural).
		// Where(&model.PejabatStruktural{ FlagAktif: 1, // untuk sk pengangkatan tampilkan semua}).
		Find(&pp).
		Error

	if err != nil {
		return nil, fmt.Errorf("error get pejabat struktural: %w", err)
	}

	return pp, nil
}

func GetPejabatStrukturalByUUID(a *app.App, ctx context.Context, uuid string) (*model.PejabatStruktural, error) {
	var ps model.PejabatStruktural
	err := a.GormDB.
		WithContext(ctx).
		Where(&model.PejabatStruktural{
			// FlagAktif: 1, // untuk sk pengangkatan tampilkan semua
			Uuid: uuid,
		}).
		Find(&ps).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("error get pejabat struktural by uuid: %w", err)
	}

	return &ps, nil
}
