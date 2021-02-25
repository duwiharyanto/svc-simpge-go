package model

type LokasiKerja struct {
	ID          string `json:"-"`
	LokasiKerja string `json:"lokasi_kerja"`
	LokasiDesc  string `json:"lokasi_desc"`
	UserInput   string `json:"-"`
	UserUpdate  string `json:"-"`
	UUID        string `json:"uuid"`
}

type LokasiKerjaResponse struct {
	Data []LokasiKerja `json:"data"`
}
