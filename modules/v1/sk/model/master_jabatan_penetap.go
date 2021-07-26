package model

type JabatanPenetap struct {
	KdJabatanPenetap string `json:"kd_jabatan_fungsional"`
	JabatanPenetap   string `json:"jabatan_fungsional"`
	UUID             string `json:"uuid"`
	ID               uint64 `json:"-"`
}

type JabatanPenetapResponse struct {
	Data []JabatanPenetap `json:"data"`
}
