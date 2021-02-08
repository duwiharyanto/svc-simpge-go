package repo

import (
	"database/sql"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-jenis-pegawai/model"
)

func GetAllJenisPegawai(a app.App) ([]model.JenisPegawai, error) {
	sqlQuery := getJenisPegawaiQuery()
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get jenis pegawai, %s", err.Error())
	}
	defer rows.Close()

	JenisPegawai := []model.JenisPegawai{}
	for rows.Next() {
		var s model.JenisPegawai
		err := rows.Scan(&s.KDJenisPegawai, &s.JenisPegawai, &s.UUID)
		if err != nil {
			return nil, fmt.Errorf("error scan jenis pegawai row, %s", err.Error())
		}
		JenisPegawai = append(JenisPegawai, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error jenis pegawai rows, %s", err.Error())
	}

	return JenisPegawai, nil
}
func GetJenisPegawaiByUUID(a app.App, uuid string) (*model.JenisPegawai, error) {
	sqlQuery := getJenisPegawaiQueryByUUID(uuid)
	var jenisPegawai model.JenisPegawai
	err := a.DB.QueryRow(sqlQuery).Scan(&jenisPegawai.ID, &jenisPegawai.KDJenisPegawai, &jenisPegawai.JenisPegawai, &jenisPegawai.UUID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error querying get jenis pegawai by uuid %s", err.Error())
	}
	return &jenisPegawai, nil
}
func GetJenisPegawaiByKD(a app.App, KDJenisPegawai string) (*model.JenisPegawai, error) {
	sqlQuery := getJenisPegawaiQueryByKdJenisPegawai(KDJenisPegawai)
	var jenisPegawai model.JenisPegawai
	err := a.DB.QueryRow(sqlQuery).Scan(&jenisPegawai.ID, &jenisPegawai.KDJenisPegawai, &jenisPegawai.JenisPegawai, &jenisPegawai.UUID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error querying get jenis pegawai by uuid %s", err.Error())
	}
	return &jenisPegawai, nil
}
