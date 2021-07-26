package usecase

import (
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/sk/model"
	"svc-insani-go/modules/v1/sk/repo"

	"github.com/labstack/echo/v4"
)

func HandleGetAllSKPegawaiV0(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		req := &model.SKPegawaiRequest{}
		err := c.Bind(req)
		if err != nil {
			fmt.Printf("[WARNING] binding sk pegawai request: %s\n", err.Error())
		}
		if req.KdJenisPegawai == "" || req.UUIDPegawai == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "uuid pegawai dan kd jenis pegawai wajib diisi"})
		}
		ss, err := repo.GetAllSKPegawai(a, req)
		if err != nil {
			fmt.Printf("[ERROR] repo get all sk pegawai: %s\n", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		res := model.SKPegawaiResponse{
			Data: []model.SKPegawai{},
		}
		if len(ss) == 0 {
			return c.JSON(http.StatusOK, res)
		}
		res.Data = ss

		return c.JSON(http.StatusOK, res)
	}
	return echo.HandlerFunc(h)
}

func HandleGetAllSKPegawai(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		req := &model.SKPegawaiRequest{}
		err := c.Bind(req)
		if err != nil {
			fmt.Printf("[WARNING] binding sk pegawai request: %s\n", err.Error())
		}
		if req.KdJenisPegawai == "" || req.UUIDPegawai == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "uuid pegawai dan kd jenis pegawai wajib diisi"})
		}
		ss, err := repo.GetAllSKPegawai(a, req)
		if err != nil {
			fmt.Printf("[ERROR] repo get all sk pegawai: %s\n", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		res := model.SKPegawaiResponse{
			Data: []model.SKPegawai{},
		}
		if len(ss) == 0 {
			return c.JSON(http.StatusOK, res)
		}
		res.Data = ss

		return c.JSON(http.StatusOK, res)
	}
	return echo.HandlerFunc(h)
}
