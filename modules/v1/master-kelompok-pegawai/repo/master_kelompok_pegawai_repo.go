package repo

import (
	"database/sql"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-kelompok-pegawai/model"
)

func GetAllKelompokPegawai(a app.App, IDJenisPegawai string) ([]model.KelompokPegawai, error) {
	// c.Param("kd_jenis_pegawai")
	sqlQuery := getKelompokPegawaiQuery(IDJenisPegawai)
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get kelompok pegawai, %s", err.Error())
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

func GetAllKelompokPegawaiByUUID(a app.App, uuid string) ([]model.KelompokPegawai, error) {
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

func GetKelompokPegawaiByUUID(a app.App, uuid string) (*model.KelompokPegawai, error) {
	var k model.KelompokPegawai
	sqlQuery := getKelompokPegawaiByUUID(uuid)
	err := a.DB.QueryRow(sqlQuery).Scan(
		&k.ID,
		&k.KdStatusPegawai,
		&k.KdJenisPegawai,
		&k.KdKelompokPegawai,
		&k.UUID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error querying get kelompok pegawai by uuid, %s", err.Error())
	}

	return &k, nil
}
