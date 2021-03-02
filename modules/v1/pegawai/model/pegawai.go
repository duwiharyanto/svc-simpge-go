package model

import (
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
