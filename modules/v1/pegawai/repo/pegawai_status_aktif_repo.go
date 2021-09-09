package repo

import (
	"database/sql"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/app/helper"
	"svc-insani-go/modules/v1/pegawai/model"
)

func GetStatusPegawaiAktif(a *app.App, uuid string) (*model.StatusAktif, error) {
	sqlQuery := getStatusPegawaiAktifQuery(uuid)
	var statusAktif model.StatusAktif

	err := a.DB.QueryRow(sqlQuery).Scan(
		&statusAktif.FlagAktifPegawai,
		&statusAktif.KdStatusAktifPegawai,
		&statusAktif.StatusAktifPegawai,
		&statusAktif.UuidStatusAktifPegawai,
		&statusAktif.TglStatusAktifPegawai,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("error querying and scanning status pegawai aktif: %s", err.Error())
	}

	_, statusAktif.TglStatusAktifPegawaiIdn = helper.GetIndonesianDate("2006-01-02", statusAktif.TglStatusAktifPegawai)

	return &statusAktif, nil
}
