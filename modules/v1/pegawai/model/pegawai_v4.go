package model

type PegawaiV4 struct {
	Id   uint64
	Nik  string
	Nama string
	Uuid string
}

func (*PegawaiV4) TableName() string {
	return "pegawai_v2"
}
