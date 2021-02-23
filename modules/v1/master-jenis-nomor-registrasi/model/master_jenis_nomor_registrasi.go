package model

import (
	jenisPegawai "svc-insani-go/modules/v1/master-jenis-pegawai/model"
	kelompokPegawai "svc-insani-go/modules/v1/master-kelompok-pegawai/model"
	statusPegawai "svc-insani-go/modules/v1/master-status-pegawai/model"
	unitPegawai "svc-insani-go/modules/v1/master-unit-kerja/model"
)

type Pegawai struct {
	ID                              string `json:"-"`
	NIK                             string `json:"nik"`
	Nama                            string `json:"nama"`
	GelarDepan                      string `json:"gelar_depan"`
	GelarBelakang                   string `json:"gelar_belakang"`
	FlagDosen                       int    `json:"flag_dosen"`
	jenisPegawai.JenisPegawai       `json:"jenis_pegawai"`
	kelompokPegawai.KelompokPegawai `json:"kelompok_pegawai"`
	statusPegawai.StatusPegawai     `json:"status_pegawai"`
	unitPegawai.UnitKerja           `json:"unit_kerja"`
	UserInput                       string `json:"-"`
	UserUpdate                      string `json:"-"`
	UUID                            string `json:"uuid"`
}

func (p *Pegawai) SetFlagDosen() {
	if !p.JenisPegawai.IsEmpty() && p.JenisPegawai.KDJenisPegawai == "ED" {
		p.FlagDosen = 1
	}
}

type PegawaiRequest struct {
	KdUnitKerja       string `query:"kd_unit_kerja"`
	KdKelompokPegawai string `query:"kd_kelompok_pegawai"`
	Limit             int    `query:"limit"`
	Offset            int    `query:"offset"`
	Cari              string `query:"cari"`
}
type PegawaiResponse struct {
	Count  int       `json:"count"`
	Data   []Pegawai `json:"data"`
	Limit  int       `json:"limit"`
	Offset int       `json:"offset"`
}
