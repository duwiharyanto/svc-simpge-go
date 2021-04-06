package usecase

import (
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-jenis-nomor-registrasi/model"
	"svc-insani-go/modules/v1/master-jenis-nomor-registrasi/repo"

	"github.com/labstack/echo"
)

func HandleGetJenisNoRegis(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		JenisNoRegis, err := repo.GetJenisNomorRegistrasi(a)
		if err != nil {
			fmt.Printf("[ERROR] %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		return c.JSON(http.StatusOK, model.JenisNomorRegistrasiResponse{
			Data: JenisNoRegis,
		})
	}
	return echo.HandlerFunc(h)
}

func HandleJenisNoRegisByUUID(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		uuid := c.QueryParam("uuid")
		pp, err := repo.GetJenisNoRegisByUUID(a, c.Request().Context(), uuid)
		if err != nil {
			fmt.Printf("[ERROR] repo get nomor registrasi by uuid, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		return c.JSON(http.StatusOK, pp)
	}
	return echo.HandlerFunc(h)
}
