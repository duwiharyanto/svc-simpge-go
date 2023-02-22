package usecase

import (
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-bidang-sub-bidang/model"
	"svc-insani-go/modules/v1/master-bidang-sub-bidang/repo"

	"github.com/labstack/echo/v4"
)

func HandleGetBidang(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		Bidang, err := repo.GetBidang(a, c.Request().Context())

		if err != nil {
			fmt.Printf("[ERROR] %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		return c.JSON(http.StatusOK, model.BidangResponse{
			Data: Bidang,
		})
	}
	return echo.HandlerFunc(h)
}

func HandleGetBidangByUUID(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		uuid := c.QueryParam("uuid")
		pp, err := repo.GetBidangByUUID(a, c.Request().Context(), uuid)
		if err != nil {
			fmt.Printf("[ERROR] repo get bidang by uuid, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		return c.JSON(http.StatusOK, pp)
	}
	return echo.HandlerFunc(h)
}

func HandleGetSubBidang(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		SubBidang, err := repo.GetSubBidang(a, c.Request().Context())

		if err != nil {
			fmt.Printf("[ERROR] %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		return c.JSON(http.StatusOK, model.SubBidangResponse{
			Data: SubBidang,
		})
	}
	return echo.HandlerFunc(h)
}

func HandleGetSubBidangByUUID(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		uuid := c.QueryParam("uuid")
		pp, err := repo.GetSubBidangByUUID(a, c.Request().Context(), uuid)
		if err != nil {
			fmt.Printf("[ERROR] repo get sub bidang by uuid, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		return c.JSON(http.StatusOK, pp)
	}
	return echo.HandlerFunc(h)
}
