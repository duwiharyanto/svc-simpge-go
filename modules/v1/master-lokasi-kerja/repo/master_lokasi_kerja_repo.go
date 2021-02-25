package repo

import (
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-lokasi-kerja/model"
)

func GetLokasiKerja(a app.App) ([]model.LokasiKerja, error) {
	sqlQuery := getLokasiKerjaQuery()
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get lokasi kerja, %w", err)
	}
	defer rows.Close()

	pp := []model.LokasiKerja{}
	for rows.Next() {
		var p model.LokasiKerja
		err := rows.Scan(
			&p.ID,
			&p.LokasiKerja,
			&p.LokasiDesc,
			&p.UUID,
		)
		if err != nil {
			return nil, fmt.Errorf("error scan lokasi kerja row, %s", err.Error())
		}
		pp = append(pp, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error lokasi kerja rows, %s", err.Error())
	}

	return pp, nil
}
