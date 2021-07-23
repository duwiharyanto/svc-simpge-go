package repo

import (
	"context"
	"errors"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-status-pegawai-aktif/model"

	"gorm.io/gorm"
)

func GetStatusPegawaiAktif(a *app.App, FlagStatus string) ([]model.StatusPegawaiAktif, error) {

	sqlQuery := getStatusPegawaiAktifAllQuery()

	if FlagStatus != "" {
		sqlQuery = getStatusPegawaiAktifQuery(FlagStatus)
	}

	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get status pegawai aktif, %w", err)
	}
	defer rows.Close()

	pp := []model.StatusPegawaiAktif{}
	for rows.Next() {
		var p model.StatusPegawaiAktif
		err := rows.Scan(
			&p.ID,
			&p.KdStatusAktif,
			&p.StatusAktif,
			&p.UUID,
		)
		if err != nil {
			return nil, fmt.Errorf("error scan status pegawai aktif row, %s", err.Error())
		}
		pp = append(pp, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error status pegawai aktif rows, %s", err.Error())
	}

	return pp, nil
}

func GetStatusPegawaiAktifByUUID(a *app.App, ctx context.Context, uuid string) (*model.StatusPegawaiAktif, error) {
	var StatusPegawaiAktif model.StatusPegawaiAktif

	err := a.GormDB.
		WithContext(ctx).
		Where("flag_aktif = 1 AND uuid = ?", uuid).
		First(&StatusPegawaiAktif).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error get status pegawai aktif by uuid %s", err)
	}
	return &StatusPegawaiAktif, nil
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
