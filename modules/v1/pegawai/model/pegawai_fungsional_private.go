package model

type PegawaiFungsionalPrivate struct {
	IdPegawai           uint64 `json:"id_pegawai,omitempty"`
	JabatanFungsional   string `json:"jabatan_fungsional"`
	IdJabatanFungsional uint64 `json:"id_jabatan_fungsional"`
	KdJabatanFungsional string `json:"kd_jabatan_fungsional"`
}
