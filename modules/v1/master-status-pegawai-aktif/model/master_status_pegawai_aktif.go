package model

const (
	statusActive = "AKT"
)

type StatusPegawaiAktif struct {
	ID              uint64 `json:"-" gorm:"primaryKey"`
	KdStatusAktif   string `json:"kd_status_aktif" gorm:"column:kd_status"`
	StatusAktif     string `json:"status_aktif" gorm:"column:status"`
	FlagStatusAktif int    `json:"-"`
	UserInput       string `json:"-"`
	UserUpdate      string `json:"-"`
	UUID            string `json:"uuid"`
}

func (s StatusPegawaiAktif) IsActive() bool {
	return s.FlagStatusAktif == 1
}

type StatusPegawaiAktifResponse struct {
	Data []StatusPegawaiAktif `json:"data"`
}
