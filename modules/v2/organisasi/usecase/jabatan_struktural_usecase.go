package usecase

import (
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v2/organisasi/model"
	"svc-insani-go/modules/v2/organisasi/repo"

	"github.com/labstack/echo/v4"
)

func HandleGetAllJabatanStruktural(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		ctx := c.Request().Context()
		jj, err := repo.GetAllJabatanStruktural(a, ctx)
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				echo.NewHTTPError(
					http.StatusInternalServerError,
					err,
				))
		}

		return c.JSON(http.StatusOK, map[string][]model.JabatanStruktural{"data": jj})
	}
	return echo.HandlerFunc(h)
}

func HandleGetPejabatStruktural(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		ctx := c.Request().Context()
		uuidJabatanStruktural := c.QueryParam("uuid_jabatan_struktural")
		pp, err := repo.GetPejabatStruktural(a, ctx, uuidJabatanStruktural)
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				echo.NewHTTPError(
					http.StatusInternalServerError,
					err,
				))
		}

		return c.JSON(http.StatusOK, map[string][]model.PejabatStruktural{"data": pp})
	}
	return echo.HandlerFunc(h)
}
