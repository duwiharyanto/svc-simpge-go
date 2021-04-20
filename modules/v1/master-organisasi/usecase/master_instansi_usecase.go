package usecase

import (
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-organisasi/repo"

	"github.com/labstack/echo"
)

func HandleSearchInstansi(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		nama := c.QueryParam("nama_instansi")
		pp, err := repo.SearchInstansi(a, c.Request().Context(), nama)
		if err != nil {
			fmt.Printf("[ERROR] repo search instansi, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		return c.JSON(http.StatusOK, pp)
	}
	return echo.HandlerFunc(h)
}
