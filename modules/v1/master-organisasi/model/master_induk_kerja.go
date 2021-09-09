package model

type IndukKerja struct {
	ID           uint64 `json:"-"`
	KdUnit       string `json:"kd_unit,omitempty"`
	Unit         string `json:"unit,omitempty"`
	Keterangan   string `json:"keterangan"`
	UserInput    string `json:"-"`
	UserUpdate   string `json:"-"`
	UUID         string `json:"uuid"`
	KdIndukKerja string `json:"kd_induk_kerja,omitempty"`
	IndukKerja   string `json:"induk_kerja,omitempty"`
}

type IndukKerjaResponse struct {
	Data []IndukKerja `json:"data"`
}

type UnitKerjaResponse struct {
	Data []Unit2 `json:"data"`
}

type BagianKerjaResponse struct {
	Data []Unit3 `json:"data"`
}

type Unit1 struct {
	ID          uint64 `json:"-" gorm:"primaryKey"`
	KdUnit1     string `json:"kd_unit"`
	Unit1       string `json:"unit"`
	Keterangan1 string `json:"-"`
	UserInput   string `json:"-"`
	UserUpdate  string `json:"-"`
	UUID        string `json:"uuid"`
}

type Unit2 struct {
	ID          uint64 `json:"-" gorm:"primaryKey"`
	KdUnit2     string `json:"kd_unit"`
	Unit2       string `json:"unit"`
	Keterangan2 string `json:"-"`
	KdPddikti   string `json:"kd_pddikti"`
	UserInput   string `json:"-"`
	UserUpdate  string `json:"-"`
	UUID        string `json:"uuid"`

	Keterangan1 string `json:"-"`
	Unit1       Unit1  `json:"-" gorm:"foreignKey:Keterangan1;references:KdUnit1"`
}

type Unit3 struct {
	ID          uint64 `json:"-" gorm:"primaryKey"`
	KdUnit3     string `json:"kd_unit"`
	Unit3       string `json:"unit"`
	Keterangan1 string `json:"-"`
	UserInput   string `json:"-"`
	UserUpdate  string `json:"-"`
	UUID        string `json:"uuid"`
}

type Homebase struct {
	ID          uint64 `json:"-" gorm:"primaryKey"`
	KdUnit2     string `json:"kd_homebase_uii" json:"column:kd_unit2"`
	Unit2       string `json:"unit2"`
	Keterangan1 string `json:"keterangan1"`
	Keterangan2 string `json:"keterangan2"`
	KdPddikti   string `json:"kd_pddikti"`
	UUID        string `json:"uuid"`
}

func (*Homebase) TableName() string {
	return "unit2"
}

type HomebaseResponse struct {
	Data []Homebase `json:"data"`
}
