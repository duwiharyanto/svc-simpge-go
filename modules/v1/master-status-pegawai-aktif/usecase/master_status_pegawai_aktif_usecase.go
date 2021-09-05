package usecase

import (
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-status-pegawai-aktif/model"
	"svc-insani-go/modules/v1/master-status-pegawai-aktif/repo"

	"github.com/labstack/echo/v4"
)

func HandleGetStatusPegawaiAktif(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		FlagStatus := c.QueryParam("flag_status")
		StatusAktif, err := repo.GetStatusPegawaiAktif(a, FlagStatus)
		if err != nil {
			fmt.Printf("[ERROR] %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		return c.JSON(http.StatusOK, model.StatusPegawaiAktifResponse{
			Data: StatusAktif,
		})
	}
	return echo.HandlerFunc(h)
}

func HandleGetStatusPegawaiAktifByUUID(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		uuid := c.QueryParam("uuid")
		pp, err := repo.GetStatusPegawaiAktifByUUID(a, c.Request().Context(), uuid)
		if err != nil {
			fmt.Printf("[ERROR] repo get status pegawai aktif by uuid, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		return c.JSON(http.StatusOK, pp)
	}
	return echo.HandlerFunc(h)
}
