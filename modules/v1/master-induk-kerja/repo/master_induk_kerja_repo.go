package repo

import (
	"database/sql"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-induk-kerja/model"
)

func GetIndukKerja(a app.App) ([]model.IndukKerja, error) {
	sqlQuery := getIndukKerjaQuery()
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get induk kerja, %w", err)
	}
	defer rows.Close()

	pp := []model.IndukKerja{}
	for rows.Next() {
		var p model.IndukKerja
		err := rows.Scan(
			&p.ID,
			&p.KdUnit,
			&p.Unit,
			&p.Keterangan,
			&p.UUID,
		)
		if err != nil {
			return nil, fmt.Errorf("error scan induk kerja row, %s", err.Error())
		}
		pp = append(pp, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error induk kerja rows, %s", err.Error())
	}

	return pp, nil
}

func GetUnitKerja(a app.App, IndukKerja string) ([]model.IndukKerja, error) {
	sqlQuery := getUnitKerjaQuery(IndukKerja)
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get induk kerja, %w", err)
	}
	defer rows.Close()

	pp := []model.IndukKerja{}
	for rows.Next() {
		var p model.IndukKerja
		err := rows.Scan(
			&p.ID,
			&p.KdUnit,
			&p.Unit,
			&p.Keterangan,
			&p.UUID,
		)
		if err != nil {
			return nil, fmt.Errorf("error scan induk kerja row, %s", err.Error())
		}
		pp = append(pp, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error induk kerja rows, %s", err.Error())
	}

	return pp, nil
}

func GetBagianKerja(a app.App, UnitKerja string) ([]model.IndukKerja, error) {
	sqlQuery := getBagianKerjaQuery(UnitKerja)
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get induk kerja, %w", err)
	}
	defer rows.Close()

	pp := []model.IndukKerja{}
	for rows.Next() {
		var p model.IndukKerja
		err := rows.Scan(
			&p.ID,
			&p.KdUnit,
			&p.Unit,
			&p.Keterangan,
			&p.UUID,
		)
		if err != nil {
			return nil, fmt.Errorf("error scan induk kerja row, %s", err.Error())
		}
		pp = append(pp, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error induk kerja rows, %s", err.Error())
	}

	return pp, nil
}

func GetIndukKerjaByUUID(a app.App, uuid string) (*model.IndukKerja, error) {
	sqlQuery := getIndukKerjaQueryByUUID(uuid)
	//fmt.Printf("log query induk kerja : %s\n", sqlQuery)
	var indukKerja model.IndukKerja
	err := a.DB.QueryRow(sqlQuery).Scan(&indukKerja.ID, &indukKerja.KdIndukKerja, &indukKerja.IndukKerja, &indukKerja.UUID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error querying get jenis induk kerja pegawai by uuid %s", err.Error())
	}
	//fmt.Printf("log query induk kerja datanya ada tidak : %+v\n", indukKerja)
	return &indukKerja, nil
}
