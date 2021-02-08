package model

type KelompokPegawai struct {
	KdKelompokPegawai string `json:"kd_kelompok_pegawai"`
	KdStatusPegawai   string `json:"-"`
	KdJenisPegawai    string `json:"-"`
	KelompokPegawai   string `json:"kelompok_pegawai"`
	UUID              string `json:"uuid"`
}

type KelompokPegawaiResponse struct {
	Data []KelompokPegawai `json:"data"`
}
