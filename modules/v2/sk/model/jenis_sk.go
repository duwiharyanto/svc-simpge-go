package model

type JenisSk struct {
	Id        uint64 `json:"-"`
	KdJenisSk string `json:"kd_jenis_sk"`
	NamaSk    string `json:"jenis_sk"`
	FlagAktif int    `json:"-"`
	Uuid      string `json:"uuid"`
}
