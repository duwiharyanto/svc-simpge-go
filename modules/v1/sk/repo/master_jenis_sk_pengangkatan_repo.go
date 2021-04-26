package repo

import (
	"database/sql"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/sk/model"
)

func GetAllJenisSKPengangkatan(a app.App, kdJenisPegawai string) ([]model.JenisSKPengangkatan, error) {
	sqlQuery := getJenisSKPengangkatanQuery(kdJenisPegawai)
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get jenis sk pengangkatan pegawai, %s", err.Error())
	}
	defer rows.Close()

	jenisSkPengangkatan := []model.JenisSKPengangkatan{}
	for rows.Next() {
		var s model.JenisSKPengangkatan
		err := rows.Scan(&s.KdJenisSKPengangkatan, &s.JenisSKPengangkatan, &s.KdKelompokPegawai, &s.UUID)
		if err != nil {
			return nil, fmt.Errorf("error scan jenis sk pengangkatan pegawai row, %s", err.Error())
		}
		jenisSkPengangkatan = append(jenisSkPengangkatan, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error jenis sk pengangkatan pegawai rows, %s", err.Error())
	}

	return jenisSkPengangkatan, nil
}

func GetJenisSKPengangkatanByUUID(a app.App, uuid string) (*model.JenisSKPengangkatan, error) {
	sqlQuery := getJenisSKPengangkatanQueryByUUID(uuid)
	var jenisSKPengangkatan model.JenisSKPengangkatan
	err := a.DB.QueryRow(sqlQuery).Scan(&jenisSKPengangkatan.ID, &jenisSKPengangkatan.KdJenisSKPengangkatan, &jenisSKPengangkatan.JenisSKPengangkatan, &jenisSKPengangkatan.KdKelompokPegawai, &jenisSKPengangkatan.UUID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error querying get jenis sk pengangkatan by uuid %s", err.Error())
	}
	return &jenisSKPengangkatan, nil
}
