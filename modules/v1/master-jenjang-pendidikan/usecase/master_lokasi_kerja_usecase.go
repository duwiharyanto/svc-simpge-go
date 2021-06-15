package usecase

import (
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-jenjang-pendidikan/model"
	"svc-insani-go/modules/v1/master-jenjang-pendidikan/repo"

	"github.com/labstack/echo"
)

func HandleGetLokasiKerja(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		JenjangPendidikan, err := repo.GetJenjangPendidikan(a)
		if err != nil {
			fmt.Printf("[ERROR] %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		return c.JSON(http.StatusOK, model.JenjangPendidikanResponse{
			Data: JenjangPendidikan,
		})
	}
	return echo.HandlerFunc(h)
}

func HandleJenjangPendidikanByUUID(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		uuid := c.QueryParam("uuid")
		pp, err := repo.GetJenjangPendidikanByUUID(a, c.Request().Context(), uuid)
		if err != nil {
			fmt.Printf("[ERROR] repo get jenjang pendidikan by uuid, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		return c.JSON(http.StatusOK, pp)
	}
	return echo.HandlerFunc(h)
}
