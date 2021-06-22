package repo

import (
	"fmt"
)

func getJenjangPendidikanQuery() string {
	return fmt.Sprintf(`SELECT id, COALESCE(kd_jenjang,''), COALESCE(jenjang,''), COALESCE(nama_jenjang,''), uuid FROM jenjang_pendidikan WHERE flag_aktif=1`)
}
