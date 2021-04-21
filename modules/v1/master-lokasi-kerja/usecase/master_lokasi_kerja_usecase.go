package usecase

import (
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-lokasi-kerja/model"
	"svc-insani-go/modules/v1/master-lokasi-kerja/repo"

	"github.com/labstack/echo"
)

func HandleGetLokasiKerja(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		LokasiKerja, err := repo.GetLokasiKerja(a)
		if err != nil {
			fmt.Printf("[ERROR] %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		return c.JSON(http.StatusOK, model.LokasiKerjaResponse{
			Data: LokasiKerja,
		})
	}
	return echo.HandlerFunc(h)
}

func HandleGetLokasiKerjaDummy(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		return c.JSONBlob(http.StatusOK, []byte(dummyLokasiKerja))
	}
	return echo.HandlerFunc(h)
}

func HandleLokasiKerjaByUUID(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		uuid := c.QueryParam("uuid")
		pp, err := repo.GetLokasiKerjaByUUID(a, c.Request().Context(), uuid)
		if err != nil {
			fmt.Printf("[ERROR] repo get lokasi kerja by uuid, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		return c.JSON(http.StatusOK, pp)
	}
	return echo.HandlerFunc(h)
}

const dummyLokasiKerja = `{
    "data": [
        {
            "lokasi_kerja": "330",
            "lokasi_kerja_desc": "Fakultas Ilmu Sosial Dan Budaya",
            "uuid": "a818aed6-4fff-11eb-bf95-a74048ab8082"
        },
		{
            "lokasi_kerja": "410",
            "lokasi_kerja_desc": "Fakultas Hukum",
            "uuid": "a818b6f6-4fff-11eb-bf95-a74048ab8082"
        },
		{
            "lokasi_kerja": "100",
            "lokasi_kerja_desc": "Rektorat",
            "uuid": "a818bd36-4fff-11eb-bf95-a74048ab8082"
        },
}`
