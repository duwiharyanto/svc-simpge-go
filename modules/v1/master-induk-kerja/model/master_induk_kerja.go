package model

type IndukKerja struct {
	ID           string `json:"-"`
	KdUnit       string `json:"kd_unit,omitempty"`
	Unit         string `json:"unit,omitempty"`
	Keterangan   string `json:"keterangan"`
	UserInput    string `json:"-"`
	UserUpdate   string `json:"-"`
	UUID         string `json:"uuid"`
	KdIndukKerja string `json:"kd_induk_kerja,omitempty"`
	IndukKerja   string `json:"induk_kerja,omitempty"`
}

type IndukKerjaResponse struct {
	Data []IndukKerja `json:"data"`
}
