package model

import (
	"fmt"
	"mime/multipart"
	"strings"
	jenisPegawai "svc-insani-go/modules/v1/master-jenis-pegawai/model"
	kelompokPegawai "svc-insani-go/modules/v1/master-kelompok-pegawai/model"
	statusPegawai "svc-insani-go/modules/v1/master-status-pegawai/model"
	unitKerja "svc-insani-go/modules/v1/master-unit-kerja/model"
	"time"
)

type Pegawai struct {
	ID                              string `json:"-" gorm:"primaryKey"`
	NIK                             string `json:"nik" gorm:"type:varchar;not null"`
	Nama                            string `json:"nama" gorm:"type:varchar;not null"`
	GelarDepan                      string `json:"gelar_depan" gorm:"type:varchar"`
	GelarBelakang                   string `json:"gelar_belakang" gorm:"type:varchar"`
	FlagDosen                       int    `json:"flag_dosen"`
	KdUnit2                         int    `json:"kd_unit2"`
	jenisPegawai.JenisPegawai       `json:"jenis_pegawai"`
	kelompokPegawai.KelompokPegawai `json:"kelompok_pegawai"`
	statusPegawai.StatusPegawai     `json:"status_pegawai"`
	UnitKerja                       unitKerja.UnitKerja `json:"unit_kerja" gorm:"foreignKey:KdUnit2"`
	UserInput                       string              `json:"-"`
	UserUpdate                      string              `json:"-"`
	UUID                            string              `json:"uuid"`
}

func (p *Pegawai) SetFlagDosen() {
	if !p.JenisPegawai.IsEmpty() && p.JenisPegawai.KDJenisPegawai == "ED" {
		p.FlagDosen = 1
	}
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
	KdUnitKerja       string `query:"kd_unit_kerja"`
	KdKelompokPegawai string `query:"kd_kelompok_pegawai"`
	Limit             int    `query:"limit"`
	Offset            int    `query:"offset"`
	Cari              string `query:"cari"`
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
	ID              string `json:"-"`
	NIK             string `json:"nik"`
	Nama            string `json:"nama"`
	JenisPegawai    string `json:"jenis_pegawai"`
	KelompokPegawai string `json:"kelompok_pegawai"`
	UnitKerja       string `json:"unit_kerja"`
	UserInput       string `json:"-"`
	UserUpdate      string `json:"-"`
	UUID            string `json:"uuid"`
}

type PegawaiDetail struct {
	PegawaiPribadi    *PegawaiPribadi     `json:"pribadi"`
	JenjangPendidikan []JenjangPendidikan `json:"pendidikan"`
	PegawaiYayasan    *PegawaiYayasan     `json:"kepegawaian"`
	UnitKerjaPegawai  *UnitKerjaPegawai   `json:"unit_kerja"`
	PegawaiPNSPTT     *PegawaiPNSPTT      `json:"negara_ptt"`
	StatusAktif       *StatusAktif        `json:"status_aktif"`
}

type PegawaiYayasan struct {
	ID                   string `json:"-" gorm:"primaryKey;not null"`
	KDJenisPegawai       string `json:"kd_jenis_pegawai"`
	JenisPegawai         string `json:"jenis_pegawai"`
	KdKelompokPegawai    string `json:"kd_kelompok_pegawai"`
	KelompokPegawai      string `json:"kelompok_pegawai"`
	KDStatusPegawai      string `json:"kd_status_pegawai"`
	StatusPegawai        string `json:"status_pegawai"`
	KdPangkat            string `json:"kd_pangkat_golongan"`
	Pangkat              string `json:"pangkat"`
	Golongan             string `json:"golongan"`
	TmtPangkatGolongan   string `json:"tmt_pangkat_gol_ruang_pegawai"`
	KdJabatanFungsional  string `json:"kd_jabatan_fungsional"`
	JabatanFungsional    string `json:"jabatan_fungsional"`
	TmtJabatan           string `json:"tmt_jabatan"`
	MasaKerjaBawaanTahun string `json:"masa_kerja_bawaan_tahun"`
	MasaKerjaBawaanBulan string `json:"masa_kerja_bawaan_bulan"`
	MasaKerjaGajiTahun   string `json:"masa_kerja_gaji_tahun"`
	MasaKerjaGajiBulan   string `json:"masa_kerja_gaji_bulan"`
	MasaKerjaTotalahun   string `json:"masa_kerja_total_tahun"`
	MasaKerjaTotalBulan  string `json:"masa_kerja_total_bulan"`
	AngkaKredit          string `json:"angka_kredit"`
	NoSertifikasi        string `json:"nomor_sertifikasi_pegawai"`
	KdJenisRegis         string `json:"kd_jenis_regis"`
	JenisNomorRegis      string `json:"jenis_no_regis"`
	NomorRegis           string `json:"no_regis"`
}

func (*PegawaiYayasan) TableName() string {
	return "pegawai"
}

type UnitKerjaPegawai struct {
	KdIndukKerja  string `json:"kd_induk_kerja"`
	IndukKerja    string `json:"induk_kerja"`
	KdUnitKerja   string `json:"kd_unit_kerja"`
	UnitKerja     string `json:"unit_kerja"`
	KdBagianKerja string `json:"kd_bagian_kerja"`
	BagianKerja   string `json:"bagian_kerja"`
	LokasiKerja   string `json:"kd_lokasi_kerja"`
	LokasiDesc    string `json:"lokasi_kerja"`
	NoSkPertama   string `json:"nomor_sk_pertama_unit_kerja"`
	TmtSkPertama  string `json:"tmt_sk_pertama_unit_kerja"`
}

type PegawaiPNSPTT struct {
	FlagPns               string `json:"flag_pns"`
	NipPNS                string `json:"nip_pns"`
	NoKartuPegawai        string `json:"no_kartu_pegawai"`
	PangkatGolongan       string `json:"pangkat_gol_ruang_pns"`
	KdPangkatGolongan     string `json:"kd_pangkat_golongan"`
	PangkatPNS            string `json:"pangkat_pns"`
	GolonganPNS           string `json:"golongan_pns"`
	UuidPangkatGolongan   string `json:"uuid_pangkat_gol_ruang_pns"`
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
	KdJenisPTT            string `json:"kd_jenis_ptt"`
	JenisPTT              string `json:"jenis_ptt"`
	InstansiAsalPtt       string `json:"instansi_asal_ptt"`
	KeteranganPtt         string `json:"keterangan_ptt"`
}

type StatusAktif struct {
	FlagAktifPegawai       string `json:"flag_aktif_pegawai"`
	StatusAktifPegawai     string `json:"status_aktif_pegawai"`
	KdStatusAktifPegawai   string `json:"kd_status_aktif_pegawai"`
	UuidStatusAktifPegawai string `json:"uuid_status_aktif_pegawai"`
}

type PegawaiPendidikan struct {
	UuidPendidikan          string                `form:"uuid_pendidikan" json:"uuid_pendidikan"`
	IdPendidikan            string                `form:"id_pendidikan" json:"id_pendidikan" gorm:"primaryKey;column:id"`
	IdPersonalDataPribadi   string                `form:"id_personal_data_pribadi" json:"-"`
	KdJenjang               string                `json:"kd_jenjang_pendidikan"`
	IDJenjang               string                `json:"id_jenjang"`
	UrutanJenjang           string                `json:"-"`
	NamaInstitusi           string                `json:"nama_institusi"`
	Jurusan                 string                `json:"jurusan"`
	TglKelulusan            string                `json:"tgl_kelulusan"`
	TglKelulusanIDN         string                `json:"tgl_kelulusan_idn"`
	FlagIjazahDiakui        string                `form:"flag_ijazah_tertinggi_diakui json:"flag_ijazah_tertinggi_diakui"`
	FlagIjazahTerakhir      string                `form:"flag_ijazah_terakhir json:"flag_ijazah_terakhir"`
	Akreditasi              string                `json:"akreditasi"`
	KonsentrasiBidang       string                `json:"konsentrasi_bidang_ilmu"`
	FlagPerguruanTinggi     int                   `json:"flag_perguruan_tinggi"`
	Gelar                   string                `json:"gelar"`
	NomorInduk              string                `json:"nomor_induk"`
	TahunMasuk              string                `json:"tahun_masuk"`
	JudulTugasAkhir         string                `json:"judul_tugas_akhir"`
	FlagInstitusiLuarNegeri int                   `json:"flag_institusi_luar_negeri"`
	NomorIjazah             string                `json:"nomor_ijazah"`
	TglIjazah               string                `json:"tgl_ijazah"`
	TglIjazahIDN            string                `json:"tgl_ijazah_idn"`
	PathIjazah              string                `json:"path_ijazah"`
	URLIjazah               string                `json:"url_ijazah"`
	NamaFileIjazah          string                `json:"nama_file_ijazah"`
	FlagIjazahTerverifikasi int                   `json:"flag_ijazah_terverifikasi"`
	Nilai                   float64               `json:"nilai"`
	JumlahPelajaran         int                   `json:"jumlah_pelajaran"`
	NamaFileSKPenyetaraan   string                `json:"nama_file_sk_penyetaraan"`
	NomorSKPenyetaraan      string                `json:"nomor_sk_penyetaraan"`
	TglSKPenyetaraan        string                `json:"tgl_sk_penyetaraan"`
	TglSKPenyetaraanIDN     string                `json:"tgl_sk_penyetaraan_idn"`
	PathSKPenyetaraan       string                `json:"path_sk_penyetaraan"`
	URLSKPenyetaraan        string                `json:"url_sk_penyetaraan"`
	FileSKPenyetaraan       *multipart.FileHeader `form:"-" json:"-"`
	UUIDPersonal            string                `form:"-" json:"uuid_personal"`
	UserInput               string                `form:"-" json:"-"`
	TglInput                string                `form:"-" json:"-"`
	UserUpdate              string                `form:"-" json:"-"`
	BerkasPendukungList     `form:"-" json:"berkas_pendukung"`
	BerkasPendukung         BerkasPendukungMap  `form:"-" json:"-"`
	OldBerkasPendukungList  BerkasPendukungList `form:"-" json:"-"`
}

type BerkasPendukung struct {
	IDPendidikan string `json:"-"`
	IDPersonal   string `json:"-"`
	NamaPersonal string `json:"-"`
	IDJenisFile  string `json:"-"`
	KdJenisFile  string `json:"kd_jenis_file_pendidikan"`

	JenisFile     string                `json:"jenis_file_pendidikan"`
	UUIDJenisFile string                `json:"uuid_jenis_file_pendidikan"`
	File          *multipart.FileHeader `json:"-"`
	Folder        string                `json:"-"`
	SubFolder     string                `json:"-"`
	NamaFile      string                `json:"nama_file_pendidikan"`
	PathFile      string                `json:"-"`
	URLFile       string                `json:"url_file_pendidikan"`
	UUIDFile      string                `json:"uuid_file_pendidikan"`
	UserUpdate    string                `json:"-"`
	UUIDPersonal  string                `json:"-"`
}

type JenjangPendidikan struct {
	JenjangPendidikan string              `json:"jenjang"`
	UrutanJenjang     string              `json:"-"`
	Data              []PegawaiPendidikan `json:"data"`
}

type PegawaiUpdate struct {
	Id                      int                     `form:"id" gorm:"primaryKey"`
	IdPersonalDataPribadi   string                  `form:"id_personal_data_pribadi" gorm:"<-:create"`
	FlagAktif               int                     `form:"flag_aktif" gorm:"->"`
	Nik                     string                  `form:"nik" gorm:"->"`
	NikKtp                  string                  `form:"nik_ktp" gorm:"->"`
	Nama                    string                  `form:"nama" gorm:"->"`
	GelarDepan              string                  `form:"gelar_depan" gorm:"<-:create"`
	GelarBelakang           string                  `form:"gelar_belakang" gorm:"<-:create"`
	TempatLahir             string                  `form:"tempat_lahir" gorm:"<-:create"`
	TglLahir                string                  `form:"tgl_lahir" gorm:"<-:create"`
	JenisKelamin            string                  `form:"jenis_kelamin" gorm:"<-:create"`
	IdAgama                 string                  `form:"id_agama" gorm:"<-:create"`
	KdAgama                 string                  `form:"kd_agama" gorm:"<-:create"`
	IdGolonganDarah         int                     `form:"id_golongan_darah" gorm:"<-:create"`
	KdGolonganDarah         string                  `form:"kd_golongan_darah" gorm:"<-:create"`
	IdStatusPerkawinan      int                     `form:"id_status_perkawinan" gorm:"<-:create"`
	KdStatusPerkawinan      string                  `form:"kd_status_perkawinan" gorm:"<-:create"`
	IdPendidikanMasuk       int                     `form:"id_pendidikan_masuk" gorm:"<-:create"`
	KdPendidikanMasuk       string                  `form:"kd_pendidikan_masuk" gorm:"<-:create"`
	IdStatusPendidikanMasuk int                     `form:"id_status_pendidikan_masuk" gorm:"<-:create"`
	KdStatusPendidikanMasuk string                  `form:"kd_status_pendidikan_masuk" gorm:"<-:create"`
	IdPendidikanTerakhir    int                     `form:"id_pendidikan_terakhir" gorm:"<-:create"`
	KdPendidikanTerakhir    string                  `form:"kd_pendidikan_terakhir" gorm:"<-:create"`
	IdJenisPendidikan       int                     `form:"id_jenis_pendidikan" gorm:"<-:create"`
	kdJenisPendidikan       string                  `form:"kd_jenis_pendidikan" gorm:"<-:create"`
	UuidJenisPegawai        string                  `form:"uuid_jenis_pegawai" gorm:"-"`
	IdJenisPegawai          int                     `form:"id_jenis_pegawai"`
	KdJenisPegawai          string                  `form:"kd_jenis_pegawai"`
	UuidStatusPegawai       string                  `form:"uuid_status_pegawai" gorm:"-"`
	IdStatusPegawai         int                     `form:"id_status_pegawai"`
	KdStatusPegawai         string                  `form:"kd_status_pegawai"`
	UuidKelompokPegawai     string                  `form:"uuid_kelompok_pegawai" gorm:"-"`
	IdKelompokPegawai       int                     `form:"id_kelompok_pegawai"`
	KdKelompokPegawai       string                  `form:"kd_kelompok_pegawai"`
	UuidGolongan            string                  `form:"uuid_golongan" gorm:"-"`
	IdGolongan              int                     `form:"id_golongan"`
	KdGolongan              string                  `form:"kd_golongan"`
	UuidRuang               string                  `form:"uuid_ruang" gorm:"-"`
	IdRuang                 int                     `form:"id_ruang"`
	KdRuang                 string                  `form:"kd_ruang"`
	UuidUnitKerja1          string                  `form:"uuid_unit_kerja1" gorm:"-"`
	IdUnitKerja1            int                     `form:"id_unit_kerja1"`
	KdUnit1                 string                  `form:"kd_unit1"`
	UuidUnitKerja2          string                  `form:"uuid_unit_kerja2" gorm:"-"`
	IdUnitKerja2            int                     `form:"id_unit_kerja2"`
	KdUnit2                 string                  `form:"kd_unit2"`
	UuidUnitKerja3          string                  `form:"uuid_unit_kerja3" gorm:"-"`
	IdUnitKerja3            int                     `form:"id_unit_kerja3"`
	KdUnit3                 string                  `form:"kd_unit3"`
	IdUnitKerjaLokasi       int                     `form:"id_unit_kerja_lokasi"`
	LokasiKerja             string                  `form:"lokasi_kerja"`
	UuidUnitKerjaLokasi     string                  `form:"uuid_unit_kerja_lokasi" gorm:"-"`
	FlagPensiun             string                  `form:"flag_pensiun" gorm:"->"`
	TglPensiun              string                  `form:"tgl_pensiun" gorm:"<-:create" gorm:"->"`
	FlagMeninggal           string                  `form:"flag_meninggal" gorm:"->"`
	TglInput                string                  `form:"tgl_input" gorm:"->" gorm:"->"`
	UserInput               string                  `form:"user_input" gorm:"->"`
	TglUpdate               string                  `form:"tgl_update" gorm:"->"`
	UserUpdate              string                  `form:"user_update"`
	Uuid                    string                  `form:"uuid" gorm:"->"`
	PegawaiFungsional       PegawaiFungsionalUpdate `gorm:"foreignkey:IdPegawai;references:Id"`
	PegawaiPNS              PegawaiPNSUpdate        `gorm:"foreignkey:IdPegawai;references:Id"`
	// PegawaiPendidikan       []PegawaiPendidikan     `gorm:"foreignkey:IdPersonal;references:IdPersonalDataPribadi"`
}

func (*PegawaiUpdate) TableName() string {
	return "pegawai"
}

type PegawaiFungsionalUpdate struct {
	Id                       int        `form:"-"`
	IdKafka                  int        `form:"-"`
	IdPegawai                int        `form:"-"`
	UuidPangkatGolongan      string     `form:"uuid_pangkat_golongan" gorm:"-"`
	IdPangkatGolongan        int        `form:"id_pangkat_golongan"`
	KdPangkatGolongan        string     `form:"kd_pangkat_golongan"`
	UuidJabatanFungsional    string     `form:"uuid_jabatan_fungsional" gorm:"-"`
	IdJabatanFungsional      int        `form:"id_jabatan_fungsional"`
	KdJabatanFungsional      string     `form:"kd_jabatan_fungsional"`
	TmtPangkatGolongan       string     `form:"tmt_pangkat_golongan"`
	TmtJabatan               *time.Time `form:"tmt_jabatan"`
	MasaKerjaBawaanTahun     string     `form:"masa_kerja_bawaan_tahun"`
	MasaKerjaBawaanBulan     string     `form:"masa_kerja_bawaan_bulan"`
	MasaKerjaGajiTahun       string     `form:"masa_kerja_gaji_tahun"`
	MasaKerjaGajiBulan       string     `form:"masa_kerja_gaji_bulan"`
	MasaKerjaTotalTahun      string     `form:"masa_kerja_total_tahun"`
	MasaKerjaTotalBulan      string     `form:"masa_kerja_total_bulan"`
	AngkaKredit              string     `form:"angka_kredit"`
	NomorSertifikasi         string     `form:"nomor_sertifikasi"`
	UuidJenisNomorRegistrasi string     `form:"uuid_jenis_nomor_registrasi" gorm:"-"`
	IdJenisNomorRegistrasi   int        `form:"id_jenis_nomor_registrasi"`
	KdJenisNomorRegistrasi   string     `form:"kd_jenis_nomor_registrasi"`
	NomorRegistrasi          string     `form:"nomor_registrasi"`
	NomorSkPertama           string     `form:"nomor_sk_pertama"`
	TmtSkPertama             string     `form:"tmt_sk_pertama"`
	UuidStatusPegawaiAktif   string     `form:"uuid_status_pegawai_aktif" gorm:"-"`
	IdStatusPegawaiAktif     int        `form:"id_status_pegawai_aktif"`
	KdStatusPegawaiAktif     string     `form:"kd_status_pegawai_aktif"`
	TglInput                 string     `form:"-" gorm:"-"`
	UserInput                string     `form:"-" gorm:"-"`
	TglUpdate                string     `form:"-" gorm:"-"`
	UserUpdate               string     `form:"-"`
	FlagAktif                int        `form:"-" gorm:"-"`
	// Uuid                     string `form:"-"`
}

func (*PegawaiFungsionalUpdate) TableName() string {
	return "pegawai_fungsional"
}

type PegawaiPNSUpdate struct {
	Id                    int    `form:"-" gorm:"primaryKey"`
	IdPegawai             int    `form:"-"`
	NipPns                string `form:"nip_pns"`
	NoKartuPegawai        string `form:"no_kartu_pegawai"`
	UuidJenisPtt          string `form:"uuid_jenis_ptt" gorm:"-"`
	IdJenisPtt            int    `form:"id_jenis_ptt"`
	KdJenisPtt            string `form:"kd_jenis_ptt"`
	InstansiAsal          string `form:"instansi_asal"`
	UuidPangkatGolongan   string `form:"uuid_pangkat_golongan" gorm:"-"`
	IdPangkatGolongan     int    `form:"id_pangkat_golongan"`
	KdPangkatGolongan     string `form:"kd_pangkat_golongan"`
	TmtPangkatGolongan    string `form:"tmt_pangkat_golongan"`
	UuidJabatanFungsional string `form:"uuid_jabatan_fungsional" gorm:"-"`
	IdJabatanFungsional   int    `form:"id_jabatan_fungsional"`
	KdJabatanFungsional   string `form:"kd_jabatan_fungsional"`
	TmtJabatan            string `form:"tmt_jabatan"`
	MasaKerjaTahun        string `form:"masa_kerja_tahun"`
	MasaKerjaBulan        string `form:"masa_kerja_bulan"`
	AngkaKredit           string `form:"angka_kredit"`
	Keterangan            string `form:"keterangan_pns" gorm:"column:keterangan"`
	TglInput              string `form:"-" gorm:"-"`
	UserInput             string `form:"-" gorm:"-"`
	TglUpdate             string `form:"-" gorm:"-"`
	UserUpdate            string `form:"-"`
	FlagAktif             int    `form:"-" gorm:"-"`
	Uuid                  string `form:"-"`
}

func (*PegawaiPNSUpdate) TableName() string {
	return "pegawai_pns"
}

// type PegawaiPTTUpdate struct {
// 	Id           string `form:"-"`
// 	IdPegawai    int    `form:"-"`
// 	UuidJenisPtt string `form:"uuid_jenis_ptt" gorm:"-"`
// 	IdJenisPtt   int    `form:"id_jenis_ptt"`
// 	KdJenisPtt   string `form:"kd_jenis_ptt"`
// 	InstansiAsal string `form:"instansi_asal"`
// 	Keterangan   string `form:"keterangan"`
// 	TglInput     string `form:"-" gorm:"-"`
// 	UserInput    string `form:"-"`
// 	TglUpdate    string `form:"-" gorm:"-"`
// 	UserUpdate   string `form:"-"`
// 	FlagAktif    int    `form:"-"`
// 	Uuid         string `form:"-"`
// }

// func (*PegawaiPTTUpdate) TableName() string {
// 	return "pegawai_tidak_tetap"
// }

type PegawaiPendidikanUpdate struct {
	UuidPendidikan        string `form:"uuid_pendidikan" json:"uuid_pendidikan"`
	IdPendidikan          string `form:"id_pendidikan" json:"id_pendidikan" gorm:"primaryKey;column:id"`
	IdPersonalDataPribadi string `form:"id_personal_data_pribadi" json:"-"`
	FlagIjazahDiakui      string `form:"flag_ijazah_tertinggi_diakui json:"flag_ijazah_tertinggi_diakui"`
	FlagIjazahTerakhir    string `form:"flag_ijazah_terakhir json:"flag_ijazah_terakhir"`
	NomorIjazah           string `json:"nomor_ijazah"`
	UUIDPersonal          string `form:"-" json:"uuid_personal"`
	UserUpdate            string `form:"-" json:"-"`
}

func (*PegawaiPendidikanUpdate) TableName() string {
	return "pegawai_pendidikan"
}

type BerkasPendukungList []BerkasPendukung

func (list BerkasPendukungList) MapByIdPendidikan() map[string][]BerkasPendukung {
	m := make(map[string][]BerkasPendukung)
	for _, pendidikan := range list {
		m[pendidikan.IDPendidikan] = append(m[pendidikan.IDPendidikan], pendidikan)
	}
	return m
}

type BerkasPendukungMap map[int]BerkasPendukung

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
	// now := time.Now().In(loc)
	// datetime := now.Format("2006-01-02 150405")
	splittedPath := strings.Split(b.PathFile, ".")
	fileExtension := splittedPath[1]
	b.NamaFile = fmt.Sprintf("%s.%s", b.JenisFile, fileExtension)
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
