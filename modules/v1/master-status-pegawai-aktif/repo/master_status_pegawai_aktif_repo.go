package repo

import (
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-status-pegawai-aktif/model"
)

func GetStatusPegawaiAktif(a app.App, FlagStatus string) ([]model.StatusPegawaiAktif, error) {
	sqlQuery := getStatusPegawaiAktifQuery(FlagStatus)
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get status pegawai aktif, %w", err)
	}
	defer rows.Close()

	pp := []model.StatusPegawaiAktif{}
	for rows.Next() {
		var p model.StatusPegawaiAktif
		err := rows.Scan(
			&p.ID,
			&p.KdStatusAktif,
			&p.StatusAktif,
			&p.UUID,
		)
		if err != nil {
			return nil, fmt.Errorf("error scan status pegawai aktif row, %s", err.Error())
		}
		pp = append(pp, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error status pegawai aktif rows, %s", err.Error())
	}

	return pp, nil
}
