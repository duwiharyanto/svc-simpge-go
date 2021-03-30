package repo

import (
	"fmt"
	"svc-insani-go/modules/v1/sk-pengangkatan/model"
)

func getAllPengangkatanTendikQuery(IDPegawai string) string {
	query := fmt.Sprintf(`SELECT id, id_jenis_sk, nomor_sk, tentang_sk, tmt, user_input, tgl_input, user_update, tgl_update, flag_aktif, uuid, id_pegawai FROM sk_pegawai AND sk_pegawai.id_pegawai = %q`, IDPegawai)
	return query
	// return helper.FlatQuery(query)
}
func getSKPengangkatanTendik() string {
	return "SELECT sk_pengangkatan.id_unit_pengangkat, sk_pengangkatan.id_unit_pegawai , sk_pengangkatan.masa_kerja_ril_bulan , sk_pengangkatan.masa_kerja_ril_tahun , sk_pengangkatan.masa_kerja_gaji_bulan , sk_pengangkatan.masa_kerja_gaji_tahun , sk_pengangkatan.id_pangkat_pegawai , sk_pengangkatan.id_golongan_pegawai , sk_pengangkatan.id_status_pengangkatan , sk_pengangkatan.id_jenis_ijazah , sk_pengangkatan.path_ijazah FROM sk_pengangkatan JOIN sk_pegawai ON sk_pegawai.id = sk_pengangkatan.id_sk_pegawai WHERE sk_pengangkatan.flag_aktif=1"
}

func getPengangkatanTendikQuery(IDSKPegawai string) string {
	return fmt.Sprintf(`SELECT sk_pengangkatan.id_unit_pengangkat, sk_pengangkatan.id_unit_pegawai , sk_pengangkatan.masa_kerja_ril_bulan , sk_pengangkatan.masa_kerja_ril_tahun , sk_pengangkatan.masa_kerja_gaji_bulan , sk_pengangkatan.masa_kerja_gaji_tahun , sk_pengangkatan.id_pangkat_pegawai , sk_pengangkatan.id_golongan_pegawai , sk_pengangkatan.id_status_pengangkatan , sk_pengangkatan.id_jenis_ijazah , sk_pengangkatan.path_ijazah FROM sk_pengangkatan JOIN sk_pegawai ON sk_pegawai.id = sk_pengangkatan.id_sk_pegawai WHERE sk_pengangkatan.flag_aktif=1 AND sk_pengangkatan.id_sk_pegawai = %q`, IDSKPegawai)
}

func createSKPengangkatanTendikQuery(skPengangkatanTendik model.SKPengangkatanTendik, IDSKPegawai string) string {
	return fmt.Sprintf(`INSERT INTO sk_pengangkatan_tendik (id, id_kelompok_sk_pengangkatan, id_unit_pengangkat, id_unit_pegawai, id_jabatan_fungsional, gaji_pokok, masa_kerja_ril_bulan, masa_kerja_ril_tahun, masa_kerja_gaji_bulan, masa_kerja_gaji_tahun, masa_kerja_diakui_bulan, masa_kerja_diakui_tahun, id_pangkat_golongan_pegawai, id_status_pengangkatan, id_jenis_ijazah, tanggal_ditetapkan, path_sk, user_input, user_update, id_sk_pegawai, uuid) VALUES(UUID_SHORT(), %q, %q, %q, %q, %d, %d, %d, %d, %d, %d, %d, %q, %q, %q, %q, %q, %q, %q, %q, UUID())`,
		skPengangkatanTendik.IDKelompokSKPengangkatan,
		skPengangkatanTendik.IDUnitPengangkat,
		skPengangkatanTendik.IDUnitPegawai,
		skPengangkatanTendik.IDJabatanFungsional,
		skPengangkatanTendik.GajiPokok,
		skPengangkatanTendik.MasaRilBulan,
		skPengangkatanTendik.MasaRilTahun,
		skPengangkatanTendik.MasaGajiBulan,
		skPengangkatanTendik.MasaGajiTahun,
		skPengangkatanTendik.MasaKerjaDiakuiBulan,
		skPengangkatanTendik.MasaKerjaDiakuiTahun,
		skPengangkatanTendik.IDPangkatGolonganPegawai,
		skPengangkatanTendik.IDStatusPengangkatan,
		skPengangkatanTendik.IDJenisIjazah,
		skPengangkatanTendik.TanggalDitetapkan,
		skPengangkatanTendik.PathSKTendik,
		skPengangkatanTendik.UserInput,
		skPengangkatanTendik.UserUpdate,
		IDSKPegawai,
	)
}

func getDetailSKPengangkatanTendikQuery(UUIDSKPengangkatanTendik string) string {
	return fmt.Sprintf(`SELECT
	COALESCE(k.nama, ''), COALESCE(k.nik, ''), COALESCE(k.tempat_lahir, ''), COALESCE(k.tgl_lahir, ''),
	COALESCE(c.kd_jenis_sk, ''), COALESCE(c.nama_sk, ''), COALESCE(c.uuid, ''),
	COALESCE(d.kd_kelompok_pegawai, ''), COALESCE(d.kelompok_pegawai, ''), COALESCE(d.uuid, ''),
	COALESCE(a.nomor_sk, ''),
	COALESCE(e.kd_unit_kerja, ''), COALESCE(e.nama_unit_kerja, ''), COALESCE(e.uuid, ''),
	COALESCE(a.tmt, ''),
	COALESCE(f.kd_unit_kerja, ''), COALESCE(f.nama_unit_kerja, ''), COALESCE(f.uuid, ''),
	COALESCE(g.kd_fungsional, ''), COALESCE(g.fungsional, ''), COALESCE(g.uuid, ''),
	COALESCE(h.pangkat, ''), COALESCE(h.golongan, ''), COALESCE(h.uuid, ''),
	COALESCE(b.gaji_pokok, 0), 
	COALESCE(b.masa_kerja_diakui_tahun, '0'), COALESCE(b.masa_kerja_diakui_bulan, '0'), 
	COALESCE(b.masa_kerja_ril_tahun, '0'), COALESCE(b.masa_kerja_ril_bulan, '0'), 
	COALESCE(b.masa_kerja_gaji_tahun, '0'), COALESCE(b.masa_kerja_gaji_bulan, '0'), 
	COALESCE(i.kd_status_pengangkatan, ''), COALESCE(i.status_pengangkatan, ''), COALESCE(i.uuid, ''),
	COALESCE(j.kd_jenis_ijazah, ''), COALESCE(j.jenis_ijazah, ''), COALESCE(j.uuid, ''),
	COALESCE(b.tanggal_ditetapkan, ''), COALESCE(b.path_sk, ''), COALESCE(b.uuid, '')
	FROM sk_pegawai a
	JOIN sk_pengangkatan_tendik b ON a.id = b.id_sk_pegawai
	LEFT JOIN jenis_sk c ON a.id_jenis_sk = c.id
	LEFT JOIN kelompok_pegawai d ON b.id_kelompok_sk_pengangkatan = d.id
	LEFT JOIN unit_kerja e ON b.id_unit_pengangkat = e.id
	LEFT JOIN unit_kerja f ON b.id_unit_pegawai = f.id
	LEFT JOIN jabatan_fungsional g ON b.id_jabatan_fungsional = g.id
	LEFT JOIN pangkat_golongan_pegawai h ON b.id_pangkat_golongan_pegawai = h.id
	LEFT JOIN status_pengangkatan i ON b.id_status_pengangkatan = i.id
	LEFT JOIN jenis_ijazah j ON b.id_jenis_ijazah = j.id
	LEFT JOIN pegawai k ON a.id_pegawai = k.id
	WHERE b.flag_aktif = 1 AND b.uuid = %q`, UUIDSKPengangkatanTendik)
}

// kurang
// func getDetailSKPengangkatanTendikQuery(UUIDSKPengangkatanTendik string) string {
// 	return fmt.Sprintf(`SELECT sk_pengangkatan_tendik.id_unit_pengangkat, sk_pengangkatan_tendik.id_unit_pegawai, sk_pengangkatan_tendik.masa_kerja_ril_bulan, sk_pengangkatan_tendik.masa_kerja_ril_tahun, sk_pengangkatan_tendik.masa_kerja_gaji_bulan, sk_pengangkatan_tendik.masa_kerja_gaji_tahun,
// 	sk_pengangkatan_tendik.id_pangkat_golongan_pegawai, sk_pengangkatan_tendik.id_status_pengangkatan, sk_pengangkatan_tendik.id_jenis_ijazah, sk_pengangkatan_tendik.tanggal_ditetapkan, sk_pengangkatan_tendik.path_sk, sk_pengangkatan_tendik.user_input, sk_pengangkatan_tendik.user_update, sk_pengangkatan_tendik.uuid, sk_pengangkatan_tendik.id_sk_pegawai
// 		FROM hcm_insani.sk_pengangkatan_tendik
// 		JOIN hcm_insani.sk_pegawai ON sk_pengangkatan_tendik.id_sk_pegawai=sk_pegawai.id
// 		AND sk_pengangkatan_tendik.uuid = %q`, UUIDSKPengangkatanTendik)
// }

func updateSKPengangkatanTendikQuery(skPengangkatanTendik model.SKPengangkatanTendik, IDSKPegawai string) string {
	return fmt.Sprintf(`UPDATE sk_pengangkatan_tendik
	SET id=?, id_unit_pengangkat=?, id_unit_pegawai=?, masa_kerja_ril_bulan=?, masa_kerja_ril_tahun=?, masa_kerja_gaji_bulan=?, masa_kerja_gaji_tahun=?, id_pangkat_golongan_pegawai=?, id_status_pengangkatan=?, id_jenis_ijazah=?, tanggal_ditetapkan=?, path_sk=?, user_input='system', tgl_input=CURRENT_TIMESTAMP, user_update=?, tgl_update=CURRENT_TIMESTAMP, flag_aktif=1, id_sk_pegawai=?
	WHERE uuid=?
	`)
}

func deleteSKPengangkatanTendikByUUID() string {
	return "DELETE FROM sk_pengangkatan_tendik WHERE uuid=?"
}
