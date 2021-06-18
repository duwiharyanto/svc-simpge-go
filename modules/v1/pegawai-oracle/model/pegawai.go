package model

type PegawaiSimpeg struct {
	NIP         string `json:"-"`
	Nama        string `json:"nama,omitempty"`
	KdKelamin   string `json:"kd_kelamin,omitempty"`
	KdNikah     string `json:"kd_nikah,omitempty"`
	JmlKeluarga int    `json:"jml_keluarga,omitempty"`
	UserUpdate  string `json:"user_update,omitempty"`
	TglUpdate   string `json:"tgl_update,omitempty"`
}

type GetPegawaiSimpegResult struct {
	Result bool          `json:"result"`
	Data   PegawaiSimpeg `json:"data"`
}

type PegawaiStatusSimpeg struct {
	NIP           string  `json:"-"`
	FlagSekantor  string  `json:"flag_sekantor"`
	NipSuamiIstri *string `json:"nip_suami_istri"`
	UserUpdate    string  `json:"user_update,omitempty"`
	TglUpdate     string  `json:"tgl_update,omitempty"`
}

type GetPegawaiStatusSimpegResult struct {
	Result bool                `json:"result"`
	Data   PegawaiStatusSimpeg `json:"data"`
}

type KepegawaianYayasanSimpeg struct {
	NIP               string `json:"-"`
	FlagPensiun       string `json:"flag_pensiun"`
	KdStatusHidup     string `json:"kd_status_hidup"`
	UserUpdate        string `json:"user_update,omitempty"`
	TglUpdate         string `json:"tgl_update,omitempty"`
	NipKopertis       string `json:"nip_kopertis"`
	KdPendidikanMasuk string `json:"kd_pendidikan_masuk"`
	KdPendidikan      string `json:"kd_pendidikan"`
	*InstansiAsalPtt  `json:"instansi_asal_ptt"`
	*JenisPegawai     `json:"jenis_pegawai,omitempty"`
	*KelompokPegawai  `json:"kelompok_pegawai,omitempty"`
	*LokasiKerja      `json:"lokasi_kerja,omitempty"`
	*StatusPegawai    `json:"status_pegawai,omitempty"`
	*PegawaiStatus    `json:"pegawai_status,omitempty"`
	*Unit1            `json:"unit1,omitempty"`
	*Unit2            `json:"unit2,omitempty"`
	*Unit3            `json:"unit3,omitempty"`
}

type InstansiAsalPtt struct {
	Instansi   string `json:"instansi"`
	Keterangan string `json:"keterangan"`
}

type JenisPegawai struct {
	KdJenisPegawai string `json:"kd_jenis_pegawai,omitempty"`
	JenisPegawai   string `json:"jenis_pegawai,omitempty"`
}

type KelompokPegawai struct {
	KdKelompokPegawai string `json:"kd_kelompok_pegawai,omitempty"`
	KelompokPegawai   string `json:"kelompok_pegawai,omitempty"`
}

type LokasiKerja struct {
	KdLokasi string `json:"kd_lokasi,omitempty"`
	Lokasi   string `json:"lokasi,omitempty"`
}

type Unit1 struct {
	KdUnit1 string `json:"kd_unit1,omitempty"`
	Unit1   string `json:"unit1,omitempty"`
}

type Unit2 struct {
	KdUnit1 string `json:"kd_unit1,omitempty"`
	KdUnit2 string `json:"kd_unit2,omitempty"`
	Unit2   string `json:"unit2,omitempty"`
}

type Unit3 struct {
	KdUnit2 string `json:"kd_unit2,omitempty"`
	KdUnit3 string `json:"kd_unit3,omitempty"`
	Unit3   string `json:"unit3,omitempty"`
}

type StatusPegawai struct {
	KdStatusPegawai string `json:"kd_status_pegawai,omitempty"`
	StatusPegawai   string `json:"status_pegawai,omitempty"`
}

type PegawaiStatus struct {
	AngkaKreditFungsional     float64            `json:"ak_fungsional"`
	AngkaKreditKopertis       float64            `json:"ak_fungsional_kopertis"`
	FlagMengajar              string             `json:"flag_mengajar"`
	FlagSekolah               string             `json:"flag_sekolah"`
	KdHomebasePddikti         string             `json:"kd_homebase_pddikti,omitempty"`
	KdHomebaseUii             string             `json:"kd_homebase_uii,omitempty"`
	MasaKerjaGajiTahun        int                `json:"masa_kerja_gaji_tahun"`
	MasaKerjaGajiBulan        int                `json:"masa_kerja_gaji_bulan"`
	MasaKerjaKopertisTahun    int                `json:"masa_kerja_kopertis_tahun"`
	MasaKerjaKopertisBulan    int                `json:"masa_kerja_kopertis_bulan"`
	NoKarpeg                  string             `json:"no_karpeg"`
	NoSkPertama               string             `json:"no_sk_pertama,omitempty"`
	TglSkPertama              string             `json:"tgl_sk_pertama,omitempty"`
	PangkatKopertis           *Pangkat           `json:"pangkat_kopertis,omitempty"`
	PangkatYayasan            *Pangkat           `json:"pangkat_yayasan,omitempty"`
	JabatanFungsionalKopertis *JabatanFungsional `json:"jabatan_fungsional_kopertis,omitempty"`
	*JabatanFungsional        `json:"jabatan_fungsional,omitempty"`
}

type Pangkat struct {
	KdGolongan string `json:"kd_golongan,omitempty"`
	KdRuang    string `json:"kd_ruang,omitempty"`
	TmtPangkat string `json:"tmt_pangkat"`
}

type JabatanFungsional struct {
	KdFungsional  string `json:"kd_fungsional,omitempty"`
	TmtFungsional string `json:"tmt_fungsional"`
}

type GetKepegawaianYayasanSimpegResult struct {
	Result bool                     `json:"result"`
	Data   KepegawaianYayasanSimpeg `json:"data"`
}
