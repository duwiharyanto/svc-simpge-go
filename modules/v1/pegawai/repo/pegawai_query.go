package repo

import (
	"fmt"
	"svc-insani-go/modules/v1/pegawai/model"
)

func getListAllPegawaiQuery(req *model.PegawaiRequest) string {
	var kdKelompokFilterQuery string
	if req.KdKelompokPegawai != "" {
		kdKelompokFilterQuery = fmt.Sprintf("AND a.kd_kelompok_pegawai = '%s'", req.KdKelompokPegawai)
	}
	var kdUnitKerjaFilterQuery string
	if req.KdUnitKerja != "" {
		kdUnitKerjaFilterQuery = fmt.Sprintf("AND a.kd_unit2 = '%s'", req.KdUnitKerja)
	}
	var namaFilterQuery string
	if req.Cari != "" {
		namaFilterQuery = fmt.Sprintf("AND a.nama LIKE '%%%s%%'", req.Cari)
	}
	var paginationFilterQuery string
	if req.Limit != 0 {
		paginationFilterQuery = fmt.Sprintf("LIMIT %d,%d", req.Offset, req.Limit)
	}
	return fmt.Sprintf(`SELECT a.nik, a.nama, COALESCE(a.gelar_depan,''), COALESCE(a.gelar_belakang,''),
	COALESCE(a.kd_kelompok_pegawai,''), COALESCE(c.kelompok_pegawai,''), COALESCE(c.uuid,''),
	COALESCE(a.kd_unit2,''), COALESCE(b.unit2,''), COALESCE(b.uuid,''),
	COALESCE(d.kd_jenis_pegawai,''), COALESCE(d.nama_jenis_pegawai,''), COALESCE(d.uuid,''),
	COALESCE(e.kd_status_pegawai,''), COALESCE(e.status_pegawai,''), COALESCE(e.uuid,''),
	a.uuid
	FROM pegawai a
	LEFT JOIN unit2 b ON a.kd_unit2 = b.kd_unit2
	LEFT JOIN kelompok_pegawai c ON a.kd_kelompok_pegawai = c.kd_kelompok_pegawai
	LEFT JOIN jenis_pegawai d ON a.id_jenis_pegawai = d.id
	LEFT JOIN status_pegawai e ON a.id_status_pegawai = e.id
	WHERE a.flag_aktif=1 %s %s %s ORDER BY a.nama %s`,
		kdKelompokFilterQuery,
		kdUnitKerjaFilterQuery,
		namaFilterQuery,
		paginationFilterQuery,
	)
}

func countPegawaiQuery(req *model.PegawaiRequest) string {
	var kdKelompokFilterQuery string
	if req.KdKelompokPegawai != "" {
		kdKelompokFilterQuery = fmt.Sprintf("AND a.kd_kelompok_pegawai = '%s'", req.KdKelompokPegawai)
	}
	var kdUnitKerjaFilterQuery string
	if req.KdUnitKerja != "" {
		kdUnitKerjaFilterQuery = fmt.Sprintf("AND a.kd_unit2 = '%s'", req.KdUnitKerja)
	}
	var namaFilterQuery string
	if req.Cari != "" {
		namaFilterQuery = fmt.Sprintf("AND a.nama LIKE '%%%s%%'", req.Cari)
	}
	return fmt.Sprintf(`SELECT COUNT(*)
	FROM pegawai a
	WHERE a.flag_aktif = 1 %s %s %s`,
		kdKelompokFilterQuery,
		kdUnitKerjaFilterQuery,
		namaFilterQuery,
	)
}

func getPegawaiByUUID(uuid string) string {
	return fmt.Sprintf(`SELECT id, nik, nama, COALESCE(gelar_depan,''), COALESCE(gelar_belakang,''), COALESCE(kd_kelompok_pegawai,''), COALESCE(kd_unit2,''), uuid FROM pegawai WHERE flag_aktif=1 AND uuid = %q`, uuid)
}
