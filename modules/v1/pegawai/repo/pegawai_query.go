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

func getListAllPegawaiPrivateQuery(req *model.PegawaiPrivateRequest) string {
	var nikFilterQuery string
	if req.Nik != "" {
		nikFilterQuery = fmt.Sprintf("AND p.nik IN (%s) ", req.Nik)
	}
	var namaFilterQuery string
	if req.Nama != "" {
		namaFilterQuery = fmt.Sprintf("AND p.nama like '%%%s%%' ", req.Nama)
	}
	var kdJenisPegawaiFilterQuery string
	if req.KdJenisPegawai != "" {
		kdJenisPegawaiFilterQuery = fmt.Sprintf("AND jpeg.kd_jenis_pegawai = '%s'", req.KdJenisPegawai)
	}
	var kdKelompokFilterQuery string
	if req.KdKelompokPegawai != "" {
		kdKelompokFilterQuery = fmt.Sprintf("AND kp.kd_kelompok_pegawai = '%s'", req.KdKelompokPegawai)
	}
	var kdIndukKerjaFilterQuery string
	if req.KdIndukKerja != "" {
		kdIndukKerjaFilterQuery = fmt.Sprintf("AND p.kd_unit1 = '%s'", req.KdIndukKerja)
	}

	return fmt.Sprintf(`SELECT
	COALESCE(p.id,0) id_pegawai,
	COALESCE(p.id_personal_data_pribadi,0),
	p.nama,
	p.nik,
	COALESCE(jpeg.nama_jenis_pegawai,'') jenis_pegawai,
	COALESCE(jpeg.id,0) id_jenis_pegawai,
	COALESCE(jpeg.kd_jenis_pegawai,''),
	COALESCE(kp.kelompok_pegawai,''),
	COALESCE(kp.id,0) id_kelompok_pegawai,
	COALESCE(kp.kd_kelompok_pegawai,''),
	COALESCE(kp.id_kategori_kelompok_pegawai,0),
	COALESCE(kkp.kd_kategori_kelompok_pegawai,''),
	COALESCE(g.golongan,''),
	COALESCE(pgp.id_golongan,0) id_golongan,
	COALESCE(pgp.kd_golongan,'') kd_golongan,
	COALESCE(g1.golongan,'') golongan_negara,
	COALESCE(pgppns.id_golongan,0) id_golongan_negara,
	COALESCE(pgppns.kd_golongan,'') kd_golongan_negara,
	COALESCE(r.nama_ruang,'') ruang,
	COALESCE(r.id,0) id_ruang,
	COALESCE(r.kd_ruang,'') kd_ruang,
	COALESCE(r_ngr.nama_ruang,'') ruang_negara,
	COALESCE(r_ngr.id,0) id_ruang_negara,
	COALESCE(r_ngr.kd_ruang,'') kd_ruang_negara,
	COALESCE(u2.unit2,'') unit_kerja,
	COALESCE(u2.id,0) id_unit,
	COALESCE(u2.kd_unit2,'') kd_unit,
	COALESCE(u1.unit1,'') induk_kerja,
	COALESCE(u1.id,0) id_induk_kerja,
	COALESCE(u1.kd_unit1,'') kd_induk_kerja,
	COALESCE(pf.id_status_pegawai_aktif,0),
	COALESCE(spa.status,'') status_pegawai_aktif,
	COALESCE(spa.kd_status,'') kd_status_pegawai_aktif,
	COALESCE(sp.status_pegawai,'') status_pegawai,
	COALESCE(p.id_status_pegawai,0),
	COALESCE(p.kd_status_pegawai,''),
	COALESCE(p.jenis_kelamin,''),
	COALESCE(jf.fungsional ,'') jabatan_fungsional_yayasan,
	COALESCE(pf.id_jabatan_fungsional,0) id_jabatan_fungsional_yayasan,
	COALESCE(pf.kd_jabatan_fungsional,'') kd_jabatan_fungsional_yayasan,
	COALESCE(jf2.fungsional ,'') jabatan_fungsional_negara,
	COALESCE(pp.id_jabatan_fungsional ,0) id_jabatan_fungsional_negara,
	COALESCE(pp.kd_jabatan_fungsional ,'') kd_jabatan_fungsional_negara,
	COALESCE(p.id_detail_profesi,0),
	COALESCE(dp.detail_profesi,''),
	COALESCE(jp.id,0) id_jenjang_pendidikan,
	COALESCE(jp.kd_jenjang,'') kd_jenjang_pendidikan,
	COALESCE(jp.jenjang,'') jenjang_pendidikan,
	COALESCE(pf.tmt_sk_pertama,'') tmt_sk_pertama,
	COALESCE(pf.masa_kerja_awal_kepegawaian_tahun,'') masa_kerja_tahun,
	COALESCE(pf.masa_kerja_awal_kepegawaian_bulan,'') masa_kerja_bulan,
	COALESCE((SELECT COUNT(*) FROM hcm_personal.personal_hubungan_keluarga phk WHERE phk.id_personal_data_pribadi =  p.id_personal_data_pribadi AND phk.kd_hubungan_keluarga IN ('AAK','AT','AN','SUA','IST') AND phk.flag_aktif = 1),'') jumlah_keluarga,
	COALESCE((SELECT COUNT(*) from hcm_personal.personal_hubungan_keluarga phk WHERE phk.id_personal_data_pribadi = p.id_personal_data_pribadi AND phk.kd_hubungan_keluarga IN ('AAK','AT','AN') AND phk.flag_aktif = 1),'') jumlah_anak,
	COALESCE((SELECT DISTINCT pi.npwp from hcm_personal.personal_identitas pi WHERE pi.id_personal_data_pribadi = p.id_personal_data_pribadi AND pi.flag_aktif = 1),'') npwp,
	COALESCE(spn.id,0) id_status_nikah,
	COALESCE(spn.kd_status,'') kd_status_nikah,
	COALESCE(spn.status,'') status_nikah,
	COALESCE(p.nik_suami_istri,''),
	COALESCE(p.nik_ktp,'') nik_ktp,
	COALESCE((SELECT pdp.flag_ptkp from hcm_tanggungan.personal_data_pribadi pdp WHERE pdp.id = p.id_personal_data_pribadi AND pdp.flag_aktif = 1),0) flag_klaim_tanggungan,
	COALESCE(p.flag_pensiun,0),
	COALESCE(p.flag_meninggal,0),
	COALESCE((SELECT DISTINCT phk.flag_sekantor from hcm_personal.personal_hubungan_keluarga phk WHERE phk.id_personal_data_pribadi = p.id_personal_data_pribadi AND phk.kd_hubungan_keluarga in ('SUA','IST') AND phk.flag_aktif = 1),0) flag_suami_istri_sekantor,
	COALESCE((CASE WHEN pf.id_jabatan_fungsional != '' OR pp.id_jabatan_fungsional != '' THEN 1 END ),0) is_fungsional,
	COALESCE((SELECT COUNT(*) from hcm_organisasi.pejabat_organisasi po JOIN hcm_organisasi.unit u ON u.id = po.id_unit WHERE po.id_pegawai = p.id AND po.flag_aktif =1),'') is_struktural
	from
	pegawai p
	LEFT JOIN
		jenis_pegawai jpeg ON p.kd_jenis_pegawai = jpeg.kd_jenis_pegawai
	LEFT JOIN
		kelompok_pegawai kp ON p.kd_kelompok_pegawai = kp.kd_kelompok_pegawai
	LEFT JOIN
		kategori_kelompok_pegawai kkp on kp.id_kategori_kelompok_pegawai = kkp.id
	LEFT JOIN
		pegawai_pns pp on pp.id_pegawai = p.id
	LEFT JOIN
		pegawai_fungsional pf ON p.id = pf.id_pegawai
	LEFT JOIN
		pangkat_golongan_pegawai pgp ON pf.id_pangkat_golongan = pgp.id
	LEFT JOIN
		pangkat_golongan_pegawai pgppns ON pp.id_pangkat_golongan = pgppns.id
	LEFT JOIN
		golongan g on pgp.id_golongan = g.id
	LEFT JOIN
		golongan g1 on pgppns.id_golongan = g1.id
	LEFT JOIN
		ruang r on pgp.id_ruang = r.id
	LEFT JOIN
		ruang r_ngr on pgppns.id_ruang = r_ngr.id
	LEFT JOIN
		unit1 u1 ON p.kd_unit1 = u1.kd_unit1
	LEFT JOIN
		unit2 u2 ON p.kd_unit2  = u2.kd_unit2
	LEFT JOIN
		status_pegawai_aktif spa ON spa.id = pf.id_status_pegawai_aktif
	LEFT JOIN
		status_pegawai sp ON p.kd_status_pegawai = sp.kd_status_pegawai
	LEFT JOIN
		jabatan_fungsional jf ON pf.id_jabatan_fungsional = jf.id
	LEFT JOIN
		jabatan_fungsional jf2 ON pp.id_jabatan_fungsional = jf2.id
	LEFT JOIN
		jenjang_pendidikan jp ON p.id_pendidikan_terakhir = jp.id
	LEFT JOIN
		status_pernikahan spn ON p.kd_status_perkawinan = spn.kd_status
	LEFT JOIN
		detail_profesi dp ON p.id_detail_profesi = dp.id
	WHERE p.flag_aktif=1 %s %s %s %s %s`, nikFilterQuery, namaFilterQuery, kdJenisPegawaiFilterQuery, kdKelompokFilterQuery, kdIndukKerjaFilterQuery)

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
		COALESCE(kp2.uuid,''),
		COALESCE(kp2.id, 0),
		COALESCE(kp2.kd_kelompok_pegawai,''),
		COALESCE(kp2.kelompok_pegawai,''),
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
		COALESCE(pf.masa_kerja_awal_kepegawaian_tahun,''),
		COALESCE(pf.masa_kerja_awal_kepegawaian_bulan,''),
		COALESCE(pf.masa_kerja_awal_pensiun_tahun,''),
		COALESCE(pf.masa_kerja_awal_pensiun_bulan,''),
		COALESCE(pf.angka_kredit,''),
		COALESCE(pf.nomor_sertifikasi,''),
		COALESCE(pf.nidn,''),
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
		kelompok_pegawai kp2 ON p.id_kelompok_pegawai_payroll = kp2.id
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
		COALESCE(pf.nomor_surat_kontrak,''),
		COALESCE(pf.tmt_surat_kontrak,''),
		COALESCE(pf.tgl_surat_kontrak,''),
		COALESCE(pf.tmt_awal_kontrak,''),
		COALESCE(pf.tmt_akhir_kontrak,''),
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
		COALESCE(pp.nomor_pangkat_golongan,''),
		COALESCE(jf.uuid,''),
		COALESCE(jf.kd_fungsional,''),
		COALESCE(jf.fungsional,''),
		COALESCE(pp.tmt_jabatan,''),
		COALESCE(pp.nomor_jabatan_fungsional,''),
		COALESCE(pp.masa_kerja_tahun,''),
		COALESCE(pp.masa_kerja_bulan,''),
		COALESCE(pp.masa_kerja_golongan_tahun,''),
		COALESCE(pp.masa_kerja_golongan_bulan,''),
		COALESCE(pp.angka_kredit,''),
		COALESCE(pp.nomor_pak,''),
		COALESCE(pp.tmt_pak,''),
		COALESCE(pp.nomor_sk_pensiun,''),
		COALESCE(pp.tmt_sk_pensiun,''),
		COALESCE(pp.masa_kerja_pensiun_tahun,''),
		COALESCE(pp.masa_kerja_pensiun_bulan,''),
		COALESCE(pp.keterangan,''),
		COALESCE(pp.nira,'')
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
		COALESCE(jd1.kd_detail,''),
		COALESCE(jd2.kd_detail,''),
		COALESCE(pdp.path_file_foto_personal,''),
		COALESCE(p.user_input,''), 
		COALESCE(p.user_update,''), 
		COALESCE(p.uuid,'') 
	FROM pegawai p
	LEFT JOIN agama ag ON p.id_agama = ag.id 
	LEFT JOIN golongan_darah gd ON p.id_golongan_darah = gd.id 
	LEFT JOIN status_pernikahan sp ON p.id_status_perkawinan = sp.id 
	LEFT JOIN jenjang_pendidikan_detail jd1 ON p.id_status_pendidikan_masuk = jd1.id 
	LEFT JOIN jenjang_pendidikan_detail jd2 ON p.id_jenis_pendidikan = jd2.id 
	LEFT JOIN jenis_pegawai jp ON p.id_jenis_pegawai = jp.id 
	LEFT JOIN kelompok_pegawai kp ON p.id_kelompok_pegawai = kp.id 
	LEFT JOIN unit2 u2 ON p.id_unit_kerja2 = u2.id 
	LEFT JOIN personal_data_pribadi pdp ON p.id_personal_data_pribadi = pdp.id 
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

func getPegawaiByNik(nik string) string {
	q := fmt.Sprintf(`SELECT
	p.nama,
	COALESCE(p.gelar_depan,''),
	COALESCE(p.gelar_belakang,''),
	p.nik,
	p.tempat_lahir,
	p.tgl_lahir,
	p.jenis_kelamin,
	COALESCE(jp.kd_pendidikan_simpeg,''),
	COALESCE(sp.kd_status_pegawai,''),
	COALESCE(sp.status_pegawai,''),
	COALESCE(kp.kd_kelompok_pegawai,''),
	COALESCE(kp.kelompok_pegawai,''),
	COALESCE(pgp.kd_pangkat_gol,''),
	COALESCE(pgp.pangkat,''),
	COALESCE(pgp.kd_golongan,''),
	COALESCE(pgp.golongan,''),
	COALESCE(pgp.kd_ruang,''),
	COALESCE(pf.tmt_pangkat_golongan,''),
	COALESCE(jf.kd_fungsional,''),
	COALESCE(jf.fungsional,''),
	COALESCE(pf.tmt_jabatan,''),
	COALESCE(u1.kd_unit1,''),
	COALESCE(u1.unit1,''),
	COALESCE(u2.kd_unit2,''),
	COALESCE(u2.unit2,'')
	from
	pegawai p
	LEFT JOIN 
		jenjang_pendidikan jp ON p.id_pendidikan_terakhir = jp.id
	LEFT JOIN 
		status_pegawai sp ON p.id_status_pegawai = sp.id 
	LEFT JOIN
		kelompok_pegawai kp ON p.id_kelompok_pegawai = kp.id 
	LEFT JOIN
		pegawai_fungsional pf ON p.id = pf.id_pegawai 
	LEFT JOIN
		pangkat_golongan_pegawai pgp ON pf.id_pangkat_golongan = pgp.id 
	LEFT JOIN 
		jabatan_fungsional jf ON pf.id_jabatan_fungsional = jf.id 
	LEFT JOIN
		unit1 u1 ON p.id_unit_kerja1 = u1.id 
	LEFT JOIN
		unit2 u2 ON p.id_unit_kerja2 = u2.id 
	WHERE p.flag_aktif=1 AND p.nik = %q`, nik)
	return helper.FlatQuery(q)
}
