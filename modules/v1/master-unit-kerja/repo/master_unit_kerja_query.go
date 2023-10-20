package repo

import "fmt"

func getUnitKerjaQuery() string {
	return "SELECT kd_unit_kerja,nama_unit_kerja, uuid FROM unit_kerja WHERE flag_aktif ORDER BY nama_unit_kerja"
}

func getUnitKerjaByUUID(uuid string) string {
	return fmt.Sprintf(`SELECT id, kd_unit_kerja, nama_unit_kerja, uuid FROM unit_kerja WHERE flag_aktif=1 AND uuid = %q`, uuid)
}

func getUnit2ByUUID(uuid string) string {
	return fmt.Sprintf(`SELECT id, kd_unit2, unit2, uuid FROM unit2 WHERE uuid = %q`, uuid)
}
