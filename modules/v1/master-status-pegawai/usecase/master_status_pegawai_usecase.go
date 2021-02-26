package usecase

import (
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-status-pegawai/model"
	"svc-insani-go/modules/v1/master-status-pegawai/repo"

	"github.com/labstack/echo"
)

func HandleGetAllStatusPegawai(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		StatusPegawai, err := repo.GetAllStatusPegawai(a)
		if err != nil {
			fmt.Printf("[ERROR] %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		return c.JSON(http.StatusOK, model.StatusPegawaiResponse{
			Data: StatusPegawai,
		})
	}
	return echo.HandlerFunc(h)
}
