package model

import (
	kepegawaian "svc-insani-go/modules/v2/kepegawaian/model"
)

type SkPegawai struct {
	Id         uint64
	NomorSk    string
	TentangSk  string
	Tmt        string
	FlagAktif  int
	UserUpdate string
	Uuid       string

	IdPegawai uint64
	Pegawai   kepegawaian.Pegawai `gorm:"foreignKey:IdPegawai"`
}

type JenisIjazah struct {
	Id          uint64
	JenisIjazah string
	FlagAktif   int
	UserInput   string
	UserUpdate  string
	Uuid        string
}
