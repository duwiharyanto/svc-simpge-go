package usecase

import (
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-detail-profesi/model"
	"svc-insani-go/modules/v1/master-detail-profesi/repo"

	"github.com/labstack/echo"
)

func HandleGetDetailProfesi(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		JenjangPendidikan, err := repo.GetDetailProfesi(a, c.Request().Context())
		if err != nil {
			fmt.Printf("[ERROR] %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		return c.JSON(http.StatusOK, model.DetailProfesiResponse{
			Data: JenjangPendidikan,
		})
	}
	return echo.HandlerFunc(h)
}

func HandleDetailProfesiByUUID(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		uuid := c.QueryParam("uuid")
		pp, err := repo.GetDetailProfesiByUUID(a, c.Request().Context(), uuid)
		if err != nil {
			fmt.Printf("[ERROR] repo get jenjang pendidikan by uuid, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		return c.JSON(http.StatusOK, pp)
	}
	return echo.HandlerFunc(h)
}
