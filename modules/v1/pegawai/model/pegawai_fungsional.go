package model

import statusPegawaiAktif "svc-insani-go/modules/v1/master-status-pegawai-aktif/model"

type PegawaiFungsional struct {
	Id        uint64 `json:"-" gorm:"primaryKey"`
	IdPegawai uint64 `json:"-"`

	IdStatusPegawaiAktif uint64                                `json:"-"`
	StatusPegawaiAktif   statusPegawaiAktif.StatusPegawaiAktif `json:"-" gorm:"foreignKey:IdStatusPegawaiAktif"`
}

type PegawaiFungsionalCreate struct {
	Id                            uint64  `form:"-" gorm:"primaryKey"`
	IdKafka                       uint64  `form:"-" gorm:"-"`
	IdPegawai                     uint64  `form:"-"`
	UuidPangkatGolongan           string  `form:"uuid_pangkat_golongan" gorm:"-"`
	IdPangkatGolongan             uint64  `form:"-" gorm:"default:null"`
	KdPangkatGolongan             string  `form:"kd_pangkat_golongan" gorm:"default:null"`
	UuidJabatanFungsional         string  `form:"uuid_jabatan_fungsional" gorm:"-"`
	IdJabatanFungsional           uint64  `form:"-" gorm:"default:null"`
	KdJabatanFungsional           string  `form:"kd_jabatan_fungsional" gorm:"default:null"`
	TmtPangkatGolongan            *string `form:"tmt_pangkat_golongan" gorm:"default:null"`
	TmtPangkatGolonganIDN         string  `form:"tmt_pangkat_golongan_idn" gorm:"-"`
	TmtJabatan                    *string `form:"tmt_jabatan" gorm:"default:null"`
	TmtJabatanIDN                 string  `form:"tmt_jabatan_idn" gorm:"-"`
	MasaKerjaBawaanTahun          string  `form:"masa_kerja_bawaan_tahun" gorm:"default:null"`
	MasaKerjaBawaanBulan          string  `form:"masa_kerja_bawaan_bulan" gorm:"default:null"`
	MasaKerjaGajiTahun            string  `form:"masa_kerja_gaji_tahun" gorm:"default:null"`
	MasaKerjaGajiBulan            string  `form:"masa_kerja_gaji_bulan" gorm:"default:null"`
	MasaKerjaAwalKepegawaianTahun *string `form:"masa_kerja_awal_kepegawaian_tahun" gorm:"default:null"`
	MasaKerjaAwalKepegawaianBulan *string `form:"masa_kerja_awal_kepegawaian_bulan" gorm:"default:null"`
	MasaKerjaAwalPensiunTahun     *string `form:"masa_kerja_awal_pensiun_tahun" gorm:"default:null"`
	MasaKerjaAwalPensiunBulan     *string `form:"masa_kerja_awal_pensiun_bulan" gorm:"default:null"`
	AngkaKredit                   string  `form:"angka_kredit" gorm:"default:null"`
	NomorSertifikasi              string  `form:"nomor_sertifikasi" gorm:"default:null"`
	UuidJenisNomorRegistrasi      string  `form:"uuid_jenis_nomor_registrasi" gorm:"-"`
	IdJenisNomorRegistrasi        uint64  `form:"-" gorm:"default:null"`
	KdJenisNomorRegistrasi        string  `form:"kd_jenis_nomor_registrasi" gorm:"default:null"`
	NomorRegistrasi               string  `form:"nomor_registrasi" gorm:"default:null"`
	NomorSkPertama                string  `form:"nomor_sk_pertama" gorm:"default:null"`
	TmtSkPertama                  *string `form:"tmt_sk_pertama" gorm:"default:null"`
	TmtSkPertamaIDN               string  `form:"tmt_sk_pertama_idn" gorm:"-"`
	UuidStatusPegawaiAktif        string  `form:"uuid_status_pegawai_aktif" gorm:"-"`
	IdStatusPegawaiAktif          uint64  `form:"-" gorm:"default:null"`
	KdStatusPegawaiAktif          string  `form:"kd_status_pegawai_aktif" gorm:"default:null"`
	UuidHomebasePddikti           string  `form:"uuid_homebase_pddikti" gorm:"-"`
	IdHomebasePddikti             uint64  `form:"-" gorm:"default:null"`
	KdHomebasePddikti             string  `form:"-" gorm:"default:null"`
	UuidHomebaseUii               string  `form:"uuid_homebase_uii" gorm:"-"`
	IdHomebaseUii                 uint64  `form:"-" gorm:"default:null"`
	KdHomebaseUii                 string  `form:"-" gorm:"default:null"`
	NomorSk                       string  `form:"nomor_sk" gorm:"default:null"`
	TmtSk                         string  `form:"tmt_sk_idn" gorm:"-"`
	TmtSkIDN                      string  `form:"tmt_sk_idn" gorm:"default:null"`
	TglSk                         string  `form:"tgl_sk" gorm:"default:null"`
	TglSkIDN                      string  `form:"tgl_sk_idn" gorm:"default:null"`
	TmtAwalKontrak                string  `form:"tmt_awal_kontrak" gorm:"default:null"`
	TmtAwalKontrakIDN             string  `form:"tmt_awal_kontrak_idn" gorm:"default:null"`
	TmtAkhirKontrak               string  `form:"tmt_akhir_kontrak" gorm:"default:null"`
	TmtAkhirKontrakIDN            string  `form:"tmt_akhir_kontrak_idn" gorm:"default:null"`
	TglInput                      string  `form:"-" gorm:"-"`
	UserInput                     string  `form:"-"`
	TglUpdate                     string  `form:"-" gorm:"-"`
	UserUpdate                    string  `form:"-"`
	FlagAktif                     uint64  `form:"-" gorm:"-"`
}

func (*PegawaiFungsionalCreate) TableName() string {
	return "pegawai_fungsional"
}

type PegawaiFungsionalUpdate struct {
	Id                            *uint64 `form:"-"`
	IdKafka                       *uint64 `form:"-"`
	IdPegawai                     *uint64 `form:"-"`
	UuidPangkatGolongan           *string `form:"uuid_pangkat_golongan" gorm:"-"`
	IdPangkatGolongan             *uint64 `-:"id_pangkat_golongan" gorm:"default:null"`
	KdPangkatGolongan             *string `form:"kd_pangkat_golongan" gorm:"default:null"`
	UuidJabatanFungsional         *string `form:"uuid_jabatan_fungsional" gorm:"-"`
	IdJabatanFungsional           *uint64 `form:"-" gorm:"default:null"`
	KdJabatanFungsional           *string `form:"kd_jabatan_fungsional" gorm:"default:null"`
	TmtPangkatGolongan            *string `form:"tmt_pangkat_golongan" gorm:"default:null"`
	TmtPangkatGolonganIDN         *string `form:"tmt_pangkat_golongan_idn" gorm:"-"`
	TmtJabatan                    *string `form:"tmt_jabatan" gorm:"default:null"`
	TmtJabatanIDN                 *string `form:"tmt_jabatan_idn" gorm:"-"`
	MasaKerjaBawaanTahun          *string `form:"masa_kerja_bawaan_tahun" gorm:"default:null"`
	MasaKerjaBawaanBulan          *string `form:"masa_kerja_bawaan_bulan" gorm:"default:null"`
	MasaKerjaGolonganTahun        *string `form:"masa_kerja_golongan_tahun" gorm:"default:null"`
	MasaKerjaGolonganBulan        *string `form:"masa_kerja_golongan_bulan" gorm:"default:null"`
	MasaKerjaGajiTahun            *string `form:"masa_kerja_gaji_tahun" gorm:"default:null"`
	MasaKerjaGajiBulan            *string `form:"masa_kerja_gaji_bulan" gorm:"default:null"`
	MasaKerjaAwalKepegawaianTahun *string `form:"masa_kerja_awal_kepegawaian_tahun" gorm:"default:null"`
	MasaKerjaAwalKepegawaianBulan *string `form:"masa_kerja_awal_kepegawaian_bulan" gorm:"default:null"`
	MasaKerjaAwalPensiunTahun     *string `form:"masa_kerja_awal_pensiun_tahun" gorm:"default:null"`
	MasaKerjaAwalPensiunBulan     *string `form:"masa_kerja_awal_pensiun_bulan" gorm:"default:null"`
	// MasaKerjaTotalTahun      string  `form:"masa_kerja_total_tahun"`
	// MasaKerjaTotalBulan      string  `form:"masa_kerja_total_bulan"`
	AngkaKredit              *string `form:"angka_kredit" gorm:"default:null"`
	NomorSertifikasi         *string `form:"nomor_sertifikasi" gorm:"default:null"`
	Nidn                     *string `form:"nidn" gorm:"default:null"`
	UuidJenisNomorRegistrasi *string `form:"uuid_jenis_nomor_registrasi" gorm:"-"`
	IdJenisNomorRegistrasi   *uint64 `form:"-" gorm:"default:null"`
	KdJenisNomorRegistrasi   *string `form:"kd_jenis_nomor_registrasi" gorm:"default:null"`
	NomorRegistrasi          *string `form:"nomor_registrasi" gorm:"default:null"`
	NomorSkPertama           *string `form:"nomor_sk_pertama" gorm:"default:null"`
	TmtSkPertama             *string `form:"tmt_sk_pertama" gorm:"default:null"`
	TmtSkPertamaIDN          *string `form:"tmt_sk_pertama_idn" gorm:"-"`
	UuidStatusPegawaiAktif   *string `form:"uuid_status_pegawai_aktif" gorm:"-"`
	IdStatusPegawaiAktif     *uint64 `form:"-" gorm:"default:null"`
	KdStatusPegawaiAktif     *string `form:"kd_status_pegawai_aktif" gorm:"default:null"`
	TglStatusPegawaiAktif    *string `form:"tgl_status_aktif" gorm:"default:null"`
	UuidHomebasePddikti      *string `form:"uuid_homebase_pddikti" gorm:"-"`
	IdHomebasePddikti        *uint64 `form:"-" gorm:"default:null"`
	KdHomebasePddikti        *string `form:"-" gorm:"default:null"`
	UuidHomebaseUii          *string `form:"uuid_homebase_uii" gorm:"-"`
	IdHomebaseUii            *uint64 `form:"-" gorm:"default:null"`
	KdHomebaseUii            *string `form:"-" gorm:"default:null"`
	NomorSk                  *string `form:"nomor_sk" gorm:"default:null"`
	TmtSk                    *string `form:"tmt_sk" gorm:"default:null"`
	TmtSkIDN                 *string `form:"tmt_sk_idn" gorm:"-"`
	TglSk                    *string `form:"tgl_sk" gorm:"default:null"`
	TglSkIDN                 *string `form:"tgl_sk_idn" gorm:"-"`
	TmtAwalKontrak           *string `form:"tmt_awal_kontrak" gorm:"default:null"`
	TmtAkhirKontrak          *string `form:"tmt_akhir_kontrak" gorm:"default:null"`
	TglInput                 *string `form:"-" gorm:"-"`
	UserInput                *string `form:"-" gorm:"-"`
	TglUpdate                *string `form:"-" gorm:"-"`
	UserUpdate               *string `form:"-"`
	FlagAktif                *uint64 `form:"-" gorm:"-"`
	// Uuid                     string `form:"-"`
}

func (*PegawaiFungsionalUpdate) TableName() string {
	return "pegawai_fungsional"
}
