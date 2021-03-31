package repo

import (
	"fmt"
	"svc-insani-go/modules/v1/sk-pengangkatan/model"
)

func createSKPengangkatanDosenQuery(skPengangkatanDosen model.SKPengangkatanDosen, IDSKPegawai string) string {
	return fmt.Sprintf(`INSERT INTO sk_pengangkatan_dosen
	(id_jenis_sk_pengangkatan, masa_kerja_diakui_bulan_lama, masa_kerja_diakui_tahun_baru, id_pangkat_gol_lama, id_jabatan_fungsional_lama, id_induk_kerja_baru, id_pangkat_gol_baru, gaji_pokok, tunjangan_beras, tunjangan_khusus, sks_mengajar, bantuan_komunikasi, tunjangan_tahunan, masa_kerja_riil_bulan_baru, masa_kerja_riil_tahun_baru, masa_kerja_gaji_bulan_baru, masa_kerja_gaji_tahun_baru, id_jabatan_penetap, id_pegawai_penetap, tgl_ditetapkan, id_unit_kerja, id_jabatan_fungsional, id_ijazah_pendidikan, instansi_kerja, tgl_berakhir, jangka_waktu_evaluasi, id_sk_pegawai, path_sk, user_input, user_update)
	VALUES(%q, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q, %q)`,
		skPengangkatanDosen.IDJenisSKPengangkatan,
		skPengangkatanDosen.MasaKerjaDiakuiBulanLama,
		skPengangkatanDosen.MasaKerjaDiakuiTahunBaru,
		skPengangkatanDosen.IDPangkatGolonganPegawaiLama,
		skPengangkatanDosen.IDJabatanFungsionalLama,
		skPengangkatanDosen.IDIndukKerjaBaru,
		skPengangkatanDosen.IDPangkatGolonganPegawaiBaru,
		skPengangkatanDosen.GajiPokok,
		skPengangkatanDosen.TunjanganBeras,
		skPengangkatanDosen.TunjanganKhusus,
		skPengangkatanDosen.SKSMengajar,
		skPengangkatanDosen.BantuanKomunikasi,
		skPengangkatanDosen.TunjanganTahunan,
		skPengangkatanDosen.MasaKerjaRilBulanBaru,
		skPengangkatanDosen.MasaKerjaRilTahunBaru,
		skPengangkatanDosen.MasaKerjaGajiBulanBaru,
		skPengangkatanDosen.MasaKerjaGajiTahunBaru,
		skPengangkatanDosen.IDJabatanPenetap,
		skPengangkatanDosen.IDPegawaiPenetap,
		skPengangkatanDosen.TanggalDitetapkan,
		skPengangkatanDosen.IDUnitKerja,
		skPengangkatanDosen.IDJabatanFungsionalBaru,
		skPengangkatanDosen.IDJenisIjazah,
		skPengangkatanDosen.InstansiKerja,
		skPengangkatanDosen.TglBerakhir,
		skPengangkatanDosen.JangkaWaktuEvaluasi,
		IDSKPegawai,
		skPengangkatanDosen.PathSKDosen,
		skPengangkatanDosen.UserInput,
		skPengangkatanDosen.UserUpdate,
	)
}

func createPivotMakul(idSkPengangkatanDosen, idMakul, userUpdate string) string {
	return fmt.Sprintf(`INSERT INTO pivot_makul_skpengangkatan (id_sk_pegawai, id_makul, user_input, user_update) VALUES (%q, %q, %q, %q)`,
		idSkPengangkatanDosen,
		idMakul,
		userUpdate,
		userUpdate,
	)
}
func getDetailSKPengangkatanDosenQuery(UUIDSKPengangkatanDosen string) string {
	return fmt.Sprintf(`SELECT COALESCE(sk_pengangkatan_dosen.id_jenis_sk_pengangkatan,''), COALESCE(sk_pengangkatan_dosen.masa_kerja_diakui_bulan_lama,''), 
	COALESCE(sk_pengangkatan_dosen.masa_kerja_diakui_tahun_baru,''), COALESCE(sk_pengangkatan_dosen.id_pangkat_gol_lama,''), 
	COALESCE(sk_pengangkatan_dosen.id_jabatan_fungsional_lama,''), COALESCE(sk_pengangkatan_dosen.id_induk_kerja_baru,''), COALESCE(sk_pengangkatan_dosen.id_pangkat_gol_baru,''), gaji_pokok, 
	COALESCE(sk_pengangkatan_dosen.tunjangan_beras,''), COALESCE(sk_pengangkatan_dosen.tunjangan_khusus,''), COALESCE(sk_pengangkatan_dosen.sks_mengajar,''), 
	COALESCE(sk_pengangkatan_dosen.bantuan_komunikasi,''), COALESCE(sk_pengangkatan_dosen.tunjangan_tahunan,''), COALESCE(sk_pengangkatan_dosen.masa_kerja_riil_bulan_baru,''), 
	COALESCE(sk_pengangkatan_dosen.masa_kerja_riil_tahun_baru,''), COALESCE(sk_pengangkatan_dosen.masa_kerja_gaji_bulan_baru,''), COALESCE(sk_pengangkatan_dosen.masa_kerja_gaji_tahun_baru,''), 
	COALESCE(sk_pengangkatan_dosen.id_jabatan_penetap,''), COALESCE(sk_pengangkatan_dosen.id_pegawai_penetap,''), 
	COALESCE(sk_pengangkatan_dosen.tgl_ditetapkan,''), COALESCE(sk_pengangkatan_dosen.id_unit_kerja,''), COALESCE(sk_pengangkatan_dosen.id_jabatan_fungsional,''), 
	COALESCE(sk_pengangkatan_dosen.id_ijazah_pendidikan,''), COALESCE(sk_pengangkatan_dosen.instansi_kerja,''), COALESCE(sk_pengangkatan_dosen.tgl_berakhir,''), 
	COALESCE(sk_pengangkatan_dosen.jangka_waktu_evaluasi,''), COALESCE(sk_pengangkatan_dosen.path_sk,''), COALESCE(sk_pengangkatan_dosen.user_input,''), COALESCE(sk_pengangkatan_dosen.user_update,''), 
	sk_pengangkatan_dosen.uuid, sk_pengangkatan_dosen.id_sk_pegawai
	FROM sk_pengangkatan_dosen JOIN sk_pegawai ON sk_pengangkatan_dosen.id_sk_pegawai = sk_pegawai.id AND sk_pengangkatan_dosen.uuid = %q`, UUIDSKPengangkatanDosen)
}

func deleteSKPengangkatanDosenByUUID() string {
	return "DELETE FROM sk_pengangkatan_dosen WHERE uuid=?"
}

func updateSKPengangkatanDosenQuery(skPengangkatanDosen model.SKPengangkatanDosen, IDSKPegawai string) string {
	return fmt.Sprintf(`UPDATE sk_pengangkatan_tendik
	SET id=?, id_unit_pengangkat=?, id_unit_pegawai=?, masa_kerja_ril_bulan=? masa_kerja_ril_tahun=? masa_kerja_gaji_bulan=? masa_kerja_gaji_tahun=? id_pangkat_golongan_pegawai=?, id_status_pengangkatan=?, id_jenis_ijazah=?, tanggal_ditetapkan=? path_sk=? user_input='system', tgl_input=CURRENT_TIMESTAMP, user_update=? tgl_update=CURRENT_TIMESTAMP, flag_aktif=1, id_sk_pegawai=?
	WHERE uuid=?`)
}
