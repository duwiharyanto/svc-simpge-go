package model

type PejabatStrukturalPrivate struct {
	IdPegawai      uint64 `json:"id_pegawai,omitempty"`
	IdJenisUnit    uint64 `json:"id_jenis_unit"`
	IdJenisJabatan uint64 `json:"id_jenis_jabatan"`
	IdUnit         uint64 `json:"id_unit"`
}
