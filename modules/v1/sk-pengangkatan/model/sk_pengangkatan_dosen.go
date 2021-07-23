package model

import "mime/multipart"

func EmptySKPengangkatanDosen() SKPengangkatanTendik {
	return SKPengangkatanTendik{}
}

type SKPengangkatanDosen struct {
	IDJenisSKPengangkatan          uint64                `json:"id_jenis_sk_pengangkatan" form:"id_jenis_sk_pengangkatan"`
	UUIDJenisSKPengangkatan        string                `json:"uuid_jenis_sk_pengangkatan" form:"uuid_jenis_sk_pengangkatan"`
	IDKelompokPegawai              uint64                `json:"id_kelompok_pegawai" form:"id_kelompok_pegawai"`
	UUIDKelompokPegawai            string                `json:"uuid_kelompok_pegawai" form:"uuid_kelompok_pegawai"`
	MasaKerjaDiakuiBulanLama       string                `json:"masa_kerja_diakui_bulan_lama" form:"masa_kerja_diakui_bulan_lama"`
	MasaKerjaDiakuiTahunBaru       string                `json:"masa_kerja_diakui_tahun_baru" form:"masa_kerja_diakui_tahun_baru"`
	IDPangkatGolonganPegawaiLama   uint64                `json:"id_pangkat_gol_lama" form:"id_pangkat_gol_lama"`
	UUIDPangkatGolonganPegawaiLama string                `json:"uuid_pangkat_gol_lama" form:"uuid_pangkat_gol_lama"`
	IDJabatanFungsionalLama        uint64                `json:"id_jabatan_fungsional_lama" form:"id_jabatan_fungsional_lama"`
	UUIDJabatanFungsionalLama      string                `json:"uuid_jabatan_fungsional_lama" form:"uuid_jabatan_fungsional_lama"`
	IDMataKuliah                   []string              `json:"-" form:"-"`
	UUIDMataKuliahStr              string                `json:"-" form:"uuid_mata_kuliah"`
	UUIDMataKuliah                 []string              `json:"-" form:"-"`
	IDIndukKerjaBaru               uint64                `json:"id_induk_kerja_baru" form:"id_induk_kerja_baru"`
	UUIDIndukKerjaBaru             string                `json:"uuid_induk_kerja_baru" form:"uuid_induk_kerja_baru"`
	IDPangkatGolonganPegawaiBaru   uint64                `json:"id_pangkat_gol_baru" form:"id_pangkat_gol_baru"`
	UUIDPangkatGolonganPegawaiBaru string                `json:"uuid_pangkat_gol_baru" form:"uuid_pangkat_gol_baru"`
	GajiPokok                      string                `json:"gaji_pokok" form:"gaji_pokok"`
	TunjanganBeras                 string                `json:"tunjangan_beras" form:"tunjangan_beras"`
	TunjanganKhusus                string                `json:"tunjangan_khusus" form:"tunjangan_khusus"`
	SKSMengajar                    string                `json:"sks_mengajar" form:"sks_mengajar"`
	BantuanKomunikasi              string                `json:"bantuan_komunikasi" form:"bantuan_komunikasi"`
	TunjanganTahunan               string                `json:"tunjangan_tahunan" form:"tunjangan_tahunan"`
	MasaKerjaRilBulanBaru          string                `json:"masa_kerja_riil_bulan_baru" form:"masa_kerja_riil_bulan_baru"`
	MasaKerjaRilTahunBaru          string                `json:"masa_kerja_riil_tahun_baru" form:"masa_kerja_riil_tahun_baru"`
	MasaKerjaGajiBulanBaru         string                `json:"masa_kerja_gaji_bulan_baru" form:"masa_kerja_gaji_bulan_baru"`
	MasaKerjaGajiTahunBaru         string                `json:"masa_kerja_gaji_tahun_baru" form:"masa_kerja_gaji_tahun_baru"`
	IDJabatanPenetap               uint64                `json:"id_jabatan_penetap" form:"id_jabatan_penetap"`
	UUIDJabatanPenetap             string                `json:"uuid_jabatan_penetap" form:"uuid_jabatan_penetap"`
	IDPegawaiPenetap               uint64                `json:"id_pegawai_penetap" form:"id_pegawai_penetap"`
	UUIDPegawaiPenetap             string                `json:"uuid_pegawai_penetap" form:"uuid_pegawai_penetap"`
	TanggalDitetapkan              string                `json:"tgl_ditetapkan" form:"tgl_ditetapkan"`
	IDUnitKerja                    uint64                `json:"id_unit_kerja" form:"id_unit_kerja"`
	UUIDUnitKerja                  string                `json:"uuid_unit_kerja" form:"uuid_unit_kerja"`
	IDJabatanFungsionalBaru        uint64                `json:"id_jabatan_fungsional" form:"id_jabatan_fungsional"`
	UUIDJabatanFungsionalBaru      string                `json:"uuid_jabatan_fungsional" form:"uuid_jabatan_fungsional"`
	IDJenisIjazah                  uint64                `json:"id_ijazah_pendidikan" form:"id_ijazah_pendidikan"`
	UUIDJenisIjazah                string                `json:"uuid_ijazah_pendidikan" form:"uuid_ijazah_pendidikan"`
	InstansiKerja                  string                `json:"instansi_kerja" form:"instansi_kerja"`
	TglBerakhir                    string                `json:"tgl_berakhir" form:"tgl_berakhir"`
	JangkaWaktuEvaluasi            string                `json:"jangka_waktu_evaluasi" form:"jangka_waktu_evaluasi"`
	PathSKDosen                    string                `json:"-"`
	URLSKDosen                     string                `json:"url_sk_pengangkatan" form:"url_sk_pengangkatan"`
	FileSKDosen                    *multipart.FileHeader `form:"-"`
	UserInput                      string                `json:"user_input" form:"user_input"`
	UserUpdate                     string                `json:"user_update" form:"user_update"`
	UUID                           string                `json:"uuid" form:"uuid"`
	IDSKPegawai                    uint64                `json:"id_sk_pegawai" form:"id_sk_pegawai"`
	UUIDSKPengangkatanDosen        string                `json:"uuid_sk_pengangkatan_dosen" form:"uuid_sk_pengangkatan_dosen"`
}
type PivotMakulDosen struct {
	IDMakulSKPengangkatan   uint64 `json:"id_makul_skpengangkatan" form:"id_makul_skpengangkatan"`
	UUIDMakulSKPengangkatan string `json:"uuid_makul_skpengangkatan" form:"uuid_makul_skpengangkatan"`
	IDSKPengangkatan        uint64 `json:"id_sk_pengangkatan" form:"id_sk_pengangkatan"`
}

func (sk SKPengangkatanDosen) IsEmpty() bool {
	return sk.IDJenisSKPengangkatan == 0
}

type SKPengangkatanDosenResponse struct {
	Data []SKPengangkatanDosen `json:"data"`
}
