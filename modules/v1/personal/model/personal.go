package model

type PersonalDataPribadi struct {
	Id            int    `json:"id" gorm:"primaryKey"`
	IdKafka       int    `json:"id_kafka"`
	NamaLengkap   string `json:"nama_lengkap"`
	GelarDepan    string `json:"gelar_depan"`
	GelarBelakang string `json:"gelar_belakang"`
	NikKtp        string `json:"nik_ktp"`
	NikPegawai    string `json:"nik_pegawai"`
	Uuid          string `json:"uuid"`
}
