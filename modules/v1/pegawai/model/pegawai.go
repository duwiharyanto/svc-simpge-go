package model

import (
	"fmt"
	"strings"
	jenisPegawai "svc-insani-go/modules/v1/master-jenis-pegawai/model"
	kelompokPegawai "svc-insani-go/modules/v1/master-kelompok-pegawai/model"
	indukKerja "svc-insani-go/modules/v1/master-organisasi/model"

	statusPegawaiAktif "svc-insani-go/modules/v1/master-status-pegawai-aktif/model"
	statusPegawai "svc-insani-go/modules/v1/master-status-pegawai/model"
	unitKerja "svc-insani-go/modules/v1/master-unit-kerja/model"
	"time"

	"github.com/cstockton/go-conv"
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

type CreatePegawai struct {
	ID   string `json:"-" gorm:"primaryKey"`
	NIK  string `json:"nik" gorm:"type:varchar;not null"`
	Nama string `json:"nama" gorm:"type:varchar;not null"`
	UUID string `json:"uuid"`
}

func (*CreatePegawai) TableName() string {
	return "pegawai"
}

type PegawaiRequestParam struct {
	ID                              string `json:"-"`
	NIK                             string `json:"nik"`
	jenisPegawai.JenisPegawai       `json:"jenis_pegawai"`
	kelompokPegawai.KelompokPegawai `json:"kelompok_pegawai"`
	indukKerja.IndukKerja           `json:"induk_kerja"`
	Limit                           int
	Offset                          int
}

func (p *Pegawai) SetFlagDosen() {
	if !p.JenisPegawai.IsEmpty() && p.JenisPegawai.KDJenisPegawai == "ED" {
		p.FlagDosen = 1
	}
}

type PegawaiFungsional struct {
	Id        uint64 `json:"-" gorm:"primaryKey"`
	IdPegawai uint64 `json:"-"`

	IdStatusPegawaiAktif uint64                                `json:"-"`
	StatusPegawaiAktif   statusPegawaiAktif.StatusPegawaiAktif `json:"-" gorm:"foreignKey:IdStatusPegawaiAktif"`
}

type Pegawai2 struct {
	ID            string `gorm:"primaryKey;not null"`
	NIK           string `gorm:"type:varchar;not null"`
	Nama          string `gorm:"type:varchar;not null"`
	GelarDepan    string `gorm:"type:varchar"`
	GelarBelakang string `gorm:"type:varchar"`
	KdUnit2       int    `gorm:"type:varchar; column:kd_unit2"`
	Unit2         *Unit2 `gorm:"foreignKey:KdUnit2"`
	UserInput     string `gorm:"type:varchar"`
	UserUpdate    string `gorm:"type:varchar"`
	UUID          string `gorm:"type:varchar"`
}

func (*Pegawai2) TableName() string {
	return "pegawai"
}

type Unit2 struct {
	ID            string `gorm:"primaryKey;not null"`
	KdUnit2       string `gorm:"type:varchar;unique;column:kd_unit2"`
	NamaUnitKerja string `gorm:"type:varchar;column:unit2"`
	UUID          string `gorm:"type:varchar"`
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

type PegawaiResponseTest struct {
	Count  int        `json:"count"`
	Data   []Pegawai2 `json:"data"`
	Limit  int        `json:"limit"`
	Offset int        `json:"offset"`
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
	JenisPegawai    string `json:"jenis_pegawai"`
	KelompokPegawai string `json:"kelompok_pegawai"`
	UnitKerja       string `json:"unit_kerja"`
	UserInput       string `json:"-"`
	UserUpdate      string `json:"-"`
	UUID            string `json:"uuid"`
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

type PegawaiYayasan struct {
	ID                         uint64 `json:"-" gorm:"primaryKey;not null"`
	IDJenisPegawai             uint64 `json:"-"`
	KDJenisPegawai             string `json:"kd_jenis_pegawai"`
	UuidJenisPegawai           string `json:"uuid_jenis_pegawai"`
	JenisPegawai               string `json:"jenis_pegawai"`
	UuidKelompokPegawai        string `json:"uuid_kelompok_pegawai"`
	IdKelompokPegawai          uint64 `json:"-"`
	KdKelompokPegawai          string `json:"kd_kelompok_pegawai"`
	KelompokPegawai            string `json:"kelompok_pegawai"`
	UuidPendidikanMasuk        string `json:"-" form:"uuid_pendidikan_masuk"`
	IdPendidikanMasuk          uint64 `json:"-"`
	KdPendidikanMasuk          string `json:"-"`
	IdPendidikanMasukSimpeg    uint64 `json:"-"`
	KdPendidikanMasukSimpeg    string `json:"-"`
	PendidikanMasuk            string `json:"-"`
	UuidPendidikanTerakhir     string `json:"-" form:"uuid_pendidikan_terakhir"`
	IdPendidikanTerakhir       uint64 `json:"-"`
	KdPendidikanTerakhir       string `json:"-"`
	IdPendidikanTerakhirSimpeg uint64 `json:"-"`
	KdPendidikanTerakhirSimpeg string `json:"-"`
	PendidikanTerakhir         string `json:"-"`
	UuidStatusPegawai          string `json:"uuid_status_pegawai"`
	IDStatusPegawai            uint64 `json:"-"`
	KDStatusPegawai            string `json:"kd_status_pegawai"`
	StatusPegawai              string `json:"status_pegawai"`
	UuidPangkatGolongan        string `json:"uuid_pangkat_golongan"`
	IdPangkat                  uint64 `json:"-"`
	KdPangkat                  string `json:"kd_pangkat_golongan"`
	Pangkat                    string `json:"pangkat"`
	Golongan                   string `json:"golongan"`
	IdGolongan                 uint64 `json:"-"`
	KdGolongan                 string `json:"kd_golongan"`
	IdRuang                    uint64 `json:"-"`
	KdRuang                    string `json:"kd_ruang"`
	TmtPangkatGolongan         string `json:"tmt_pangkat_gol_ruang_pegawai"`
	TmtPangkatGolonganIdn      string `json:"tmt_pangkat_gol_ruang_pegawai_idn"`
	UuidJabatanFungsional      string `json:"uuid_jabatan_fungsional"`
	IdJabatanFungsional        uint64 `json:"-"`
	KdJabatanFungsional        string `json:"kd_jabatan_fungsional"`
	JabatanFungsional          string `json:"jabatan_fungsional"`
	TmtJabatan                 string `json:"tmt_jabatan"`
	TmtJabatanIdn              string `json:"tmt_jabatan_idn"`
	MasaKerjaBawaanTahun       string `json:"masa_kerja_bawaan_tahun"`
	MasaKerjaBawaanBulan       string `json:"masa_kerja_bawaan_bulan"`
	MasaKerjaGajiTahun         string `json:"masa_kerja_gaji_tahun"`
	MasaKerjaGajiBulan         string `json:"masa_kerja_gaji_bulan"`
	MasaKerjaTotalTahun        string `json:"masa_kerja_total_tahun"`
	MasaKerjaTotalBulan        string `json:"masa_kerja_total_bulan"`
	AngkaKredit                string `json:"angka_kredit"`
	NoSertifikasi              string `json:"nomor_sertifikasi_pegawai"`
	UuidJenisRegis             string `json:"uuid_jenis_regis"`
	IdJenisRegis               uint64 `json:"-"`
	KdJenisRegis               string `json:"kd_jenis_regis"`
	JenisNomorRegis            string `json:"jenis_no_regis"`
	NomorRegis                 string `json:"no_regis"`
}

func (*PegawaiYayasan) TableName() string {
	return "pegawai"
}

type UnitKerjaPegawai struct {
	UuidIndukKerja      string `json:"uuid_induk_kerja"`
	KdIndukKerja        string `json:"kd_induk_kerja"`
	IndukKerja          string `json:"induk_kerja"`
	UuidUnitKerja       string `json:"uuid_unit_kerja"`
	KdUnitKerja         string `json:"kd_unit_kerja"`
	UnitKerja           string `json:"unit_kerja"`
	UuidBagianKerja     string `json:"uuid_bagian_kerja"`
	KdBagianKerja       string `json:"kd_bagian_kerja"`
	BagianKerja         string `json:"bagian_kerja"`
	UuidLokasiKerja     string `json:"uuid_lokasi_kerja"`
	LokasiKerja         string `json:"kd_lokasi_kerja"`
	LokasiDesc          string `json:"lokasi_kerja"`
	NoSkPertama         string `json:"nomor_sk_pertama_unit_kerja"`
	TmtSkPertama        string `json:"tmt_sk_pertama_unit_kerja"`
	TmtSkPertamaIdn     string `json:"tmt_sk_pertama_unit_kerja_idn"`
	KdHomebasePddikti   string `json:"kd_homebase_pddikti"`
	UuidHomebasePddikti string `json:"uuid_homebase_pddikti"`
	KdHomebaseUii       string `json:"kd_homebase_uii"`
	UuidHomebaseUii     string `json:"uuid_homebase_uii"`
}

type PegawaiPNSPTT struct {
	NipPNS                string `json:"nip_pns"`
	NoKartuPegawai        string `json:"no_kartu_pegawai"`
	UuidDetailProfesi     string `json:"uuid_detail_profesi" form:"uuid_detail_profesi"`
	IdDetailProfesi       int    `json:"-"`
	DetailProfesi         string `json:"detail_profesi"`
	UuidPangkatGolongan   string `json:"uuid_pangkat_gol_ruang_pns"`
	KdPangkatGolonganPns  string `json:"kd_pangkat_golongan"`
	PangkatPNS            string `json:"pangkat_pns"`
	GolonganPNS           string `json:"golongan_pns"`
	KdGolonganPNS         string `json:"kd_golongan_pns"`
	KdRuangPNS            string `json:"kd_ruang_pns"`
	TmtPangkatGolongan    string `json:"tmt_pangkat_gol_ruang_pns"`
	TmtPangkatGolonganIdn string `json:"tmt_pangkat_gol_ruang_pns_idn"`
	KdJabatanPns          string `json:"kd_jabatan_pns"`
	JabatanPns            string `json:"jabatan_pns"`
	UuidJabatanPns        string `json:"uuid_jabatan_pns"`
	TmtJabatanPns         string `json:"tmt_jabatan_pns"`
	TmtJabatanPnsIdn      string `json:"tmt_jabatan_pns_idn"`
	MasaKerjaPnsTahun     string `json:"masa_kerja_pns_tahun"`
	MasaKerjaPnsBulan     string `json:"masa_kerja_pns_bulan"`
	AngkaKreditPns        string `json:"angka_kredit_pns"`
	KeteranganPNS         string `json:"keterangan_pns"`
	UuidJenisPTT          string `json:"uuid_jenis_ptt"`
	KdJenisPTT            string `json:"kd_jenis_ptt"`
	JenisPTT              string `json:"jenis_ptt"`
	InstansiAsalPtt       string `json:"instansi_asal_ptt"`
}

type StatusAktif struct {
	FlagAktifPegawai       string `json:"flag_aktif_pegawai"`
	StatusAktifPegawai     string `json:"status_aktif_pegawai"`
	KdStatusAktifPegawai   string `json:"kd_status_aktif_pegawai"`
	UuidStatusAktifPegawai string `json:"uuid_status_aktif_pegawai"`
}

type PegawaiUpdate struct {
	Id                      uint64                  `form:"id" gorm:"primaryKey;<-false"`
	IdPersonalDataPribadi   uint64                  `form:"id_personal_data_pribadi" gorm:"<-:create"`
	FlagAktif               int                     `form:"flag_aktif" gorm:"->"`
	Nik                     string                  `form:"nik" gorm:"->;<-:create"`
	NikKtp                  string                  `form:"nik_ktp" gorm:"->"`
	Nama                    string                  `form:"nama" gorm:"->;<-:create"`
	GelarDepan              string                  `form:"gelar_depan" gorm:"<-:create"`
	GelarBelakang           string                  `form:"gelar_belakang" gorm:"<-:create"`
	TempatLahir             string                  `form:"tempat_lahir" gorm:"<-:create"`
	TglLahir                string                  `form:"tgl_lahir" gorm:"<-:create"`
	JenisKelamin            string                  `form:"jenis_kelamin" gorm:"<-:create"`
	IdAgama                 uint64                  `form:"id_agama" gorm:"<-:create"`
	KdAgama                 string                  `form:"kd_agama" gorm:"<-:create"`
	IdGolonganDarah         uint64                  `form:"id_golongan_darah" gorm:"<-:create"`
	KdGolonganDarah         string                  `form:"kd_golongan_darah" gorm:"<-:create"`
	IdStatusPerkawinan      uint64                  `form:"id_status_perkawinan" gorm:"<-:create"`
	KdStatusPerkawinan      string                  `form:"kd_status_perkawinan" gorm:"<-:create"`
	UuidPendidikanMasuk     string                  `form:"uuid_pendidikan_masuk" gorm:"-"`
	IdPendidikanMasuk       uint64                  `form:"id_pendidikan_masuk"`
	KdPendidikanMasuk       string                  `form:"kd_pendidikan_masuk"`
	UuidPendidikanTerakhir  string                  `form:"uuid_pendidikan_terakhir" gorm:"-"`
	IdPendidikanTerakhir    uint64                  `form:"id_pendidikan_terakhir"`
	KdPendidikanTerakhir    string                  `form:"kd_pendidikan_terakhir"`
	IdStatusPendidikanMasuk uint64                  `form:"id_status_pendidikan_masuk" gorm:"<-:create"`
	KdStatusPendidikanMasuk string                  `form:"kd_status_pendidikan_masuk" gorm:"<-:create"`
	IdJenisPendidikan       uint64                  `form:"id_jenis_pendidikan" gorm:"<-:create"`
	kdJenisPendidikan       string                  `form:"kd_jenis_pendidikan" gorm:"<-:create"`
	UuidJenisPegawai        string                  `form:"uuid_jenis_pegawai" gorm:"-"`
	IdJenisPegawai          uint64                  `form:"id_jenis_pegawai"`
	KdJenisPegawai          string                  `form:"kd_jenis_pegawai"`
	UuidStatusPegawai       string                  `form:"uuid_status_pegawai" gorm:"-"`
	IdStatusPegawai         uint64                  `form:"id_status_pegawai"`
	KdStatusPegawai         string                  `form:"kd_status_pegawai"`
	UuidKelompokPegawai     string                  `form:"uuid_kelompok_pegawai" gorm:"-"`
	IdKelompokPegawai       uint64                  `form:"id_kelompok_pegawai"`
	KdKelompokPegawai       string                  `form:"kd_kelompok_pegawai"`
	UuidDetailProfesi       string                  `form:"uuid_detail_profesi"  gorm:"-"`
	IdDetailProfesi         uint64                  `form:"id_detail_profesi"`
	UuidGolongan            string                  `form:"uuid_golongan" gorm:"-"`
	IdGolongan              uint64                  `form:"id_golongan"`
	KdGolongan              string                  `form:"kd_golongan"`
	UuidRuang               string                  `form:"uuid_ruang" gorm:"-"`
	IdRuang                 uint64                  `form:"id_ruang"`
	KdRuang                 string                  `form:"kd_ruang"`
	UuidUnitKerja1          string                  `form:"uuid_induk_kerja" gorm:"-"` //Perubahan
	IdUnitKerja1            uint64                  `form:"id_unit_kerja1"`
	KdUnit1                 string                  `form:"kd_unit1"`
	UuidUnitKerja2          string                  `form:"uuid_unit_kerja" gorm:"-"` //Perubahan
	IdUnitKerja2            uint64                  `form:"id_unit_kerja2"`
	KdUnit2                 string                  `form:"kd_unit2"`
	UuidUnitKerja3          string                  `form:"uuid_bagian_kerja" gorm:"-"` //Perubahan
	IdUnitKerja3            uint64                  `form:"id_unit_kerja3"`
	KdUnit3                 string                  `form:"kd_unit3"`
	IdUnitKerjaLokasi       uint64                  `form:"id_unit_kerja_lokasi"`
	LokasiKerja             string                  `form:"lokasi_kerja"`
	UuidLokasiKerja         string                  `form:"uuid_lokasi_kerja" gorm:"-"`
	FlagPensiun             string                  `form:"flag_pensiun" gorm:"->"`
	TglPensiun              string                  `form:"tgl_pensiun" gorm:"->"`
	FlagMeninggal           string                  `form:"flag_meninggal" gorm:"->"`
	TglInput                string                  `form:"tgl_input" gorm:"->"`
	UserInput               string                  `form:"user_input" gorm:"->"`
	TglUpdate               string                  `form:"tgl_update" gorm:"->"`
	UserUpdate              string                  `form:"user_update"`
	Uuid                    string                  `form:"uuid" gorm:"->;<-false"`
	PegawaiFungsional       PegawaiFungsionalUpdate `gorm:"foreignkey:IdPegawai;references:Id"`
	PegawaiPNS              PegawaiPNSUpdate        `gorm:"foreignkey:IdPegawai;references:Id"`
}

func (*PegawaiUpdate) TableName() string {
	return "pegawai"
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

type PegawaiPNSUpdate struct {
	Id                    uint64  `form:"-" gorm:"primaryKey"`
	IdPegawai             uint64  `form:"-"`
	NipPns                string  `form:"nip_pns"`
	NoKartuPegawai        string  `form:"no_kartu_pegawai"`
	UuidJenisPtt          string  `form:"uuid_jenis_ptt" gorm:"-"`
	IdJenisPtt            uint64  `form:"id_jenis_ptt"`
	KdJenisPtt            string  `form:"kd_jenis_ptt"`
	InstansiAsal          string  `form:"instansi_asal_ptt" gorm:"column:instansi_asal"`
	UuidPangkatGolongan   string  `form:"uuid_pangkat_gol_ruang_pns" gorm:"-"` //Perubahan
	IdPangkatGolongan     uint64  `form:"id_pangkat_golongan"`
	KdPangkatGolongan     string  `form:"kd_pangkat_golongan"`
	TmtPangkatGolongan    *string `form:"tmt_pangkat_gol_ruang_pns" gorm:"column:tmt_pangkat_golongan"` //Perubahan
	TmtPangkatGolonganIDN string  `form:"tmt_pangkat_gol_ruang_pns_idn" gorm:"-"`                       //Perubahan
	UuidJabatanFungsional string  `form:"uuid_jabatan_pns" gorm:"-"`                                    //Perubahan
	IdJabatanFungsional   uint64  `form:"id_jabatan_fungsional"`
	KdJabatanFungsional   string  `form:"kd_jabatan_fungsional"`
	TmtJabatan            *string `form:"tmt_jabatan_pns" gorm:"tmt_jabatan"`                  //Perubahan
	TmtJabatanIDN         string  `form:"tmt_jabatan_pns_idn" gorm:"-"`                        //Perubahan
	MasaKerjaTahun        string  `form:"masa_kerja_pns_tahun" gorm:"column:masa_kerja_tahun"` //Perubahan
	MasaKerjaBulan        string  `form:"masa_kerja_pns_bulan" gorm:"column:masa_kerja_bulan"` //Perubahan
	AngkaKredit           string  `form:"angka_kredit_pns" gorm:"column:angka_kredit"`         //Perubahan
	Keterangan            string  `form:"keterangan_pns" gorm:"column:keterangan"`             //Perubahan
	TglInput              string  `form:"-" gorm:"-"`
	UserInput             string  `form:"-" gorm:"-"`
	TglUpdate             string  `form:"-" gorm:"-"`
	UserUpdate            string  `form:"-"`
	FlagAktif             uint64  `form:"-" gorm:"-"`
	Uuid                  string  `form:"-" gorm:"-"`
}

func (*PegawaiPNSUpdate) TableName() string {
	return "pegawai_pns"
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

func (p *PegawaiPendidikan) SetTanggalIDN() {
	p.TglSKPenyetaraanIDN = GetIndonesianDate(p.TglSKPenyetaraan)
	p.TglKelulusanIDN = GetIndonesianDate(p.TglKelulusan)
	p.TglIjazahIDN = GetIndonesianDate(p.TglIjazah)
}

func (p *PegawaiYayasan) SetTanggalIDN() {
	p.TmtPangkatGolonganIdn = GetIndonesianDate(p.TmtPangkatGolongan)
	p.TmtJabatanIdn = GetIndonesianDate(p.TmtJabatan)
}

func (p *UnitKerjaPegawai) SetTanggalIDN() {
	p.TmtSkPertamaIdn = GetIndonesianDate(p.TmtSkPertama)
}

func (p *PegawaiPNSPTT) SetTanggalIDN() {
	p.TmtPangkatGolonganIdn = GetIndonesianDate(p.TmtPangkatGolongan)
	p.TmtJabatanPnsIdn = GetIndonesianDate(p.TmtJabatanPns)
}

func (p *PegawaiPendidikan) SetNamaFileIjazah() {
	if p.PathIjazah == "" {
		return
	}
	uploadedFileName := strings.Split(p.PathIjazah, "/")[2]
	splittedPathIjazah := strings.Split(p.PathIjazah, ".")
	fileExtensionIjazah := splittedPathIjazah[1]
	p.NamaFileIjazah = fmt.Sprintf("%s.%s", uploadedFileName, fileExtensionIjazah)
}

func (p *PegawaiPendidikan) SetNamaFilePenyetaraan() {
	if p.PathSKPenyetaraan == "" {
		return
	}
	uploadedFileName := strings.Split(p.PathSKPenyetaraan, "/")[2]
	splittedPathPenyetaraan := strings.Split(p.PathSKPenyetaraan, ".")
	fileExtensionPenyetaraan := splittedPathPenyetaraan[1]
	p.NamaFileSKPenyetaraan = fmt.Sprintf("%s.%s", uploadedFileName, fileExtensionPenyetaraan)
}

func (b *BerkasPendukung) SetDownloadFileName(loc *time.Location) {
	if b.PathFile == "" {
		return
	}
	now := time.Now().In(loc)
	datetime := now.Format("2006-01-02 150405")
	splittedPath := strings.Split(b.PathFile, ".")
	fileExtension := splittedPath[1]
	b.NamaFile = fmt.Sprintf("%s %s %s.%s", datetime, b.NamaPersonal, b.JenisFile, fileExtension)
}

func (b *PegawaiPendidikan) SetDownloadFileNamePendidikan(loc *time.Location) {
	if b.PathIjazah == "" || b.PathSKPenyetaraan == "" {
		return
	}

	// now := time.Now().In(loc)
	// datetime := now.Format("2006-01-02 150405")
	ijazah := "Ijazah"
	splittedPathIjazah := strings.Split(b.PathIjazah, ".")
	fileExtensionIjazah := splittedPathIjazah[1]
	b.NamaFileIjazah = fmt.Sprintf("%s.%s", ijazah, fileExtensionIjazah)

	penyetaraan := "SK Penyetaraan"
	splittedPathPenyetaraan := strings.Split(b.PathSKPenyetaraan, ".")
	fileExtensionPenyetaraan := splittedPathPenyetaraan[1]
	b.NamaFileSKPenyetaraan = fmt.Sprintf("%s.%s", penyetaraan, fileExtensionPenyetaraan)
}

func (a *PegawaiYayasan) SetMasaKerjaTotal(b *UnitKerjaPegawai) {
	now := time.Now()
	yearNow, _ := conv.Int(now.Year())
	monthNow, _ := conv.Int(now.Month())
	dayNow, _ := conv.Int(now.Day())
	maxMonth := 12

	var yearSk int
	var monthSk int
	var daySk int

	if b.TmtSkPertama != "" {
		dateSk, err := time.Parse("2006-01-02", b.TmtSkPertama)
		if err != nil {
			return
		}
		yearSk, _ = conv.Int(dateSk.Year())
		monthSk, _ = conv.Int(dateSk.Month())
		daySk, _ = conv.Int(dateSk.Day())
	}

	yearSkNow := yearNow - yearSk
	monthSkNow := monthNow - monthSk

	if monthSkNow < 0 {
		yearSkNow = yearSkNow - 1
		monthSkNow = maxMonth + monthSkNow
	}

	daySkNow := dayNow - daySk

	if daySkNow < 0 {
		monthSkNow = monthSkNow - 1
	}

	// fmt.Println("Year Sk Now : ", yearSkNow)
	// fmt.Println("Month Sk Now : ", monthSkNow)

	var MasaBawaanBulan, _ = conv.Int(a.MasaKerjaBawaanBulan)
	var MasaBawaanTahun, _ = conv.Int(a.MasaKerjaBawaanTahun)
	var MasaTotalTahun = 0
	var MasaTotalBulan = 0

	// fmt.Println("Year Bawaan Tahun : ", MasaBawaanTahun)
	// fmt.Println("Month Bawaan Tahun : ", MasaBawaanBulan)

	if MasaBawaanTahun != 0 || yearSkNow != 0 {
		MasaTotalTahun = MasaBawaanTahun + yearSkNow
	}

	if MasaBawaanBulan != 00 || monthSkNow != 0 {
		MasaTotalBulan = MasaBawaanBulan + monthSkNow
		if MasaTotalBulan >= 12 {
			MasaTotalBulan = MasaTotalBulan - 12
			MasaTotalTahun = MasaTotalTahun + 1
		}
	}
	// fmt.Println("Year Total Tahun : ", MasaTotalTahun)
	// fmt.Println("Month Total Tahun : ", MasaTotalBulan)

	a.MasaKerjaTotalBulan, _ = conv.String(MasaTotalBulan)
	a.MasaKerjaTotalTahun, _ = conv.String(MasaTotalTahun)
}
