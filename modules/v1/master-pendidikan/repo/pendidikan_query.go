package repo

import (
	"fmt"
	"svc-insani-go/app/helper"
)

func getGelarDepanQuery() string {
	q := fmt.Sprintf(`SELECT gelar FROM gelar WHERE flag_aktif = 1 AND kd_posisi_gelar = '1'`)
	return helper.FlatQuery(q)
}

func getGelarBelakangQuery() string {
	q := fmt.Sprintf(`SELECT gelar FROM gelar WHERE flag_aktif = 1 AND kd_posisi_gelar = '2'`)
	return helper.FlatQuery(q)
}

func getJenjangPendidikanQuery() string {
	q := fmt.Sprintf(`SELECT
	id, COALESCE(kd_jenjang,''), COALESCE(jenjang,''), COALESCE(nama_jenjang,''),
	COALESCE(kd_pendidikan_simpeg,''), COALESCE(nama_pendidikan_simpeg,''),	uuid
	FROM jenjang_pendidikan WHERE flag_aktif = 1 AND kd_pendidikan_simpeg IS NOT NULL`)
	return helper.FlatQuery(q)
}

func getJenjangPendidikanDetailQuery(kdJenjangPendidikan string) string {
	var kdJenjangPendidikanFilter string
	if kdJenjangPendidikan != "" {
		kdJenjangPendidikanFilter = fmt.Sprintf(`AND kd_jenjang_pendidikan = %q`, kdJenjangPendidikan)
	}
	q := fmt.Sprintf(`SELECT id, kd_jenjang_pendidikan, kd_detail, nama_detail, keterangan, uuid
	FROM jenjang_pendidikan_detail WHERE flag_aktif = 1 %s`,
		kdJenjangPendidikanFilter,
	)
	return helper.FlatQuery(q)
}
