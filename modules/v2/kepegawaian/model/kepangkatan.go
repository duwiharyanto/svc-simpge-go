package model

type PangkatGolonganRuang struct {
	Id         uint64 `form:"-" json:"-"`
	Pangkat    string `form:"-" json:"pangkat"`
	IdGolongan uint64 `form:"-" json:"-"`
	Golongan   string `form:"-" json:"golongan"`
	IdRuang    uint64 `form:"-" json:"-"`
	KdRuang    string `form:"-" json:"ruang"`
	FlagAktif  int    `form:"-" json:"-"`
	UserUpdate string `form:"-" json:"-"`
	Uuid       string `form:"uuid_pangkat_golongan_pegawai" json:"uuid_pangkat_golongan_pegawai"`
}

func (PangkatGolonganRuang) TableName() string {
	return "pangkat_golongan_pegawai"
}
