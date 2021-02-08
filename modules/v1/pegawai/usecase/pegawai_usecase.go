package usecase

import (
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/pegawai/model"
	"svc-insani-go/modules/v1/pegawai/repo"

	"github.com/labstack/echo"
)

func HandleGetPegawai(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		req := &model.PegawaiRequest{}
		err := c.Bind(req)
		if err != nil {
			fmt.Printf("[WARNING] binding pegawai request: %s\n", err.Error())
		}
		res := model.PegawaiResponse{
			Data: []model.Pegawai{},
		}
		count, err := repo.CountPegawai(a, req)
		if err != nil {
			fmt.Printf("[ERROR] repo count pegawai, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		if count == 0 {
			return c.JSON(http.StatusOK, res)
		}
		pp, err := repo.GetAllPegawai(a, req)
		if err != nil {
			fmt.Printf("[ERROR] repo get all pegawai, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		res.Count = count
		res.Data = pp
		res.Limit = req.Limit
		res.Offset = req.Offset
		return c.JSON(http.StatusOK, res)
	}
	return echo.HandlerFunc(h)
}

func HandleGetSimpegPegawaiByUUID(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		return c.JSONBlob(http.StatusOK, []byte(dummySimpegPegawaiDetail))
	}
	return echo.HandlerFunc(h)
}

func HandleUpdateSimpegPegawaiByUUID(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Perubahan pegawai berhasil disimpan"})
	}
	return echo.HandlerFunc(h)
}

const dummySimpegPegawaiDetail = `{
    "pendidikan": {
        "tingkat_pdd_pertama": "Sekolah Menengah Umum",
        "uuid_tingkat_pdd_pertama": "uuid-sekolah-menengah-umum",
        "jenis_pdd_pertama": "SMU Umum",
        "uuid_jenis_pdd_pertama": "uuid-jenis-smu-umum"
    },
    "kepegawaian": {
        "jenis_pegawai": "Administratif",
        "kd_jenis_pegawai": "ED",
        "uuid_jenis_pegawai": "uuid-administratif",
        "status_pegawai": "Pegawai tetap",
        "kd_status_pegawai": "PT",
        "uuid_status_pegawai": "uuid-pegawai-tetap",
        "kelompok_pegawai": "Administrasi tetap",
        "kd_kelompok_pegawai": "02",
        "uuid_kelompok_pegawai": "uuid-administrasiitetap",
        "pangkat_gol_ruang": "Pengatur, II/C",
        "uuid_pangkat_gol_ruang": "uuid-pengatur-ii-c",
        "tmt_pangkat_gol_ruang": "1989-01-12",
        "tmt_pangkat_gol_ruang_idn": "12 Januari 1989",
        "jabatan": "",
        "uuid_jabatan": "",
        "tmt_jabatan": "1989-01-12",
        "tmt_jabatan_idn": "12 Januari 1989",
        "masa_kerja_gaji_tahun": 3,
        "masa_kerja_gaji_bulan": 0,
        "angka_kredit": 144
    },
    "unit_kerja": {
        "unit_kontrak_fakultas": "Rektoriat",
        "uuid_unit_kontrak_fakultas": "uuid-rektoriat",
        "unit_kerja_jurusan": "Direktorat Organisai dan Sumber Daya Manusia",
        "uuid_unit_kerja_jurusan": "uuid-direktorat-organisai-dan-sumber-daya-manusia",
        "bagian": "Sistem Informasi",
        "uuid_bagian": "uuid-sistem-informasi",
        "lokasi_kerja": "Rektorat",
        "uuid_lokasi_kerja": "uuid-rektorat",
        "nomor_sk_pertama": "714/SK-Rek/DOSDM/IX",
        "tmt_sk_pertama": "1998-01-12",
        "tmt_sk_pertama_idn": "12 Januari 1998"
    },
    "negara_ptt": {
        "flag_pns": 1,
        "nip_pns": "294197242",
        "pangkat_gol_ruang_pns": "Pengatur, II/C",
        "uuid_pangkat_gol_ruang_pns": "uuid-pengatur-ii-c",
        "tmt_pangkat_gol_ruang_pns": "1989-01-12",
        "tmt_pangkat_gol_ruang_pns_idn": "12 Januari 1989",
        "jabatan_pns": "",
        "uuid_jabatan_pns": "",
        "tmt_jabatan_pns": "1989-01-12",
        "tmt_jabatan_pns_idn": "12 Januari 1989",
        "masa_kerja_gaji_tahun": 3,
        "masa_kerja_gaji_bulan": 0,
        "angka_kredit": 144,
        "instansi_asal_ptt": "Pengatur, II/C",
        "uuid_instansi_asal_ptt": "uuid-pengatur-ii-c",
        "keterangan_ptt": ""
    }
}`
