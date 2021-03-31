package repo

import "fmt"

// func getJenisSKPengangkatanQuery() string {
// 	return "SELECT kd_jenis_sk_pengangkatan,jenis_sk_pengangkatan, kd_kelompok_pegawai, uuid FROM jenis_sk_pengangkatan WHERE flag_aktif=1"
// }
func getJenisSKPengangkatanQuery(kdJenisPegawai string) string {
	var kdJenisPegawaiFilter string
	if kdJenisPegawai == "ED" {
		kdJenisPegawaiFilter = fmt.Sprintf("AND kd_jenis_pegawai = '%s'", kdJenisPegawai)
	} else if kdJenisPegawai == "TPA" || kdJenisPegawai == "AD" {
		kdJenisPegawaiFilter = fmt.Sprintf("AND kd_jenis_pegawai != 'ED'")
	}
	return fmt.Sprintf("SELECT kd_kelompok_pegawai,kelompok_pegawai, kd_kelompok_pegawai,uuid FROM kelompok_pegawai WHERE flag_aktif=1 %s", kdJenisPegawaiFilter)
}
func getJenisSKPengangkatanQueryByUUID(uuid string) string {
	return fmt.Sprintf(`SELECT kd_jenis_sk_pengangkatan, jenis_sk_pengangkatan, kd_kelompok_pegawai, uuid FROM jenis_sk_pengangkatan WHERE uuid = %q`, uuid)
}
