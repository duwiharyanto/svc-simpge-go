package model

type Unit2 struct {
	Id          uint64 `form:"-"`
	RowId       uint64 `form:"-"`
	KdUnit2     string `form:"-"`
	Unit2       string `form:"-"`
	Keterangan1 string `form:"-"`
	FlagAktif   int    `form:"-"`
	UserInput   string `form:"-"`
	UserUpdate  string `form:"-"`
	Uuid        string `form:"-"`
}
