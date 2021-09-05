package model

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
