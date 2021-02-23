package repo

import (
	"database/sql"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-pangkat-golongan-pegawai/model"
)

func GetAllPangkatPegawai(a app.App) ([]model.PangkatPegawai, error) {
	sqlQuery := getPangkatPegawaiQuery()
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get pangkat pegawai, %s", err.Error())
	}
	defer rows.Close()

	PangkatPegawai := []model.PangkatPegawai{}
	for rows.Next() {
		var s model.PangkatPegawai
		err := rows.Scan(&s.KdPangkat, &s.NamaPangkat, &s.UUID)
		if err != nil {
			return nil, fmt.Errorf("error scan pangkat pegawai row, %s", err.Error())
		}
		PangkatPegawai = append(PangkatPegawai, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error status pangkat pegawai rows, %s", err.Error())
	}

	return PangkatPegawai, nil
}
func GetPangkatPegawaiByUUID(a app.App, uuid string) (*model.PangkatPegawai, error) {
	sqlQuery := getPangkatPegawaiByUUIDQuery(uuid)
	var pangkatPegawai model.PangkatPegawai
	err := a.DB.QueryRow(sqlQuery).Scan(&pangkatPegawai.ID, &pangkatPegawai.KdPangkat, &pangkatPegawai.NamaPangkat, &pangkatPegawai.UUID)
	if err != nil {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error querying get golongan pegawai by uuid, %s", err.Error())
	}
	return &pangkatPegawai, nil
}

func GetPangkatGolonganPegawai(a app.App) ([]model.PangkatGolonganPegawai, error) {
	sqlQuery := getPangkatGolonganPegawaiQuery()
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get pangkat golongan pegawai, %s", err.Error())
	}
	defer rows.Close()

	PangkatGolPegawai := []model.PangkatGolonganPegawai{}
	for rows.Next() {
		var s model.PangkatGolonganPegawai
		err := rows.Scan(&s.Pangkat, &s.Golongan, &s.UUID)
		if err != nil {
			return nil, fmt.Errorf("error scan pangkat golongan pegawai row, %s", err.Error())
		}
		PangkatGolPegawai = append(PangkatGolPegawai, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error status pangkat golongan pegawai rows, %s", err.Error())
	}

	return PangkatGolPegawai, nil
}

func GetPangkatGolonganPegawaiByUUID(a app.App, uuid string) (*model.PangkatGolonganPegawai, error) {
	sqlQuery := getPangkatGolonganPegawaiQueryByUUID(uuid)
	var pangkatGolonganPegawai model.PangkatGolonganPegawai
	err := a.DB.QueryRow(sqlQuery).Scan(&pangkatGolonganPegawai.ID, &pangkatGolonganPegawai.KdPangkat, &pangkatGolonganPegawai.Pangkat, &pangkatGolonganPegawai.Golongan, &pangkatGolonganPegawai.UUID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error querying get pangkat golongan pegawai by uuid, %s", err.Error())
	}
	return &pangkatGolonganPegawai, nil
}
