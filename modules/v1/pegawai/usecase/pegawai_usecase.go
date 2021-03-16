package usecase

import (
	"fmt"
	"strconv"

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

	pegawaiPendidikan, err := repo.GetPegawaiPendidikan(a, uuidPegawai)
	if err != nil {
		return model.PegawaiDetail{}, fmt.Errorf("error repo get pendidikan pegawai by uuid, %w", err)
	}

	// pegawaiFilePendidikan, err := repo.GetPegawaiFilePendidikan(a, uuidPegawai)
	// if err != nil {
	// 	return model.PegawaiDetail{}, fmt.Errorf("error repo get file pendidikan pegawai by uuid, %w", err)
	// }

	pegawaiDetail.PegawaiYayasan = kepegawaianYayasan
	pegawaiDetail.UnitKerjaPegawai = unitKerjaPegawai
	pegawaiDetail.PegawaiPNSPTT = pegawaiPNS
	pegawaiDetail.PegawaiPNSPTT = pegawaiPTT
	pegawaiDetail.StatusAktif = statusPegawaiAktif
	pegawaiDetail.PegawaiPribadi = pegawaiPribadi
	pegawaiDetail.JenjangPendidikan = pegawaiPendidikan
	// pegawaiDetail.PegawaiPendidikan.BerkasPendukung = pegawaiFilePendidikan

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
    "pribadi": {
        "nik": "785110101",
        "nama": "Abdul Kadir Aboe",
        "jenis_pegawai": "Administratif",
        "kelompok_pegawai": "Dosen Tidak Tetap Perjanjian Kerja",
        "unit_kerja": "Direktorat Sumber Daya Manusia",
        "uuid": "d8c26983-1437-11eb-a014-7eb0d4a3c7a0"
    },
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
                    "tgl_kelulusan_idn": "2 Februari 2021",
                    "flag_ijazah_tertinggi_diakui" : "0",
                    "flag_ijazah_terakhir" : "0",
                    "uuid_pendidikan" : "uuid-pendidikan",
                    "akreditasi": "A",
                    "uuid_akreditasi": "354007de-4bcf-11ea-971c-7eb0d4a3c7a0",
                    "konsentrasi_bidang_ilmu": "",
                    "flag_perguruan_tinggi": 0,
                    "uuid_jenjang": "bdaa4faf-495a-11ea-971c-7eb0d4a3c7a0",
                    "gelar": "",
                    "nomor_induk": "23434343434",
                    "tahun_masuk": "2019",
                    "tahun_lulus": "2021",
                    "judul_tugas_akhir": "",
                    "flag_institusi_luar_negeri": 0,
                    "nomor_ijazah": "2343434",
                    "tgl_ijazah": "2021-02-02",
                    "tgl_ijazah_idn": "2 Februari 2021",
                    "url_ijazah": "https://s3-dev.uii.ac.id/personal/c444cd35-b5cb-11ea-af8b-000c29d8230c/pendidikan/ijazah/1612844993.pdf?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=lmZPXbUgOtkgHa7yiTO6%2F20210305%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20210305T041238Z&X-Amz-Expires=36000&X-Amz-SignedHeaders=host&response-content-disposition=filename%3D%222021-03-05%20111238%20Aan%20Kurniawan%20ijazah.pdf%22&X-Amz-Signature=db726c3c3779a6d3e86439631667718141615cf72cb02354821be95169347d6f",
                    "nama_file_ijazah": "2021-03-05 111238 Aan Kurniawan ijazah.pdf",
                    "flag_ijazah_terverifikasi": 0,
                    "nilai": 23,
                    "jumlah_pelajaran": 32,
                    "url_sk_penyetaraan": "",
                    "nama_file_sk_penyetaraan": "",
                    "nomor_sk_penyetaraan": "",
                    "tgl_sk_penyetaraan": "",
                    "tgl_sk_penyetaraan_idn": "",
                    "berkas_pendukung": [
                        {
                            "kd_jenis_file_pendidikan": "TRN",
                            "jenis_file_pendidikan": "Transkrip nilai",
                            "uuid_jenis_file_pendidikan": "431223fb-385e-11eb-a014-7eb0d4a3c7a0",
                            "nama_file_pendidikan": "2021-03-05 111238 Aan Kurniawan Transkrip nilai.pdf",
                            "url_file_pendidikan": "https://s3-dev.uii.ac.id/personal/c444cd35-b5cb-11ea-af8b-000c29d8230c/pendidikan/transkrip-nilai/1612844993.pdf?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=lmZPXbUgOtkgHa7yiTO6%2F20210305%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20210305T041238Z&X-Amz-Expires=36000&X-Amz-SignedHeaders=host&response-content-disposition=filename%3D%222021-03-05%20111238%20Aan%20Kurniawan%20Transkrip%20nilai.pdf%22&X-Amz-Signature=454eee1159a6f9e49ba6fed81c8d338e9e79dfaab615a8f2c439db3a65fffd71",
                            "uuid_file_pendidikan": "74aea289-6a8f-11eb-9f3c-7eb0d4a3c7a0"
                        }
                    ]
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
                    "uuid_pendidikan" : "uuid-pendidikan",
                    "uuid_akreditasi": "354007de-4bcf-11ea-971c-7eb0d4a3c7a0",
                    "konsentrasi_bidang_ilmu": "",
                    "flag_perguruan_tinggi": 0,
                    "uuid_jenjang": "bdaa4faf-495a-11ea-971c-7eb0d4a3c7a0",
                    "gelar": "",
                    "nomor_induk": "23434343434",
                    "tahun_masuk": "2019",
                    "tahun_lulus": "2021",
                    "judul_tugas_akhir": "",
                    "flag_institusi_luar_negeri": 0,
                    "nomor_ijazah": "2343434",
                    "tgl_ijazah": "2021-02-02",
                    "tgl_ijazah_idn": "2 Februari 2021",
                    "url_ijazah": "https://s3-dev.uii.ac.id/personal/c444cd35-b5cb-11ea-af8b-000c29d8230c/pendidikan/ijazah/1612844993.pdf?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=lmZPXbUgOtkgHa7yiTO6%2F20210305%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20210305T041238Z&X-Amz-Expires=36000&X-Amz-SignedHeaders=host&response-content-disposition=filename%3D%222021-03-05%20111238%20Aan%20Kurniawan%20ijazah.pdf%22&X-Amz-Signature=db726c3c3779a6d3e86439631667718141615cf72cb02354821be95169347d6f",
                    "nama_file_ijazah": "2021-03-05 111238 Aan Kurniawan ijazah.pdf",
                    "flag_ijazah_terverifikasi": 0,
                    "nilai": 23,
                    "jumlah_pelajaran": 32,
                    "tgl_kelulusan_idn": "2 Februari 2021",
                    "url_sk_penyetaraan": "",
                    "nama_file_sk_penyetaraan": "",
                    "nomor_sk_penyetaraan": "",
                    "tgl_sk_penyetaraan": "",
                    "tgl_sk_penyetaraan_idn": "",
                    "berkas_pendukung": [
                        {
                            "kd_jenis_file_pendidikan": "TRN",
                            "jenis_file_pendidikan": "Transkrip nilai",
                            "uuid_jenis_file_pendidikan": "431223fb-385e-11eb-a014-7eb0d4a3c7a0",
                            "nama_file_pendidikan": "2021-03-05 111238 Aan Kurniawan Transkrip nilai.pdf",
                            "url_file_pendidikan": "https://s3-dev.uii.ac.id/personal/c444cd35-b5cb-11ea-af8b-000c29d8230c/pendidikan/transkrip-nilai/1612844993.pdf?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=lmZPXbUgOtkgHa7yiTO6%2F20210305%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20210305T041238Z&X-Amz-Expires=36000&X-Amz-SignedHeaders=host&response-content-disposition=filename%3D%222021-03-05%20111238%20Aan%20Kurniawan%20Transkrip%20nilai.pdf%22&X-Amz-Signature=454eee1159a6f9e49ba6fed81c8d338e9e79dfaab615a8f2c439db3a65fffd71",
                            "uuid_file_pendidikan": "74aea289-6a8f-11eb-9f3c-7eb0d4a3c7a0"
                        }
                    ]
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
                    "flag_ijazah_tertinggi_diakui" : "0",
                    "flag_ijazah_terakhir" : "1",
                    "uuid_pendidikan" : "uuid-pendidikan",
                    "uuid_akreditasi": "354007de-4bcf-11ea-971c-7eb0d4a3c7a0",
                    "konsentrasi_bidang_ilmu": "",
                    "flag_perguruan_tinggi": 0,
                    "uuid_jenjang": "bdaa4faf-495a-11ea-971c-7eb0d4a3c7a0",
                    "gelar": "",
                    "nomor_induk": "23434343434",
                    "tahun_masuk": "2019",
                    "tahun_lulus": "2021",
                    "judul_tugas_akhir": "",
                    "flag_institusi_luar_negeri": 0,
                    "nomor_ijazah": "2343434",
                    "tgl_ijazah": "2021-02-02",
                    "tgl_ijazah_idn": "2 Februari 2021",
                    "url_ijazah": "https://s3-dev.uii.ac.id/personal/c444cd35-b5cb-11ea-af8b-000c29d8230c/pendidikan/ijazah/1612844993.pdf?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=lmZPXbUgOtkgHa7yiTO6%2F20210305%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20210305T041238Z&X-Amz-Expires=36000&X-Amz-SignedHeaders=host&response-content-disposition=filename%3D%222021-03-05%20111238%20Aan%20Kurniawan%20ijazah.pdf%22&X-Amz-Signature=db726c3c3779a6d3e86439631667718141615cf72cb02354821be95169347d6f",
                    "nama_file_ijazah": "2021-03-05 111238 Aan Kurniawan ijazah.pdf",
                    "flag_ijazah_terverifikasi": 0,
                    "nilai": 23,
                    "jumlah_pelajaran": 32,
                    "tgl_kelulusan_idn": "2 Februari 2021",
                    "url_sk_penyetaraan": "",
                    "nama_file_sk_penyetaraan": "",
                    "nomor_sk_penyetaraan": "",
                    "tgl_sk_penyetaraan": "",
                    "tgl_sk_penyetaraan_idn": "",
                    "berkas_pendukung": [
                        {
                            "kd_jenis_file_pendidikan": "TRN",
                            "jenis_file_pendidikan": "Transkrip nilai",
                            "uuid_jenis_file_pendidikan": "431223fb-385e-11eb-a014-7eb0d4a3c7a0",
                            "nama_file_pendidikan": "2021-03-05 111238 Aan Kurniawan Transkrip nilai.pdf",
                            "url_file_pendidikan": "https://s3-dev.uii.ac.id/personal/c444cd35-b5cb-11ea-af8b-000c29d8230c/pendidikan/transkrip-nilai/1612844993.pdf?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=lmZPXbUgOtkgHa7yiTO6%2F20210305%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20210305T041238Z&X-Amz-Expires=36000&X-Amz-SignedHeaders=host&response-content-disposition=filename%3D%222021-03-05%20111238%20Aan%20Kurniawan%20Transkrip%20nilai.pdf%22&X-Amz-Signature=454eee1159a6f9e49ba6fed81c8d338e9e79dfaab615a8f2c439db3a65fffd71",
                            "uuid_file_pendidikan": "74aea289-6a8f-11eb-9f3c-7eb0d4a3c7a0"
                        }
                    ]
                },
                {
                    "kd_jenjang_pendidikan": "S1",
                    "nama_institusi": "Universitas Gajah Mada",
                    "jurusan": "Magister Teknologi Informasi",
                    "tgl_kelulusan": "2020-01-01",
                    "flag_ijazah_tertinggi_diakui" : "0",
                    "flag_ijazah_terakhir" : "0",
                    "uuid_pendidikan" : "uuid-pendidikan",
                    "uuid_akreditasi": "354007de-4bcf-11ea-971c-7eb0d4a3c7a0",
                    "konsentrasi_bidang_ilmu": "",
                    "flag_perguruan_tinggi": 0,
                    "uuid_jenjang": "bdaa4faf-495a-11ea-971c-7eb0d4a3c7a0",
                    "gelar": "",
                    "nomor_induk": "23434343434",
                    "tahun_masuk": "2019",
                    "tahun_lulus": "2021",
                    "judul_tugas_akhir": "",
                    "flag_institusi_luar_negeri": 0,
                    "nomor_ijazah": "2343434",
                    "tgl_ijazah": "2021-02-02",
                    "tgl_ijazah_idn": "2 Februari 2021",
                    "url_ijazah": "https://s3-dev.uii.ac.id/personal/c444cd35-b5cb-11ea-af8b-000c29d8230c/pendidikan/ijazah/1612844993.pdf?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=lmZPXbUgOtkgHa7yiTO6%2F20210305%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20210305T041238Z&X-Amz-Expires=36000&X-Amz-SignedHeaders=host&response-content-disposition=filename%3D%222021-03-05%20111238%20Aan%20Kurniawan%20ijazah.pdf%22&X-Amz-Signature=db726c3c3779a6d3e86439631667718141615cf72cb02354821be95169347d6f",
                    "nama_file_ijazah": "2021-03-05 111238 Aan Kurniawan ijazah.pdf",
                    "flag_ijazah_terverifikasi": 0,
                    "nilai": 23,
                    "jumlah_pelajaran": 32,
                    "tgl_kelulusan_idn": "2 Februari 2021",
                    "url_sk_penyetaraan": "",
                    "nama_file_sk_penyetaraan": "",
                    "nomor_sk_penyetaraan": "",
                    "tgl_sk_penyetaraan": "",
                    "tgl_sk_penyetaraan_idn": "",
                    "berkas_pendukung": [
                        {
                            "kd_jenis_file_pendidikan": "TRN",
                            "jenis_file_pendidikan": "Transkrip nilai",
                            "uuid_jenis_file_pendidikan": "431223fb-385e-11eb-a014-7eb0d4a3c7a0",
                            "nama_file_pendidikan": "2021-03-05 111238 Aan Kurniawan Transkrip nilai.pdf",
                            "url_file_pendidikan": "https://s3-dev.uii.ac.id/personal/c444cd35-b5cb-11ea-af8b-000c29d8230c/pendidikan/transkrip-nilai/1612844993.pdf?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=lmZPXbUgOtkgHa7yiTO6%2F20210305%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20210305T041238Z&X-Amz-Expires=36000&X-Amz-SignedHeaders=host&response-content-disposition=filename%3D%222021-03-05%20111238%20Aan%20Kurniawan%20Transkrip%20nilai.pdf%22&X-Amz-Signature=454eee1159a6f9e49ba6fed81c8d338e9e79dfaab615a8f2c439db3a65fffd71",
                            "uuid_file_pendidikan": "74aea289-6a8f-11eb-9f3c-7eb0d4a3c7a0"
                        }
                    ]
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

func HandleGetPegawaix(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		limit, err := strconv.Atoi(c.QueryParam("limit"))
		if err != nil {
			fmt.Printf("[ERROR] convert string to int, %s\n", err.Error())
		}

		offset, err := strconv.Atoi(c.QueryParam("offset"))
		if err != nil {
			fmt.Printf("[ERROR] convert string to int, %s\n", err.Error())
		}

		pp, err := repo.GetAllPegawaix(a, c.Request().Context(), limit, offset)
		if err != nil {
			fmt.Printf("[ERROR] repo get all pegawai, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		return c.JSON(http.StatusOK, pp)
	}
	return echo.HandlerFunc(h)
}

func HandleGetPegawaiByUUIDx(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		uuidPersonal := c.Param("uuidPersonal")
		pp, err := repo.GetPegawaiByUUIDx(a, c.Request().Context(), uuidPersonal)
		if err != nil {
			fmt.Printf("[ERROR] repo get pegawai by uuid, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		return c.JSON(http.StatusOK, pp)
	}
	return echo.HandlerFunc(h)
}

func HandleUpdatePegawaix(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		pegawaiRequest := new(model.PegawaiUpdate)
		err := c.Bind(pegawaiRequest)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, nil)
	}

	return echo.HandlerFunc(h)
}
