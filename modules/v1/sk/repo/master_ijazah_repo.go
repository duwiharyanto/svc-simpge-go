package repo

import (
	"database/sql"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/sk/model"
)

func GetAllJenisIjazah(a app.App) ([]model.JenisIjazah, error) {
	// c.Param("kd_jenis_pegawai")
	sqlQuery := getJenisIjazahQuery()
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get ijazah pegawai, %s", err.Error())
	}
	defer rows.Close()

	JenisIjazah := []model.JenisIjazah{}
	for rows.Next() {
		var s model.JenisIjazah
		err := rows.Scan(&s.KdJenisIjazah, &s.JenisIjazah, &s.UUID)
		if err != nil {
			return nil, fmt.Errorf("error scan ijazah pegawai row, %s", err.Error())
		}
		JenisIjazah = append(JenisIjazah, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error status ijazah pegawai rows, %s", err.Error())
	}

	return JenisIjazah, nil
}

func GetJenisIjazahByUUID(a app.App, uuid string) (*model.JenisIjazah, error) {
	sqlQuery := getJenisIjazahByUUID(uuid)
	//fmt.Printf("log query : \n %s\n ", sqlQuery)
	var jenisIjazah model.JenisIjazah
	err := a.DB.QueryRow(sqlQuery).Scan(&jenisIjazah.ID, &jenisIjazah.KdJenisIjazah, &jenisIjazah.JenisIjazah, &jenisIjazah.UUID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error querying get jenis ijazah by uuid %s", err.Error())
	}
	return &jenisIjazah, nil
}
