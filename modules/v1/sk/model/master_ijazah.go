package model

type JenisIjazah struct {
	KdJenisIjazah string `json:"kd_jenis_ijazah"`
	JenisIjazah   string `json:"jenis_ijazah"`
	UUID          string `json:"uuid"`
	ID            string `json:"-"`
}

type JenisIjazahResponse struct {
	Data []JenisIjazah `json:"data"`
}
