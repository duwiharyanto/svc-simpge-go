package model

type StatusPegawaiAktif struct {
	ID            string `json:"-"`
	KdStatusAktif string `json:"kd_status_aktif"`
	StatusAktif   string `json:"status_aktif"`
	UserInput     string `json:"-"`
	UserUpdate    string `json:"-"`
	UUID          string `json:"uuid"`
}

type StatusPegawaiAktifResponse struct {
	Data []StatusPegawaiAktif `json:"data"`
}
