package repo

import (
	"database/sql"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/pegawai/model"
)

func GetAllPegawai(a app.App, req *model.PegawaiRequest) ([]model.Pegawai, error) {
	sqlQuery := getListAllPegawaiQuery(req)
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get pegawai, %w", err)
	}
	defer rows.Close()

	pp := []model.Pegawai{}
	for rows.Next() {
		var p model.Pegawai
		err := rows.Scan(
			&p.NIK,
			&p.Nama,
			&p.GelarDepan,
			&p.GelarBelakang,
			&p.KelompokPegawai.KdKelompokPegawai,
			&p.KelompokPegawai.KelompokPegawai,
			&p.KelompokPegawai.UUID,
			&p.UnitKerja.KdUnitKerja,
			&p.UnitKerja.NamaUnitKerja,
			&p.UnitKerja.UUID,
			&p.JenisPegawai.KDJenisPegawai,
			&p.JenisPegawai.JenisPegawai,
			&p.JenisPegawai.UUID,
			&p.StatusPegawai.KDStatusPegawai,
			&p.StatusPegawai.StatusPegawai,
			&p.StatusPegawai.UUID,
			&p.UUID,
		)
		if err != nil {
			return nil, fmt.Errorf("error scan pegawai row, %s", err.Error())
		}
		p.SetFlagDosen()
		pp = append(pp, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error pegawai rows, %s", err.Error())
	}

	return pp, nil
}

func CountPegawai(a app.App, req *model.PegawaiRequest) (int, error) {
	sqlQuery := countPegawaiQuery(req)
	var count int
	err := a.DB.QueryRow(sqlQuery).Scan(&count)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	if err != nil {
		return 0, fmt.Errorf("error querying count pegawai, %w", err)
	}

	return count, nil

}
func GetPegawaiByUUID(a app.App, uuid string) (*model.Pegawai, error) {
	sqlQuery := getPegawaiByUUID(uuid)
	var pegawai model.Pegawai
	err := a.DB.QueryRow(sqlQuery).Scan(&pegawai.ID, &pegawai.NIK, &pegawai.Nama, &pegawai.GelarDepan, &pegawai.GelarBelakang, &pegawai.UnitKerja.KdUnitKerja, &pegawai.KelompokPegawai.KdKelompokPegawai, &pegawai.UUID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error querying get pegawai by uuid, %s", err.Error())
	}

	return &pegawai, nil

}
