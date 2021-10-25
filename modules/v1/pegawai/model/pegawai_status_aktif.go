package model

type StatusAktif struct {
	FlagAktifPegawai         string `json:"flag_aktif_pegawai"`
	StatusAktifPegawai       string `json:"status_aktif_pegawai"`
	KdStatusAktifPegawai     string `json:"kd_status_aktif_pegawai"`
	UuidStatusAktifPegawai   string `json:"uuid_status_aktif_pegawai"`
	TglStatusAktifPegawai    string `json:"tgl_status_aktif"`
	TglStatusAktifPegawaiIdn string `json:"tgl_status_aktif_idn"`
}
