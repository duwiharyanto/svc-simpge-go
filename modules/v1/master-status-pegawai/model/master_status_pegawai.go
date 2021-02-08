package model

type StatusPegawai struct {
	KDStatusPegawai string `json:"kd_status_pegawai"`
	StatusPegawai   string `json:"status_pegawai"`
	UUID            string `json:"uuid"`
	ID              string `json:"-"`
}

type StatusPegawaiResponse struct {
	Data []StatusPegawai `json:"data"`
}
