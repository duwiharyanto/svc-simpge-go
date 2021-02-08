package repo

import (
	"database/sql"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-unit-kerja/model"
)

func GetAllUnitKerja(a app.App) ([]model.UnitKerja, error) {
	// c.Param("kd_jenis_pegawai")
	sqlQuery := getUnitKerjaQuery()
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get unit kerja pegawai, %s", err.Error())
	}
	defer rows.Close()

	UnitKerja := []model.UnitKerja{}
	for rows.Next() {
		var s model.UnitKerja
		err := rows.Scan(&s.KdUnitKerja, &s.NamaUnitKerja, &s.UUID)
		if err != nil {
			return nil, fmt.Errorf("error scan unit kerja pegawai row, %s", err.Error())
		}
		UnitKerja = append(UnitKerja, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error status unit pegawai rows, %s", err.Error())
	}

	return UnitKerja, nil
}

func GetUnitKerjaByUUID(a app.App, uuid string) (*model.UnitKerja, error) {
	sqlQuery := getUnitKerjaByUUID(uuid)
	fmt.Printf("\n\n[DEBUG] query unit kerja : \n%s\n", sqlQuery)
	var unitKerja model.UnitKerja
	err := a.DB.QueryRow(sqlQuery).Scan(&unitKerja.ID, &unitKerja.KdUnitKerja, &unitKerja.NamaUnitKerja, &unitKerja.UUID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error querying get master unit kerja pegawai by uuid %s", err.Error())
	}
	return &unitKerja, nil
}
