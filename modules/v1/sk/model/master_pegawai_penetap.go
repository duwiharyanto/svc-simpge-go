package model

type PegawaiPengangkat struct {
	ID                uint64 `json:"-"`
	NIK               string `json:"nik"`
	Nama              string `json:"nama"`
	GelarDepan        string `json:"gelar_depan"`
	GelarBelakang     string `json:"gelar_belakang"`
	KdKelompokPegawai string `json:"kd_kelompok_pegawai"`
	KdUnit            string `json:"kd_unit"`

	UserInput  string `json:"user_input"`
	UserUpdate string `json:"user_update"`
	UUID       string `json:"uuid"`
}

type PegawaiPengangkatResponse struct {
	Data []PegawaiPengangkat `json:"data"`
}
