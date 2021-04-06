package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/app/minio"
	kepegawaianRepo "svc-insani-go/modules/v2/kepegawaian/repo"
	organisasiRepo "svc-insani-go/modules/v2/organisasi/repo"
	"svc-insani-go/modules/v2/sk/model"
	"svc-insani-go/modules/v2/sk/repo"

	"github.com/labstack/echo"
)

func HandleUpdateSkPengangkatanTendik(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		ctx := c.Request().Context()
		uuid := c.QueryParam("uuid_sk_pengangkatan_tendik")
		sk, err := repo.GetSkPengangkatanTendik(a, ctx, uuid)
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				echo.NewHTTPError(
					http.StatusInternalServerError,
					"error repo get sk pengangkatan tendik",
				))
		}

		if sk == nil {
			return c.JSON(
				http.StatusBadRequest,
				echo.NewHTTPError(
					http.StatusBadRequest,
					"Data SK pengangkatan tendik tidak ditemukan",
				))
		}

		skRequest := *sk
		err = c.Bind(&skRequest)
		if err != nil {
			fmt.Printf("[DEBUG] err binding sk pakt: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		// fmt.Printf("\n[DEBUG] old sk: %+v\n", sk)

		if skRequest.JabatanFungsional.Uuid != sk.JabatanFungsional.Uuid {
			jabfung, err := kepegawaianRepo.GetJabatanFungsional(a, ctx, skRequest.JabatanFungsional.Uuid)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": app.ErrInternalServer})
			}
			if jabfung == nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"message": "jabatan fungsional tidak ditemukan"})
			}
			skRequest.JabatanFungsional = *jabfung
		}

		uuidJabatanPenetap := c.FormValue("uuid_jabatan_penetap")
		if uuidJabatanPenetap != sk.JabatanPenetap.Uuid {
			jabPenetap, err := organisasiRepo.GetJabatanStruktural(a, ctx, uuidJabatanPenetap)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": app.ErrInternalServer})
			}
			if jabPenetap == nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"message": "jabatan penetap tidak ditemukan"})
			}
			skRequest.JabatanPenetap = *jabPenetap
		}

		if skRequest.JenisIjazah.Uuid != sk.JenisIjazah.Uuid {
			jenisIjazah, err := repo.GetJenisIjazah(a, ctx, skRequest.JenisIjazah.Uuid)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": app.ErrInternalServer})
			}
			if jenisIjazah == nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"message": "jenis ijazah tidak ditemukan"})
			}
			skRequest.JenisIjazah = *jenisIjazah
		}

		if skRequest.KelompokSkPengangkatan.Uuid != sk.KelompokSkPengangkatan.Uuid {
			kelompokSk, err := repo.GetKelompokSkPengangkatan(a, ctx, skRequest.KelompokSkPengangkatan.Uuid)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": app.ErrInternalServer})
			}
			if kelompokSk == nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"message": "kelompok SK pengangkatan tidak ditemukan"})
			}
			skRequest.KelompokSkPengangkatan = *kelompokSk
		}

		if skRequest.PangkatGolonganRuang.Uuid != sk.PangkatGolonganRuang.Uuid {
			pgr, err := kepegawaianRepo.GetPangkatGolonganRuang(a, ctx, skRequest.PangkatGolonganRuang.Uuid)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": app.ErrInternalServer})
			}
			if pgr == nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"message": "pangkat golongan ruang tidak ditemukan"})
			}
			skRequest.PangkatGolonganRuang = *pgr
		}

		uuidPejabatPenetap := c.FormValue("uuid_pejabat_penetap")
		if uuidPejabatPenetap != sk.PejabatPenetap.Uuid {
			pejabPenetap, err := organisasiRepo.GetPejabatStrukturalByUUID(a, ctx, uuidPejabatPenetap)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": app.ErrInternalServer})
			}
			if pejabPenetap == nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"message": "pejabat penetap tidak ditemukan"})
			}
			skRequest.PejabatPenetap = *pejabPenetap
		}

		if skRequest.StatusPengangkatan.Uuid != sk.StatusPengangkatan.Uuid {
			statusPengangkatan, err := repo.GetStatusPengangkatan(a, ctx, skRequest.StatusPengangkatan.Uuid)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": app.ErrInternalServer})
			}
			if statusPengangkatan == nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"message": "status pengangkatan tidak ditemukan"})
			}
			skRequest.StatusPengangkatan = *statusPengangkatan
		}

		uuidUnitPengangkat := c.FormValue("uuid_unit_pengangkat")
		if uuidUnitPengangkat != sk.UnitPengangkat.Uuid {
			unitPengangkat, err := organisasiRepo.GetUnit2(a, ctx, uuidUnitPengangkat)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": app.ErrInternalServer})
			}
			if unitPengangkat == nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"message": "unit pengangkat tidak ditemukan"})
			}
			skRequest.UnitPengangkat = *unitPengangkat
		}

		uuidUnitKerja := c.FormValue("uuid_unit_kerja")
		if uuidUnitKerja != sk.UnitKerja.Uuid {
			unitKerja, err := organisasiRepo.GetUnit2(a, ctx, uuidUnitKerja)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": app.ErrInternalServer})
			}
			if unitKerja == nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"message": "unit kerja tidak ditemukan"})
			}
			skRequest.UnitKerja = *unitKerja
		}

		skRequest.UserUpdate = c.Request().Header.Get("X-Member")
		skRequest.SkPegawai.UserUpdate = c.Request().Header.Get("X-Member")

		fileSk, _ := c.FormFile("file_sk")

		// TODO: validasi file sk

		if fileSk != nil {
			f, err := fileSk.Open()
			if err != nil {
				f.Close()
				return c.JSON(
					http.StatusInternalServerError,
					echo.NewHTTPError(
						http.StatusInternalServerError,
						"error open form file sk: "+err.Error(),
					))
			}

			formFile := minio.NewFormFile(&a.MinioClient)
			formFile.Append(a.MinioBucketName, "file_sk", "", fileSk.Header.Get("Content-Type"), fileSk.Size, f)
			f.Close()
			skRequest.PathSk = formFile.GenerateObjectName("file_sk", "sk", "pengangkatan", skRequest.SkPegawai.Pegawai.Uuid)
			err = formFile.Upload()
			if err != nil {
				return c.JSON(
					http.StatusInternalServerError,
					echo.NewHTTPError(
						http.StatusInternalServerError,
						"error upload form file sk pengangkatan tendik: "+err.Error(),
					))
			}
		}

		err = repo.UpdateSkPengangkatanTendik(a, ctx, &skRequest)
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				echo.NewHTTPError(
					http.StatusInternalServerError,
					err,
				))
		}

		// fmt.Printf("\n[DEBUG] sk after req: %+v\n", skRequest)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Berhasil ubah SK pengangkatan tendik",
			"data":    sk.SkPegawai.Pegawai,
		})
	}
	return echo.HandlerFunc(h)
}

func HandleGetSkPengangkatanTendik(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		ctx := c.Request().Context()
		uuidSkPengangkatanTendik := c.QueryParam("uuid_sk_pengangkatan_tendik")
		sk, err := repo.GetSkPengangkatanTendik(a, ctx, uuidSkPengangkatanTendik)
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				echo.NewHTTPError(
					http.StatusInternalServerError,
					"error get sk pengangkatan tendik: "+err.Error(),
				))
		}

		if sk.PathSk != "" {
			formFile := minio.NewFormFile(&a.MinioClient)
			formFile.Append(a.MinioBucketName, "file_sk", sk.PathSk, "", 0, nil)
			err = formFile.GenerateUrl()
			if err != nil {
				c.Logger().Debug("error generate url file sk:", err.Error())
			}
			sk.UrlFileSk = fmt.Sprintf(`%s`, formFile.GetUrl("file_sk"))
		}

		res := map[string][]*model.SkPengangkatanTendik{
			"data": []*model.SkPengangkatanTendik{},
		}

		if sk == nil {
			return c.JSON(http.StatusOK, res)
		}

		res["data"] = append(res["data"], sk)

		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		enc.SetEscapeHTML(false) // agar url file tidak diescape
		err = enc.Encode(res)
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				echo.NewHTTPError(
					http.StatusInternalServerError,
					"error encoding result: "+err.Error(),
				))
		}

		// return c.JSON(http.StatusOK, res)
		return c.JSONBlob(http.StatusOK, buf.Bytes())
	}
	return echo.HandlerFunc(h)

}
