package model

type GelarDepan struct {
	Gelar string `json:"gelar"`
}

type GelarBelakang struct {
	Gelar string `json:"gelar"`
}

type JenisPendidikan struct {
	ID                  uint64  `json:"-" gorm:"primaryKey"`
	KdJenjangPendidikan string  `json:"kd_jenjang_pendidikan"`
	KdJenis             string  `json:"kd_jenis"`
	NamaJenis           string  `json:"nama_jenis"`
	Keterangan          *string `json:"keterangan"`
	UUID                string  `json:"uuid"`
}

type JenjangPendidikan struct {
	ID                   uint64 `json:"-" gorm:"primaryKey"`
	KdJenjang            string `json:"kd_jenjang"`
	KdPendidikanSimpeg   string `json:"-"`
	Jenjang              string `json:"jenjang"`
	NamaJenjang          string `json:"nama_jenjang"`
	NamaPendidikanSimpeg string `json:"-"`
	UserInput            string `json:"-"`
	UserUpdate           string `json:"-"`
	UUID                 string `json:"uuid"`
}

type JenjangPendidikanDetail struct {
	ID                  uint64  `json:"-" gorm:"primaryKey"`
	KdJenjangPendidikan string  `json:"kd_jenjang_pendidikan"`
	KdDetail            string  `json:"kd_detail"`
	NamaDetail          string  `json:"nama_detail"`
	Keterangan          *string `json:"keterangan"`
	UUID                string  `json:"uuid"`
}

type JenjangPendidikanList []JenjangPendidikan

func (jj JenjangPendidikanList) MapByKdPendidikanSimpeg() map[string]JenjangPendidikan {
	m := make(map[string]JenjangPendidikan)
	for _, j := range jj {
		m[j.KdPendidikanSimpeg] = j
	}
	return m
}

type JenjangPendidikanResponse struct {
	Data []JenjangPendidikan `json:"data"`
}
