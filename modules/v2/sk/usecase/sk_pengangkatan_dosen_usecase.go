package usecase

import (
	"encoding/json"
	"fmt"
	"net/http"
	"svc-insani-go/app"

	"github.com/labstack/echo"
)

// still dummy
func HandleCreateSkPengangkatanDosen(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		res := []byte(fmt.Sprintf(dummyUpdateSkPengangkatanDosen, "tambah"))

		return c.JSONBlob(http.StatusOK, res)
		// 	ctx := c.Request().Context()
		// 	kdKelompokPegawai := c.QueryParam("kd_kelompok_pegawai")
		// 	if kdKelompokPegawai == "ED" {
		// 		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Tambah SK pengangkatan dosen saat ini belum tersedia"})
		// 	}

		// 	skRequest := new(model.SkPengangkatanDosen)
		// 	err := c.Bind(skRequest)
		// 	if err != nil {
		// 		fmt.Printf("[DEBUG] err binding sk pakt: %s\n", err.Error())
		// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		// 	}

		// 	uuidPegawai := c.QueryParam("uuid_pegawai")
		// 	pegawai, err := kepegawaianRepo.GetPegawai(a, ctx, uuidPegawai)
		// 	if err != nil {
		// 		return c.JSON(
		// 			http.StatusInternalServerError,
		// 			echo.NewHTTPError(
		// 				http.StatusInternalServerError,
		// 				"error get pegawai by uuid: "+err.Error(),
		// 			))
		// 	}
		// 	if pegawai == nil {
		// 		return c.JSON(http.StatusBadRequest, map[string]string{"message": "pegawai tidak ditemukan"})
		// 	}

		// 	if skRequest.JabatanFungsional.Uuid != "" {
		// 		jabfung, err := kepegawaianRepo.GetJabatanFungsional(a, ctx, skRequest.JabatanFungsional.Uuid)
		// 		if err != nil {
		// 			return c.JSON(http.StatusInternalServerError, map[string]string{"message": app.ErrInternalServer})
		// 		}
		// 		if jabfung == nil {
		// 			return c.JSON(http.StatusBadRequest, map[string]string{"message": "jabatan fungsional tidak ditemukan"})
		// 		}
		// 		skRequest.JabatanFungsional = *jabfung
		// 	}

		// 	uuidJabatanPenetap := c.FormValue("uuid_jabatan_penetap")
		// 	if uuidJabatanPenetap != "" {
		// 		jabPenetap, err := organisasiRepo.GetJabatanStruktural(a, ctx, uuidJabatanPenetap)
		// 		if err != nil {
		// 			return c.JSON(http.StatusInternalServerError, map[string]string{"message": app.ErrInternalServer})
		// 		}
		// 		if jabPenetap == nil {
		// 			return c.JSON(http.StatusBadRequest, map[string]string{"message": "jabatan penetap tidak ditemukan"})
		// 		}
		// 		skRequest.JabatanPenetap = *jabPenetap
		// 	}

		// 	if skRequest.JenisIjazah.Uuid != "" {
		// 		jenisIjazah, err := repo.GetJenisIjazah(a, ctx, skRequest.JenisIjazah.Uuid)
		// 		if err != nil {
		// 			return c.JSON(http.StatusInternalServerError, map[string]string{"message": app.ErrInternalServer})
		// 		}
		// 		if jenisIjazah == nil {
		// 			return c.JSON(http.StatusBadRequest, map[string]string{"message": "jenis ijazah tidak ditemukan"})
		// 		}
		// 		skRequest.JenisIjazah = *jenisIjazah
		// 	}

		// 	if skRequest.KelompokSkPengangkatan.Uuid != "" {
		// 		kelompokSk, err := repo.GetKelompokSkPengangkatan(a, ctx, skRequest.KelompokSkPengangkatan.Uuid)
		// 		if err != nil {
		// 			return c.JSON(http.StatusInternalServerError, map[string]string{"message": app.ErrInternalServer})
		// 		}
		// 		if kelompokSk == nil {
		// 			return c.JSON(http.StatusBadRequest, map[string]string{"message": "kelompok SK pengangkatan tidak ditemukan"})
		// 		}
		// 		skRequest.KelompokSkPengangkatan = *kelompokSk
		// 	}

		// 	if skRequest.PangkatGolonganRuang.Uuid != "" {
		// 		pgr, err := kepegawaianRepo.GetPangkatGolonganRuang(a, ctx, skRequest.PangkatGolonganRuang.Uuid)
		// 		if err != nil {
		// 			return c.JSON(http.StatusInternalServerError, map[string]string{"message": app.ErrInternalServer})
		// 		}
		// 		if pgr == nil {
		// 			return c.JSON(http.StatusBadRequest, map[string]string{"message": "pangkat golongan ruang tidak ditemukan"})
		// 		}
		// 		skRequest.PangkatGolonganRuang = *pgr
		// 	}

		// 	uuidPejabatPenetap := c.FormValue("uuid_pejabat_penetap")
		// 	if uuidPejabatPenetap != "" {
		// 		pejabPenetap, err := organisasiRepo.GetPejabatStrukturalByUUID(a, ctx, uuidPejabatPenetap)
		// 		if err != nil {
		// 			return c.JSON(http.StatusInternalServerError, map[string]string{"message": app.ErrInternalServer})
		// 		}
		// 		if pejabPenetap == nil {
		// 			return c.JSON(http.StatusBadRequest, map[string]string{"message": "pejabat penetap tidak ditemukan"})
		// 		}
		// 		skRequest.PejabatPenetap = *pejabPenetap
		// 	}

		// 	if skRequest.StatusPengangkatan.Uuid != "" {
		// 		statusPengangkatan, err := repo.GetStatusPengangkatan(a, ctx, skRequest.StatusPengangkatan.Uuid)
		// 		if err != nil {
		// 			return c.JSON(http.StatusInternalServerError, map[string]string{"message": app.ErrInternalServer})
		// 		}
		// 		if statusPengangkatan == nil {
		// 			return c.JSON(http.StatusBadRequest, map[string]string{"message": "status pengangkatan tidak ditemukan"})
		// 		}
		// 		skRequest.StatusPengangkatan = *statusPengangkatan
		// 	}

		// 	uuidUnitPengangkat := c.FormValue("uuid_unit_pengangkat")
		// 	if uuidUnitPengangkat != "" {
		// 		unitPengangkat, err := organisasiRepo.GetUnit2(a, ctx, uuidUnitPengangkat)
		// 		if err != nil {
		// 			return c.JSON(http.StatusInternalServerError, map[string]string{"message": app.ErrInternalServer})
		// 		}
		// 		if unitPengangkat == nil {
		// 			return c.JSON(http.StatusBadRequest, map[string]string{"message": "unit pengangkat tidak ditemukan"})
		// 		}
		// 		skRequest.UnitPengangkat = *unitPengangkat
		// 	}

		// 	uuidUnitKerja := c.FormValue("uuid_unit_kerja")
		// 	if uuidUnitKerja != "" {
		// 		unitKerja, err := organisasiRepo.GetUnit2(a, ctx, uuidUnitKerja)
		// 		if err != nil {
		// 			return c.JSON(http.StatusInternalServerError, map[string]string{"message": app.ErrInternalServer})
		// 		}
		// 		if unitKerja == nil {
		// 			return c.JSON(http.StatusBadRequest, map[string]string{"message": "unit kerja tidak ditemukan"})
		// 		}
		// 		skRequest.UnitKerja = *unitKerja
		// 	}

		// 	jenisSk, err := repo.GetJenisSk(a, ctx, kdJenisSkPengangkatan)
		// 	if err != nil {
		// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": app.ErrInternalServer})
		// 	}
		// 	if jenisSk == nil {
		// 		fmt.Printf("[ERROR] jenis sk with code %s is not found\n", kdJenisSkPengangkatan)
		// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		// 	}

		// 	skRequest.UserUpdate = c.Request().Header.Get("X-Member")
		// 	skRequest.UserInput = skRequest.UserUpdate
		// 	skRequest.FlagAktif = 1
		// 	skRequest.Uuid = guuid.New().String()

		// 	skRequest.SkPegawai.Pegawai = *pegawai
		// 	skRequest.SkPegawai.Id = uint64(guuid.New().ID())
		// 	skRequest.SkPegawai.IdPegawai = pegawai.Id
		// 	skRequest.SkPegawai.IdJenisSk = jenisSk.Id
		// 	skRequest.SkPegawai.Uuid = guuid.New().String()
		// 	skRequest.SkPegawai.UserUpdate = c.Request().Header.Get("X-Member")
		// 	skRequest.SkPegawai.UserInput = skRequest.SkPegawai.UserUpdate
		// 	skRequest.SkPegawai.FlagAktif = 1

		// 	fileSk, _ := c.FormFile("file_sk")

		// 	// TODO: validasi file sk

		// 	if fileSk != nil {
		// 		f, err := fileSk.Open()
		// 		if err != nil {
		// 			f.Close()
		// 			return c.JSON(
		// 				http.StatusInternalServerError,
		// 				echo.NewHTTPError(
		// 					http.StatusInternalServerError,
		// 					"error open form file sk: "+err.Error(),
		// 				))
		// 		}

		// 		formFile := minio.NewFormFile(&a.MinioClient)
		// 		formFile.Append(a.MinioBucketName, "file_sk", "", fileSk.Header.Get("Content-Type"), fileSk.Size, f)
		// 		f.Close()
		// 		skRequest.PathSk = formFile.GenerateObjectName("file_sk", "sk", "pengangkatan", skRequest.SkPegawai.Pegawai.Uuid)
		// 		err = formFile.Upload()
		// 		if err != nil {
		// 			return c.JSON(
		// 				http.StatusInternalServerError,
		// 				echo.NewHTTPError(
		// 					http.StatusInternalServerError,
		// 					"error upload form file sk pengangkatan tendik: "+err.Error(),
		// 				))
		// 		}
		// 	}

		// 	err = repo.CreateSkPengangkatanDosen(a, ctx, skRequest)
		// 	if err != nil {
		// 		return c.JSON(
		// 			http.StatusInternalServerError,
		// 			echo.NewHTTPError(
		// 				http.StatusInternalServerError,
		// 				"error repo create sk pengangkatan tendik: "+err.Error(),
		// 			))
		// 	}

		// 	return c.JSON(http.StatusOK, map[string]interface{}{
		// 		"message": "Berhasil tambah SK pengangkatan tendik",
		// 		"data":    skRequest.SkPegawai.Pegawai,
		// 	})
	}
	return echo.HandlerFunc(h)
}

// still dummy
func HandleUpdateSkPengangkatanDosen(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		type Req struct {
			Tmt               string   `form:"tmt"`
			UuidMataKuliah    string   `form:"uuid_mata_kuliah"`
			UuidMataKuliahArr []string `form:"-"`
		}

		req := new(Req)
		err := c.Bind(req)
		if err != nil {
			fmt.Printf("[ERROR] binding request: %s\n", err.Error())
		}

		err = json.Unmarshal([]byte(req.UuidMataKuliah), &req.UuidMataKuliahArr)
		if err != nil {
			fmt.Printf("[ERROR] unmarshal uuid mata kuliah: %s\n", err.Error())
		}

		fmt.Printf("[DEBUG] req: %+v\n", req)
		res := []byte(fmt.Sprintf(dummyUpdateSkPengangkatanDosen, "ubah"))

		return c.JSONBlob(http.StatusOK, res)
	}
	return echo.HandlerFunc(h)
}

// still dummy
func HandleGetSkPengangkatanDosen(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		return c.JSONBlob(http.StatusOK, []byte(dummyGetSkPengangkatanDosen))
		// 	ctx := c.Request().Context()
		// 	uuidSkPengangkatanDosen := c.QueryParam("uuid_sk_pengangkatan_tendik")
		// 	sk, err := repo.GetSkPengangkatanDosen(a, ctx, uuidSkPengangkatanDosen)
		// 	if err != nil {
		// 		return c.JSON(
		// 			http.StatusInternalServerError,
		// 			echo.NewHTTPError(
		// 				http.StatusInternalServerError,
		// 				"error get sk pengangkatan tendik: "+err.Error(),
		// 			))
		// 	}

		// 	if sk.PathSk != "" {
		// 		formFile := minio.NewFormFile(&a.MinioClient)
		// 		formFile.Append(a.MinioBucketName, sk.SkPegawai.NomorSk, sk.PathSk, "", 0, nil)
		// 		err = formFile.GenerateUrl()
		// 		if err != nil {
		// 			c.Logger().Debug("error generate url file sk:", err.Error())
		// 		}
		// 		sk.UrlFileSk, sk.NamaFileSk = formFile.GetUrl(sk.SkPegawai.NomorSk)
		// 	}

		// 	res := map[string][]*model.SkPengangkatanDosen{
		// 		"data": []*model.SkPengangkatanDosen{},
		// 	}

		// 	if sk == nil {
		// 		return c.JSON(http.StatusOK, res)
		// 	}

		// 	res["data"] = append(res["data"], sk)

		// 	var buf bytes.Buffer
		// 	enc := json.NewEncoder(&buf)
		// 	enc.SetEscapeHTML(false) // agar url file tidak diescape
		// 	err = enc.Encode(res)
		// 	if err != nil {
		// 		return c.JSON(
		// 			http.StatusInternalServerError,
		// 			echo.NewHTTPError(
		// 				http.StatusInternalServerError,
		// 				"error encoding result: "+err.Error(),
		// 			))
		// 	}

		// 	// return c.JSON(http.StatusOK, res)
		// 	return c.JSONBlob(http.StatusOK, buf.Bytes())
	}
	return echo.HandlerFunc(h)

}

// still dummy
func HandleDeleteSkPengangkatanDosen(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		res := []byte(fmt.Sprintf(dummyUpdateSkPengangkatanDosen, "hapus"))

		return c.JSONBlob(http.StatusOK, res)
		// 	uuid := c.QueryParam("uuid_sk_pengangkatan_tendik")
		// 	if uuid == "" {
		// 		return c.JSON(http.StatusBadRequest, map[string]string{"message": "uuid sk pengangkatan tendik wajib diisi"})
		// 	}

		// 	ctx := c.Request().Context()
		// 	skpt, err := repo.GetSkPengangkatanDosen(a, ctx, uuid)
		// 	if err != nil {
		// 		return c.JSON(
		// 			http.StatusInternalServerError,
		// 			echo.NewHTTPError(
		// 				http.StatusInternalServerError,
		// 				"error get sk pengangkatan tendik by uuid: "+uuid,
		// 			))
		// 	}

		// 	if skpt == nil {
		// 		return c.JSON(http.StatusBadRequest, map[string]string{"message": "sk pengangkatan tendik tidak ditemukan"})
		// 	}

		// 	skpt.SkPegawai.FlagAktif = 0
		// 	skpt.SkPegawai.UserUpdate = c.Request().Header.Get("X-Member")
		// 	skpt.FlagAktif = 0
		// 	skpt.UserUpdate = c.Request().Header.Get("X-Member")

		// 	err = repo.UpdateSkPengangkatanDosen(a, ctx, skpt)
		// 	if err != nil {
		// 		return c.JSON(
		// 			http.StatusInternalServerError,
		// 			echo.NewHTTPError(
		// 				http.StatusInternalServerError,
		// 				"error delete sk pengangkatan tendik: "+err.Error(),
		// 			))
		// 	}
		// 	return c.JSON(http.StatusOK, map[string]interface{}{
		// 		"message": "Berhasil hapus sk pengangkatan tendik",
		// 		"data":    skpt.SkPegawai.Pegawai,
		// 	})
	}
	return echo.HandlerFunc(h)
}

const dummyGetSkPengangkatanDosen = `{
	"data": [
			{
					"nama_pegawai": "Abdul Kadir Aboe",
					"nik_pegawai": "091002120",
					"ttl": "Tegal, 24 November 1979",
					"bantuan_komunikasi": 1000000,
					"gaji_pokok": 1000000,
					"induk_kerja": {
							"kd_unit": "000",
							"unit": "Pengurus Yayasan Badan Wakaf",
							"keterangan": "",
							"uuid": "7f994497-1fd7-11eb-a014-7eb0d4a3c7a0"
					},
					"instansi_kerja": "National University of Singapore",
					"jabatan_fungsional": {
							"kd_jabatan_fungsional": "06",
							"jabatan_fungsional": "Asisten Ahli",
							"uuid": "aeb51169-2fc6-11eb-a014-7eb0d4a3c7a0"
					},
					"jabatan_fungsional_lama": {
							"kd_jabatan_fungsional": "12",
							"jabatan_fungsional": "Lektor",
							"uuid": "aeb51718-2fc6-11eb-a014-7eb0d4a3c7a0"
					},
					"jabatan_penetap": {
							"jenis_jabatan": "Rektor",
							"jenis_unit": "Universitas",
							"unit": "Universitas Islam Indonesia",
							"kd_unit": "100",
							"uuid": "82dbe329-9461-11eb-b06a-000c2977b907"
					},
					"jangka_waktu_evaluasi": "Setiap 2 tahun dari pengangkatan atau sekurang-kurangnya sesuai masa kontrak",
					"jenis_ijazah": {
							"kd_jenis_ijazah": "6",
							"jenis_ijazah": "S2 ",
							"uuid": "74d325b7-ee86-11ea-8c77-7eb0d4a3c7a0"
					},
					"jenis_sk": {
							"kd_jenis_sk": "1",
							"jenis_sk": "Pengangkatan",
							"uuid": "ebc9e2c0-ee60-11ea-8c77-7eb0d4a3c7a0"
					},
					"kelompok_sk_pengangkatan": {
							"kelompok_sk_pengangkatan": "Dosen Tidak Tetap Profesi",
							"uuid_kelompok_sk_pengangkatan": "74201d68-4fea-11eb-bf95-a74048ab8082"
					},
					"mata_kuliah": [
							{
									"kd_matakuliah": "2113001",
									"nama_matakuliah": "Akuntansi Pengantar",
									"nama_matakuliah_en": "Introduction to Accounting",
									"nama_singkat_matakuliah": "Akuntansi Pengantar",
									"uuid": "5833a3de-213b-11ea-889a-506b8da96a87"
							},
							{
									"kd_matakuliah": "STF101",
									"nama_matakuliah": "Pengantar Informatika",
									"nama_matakuliah_en": "Introduction to Informatics",
									"nama_singkat_matakuliah": "PINF",
									"uuid": "4576b120-5b78-11eb-831c-7eb0d4a3c7a0"
							},
							{
									"kd_matakuliah": "41012009",
									"nama_matakuliah": "Hukum Agraria",
									"nama_matakuliah_en": "Agrarian Law",
									"nama_singkat_matakuliah": "Hk. Agraria",
									"uuid": "9542e427-2139-11ea-889a-506b8da96a87"
							}
					],
					"masa_kerja_diakui_tahun_baru": 1,
					"masa_kerja_diakui_bulan_lama": 1,
					"masa_kerja_gaji_bulan": 1,
					"masa_kerja_gaji_tahun": 1,
					"masa_kerja_riil_bulan": 1,
					"masa_kerja_riil_tahun": 1,
					"nomor_sk": "sk-dosen-1",
					"nama_file_sk_pengangkatan": "sk-dosen-1.pdf",
					"pangkat_gol": {
							"pangkat": "Pengatur muda",
							"golongan": "II/A",
							"uuid": "5e40d320-ee83-11ea-8c77-7eb0d4a3c7a0"
					},
					"pangkat_gol_lama": {
							"pangkat": "Juru tingkat 1",
							"golongan": "I/D",
							"uuid": "c6157b5b-09e3-11eb-8c77-7eb0d4a3c7a0"
					},
					"pejabat_penetap": {
							"nama": "Kariyam",
							"gelar_depan": "",
							"gelar_belakang": "S.Si., M.Si.",
							"uuid": "0e67eb57-9463-11eb-b06a-000c2977b907"
					},
					"sks_mengajar": 2400000,
					"tgl_berakhir": "2020-09-08",
					"tgl_berakhir_idn": "8 September 2020",
					"tgl_ditetapkan": "2020-09-08",
					"tgl_ditetapkan_idn": "8 September 2020",
					"tmt": "2020-09-08",
					"tmt_idn": "8 September 2020",
					"tunjangan_beras": 1000000,
					"tunjangan_khusus": 1000000,
					"tunjangan_tahunan": "Akomodasi, asuransi (BPJS Kesehatan), visa dan tiket pesawat tidak berlaku untuk mobilitas daring (dapat berubah tergantung pada situasi pandemi covid 19)",
					"unit_kerja": {
							"kd_unit": "000",
							"unit": "Pengurus Yayasan Badan Wakaf",
							"kd_induk_unit": "000",
							"uuid": "798c791e-1fd3-11eb-a014-7eb0d4a3c7a0"
					},
					"uuid_sk_pengangkatan_dosen": "2252388283",
					"url_sk_pengangkatan": "s3-dev.minio.io/insani/sk-pengangkatan/sk-dosen-1.pdf"
			}
	]
}`

const dummyUpdateSkPengangkatanDosen = `{
	"message": "Berhasil %s SK pengangkatan dosen",
	"data":    
	{
		"nik": "785110101",
		"nama": "Abdul Kadir Aboe",
		"flag_dosen": 1,
		"kd_unit2": 0,
		"jenis_pegawai": {
			"kd_jenis_pegawai": "ED",
			"jenis_pegawai": "Edukatif",
			"uuid": "06e83088-0467-11eb-8c77-7eb0d4a3c7a0"
		},
		"kelompok_pegawai": {
			"kd_kelompok_pegawai": "01",
			"kd_status_pegawai": "",
			"kd_jenis_pegawai": "",
			"kelompok_pegawai": "Dosen Tetap Yayasan",
			"uuid": "bfe1a7d9-f3dc-11ea-8c77-7eb0d4a3c7a0"
		},
		"unit_kerja": {
			"kd_unit_kerja": "511",
			"nama_unit_kerja": "Teknik Sipil S1",
			"uuid": "7996804b-1fd3-11eb-a014-7eb0d4a3c7a0"
		},
		"uuid": "d8c26983-1437-11eb-a014-7eb0d4a3c7a0"
	}
}`
