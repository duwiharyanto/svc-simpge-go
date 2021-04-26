package repo

import (
	"database/sql"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/sk/model"
)

func GetAllPegawaiPenetapSK(a app.App) ([]model.PegawaiPengangkat, error) {
	// c.Param("kd_jenis_pegawai")
	sqlQuery := getPegawaiPenetapSKQuery()
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get pangkat pegawai, %s", err.Error())
	}
	defer rows.Close()

	pegawaiPengangkat := []model.PegawaiPengangkat{}
	for rows.Next() {
		var s model.PegawaiPengangkat
		err := rows.Scan(&s.NIK, &s.Nama, &s.GelarDepan, &s.GelarBelakang, &s.KdKelompokPegawai, &s.UUID)
		if err != nil {
			return nil, fmt.Errorf("error scan pegawai pengangkat sk row, %s", err.Error())
		}
		pegawaiPengangkat = append(pegawaiPengangkat, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error status pangkat pegawai rows, %s", err.Error())
	}

	return pegawaiPengangkat, nil
}
func GetPegawaiPenetapSKByUUID(a app.App, uuid string) (*model.PegawaiPengangkat, error) {
	sqlQuery := getPegawaiPenetapSKQueryByUUID(uuid)
	//fmt.Printf("\n[DEBUG] QUERY : %s \n", sqlQuery)
	var penangkatPegawai model.PegawaiPengangkat
	err := a.DB.QueryRow(sqlQuery).Scan(&penangkatPegawai.ID, &penangkatPegawai.NIK, &penangkatPegawai.Nama, &penangkatPegawai.GelarDepan, &penangkatPegawai.GelarBelakang, &penangkatPegawai.KdKelompokPegawai, &penangkatPegawai.KdUnit, &penangkatPegawai.UUID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error querying get pegawai pegawai by uuid, %s", err.Error())
	}
	//fmt.Printf("\n[DEBUG] penangkatPegawai : %+v \n", penangkatPegawai)
	return &penangkatPegawai, nil
}
