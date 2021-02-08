package model

type JenisPegawai struct {
	KDJenisPegawai string `json:"kd_jenis_pegawai"`
	JenisPegawai   string `json:"jenis_pegawai"`
	UUID           string `json:"uuid"`
	ID             string `json:"-"`
}

func (j JenisPegawai) IsEmpty() bool {
	return j == JenisPegawai{}
}

type JenisPegawaiResponse struct {
	Data []JenisPegawai `json:"data"`
}
