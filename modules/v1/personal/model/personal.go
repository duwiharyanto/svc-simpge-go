package model

type PersonalDataPribadi struct {
	Id            uint64 `json:"-" gorm:"primaryKey"`
	IdKafka       int    `json:"-"`
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
	Id                 int    `json:"id" gorm:"primaryKey"`
	NamaLengkap        string `json:"nama_lengkap"`
	NikKtp             string `json:"nik_ktp"`
	NikPegawai         string `json:"nik_pegawai"`
	TglLahir           string `json:"tgl_lahir"`
	TempatLahir        string `json:"tempat_lahir"`
	IdAgama            int    `json:"id_agama"`
	KdAgama            string `json:"kd_agama"`
	JenisKelamin       string `json:"jenis_kelamin"`
	IdGolonganDarah    int    `json:"id_golongan_darah"`
	KdGolonganDarah    string `json:"kd_golongan_darah"`
	IdStatusPernikahan int    `json:"id_status_pernikahan"`
}

func (*PersonalDataPribadiId) TableName() string {
	return "personal_data_pribadi"
}
