package model

import (
	"fmt"
	"mime/multipart"
	"strings"
	"time"
)

type PegawaiPendidikan struct {
	UuidPendidikan               string                `form:"uuid_pendidikan" json:"uuid_pendidikan"`
	IdPendidikan                 string                `form:"id_pendidikan" json:"-" gorm:"primaryKey;column:id"`
	IdPersonalDataPribadi        string                `form:"id_personal_data_pribadi" json:"-"`
	KdJenjang                    string                `json:"kd_jenjang_pendidikan"`
	IDJenjang                    string                `json:"-"`
	IDJenjangPddDetailDiakui     string                `json:"-"`
	KdJenjangPddDetailDiakui     string                `json:"kd_jenis_pdd_diakui"`
	NamaJenjangPddDetailDiakui   string                `json:"nama_jenis_pdd_diakui"`
	UuidJenjangPddDetailDiakui   string                `json:"uuid_jenis_pdd_diakui"`
	IDJenjangPddDetailTerakhir   string                `json:"-"`
	KdJenjangPddDetailTerakhir   string                `json:"kd_jenis_pdd_terakhir"`
	NamaJenjangPddDetailTerakhir string                `json:"nama_jenis_pdd_terakhir"`
	UuidJenjangPddDetailTerakhir string                `json:"uuid_jenis_pdd_terakhir"`
	UrutanJenjang                string                `json:"-"`
	NamaInstitusi                string                `json:"nama_institusi"`
	Jurusan                      string                `json:"jurusan"`
	TglKelulusan                 string                `json:"tgl_kelulusan"`
	TglKelulusanIDN              string                `json:"tgl_kelulusan_idn"`
	FlagIjazahDiakui             string                `form:"flag_ijazah_tertinggi_diakui" json:"flag_ijazah_tertinggi_diakui"`
	FlagIjazahTerakhir           string                `form:"flag_ijazah_terakhir" json:"flag_ijazah_terakhir"`
	Akreditasi                   string                `json:"akreditasi"`
	KonsentrasiBidang            string                `json:"konsentrasi_bidang_ilmu"`
	FlagPerguruanTinggi          int                   `json:"flag_perguruan_tinggi"`
	Gelar                        string                `json:"gelar"`
	NomorInduk                   string                `json:"nomor_induk"`
	TahunMasuk                   string                `json:"tahun_masuk"`
	JudulTugasAkhir              string                `json:"judul_tugas_akhir"`
	FlagInstitusiLuarNegeri      int                   `json:"flag_institusi_luar_negeri"`
	NomorIjazah                  string                `json:"nomor_ijazah"`
	TglIjazah                    string                `json:"tgl_ijazah"`
	TglIjazahIDN                 string                `json:"tgl_ijazah_idn"`
	PathIjazah                   string                `json:"path_ijazah"`
	URLIjazah                    string                `json:"url_ijazah"`
	NamaFileIjazah               string                `json:"nama_file_ijazah"`
	FlagIjazahTerverifikasi      int                   `json:"flag_ijazah_terverifikasi"`
	Nilai                        float64               `json:"nilai"`
	JumlahPelajaran              int                   `json:"jumlah_pelajaran"`
	NamaFileSKPenyetaraan        string                `json:"nama_file_sk_penyetaraan"`
	NomorSKPenyetaraan           string                `json:"nomor_sk_penyetaraan"`
	TglSKPenyetaraan             string                `json:"tgl_sk_penyetaraan"`
	TglSKPenyetaraanIDN          string                `json:"tgl_sk_penyetaraan_idn"`
	PathSKPenyetaraan            string                `json:"path_sk_penyetaraan"`
	URLSKPenyetaraan             string                `json:"url_sk_penyetaraan"`
	FileSKPenyetaraan            *multipart.FileHeader `form:"-" json:"-"`
	UUIDPersonal                 string                `form:"-" json:"uuid_personal"`
	UserInput                    string                `form:"-" json:"-"`
	TglInput                     string                `form:"-" json:"-"`
	UserUpdate                   string                `form:"-" json:"-"`
	BerkasPendukungList          `form:"-" json:"berkas_pendukung"`
	BerkasPendukung              BerkasPendukungMap  `form:"-" json:"-"`
	OldBerkasPendukungList       BerkasPendukungList `form:"-" json:"-"`
}

type PegawaiPendidikanList []PegawaiPendidikan

func (pp PegawaiPendidikanList) MapByUuid() map[string]PegawaiPendidikan {
	m := make(map[string]PegawaiPendidikan)
	for _, p := range pp {
		m[p.UuidPendidikan] = p
	}
	return m
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

type DataPendidikanDetail struct {
	UuidPendidikanMasuk string              `json:"uuid_pendidikan_masuk"`
	KdPendidikanMasuk   string              `json:"kd_pendidikan_masuk"`
	PendidikanMasuk     string              `json:"pendidikan_masuk"`
	Data                []JenjangPendidikan `json:"data_pendidikan"`
}

func (pendidikanDetail DataPendidikanDetail) IsEmpty() bool {
	return pendidikanDetail.KdPendidikanMasuk == "" && pendidikanDetail.PendidikanMasuk == ""
}

type JenjangPendidikan struct {
	JenjangPendidikan string              `json:"jenjang"`
	UrutanJenjang     string              `json:"-"`
	Data              []PegawaiPendidikan `json:"data"`
}

type JenjangPendidikanList []JenjangPendidikan

func (jj JenjangPendidikanList) PegawaiPendidikan() PegawaiPendidikanList {
	educations := []PegawaiPendidikan{}
	for _, j := range jj {
		educations = append(educations, j.Data...)
	}
	return educations
}

type PendidikanPersonal struct {
	Data []JenjangPendidikan `json:"data_pendidikan"`
}

type PegawaiPendidikanUpdate struct {
	UuidPendidikan             string `form:"uuid_pendidikan" json:"uuid_pendidikan"`
	IdPendidikan               uint64 `form:"id_pendidikan" json:"id_pendidikan" gorm:"primaryKey;column:id"`
	IdPersonalDataPribadi      uint64 `form:"id_personal_data_pribadi" json:"-"`
	FlagIjazahDiakui           string `form:"flag_ijazah_diakui" json:"flag_ijazah_diakui"`
	FlagIjazahTerakhir         string `form:"flag_ijazah_terakhir" json:"flag_ijazah_terakhir"`
	IdJenjangPddDetailDiakui   uint64 `form:"-" json:"-" gorm:"default:null"`
	IdJenjangPddDetailTerakhir uint64 `form:"-" json:"-" gorm:"default:null"`
	NomorIjazah                string `json:"nomor_ijazah"`
	UUIDPersonal               string `form:"-" json:"uuid_personal"`
	UserUpdate                 string `form:"-" json:"-"`
}

func (*PegawaiPendidikanUpdate) TableName() string {
	return "pegawai_pendidikan"
}

type PegawaiPendidikanRequest struct {
	UuidPendidikanDiakui                 string // uuid dari pendidikan tertinggi diakui
	UuidPendidikanTerakhir               string // uuid dari pendidikan terakhir
	IdJenjangPendidikanDetailDiakui      *uint64
	IdJenjangPendidikanDetailTerakhir    *uint64
	UuidJenjangPendidikanTertinggiDiakui string
	IdPersonalPegawai                    uint64
	UserUpdate                           *string
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
