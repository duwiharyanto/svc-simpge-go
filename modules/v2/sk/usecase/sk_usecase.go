package usecase

import (
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v2/sk/model"
	"svc-insani-go/modules/v2/sk/repo"

	"github.com/labstack/echo"
)

func HandleGetAllJenisIjazah(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		ctx := c.Request().Context()
		jj, err := repo.GetAllJenisIjazah(a, ctx)
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				echo.NewHTTPError(
					http.StatusInternalServerError,
					err,
				))
		}

		res := map[string][]model.JenisIjazah{
			"data": []model.JenisIjazah{},
		}

		if jj == nil {
			return c.JSON(http.StatusOK, res)
		}

		res["data"] = jj
		return c.JSON(http.StatusOK, res)
	}
	return echo.HandlerFunc(h)
}

func HandleGetAllJenisSk(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		ctx := c.Request().Context()
		jj, err := repo.GetAllJenisSk(a, ctx)
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				echo.NewHTTPError(
					http.StatusInternalServerError,
					err,
				))
		}

		res := map[string][]model.JenisSk{
			"data": []model.JenisSk{},
		}

		if jj == nil {
			return c.JSON(http.StatusOK, res)
		}

		res["data"] = jj
		return c.JSON(http.StatusOK, res)
	}

	return echo.HandlerFunc(h)
}

func HandleGetAllKelompokSkPengangkatan(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		ctx := c.Request().Context()
		jj, err := repo.GetAllKelompokSkPengangkatan(a, ctx)
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				echo.NewHTTPError(
					http.StatusInternalServerError,
					err,
				))
		}

		res := map[string][]model.KelompokSkPengangkatan{
			"data": []model.KelompokSkPengangkatan{},
		}

		if jj == nil {
			return c.JSON(http.StatusOK, res)
		}

		res["data"] = jj
		return c.JSON(http.StatusOK, res)
	}
	return echo.HandlerFunc(h)
}

func HandleGetAllMataKuliah(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		// dummy
		mk := []byte(`{
			"data": [
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
			]
	}`)
		return c.JSONBlob(http.StatusOK, mk)
	}
	return echo.HandlerFunc(h)
}
