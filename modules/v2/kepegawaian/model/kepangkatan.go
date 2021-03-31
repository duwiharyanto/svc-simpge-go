package model

type PangkatGolonganRuang struct {
	Id        uint64
	Pangkat   string
	Golongan  string
	FlagAktif int
	Uuid      string
}

func (PangkatGolonganRuang) TableName() string {
	return "pangkat_golongan_pegawai"
}
