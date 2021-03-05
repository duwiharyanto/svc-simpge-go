package model

import (
	"mime/multipart"
	jenisPegawai "svc-insani-go/modules/v1/master-jenis-pegawai/model"
	kelompokPegawai "svc-insani-go/modules/v1/master-kelompok-pegawai/model"
	statusPegawai "svc-insani-go/modules/v1/master-status-pegawai/model"
	unitKerja "svc-insani-go/modules/v1/master-unit-kerja/model"
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
	PegawaiPribadi   *PegawaiPribadi   `json:"pribadi"`
	PegawaiYayasan   *PegawaiYayasan   `json:"kepegawaian"`
	UnitKerjaPegawai *UnitKerjaPegawai `json:"unit_kerja"`
	PegawaiPNSPTT    *PegawaiPNSPTT    `json:"negara_ptt"`
	StatusAktif      *StatusAktif      `json:"status_aktif"`
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
	ID                      string                `form:"-" json:"-"`
	IDPersonal              string                `form:"-" json:"-"`
	NamaPersonal            string                `form:"-" json:"-"`
	Akreditasi              string                `form:"-" json:"akreditasi"`
	UUIDAkreditasi          string                `form:"uuid_akreditasi" json:"uuid_akreditasi"`
	KonsentrasiBidang       string                `form:"konsentrasi_bidang_ilmu" json:"konsentrasi_bidang_ilmu"`
	Jurusan                 string                `form:"jurusan" json:"jurusan"`
	KdJenjang               string                `form:"-" json:"kd_jenjang"`
	Jenjang                 string                `form:"-" json:"jenjang"`
	KeteranganJenjang       string                `form:"-" json:"-"`
	FlagPerguruanTinggi     int                   `form:"-" json:"flag_perguruan_tinggi"`
	UUIDJenjang             string                `form:"uuid_jenjang" json:"uuid_jenjang"`
	Gelar                   string                `form:"gelar" json:"gelar"`
	NomorInduk              string                `form:"nomor_induk" json:"nomor_induk"`
	TahunMasuk              string                `form:"tahun_masuk" json:"tahun_masuk"`
	TahunLulus              string                `form:"-" json:"tahun_lulus"`
	JudulTugasAkhir         string                `form:"judul_tugas_akhir" json:"judul_tugas_akhir"`
	NamaInstitusi           string                `form:"nama_institusi" json:"nama_institusi"`
	FlagInstitusiLuarNegeri int                   `form:"flag_institusi_luar_negeri" json:"flag_institusi_luar_negeri"`
	NomorIjazah             string                `form:"nomor_ijazah" json:"nomor_ijazah"`
	TglIjazah               string                `form:"tgl_ijazah" json:"tgl_ijazah"`
	TglIjazahIDN            string                `form:"-" json:"tgl_ijazah_idn"`
	PathIjazah              string                `form:"-" json:"-"`
	FileIjazah              *multipart.FileHeader `form:"-" json:"-"`
	URLIjazah               string                `form:"-" json:"url_ijazah"`
	NamaFileIjazah          string                `form:"-" json:"nama_file_ijazah"`
	FlagIjazahTerverifikasi int                   `form:"flag_ijazah_terverifikasi" json:"flag_ijazah_terverifikasi"`
	Nilai                   float64               `form:"nilai" json:"nilai"`
	JumlahPelajaran         int                   `form:"jumlah_pelajaran" json:"jumlah_pelajaran"`
	TglKelulusan            string                `form:"tgl_kelulusan" json:"tgl_kelulusan"`
	TglKelulusanIDN         string                `form:"-" json:"tgl_kelulusan_idn"`
	PathSKPenyetaraan       string                `form:"-" json:"-"`
	FileSKPenyetaraan       *multipart.FileHeader `form:"-" json:"-"`
	URLSKPenyetaraan        string                `form:"-" json:"url_sk_penyetaraan"`
	NamaFileSKPenyetaraan   string                `form:"-" json:"nama_file_sk_penyetaraan"`
	NomorSKPenyetaraan      string                `form:"nomor_sk_penyetaraan" json:"nomor_sk_penyetaraan"`
	TglSKPenyetaraan        string                `form:"tgl_sk_penyetaraan" json:"tgl_sk_penyetaraan"`
	TglSKPenyetaraanIDN     string                `form:"-" json:"tgl_sk_penyetaraan_idn"`
	UUIDPersonal            string                `form:"-" json:"uuid_personal"`
	UserInput               string                `form:"-" json:"-"`
	TglInput                string                `form:"-" json:"-"`
	UserUpdate              string                `form:"-" json:"-"`
	IsUser                  string                `form:"user" json:"-"`
	UUID                    string                `form:"uuid" json:"uuid"`
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
	SubFolder     string                `json:"-"` // folder terakhir di mana file berada
	NamaFile      string                `json:"nama_file_pendidikan"`
	PathFile      string                `json:"-"`
	URLFile       string                `json:"url_file_pendidikan"`
	UUIDFile      string                `json:"uuid_file_pendidikan"`
	UserUpdate    string                `json:"-"`
	UUIDPersonal  string                `json:"-"`
}

type BerkasPendukungList []BerkasPendukung

type BerkasPendukungMap map[int]BerkasPendukung
