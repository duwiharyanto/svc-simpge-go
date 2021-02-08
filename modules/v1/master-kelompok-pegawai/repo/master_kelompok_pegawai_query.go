package repo

import "fmt"

func getKelompokPegawaiQuery(IdJenisPegawai string) string {
	var IdJenisPegawaiFilter string
	if IdJenisPegawai != "" {
		IdJenisPegawaiFilter = fmt.Sprintf("AND id_jenis_pegawai = %s", IdJenisPegawai)
	}
	return fmt.Sprintf("SELECT kd_status_pegawai, kd_jenis_pegawai, kelompok_pegawai, kd_kelompok_pegawai, uuid FROM kelompok_pegawai WHERE flag_aktif=1 %s ORDER BY kelompok_pegawai", IdJenisPegawaiFilter)
}
func getKelompokPegawaiByUUID(uuid string) string {
	return fmt.Sprintf(`SELECT id, kd_status_pegawai, kd_jenis_pegawai, kelompok_pegawai, kd_kelompok_pegawai, uuid FROM kelompok_pegawai WHERE flag_aktfif=1 AND uuid = %q`, uuid)
}
