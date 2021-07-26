package model

type JenisNomorRegistrasi struct {
	ID              uint64 `json:"-" gorm:"primaryKey"`
	JenisNomorRegis string `json:"jenis_no_regis" gorm:"column:jenis_no_regis"`
	KdJenisRegis    string `json:"kd_jenis_regis"`
	UserInput       string `json:"-"`
	UserUpdate      string `json:"-"`
	UUID            string `json:"uuid"`
}

func (*JenisNomorRegistrasi) TableName() string {
	return "jenis_nomor_registrasi"
}

type JenisNomorRegistrasiResponse struct {
	Data []JenisNomorRegistrasi `json:"data"`
}
