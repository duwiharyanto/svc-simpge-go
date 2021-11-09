package model

import (
	"fmt"
	jenisPegawai "svc-insani-go/modules/v1/master-jenis-pegawai/model"
	kelompokPegawai "svc-insani-go/modules/v1/master-kelompok-pegawai/model"
	indukKerja "svc-insani-go/modules/v1/master-organisasi/model"
	statusPegawai "svc-insani-go/modules/v1/master-status-pegawai/model"
	unitKerja "svc-insani-go/modules/v1/master-unit-kerja/model"
	"time"

	ptr "github.com/openlyinc/pointy"
)

type Pegawai struct {
	Id                              uint64 `json:"-" gorm:"primaryKey"`
	IdPersonalDataPribadi           uint64 `json:"-"`
	NIK                             string `json:"nik" gorm:"type:varchar;not null"`
	Nama                            string `json:"nama" gorm:"type:varchar;not null"`
	GelarDepan                      string `json:"gelar_depan" gorm:"type:varchar"`
	GelarBelakang                   string `json:"gelar_belakang" gorm:"type:varchar"`
	FlagDosen                       int    `json:"flag_dosen" gorm:"-"`
	KdUnit2                         string `json:"kd_unit2"`
	UserInput                       string `json:"-"`
	UserUpdate                      string `json:"-"`
	UUID                            string `json:"uuid"`
	jenisPegawai.JenisPegawai       `json:"jenis_pegawai" gorm:"-"`
	kelompokPegawai.KelompokPegawai `json:"kelompok_pegawai" gorm:"-"`
	statusPegawai.StatusPegawai     `json:"status_pegawai" gorm:"-"`
	UnitKerja                       unitKerja.UnitKerja   `json:"unit_kerja" gorm:"foreignKey:KdUnit2"`
	IndukKerja                      indukKerja.IndukKerja `json:"induk_kerja" gorm:"-"`

	PegawaiFungsional PegawaiFungsional `json:"-" gorm:"foreignKey:IdPegawai"`
}

type PegawaiCreate struct {
	Id                      uint64                  `form:"-" gorm:"primaryKey"`
	Uuid                    string                  `form:"-"`
	IdPersonalDataPribadi   uint64                  `form:"-"`
	FlagAktif               int                     `form:"flag_aktif" gorm:"->"`
	Nik                     string                  `form:"nik" gorm:"uniqueIndex"`
	NikKtp                  string                  `form:"nik_ktp" gorm:"default:null"`
	Nama                    string                  `form:"nama"`
	GelarDepan              string                  `form:"gelar_depan" gorm:"default:null"`
	GelarBelakang           string                  `form:"gelar_belakang" gorm:"default:null"`
	TempatLahir             string                  `form:"tempat_lahir" gorm:"default:null"`
	TglLahir                string                  `form:"tgl_lahir" gorm:"default:null"`
	JenisKelamin            string                  `form:"jenis_kelamin" gorm:"default:null"`
	IdAgama                 uint64                  `form:"-" gorm:"default:null"`
	KdAgama                 string                  `form:"kd_agama" gorm:"default:null"`
	IdGolonganDarah         uint64                  `form:"-" gorm:"default:null"`
	KdGolonganDarah         string                  `form:"kd_golongan_darah" gorm:"default:null"`
	IdStatusPerkawinan      uint64                  `form:"-" gorm:"default:null"`
	KdStatusPerkawinan      string                  `form:"kd_status_perkawinan" gorm:"default:null"`
	UuidPendidikanMasuk     string                  `form:"uuid_pendidikan_masuk" gorm:"-"`
	IdPendidikanMasuk       uint64                  `form:"-" gorm:"default:null"`
	KdPendidikanMasuk       string                  `form:"kd_pendidikan_masuk" gorm:"default:null"`
	IdStatusPendidikanMasuk uint64                  `form:"-" gorm:"default:null"`
	KdStatusPendidikanMasuk string                  `form:"kd_status_pendidikan_masuk" gorm:"default:null"`
	UuidPendidikanTerakhir  string                  `form:"uuid_pendidikan_terakhir" gorm:"-"`
	IdPendidikanTerakhir    uint64                  `form:"-" gorm:"default:null"`
	KdPendidikanTerakhir    string                  `form:"kd_pendidikan_terakhir" gorm:"default:null"`
	IdJenisPendidikan       uint64                  `form:"-" gorm:"default:null"`
	KdJenisPendidikan       string                  `form:"kd_jenis_pendidikan" gorm:"default:null"`
	UuidJenisPegawai        string                  `form:"uuid_jenis_pegawai" gorm:"-"`
	IdJenisPegawai          uint64                  `form:"-" gorm:"default:null"`
	KdJenisPegawai          string                  `form:"kd_jenis_pegawai" gorm:"default:null"`
	UuidStatusPegawai       string                  `form:"uuid_status_pegawai" gorm:"-"`
	IdStatusPegawai         uint64                  `form:"-" gorm:"default:null"`
	KdStatusPegawai         string                  `form:"kd_status_pegawai" gorm:"default:null"`
	UuidKelompokPegawai     string                  `form:"uuid_kelompok_pegawai" gorm:"-"`
	IdKelompokPegawai       uint64                  `form:"-" gorm:"default:null"`
	KdKelompokPegawai       string                  `form:"kd_kelompok_pegawai" gorm:"default:null"`
	UuidDetailProfesi       string                  `form:"uuid_detail_profesi"  gorm:"-"`
	IdDetailProfesi         uint64                  `form:"-" gorm:"default:null"`
	UuidGolongan            string                  `form:"uuid_golongan" gorm:"-"`
	IdGolongan              uint64                  `form:"-" gorm:"default:null"`
	KdGolongan              string                  `form:"kd_golongan" gorm:"default:null"`
	UuidRuang               string                  `form:"uuid_ruang" gorm:"-"`
	IdRuang                 uint64                  `form:"-" gorm:"default:null"`
	KdRuang                 string                  `form:"kd_ruang" gorm:"default:null"`
	UuidUnitKerja1          string                  `form:"uuid_induk_kerja" gorm:"-"`
	IdUnitKerja1            uint64                  `form:"-" gorm:"default:null"`
	KdUnit1                 string                  `form:"kd_unit1" gorm:"default:null"`
	UuidUnitKerja2          string                  `form:"uuid_unit_kerja" gorm:"-"`
	IdUnitKerja2            uint64                  `form:"-" gorm:"default:null"`
	KdUnit2                 string                  `form:"kd_unit2" gorm:"default:null"`
	UuidUnitKerja3          string                  `form:"uuid_bagian_kerja" gorm:"-"`
	IdUnitKerja3            uint64                  `form:"-" gorm:"default:null"`
	KdUnit3                 string                  `form:"kd_unit3" gorm:"default:null"`
	IdUnitKerjaLokasi       uint64                  `form:"-" gorm:"default:null"`
	LokasiKerja             string                  `form:"lokasi_kerja" gorm:"default:null"`
	UuidLokasiKerja         string                  `form:"uuid_lokasi_kerja" gorm:"-"`
	FlagPensiun             string                  `form:"flag_pensiun" gorm:"->"`
	TglPensiun              string                  `form:"tgl_pensiun" gorm:"->"`
	FlagMeninggal           string                  `form:"flag_meninggal" gorm:"->"`
	FlagSekolah             string                  `form:"-" gorm:"default:0"`
	FlagMengajar            string                  `form:"-" gorm:"default:0"`
	TglInput                string                  `form:"tgl_input" gorm:"->"`
	UserInput               string                  `form:"user_input"`
	TglUpdate               string                  `form:"tgl_update" gorm:"->"`
	UserUpdate              string                  `form:"user_update"`
	UuidPersonal            string                  `form:"uuid_personal" gorm:"-"`
	PegawaiFungsional       PegawaiFungsionalCreate `gorm:"foreignKey:Id"`
	PegawaiPNS              PegawaiPNSCreate        `gorm:"foreignKey:Id"`

	UuidJenisPresensi string `form:"uuid_jenis_presensi" gorm:"-"`
	KdJenisPresensi   string `form:"kd_jenis_presensi" gorm:"-"`
}

func (*PegawaiCreate) TableName() string {
	return "pegawai"
}

func (p PegawaiCreate) IsLecturer() bool {
	return p.KdJenisPegawai == kdJenisPegawaiDosen
}

type CreatePegawai struct {
	ID   string `json:"-" gorm:"primaryKey"`
	NIK  string `json:"nik" gorm:"type:varchar;not null"`
	Nama string `json:"nama" gorm:"type:varchar;not null"`
	UUID string `json:"uuid"`
}

func (*CreatePegawai) TableName() string {
	return "pegawai"
}

func (p *Pegawai) SetFlagDosen() {
	if !p.JenisPegawai.IsEmpty() && p.JenisPegawai.KDJenisPegawai == "ED" {
		p.FlagDosen = 1
	}
}

type PegawaiRequest struct {
	UuidJenisPegawai    string `query:"uuid_jenis_pegawai"`
	UuidUnitKerja       string `query:"uuid_unit_kerja"`
	UuidKelompokPegawai string `query:"uuid_kelompok_pegawai"`
	UuidStatusAktif     string `query:"uuid_status_aktif"`
	// KdJenisPegawai      string `query:"kd_jenis_pegawai"`
	// KdUnitKerja         string `query:"kd_unit_kerja"`
	// KdKelompokPegawai   string `query:"kd_kelompok_pegawai"`
	Limit  int    `query:"limit"`
	Offset int    `query:"offset"`
	Cari   string `query:"cari"`
}

type PegawaiResponse struct {
	Count  int       `json:"count"`
	Data   []Pegawai `json:"data"`
	Limit  int       `json:"limit"`
	Offset int       `json:"offset"`
}

type PegawaiPribadi struct {
	ID                 uint64 `json:"-"`
	NIK                string `json:"nik"`
	Nama               string `json:"nama"`
	IdAgama            uint64 `json:"-"`
	KdAgama            string `json:"-"`
	KdItemAgama        string `json:"-"`
	IdGolonganDarah    uint64 `json:"-"`
	KdGolonganDarah    string `json:"-"`
	GolonganDarah      string `json:"-"`
	KdKelamin          string `json:"-"`
	IdStatusPerkawinan uint64 `json:"-"`
	KdNikah            string `json:"-"`
	TempatLahir        string `json:"-"`
	TanggalLahir       string `json:"-"`
	FlagPensiun        string `json:"-"`
	GelarDepan         string `json:"gelar_depan"`
	GelarBelakang      string `json:"gelar_belakang"`
	// JumlahAnak       string `json:"-"`
	// JumlahDitanggung string `json:"-"`
	// JumlahKeluarga   string `json:"-"`
	NoKTP string `json:"-"`
	// NoTelepon       string `json:"no_telpon"`
	JenisPegawai            string `json:"jenis_pegawai"`
	KelompokPegawai         string `json:"kelompok_pegawai"`
	UnitKerja               string `json:"unit_kerja"`
	KdStatusPendidikanMasuk string `json:"kd_status_pendidikan_masuk"`
	KdJenisPendidikan       string `json:"kd_jenis_pendidikan"`
	UserInput               string `json:"-"`
	UserUpdate              string `json:"-"`
	UUID                    string `json:"uuid"`
}

type PegawaiDetail struct {
	PegawaiPribadi    *PegawaiPribadi      `json:"pribadi"`
	JenjangPendidikan DataPendidikanDetail `json:"pendidikan"`
	PegawaiYayasan    *PegawaiYayasan      `json:"kepegawaian"`
	UnitKerjaPegawai  *UnitKerjaPegawai    `json:"unit_kerja"`
	PegawaiPNSPTT     *PegawaiPNSPTT       `json:"negara_ptt"`
	StatusAktif       *StatusAktif         `json:"status_aktif"`
}

func (pd PegawaiDetail) IsEmpty() bool {
	return pd.PegawaiPribadi == nil
}

type PegawaiUpdate struct {
	Id                         uint64                  `form:"-" gorm:"primaryKey;<-false"`
	IdPersonalDataPribadi      *uint64                 `form:"-" gorm:"<-:create"`
	FlagAktif                  *int                    `form:"flag_aktif" gorm:"->"`
	Nik                        *string                 `form:"nik" gorm:"->;<-:create"`
	NikKtp                     *string                 `form:"nik_ktp" gorm:"->"`
	Nama                       *string                 `form:"nama" gorm:"->;<-:create"`
	GelarDepan                 *string                 `form:"gelar_depan" gorm:"default:null"`
	GelarBelakang              *string                 `form:"gelar_belakang" gorm:"default:null"`
	TempatLahir                *string                 `form:"tempat_lahir" gorm:"<-:create"`
	TglLahir                   *string                 `form:"tgl_lahir" gorm:"<-:create"`
	JenisKelamin               *string                 `form:"jenis_kelamin" gorm:"<-:create"`
	IdAgama                    *uint64                 `form:"-" gorm:"<-:create"`
	KdAgama                    *string                 `form:"kd_agama" gorm:"<-:create"`
	IdGolonganDarah            *uint64                 `form:"-" gorm:"<-:create"`
	KdGolonganDarah            *string                 `form:"kd_golongan_darah" gorm:"<-:create"`
	IdStatusPerkawinan         *uint64                 `form:"-" gorm:"<-:create"`
	KdStatusPerkawinan         *string                 `form:"kd_status_perkawinan" gorm:"<-:create"`
	UuidPendidikanMasuk        *string                 `form:"uuid_pendidikan_masuk" gorm:"-"` // UUID dari jenjang pendidikan tertinggi diakui
	IdPendidikanMasuk          *uint64                 `form:"-" gorm:"default:null"`
	KdPendidikanMasuk          *string                 `form:"kd_pendidikan_masuk" gorm:"default:null"`
	UuidPendidikanTerakhir     *string                 `form:"uuid_pendidikan_terakhir" gorm:"-"` // UUID dari jenjang pendidikan terakhir
	IdPendidikanTerakhir       *uint64                 `form:"-" gorm:"default:null"`
	KdPendidikanTerakhir       *string                 `form:"kd_pendidikan_terakhir" gorm:"default:null"`
	IdStatusPendidikanMasuk    *uint64                 `form:"-" gorm:"default:null"`
	KdStatusPendidikanMasuk    *string                 `form:"kd_status_pendidikan_masuk" gorm:"default:null"`
	UuidStatusPendidikanMasuk  *string                 `form:"uuid_jenis_pdd_diakui" gorm:"-"`
	IdJenisPendidikan          *uint64                 `form:"-" gorm:"default:null"`
	KdJenisPendidikan          *string                 `form:"kd_jenis_pendidikan" gorm:"default:null"`
	UuidJenisPendidikan        *string                 `form:"uuid_jenis_pdd_terakhir" gorm:"-"`
	UuidJenisPegawai           *string                 `form:"uuid_jenis_pegawai" gorm:"-"`
	IdJenisPegawai             *uint64                 `form:"-" gorm:"default:null"`
	KdJenisPegawai             *string                 `form:"kd_jenis_pegawai" gorm:"default:null"`
	UuidStatusPegawai          *string                 `form:"uuid_status_pegawai" gorm:"-"`
	IdStatusPegawai            *uint64                 `form:"-" gorm:"default:null"`
	KdStatusPegawai            *string                 `form:"kd_status_pegawai" gorm:"default:null"`
	UuidKelompokPegawai        *string                 `form:"uuid_kelompok_pegawai" gorm:"-"`
	IdKelompokPegawai          *uint64                 `form:"-" gorm:"default:null"`
	KdKelompokPegawai          *string                 `form:"kd_kelompok_pegawai" gorm:"default:null"`
	UuidKelompokPegawaiPayroll *string                 `form:"uuid_kelompok_pegawai_payroll" gorm:"-"`
	IdKelompokPegawaiPayroll   *uint64                 `form:"-" gorm:"default:null"`
	KdKelompokPegawaiPayroll   *string                 `form:"kd_kelompok_pegawai_payroll" gorm:"default:null"`
	UuidDetailProfesi          *string                 `form:"uuid_detail_profesi"  gorm:"-"`
	IdDetailProfesi            *uint64                 `form:"-" gorm:"default:null"`
	UuidGolongan               *string                 `form:"uuid_golongan" gorm:"-"`
	KdGolongan                 *string                 `form:"kd_golongan" gorm:"default:null"`
	IdGolongan                 *uint64                 `form:"-" gorm:"default:null"`
	UuidRuang                  *string                 `form:"uuid_ruang" gorm:"-"`
	IdRuang                    *uint64                 `form:"-" gorm:"default:null"`
	KdRuang                    *string                 `form:"kd_ruang" gorm:"default:null"`
	UuidUnitKerja1             *string                 `form:"uuid_induk_kerja" gorm:"-"` //Perubahan
	IdUnitKerja1               *uint64                 `form:"-" gorm:"default:null"`
	KdUnit1                    *string                 `form:"kd_unit1" gorm:"default:null"`
	UuidUnitKerja2             *string                 `form:"uuid_unit_kerja" gorm:"-"` //Perubahan
	IdUnitKerja2               *uint64                 `form:"-" gorm:"default:null"`
	KdUnit2                    *string                 `form:"kd_unit2" gorm:"default:null"`
	UuidUnitKerja3             *string                 `form:"uuid_bagian_kerja" gorm:"-"` //Perubahan
	IdUnitKerja3               *uint64                 `form:"-" gorm:"default:null"`
	KdUnit3                    *string                 `form:"kd_unit3" gorm:"default:null"`
	IdUnitKerjaLokasi          *uint64                 `form:"-" gorm:"default:null"`
	LokasiKerja                *string                 `form:"lokasi_kerja" gorm:"default:null"`
	UuidLokasiKerja            *string                 `form:"uuid_lokasi_kerja" gorm:"-"`
	FlagPensiun                *string                 `form:"flag_pensiun" gorm:"default:null"`
	TglPensiun                 *string                 `form:"tgl_pensiun" gorm:"->"`
	FlagMeninggal              *string                 `form:"flag_meninggal" gorm:"default:null"`
	FlagSekolah                *string                 `form:"-" gorm:"default:0"`
	FlagMengajar               *string                 `form:"-" gorm:"default:0"`
	TglInput                   *string                 `form:"tgl_input" gorm:"->"`
	UserInput                  *string                 `form:"user_input" gorm:"->"`
	TglUpdate                  *string                 `form:"tgl_update" gorm:"->"`
	UserUpdate                 *string                 `form:"user_update"`
	Uuid                       *string                 `form:"uuid" gorm:"->;<-false"`
	PegawaiFungsional          PegawaiFungsionalUpdate `gorm:"foreignkey:IdPegawai;references:Id"`
	PegawaiPNS                 PegawaiPNSUpdate        `gorm:"foreignkey:IdPegawai;references:Id"`
}

func (*PegawaiUpdate) TableName() string {
	return "pegawai"
}

var (
	kdJenisPegawaiDosen = "ED"
)

func (p PegawaiUpdate) IsLecturer() bool {
	return ptr.StringValue(p.KdJenisPegawai, "") == kdJenisPegawaiDosen
}

var indonesianMonths = [...]string{
	"Januari",
	"Februari",
	"Maret",
	"April",
	"Mei",
	"Juni",
	"Juli",
	"Agustus",
	"September",
	"Oktober",
	"November",
	"Desember",
}

func GetIndonesianMonth(date string) string {
	t, _ := time.Parse("2006-01-02", date)
	month := t.Month()
	var idnMonth string
	if time.January <= month && month <= time.December {
		idnMonth = indonesianMonths[month-1]
	}
	return idnMonth
}

func GetIndonesianDate(date string) string {
	dateTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		return ""
	}
	idnMonth := GetIndonesianMonth(date)
	return fmt.Sprintf("%d %s %d", dateTime.Day(), idnMonth, dateTime.Year())
}
