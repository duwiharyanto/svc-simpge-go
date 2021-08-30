package model

import (
	pegawai "svc-insani-go/modules/v1/pegawai/model"
)

type PersonalDataPribadi struct {
	Id            uint64 `json:"-" gorm:"primaryKey"`
	IdKafka       uint64 `json:"-"`
	NamaLengkap   string `json:"nama_lengkap"`
	GelarDepan    string `json:"gelar_depan"`
	GelarBelakang string `json:"gelar_belakang"`
	NikKtp        string `json:"nik_ktp"`
	NikPegawai    string `json:"nik_pegawai"`
	Uuid          string `json:"uuid"`
}

type PersonalDataPribadiResponse struct {
	Data []PersonalDataPribadi `json:"data"`
}

type PersonalDataPribadiId struct {
	Id                 uint64 `json:"id" gorm:"primaryKey"`
	NamaLengkap        string `json:"nama_lengkap"`
	GelarDepan         string `json:"gelar_depan"`
	GelarBelakang      string `json:"gelar_belakang"`
	NikKtp             string `json:"nik_ktp"`
	NikPegawai         string `json:"nik_pegawai"`
	TglLahir           string `json:"tgl_lahir"`
	TempatLahir        string `json:"tempat_lahir"`
	IdAgama            uint64 `json:"id_agama"`
	KdAgama            string `json:"kd_agama"`
	JenisKelamin       string `json:"jenis_kelamin"`
	IdGolonganDarah    uint64 `json:"id_golongan_darah"`
	KdGolonganDarah    string `json:"kd_golongan_darah"`
	IdStatusPernikahan uint64 `json:"id_status_pernikahan"`
	KdStatusPerkawinan string `json:"kd_status_perkawinan"`

	Pegawai          pegawai.Pegawai  `json:"-" gorm:"foreignKey:IdPersonalDataPribadi"`
	Agama            Agama            `json:"-" gorm:"foreignKey:IdAgama"`
	GolonganDarah    GolonganDarah    `json:"-" gorm:"foreignKey:IdGolonganDarah"`
	StatusPernikahan StatusPernikahan `json:"-" gorm:"foreignKey:IdStatusPernikahan"`
}

func (*PersonalDataPribadiId) TableName() string {
	return "personal_data_pribadi"
}

type Agama struct {
	Id     uint64 `gorm:"primaryKey"`
	KdItem string
	Agama  string
	Uuid   string
}

const UnknownBloodType = "belum diketahui"

type GolonganDarah struct {
	Id            uint64 `gorm:"primaryKey"`
	GolonganDarah string
	Uuid          string
}

type StatusPernikahan struct {
	Id       uint64 `gorm:"primaryKey"`
	KdStatus string
	Status   string
	Uuid     string
}
