package repo

import "fmt"

func getPegawaiPenetapSKQuery() string {
	return "SELECT nik, nama, COALESCE(gelar_depan,''), COALESCE(gelar_belakang,''), COALESCE(kd_kelompok_pegawai,''), COALESCE(kd_unit,''), uuid FROM pegawai_penetap_sk WHERE flag_aktif=1"
}
func getPegawaiPenetapSKQueryByUUID(uuid string) string {
	return fmt.Sprintf(`SELECT id, nik, nama, COALESCE(gelar_depan,''), COALESCE(gelar_belakang,''), COALESCE(kd_kelompok_pegawai,''), COALESCE(kd_unit,''), uuid FROM pegawai_penetap_sk WHERE flag_aktif=1 AND uuid = %q`, uuid)
}
