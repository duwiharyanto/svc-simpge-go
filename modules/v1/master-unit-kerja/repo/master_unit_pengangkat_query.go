package repo

import "fmt"

func getUnitPengangkatQuery() string {
	return "SELECT kd_unit_kerja,nama_unit_kerja, uuid FROM unit_kerja WHERE flag_aktif=1"
}
func getUnitPengangkatQueryByUUID(uuid string) string {
	return fmt.Sprintf(`SELECT id, kd_unit_kerja, nama_unit_kerja, uuid FROM unit_kerja WHERE flag_aktif=1 AND uuid = %q`, uuid)
}
