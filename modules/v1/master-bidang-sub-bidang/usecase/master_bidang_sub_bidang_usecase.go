package usecase

import (
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-bidang-sub-bidang/model"
	"svc-insani-go/modules/v1/master-bidang-sub-bidang/repo"

	"github.com/labstack/echo/v4"
)

func HandleGetBidangSubBidang(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		BidangSubBidang, err := repo.GetBidangSubBidang(a, c.Request().Context())

		if err != nil {
			fmt.Printf("[ERROR] %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		return c.JSON(http.StatusOK, model.BidangSubBidangResponse{
			Data: BidangSubBidang,
		})
	}
	return echo.HandlerFunc(h)
}

func HandleGetBidangSubBidangByUUID(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		uuid := c.QueryParam("uuid")
		pp, err := repo.GetBidangSubBidangByUUID(a, c.Request().Context(), uuid)
		if err != nil {
			fmt.Printf("[ERROR] repo get bidang sub bidang by uuid, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		return c.JSON(http.StatusOK, pp)
	}
	return echo.HandlerFunc(h)
}
