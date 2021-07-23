package model

type JenisSK struct {
	KdJenisSK          string `json:"kd_jenis_sk_pengangkatan"`
	JeniSKPengangkatan string `json:"jenis_sk_pengangkatan"`
	UUID               string `json:"uuid"`
	ID                 uint64 `json:"-"`
}

type JenisSKResponse struct {
	Data []JenisSK `json:"data"`
}
