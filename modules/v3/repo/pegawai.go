package repo

import (
	"context"
	"errors"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v3/model"

	"gorm.io/gorm"
)

func CreatePegawai(a *app.App, ctx context.Context, pegawai *model.Pegawai) error {
	return nil
}

func GetKelompokPegawaiByUUID(a *app.App, ctx context.Context, uuid string) (*model.KelompokPegawai, error) {
	var kelompokPegawai model.KelompokPegawai

	err := a.GormDB.
		WithContext(ctx).
		Joins("JenisPegawai").
		Joins("StatusPegawai").
		Where("kelompok_pegawai.flag_aktif = 1 AND kelompok_pegawai.uuid = ?", uuid).
		First(&kelompokPegawai).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("error get kelompok pegawai by uuid: %w", err)
	}
	return &kelompokPegawai, nil
}

func GetStatusPegawaiAktifByCode(a *app.App, ctx context.Context, code string) (*model.StatusPegawaiAktif, error) {
	var StatusPegawaiAktif model.StatusPegawaiAktif

	err := a.GormDB.
		WithContext(ctx).
		Where("flag_aktif = 1 AND kd_status = ?", code).
		First(&StatusPegawaiAktif).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("error get status pegawai aktif by code: %w", err)
	}
	return &StatusPegawaiAktif, nil
}
