package model

type JabatanFungsional struct {
	Id             uint64 `form:"-" json:"-"`
	KdFungsional   string `form:"-" json:"kd_jabatan_fungsional"`
	KdJenisPegawai string `form:"-" json:"-"`
	Fungsional     string `form:"-" json:"jabatan_fungsional"`
	FlagAktif      int    `form:"-" json:"-"`
	UserInput      string `form:"-" json:"-"`
	UserUpdate     string `form:"-" json:"-"`
	Uuid           string `form:"uuid_jabatan_fungsional" json:"uuid_jabatan_fungsional"`
}
