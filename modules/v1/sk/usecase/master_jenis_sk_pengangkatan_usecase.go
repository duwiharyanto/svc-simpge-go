package usecase

import (
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/sk/model"
	"svc-insani-go/modules/v1/sk/repo"

	"github.com/labstack/echo/v4"
)

func HandleGetAllJenisSKPengangkatan(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		kdJenisPegawai := c.QueryParam("kd_jenis_pegawai")
		jenisSKPengangkatan, err := repo.GetAllJenisSKPengangkatan(a, kdJenisPegawai)
		if err != nil {
			fmt.Printf("[ERROR] %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		return c.JSON(http.StatusOK, model.JenisSKPengangkatanResponse{
			Data: jenisSKPengangkatan,
		})
	}
	return echo.HandlerFunc(h)
}
