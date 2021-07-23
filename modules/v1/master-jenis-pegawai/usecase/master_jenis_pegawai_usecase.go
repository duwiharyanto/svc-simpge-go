package usecase

import (
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-jenis-pegawai/model"
	"svc-insani-go/modules/v1/master-jenis-pegawai/repo"

	"github.com/labstack/echo/v4"
)

func HandleGetAllJenisPegawai(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		JenisPegawai, err := repo.GetAllJenisPegawai(a)
		if err != nil {
			fmt.Printf("[ERROR] %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		return c.JSON(http.StatusOK, model.JenisPegawaiResponse{
			Data: JenisPegawai,
		})
	}
	return echo.HandlerFunc(h)
}
