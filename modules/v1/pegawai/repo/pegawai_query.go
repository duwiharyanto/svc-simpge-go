package repo

import (
	"fmt"
	"strings"
	"svc-insani-go/helper"
	"svc-insani-go/modules/v1/pegawai/model"
)

func getListAllPegawaiQuery(req *model.PegawaiRequest) string {
	var kdJenisPegawaiFilterQuery string
	if req.UuidJenisPegawai != "" {
		kdJenisPegawaiFilterQuery = fmt.Sprintf("AND d.uuid = '%s'", req.UuidJenisPegawai)
	}
	var kdKelompokFilterQuery string
	if req.UuidKelompokPegawai != "" {
		kdKelompokFilterQuery = fmt.Sprintf("AND c.uuid = '%s'", req.UuidKelompokPegawai)
	}
	var kdUnitKerjaFilterQuery string
	if req.UuidUnitKerja != "" {
		kdUnitKerjaFilterQuery = fmt.Sprintf("AND b.uuid = '%s'", req.UuidUnitKerja)
	}
	var kdStatusAktifFilterQuery string
	if req.UuidStatusAktif != "" {
		kdStatusAktifFilterQuery = fmt.Sprintf("AND e.uuid = '%s'", req.UuidStatusAktif)
	}
	var namaFilterQuery string
	if req.Cari != "" {
		namaFilterQuery = fmt.Sprintf("AND (a.nama LIKE '%%%s%%' OR a.nik LIKE '%%%s%%')", req.Cari, req.Cari)
	}
	var paginationFilterQuery string
	if req.Limit != 0 {
		paginationFilterQuery = fmt.Sprintf("LIMIT %d,%d", req.Offset, req.Limit)
	}
	return fmt.Sprintf(`SELECT a.nik, a.nama, COALESCE(a.gelar_depan,''), COALESCE(a.gelar_belakang,''),
	COALESCE(a.kd_kelompok_pegawai,''), COALESCE(c.kelompok_pegawai,''), COALESCE(c.uuid,''),
	COALESCE(a.kd_unit2,''), COALESCE(b.unit2,''), COALESCE(b.uuid,''),
	COALESCE(d.kd_jenis_pegawai,''), COALESCE(d.nama_jenis_pegawai,''), COALESCE(d.uuid,''),
	COALESCE(e.kd_status,''), COALESCE(e.status,''), COALESCE(e.uuid,''),
	a.uuid
	FROM pegawai a
	LEFT JOIN pegawai_fungsional f ON a.id = f.id_pegawai
	LEFT JOIN unit2 b ON a.id_unit_kerja2 = b.id
	LEFT JOIN kelompok_pegawai c ON a.id_kelompok_pegawai = c.id
	LEFT JOIN jenis_pegawai d ON a.id_jenis_pegawai = d.id
	LEFT JOIN status_pegawai_aktif e ON f.id_status_pegawai_aktif = e.id
	WHERE a.flag_aktif=1 %s %s %s %s %s ORDER BY a.nama %s`,
		kdJenisPegawaiFilterQuery,
		kdKelompokFilterQuery,
		kdUnitKerjaFilterQuery,
		kdStatusAktifFilterQuery,
		namaFilterQuery,
		paginationFilterQuery,
	)
}

func countPegawaiQuery(req *model.PegawaiRequest) string {
	var uuidJenisPegawaiFilterQuery string
	if req.UuidJenisPegawai != "" {
		uuidJenisPegawaiFilterQuery = fmt.Sprintf("AND d.uuid = %q", req.UuidJenisPegawai)
	}
	var uuidKelompokFilterQuery string
	if req.UuidKelompokPegawai != "" {
		uuidKelompokFilterQuery = fmt.Sprintf("AND c.uuid = %q", req.UuidKelompokPegawai)
	}
	var uuidUnitKerjaFilterQuery string
	if req.UuidUnitKerja != "" {
		uuidUnitKerjaFilterQuery = fmt.Sprintf("AND b.uuid = %q", req.UuidUnitKerja)
	}
	var kdStatusAktifFilterQuery string
	if req.UuidStatusAktif != "" {
		kdStatusAktifFilterQuery = fmt.Sprintf("AND e.uuid = %q", req.UuidStatusAktif)
	}
	var namaFilterQuery string
	if req.Cari != "" {
		namaFilterQuery = fmt.Sprintf(`AND (a.nama LIKE "%%%s%%" OR a.nik LIKE "%%%s%%")`, req.Cari, req.Cari)
	}
	return fmt.Sprintf(`SELECT COUNT(*)
	FROM pegawai a
	LEFT JOIN pegawai_fungsional f ON a.id = f.id_pegawai
	LEFT JOIN unit2 b ON a.id_unit_kerja2 = b.id
	LEFT JOIN kelompok_pegawai c ON a.id_kelompok_pegawai = c.id
	LEFT JOIN jenis_pegawai d ON a.id_jenis_pegawai = d.id
	LEFT JOIN status_pegawai_aktif e ON f.id_status_pegawai_aktif = e.id
	WHERE a.flag_aktif = 1 %s %s %s %s %s`,
		uuidJenisPegawaiFilterQuery,
		uuidKelompokFilterQuery,
		uuidUnitKerjaFilterQuery,
		kdStatusAktifFilterQuery,
		namaFilterQuery,
	)
}

func getPegawaiByUUID(uuid string) string {
	q := fmt.Sprintf(`SELECT
	a.id,
	a.nik,
	a.nama,
	COALESCE(a.gelar_depan,''),
	COALESCE(a.gelar_belakang,''),
	COALESCE(b.kd_jenis_pegawai,''),
	COALESCE(b.nama_jenis_pegawai,''),
	COALESCE(b.uuid,''),
	COALESCE(c.kd_kelompok_pegawai,''),
	COALESCE(c.kelompok_pegawai,''),
	COALESCE(c.uuid,''),
	COALESCE(d.kd_unit2,''),
	COALESCE(d.unit2,''),
	COALESCE(d.uuid,''),
	a.uuid FROM
	pegawai a
	LEFT JOIN jenis_pegawai b ON a.id_jenis_pegawai = b.id
	LEFT JOIN kelompok_pegawai c ON a.id_kelompok_pegawai = c.id
	LEFT JOIN unit2 d ON a.id_unit_kerja2 = d.id
	WHERE a.flag_aktif=1 AND a.uuid = %q`, uuid)
	return helper.FlatQuery(q)
}

func getPegawaiYayasanQuery(uuid string) string {
	q := fmt.Sprintf(`
	SELECT
		p.id,
		COALESCE(jp.uuid,''),
		COALESCE(jp.id, 0),
		COALESCE(jp.kd_jenis_pegawai,''),
		COALESCE(jp.nama_jenis_pegawai,''),
		COALESCE(sp.uuid,''),
		COALESCE(sp.id, 0),
		COALESCE(sp.kd_status_pegawai,''),
		COALESCE(sp.status_pegawai,''),
		COALESCE(kp.uuid,''),
		COALESCE(kp.id, 0),
		COALESCE(kp.kd_kelompok_pegawai,''),
		COALESCE(kp.kelompok_pegawai,''),
		COALESCE(jp2.id, 0),
		COALESCE(jp2.uuid,''),
		COALESCE(jp2.kd_jenjang,''),
		COALESCE(jp2.kd_pendidikan_simpeg,''),
		COALESCE(jp2.nama_pendidikan_simpeg,''),
		COALESCE(jp3.uuid,''),
		COALESCE(jp3.id, 0),
		COALESCE(jp3.kd_jenjang,''),
		COALESCE(jp3.kd_pendidikan_simpeg,''),
		COALESCE(jp3.nama_pendidikan_simpeg,''),
		COALESCE(pgp.uuid,''),
		COALESCE(pgp.kd_pangkat_gol,''),
		COALESCE(pgp.pangkat,''),
		COALESCE(pgp.kd_golongan,''),
		COALESCE(pgp.golongan,''),
		COALESCE(pgp.kd_ruang,''),
		COALESCE(pf.tmt_pangkat_golongan,''),
		COALESCE(jf.uuid,''),
		COALESCE(jf.kd_fungsional,''),
		COALESCE(jf.fungsional,''),
		COALESCE(pf.tmt_jabatan,''),
		COALESCE(pf.masa_kerja_bawaan_tahun,''),
		COALESCE(pf.masa_kerja_bawaan_bulan,''),
		COALESCE(pf.masa_kerja_gaji_tahun,''),
		COALESCE(pf.masa_kerja_gaji_bulan,''),
		COALESCE(pf.angka_kredit,''),
		COALESCE(pf.nomor_sertifikasi,''),
		COALESCE(jnr.uuid,''),
		COALESCE(jnr.id, 0),
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
	LEFT JOIN 
		jenjang_pendidikan jp2 ON p.id_pendidikan_masuk = jp2.id
	LEFT JOIN 
		jenjang_pendidikan jp3 ON p.id_pendidikan_terakhir = jp3.id
	WHERE
		p.uuid = %q AND p.flag_aktif = 1`, uuid)

	return helper.FlatQuery(q)
}

func getUnitKerjaPegawaiQuery(uuid string) string {
	q := fmt.Sprintf(`
	SELECT 
		COALESCE(u1.uuid,''),
		COALESCE(u1.kd_unit1,''),
		COALESCE(u1.unit1,''),
		COALESCE(u2.uuid,''),
		COALESCE(u2.kd_unit2,''),
		COALESCE(u2.unit2,''),
		COALESCE(u3.uuid,''),
		COALESCE(u3.kd_unit3,''),
		COALESCE(u3.unit3,''),
		COALESCE(lk.uuid,''),
		COALESCE(lk.lokasi_kerja,''),
		COALESCE(lk.lokasi_desc,''),
		COALESCE(pf.nomor_sk_pertama,''),
		COALESCE(pf.tmt_sk_pertama,''),
		COALESCE(u22.uuid,''),
		COALESCE(u22.kd_pddikti,''),
		COALESCE(u22.uuid,''),
		COALESCE(u22.kd_unit2,'')
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
	LEFT JOIN
		unit2 u22 ON pf.id_homebase_uii = u22.id 
	WHERE 
		p.uuid = %q AND p.flag_aktif = 1`, uuid)

	return helper.FlatQuery(q)
}

func getPegawaiPNSQuery(uuid string) string {
	q := fmt.Sprintf(`
	SELECT 
		COALESCE(pp.nip_pns,''),
		COALESCE(pp.no_kartu_pegawai,''),
		COALESCE(df.uuid,''),
		COALESCE(df.detail_profesi,''),
		COALESCE(jptt.uuid,''),
		COALESCE(jptt.kd_jenis_ptt,''),
		COALESCE(jptt.jenis_ptt,''),
		COALESCE(pp.instansi_asal,''),
		COALESCE(pgp.uuid,''),
		COALESCE(pgp.kd_pangkat_gol,''),
		COALESCE(pgp.pangkat,''),
		COALESCE(pgp.golongan,''),
		COALESCE(pgp.kd_golongan,''),
		COALESCE(pgp.kd_ruang,''),
		COALESCE(pp.tmt_pangkat_golongan,''),
		COALESCE(jf.uuid,''),
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
		pangkat_golongan_pegawai pgp ON pp.id_pangkat_golongan = pgp.id
	LEFT JOIN
		jabatan_fungsional jf ON pp.id_jabatan_fungsional = jf.id
	LEFT JOIN
		jenis_pegawai_tidak_tetap jptt ON pp.id_jenis_ptt = jptt.id
	LEFT JOIN 
		detail_profesi df ON p.id_detail_profesi = df.id
	WHERE 
		p.uuid = %q AND p.flag_aktif = 1`, uuid)

	return helper.FlatQuery(q)
}

func getStatusPegawaiAktifQuery(uuid string) string {
	q := fmt.Sprintf(`
	SELECT 
		COALESCE(spa.flag_status_aktif,''),
		COALESCE(spa.kd_status,''),
		COALESCE(spa.status,''),
		COALESCE(spa.uuid,''),
		COALESCE(pf.tgl_status_pegawai_aktif,'')
	FROM 
		pegawai p
	LEFT JOIN
		pegawai_fungsional pf ON p.id = pf.id_pegawai
	LEFT JOIN
		status_pegawai_aktif spa ON pf.id_status_pegawai_aktif = spa.id 
	WHERE
		p.uuid = %q AND p.flag_aktif = 1`, uuid)

	return helper.FlatQuery(q)
}

func getPegawaiPribadiQuery(uuid string) string {
	q := fmt.Sprintf(`
	SELECT 
		p.id,
		COALESCE(p.nama,''),
		COALESCE(p.nik,''),
		COALESCE(p.id_agama,0),
		COALESCE(p.kd_agama,''),
		COALESCE(ag.kd_item,''),
		COALESCE(p.id_golongan_darah,0),
		COALESCE(p.kd_golongan_darah,''),
		COALESCE(gd.golongan_darah,''),
		COALESCE(p.jenis_kelamin,''),
		COALESCE(p.id_status_perkawinan, 0),
		COALESCE(sp.kd_status,''),
		COALESCE(p.tempat_lahir,''),
		COALESCE(pdp.tgl_lahir,''),
		COALESCE(p.flag_pensiun,''),
		COALESCE(p.gelar_depan ,''),
		COALESCE(p.gelar_belakang ,''),
		COALESCE(p.nik_ktp,''),
		COALESCE(jp.nama_jenis_pegawai,''),
		COALESCE(kp.kelompok_pegawai,''),
		COALESCE(u2.unit2,''),
		COALESCE(p.user_input,''), 
		COALESCE(p.user_update,''), 
		COALESCE(p.uuid,'') 
	FROM 
		pegawai p
	LEFT JOIN 
		agama ag ON p.id_agama = ag.id 
	LEFT JOIN 
		golongan_darah gd ON p.id_golongan_darah = gd.id 
	LEFT JOIN 
		status_pernikahan sp ON p.id_status_perkawinan = sp.id 
	LEFT JOIN 
		jenis_pegawai jp ON p.id_jenis_pegawai = jp.id 
	LEFT JOIN 
		kelompok_pegawai kp ON p.id_kelompok_pegawai = kp.id 
	LEFT JOIN
		unit2 u2 ON p.id_unit_kerja2 = u2.id 
	LEFT JOIN
		personal_data_pribadi pdp ON p.id_personal_data_pribadi = pdp.id 
	WHERE
		p.uuid = %q AND p.flag_aktif = 1`, uuid)

	return helper.FlatQuery(q)
}

func getPegawaiPendidikanQuery(uuid string) string {
	q := fmt.Sprintf(`
	SELECT 
		COALESCE(pp.uuid,''),
		COALESCE(pp.id,''),
		COALESCE(pp.kd_jenjang,''),
		COALESCE(LPAD(pp.urutan_jenjang,2,"0"),''),
		COALESCE(pp.id_jenjang_pdd_detail_diakui,''),
		COALESCE(d1.kd_detail,''),
		COALESCE(d1.nama_detail,''),
		COALESCE(d1.uuid,''),
		COALESCE(pp.id_jenjang_pdd_detail_terakhir,''),
		COALESCE(d2.kd_detail,''),
		COALESCE(d2.nama_detail,''),
		COALESCE(d2.uuid,''),
		COALESCE(pp.nama_institusi,''),
		COALESCE(pp.jurusan,''),
		COALESCE(pp.tgl_kelulusan,''),
		COALESCE(pp.flag_ijazah_diakui,''),
		COALESCE(pp.flag_ijazah_terakhir,''),
		COALESCE(pp.kd_akreditas,''),
		COALESCE(pp.konsentrasi_bidang_ilmu,''),
		COALESCE(pp.gelar,''),
		COALESCE(pp.nomor_induk,''),
		COALESCE(pp.tahun_masuk,''),
		COALESCE(pp.judul_tugas_akhir,''),
		COALESCE(pp.flag_institusi_luar_negeri,''),
		COALESCE(pp.nomor_ijazah,''),
		COALESCE(pp.tgl_ijazah,''),
		COALESCE(pp.path_ijazah,''),
		COALESCE(pp.flag_ijazah_terverifikasi,''),
		COALESCE(pp.nilai,''),
		COALESCE(pp.jumlah_pelajaran,''),
		COALESCE(pp.path_sk_penyetaraan,''),
		COALESCE(pp.nomor_sk_penyetaraan,''),
		COALESCE(pp.tgl_sk_penyetaraan,''),
		COALESCE(pp.uuid_personal,'')
	FROM pegawai_pendidikan pp
	LEFT JOIN pegawai p ON pp.id_personal_data_pribadi = p.id_personal_data_pribadi
	LEFT JOIN jenjang_pendidikan_detail d1 ON pp.id_jenjang_pdd_detail_diakui = d1.id
	LEFT JOIN jenjang_pendidikan_detail d2 ON pp.id_jenjang_pdd_detail_terakhir = d2.id
	WHERE p.uuid = %q AND pp.flag_aktif = 1`, uuid)

	return helper.FlatQuery(q)
}

func getPegawaiPendidikanPersonalQuery(uuid string) string {
	q := fmt.Sprintf(`
	SELECT 
		COALESCE(pp.uuid,''),
		COALESCE(pp.id,''),
		COALESCE(pp.kd_jenjang,''),
		COALESCE(LPAD(pp.urutan_jenjang,2,"0"),''),
		COALESCE(pp.nama_institusi,''),
		COALESCE(pp.jurusan,''),
		COALESCE(pp.tgl_kelulusan,''),
		COALESCE(pp.flag_ijazah_diakui,''),
		COALESCE(pp.flag_ijazah_terakhir,''),
		COALESCE(pp.kd_akreditas,''),
		COALESCE(pp.konsentrasi_bidang_ilmu,''),
		COALESCE(pp.gelar,''),
		COALESCE(pp.nomor_induk,''),
		COALESCE(pp.tahun_masuk,''),
		COALESCE(pp.judul_tugas_akhir,''),
		COALESCE(pp.flag_institusi_luar_negeri,''),
		COALESCE(pp.nomor_ijazah,''),
		COALESCE(pp.tgl_ijazah,''),
		COALESCE(pp.path_ijazah,''),
		COALESCE(pp.flag_ijazah_terverifikasi,''),
		COALESCE(pp.nilai,''),
		COALESCE(pp.jumlah_pelajaran,''),
		COALESCE(pp.path_sk_penyetaraan,''),
		COALESCE(pp.nomor_sk_penyetaraan,''),
		COALESCE(pp.tgl_sk_penyetaraan,''),
		COALESCE(pp.uuid_personal,'')
	FROM 
		pegawai_pendidikan pp
	LEFT JOIN
		personal_data_pribadi p ON pp.id_personal_data_pribadi = p.id
	WHERE
		p.uuid = %q AND pp.flag_aktif = 1`, uuid)

	return helper.FlatQuery(q)
}

func getPegawaiFilePendidikanQuery(idList ...string) string {
	joinedId := strings.Join(idList, "', '")
	q := fmt.Sprintf(`
	SELECT
		COALESCE(kd_jenis_file,''),
		COALESCE(jenis_file,''),
		COALESCE(path_file_pendidikan,''),
		COALESCE(id_personal_pendidikan,'')
	FROM 
		pegawai_file_pendidikan
	WHERE
		id_personal_pendidikan IN ('%s') AND flag_aktif = 1`, joinedId)

	return helper.FlatQuery(q)
}

func updatePegawaiQuery(pegwaiUpdate model.PegawaiUpdate) string {
	return fmt.Sprintf(``)
}
