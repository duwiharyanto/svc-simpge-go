package model

type Pegawai struct {
	Id        uint64 `form:"-" json:"-"`
	Nik       string `form:"-" json:"nik_pegawai"`
	Nama      string `form:"-" json:"nama_pegawai"`
	FlagAktif int    `form:"-" json:"-"`
	Uuid      string `form:"-" json:"-"`
}
