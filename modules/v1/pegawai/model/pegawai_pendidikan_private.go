package model

type PegawaiPendidikanPrivate struct {
	IdPersonal       uint64 `form:"id_personal_data_pribadi" json:"-" gorm:"primaryKey;column:id"`
	KdJenjang        string `json:"kd_jenjang_pendidikan"`
	FlagIjazahDiakui string `form:"flag_ijazah_tertinggi_diakui" json:"flag_ijazah_tertinggi_diakui"`
	NamaInstitusi    string `json:"nama_institusi"`
}
