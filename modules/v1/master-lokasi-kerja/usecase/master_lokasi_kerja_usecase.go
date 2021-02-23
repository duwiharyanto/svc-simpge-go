package usecase

import (
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/pegawai/model"
	"svc-insani-go/modules/v1/pegawai/repo"

	"github.com/labstack/echo"
)

func HandleGetLokasiKerjaAll(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		req := &model.PegawaiRequest{}
		err := c.Bind(req)
		if err != nil {
			fmt.Printf("[WARNING] binding pegawai request: %s\n", err.Error())
		}
		res := model.PegawaiResponse{
			Data: []model.Pegawai{},
		}
		count, err := repo.CountPegawai(a, req)
		if err != nil {
			fmt.Printf("[ERROR] repo count pegawai, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		if count == 0 {
			return c.JSON(http.StatusOK, res)
		}
		pp, err := repo.GetAllPegawai(a, req)
		if err != nil {
			fmt.Printf("[ERROR] repo get all pegawai, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		res.Count = count
		res.Data = pp
		res.Limit = req.Limit
		res.Offset = req.Offset
		return c.JSON(http.StatusOK, res)
	}
	return echo.HandlerFunc(h)
}

func HandleGetLokasiKerjaDummy(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		return c.JSONBlob(http.StatusOK, []byte(dummyLokasiKerja))
	}
	return echo.HandlerFunc(h)
}

const dummyLokasiKerja = `{
    "data": [
        {
            "lokasi_kerja": "330",
            "lokasi_kerja_desc": "Fakultas Ilmu Sosial Dan Budaya",
            "uuid": "a818aed6-4fff-11eb-bf95-a74048ab8082"
        },
		{
            "lokasi_kerja": "410",
            "lokasi_kerja_desc": "Fakultas Hukum",
            "uuid": "a818b6f6-4fff-11eb-bf95-a74048ab8082"
        },
		{
            "lokasi_kerja": "100",
            "lokasi_kerja_desc": "Rektorat",
            "uuid": "a818bd36-4fff-11eb-bf95-a74048ab8082"
        },
}`
