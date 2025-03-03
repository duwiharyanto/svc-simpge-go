package model

type PegawaiPNSUpdate struct {
	Id                     *uint64 `form:"-" gorm:"primaryKey"`
	IdPegawai              *uint64 `form:"-"`
	NipPns                 *string `form:"nip_pns" gorm:"default:null"`
	NoKartuPegawai         *string `form:"no_kartu_pegawai" gorm:"default:null"`
	UuidJenisPtt           *string `form:"uuid_jenis_ptt" gorm:"-"`
	IdJenisPtt             *uint64 `form:"id_jenis_ptt" gorm:"default:null"`
	KdJenisPtt             *string `form:"kd_jenis_ptt" gorm:"default:null"`
	InstansiAsal           *string `form:"instansi_asal_ptt" gorm:"column:instansi_asal"`
	UuidPangkatGolongan    *string `form:"uuid_pangkat_gol_ruang_pns" gorm:"-"` //Perubahan
	IdPangkatGolongan      *uint64 `form:"id_pangkat_golongan" gorm:"default:null"`
	KdPangkatGolongan      *string `form:"kd_pangkat_golongan" gorm:"default:null"`
	TmtPangkatGolongan     *string `form:"tmt_pangkat_gol_ruang_pns" gorm:"column:tmt_pangkat_golongan;default:null"` //Perubahan
	TmtPangkatGolonganIDN  *string `form:"tmt_pangkat_gol_ruang_pns_idn" gorm:"-"`                                    //Perubahan
	NomorPangkatGolongan   *string `form:"nomor_pangkat_golongan" gorm:"column:nomor_pangkat_golongan;default:null"`  //Tambahan
	UuidJabatanFungsional  *string `form:"uuid_jabatan_pns" gorm:"-"`                                                 //Perubahan
	IdJabatanFungsional    *uint64 `form:"id_jabatan_fungsional" gorm:"default:null"`
	KdJabatanFungsional    *string `form:"kd_jabatan_fungsional" gorm:"default:null"`
	TmtJabatan             *string `form:"tmt_jabatan_pns" gorm:"tmt_jabatan;default:null"`                                //Perubahan
	TmtJabatanIDN          *string `form:"tmt_jabatan_pns_idn" gorm:"-"`                                                   //Perubahan
	NomorJabatanFungsional *string `form:"nomor_jabatan_fungsional" gorm:"column:nomor_jabatan_fungsional;default:null"`   //Tambahan
	MasaKerjaTahun         *string `form:"masa_kerja_pns_tahun" gorm:"column:masa_kerja_tahun;default:null"`               //Perubahan
	MasaKerjaBulan         *string `form:"masa_kerja_pns_bulan" gorm:"column:masa_kerja_bulan;default:null"`               //Perubahan
	MasaKerjaGolonganTahun *string `form:"masa_kerja_golongan_tahun" gorm:"column:masa_kerja_golongan_tahun;default:null"` //Tambahan
	MasaKerjaGolonganBulan *string `form:"masa_kerja_golongan_bulan" gorm:"column:masa_kerja_golongan_bulan;default:null"` //Tambahan
	AngkaKredit            *string `form:"angka_kredit_pns" gorm:"column:angka_kredit;default:null"`                       //Perubahan
	NomorPak               *string `form:"nomor_pak" gorm:"column:nomor_pak;default:null"`                                 //Tambahan
	TmtPak                 *string `form:"tmt_pak" gorm:"tmt_pak;default:null"`                                            //Tambahan
	NomorSkPensiun         *string `form:"nomor_sk_pensiun" gorm:"column:nomor_sk_pensiun;default:null"`                   //Tambahan
	TmtSkPensiun           *string `form:"tmt_sk_pensiun" gorm:"tmt_sk_pensiun;default:null"`                              //Tambahan
	MasaKerjaPensiunTahun  *string `form:"masa_kerja_pensiun_tahun" gorm:"masa_kerja_pensiun_tahun;default:null"`          //Tambahan
	MasaKerjaPensiunBulan  *string `form:"masa_kerja_pensiun_bulan" gorm:"masa_kerja_pensiun_bulan;default:null"`          //Tambahan
	Keterangan             *string `form:"keterangan_pns" gorm:"column:keterangan;default:null"`                           //Perubahan
	Nira                   *string `form:"nira" gorm:"column:nira;default:null"`                                           //Perubahan
	TglInput               *string `form:"-" gorm:"-"`
	UserInput              *string `form:"-" gorm:"-"`
	TglUpdate              *string `form:"-" gorm:"-"`
	UserUpdate             *string `form:"-"`
	FlagAktif              *uint64 `form:"-" gorm:"-"`
	Uuid                   *string `form:"-" gorm:"-"`
}

func (*PegawaiPNSUpdate) TableName() string {
	return "pegawai_pns"
}

type PegawaiPNSPTT struct {
	NipPNS                 string `json:"nip_pns"`
	NoKartuPegawai         string `json:"no_kartu_pegawai"`
	UuidDetailProfesi      string `json:"uuid_detail_profesi" form:"uuid_detail_profesi"`
	IdDetailProfesi        int    `json:"-"`
	DetailProfesi          string `json:"detail_profesi"`
	UuidPangkatGolongan    string `json:"uuid_pangkat_gol_ruang_pns"`
	KdPangkatGolonganPns   string `json:"kd_pangkat_golongan"`
	PangkatPNS             string `json:"pangkat_pns"`
	GolonganPNS            string `json:"golongan_pns"`
	KdGolonganPNS          string `json:"kd_golongan_pns"`
	KdRuangPNS             string `json:"kd_ruang_pns"`
	TmtPangkatGolongan     string `json:"tmt_pangkat_gol_ruang_pns"`
	TmtPangkatGolonganIdn  string `json:"tmt_pangkat_gol_ruang_pns_idn"`
	NomorPangkatGolongan   string `json:"nomor_pangkat_golongan"`
	KdJabatanPns           string `json:"kd_jabatan_pns"`
	JabatanPns             string `json:"jabatan_pns"`
	UuidJabatanPns         string `json:"uuid_jabatan_pns"`
	TmtJabatanPns          string `json:"tmt_jabatan_pns"`
	TmtJabatanPnsIdn       string `json:"tmt_jabatan_pns_idn"`
	NomorJabatanFungsional string `json:"nomor_jabatan_fungsional"`
	MasaKerjaPnsTahun      string `json:"masa_kerja_pns_tahun"`
	MasaKerjaPnsBulan      string `json:"masa_kerja_pns_bulan"`
	MasaKerjaGolonganTahun string `json:"masa_kerja_golongan_tahun"`
	MasaKerjaGolonganBulan string `json:"masa_kerja_golongan_bulan"`
	AngkaKreditPns         string `json:"angka_kredit_pns"`
	NomorPak               string `json:"nomor_pak"`
	TmtPak                 string `json:"tmt_pak"`
	NomorSkPensiun         string `json:"nomor_sk_pensiun"`
	TmtSkPensiun           string `json:"tmt_sk_pensiun"`
	MasaKerjaPensiunTahun  string `json:"masa_kerja_pensiun_tahun"`
	MasaKerjaPensiunBulan  string `json:"masa_kerja_pensiun_bulan"`
	KeteranganPNS          string `json:"keterangan_pns"`
	Nira                   string `json:"nira"`
	UuidJenisPTT           string `json:"uuid_jenis_ptt"`
	KdJenisPTT             string `json:"kd_jenis_ptt"`
	JenisPTT               string `json:"jenis_ptt"`
	InstansiAsalPtt        string `json:"instansi_asal_ptt"`
}

type PegawaiPNSCreate struct {
	Id                     uint64  `form:"-" gorm:"primaryKey"`
	IdPegawai              uint64  `form:"-"`
	NipPns                 string  `form:"nip_pns" gorm:"default:null"`
	NoKartuPegawai         string  `form:"no_kartu_pegawai" gorm:"default:null"`
	UuidJenisPtt           string  `form:"uuid_jenis_ptt" gorm:"-"`
	IdJenisPtt             uint64  `form:"id_jenis_ptt" gorm:"default:null"`
	KdJenisPtt             string  `form:"kd_jenis_ptt" gorm:"default:null"`
	InstansiAsal           string  `form:"instansi_asal_ptt" gorm:"column:instansi_asal;default:null"`
	UuidPangkatGolongan    string  `form:"uuid_pangkat_gol_ruang_pns" gorm:"-"`
	IdPangkatGolongan      uint64  `form:"id_pangkat_golongan" gorm:"default:null"`
	KdPangkatGolongan      string  `form:"kd_pangkat_golongan" gorm:"default:null"`
	TmtPangkatGolongan     *string `form:"tmt_pangkat_gol_ruang_pns" gorm:"column:tmt_pangkat_golongan;default:null"`
	TmtPangkatGolonganIDN  string  `form:"tmt_pangkat_gol_ruang_pns_idn" gorm:"-"`
	NomorPangkatGolongan   string  `form:"nomor_pangkat_golongan" gorm:"column:nomor_pangkat_golongan;default:null"`
	UuidJabatanFungsional  string  `form:"uuid_jabatan_pns" gorm:"-"`
	IdJabatanFungsional    uint64  `form:"id_jabatan_fungsional" gorm:"default:null"`
	KdJabatanFungsional    string  `form:"kd_jabatan_fungsional" gorm:"default:null"`
	TmtJabatan             *string `form:"tmt_jabatan_pns" gorm:"tmt_jabatan;default:null"`
	TmtJabatanIDN          string  `form:"tmt_jabatan_pns_idn" gorm:"-"`
	NomorJabatanFungsional string  `form:"nomor_jabatan_fungsional" gorm:"column:nomor_jabatan_fungsional;default:null"`
	MasaKerjaTahun         string  `form:"masa_kerja_pns_tahun" gorm:"column:masa_kerja_tahun;default:null"`
	MasaKerjaBulan         string  `form:"masa_kerja_pns_bulan" gorm:"column:masa_kerja_bulan;default:null"`
	AngkaKredit            string  `form:"angka_kredit_pns" gorm:"column:angka_kredit;default:null"`
	Keterangan             string  `form:"keterangan_pns" gorm:"column:keterangan;default:null"`
	TglInput               string  `form:"-" gorm:"-"`
	UserInput              string  `form:"-"`
	TglUpdate              string  `form:"-" gorm:"-"`
	UserUpdate             string  `form:"-"`
	FlagAktif              uint64  `form:"-" gorm:"-"`
}

func (*PegawaiPNSCreate) TableName() string {
	return "pegawai_pns"
}
