package model

import (
	kepegawaian "svc-insani-go/modules/v2/kepegawaian/model"
	organisasi "svc-insani-go/modules/v2/organisasi/model"
)

type KelompokSkPengangkatan struct {
	Id         uint64
	IdKafka    uint64
	KelompokSk string
	FlagAktif  int
	UserInput  string
	UserUpdate string
	Uuid       string
}

type StatusPengangkatan struct {
	Id                 uint64
	StatusPengangkatan string
	FlagAktif          int
	UserInput          string
	UserUpdate         string
	Uuid               string
}

type SkPengangkatanTendik struct {
	Id                   uint64
	GajiPokok            int
	MasaKerjaRilBulan    int
	MasaKerjaRilTahun    int
	MasaKerjaGajiBulan   int
	MasaKerjaGajiTahun   int
	MasaKerjaDiakuiBulan int
	MasaKerjaDiakuiTahun int
	TanggalDitetapkan    string
	PathSk               string
	FlagAktif            int
	UserUpdate           string
	Uuid                 string

	IdJabatanFungsional      uint64
	JabatanFungsional        kepegawaian.JabatanFungsional `gorm:"foreignKey:IdJabatanFungsional"`
	IdJenisIjazah            uint64
	JenisIjazah              JenisIjazah `gorm:"foreignKey:IdJenisIjazah"`
	IdKelompokSkPengangkatan uint64
	KelompokSkPengangkatan   KelompokSkPengangkatan `gorm:"foreignKey:IdKelompokSkPengangkatan"`
	IdPangkatGolonganPegawai uint64
	PangkatGolonganRuang     kepegawaian.PangkatGolonganRuang `gorm:"foreignKey:IdPangkatGolonganPegawai"`
	IdSkPegawai              uint64
	SkPegawai                SkPegawai `gorm:"foreignKey:IdSkPegawai"`
	IdStatusPengangkatan     uint64
	StatusPengangkatan       StatusPengangkatan `gorm:"foreignKey:IdStatusPengangkatan"`
	IdUnitPegawai            uint64
	UnitKerja                organisasi.Unit2 `gorm:"foreignKey:IdUnitPegawai"`
	IdUnitPengangkat         uint64
	UnitPengangkat           organisasi.Unit2 `gorm:"foreignKey:IdUnitPengangkat"`
}
