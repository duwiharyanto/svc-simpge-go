package usecase

import (
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-jabatan-fungsional/model"
	"svc-insani-go/modules/v1/master-jabatan-fungsional/repo"

	"github.com/labstack/echo/v4"
)

func HandleGetJabatanFungsional(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		kdJenisPegawai := c.QueryParam("kd_jenis_pegawai")
		jabatanFungsional, err := repo.GetAllJabatanFungsional(a, kdJenisPegawai)
		if err != nil {
			fmt.Printf("[ERROR] %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		return c.JSON(http.StatusOK, model.JabatanFungsionalResponse{
			Data: jabatanFungsional,
		})
	}
	return echo.HandlerFunc(h)
}

func HandleJabatanFungsionalByUUID(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		uuid := c.QueryParam("uuid")
		pp, err := repo.GetJabatanFungsionalByUUID(a, c.Request().Context(), uuid)
		if err != nil {
			fmt.Printf("[ERROR] repo get jabatan fungsional by uuid, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		return c.JSON(http.StatusOK, pp)
	}
	return echo.HandlerFunc(h)
}
