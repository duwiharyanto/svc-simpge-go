package usecase

import (
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-jenis-pegawai-tidak-tetap/model"
	"svc-insani-go/modules/v1/master-jenis-pegawai-tidak-tetap/repo"

	"github.com/labstack/echo/v4"
)

func HandleGetJenisPTT(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		JenisNoRegis, err := repo.GetJenisPTT(a)
		if err != nil {
			fmt.Printf("[ERROR] %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		return c.JSON(http.StatusOK, model.JenisPTTResponse{
			Data: JenisNoRegis,
		})
	}
	return echo.HandlerFunc(h)
}

// func HandleGetJenisNoRegisDummy(a *app.App) echo.HandlerFunc {
// 	h := func(c echo.Context) error {
// 		return c.JSONBlob(http.StatusOK, []byte(dummyJenisNoRegis))
// 	}
// 	return echo.HandlerFunc(h)
// }

func HandleJenisPTTByUUID(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		uuid := c.QueryParam("uuid")
		pp, err := repo.GetJenisPTTByUUID(a, c.Request().Context(), uuid)
		if err != nil {
			fmt.Printf("[ERROR] repo get jenis pegawai tidak tetap by uuid, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		return c.JSON(http.StatusOK, pp)
	}
	return echo.HandlerFunc(h)
}

const dummyJenisNoRegis = `{
    "data": [
        {
            "kd_jenis_regis": "001",
            "jenis_no_regis": "NIDN",
            "uuid": "05577756-e996-11e9-8f20-506b8da96a81"
        },
        {
            "kd_jenis_regis": "002",
            "jenis_no_regis": "NIDK",
            "uuid": "05577756-e996-11e9-8f20-506b8da96a82"
        },
        {
            "kd_jenis_regis": "003",
            "jenis_no_regis": "NUP",
            "uuid": "05577756-e996-11e9-8f20-506b8da96a83"
        },
        {
            "kd_jenis_regis": "004",
            "jenis_no_regis": "NITK",
            "uuid": "05577756-e996-11e9-8f20-506b8da96a84"
        },
}`
