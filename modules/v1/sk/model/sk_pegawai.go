package model

import "svc-insani-go/helper"

// sk list
type SKPegawai struct {
	UnitKerja   string `json:"unit_kerja"`
	JenisSK     string `json:"jenis_sk"`
	KdJenisSK   string `json:"kd_jenis_sk"`
	UuidJenisSK string `json:"uuid_jenis_sk"`
	NomorSK     string `json:"nomor_sk"`
	TMT         string `json:"tmt"`
	TMTIDN      string `json:"tmt_idn"`
	TglUpdate   string `json:"-"`
	UUID        string `json:"uuid"`
}

func (sk *SKPegawai) SetTMTIDN() {
	_, sk.TMTIDN = helper.GetIndonesianDate("2006-01-02", sk.TMT)
}

type SKPegawaiRequest struct {
	KdJenisPegawai string `query:"kd_jenis_pegawai"`
	UUIDPegawai    string `query:"uuid_pegawai"`
}

type SKPegawaiResponse struct {
	Count  int         `json:"-"`
	Data   []SKPegawai `json:"data"`
	Limit  int         `json:"-"`
	Offset int         `json:"-"`
}
