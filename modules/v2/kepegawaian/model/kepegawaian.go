package model

import (
	organisasi "svc-insani-go/modules/v2/organisasi/model"
)

type Pegawai struct {
	Id        uint64 `form:"-" json:"-"`
	Nik       string `form:"-" json:"nik"`
	Nama      string `form:"-" json:"nama"`
	FlagAktif int    `form:"-" json:"-"`
	Uuid      string `form:"-" json:"uuid"`

	IdJenisPegawai uint64       `form:"-" json:"-"`
	JenisPegawai   JenisPegawai `form:"-" json:"jenis_pegawai" gorm:"foreignKey:IdJenisPegawai"`

	IdUnitKerja2 uint64           `form:"-" json:"-"`
	Unit2        organisasi.Unit2 `form:"-" json:"unit_kerja" gorm:"foreignKey:IdUnitKerja2"`
}

type JenisPegawai struct {
	Id               uint64 `form:"-" json:"-"`
	KdJenisPegawai   string `form:"-" json:"kd_jenis_pegawai"`
	NamaJenisPegawai string `form:"-" json:"jenis_pegawai"`
	FlagAktif        int    `form:"-" json:"-"`
	Uuid             string `form:"-" json:"uuid"`
}
