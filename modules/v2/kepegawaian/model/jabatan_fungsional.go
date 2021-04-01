package model

type JabatanFungsional struct {
	Id             uint64 `form:"-"`
	KdFungsional   string `form:"-"`
	KdJenisPegawai string `form:"-"`
	Fungsional     string `form:"-"`
	FlagAktif      int    `form:"-"`
	UserInput      string `form:"-"`
	UserUpdate     string `form:"-"`
	Uuid           string `form:"uuid_jabatan_fungsional"`
}
