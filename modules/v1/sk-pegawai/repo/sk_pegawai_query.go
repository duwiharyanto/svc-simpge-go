package repo

import (
	"fmt"
	"svc-insani-go/modules/v1/sk-pegawai/model"
)

func createSKPegawaiQuery(skPegawai model.SKPegawai) string {
	query := fmt.Sprintf(`INSERT INTO sk_pegawai (id, id_pegawai, id_jenis_sk, nomor_sk, tentang_sk, tmt, user_input, user_update, uuid) VALUES (%q, %q, %q, %q, %q, %q, %q, %q, %q)`,
		skPegawai.ID,
		skPegawai.IDPegawai,
		skPegawai.IDJenisSK,
		skPegawai.NomorSK,
		skPegawai.TentangSK,
		skPegawai.TMT,
		skPegawai.UserInput,
		skPegawai.UserUpdate,
		skPegawai.UUID,
	)
	return query
	// return helper.FlatQuery(query)
}

func getAllSKPegawaibyUUIDQuery(IDPegawai string) string {
	query := fmt.Sprintf(`SELECT id_jenis_sk, nomor_sk, tentang_sk, tmt, uuid FROM sk_pegawai WHERE flag_aktif=1 AND id_pegawai = %q`, IDPegawai)
	return query
	// return helper.FlatQuery(query)
}

func getSKPengangkatan() string {
	return "SELECT sk_pengangkatan.kd_unit_pengangkat, sk_pengangkatan.uuid_unit_pengangkat , sk_pengangkatan.kd_unit_pegawai , sk_pengangkatan.uuid_unit_pegawai , sk_pengangkatan.masa_kerja_ril_bulan , sk_pengangkatan.masa_kerja_ril_tahun , sk_pengangkatan.masa_kerja_gaji_bulan , sk_pengangkatan.masa_kerja_gaji_tahun , sk_pengangkatan.kd_pangkat_pegawai , sk_pengangkatan.uuid_pangkat_pegawai ,sk_pengangkatan.kd_golongan_pegawai , sk_pengangkatan.uuid_golongan_pegawai , sk_pengangkatan.kd_status_pengangkatan , sk_pengangkatan.uuid_status_pengangkatan, sk_pengangkatan.kd_jenis_ijazah , sk_pengangkatan.uuid_jenis_ijazah , sk_pengangkatan.path_ijazah FROM sk_pengangkatan JOIN pegawai ON sk_pengangkatan.nik= pegawai.nik JOIN sk_pegawai ON sk_pegawai.uuid = sk_pengangkatan.uuid_sk_utama WHERE sk_pengangkatan.flag_aktif=1"
}

func getSKPengangkatanQuery(IDSKPegawai string) string {
	return fmt.Sprintf(`SELECT sk_pengangkatan.kd_unit_pengangkat, sk_pengangkatan.uuid_unit_pengangkat , sk_pengangkatan.kd_unit_pegawai , sk_pengangkatan.uuid_unit_pegawai , sk_pengangkatan.masa_kerja_ril_bulan , sk_pengangkatan.masa_kerja_ril_tahun , sk_pengangkatan.masa_kerja_gaji_bulan , sk_pengangkatan.masa_kerja_gaji_tahun , sk_pengangkatan.kd_pangkat_pegawai , sk_pengangkatan.uuid_pangkat_pegawai ,sk_pengangkatan.kd_golongan_pegawai , sk_pengangkatan.uuid_golongan_pegawai , sk_pengangkatan.kd_status_pengangkatan , sk_pengangkatan.uuid_status_pengangkatan, sk_pengangkatan.kd_jenis_ijazah , sk_pengangkatan.uuid_jenis_ijazah , sk_pengangkatan.path_ijazah FROM sk_pengangkatan WHERE sk_pengangkatan.flag_aktif=1 AND sk_pengangkatan.id_sk_pegawai = %q`, IDSKPegawai)
}

// func getSKPengangkatanDetailQuery(skPegawai model.SKPegawai) string {
// 	query := fmt.Sprintf(`SELECT id_jenis_sk, nomor_sk, tentang_sk, tmt, uuid FROM sk_pegawai WHERE flag_aktif=1 AND id_pegawai = %q`, IDPegawai)
// 	return query
// }
