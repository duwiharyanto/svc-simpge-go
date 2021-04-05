package usecase

import (
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v2/sk/model"
	"svc-insani-go/modules/v2/sk/repo"

	"github.com/labstack/echo"
)

func HandleGetAllJenisIjazah(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		ctx := c.Request().Context()
		jj, err := repo.GetAllJenisIjazah(a, ctx)
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError,
				echo.NewHTTPError(
					http.StatusInternalServerError,
					err,
				))
		}

		res := map[string][]model.JenisIjazah{
			"data": []model.JenisIjazah{},
		}

		if jj == nil {
			return c.JSON(http.StatusOK, res)
		}

		res["data"] = jj
		return c.JSON(http.StatusOK, res)
	}
	return echo.HandlerFunc(h)
}
