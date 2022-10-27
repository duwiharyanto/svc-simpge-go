package model

type PejabatStrukturalPrivate struct {
	IdPegawai      string
	IdJenisUnit    string `json:"id_jenis_unit"`
	IdJenisJabatan string `json:"id_jenis_jabatan"`
	IdUnit         string `json:"id_unit"`
}
