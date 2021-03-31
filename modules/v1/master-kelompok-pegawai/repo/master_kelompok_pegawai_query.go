package repo

import "fmt"

func getKelompokPegawaiQuery(kdJenisPegawai string) string {
	var kdJenisPegawaiFilter string
	if kdJenisPegawai == "ED" {
		kdJenisPegawaiFilter = fmt.Sprintf("AND kd_jenis_pegawai = '%s'", kdJenisPegawai)
	} else if kdJenisPegawai == "TPA" || kdJenisPegawai == "AD" {
		kdJenisPegawaiFilter = fmt.Sprintf("AND kd_jenis_pegawai != 'ED'")
	}
	return fmt.Sprintf("SELECT kd_status_pegawai,kd_jenis_pegawai, kd_kelompok_pegawai, kelompok_pegawai,uuid FROM kelompok_pegawai WHERE flag_aktif=1 %s ORDER BY kelompok_pegawai", kdJenisPegawaiFilter)
}

func getAllKelompokPegawaiByUUID(uuid string) string {
	return fmt.Sprintf(`SELECT id, kd_status_pegawai, kd_jenis_pegawai, kelompok_pegawai, kd_kelompok_pegawai, uuid FROM kelompok_pegawai WHERE flag_aktfif=1 AND uuid = %q`, uuid)
}

func getKelompokPegawaiByUUID(uuid string) string {
	return fmt.Sprintf(`SELECT id, kd_status_pegawai, kd_jenis_pegawai, kelompok_pegawai, uuid FROM kelompok_pegawai WHERE flag_aktif=1 AND uuid = %q`, uuid)
}
