package model

type JenisPTT struct {
	ID         string `json:"-" gorm:"primaryKey"`
	KdJenisPTT string `json:"kd_jenis_ptt"`
	JenisPTT   string `json:"jenis_ptt"`
	UserInput  string `json:"-"`
	UserUpdate string `json:"-"`
	UUID       string `json:"uuid"`
}

type JenisPTTResponse struct {
	Data []JenisPTT `json:"data"`
}

func (*JenisPTT) TableName() string {
	return "jenis_pegawai_tidak_tetap"
}
