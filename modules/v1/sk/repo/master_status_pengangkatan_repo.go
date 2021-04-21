package repo

import (
	"database/sql"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/sk/model"
)

func GetAllStatusPengangkatan(a app.App) ([]model.StatusPengangkatan, error) {
	// c.Param("kd_jenis_pegawai")
	sqlQuery := getStatusPengangkatanQuery()
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get ijazah pegawai, %s", err.Error())
	}
	defer rows.Close()

	statusPengangkatan := []model.StatusPengangkatan{}
	for rows.Next() {
		var s model.StatusPengangkatan
		err := rows.Scan(&s.KdStatusPengangkatan, &s.StatusPengangkatan, &s.UUID)
		if err != nil {
			return nil, fmt.Errorf("error scan status pengangkatan pegawai row, %s", err.Error())
		}
		statusPengangkatan = append(statusPengangkatan, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error status pengangkatan pegawai rows, %s", err.Error())
	}

	return statusPengangkatan, nil
}

func GetStatusPengangkatanByUUID(a app.App, uuid string) (*model.StatusPengangkatan, error) {
	sqlQuery := getStatusPengangkatanQueryByUUID(uuid)
	var statusPengangkatan model.StatusPengangkatan
	err := a.DB.QueryRow(sqlQuery).Scan(&statusPengangkatan.ID, &statusPengangkatan.KdStatusPengangkatan, &statusPengangkatan.StatusPengangkatan, &statusPengangkatan.UUID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error querying get status pengangkat by uuid %s", err.Error())
	}
	return &statusPengangkatan, nil
}
