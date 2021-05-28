package usecase

import (
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/personal/repo"

	"github.com/labstack/echo"
)

func HandleSearchPersonal(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		nama := c.QueryParam("nama")
		nikPegawai := c.QueryParam("nik_pegawai")
		pp, err := repo.SearchPersonal(a, c.Request().Context(), nama, nikPegawai)
		if err != nil {
			fmt.Printf("[ERROR] repo search personal, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		return c.JSON(http.StatusOK, pp)
	}
	return echo.HandlerFunc(h)
}

func HandleGetPersonalByID(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		idPersonal := c.Param("idPersonal")
		fmt.Println("Id Personal : ", idPersonal)
		pp, err := repo.GetPersonalById(a, c.Request().Context(), idPersonal)
		if err != nil {
			fmt.Printf("[ERROR] repo get personal by uuid, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		return c.JSON(http.StatusOK, pp)
	}
	return echo.HandlerFunc(h)
}
