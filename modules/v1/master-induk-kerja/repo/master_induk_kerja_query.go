package repo

import (
	"fmt"
)

func getIndukKerjaQuery() string {
	return fmt.Sprintf(`SELECT id, kd_unit1, unit1, keterangan2, uuid FROM unit1 WHERE flag_aktif=1`)
}

func getUnitKerjaQuery(indukKerja string) string {
	return fmt.Sprintf(`SELECT id, kd_unit2, unit2, keterangan2, uuid FROM unit2 WHERE flag_aktif=1 AND keterangan1 = %q`, indukKerja)
}

func getBagianKerjaQuery(unitKerja string) string {
	return fmt.Sprintf(`SELECT id, kd_unit3, unit3, COALESCE(keterangan2,''), uuid FROM unit3 WHERE flag_aktif="Y" AND keterangan1 = %q`, unitKerja)
}
