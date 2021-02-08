package usecase

import (
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-unit-kerja/model"
	"svc-insani-go/modules/v1/master-unit-kerja/repo"

	"github.com/labstack/echo"
)

func HandleGetUnitKerja(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		UnitKerja, err := repo.GetAllUnitKerja(a)
		if err != nil {
			fmt.Printf("[ERROR] %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		return c.JSON(http.StatusOK, model.UnitKerjaResponse{
			Data: UnitKerja,
		})
	}
	return echo.HandlerFunc(h)
}

// func HandleGetUnitKerjaByUUID(a app.App) echo.HandlerFunc {
// 	h := func(c echo.Context) error {
// 		uuid := c.Param("uuid")
// 		unitKerja, err := repo.GetUnitKerjaByUUID(a, uuid)
// 		if err != nil {
// 			fmt.Printf("[ERROR] %s\n", err.Error())
// 			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
// 		}
// 		return c.JSON(http.StatusOK, model.UnitKerjaResponse{
// 			Data: unitKerja,
// 		})
// 	}
// 	return echo.HandlerFunc(h)
// }
