package usecase

import (
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v2/organisasi/model"
	"svc-insani-go/modules/v2/organisasi/repo"

	"github.com/labstack/echo"
)

func HandleGetAllUnit2(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		ctx := c.Request().Context()
		jj, err := repo.GetAllUnit2(a, ctx)
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				echo.NewHTTPError(
					http.StatusInternalServerError,
					err,
				))
		}

		return c.JSON(http.StatusOK, map[string][]model.Unit2{"data": jj})
	}
	return echo.HandlerFunc(h)
}
