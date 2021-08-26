package model

import (
	"time"

	"github.com/cstockton/go-conv"
)

type PegawaiYayasan struct {
	ID                         uint64 `json:"-" gorm:"primaryKey;not null"`
	IDJenisPegawai             uint64 `json:"-"`
	KDJenisPegawai             string `json:"kd_jenis_pegawai"`
	UuidJenisPegawai           string `json:"uuid_jenis_pegawai"`
	JenisPegawai               string `json:"jenis_pegawai"`
	UuidKelompokPegawai        string `json:"uuid_kelompok_pegawai"`
	IdKelompokPegawai          uint64 `json:"-"`
	KdKelompokPegawai          string `json:"kd_kelompok_pegawai"`
	KelompokPegawai            string `json:"kelompok_pegawai"`
	UuidPendidikanMasuk        string `json:"-" form:"uuid_pendidikan_masuk"`
	IdPendidikanMasuk          uint64 `json:"-"`
	KdPendidikanMasuk          string `json:"-"`
	IdPendidikanMasukSimpeg    uint64 `json:"-"`
	KdPendidikanMasukSimpeg    string `json:"-"`
	PendidikanMasuk            string `json:"-"`
	UuidPendidikanTerakhir     string `json:"-" form:"uuid_pendidikan_terakhir"`
	IdPendidikanTerakhir       uint64 `json:"-"`
	KdPendidikanTerakhir       string `json:"-"`
	IdPendidikanTerakhirSimpeg uint64 `json:"-"`
	KdPendidikanTerakhirSimpeg string `json:"-"`
	PendidikanTerakhir         string `json:"-"`
	UuidStatusPegawai          string `json:"uuid_status_pegawai"`
	IDStatusPegawai            uint64 `json:"-"`
	KDStatusPegawai            string `json:"kd_status_pegawai"`
	StatusPegawai              string `json:"status_pegawai"`
	UuidPangkatGolongan        string `json:"uuid_pangkat_golongan"`
	IdPangkat                  uint64 `json:"-"`
	KdPangkat                  string `json:"kd_pangkat_golongan"`
	Pangkat                    string `json:"pangkat"`
	Golongan                   string `json:"golongan"`
	IdGolongan                 uint64 `json:"-"`
	KdGolongan                 string `json:"kd_golongan"`
	IdRuang                    uint64 `json:"-"`
	KdRuang                    string `json:"kd_ruang"`
	TmtPangkatGolongan         string `json:"tmt_pangkat_gol_ruang_pegawai"`
	TmtPangkatGolonganIdn      string `json:"tmt_pangkat_gol_ruang_pegawai_idn"`
	UuidJabatanFungsional      string `json:"uuid_jabatan_fungsional"`
	IdJabatanFungsional        uint64 `json:"-"`
	KdJabatanFungsional        string `json:"kd_jabatan_fungsional"`
	JabatanFungsional          string `json:"jabatan_fungsional"`
	TmtJabatan                 string `json:"tmt_jabatan"`
	TmtJabatanIdn              string `json:"tmt_jabatan_idn"`
	MasaKerjaBawaanTahun       string `json:"masa_kerja_bawaan_tahun"`
	MasaKerjaBawaanBulan       string `json:"masa_kerja_bawaan_bulan"`
	MasaKerjaGajiTahun         string `json:"masa_kerja_gaji_tahun"`
	MasaKerjaGajiBulan         string `json:"masa_kerja_gaji_bulan"`
	MasaKerjaTotalTahun        string `json:"masa_kerja_total_tahun"`
	MasaKerjaTotalBulan        string `json:"masa_kerja_total_bulan"`
	AngkaKredit                string `json:"angka_kredit"`
	NoSertifikasi              string `json:"nomor_sertifikasi_pegawai"`
	UuidJenisRegis             string `json:"uuid_jenis_regis"`
	IdJenisRegis               uint64 `json:"-"`
	KdJenisRegis               string `json:"kd_jenis_regis"`
	JenisNomorRegis            string `json:"jenis_no_regis"`
	NomorRegis                 string `json:"no_regis"`
}

func (*PegawaiYayasan) TableName() string {
	return "pegawai"
}

func (a *PegawaiYayasan) SetMasaKerjaTotal(b *UnitKerjaPegawai) {
	now := time.Now()
	yearNow, _ := conv.Int(now.Year())
	monthNow, _ := conv.Int(now.Month())
	dayNow, _ := conv.Int(now.Day())
	maxMonth := 12

	var yearSk int
	var monthSk int
	var daySk int

	if b.TmtSkPertama != "" {
		dateSk, err := time.Parse("2006-01-02", b.TmtSkPertama)
		if err != nil {
			return
		}
		yearSk, _ = conv.Int(dateSk.Year())
		monthSk, _ = conv.Int(dateSk.Month())
		daySk, _ = conv.Int(dateSk.Day())
	}

	yearSkNow := yearNow - yearSk
	monthSkNow := monthNow - monthSk

	if monthSkNow < 0 {
		yearSkNow = yearSkNow - 1
		monthSkNow = maxMonth + monthSkNow
	}

	daySkNow := dayNow - daySk

	if daySkNow < 0 {
		monthSkNow = monthSkNow - 1
	}

	// fmt.Println("Year Sk Now : ", yearSkNow)
	// fmt.Println("Month Sk Now : ", monthSkNow)

	var MasaBawaanBulan, _ = conv.Int(a.MasaKerjaBawaanBulan)
	var MasaBawaanTahun, _ = conv.Int(a.MasaKerjaBawaanTahun)
	var MasaTotalTahun = 0
	var MasaTotalBulan = 0

	// fmt.Println("Year Bawaan Tahun : ", MasaBawaanTahun)
	// fmt.Println("Month Bawaan Tahun : ", MasaBawaanBulan)

	if MasaBawaanTahun != 0 || yearSkNow != 0 {
		MasaTotalTahun = MasaBawaanTahun + yearSkNow
	}

	if MasaBawaanBulan != 00 || monthSkNow != 0 {
		MasaTotalBulan = MasaBawaanBulan + monthSkNow
		if MasaTotalBulan >= 12 {
			MasaTotalBulan = MasaTotalBulan - 12
			MasaTotalTahun = MasaTotalTahun + 1
		}
	}
	// fmt.Println("Year Total Tahun : ", MasaTotalTahun)
	// fmt.Println("Month Total Tahun : ", MasaTotalBulan)

	a.MasaKerjaTotalBulan, _ = conv.String(MasaTotalBulan)
	a.MasaKerjaTotalTahun, _ = conv.String(MasaTotalTahun)
}
