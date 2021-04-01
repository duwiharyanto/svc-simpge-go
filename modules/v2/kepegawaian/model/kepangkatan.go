package model

type PangkatGolonganRuang struct {
	Id         uint64
	Pangkat    string
	IdGolongan uint64
	Golongan   string
	IdRuang    uint64
	KdRuang    string
	FlagAktif  int
	UserUpdate string
	Uuid       string
}

func (PangkatGolonganRuang) TableName() string {
	return "pangkat_golongan_pegawai"
}
