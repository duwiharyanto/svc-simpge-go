package usecase

import (
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/personal/model"
	"svc-insani-go/modules/v1/personal/repo"

	"github.com/labstack/echo/v4"
)

func HandleSearchPersonal(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		cari := c.QueryParam("cari")

		pp, err := repo.SearchPersonal(a, c.Request().Context(), cari)
		if err != nil {
			fmt.Printf("[ERROR] repo search personal, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		return c.JSON(http.StatusOK, model.PersonalDataPribadiResponse{
			Data: pp,
		})
	}
	return echo.HandlerFunc(h)
}

func HandleGetPersonalByID(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		uuidPersonal := c.Param("uuidPersonal")
		fmt.Println("Uuid Personal : ", uuidPersonal)
		pp, err := repo.GetPersonalByUuid(a, c.Request().Context(), uuidPersonal)
		if err != nil {
			fmt.Printf("[ERROR] repo get personal by uuid, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		return c.JSON(http.StatusOK, pp)
	}
	return echo.HandlerFunc(h)
}
