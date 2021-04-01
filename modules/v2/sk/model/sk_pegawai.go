package model

import (
	kepegawaian "svc-insani-go/modules/v2/kepegawaian/model"
)

type SkPegawai struct {
	Id         uint64 `form:"-"`
	NomorSk    string `form:"nomor_sk"`
	TentangSk  string `form:"tentang_sk"`
	Tmt        string `form:"tmt"`
	FlagAktif  int    `form:"-"`
	UserUpdate string `form:"-"`
	Uuid       string `form:"-"`

	IdPegawai uint64
	Pegawai   kepegawaian.Pegawai `gorm:"foreignKey:IdPegawai"`
}

type JenisIjazah struct {
	Id          uint64 `form:"-"`
	JenisIjazah string `form:"-"`
	FlagAktif   int    `form:"-"`
	UserInput   string `form:"-"`
	UserUpdate  string `form:"-"`
	Uuid        string `form:"uuid_jenis_ijazah"`
}
