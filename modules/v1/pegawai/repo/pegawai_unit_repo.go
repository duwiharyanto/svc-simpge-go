package repo

import (
	"database/sql"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/pegawai/model"
)

func GetUnitKerjaPegawai(a *app.App, uuidPegawai string) (*model.UnitKerjaPegawai, error) {
	sqlQuery := getUnitKerjaPegawaiQuery(uuidPegawai)
	var unitKerjaPegawai model.UnitKerjaPegawai

	err := a.DB.QueryRow(sqlQuery).Scan(
		&unitKerjaPegawai.UuidIndukKerja,
		&unitKerjaPegawai.KdIndukKerja,
		&unitKerjaPegawai.IndukKerja,
		&unitKerjaPegawai.UuidUnitKerja,
		&unitKerjaPegawai.KdUnitKerja,
		&unitKerjaPegawai.UnitKerja,
		&unitKerjaPegawai.UuidBagianKerja,
		&unitKerjaPegawai.KdBagianKerja,
		&unitKerjaPegawai.BagianKerja,
		&unitKerjaPegawai.UuidLokasiKerja,
		&unitKerjaPegawai.LokasiKerja,
		&unitKerjaPegawai.LokasiDesc,
		&unitKerjaPegawai.NoSkPertama,
		&unitKerjaPegawai.TmtSkPertama,
		&unitKerjaPegawai.UuidHomebasePddikti,
		&unitKerjaPegawai.KdHomebasePddikti,
		&unitKerjaPegawai.UuidHomebaseUii,
		&unitKerjaPegawai.KdHomebaseUii,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("error querying and scanning unit kerja pegawai, %s", err.Error())
	}

	unitKerjaPegawai.SetTanggalIDN()

	return &unitKerjaPegawai, nil
}
