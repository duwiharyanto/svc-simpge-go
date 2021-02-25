package repo

import (
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-jenis-nomor-registrasi/model"
)

func GetJenisNomorRegistrasi(a app.App) ([]model.JenisNomorRegistrasi, error) {
	sqlQuery := getJenisNomorRegistrasiQuery()
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get jenis nomor registrasi, %w", err)
	}
	defer rows.Close()

	pp := []model.JenisNomorRegistrasi{}
	for rows.Next() {
		var p model.JenisNomorRegistrasi
		err := rows.Scan(
			&p.ID,
			&p.KdJenisRegis,
			&p.JenisNomorRegis,
			&p.UUID,
		)
		if err != nil {
			return nil, fmt.Errorf("error scan jenis nomor registrasi row, %s", err.Error())
		}
		pp = append(pp, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error jenis nomor registrasi rows, %s", err.Error())
	}

	return pp, nil
}
