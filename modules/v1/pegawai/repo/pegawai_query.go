package repo

import (
	"fmt"
	"svc-insani-go/helper"
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

func getPegawaiYayasanQuery(uuid string) string {
	q := fmt.Sprintf(`
	SELECT
		COALESCE(jp.kd_jenis_pegawai,''),
		COALESCE(jp.nama_jenis_pegawai,''),
		COALESCE(sp.kd_status_pegawai,''),
		COALESCE(sp.status_pegawai,''),
		COALESCE(kp.kd_kelompok_pegawai,''),
		COALESCE(kp.kelompok_pegawai,''),
		COALESCE(pgp.kd_pangkat_gol,''),
		COALESCE(pgp.pangkat,''),
		COALESCE(pgp.golongan,''),
		COALESCE(pf.tmt_pangkat_golongan,''),
		COALESCE(jf.kd_fungsional,''),
		COALESCE(jf.fungsional,''),
		COALESCE(pf.tmt_jabatan,''),
		COALESCE(pf.masa_kerja_bawaan_tahun,''),
		COALESCE(pf.masa_kerja_bawaan_bulan,''),
		COALESCE(pf.masa_kerja_gaji_tahun,''),
		COALESCE(pf.masa_kerja_gaji_bulan,''),
		COALESCE(pf.masa_kerja_total_tahun,''),
		COALESCE(pf.masa_kerja_total_bulan,''),
		COALESCE(pf.angka_kredit,''),
		COALESCE(pf.nomor_sertifikasi,''),
		COALESCE(jnr.kd_jenis_regis,''),
		COALESCE(jnr.jenis_no_regis,''),
		COALESCE(pf.nomor_registrasi,'')
	FROM
		pegawai p 
	LEFT JOIN
		pegawai_fungsional pf ON p.id = pf.id_pegawai 
	LEFT JOIN 
		jenis_pegawai jp ON p.id_jenis_pegawai = jp.id 
	LEFT JOIN 
		status_pegawai sp ON p.id_status_pegawai = sp.id 
	LEFT JOIN
		kelompok_pegawai kp ON p.id_kelompok_pegawai = kp.id 
	LEFT JOIN
		pangkat_golongan_pegawai pgp ON pf.id_pangkat_golongan = pgp.id 
	LEFT JOIN 
		jabatan_fungsional jf ON pf.id_jabatan_fungsional = jf.id 
	LEFT JOIN 
		jenis_nomor_registrasi jnr ON pf.id_jenis_nomor_registrasi = jnr.id
	WHERE
		p.uuid = %q`, uuid)

	return helper.FlatQuery(q)
}

func getUnitKerjaPegawaiQuery(uuid string) string {
	q := fmt.Sprintf(`
	SELECT 
		COALESCE(u1.kd_unit1,''),
		COALESCE(u1.unit1,''),
		COALESCE(u2.kd_unit2,''),
		COALESCE(u2.unit2,''),
		COALESCE(u3.kd_unit3,''),
		COALESCE(u3.unit3,''),
		COALESCE(lk.lokasi_desc,''),
		COALESCE(lk.lokasi_desc,''),
		COALESCE(pf.nomor_sk_pertama,''),
		COALESCE(pf.tmt_sk_pertama,'')
	FROM
		pegawai p
	LEFT JOIN
		pegawai_fungsional pf ON p.id = pf.id_pegawai 
	LEFT JOIN
		unit1 u1 ON p.id_unit_kerja1 = u1.id 
	LEFT JOIN
		unit2 u2 ON p.id_unit_kerja2 = u2.id 
	LEFT JOIN
		unit3 u3 ON p.id_unit_kerja3 = u3.id 
	LEFT JOIN
		lokasi_kerja lk ON p.lokasi_kerja = lk.lokasi_kerja 
	WHERE 
		p.uuid = %q`, uuid)

	return helper.FlatQuery(q)
}

func getPegawaiPNSQuery(uuid string) string {
	q := fmt.Sprintf(`
	SELECT 
		COALESCE(pp.nip_pns,''),
		COALESCE(pp.no_kartu_pegawai,''),
		COALESCE(pgp.kd_pangkat_gol,''),
		COALESCE(pgp.pangkat,''),
		COALESCE(pgp.golongan,''),
		COALESCE(pp.tmt_pangkat_golongan,''),
		COALESCE(jf.kd_fungsional,''),
		COALESCE(jf.fungsional,''),
		COALESCE(pp.tmt_jabatan,''),
		COALESCE(pp.masa_kerja_tahun,''),
		COALESCE(pp.masa_kerja_bulan,''),
		COALESCE(pp.angka_kredit,''),
		COALESCE(pp.keterangan,'')
	FROM 
		pegawai p
	LEFT JOIN
		pegawai_pns pp ON p.id = pp.id_pegawai 
	LEFT JOIN
		pangkat_golongan_pegawai pgp ON pp.id_pangkat_golongan 
	LEFT JOIN
		jabatan_fungsional jf ON pp.id_jabatan_fungsional = jf.id
	WHERE 
		p.uuid = %q`, uuid)

	return helper.FlatQuery(q)
}

func getPegawaiPTTQuery(uuid string) string {
	q := fmt.Sprintf(`
	SELECT
		COALESCE(jptt.kd_jenis_ptt,''),
		COALESCE(jptt.jenis_ptt,''),
		COALESCE(ptt.instansi_asal,''),
		COALESCE(ptt.keterangan,'')
	FROM
		pegawai p
	LEFT JOIN
		pegawai_tidak_tetap ptt ON p.id = ptt.id_pegawai
	LEFT JOIN
		jenis_pegawai_tidak_tetap jptt ON ptt.id_jenis_ptt = jptt.id
	WHERE
		p.uuid = %q`, uuid)

	return helper.FlatQuery(q)
}

func getStatusPegawaiAktifQuery(uuid string) string {
	q := fmt.Sprintf(`
	SELECT 
		COALESCE(spa.flag_status_aktif,''),
		COALESCE(spa.kd_status,''),
		COALESCE(spa.status,'')
	FROM 
		pegawai p
	LEFT JOIN
		pegawai_fungsional pf ON p.id = pf.id_pegawai
	LEFT JOIN
		status_pegawai_aktif spa ON pf.id_status_pegawai_aktif = spa.id 
	WHERE
		p.uuid = %q`, uuid)

	return helper.FlatQuery(q)
}

func getPegawaiPribadiQuery(uuid string) string {
	q := fmt.Sprintf(`
	SELECT 
		p.nama,
		p.nik,
		jp.nama_jenis_pegawai,
		kp.kelompok_pegawai,
		u2.unit2,
		p.uuid 
	FROM 
		pegawai p
	LEFT JOIN 
		jenis_pegawai jp ON p.id_jenis_pegawai = jp.id 
	LEFT JOIN 
		kelompok_pegawai kp ON p.id_kelompok_pegawai = kp.id 
	LEFT JOIN
		unit2 u2 ON p.id_unit_kerja2 = u2.id 
	WHERE
		p.uuid = %q`, uuid)

	return helper.FlatQuery(q)
}
