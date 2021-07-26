package model

type JenjangPendidikan struct {
	ID                 uint64 `json:"-" gorm:"primaryKey"`
	KdJenjang          string `json:"kd_jenjang"`
	KdPendidikanSimpeg string `json:"-"`
	Jenjang            string `json:"jenjang"`
	NamaJenjang        string `json:"nama_jenjang"`
	UserInput          string `json:"-"`
	UserUpdate         string `json:"-"`
	UUID               string `json:"uuid"`
}

type JenjangPendidikanResponse struct {
	Data []JenjangPendidikan `json:"data"`
}
