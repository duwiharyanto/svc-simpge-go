package usecase

import (
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/sk/model"
	"svc-insani-go/modules/v1/sk/repo"

	"github.com/labstack/echo"
)

func HandleGetAllStatusPengangkat(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		StatusPengangkatan, err := repo.GetAllStatusPengangkatan(a)
		if err != nil {
			fmt.Printf("[ERROR] %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		return c.JSON(http.StatusOK, model.StatusPengangkatanResponse{
			Data: StatusPengangkatan,
		})
	}
	return echo.HandlerFunc(h)
}
