package model

import (
	kepegawaian "svc-insani-go/modules/v2/kepegawaian/model"
	organisasi "svc-insani-go/modules/v2/organisasi/model"
)

type KelompokSkPengangkatan struct {
	Id         uint64 `form:"-"`
	IdKafka    uint64 `form:"-"`
	KelompokSk string `form:"-"`
	FlagAktif  int    `form:"-"`
	UserInput  string `form:"-"`
	UserUpdate string `form:"-"`
	Uuid       string `form:"uuid_kelompok_sk_pengangkatan"`
}

type StatusPengangkatan struct {
	Id                 uint64 `form:"-"`
	StatusPengangkatan string `form:"-"`
	FlagAktif          int    `form:"-"`
	UserInput          string `form:"-"`
	UserUpdate         string `form:"-"`
	Uuid               string `form:"uuid_status_pengangkatan"`
}

type SkPengangkatanTendik struct {
	Id                   uint64 `form:"-"`
	GajiPokok            int    `form:"gaji_pokok"`
	MasaKerjaRilBulan    int    `form:"masa_kerja_ril_bulan"`
	MasaKerjaRilTahun    int    `form:"masa_kerja_ril_tahun"`
	MasaKerjaGajiBulan   int    `form:"masa_kerja_gaji_bulan"`
	MasaKerjaGajiTahun   int    `form:"masa_kerja_gaji_tahun"`
	MasaKerjaDiakuiBulan int    `form:"masa_kerja_diakui_bulan"`
	MasaKerjaDiakuiTahun int    `form:"masa_kerja_diakui_tahun"`
	TanggalDitetapkan    string `form:"tanggal_ditetapkan"`
	PathSk               string `form:"-"`
	FlagAktif            int    `form:"-"`
	UserUpdate           string `form:"-"`
	Uuid                 string `form:"-" query:"uuid_sk_pengangkatan_tendik"`

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
