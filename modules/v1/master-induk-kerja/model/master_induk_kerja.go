package model

type IndukKerja struct {
	ID         string `json:"-"`
	KdUnit     string `json:"kd_unit"`
	Unit       string `json:"unit"`
	Keterangan string `json:"keterangan"`
	UserInput  string `json:"-"`
	UserUpdate string `json:"-"`
	UUID       string `json:"uuid"`
}

type IndukKerjaResponse struct {
	Data []IndukKerja `json:"data"`
}
