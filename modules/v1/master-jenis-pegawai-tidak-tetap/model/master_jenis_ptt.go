package model

type JenisPTT struct {
	ID         string `json:"-"`
	KdJenisPTT string `json:"kd_jenis_ptt"`
	JenisPTT   string `json:"jenis_ptt"`
	UserInput  string `json:"-"`
	UserUpdate string `json:"-"`
	UUID       string `json:"uuid"`
}

type JenisPTTResponse struct {
	Data []JenisPTT `json:"data"`
}
