package repo

import (
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/sk/model"
)

// sk list

func GetAllSKPegawaiV0(a *app.App, req *model.SKPegawaiRequest) ([]model.SKPegawai, error) {
	var sqlQuery string
	if req.KdJenisPegawai == "ED" {
		sqlQuery = getAllSKDosenQuery(req.UUIDPegawai)
	} else {
		sqlQuery = getAllSKTendikQuery(req.UUIDPegawai)
	}
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get sk pegawai, %w", err)
	}
	defer rows.Close()

	ss := []model.SKPegawai{}
	for rows.Next() {
		var s model.SKPegawai
		err := rows.Scan(
			&s.UnitKerja,
			&s.JenisSK,
			&s.KdJenisSK,
			&s.UuidJenisSK,
			&s.NomorSK,
			&s.TMT,
			&s.TglUpdate,
			&s.UUID,
		)
		if err != nil {
			return nil, fmt.Errorf("error scan pegawai row, %s", err.Error())
		}
		s.SetTMTIDN()
		ss = append(ss, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error pegawai rows, %s", err.Error())
	}

	return ss, nil
}

func GetAllSKPegawai(a *app.App, req *model.SKPegawaiRequest) ([]model.SKPegawai, error) {
	var sqlQuery = getAllSkQuery(req.UUIDPegawai)
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get sk pegawai, %w", err)
	}
	defer rows.Close()

	ss := []model.SKPegawai{}
	for rows.Next() {
		var s model.SKPegawai
		err := rows.Scan(
			&s.UnitKerja,
			&s.JenisSK,
			&s.KdJenisSK,
			&s.UuidJenisSK,
			&s.NomorSK,
			&s.TMT,
			&s.TglUpdate,
			&s.UUID,
		)
		if err != nil {
			return nil, fmt.Errorf("error scan pegawai row, %s", err.Error())
		}
		s.SetTMTIDN()
		ss = append(ss, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error pegawai rows, %s", err.Error())
	}

	return ss, nil

}
