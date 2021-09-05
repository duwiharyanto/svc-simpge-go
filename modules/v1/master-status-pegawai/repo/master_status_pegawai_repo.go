package repo

import (
	"context"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-status-pegawai/model"
)

func GetAllStatusPegawai(a *app.App) ([]model.StatusPegawai, error) {
	sqlQuery := getAllStatusPegawaiQuery()
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get status pegawai, %s", err.Error())
	}
	defer rows.Close()

	StatusPegawai := []model.StatusPegawai{}
	for rows.Next() {
		var s model.StatusPegawai
		err := rows.Scan(&s.KDStatusPegawai, &s.StatusPegawai, &s.UUID)
		if err != nil {
			return nil, fmt.Errorf("error scan status pegawai row, %s", err.Error())
		}
		StatusPegawai = append(StatusPegawai, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error jenis status rows, %s", err.Error())
	}

	return StatusPegawai, nil
}

func GetStatusPegawaiByUUID(a *app.App, ctx context.Context, uuid string) (*model.StatusPegawai, error) {
	var statusPegawai model.StatusPegawai

	tx := a.GormDB.WithContext(ctx)
	res := tx.First(&statusPegawai, "uuid = ?", uuid)
	if res.Error != nil {
		return nil, fmt.Errorf("error querying status pegawai by uuid %s", res.Error)
	}
	return &statusPegawai, nil
}
