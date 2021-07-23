package model

type StatusPegawaiAktif struct {
	ID            uint64 `json:"-" gorm:"primaryKey"`
	KdStatusAktif string `json:"kd_status_aktif" gorm:"column:kd_status"`
	StatusAktif   string `json:"status_aktif" gorm:"column:status"`
	UserInput     string `json:"-"`
	UserUpdate    string `json:"-"`
	UUID          string `json:"uuid"`
}

type StatusPegawaiAktifResponse struct {
	Data []StatusPegawaiAktif `json:"data"`
}
