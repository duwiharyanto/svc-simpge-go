package model

type PangkatPegawai struct {
	KdPangkat   string `json:"kd_pangkat"`
	NamaPangkat string `json:"nama_pangkat"`
	UUID        string `json:"uuid"`
	ID          uint64 `json:"-"`
}

type PangkatGolonganPegawai struct {
	ID         uint64 `json:"-"`
	KdPangkat  string `json:"-"`
	Pangkat    string `json:"pangkat"`
	Golongan   string `json:"golongan"`
	IdGolongan uint64 `json:"-"`
	KdGolongan string `json:"-"`
	IdRuang    uint64 `json:"-"`
	KdRuang    string `json:"-"`
	UUID       string `json:"uuid"`
}

type PangkatPegawaiResponse struct {
	Data []PangkatPegawai `json:"data"`
}
type PangkatGolonganPegawaiResponse struct {
	Data []PangkatGolonganPegawai `json:"data"`
}
