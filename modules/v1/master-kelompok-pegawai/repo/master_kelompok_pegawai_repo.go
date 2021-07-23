package repo

import (
	"context"
	"errors"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-kelompok-pegawai/model"

	"gorm.io/gorm"
)

func GetAllKelompokPegawai(a *app.App, kdJenisPegawai string) ([]model.KelompokPegawai, error) {
	sqlQuery := getKelompokPegawaiQuery(kdJenisPegawai)
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get kelompok pegawai, %s", err.Error())
	}
	defer rows.Close()

	KelompokPegawai := []model.KelompokPegawai{}
	for rows.Next() {
		var s model.KelompokPegawai
		err := rows.Scan(&s.KdStatusPegawai, &s.KdJenisPegawai, &s.KdKelompokPegawai, &s.KelompokPegawai, &s.UUID)
		if err != nil {
			return nil, fmt.Errorf("error scan kelompok pegawai row, %s", err.Error())
		}
		KelompokPegawai = append(KelompokPegawai, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error status kelompok pegawai rows, %s", err.Error())
	}

	return KelompokPegawai, nil
}

func GetAllKelompokPegawaiByUUID(a *app.App, uuid string) ([]model.KelompokPegawai, error) {
	//c.Param("kd_jenis_pegawai")
	sqlQuery := getAllKelompokPegawaiByUUID(uuid)
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get kelompok pegawai by uuid, %s", err.Error())
	}
	defer rows.Close()

	KelompokPegawai := []model.KelompokPegawai{}
	for rows.Next() {
		var s model.KelompokPegawai
		err := rows.Scan(&s.KdJenisPegawai, &s.KdStatusPegawai, &s.KelompokPegawai, &s.KdKelompokPegawai, &s.UUID)
		if err != nil {
			return nil, fmt.Errorf("error scan kelompok pegawai row, %s", err.Error())
		}
		KelompokPegawai = append(KelompokPegawai, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error status kelompok pegawai rows, %s", err.Error())
	}

	return KelompokPegawai, nil
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
