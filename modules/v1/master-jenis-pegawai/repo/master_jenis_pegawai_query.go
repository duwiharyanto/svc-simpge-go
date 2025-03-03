package repo

import "fmt"

func getJenisPegawaiQuery() string {
	return "SELECT kd_jenis_pegawai,nama_jenis_pegawai, uuid FROM jenis_pegawai WHERE flag_aktif=1 ORDER BY nama_jenis_pegawai"
}
func getJenisPegawaiQueryByUUID(uuid string) string {
	return fmt.Sprintf(`SELECT id, kd_jenis_pegawai, nama_jenis_pegawai, uuid FROM jenis_pegawai WHERE flag_aktif=1 AND uuid = %q`, uuid)
}
func getJenisPegawaiQueryByKdJenisPegawai(KDJenisPegawai string) string {
	return fmt.Sprintf(`SELECT id, kd_jenis_pegawai, nama_jenis_pegawai, uuid FROM jenis_pegawai WHERE flag_aktif=1 AND kd_jenis_pegawai = %q ORDER BY nama_jenis_pegawai`, KDJenisPegawai)
}
