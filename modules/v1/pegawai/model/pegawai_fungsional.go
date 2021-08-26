package model

import statusPegawaiAktif "svc-insani-go/modules/v1/master-status-pegawai-aktif/model"

type PegawaiFungsional struct {
	Id        uint64 `json:"-" gorm:"primaryKey"`
	IdPegawai uint64 `json:"-"`

	IdStatusPegawaiAktif uint64                                `json:"-"`
	StatusPegawaiAktif   statusPegawaiAktif.StatusPegawaiAktif `json:"-" gorm:"foreignKey:IdStatusPegawaiAktif"`
}

type PegawaiFungsionalCreate struct {
	Id                       uint64  `form:"-" gorm:"primaryKey"`
	IdKafka                  uint64  `form:"-"`
	IdPegawai                uint64  `form:"-"`
	UuidPangkatGolongan      string  `form:"uuid_pangkat_golongan" gorm:"-"`
	IdPangkatGolongan        uint64  `form:"id_pangkat_golongan" gorm:"default:null"`
	KdPangkatGolongan        string  `form:"kd_pangkat_golongan" gorm:"default:null"`
	UuidJabatanFungsional    string  `form:"uuid_jabatan_fungsional" gorm:"-"`
	IdJabatanFungsional      uint64  `form:"id_jabatan_fungsional" gorm:"default:null"`
	KdJabatanFungsional      string  `form:"kd_jabatan_fungsional" gorm:"default:null"`
	TmtPangkatGolongan       *string `form:"tmt_pangkat_golongan" gorm:"default:null"`
	TmtPangkatGolonganIDN    string  `form:"tmt_pangkat_golongan_idn" gorm:"-"`
	TmtJabatan               *string `form:"tmt_jabatan" gorm:"default:null"`
	TmtJabatanIDN            string  `form:"tmt_jabatan_idn" gorm:"-"`
	MasaKerjaBawaanTahun     string  `form:"masa_kerja_bawaan_tahun" gorm:"default:null"`
	MasaKerjaBawaanBulan     string  `form:"masa_kerja_bawaan_bulan" gorm:"default:null"`
	MasaKerjaGajiTahun       string  `form:"masa_kerja_gaji_tahun" gorm:"default:null"`
	MasaKerjaGajiBulan       string  `form:"masa_kerja_gaji_bulan" gorm:"default:null"`
	AngkaKredit              string  `form:"angka_kredit" gorm:"default:null"`
	NomorSertifikasi         string  `form:"nomor_sertifikasi" gorm:"default:null"`
	UuidJenisNomorRegistrasi string  `form:"uuid_jenis_nomor_registrasi" gorm:"-"`
	IdJenisNomorRegistrasi   uint64  `form:"id_jenis_nomor_registrasi" gorm:"default:null"`
	KdJenisNomorRegistrasi   string  `form:"kd_jenis_nomor_registrasi" gorm:"default:null"`
	NomorRegistrasi          string  `form:"nomor_registrasi" gorm:"default:null"`
	NomorSkPertama           string  `form:"nomor_sk_pertama" gorm:"default:null"`
	TmtSkPertama             *string `form:"tmt_sk_pertama" gorm:"default:null"`
	TmtSkPertamaIDN          string  `form:"tmt_sk_pertama_idn" gorm:"-"`
	UuidStatusPegawaiAktif   string  `form:"uuid_status_pegawai_aktif" gorm:"-"`
	IdStatusPegawaiAktif     uint64  `form:"id_status_pegawai_aktif" gorm:"default:null"`
	KdStatusPegawaiAktif     string  `form:"kd_status_pegawai_aktif" gorm:"default:null"`
	UuidHomebasePddikti      string  `form:"uuid_homebase_pddikti" gorm:"-"`
	IdHomebasePddikti        uint64  `form:"id_homebase_pddikti" gorm:"default:null"`
	UuidHomebaseUii          string  `form:"uuid_homebase_uii" gorm:"-"`
	IdHomebaseUii            uint64  `form:"id_homebase_uii" gorm:"default:null"`
	TglInput                 string  `form:"-" gorm:"-"`
	UserInput                string  `form:"-"`
	TglUpdate                string  `form:"-" gorm:"-"`
	UserUpdate               string  `form:"-"`
	FlagAktif                uint64  `form:"-" gorm:"-"`
}

func (*PegawaiFungsionalCreate) TableName() string {
	return "pegawai_fungsional"
}

type PegawaiFungsionalUpdate struct {
	Id                    uint64  `form:"-"`
	IdKafka               uint64  `form:"-"`
	IdPegawai             uint64  `form:"-"`
	UuidPangkatGolongan   string  `form:"uuid_pangkat_golongan" gorm:"-"`
	IdPangkatGolongan     uint64  `form:"id_pangkat_golongan"`
	KdPangkatGolongan     string  `form:"kd_pangkat_golongan"`
	UuidJabatanFungsional string  `form:"uuid_jabatan_fungsional" gorm:"-"`
	IdJabatanFungsional   uint64  `form:"id_jabatan_fungsional"`
	KdJabatanFungsional   string  `form:"kd_jabatan_fungsional"`
	TmtPangkatGolongan    *string `form:"tmt_pangkat_golongan" gorm:"default:null"`
	TmtPangkatGolonganIDN string  `form:"tmt_pangkat_golongan_idn" gorm:"-"`
	TmtJabatan            *string `form:"tmt_jabatan"`
	TmtJabatanIDN         string  `form:"tmt_jabatan_idn" gorm:"-"`
	MasaKerjaBawaanTahun  string  `form:"masa_kerja_bawaan_tahun"`
	MasaKerjaBawaanBulan  string  `form:"masa_kerja_bawaan_bulan"`
	MasaKerjaGajiTahun    string  `form:"masa_kerja_gaji_tahun"`
	MasaKerjaGajiBulan    string  `form:"masa_kerja_gaji_bulan"`
	// MasaKerjaTotalTahun      string  `form:"masa_kerja_total_tahun"`
	// MasaKerjaTotalBulan      string  `form:"masa_kerja_total_bulan"`
	AngkaKredit              string  `form:"angka_kredit"`
	NomorSertifikasi         string  `form:"nomor_sertifikasi"`
	UuidJenisNomorRegistrasi string  `form:"uuid_jenis_nomor_registrasi" gorm:"-"`
	IdJenisNomorRegistrasi   uint64  `form:"id_jenis_nomor_registrasi"`
	KdJenisNomorRegistrasi   string  `form:"kd_jenis_nomor_registrasi"`
	NomorRegistrasi          string  `form:"nomor_registrasi"`
	NomorSkPertama           string  `form:"nomor_sk_pertama"`
	TmtSkPertama             *string `form:"tmt_sk_pertama"`
	TmtSkPertamaIDN          string  `form:"tmt_sk_pertama_idn" gorm:"-"`
	UuidStatusPegawaiAktif   string  `form:"uuid_status_pegawai_aktif" gorm:"-"`
	IdStatusPegawaiAktif     uint64  `form:"id_status_pegawai_aktif"`
	KdStatusPegawaiAktif     string  `form:"kd_status_pegawai_aktif"`
	UuidHomebasePddikti      string  `form:"uuid_homebase_pddikti" gorm:"-"` //Perubahan
	IdHomebasePddikti        uint64  `form:"id_homebase_pddikti"`            //Perubahan
	UuidHomebaseUii          string  `form:"uuid_homebase_uii" gorm:"-"`     //Perubahan
	IdHomebaseUii            uint64  `form:"id_homebase_uii"`                //Perubahan
	TglInput                 string  `form:"-" gorm:"-"`
	UserInput                string  `form:"-" gorm:"-"`
	TglUpdate                string  `form:"-" gorm:"-"`
	UserUpdate               string  `form:"-"`
	FlagAktif                uint64  `form:"-" gorm:"-"`
	// Uuid                     string `form:"-"`
}

func (*PegawaiFungsionalUpdate) TableName() string {
	return "pegawai_fungsional"
}
