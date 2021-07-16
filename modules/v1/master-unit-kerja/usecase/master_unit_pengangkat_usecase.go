package usecase

import (
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-unit-kerja/model"
	"svc-insani-go/modules/v1/master-unit-kerja/repo"

	"github.com/labstack/echo/v4"
)

func HandleGetUnitPengangkat(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		unitPengangkat, err := repo.GetAllUnitPengangkat(a)
		if err != nil {
			fmt.Printf("[ERROR] repo get all unit pengangkat, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		return c.JSON(http.StatusOK, model.UnitPengangkatResponse{
			Data: unitPengangkat,
		})
	}
	return echo.HandlerFunc(h)
}
