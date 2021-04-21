package repo

import "fmt"

func getJenisIjazahQuery() string {
	return "SELECT kd_jenis_ijazah,jenis_ijazah, uuid FROM jenis_ijazah WHERE flag_aktif=1"
}
func getJenisIjazahByUUID(uuid string) string {
	return fmt.Sprintf(`SELECT id, kd_jenis_ijazah, jenis_ijazah, uuid FROM jenis_ijazah WHERE uuid = %q`, uuid)
}
