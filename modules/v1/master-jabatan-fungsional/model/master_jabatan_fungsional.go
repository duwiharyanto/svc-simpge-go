package model

type JabatanFungsional struct {
	KdJabatanFungsional string `json:"kd_jabatan_fungsional" gorm:"column:kd_fungsional"`
	JabatanFungsional   string `json:"jabatan_fungsional" gorm:"column:fungsional"`
	UUID                string `json:"uuid"`
	ID                  string `json:"-" gorm:"primaryKey"`
}

type JabatanFungsionalResponse struct {
	Data []JabatanFungsional `json:"data"`
}
