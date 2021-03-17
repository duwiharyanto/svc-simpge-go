package model

type IndukKerja struct {
	ID         string `json:"-"`
	KdUnit     string `json:"kd_unit"`
	Unit       string `json:"unit"`
	Keterangan string `json:"keterangan"`
	UserInput  string `json:"-"`
	UserUpdate string `json:"-"`
	UUID       string `json:"uuid"`
}

type IndukKerjaResponse struct {
	Data []IndukKerja `json:"data"`
}

type Unit1 struct {
	ID          string `json:"-" gorm:"primaryKey"`
	KdUnit1     string `json:"kd_unit1"`
	Unit1       string `json:"unit1"`
	Keterangan1 string `json:"keterangan1"`
	UserInput   string `json:"-"`
	UserUpdate  string `json:"-"`
	UUID        string `json:"uuid"`
}

type Unit2 struct {
	ID          string `json:"-" gorm:"primaryKey"`
	KdUnit2     string `json:"kd_unit2"`
	Unit2       string `json:"unit2"`
	Keterangan1 string `json:"keterangan1"`
	Keterangan2 string `json:"keterangan2"`
	UserInput   string `json:"-"`
	UserUpdate  string `json:"-"`
	UUID        string `json:"uuid"`
}

type Unit3 struct {
	ID          string `json:"-" gorm:"primaryKey"`
	KdUnit3     string `json:"kd_unit3"`
	Unit3       string `json:"unit3"`
	Keterangan1 string `json:"keterangan1"`
	UserInput   string `json:"-"`
	UserUpdate  string `json:"-"`
	UUID        string `json:"uuid"`
}
