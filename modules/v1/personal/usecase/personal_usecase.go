package usecase

import (
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/personal/model"
	"svc-insani-go/modules/v1/personal/repo"

	"github.com/labstack/echo"
)

func HandleSearchPersonal(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		nama := c.QueryParam("nama")
		nikKTP := c.QueryParam("nik_ktp")

		// if nama == "" && nikKTP == "" {
		// 	pp, err := repo.AllPersonal(a, c.Request().Context())
		// 	if err != nil {
		// 		fmt.Printf("[ERROR] repo all personal, %s\n", err.Error())
		// 		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		// 	}
		// 	return c.JSON(http.StatusOK, model.PersonalDataPribadiResponse{
		// 		Data: pp,
		// 	})
		// }

		pp, err := repo.SearchPersonal(a, c.Request().Context(), nama, nikKTP)
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

func HandleGetPersonalByID(a app.App) echo.HandlerFunc {
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
