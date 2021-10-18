package usecase

import (
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-pendidikan/model"
	"svc-insani-go/modules/v1/master-pendidikan/repo"

	"github.com/labstack/echo/v4"
)

func HandleGetGelarDepan(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		data, err := repo.GetGelarDepan(a, c.Request().Context())
		if err != nil {
			fmt.Printf("[ERROR] repo get gelar depan: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data": data,
		})
	}
	return echo.HandlerFunc(h)
}

func HandleGetGelarBelakang(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		data, err := repo.GetGelarBelakang(a, c.Request().Context())
		if err != nil {
			fmt.Printf("[ERROR] repo get gelar belakang: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data": data,
		})
	}
	return echo.HandlerFunc(h)
}

func HandleGetJenjangPendidikan(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		JenjangPendidikan, err := repo.GetJenjangPendidikan(a, c.Request().Context())
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

func HandleJenjangPendidikanByUUID(a *app.App) echo.HandlerFunc {
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

func HandleGetJenisPendidikan(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		temp, err := repo.GetJenjangPendidikanDetail(a, c.Request().Context())
		if err != nil {
			fmt.Printf("[ERROR] repo jenjang pendidikan detail: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		data := []model.JenisPendidikan{}
		for _, d := range temp {
			newData := model.JenisPendidikan{
				ID:                  d.ID,
				KdJenjangPendidikan: d.KdJenjangPendidikan,
				KdJenis:             d.KdDetail,
				NamaJenis:           d.NamaDetail,
				Keterangan:          d.Keterangan,
				UUID:                d.UUID,
			}
			data = append(data, newData)
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data": data,
		})
	}
	return echo.HandlerFunc(h)
}
