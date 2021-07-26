package model

type DetailProfesi struct {
	ID            uint64 `json:"-" gorm:"primaryKey"`
	DetailProfesi string `json:"nama_jenjang"`
	UserInput     string `json:"-"`
	UserUpdate    string `json:"-"`
	UUID          string `json:"uuid"`
}

type DetailProfesiResponse struct {
	Data []DetailProfesi `json:"data"`
}
