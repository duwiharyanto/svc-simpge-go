package usecase

import (
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-organisasi/model"
	"svc-insani-go/modules/v1/master-organisasi/repo"

	"github.com/labstack/echo"
)

func HandleGetIndukKerja(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {

		IndukKerja, err := repo.GetIndukKerja(a)
		if err != nil {
			fmt.Printf("[ERROR] %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		return c.JSON(http.StatusOK, model.IndukKerjaResponse{
			Data: IndukKerja,
		})
	}
	return echo.HandlerFunc(h)
}

func HandleGetUnitKerja(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		KdIndukKerja := c.QueryParam("kd_induk_kerja")

		if KdIndukKerja == "" {
			fmt.Printf("[ERROR] Data induk kerja kosong ,\n")
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		UnitKerja, err := repo.GetUnitKerja(a, KdIndukKerja)
		if err != nil {
			fmt.Printf("[ERROR] %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		return c.JSON(http.StatusOK, model.IndukKerjaResponse{
			Data: UnitKerja,
		})

	}
	return echo.HandlerFunc(h)
}

func HandleGetBagianKerja(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		KdUnitKerja := c.QueryParam("kd_unit_kerja")

		if KdUnitKerja == "" {
			fmt.Printf("[ERROR] Data unit kerja kosong ,\n")
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		BagianKerja, err := repo.GetBagianKerja(a, KdUnitKerja)
		if err != nil {
			fmt.Printf("[ERROR] %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		return c.JSON(http.StatusOK, model.IndukKerjaResponse{
			Data: BagianKerja,
		})

	}
	return echo.HandlerFunc(h)
}

func HandleIndukKerjaByUUID(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		uuid := c.QueryParam("uuid")
		pp, err := repo.GetBagianKerjaByUUID(a, c.Request().Context(), uuid)
		if err != nil {
			fmt.Printf("[ERROR] repo get induk kerja by uuid, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		return c.JSON(http.StatusOK, pp)
	}
	return echo.HandlerFunc(h)
}

func HandleHomebase(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		pp, err := repo.GetHomebase(a, c.Request().Context())
		if err != nil {
			fmt.Printf("[ERROR] repo get homebase, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		return c.JSON(http.StatusOK, model.HomebaseResponse{
			Data: pp,
		})
	}
	return echo.HandlerFunc(h)
}

func HandleHomebaseByUUID(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		uuid := c.QueryParam("uuid")
		pp, err := repo.GetHomebaseByUUID(a, c.Request().Context(), uuid)
		if err != nil {
			fmt.Printf("[ERROR] repo get homebase, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		return c.JSON(http.StatusOK, pp)
	}
	return echo.HandlerFunc(h)
}
