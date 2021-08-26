package repo

import (
	"fmt"
	"svc-insani-go/app/helper"
)

func getJenjangPendidikanQuery() string {
	q := fmt.Sprintf(`SELECT
	id, COALESCE(kd_jenjang,''), COALESCE(jenjang,''), COALESCE(nama_jenjang,''),
	COALESCE(kd_pendidikan_simpeg,''), COALESCE(nama_pendidikan_simpeg,''),	uuid
	FROM jenjang_pendidikan WHERE flag_aktif = 1 AND kd_pendidikan_simpeg IS NOT NULL`)
	return helper.FlatQuery(q)
}
