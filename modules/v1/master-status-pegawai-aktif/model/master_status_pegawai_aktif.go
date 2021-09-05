package model

const (
	statusActive               = "AKT"
	statusRetired              = "PEN"
	statusDied                 = "MNG"
	statusOnStudyingPermission = "IBL"
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

func (s StatusPegawaiAktif) IsRetired() bool {
	return s.KdStatusAktif == statusRetired
}

func (s StatusPegawaiAktif) IsDied() bool {
	return s.KdStatusAktif == statusDied
}

func (s StatusPegawaiAktif) IsOnStudyingPermission() bool {
	return s.KdStatusAktif == statusOnStudyingPermission
}

type StatusPegawaiAktifResponse struct {
	Data []StatusPegawaiAktif `json:"data"`
}
