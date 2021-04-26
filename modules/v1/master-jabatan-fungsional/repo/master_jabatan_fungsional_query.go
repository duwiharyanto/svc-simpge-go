package repo

import "fmt"

func getJabatanFungsionalQuery(kdJenisPegawai string) string {
	var kdJenisPegawaiFilter string
	if kdJenisPegawai == "ED" {
		kdJenisPegawaiFilter = fmt.Sprintf("AND kd_jenis_pegawai = '%s'", kdJenisPegawai)
	} else if kdJenisPegawai == "TPA" || kdJenisPegawai == "AD" {
		kdJenisPegawaiFilter = fmt.Sprintf("AND kd_jenis_pegawai != 'ED'")
	}
	return fmt.Sprintf("SELECT kd_fungsional,fungsional, uuid FROM jabatan_fungsional WHERE flag_aktif=1 %s", kdJenisPegawaiFilter)
}
func getJabatanFungsionalByUUID(uuid string) string {
	return fmt.Sprintf(`SELECT id, kd_fungsional, fungsional, uuid FROM jabatan_fungsional WHERE flag_aktif=1 AND uuid = %q`, uuid)
}
