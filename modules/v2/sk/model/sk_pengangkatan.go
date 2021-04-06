package model

import (
	kepegawaian "svc-insani-go/modules/v2/kepegawaian/model"
	organisasi "svc-insani-go/modules/v2/organisasi/model"
)

type KelompokSkPengangkatan struct {
	Id         uint64 `form:"-" json:"-"`
	KelompokSk string `form:"-" json:"kelompok_sk_pengangkatan"`
	FlagAktif  int    `form:"-" json:"-"`
	UserInput  string `form:"-" json:"-"`
	UserUpdate string `form:"-" json:"-"`
	Uuid       string `form:"uuid_kelompok_sk_pengangkatan" json:"uuid_kelompok_sk_pengangkatan"`
}

type StatusPengangkatan struct {
	Id                 uint64 `form:"-" json:"-"`
	StatusPengangkatan string `form:"-" json:"status_pengangkatan"`
	FlagAktif          int    `form:"-" json:"-"`
	UserInput          string `form:"-" json:"-"`
	UserUpdate         string `form:"-" json:"-"`
	Uuid               string `form:"uuid_status_pengangkatan" json:"uuid_status_pengangkatan"`
}

type SkPengangkatanTendik struct {
	Id                   uint64 `form:"-" json:"-"`
	GajiPokok            int    `form:"gaji_pokok" json:"gaji_pokok"`
	MasaKerjaRilBulan    int    `form:"masa_kerja_ril_bulan" json:"masa_kerja_ril_bulan"`
	MasaKerjaRilTahun    int    `form:"masa_kerja_ril_tahun" json:"masa_kerja_ril_tahun"`
	MasaKerjaGajiBulan   int    `form:"masa_kerja_gaji_bulan" json:"masa_kerja_gaji_bulan"`
	MasaKerjaGajiTahun   int    `form:"masa_kerja_gaji_tahun" json:"masa_kerja_gaji_tahun"`
	MasaKerjaDiakuiBulan int    `form:"masa_kerja_diakui_bulan" json:"masa_kerja_diakui_bulan"`
	MasaKerjaDiakuiTahun int    `form:"masa_kerja_diakui_tahun" json:"masa_kerja_diakui_tahun"`
	TanggalDitetapkan    string `form:"tanggal_ditetapkan" json:"tanggal_ditetapkan"`
	PathSk               string `form:"-" json:"-"`
	UrlFileSk            string `form:"-" json:"url_sk_pengangkatan" gorm:"-"`
	FlagAktif            int    `form:"-" json:"-"`
	UserUpdate           string `form:"-" json:"-"`
	Uuid                 string `form:"-" json:"-" query:"uuid_sk_pengangkatan_tendik"`

	IdJabatanFungsional      uint64                           `form:"-" json:"-"`
	JabatanFungsional        kepegawaian.JabatanFungsional    `json:"jabatan_fungsional" gorm:"foreignKey:IdJabatanFungsional"`
	IdStrukorgPejabatPenetap uint64                           `form:"-" json:"-"`
	JabatanPenetap           organisasi.JabatanStruktural     `json:"jabatan_penetap" gorm:"foreignKey:IdStrukorgPejabatPenetap;references:IdStrukturOrganisasi"`
	PejabatPenetap           organisasi.PejabatStruktural     `json:"pejabat_penetap" gorm:"foreignKey:IdStrukorgPejabatPenetap;references:IdStrukturOrganisasi"`
	IdJenisIjazah            uint64                           `form:"-" json:"-"`
	JenisIjazah              JenisIjazah                      `json:"jenis_ijazah" gorm:"foreignKey:IdJenisIjazah"`
	IdKelompokSkPengangkatan uint64                           `form:"-" json:"-"`
	KelompokSkPengangkatan   KelompokSkPengangkatan           `json:"kelompok_sk_pengangkatan" gorm:"foreignKey:IdKelompokSkPengangkatan"`
	IdPangkatGolonganPegawai uint64                           `form:"-" json:"-"`
	PangkatGolonganRuang     kepegawaian.PangkatGolonganRuang `json:"pangkat_golongan_pegawai" gorm:"foreignKey:IdPangkatGolonganPegawai"`
	IdSkPegawai              uint64                           `form:"-" json:"-"`
	SkPegawai                SkPegawai                        `json:"sk_pegawai" gorm:"foreignKey:IdSkPegawai"`
	IdStatusPengangkatan     uint64                           `form:"-" json:"-"`
	StatusPengangkatan       StatusPengangkatan               `json:"status_pengangkatan" gorm:"foreignKey:IdStatusPengangkatan"`
	IdUnitPegawai            uint64                           `form:"-" json:"-"`
	UnitKerja                organisasi.Unit2                 `json:"unit_kerja" gorm:"foreignKey:IdUnitPegawai"`
	IdUnitPengangkat         uint64                           `form:"-" json:"-"`
	UnitPengangkat           organisasi.Unit2                 `json:"unit_pengangkat" gorm:"foreignKey:IdUnitPengangkat"`
}
