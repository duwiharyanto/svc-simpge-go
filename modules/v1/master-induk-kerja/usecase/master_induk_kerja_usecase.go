package usecase

import (
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-induk-kerja/model"
	"svc-insani-go/modules/v1/master-induk-kerja/repo"

	"github.com/labstack/echo"
)

func HandleGetIndukKerja(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		KdIndukKerja := c.QueryParam("kd_induk_kerja")
		KdUnitKerja := c.QueryParam("kd_unit_kerja")

		if KdIndukKerja != "" {
			UnitKerja, err := repo.GetUnitKerja(a, KdIndukKerja)
			if err != nil {
				fmt.Printf("[ERROR] %s\n", err.Error())
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
			}

			return c.JSON(http.StatusOK, model.IndukKerjaResponse{
				Data: UnitKerja,
			})
		}

		if KdUnitKerja != "" {
			BagianKerja, err := repo.GetBagianKerja(a, KdUnitKerja)
			if err != nil {
				fmt.Printf("[ERROR] %s\n", err.Error())
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
			}

			return c.JSON(http.StatusOK, model.IndukKerjaResponse{
				Data: BagianKerja,
			})
		}

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
