package model

type StatusPegawai struct {
	KDStatusPegawai string `json:"kd_status_pegawai"`
	StatusPegawai   string `json:"status_pegawai"`
	UUID            string `json:"uuid"`
	ID              uint64 `json:"-" gorm:"primaryKey"`
}

type StatusPegawaiResponse struct {
	Data []StatusPegawai `json:"data"`
}
