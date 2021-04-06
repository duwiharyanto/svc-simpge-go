package model

type JenisSKPengangkatan struct {
	KdJenisSKPengangkatan string `json:"kd_jenis_sk_pengangkatan"`
	JenisSKPengangkatan   string `json:"jenis_sk_pengangkatan"`
	KdKelompokPegawai     string `json:"kd_kelompok_pegawai"`
	UUID                  string `json:"uuid"`
	ID                    string `json:"-"`
}

type JenisSKPengangkatanResponse struct {
	Data []JenisSKPengangkatan `json:"data"`
}
