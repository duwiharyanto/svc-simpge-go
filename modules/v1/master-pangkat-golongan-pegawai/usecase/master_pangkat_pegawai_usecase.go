package usecase

import (
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-pangkat-golongan-pegawai/model"
	"svc-insani-go/modules/v1/master-pangkat-golongan-pegawai/repo"

	"github.com/labstack/echo/v4"
)

func HandleGetPangkatGolonganPegawai(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		pangkatGolPegawai, err := repo.GetPangkatGolonganPegawai(a)
		if err != nil {
			fmt.Printf("[ERROR] repo get pangkat golongan pegawai, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		return c.JSON(http.StatusOK, model.PangkatGolonganPegawaiResponse{
			Data: pangkatGolPegawai,
		})
	}
	return echo.HandlerFunc(h)
}
