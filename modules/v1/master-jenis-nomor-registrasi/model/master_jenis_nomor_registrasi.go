package model

type JenisNomorRegistrasi struct {
	ID              string `json:"-"`
	JenisNomorRegis string `json:"jenis_no_regis"`
	KdJenisRegis    string `json:"kd_jenis_regis"`
	UserInput       string `json:"-"`
	UserUpdate      string `json:"-"`
	UUID            string `json:"uuid"`
}

type JenisNomorRegistrasiResponse struct {
	Data []JenisNomorRegistrasi `json:"data"`
}
