package repo

import (
	"database/sql"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/sk/model"
)

func GetAllJenisSK(a app.App) ([]model.JenisSK, error) {
	// c.Param("kd_jenis_pegawai")
	sqlQuery := getJenisSKQuery()
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get jenis sk pegawai, %s", err.Error())
	}
	defer rows.Close()

	JenisSK := []model.JenisSK{}
	for rows.Next() {
		var s model.JenisSK
		err := rows.Scan(&s.KdJenisSK, &s.JeniSKPengangkatan, &s.UUID)
		if err != nil {
			return nil, fmt.Errorf("error scan jenis sk pegawai row, %s", err.Error())
		}
		JenisSK = append(JenisSK, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error status golongan pegawai rows, %s", err.Error())
	}

	return JenisSK, nil
}

func GetJenisSKByUUID(a app.App, uuid string) (*model.JenisSK, error) {
	sqlQuery := getJenisSKByUUID(uuid)
	var jenisSK model.JenisSK
	err := a.DB.QueryRow(sqlQuery).Scan(&jenisSK.ID, &jenisSK.KdJenisSK, &jenisSK.JeniSKPengangkatan, &jenisSK.UUID)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error querying get jenis sk by uuid %s", err.Error())
	}
	return &jenisSK, nil
}
