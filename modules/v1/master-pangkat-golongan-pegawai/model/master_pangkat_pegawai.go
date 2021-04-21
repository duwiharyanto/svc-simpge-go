package model

type PangkatPegawai struct {
	KdPangkat   string `json:"kd_pangkat"`
	NamaPangkat string `json:"nama_pangkat"`
	UUID        string `json:"uuid"`
	ID          string `json:"-"`
}

type PangkatGolonganPegawai struct {
	KdPangkat string `json:"-"`
	Pangkat   string `json:"pangkat"`
	Golongan  string `json:"golongan"`
	UUID      string `json:"uuid"`
	ID        string `json:"-"`
}

type PangkatPegawaiResponse struct {
	Data []PangkatPegawai `json:"data"`
}
type PangkatGolonganPegawaiResponse struct {
	Data []PangkatGolonganPegawai `json:"data"`
}
