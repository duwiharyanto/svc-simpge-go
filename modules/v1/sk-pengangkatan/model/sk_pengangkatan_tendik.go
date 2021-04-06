package model

import (
	"mime/multipart"
	"strings"
	"svc-insani-go/app/helper"
	organisasiV2 "svc-insani-go/modules/v2/organisasi/model"
)

func EmptySKPengangkatanTendiki() SKPengangkatanTendik {
	return SKPengangkatanTendik{}
}

type SKPengangkatanTendik struct {
	// Jenis SK
	IDJenisSK   string `json:"-" form:"-"`
	UUIDJenisSK string `json:"uuid_jenis_sk" form:"uuid_jenis_sk"`
	KDJenisSK   string `json:"kd_jenis_sk" form:"-"`
	JenisSK     string `json:"jenis_sk" form:"-"`

	// Kelompok SK Pengangkatan
	IDKelompokSKPengangkatan   string `json:"-" form:"-"`
	UUIDKelompokSKPengangkatan string `json:"uuid_kelompok_sk_pengangkatan" form:"uuid_kelompok_sk_pengangkatan"`
	KDKelompokSKPengangkatan   string `json:"kd_kelompok_sk_pengangkatan" form:"-"`
	KelompokSKPengangkatan     string `json:"kelompok_sk_pengangkatan" form:"-"`

	// Unit yang mengangkat
	IDUnitPengangkat   string `json:"-" form:"-"`
	UUIDUnitPengangkat string `json:"uuid_unit_pengangkat" form:"uuid_unit_pengangkat"`
	KDUnitPengangkat   string `json:"kd_unit_pengangkat" form:"-"`
	UnitPengangkat     string `json:"unit_pengangkat" form:"-"`

	// Unit kerja
	IDUnitPegawai   string `json:"-" form:"-"`
	UUIDUnitPegawai string `json:"uuid_unit_kerja" form:"uuid_unit_kerja"`
	KDUnitPegawai   string `json:"kd_unit_kerja" form:"-"`
	UnitPegawai     string `json:"unit_kerja" form:"-"`

	// Jabatan fungsional
	IDJabatanFungsional   string `json:"-" form:"-"`
	UUIDJabatanFungsional string `json:"uuid_jabatan_fungsional" form:"uuid_jabatan_fungsional"`
	KDJabatanFungsional   string `json:"kd_jabatan_fungsional" form:"-"`
	JabatanFungsional     string `json:"jabatan_fungsional" form:"-"`

	// Pangkat golongan/ruang
	IDPangkatGolonganPegawai   string `json:"-" form:"-"`
	UUIDPangkatGolonganPegawai string `json:"uuid_pangkat_golongan_pegawai" form:"uuid_pangkat_golongan_pegawai"`
	KDPangkatGolonganPegawai   string `json:"kd_pangkat_golongan_pegawai" form:"-"`
	PangkatGolonganPegawai     string `json:"pangkat_golongan_pegawai" form:"-"`

	GajiPokok            int `json:"gaji_pokok" form:"gaji_pokok"`
	MasaRilBulan         int `json:"masa_kerja_ril_bulan" form:"masa_kerja_ril_bulan"`
	MasaRilTahun         int `json:"masa_kerja_ril_tahun" form:"masa_kerja_ril_tahun"`
	MasaGajiBulan        int `json:"masa_kerja_gaji_bulan" form:"masa_kerja_gaji_bulan"`
	MasaGajiTahun        int `json:"masa_kerja_gaji_tahun" form:"masa_kerja_gaji_tahun"`
	MasaKerjaDiakuiBulan int `json:"masa_kerja_diakui_bulan" form:"masa_kerja_diakui_bulan"`
	MasaKerjaDiakuiTahun int `json:"masa_kerja_diakui_tahun" form:"masa_kerja_diakui_tahun"`

	// Status pengangkatan
	IDStatusPengangkatan   string `json:"-" form:"-"`
	UUIDStatusPengangkatan string `json:"uuid_status_pengangkatan" form:"uuid_status_pengangkatan"`
	KDStatusPengangkatan   string `json:"kd_status_pengangkatan" form:"-"`
	StatusPengangkatan     string `json:"status_pengangkatan" form:"-"`

	// Ijazah yang diakui
	IDJenisIjazah   string `json:"-" form:"-"`
	UUIDJenisIjazah string `json:"uuid_jenis_ijazah" form:"uuid_jenis_ijazah"`
	KDJenisIjazah   string `json:"kd_jenis_ijazah" form:"-"`
	JenisIjazah     string `json:"jenis_ijazah" form:"-"`

	TanggalDitetapkan string                `json:"tanggal_ditetapkan" form:"tanggal_ditetapkan"`
	PathSKTendik      string                `json:"-"`
	URLSKTendik       string                `json:"url_sk_pengangkatan" form:"url_sk_pengangkatan"`
	FileSKTendik      *multipart.FileHeader `form:"-"`

	IDPegawaiPenetap    string `json:"-" form:"-"`
	UUIDPegawaiPenetap  string `json:"uuid_pegawai_penetapan" form:"uuid_pegawai_penetapan"`
	IDGolonganPegawai   string `json:"-" form:"-"`
	UUIDGolonganPegawai string `json:"uuid_golongan_pegawai" form:"uuid_golongan_pegawai"`

	UserInput                string `json:"-" form:"-"`
	UserUpdate               string `json:"-" form:"-"`
	UUID                     string `json:"-" form:"-"`
	IDSKPegawai              string `json:"-" form:"-"`
	UUIDSKPengangkatanTendik string `json:"uuid_sk_pengangkatan_tendik" form:"uuid_sk_pengangkatan_tendik"`
}

func (sk SKPengangkatanTendik) IsEmpty() bool {
	return sk == SKPengangkatanTendik{}
}

type SKPengangkatanTendikResponse struct {
	Data []SKPengangkatanTendik `json:"data"`
}

type SKPengangkatanTendikDetail struct {
	NamaPegawai         string `json:"nama_pegawai"`
	NIKPegawai          string `json:"nik_pegawai"`
	TempatLahirPegawai  string `json:"-"`
	TanggalLahirPegawai string `json:"-"`
	TTL                 string `json:"ttl"`
	NomorSK             string `json:"nomor_sk"`
	TMT                 string `json:"tmt"`
	TMTIDN              string `json:"tmt_idn"`

	// Jabatan Penetap
	JabatanPenetap organisasiV2.JabatanStruktural `json:"jabatan_penetap" form:"-"`

	// Jenis SK
	IDJenisSK   string `json:"-" form:"-"`
	UUIDJenisSK string `json:"uuid_jenis_sk" form:"uuid_jenis_sk"`
	KDJenisSK   string `json:"kd_jenis_sk" form:"-"`
	JenisSK     string `json:"jenis_sk" form:"-"`

	// Kelompok SK Pengangkatan
	IDKelompokSKPengangkatan   string `json:"-" form:"-"`
	UUIDKelompokSKPengangkatan string `json:"uuid_kelompok_sk_pengangkatan" form:"uuid_kelompok_sk_pengangkatan"`
	KDKelompokSKPengangkatan   string `json:"kd_kelompok_sk_pengangkatan" form:"-"`
	KelompokSKPengangkatan     string `json:"kelompok_sk_pengangkatan" form:"-"`

	// Unit yang mengangkat
	IDUnitPengangkat   string `json:"-" form:"-"`
	UUIDUnitPengangkat string `json:"uuid_unit_pengangkat" form:"uuid_unit_pengangkat"`
	KDUnitPengangkat   string `json:"kd_unit_pengangkat" form:"-"`
	UnitPengangkat     string `json:"unit_pengangkat" form:"-"`

	// Unit kerja
	IDUnitPegawai   string `json:"-" form:"-"`
	UUIDUnitPegawai string `json:"uuid_unit_kerja" form:"uuid_unit_kerja"`
	KDUnitPegawai   string `json:"kd_unit_kerja" form:"-"`
	UnitPegawai     string `json:"unit_kerja" form:"-"`

	// Jabatan fungsional
	IDJabatanFungsional   string `json:"-" form:"-"`
	UUIDJabatanFungsional string `json:"uuid_jabatan_fungsional" form:"uuid_jabatan_fungsional"`
	KDJabatanFungsional   string `json:"kd_jabatan_fungsional" form:"-"`
	JabatanFungsional     string `json:"jabatan_fungsional" form:"-"`

	// Pangkat golongan/ruang
	IDPangkatGolonganPegawai   string `json:"-" form:"-"`
	UUIDPangkatGolonganPegawai string `json:"uuid_pangkat_golongan_pegawai" form:"uuid_pangkat_golongan_pegawai"`
	KDPangkatGolonganPegawai   string `json:"pangkat" form:"-"`
	PangkatGolonganPegawai     string `json:"golongan" form:"-"`

	// Pejabat Penetap
	PejabatPenetap organisasiV2.PejabatStruktural `json:"pejabat_penetap" form:"-"`

	GajiPokok            int `json:"gaji_pokok" form:"gaji_pokok"`
	MasaRilBulan         int `json:"masa_kerja_ril_bulan" form:"masa_kerja_ril_bulan"`
	MasaRilTahun         int `json:"masa_kerja_ril_tahun" form:"masa_kerja_ril_tahun"`
	MasaGajiBulan        int `json:"masa_kerja_gaji_bulan" form:"masa_kerja_gaji_bulan"`
	MasaGajiTahun        int `json:"masa_kerja_gaji_tahun" form:"masa_kerja_gaji_tahun"`
	MasaKerjaDiakuiBulan int `json:"masa_kerja_diakui_bulan" form:"masa_kerja_diakui_bulan"`
	MasaKerjaDiakuiTahun int `json:"masa_kerja_diakui_tahun" form:"masa_kerja_diakui_tahun"`

	// Status pengangkatan
	IDStatusPengangkatan   string `json:"-" form:"-"`
	UUIDStatusPengangkatan string `json:"uuid_status_pengangkatan" form:"uuid_status_pengangkatan"`
	KDStatusPengangkatan   string `json:"kd_status_pengangkatan" form:"-"`
	StatusPengangkatan     string `json:"status_pengangkatan" form:"-"`

	// Ijazah yang diakui
	IDJenisIjazah   string `json:"-" form:"-"`
	UUIDJenisIjazah string `json:"uuid_jenis_ijazah" form:"uuid_jenis_ijazah"`
	KDJenisIjazah   string `json:"kd_jenis_ijazah" form:"-"`
	JenisIjazah     string `json:"jenis_ijazah" form:"-"`

	TanggalDitetapkan    string                `json:"tanggal_ditetapkan" form:"tanggal_ditetapkan"`
	TanggalDitetapkanIDN string                `json:"tanggal_ditetapkan_idn" form:"-"`
	PathSKTendik         string                `json:"-"`
	URLSKTendik          string                `json:"url_sk_pengangkatan" form:"url_sk_pengangkatan"`
	FileSKTendik         *multipart.FileHeader `json:"-" form:"-"`

	IDPegawaiPenetap    string `json:"-" form:"-"`
	UUIDPegawaiPenetap  string `json:"-" form:"-"`
	IDGolonganPegawai   string `json:"-" form:"-"`
	UUIDGolonganPegawai string `json:"-" form:"-"`

	UserInput                string `json:"-" form:"-"`
	UserUpdate               string `json:"-" form:"-"`
	UUID                     string `json:"-" form:"-"`
	IDSKPegawai              string `json:"-" form:"-"`
	UUIDSKPengangkatanTendik string `json:"uuid_sk_pengangkatan_tendik" form:"uuid_sk_pengangkatan_tendik"`
}

const dateFormat = "2006-01-02"

func (sk *SKPengangkatanTendikDetail) SetTTL() {
	var ttls []string
	if sk.TempatLahirPegawai != "" {
		ttls = append(ttls, sk.TempatLahirPegawai)
	}
	_, tglLahir := helper.GetIndonesianDate(dateFormat, sk.TanggalLahirPegawai)
	if tglLahir != "" {
		ttls = append(ttls, tglLahir)
	}
	sk.TTL = strings.Join(ttls, ", ")
}

func (sk *SKPengangkatanTendikDetail) SetTanggalSK() {

	_, sk.TMTIDN = helper.GetIndonesianDate(dateFormat, sk.TMT)
	_, sk.TanggalDitetapkanIDN = helper.GetIndonesianDate(dateFormat, sk.TanggalDitetapkan)
}
