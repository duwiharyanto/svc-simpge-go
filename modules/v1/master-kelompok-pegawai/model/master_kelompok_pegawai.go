package model

type KelompokPegawai struct {
	KdKelompokPegawai string `json:"kd_kelompok_pegawai"`
	KdStatusPegawai   string `json:"kd_status_pegawai"`
	KdJenisPegawai    string `json:"kd_jenis_pegawai"`
	KelompokPegawai   string `json:"kelompok_pegawai"`
	UUID              string `json:"uuid"`
}

type KelompokPegawaiResponse struct {
	Data []KelompokPegawai `json:"data"`
}
