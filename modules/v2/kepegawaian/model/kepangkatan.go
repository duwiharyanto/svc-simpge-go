package model

type PangkatGolonganRuang struct {
	Id         uint64 `form:"-"`
	Pangkat    string `form:"-"`
	IdGolongan uint64 `form:"-"`
	Golongan   string `form:"-"`
	IdRuang    uint64 `form:"-"`
	KdRuang    string `form:"-"`
	FlagAktif  int    `form:"-"`
	UserUpdate string `form:"-"`
	Uuid       string `form:"uuid_pangkat_golongan_pegawai"`
}

func (PangkatGolonganRuang) TableName() string {
	return "pangkat_golongan_pegawai"
}
