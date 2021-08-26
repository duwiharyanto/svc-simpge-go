package model

type PegawaiCreate struct {
	Id                      uint64                  `form:"id" gorm:"primaryKey"`
	Uuid                    string                  `form:"-"`
	IdPersonalDataPribadi   uint64                  `form:"id_personal_data_pribadi"`
	FlagAktif               int                     `form:"flag_aktif" gorm:"->"`
	Nik                     string                  `form:"nik" gorm:"uniqueIndex"`
	NikKtp                  string                  `form:"nik_ktp"`
	Nama                    string                  `form:"nama"`
	GelarDepan              string                  `form:"gelar_depan"`
	GelarBelakang           string                  `form:"gelar_belakang"`
	TempatLahir             string                  `form:"tempat_lahir"`
	TglLahir                string                  `form:"tgl_lahir"`
	JenisKelamin            string                  `form:"jenis_kelamin"`
	IdAgama                 uint64                  `form:"id_agama"`
	KdAgama                 string                  `form:"kd_agama"`
	IdGolonganDarah         uint64                  `form:"id_golongan_darah"`
	KdGolonganDarah         string                  `form:"kd_golongan_darah"`
	IdStatusPerkawinan      uint64                  `form:"id_status_perkawinan"`
	KdStatusPerkawinan      string                  `form:"kd_status_perkawinan"`
	UuidPendidikanMasuk     string                  `form:"uuid_pendidikan_masuk" gorm:"-"`
	IdPendidikanMasuk       uint64                  `form:"id_pendidikan_masuk"`
	KdPendidikanMasuk       string                  `form:"kd_pendidikan_masuk"`
	IdStatusPendidikanMasuk uint64                  `form:"id_status_pendidikan_masuk"`
	KdStatusPendidikanMasuk string                  `form:"kd_status_pendidikan_masuk"`
	UuidPendidikanTerakhir  string                  `form:"uuid_pendidikan_terakhir" gorm:"-"`
	IdPendidikanTerakhir    uint64                  `form:"id_pendidikan_terakhir"`
	KdPendidikanTerakhir    string                  `form:"kd_pendidikan_terakhir"`
	IdJenisPendidikan       uint64                  `form:"id_jenis_pendidikan"`
	kdJenisPendidikan       string                  `form:"kd_jenis_pendidikan"`
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
	UuidUnitKerja1          string                  `form:"uuid_induk_kerja" gorm:"-"`
	IdUnitKerja1            uint64                  `form:"id_unit_kerja1"`
	KdUnit1                 string                  `form:"kd_unit1"`
	UuidUnitKerja2          string                  `form:"uuid_unit_kerja" gorm:"-"`
	IdUnitKerja2            uint64                  `form:"id_unit_kerja2"`
	KdUnit2                 string                  `form:"kd_unit2"`
	UuidUnitKerja3          string                  `form:"uuid_bagian_kerja" gorm:"-"`
	IdUnitKerja3            uint64                  `form:"id_unit_kerja3"`
	KdUnit3                 string                  `form:"kd_unit3"`
	IdUnitKerjaLokasi       uint64                  `form:"id_unit_kerja_lokasi"`
	LokasiKerja             string                  `form:"lokasi_kerja"`
	UuidLokasiKerja         string                  `form:"uuid_lokasi_kerja" gorm:"-"`
	FlagPensiun             string                  `form:"flag_pensiun" gorm:"->"`
	TglPensiun              string                  `form:"tgl_pensiun" gorm:"->"`
	FlagMeninggal           string                  `form:"flag_meninggal" gorm:"->"`
	TglInput                string                  `form:"tgl_input" gorm:"->"`
	UserInput               string                  `form:"user_input"`
	TglUpdate               string                  `form:"tgl_update" gorm:"->"`
	UserUpdate              string                  `form:"user_update"`
	UuidPersonal            string                  `form:"uuid_personal" gorm:"-"`
	PegawaiFungsional       PegawaiFungsionalCreate `gorm:"foreignKey:Id"`
	PegawaiPNS              PegawaiPNSCreate        `gorm:"foreignKey:Id"`
}

func (*PegawaiCreate) TableName() string {
	return "pegawai"
}

type PegawaiFungsionalCreate struct {
	Id                       uint64  `form:"-" gorm:"primaryKey"`
	IdKafka                  uint64  `form:"-"`
	IdPegawai                uint64  `form:"-"`
	UuidPangkatGolongan      string  `form:"uuid_pangkat_golongan" gorm:"-"`
	IdPangkatGolongan        uint64  `form:"id_pangkat_golongan"`
	KdPangkatGolongan        string  `form:"kd_pangkat_golongan"`
	UuidJabatanFungsional    string  `form:"uuid_jabatan_fungsional" gorm:"-"`
	IdJabatanFungsional      uint64  `form:"id_jabatan_fungsional"`
	KdJabatanFungsional      string  `form:"kd_jabatan_fungsional"`
	TmtPangkatGolongan       *string `form:"tmt_pangkat_golongan"`
	TmtPangkatGolonganIDN    string  `form:"tmt_pangkat_golongan_idn" gorm:"-"`
	TmtJabatan               *string `form:"tmt_jabatan"`
	TmtJabatanIDN            string  `form:"tmt_jabatan_idn" gorm:"-"`
	MasaKerjaBawaanTahun     string  `form:"masa_kerja_bawaan_tahun"`
	MasaKerjaBawaanBulan     string  `form:"masa_kerja_bawaan_bulan"`
	MasaKerjaGajiTahun       string  `form:"masa_kerja_gaji_tahun"`
	MasaKerjaGajiBulan       string  `form:"masa_kerja_gaji_bulan"`
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
	UuidHomebasePddikti      string  `form:"uuid_homebase_pddikti" gorm:"-"`
	IdHomebasePddikti        uint64  `form:"id_homebase_pddikti"`
	UuidHomebaseUii          string  `form:"uuid_homebase_uii" gorm:"-"`
	IdHomebaseUii            uint64  `form:"id_homebase_uii"`
	TglInput                 string  `form:"-" gorm:"-"`
	UserInput                string  `form:"-"`
	TglUpdate                string  `form:"-" gorm:"-"`
	UserUpdate               string  `form:"-"`
	FlagAktif                uint64  `form:"-" gorm:"-"`
}

func (*PegawaiFungsionalCreate) TableName() string {
	return "pegawai_fungsional"
}

type PegawaiPNSCreate struct {
	Id                    uint64  `form:"-" gorm:"primaryKey"`
	IdPegawai             uint64  `form:"-"`
	NipPns                string  `form:"nip_pns"`
	NoKartuPegawai        string  `form:"no_kartu_pegawai"`
	UuidJenisPtt          string  `form:"uuid_jenis_ptt" gorm:"-"`
	IdJenisPtt            uint64  `form:"id_jenis_ptt"`
	KdJenisPtt            string  `form:"kd_jenis_ptt"`
	InstansiAsal          string  `form:"instansi_asal_ptt" gorm:"column:instansi_asal"`
	UuidPangkatGolongan   string  `form:"uuid_pangkat_gol_ruang_pns" gorm:"-"`
	IdPangkatGolongan     uint64  `form:"id_pangkat_golongan"`
	KdPangkatGolongan     string  `form:"kd_pangkat_golongan"`
	TmtPangkatGolongan    *string `form:"tmt_pangkat_gol_ruang_pns" gorm:"column:tmt_pangkat_golongan"`
	TmtPangkatGolonganIDN string  `form:"tmt_pangkat_gol_ruang_pns_idn" gorm:"-"`
	UuidJabatanFungsional string  `form:"uuid_jabatan_pns" gorm:"-"`
	IdJabatanFungsional   uint64  `form:"id_jabatan_fungsional"`
	KdJabatanFungsional   string  `form:"kd_jabatan_fungsional"`
	TmtJabatan            *string `form:"tmt_jabatan_pns" gorm:"tmt_jabatan"`
	TmtJabatanIDN         string  `form:"tmt_jabatan_pns_idn" gorm:"-"`
	MasaKerjaTahun        string  `form:"masa_kerja_pns_tahun" gorm:"column:masa_kerja_tahun"`
	MasaKerjaBulan        string  `form:"masa_kerja_pns_bulan" gorm:"column:masa_kerja_bulan"`
	AngkaKredit           string  `form:"angka_kredit_pns" gorm:"column:angka_kredit"`
	Keterangan            string  `form:"keterangan_pns" gorm:"column:keterangan"`
	TglInput              string  `form:"-" gorm:"-"`
	UserInput             string  `form:"-"`
	TglUpdate             string  `form:"-" gorm:"-"`
	UserUpdate            string  `form:"-"`
	FlagAktif             uint64  `form:"-" gorm:"-"`
}

func (*PegawaiPNSCreate) TableName() string {
	return "pegawai_pns"
}
