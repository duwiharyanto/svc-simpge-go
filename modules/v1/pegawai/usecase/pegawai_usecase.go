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
		uuidPegawai := c.Param("uuidPegawai")
		if uuidPegawai == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "parameter uuid pegawai wajib diisi"})
		}

		pegawaiDetail, err := PrepareGetSimpegPegawaiByUUID(a, uuidPegawai)
		if err != nil {
			fmt.Printf("[ERROR] repo get kepegawaian yayasan uuid, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		return c.JSON(http.StatusOK, pegawaiDetail)
	}
	return echo.HandlerFunc(h)
}

func PrepareGetSimpegPegawaiByUUID(a app.App, uuidPegawai string) (model.PegawaiDetail, error) {
	pegawaiDetail := model.PegawaiDetail{}

	pegawaiPribadi, err := repo.GetPegawaiPribadi(a, uuidPegawai)
	if err != nil {
		return model.PegawaiDetail{}, fmt.Errorf("error repo get pribadi pegawai uuid, %w", err)
	}

	kepegawaianYayasan, err := repo.GetKepegawaianYayasan(a, uuidPegawai)
	if err != nil {
		return model.PegawaiDetail{}, fmt.Errorf("error repo get kepegawaian yayasan uuid, %w", err)
	}

	unitKerjaPegawai, err := repo.GetUnitKerjaPegawai(a, uuidPegawai)
	if err != nil {
		return model.PegawaiDetail{}, fmt.Errorf("error repo get unit kerja pegawai by uuid, %w", err)
	}

	pegawaiPNS, err := repo.GetPegawaiPNS(a, uuidPegawai)
	if err != nil {
		return model.PegawaiDetail{}, fmt.Errorf("error repo get pegawai pns by uuid, %w", err)
	}

	pegawaiPTT, err := repo.GetPegawaiPTT(a, uuidPegawai)
	if err != nil {
		return model.PegawaiDetail{}, fmt.Errorf("error repo get pegawai tidak tetap by uuid, %w", err)
	}

	statusPegawaiAktif, err := repo.GetStatusPegawaiAktif(a, uuidPegawai)
	if err != nil {
		return model.PegawaiDetail{}, fmt.Errorf("error repo get status aktif pegawai by uuid, %w", err)
	}

	pegawaiDetail.PegawaiYayasan = kepegawaianYayasan
	pegawaiDetail.UnitKerjaPegawai = unitKerjaPegawai
	pegawaiDetail.PegawaiPNSPTT = pegawaiPNS
	pegawaiDetail.PegawaiPNSPTT = pegawaiPTT
	pegawaiDetail.StatusAktif = statusPegawaiAktif
	pegawaiDetail.PegawaiPribadi = pegawaiPribadi

	return pegawaiDetail, nil
}

func HandleUpdateSimpegPegawaiByUUID(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Perubahan pegawai berhasil disimpan"})
	}
	return echo.HandlerFunc(h)
}

func HandleGetSimpegPegawaiByUUIDDummy(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		return c.JSONBlob(http.StatusOK, []byte(dummySimpegPegawaiDetail))
	}
	return echo.HandlerFunc(h)
}

const dummySimpegPegawaiDetail = `{
    "pendidikan": 
        [
            {
                "jenjang" :"SMU",
                "data":[
                {
                    "kd_jenjang_pendidikan": "SMU",
                    "nama_institusi": "SMA N 1 Sleman",
                    "jurusan": "IPA",
                    "tgl_kelulusan": "2015-01-01",
                    "flag_ijazah_tertinggi_diakui" : "0",
                    "flag_ijazah_terakhir" : "0",
                    "uuid_pendidikan" : "uuid-pendidikan"
                }]
            },
            {
                "jenjang" :"S1",
                "data":[
                {
                    "kd_jenjang_pendidikan": "S1",
                    "nama_institusi": "Universitas Islam Indonesia",
                    "jurusan": "Teknik Informatika",
                    "tgl_kelulusan": "2019-01-01",
                    "flag_ijazah_tertinggi_diakui" : "1",
                    "flag_ijazah_terakhir" : "0",
                    "uuid_pendidikan" : "uuid-pendidikan"
                }]
            },
            {
                "jenjang" :"S2",
                "data":[
                {
                    "kd_jenjang_pendidikan": "S1",
                    "nama_institusi": "Universitas Islam Indonesia",
                    "jurusan": "Magister Teknik Informatika",
                    "tgl_kelulusan": "2019-01-01",
                    "flag_ijazah_tertinggi_diakui" : "1",
                    "flag_ijazah_terakhir" : "0",
                    "uuid_pendidikan" : "uuid-pendidikan"
                },
                {
                    "kd_jenjang_pendidikan": "S1",
                    "nama_institusi": "Universitas Gajah Mada",
                    "jurusan": "Magister Teknologi Informasi",
                    "tgl_kelulusan": "2020-01-01",
                    "flag_ijazah_tertinggi_diakui" : "0",
                    "flag_ijazah_terakhir" : "0",
                    "uuid_pendidikan" : "uuid-pendidikan"
                }]
            }
        ],
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
        "pangkat_gol_ruang_pegawai": "Pengatur, II/C",
        "uuid_pangkat_gol_ruang_pegawai": "uuid-pengatur-ii-c",
        "tmt_pangkat_gol_ruang_pegawai": "1989-01-12",
        "tmt_pangkat_gol_ruang_pegawai_idn": "12 Januari 1989",
        "jabatan_pegawai": "",
        "uuid_jabatan_pegawai": "",
        "tmt_jabatan_pegawai": "1989-01-12",
        "tmt_jabatan_pegawai_idn": "12 Januari 1989",
        "masa_kerja_gaji_tahun_pegawai": 3,
        "masa_kerja_gaji_bulan_pegawai": 0,
        "masa_kerja_total_tahun_pegawai": 3,
        "masa_kerja_total_bulan_pegawai": 0,
        "angka_kredit_pegawai": 144,
        "nomor_sertifikasi_pegawai": "123-no.sertf/2010",
        "jenis_no_registrasi_pegawai": "",
        "uuid_jenis_no_registrasi_pegawai": "",
        "no_registrasi_pegawai": "0123456789"
    },
    "unit_kerja": {
        "unit_kontrak_fakultas": "Rektoriat",
        "uuid_unit_kontrak_fakultas": "uuid-rektoriat",
        "unit_kerja_jurusan": "Direktorat Organisai dan Sumber Daya Manusia",
        "uuid_unit_kerja_jurusan": "uuid-direktorat-organisai-dan-sumber-daya-manusia",
        "bagian_unit_kerja": "Sistem Informasi",
        "uuid_bagian_unit_kerja": "uuid-sistem-informasi",
        "lokasi_unit_kerja": "Rektorat",
        "uuid_lokasi_unit_kerja": "uuid-rektorat",
        "nomor_sk_pertama_unit_kerja": "714/SK-Rek/DOSDM/IX",
        "tmt_sk_pertama_unit_kerja": "1998-01-12",
        "tmt_sk_pertama_unit_kerja_idn": "12 Januari 1998"
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
        "masa_kerja_pns_tahun": 3,
        "masa_kerja_pns_bulan": 0,
        "angka_kredit_pns": 144,
        "instansi_asal_ptt": "Pengatur, II/C",
        "uuid_instansi_asal_ptt": "uuid-pengatur-ii-c",
        "keterangan_ptt": ""
    },
    "status_aktif": {
        "flag_aktif_pegawai": 1,
        "status_aktif_pegawai": "Cuti diluar tanggungan",
        "uuid_status_aktif_pegawai": "uuid-cuti-diluar-tanggungan"
    }
}`
