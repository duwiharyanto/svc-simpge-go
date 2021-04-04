package model

import (
	kepegawaian "svc-insani-go/modules/v2/kepegawaian/model"
)

type SkPegawai struct {
	Id         uint64 `form:"-" json:"-"`
	NomorSk    string `form:"nomor_sk" json:"nomor_sk"`
	TentangSk  string `form:"tentang_sk" json:"tentang_sk"`
	Tmt        string `form:"tmt" json:"tmt"`
	FlagAktif  int    `form:"-" json:"-"`
	UserUpdate string `form:"-" json:"-"`
	Uuid       string `form:"-" json:"-"`

	IdPegawai uint64              `form:"-" json:"-"`
	Pegawai   kepegawaian.Pegawai `json:"pegawai" gorm:"foreignKey:IdPegawai"`
}

type JenisIjazah struct {
	Id          uint64 `form:"-"`
	JenisIjazah string `form:"jenis_ijazah"`
	FlagAktif   int    `form:"-"`
	UserInput   string `form:"-"`
	UserUpdate  string `form:"-"`
	Uuid        string `form:"uuid_jenis_ijazah"`
}
