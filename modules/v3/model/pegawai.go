package model

type Pegawai struct {
	RowId                       uint64 `json:"-" form:"-"`
	Id                          uint64 `json:"-" form:"-"`
	IdPersonalDataPribadi       uint64 `json:"-" form:"-"`
	FlagAktif                   int    `json:"-" form:"-"`
	Nik                         string `json:"-" form:"-"`
	NikKtp                      string `json:"-" form:"-"`
	Nama                        string `json:"-" form:"-"`
	GelarDepan                  string `json:"-" form:"-"`
	GelarBelakang               string `json:"-" form:"-"`
	TempatLahir                 string `json:"-" form:"-"`
	TglLahir                    string `json:"-" form:"-"`
	JenisKelamin                string `json:"-" form:"-"`
	IdAgama                     uint64 `json:"-" form:"-"`
	KdAgama                     string `json:"-" form:"-"`
	IdGolonganDarah             uint64 `json:"-" form:"-"`
	KdGolonganDarah             string `json:"-" form:"-"`
	IdStatusPerkawinan          uint64 `json:"-" form:"-"`
	KdStatusPerkawinan          string `json:"-" form:"-"`
	IdPendidikanMasuk           uint64 `json:"-" form:"-"`
	KdPendidikanMasuk           string `json:"-" form:"-"`
	IdStatusPendidikanMasuk     uint64 `json:"-" form:"-"`
	KdStatusPendidikanMasuk     string `json:"-" form:"-"`
	IdPendidikanTerakhir        uint64 `json:"-" form:"-"`
	KdPendidikanTerakhir        string `json:"-" form:"-"`
	IdJenisPendidikan           uint64 `json:"-" form:"-"`
	KdJenisPendidikan           string `json:"-" form:"-"`
	IdJenisPegawai              uint64 `json:"-" form:"-"`
	KdJenisPegawai              string `json:"-" form:"-"`
	IdStatusPegawai             uint64 `json:"-" form:"-"`
	KdStatusPegawai             string `json:"-" form:"-"`
	IdKelompokPegawai           uint64 `json:"-" form:"-"`
	KdKelompokPegawai           string `json:"-" form:"-"`
	IdGolongan                  uint64 `json:"-" form:"-"`
	KdGolongan                  string `json:"-" form:"-"`
	IdRuang                     uint64 `json:"-" form:"-"`
	KdRuang                     string `json:"-" form:"-"`
	IdUnitKerja1                uint64 `json:"-" form:"-"`
	KdUnit1                     string `json:"-" form:"-"`
	IdUnitKerja2                uint64 `json:"-" form:"-"`
	KdUnit2                     string `json:"-" form:"-"`
	IdUnitKerja3                uint64 `json:"-" form:"-"`
	KdUnit3                     string `json:"-" form:"-"`
	IdUnitKerjaLokasi           uint64 `json:"-" form:"-"`
	LokasiKerja                 string `json:"-" form:"-"`
	DetailProfesi               string `json:"-" form:"-"`
	IdDetailProfesi             uint64 `json:"-" form:"-"`
	FlagPensiun                 int    `json:"-" form:"-"`
	TglPensiun                  string `json:"-" form:"-"`
	FlagMeninggal               int    `json:"-" form:"-"`
	IdJabatanFungsional         uint64 `json:"-" form:"-"`
	TmtPangkat                  string `json:"-" form:"-"`
	TmtJabatan                  string `json:"-" form:"-"`
	MasaKerjaBawaanTahun        string `json:"-" form:"-"`
	MasaKerjaBawaanBulan        string `json:"-" form:"-"`
	MasaKerjaGajiTahun          string `json:"-" form:"-"`
	MasaKerjaGajiBulan          string `json:"-" form:"-"`
	AngkaKredit                 string `json:"-" form:"-"`
	NomorSertifikasi            string `json:"-" form:"-"`
	IdJenisNomorRegistrasi      uint64 `json:"-" form:"-"`
	NomorRegistrasi             string `json:"-" form:"-"`
	NomorSkPertama              string `json:"-" form:"-"`
	TmtSkPertama                string `json:"-" form:"-"`
	IdHomebasePddikti           uint64 `json:"-" form:"-"`
	IdHomebaseUii               uint64 `json:"-" form:"-"`
	IdStatusPegawaiAktif        uint64 `json:"-" form:"-"`
	NipPns                      string `json:"-" form:"-"`
	NoKartuPegawai              string `json:"-" form:"-"`
	IdJenisPtt                  uint64 `json:"-" form:"-"`
	InstansiAsal                string `json:"-" form:"-"`
	IdGolonganKopertis          uint64 `json:"-" form:"-"`
	IdRuangKopertis             uint64 `json:"-" form:"-"`
	TmtPangkatKopertis          string `json:"-" form:"-"`
	IdJabatanFungsionalKopertis uint64 `json:"-" form:"-"`
	TmtJabatanKopertis          string `json:"-" form:"-"`
	MasaKerjaTahunKopertis      string `json:"-" form:"-"`
	MasaKerjaBulanKopertis      string `json:"-" form:"-"`
	AngkaKreditKopertis         string `json:"-" form:"-"`
	Keterangan                  string `json:"-" form:"-"`
	TglInput                    string `json:"-" form:"-"`
	UserInput                   string `json:"-" form:"-"`
	TglUpdate                   string `json:"-" form:"-"`
	UserUpdate                  string `json:"-" form:"-"`
	Uuid                        string `json:"-" form:"-"`

	StatusPegawaiAktif *StatusPegawaiAktif `gorm:"foreignKey:IdStatusPegawaiAktif"`
}

func (*Pegawai) TableName() string {
	return "pegawai_v2"
}

type KelompokPegawai struct {
	Id                uint64
	KdKelompokPegawai string
	KdStatusPegawai   string
	KdJenisPegawai    string
	KelompokPegawai   string
	FlagAktif         int
	Uuid              string

	JenisPegawai  *JenisPegawai  `gorm:"foreignKey:KdJenisPegawai;references:KdJenisPegawai"`
	StatusPegawai *StatusPegawai `gorm:"foreignKey:KdStatusPegawai;references:KdStatusPegawai"`
}

type JenisPegawai struct {
	Id               uint64
	KdJenisPegawai   string
	NamaJenisPegawai string
	FlagAktif        string
	Uuid             string
}

type StatusPegawai struct {
	Id              uint64
	KdStatusPegawai string
	StatusPegawai   string
	Keterangan      string
	FlagAktif       int
	Uuid            string
}

type StatusPegawaiAktif struct {
	ID              uint64
	KdStatus        string
	Status          string
	FlagStatusAktif int
	UUID            string
}

func (s StatusPegawaiAktif) IsActive() bool {
	return s.FlagStatusAktif == 1
}
