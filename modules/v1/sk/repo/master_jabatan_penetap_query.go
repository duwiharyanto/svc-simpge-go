package repo

import "fmt"

func getJabatanPenetapQuery() string {
	return "SELECT id, kd_jabatan_penetap, jabatan_penetap, uuid FROM hcm_insani.jabatan_penetap WHERE flag_aktif=1"
}
func getJabatanPenetapQueryByUUID(uuid string) string {
	return fmt.Sprintf(`SELECT id, kd_jabatan_penetap, jabatan_penetap, uuid FROM jabatan_penetap WHERE flag_aktif=1 AND uuid = %q`, uuid)
}
