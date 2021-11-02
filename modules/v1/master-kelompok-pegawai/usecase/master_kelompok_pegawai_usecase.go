package usecase

import (
	"fmt"
	"net/http"
	"svc-insani-go/app"

	// JenisPegawaiModel "svc-insani-go/modules/v1/master-jenis-pegawai/model"
	// JenisPegawaiRepo "svc-insani-go/modules/v1/master-jenis-pegawai/repo"
	"svc-insani-go/modules/v1/master-kelompok-pegawai/model"
	"svc-insani-go/modules/v1/master-kelompok-pegawai/repo"

	"github.com/labstack/echo/v4"
)

func HandleGetKelompokPegawai(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		kdJenisPegawai := c.QueryParam("kd_jenis_pegawai")
		kdStatusPegawai := c.QueryParam("kd_status_pegawai")
		kk, err := repo.GetAllKelompokPegawai(a, kdJenisPegawai, kdStatusPegawai)
		if err != nil {
			fmt.Printf("[ERROR] repo get all kelompok pegawai, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		return c.JSON(http.StatusOK, model.KelompokPegawaiResponse{
			Data: kk,
		})
	}
	return echo.HandlerFunc(h)
}

func HandleKelompokPegawaiByUUID(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		uuid := c.QueryParam("uuid")
		pp, err := repo.GetKelompokPegawaiByUUID(a, c.Request().Context(), uuid)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		return c.JSON(http.StatusOK, pp)
	}
	return echo.HandlerFunc(h)
}
