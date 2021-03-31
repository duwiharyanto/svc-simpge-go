package usecase

import (
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-jabatan-fungsional/model"
	"svc-insani-go/modules/v1/master-jabatan-fungsional/repo"

	"github.com/labstack/echo"
)

func HandleGetJabatanPenetap(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		jabatanFungsional, err := repo.GetAllJabatanFungsional(a, "")
		if err != nil {
			fmt.Printf("[ERROR] %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		return c.JSON(http.StatusOK, model.JabatanFungsionalResponse{
			Data: jabatanFungsional,
		})
	}
	return echo.HandlerFunc(h)
}
