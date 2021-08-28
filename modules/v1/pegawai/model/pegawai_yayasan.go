package model

import (
	"fmt"
	"strconv"
	"time"
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
	masaKerjaBawaanTahunInt, _ := strconv.Atoi(a.MasaKerjaBawaanTahun)
	masaKerjaBawaanBulanInt, _ := strconv.Atoi(a.MasaKerjaBawaanBulan)

	tmtSkPertamaTime, err := time.Parse("2006-01-02", b.TmtSkPertama)
	var tmtSkPertamaDuration time.Duration
	if err == nil {
		tmtSkPertamaDuration = time.Now().Sub(tmtSkPertamaTime)
	}
	tmtSkPertamaDurationDays := tmtSkPertamaDuration.Hours() / 24
	// tmtSkPertamaDurationYears := int(tmtSkPertamaDurationDays / 365)
	tmtSkPertamaDurationRealMonths := int(tmtSkPertamaDurationDays / 365 * 12)
	// tmtSkPertamaDurationMonths := int(tmtSkPertamaDurationDays / 30 % 12)

	masaKerjaTotalRealBulan := ((masaKerjaBawaanTahunInt * 12) + masaKerjaBawaanBulanInt) + tmtSkPertamaDurationRealMonths
	a.MasaKerjaTotalTahun = fmt.Sprintf("%d", masaKerjaTotalRealBulan/12)
	a.MasaKerjaTotalBulan = fmt.Sprintf("%d", masaKerjaTotalRealBulan%12)

}
