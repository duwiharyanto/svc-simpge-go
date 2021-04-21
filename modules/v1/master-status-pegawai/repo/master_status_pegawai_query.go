package repo

func getAllStatusPegawaiQuery() string {
	return "SELECT kd_status_pegawai, status_pegawai, uuid FROM status_pegawai WHERE flag_aktif = 1"
}
