package repo

import (
	"database/sql"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-unit-kerja/model"
)

func GetAllUnitPengangkat(a app.App) ([]model.UnitPengangkat, error) {
	// c.Param("kd_jenis_pegawai")
	sqlQuery := getUnitPengangkatQuery()
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get jenis sk pegawai, %s", err.Error())
	}
	defer rows.Close()

	unitPengangkat := []model.UnitPengangkat{}
	for rows.Next() {
		var s model.UnitPengangkat
		err := rows.Scan(&s.KdUnitPengangkat, &s.UnitPengangkat, &s.UUID)
		if err != nil {
			return nil, fmt.Errorf("error scan jenis sk pegawai row, %s", err.Error())
		}
		unitPengangkat = append(unitPengangkat, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error status golongan pegawai rows, %s", err.Error())
	}

	return unitPengangkat, nil
}
func GetUnitPengangkatByUUID(a app.App, uuid string) (*model.UnitPengangkat, error) {
	sqlQuery := getUnitPengangkatQueryByUUID(uuid)
	var unitPengangkat model.UnitPengangkat
	err := a.DB.QueryRow(sqlQuery).Scan(&unitPengangkat.ID, &unitPengangkat.KdUnitPengangkat, &unitPengangkat.UnitPengangkat, &unitPengangkat.UUID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error querying get unit pengangkat pegawai sk by uuid %s", err.Error())
	}
	return &unitPengangkat, nil
}
