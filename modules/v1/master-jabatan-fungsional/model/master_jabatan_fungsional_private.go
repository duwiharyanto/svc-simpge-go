package model

type JabatanFungsionalPrivate struct {
	JabatanFungsional   string `json:"jabatan_fungsional,omitempty" gorm:"column:fungsional"`
	ID                  string `json:"id_jabatan_fungsional,omitempty" gorm:"primaryKey"`
	KdJabatanFungsional string `json:"kd_jabatan_fungsional,omitempty" gorm:"column:kd_fungsional"`
}
