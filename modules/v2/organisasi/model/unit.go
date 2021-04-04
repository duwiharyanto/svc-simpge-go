package model

type Unit2 struct {
	Id          uint64 `form:"-" json:"-"`
	RowId       uint64 `form:"-" json:"-"`
	KdUnit2     string `form:"-" json:"kd_unit"`
	Unit2       string `form:"-" json:"unit"`
	Keterangan1 string `form:"-" json:"kd_induk_unit"`
	FlagAktif   int    `form:"-" json:"-"`
	UserInput   string `form:"-" json:"-"`
	UserUpdate  string `form:"-" json:"-"`
	Uuid        string `form:"-" json:"uuid"`
}
