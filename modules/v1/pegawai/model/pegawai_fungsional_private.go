package model

type PegawaiFungsionalDataItemY struct {
	FullPegawaiFungsional
}

type FullPegawaiFungsional struct {
	*PegawaiFungsionalYayasan `json:"jabatan_fungsional_yayasan"`
	*PegawaiFungsionalNegara  `json:"jabatan_fungsional_negara"`
}

type PegawaiFungsionalYayasan struct {
	IdPegawai           uint64 `json:"id_pegawai,omitempty"`
	JabatanFungsional   string `json:"jabatan_fungsional"`
	IdJabatanFungsional uint64 `json:"id_jabatan_fungsional"`
	KdJabatanFungsional string `json:"kd_jabatan_fungsional"`
}

type PegawaiFungsionalNegara struct {
	IdPegawai           uint64 `json:"id_pegawai,omitempty"`
	JabatanFungsional   string `json:"jabatan_fungsional"`
	IdJabatanFungsional uint64 `json:"id_jabatan_fungsional"`
	KdJabatanFungsional string `json:"kd_jabatan_fungsional"`
}
