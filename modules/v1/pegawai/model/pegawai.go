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
	ID                              string `json:"-"`
	NIK                             string `json:"nik"`
	Nama                            string `json:"nama"`
	GelarDepan                      string `json:"gelar_depan"`
	GelarBelakang                   string `json:"gelar_belakang"`
	FlagDosen                       int    `json:"flag_dosen"`
	jenisPegawai.JenisPegawai       `json:"jenis_pegawai"`
	kelompokPegawai.KelompokPegawai `json:"kelompok_pegawai"`
	statusPegawai.StatusPegawai     `json:"status_pegawai"`
	unitKerja.UnitKerja             `json:"unit_kerja"`
	UserInput                       string `json:"-"`
	UserUpdate                      string `json:"-"`
	UUID                            string `json:"uuid"`
}

func (p *Pegawai) SetFlagDosen() {
	if !p.JenisPegawai.IsEmpty() && p.JenisPegawai.KDJenisPegawai == "ED" {
		p.FlagDosen = 1
	}
}

type PegawaiRequest struct {
	KdUnitKerja       string `query:"kd_unit_kerja"`
	KdKelompokPegawai string `query:"kd_kelompok_pegawai"`
	Limit             int    `query:"limit"`
	Offset            int    `query:"offset"`
	Cari              string `query:"cari"`
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
	UuidPendidikan          string                `json:"uuid_pendidikan"`
	IdPendidikan            string                `json:"id_pendidikan"`
	KdJenjang               string                `json:"kd_jenjang_pendidikan"`
	IDJenjang               string                `json:"id_jenjang"`
	UrutanJenjang           string                `json:"-"`
	NamaInstitusi           string                `json:"nama_institusi"`
	Jurusan                 string                `json:"jurusan"`
	TglKelulusan            string                `json:"tgl_kelulusan"`
	TglKelulusanIDN         string                `json:"tgl_kelulusan_idn"`
	FlagIjazahDiakui        string                `json:"flag_ijazah_tertinggi_diakui"`
	FlagIjazahTerakhir      string                `json:"flag_ijazah_terakhir"`
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
	IDPendidikan  string                `json:"-"`
	IDPersonal    string                `json:"-"`
	NamaPersonal  string                `json:"-"`
	IDJenisFile   string                `json:"-"`
	KdJenisFile   string                `json:"kd_jenis_file_pendidikan"`
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
