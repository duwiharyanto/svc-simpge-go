package model

type JabatanFungsional struct {
	KdJabatanFungsional string `json:"kd_jabatan_fungsional"`
	JabatanFungsional   string `json:"jabatan_fungsional"`
	UUID                string `json:"uuid"`
	ID                  string `json:"-"`
}

type JabatanFungsionalResponse struct {
	Data []JabatanFungsional `json:"data"`
}
