package model

type MataKuliah struct {
	KdMakul   string `json:"kd_makul"`
	NamaMakul string `json:"nama_makul"`
	UUID      string `json:"uuid"`
	ID        string `json:"-"`
}

type MataKuliahResponse struct {
	Data []MataKuliah `json:"data"`
}
