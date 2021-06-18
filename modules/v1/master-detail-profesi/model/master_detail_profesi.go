package model

type DetailProfesi struct {
	ID            string `json:"-" gorm:"primaryKey"`
	DetailProfesi string `json:"nama_jenjang"`
	UserInput     string `json:"-"`
	UserUpdate    string `json:"-"`
	UUID          string `json:"uuid"`
}

type DetailProfesiResponse struct {
	Data []DetailProfesi `json:"data"`
}
