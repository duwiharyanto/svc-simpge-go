package repo

import "fmt"

func getPangkatPegawaiQuery() string {
	return "SELECT kd_pangkat,nama_pangkat, uuid FROM pangkat WHERE flag_aktif=1"
}
func getPangkatPegawaiByUUIDQuery(uuid string) string {
	return fmt.Sprintf(`SELECT id, kd_pangkat, nama_pangkat, uuid FROM pangkat WHERE flag_aktif=1 AND uuid = %q`, uuid)
}
func getPangkatGolonganPegawaiQuery() string {
	return "SELECT pangkat, golongan, uuid FROM pangkat_golongan_pegawai WHERE flag_aktif=1"
}
func getPangkatGolonganPegawaiQueryByUUID(uuid string) string {
	return fmt.Sprintf(`SELECT id, COALESCE(kd_pangkat_gol, ''), pangkat, golongan, uuid FROM pangkat_golongan_pegawai WHERE flag_aktif=1 AND uuid = %q`, uuid)
}
