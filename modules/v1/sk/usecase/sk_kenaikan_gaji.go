package usecase

import (
	"net/http"
	"svc-insani-go/app"

	"github.com/labstack/echo"
)

func HandleUpdateSkKenaikanGaji(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Perubahan sk kenaikan gaji berhasil disimpan"})
	}
	return echo.HandlerFunc(h)
}

func HandleGetSkKenaikanGajiDummy(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		return c.JSONBlob(http.StatusOK, []byte(dummySkKenaikanGajiDetail))
	}
	return echo.HandlerFunc(h)
}

const dummySkKenaikanGajiDetail = `{
		"nama_pegawai": "Nama Dummy",
		"nik_pegawai": "091002120",
		"ttl": "Kota Padang, 3 November 1946",
		"uuid_jenis_sk": "ebc9e2c0-ee60-11ea-8c77-7eb0d4a3c7a0",
		"kd_jenis_sk": "4",
		"jenis_sk": "Kenaikan Gaji Berkala",
		"uuid_kelompok_sk_kgb": "46994cd3-ec0a-11ea-8c77-7eb0d4a3c7a0",
		"kd_kelompok_sk_kgb": "32",
		"kelompok_sk_kgb": "Pustakawan Tidak Tetap Universitas",
		"nomor_sk": "536/A.II/YBW.IX/2009",
		"uuid_jabatan_fungsional_lama": "aeb51169-2fc6-11eb-a014-7eb0d4a3c7a0",
		"kd_jabatan_fungsional_lama": "06",
		"jabatan_fungsional_lama": "Asisten Ahli",
		"uuid_pangkat_golongan_pegawai_lama": "c6101e45-09e3-11eb-8c77-7eb0d4a3c7a0",
		"pangkat_lama": "Juru muda",
		"golongan_lama": "I/A",
		"tmt_kgb_lama": "2019-01-05",
		"tmt_kgb_lama_idn": "5 Januari 2019",
		"gaji_pokok_lama": 1600000,
		"masa_kerja_ril_bulan_lama": 3,
		"masa_kerja_ril_tahun_lama": 10,
		"masa_kerja_gaji_bulan_lama": 1,
		"masa_kerja_gaji_tahun_lama": 12,
		"uuid_jabatan_fungsional_baru": "aeb51169-2fc6-11eb-a014-7eb0d4a3c7a0",
		"kd_jabatan_fungsional_baru": "06",
		"jabatan_fungsional_baru": "Asisten Ahli",
		"uuid_pangkat_golongan_pegawai_baru": "c6101e45-09e3-11eb-8c77-7eb0d4a3c7a0",
		"pangkat_baru": "Juru muda",
		"golongan_baru": "I/A",
		"tmt_kgb_baru": "2020-01-05",
		"tmt_kgb_baru_idn": "5 Januari 2020",
		"gaji_pokok_baru": 2600000,
		"masa_kerja_ril_bulan_baru": 3,
		"masa_kerja_ril_tahun_baru": 10,
		"masa_kerja_gaji_bulan_baru": 1,
		"masa_kerja_gaji_tahun_baru": 12,
		"tgl_kgb_berikutnya": "2020-01-05",
		"tgl_kgb_berikutnya_idn": "5 Januari 2020",
		"uuid_status_pegawai": "47dd67dc-0479-11eb-8c77-7eb0d4a3c7a0",
		"kd_status_pegawai": "000",
		"status_pegawai": "Pegawai Administratif Tetap UII",
		"uuid_unit_kerja": "05577f32-e996-11e9-8f20-506b8da96a87",
		"kd_unit_kerja": "131",
		"unit_kerja": "Fakultas Kedokteran",
		"uuid_jabatan_penetap": "05577f32-e996-11e9-8f20-506b8da96a87",
		"kd_jabatan_penetap": "001",
		"jabatan_penetap": "Rektor",
		"penetap": "Dr. Ir. Luthfi Hasan, M.S.",
		"tanggal_ditetapkan": "2020-02-05",
		"tanggal_ditetapkan_idn": "5 Februari 2020",
		"url_sk_kgb": "",
		"uuid_sk_kgb": "c76a98f3-2ffd-11eb-a014-7eb0d4a3c7a0"
}`
