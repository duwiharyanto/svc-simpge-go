package repo

import (
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-jenis-pegawai-tidak-tetap/model"
)

func GetJenisPTT(a app.App) ([]model.JenisPTT, error) {
	sqlQuery := getJenisNomorRegistrasiQuery()
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get jenis pegawai tidak tetap, %w", err)
	}
	defer rows.Close()

	pp := []model.JenisPTT{}
	for rows.Next() {
		var p model.JenisPTT
		err := rows.Scan(
			&p.ID,
			&p.KdJenisPTT,
			&p.JenisPTT,
			&p.UUID,
		)
		if err != nil {
			return nil, fmt.Errorf("error scan jenis pegawai tidak tetap row, %s", err.Error())
		}
		pp = append(pp, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error jenis pegawai tidak tetap rows, %s", err.Error())
	}

	return pp, nil
}
