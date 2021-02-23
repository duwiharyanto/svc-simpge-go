package usecase

import (
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/pegawai/model"
	"svc-insani-go/modules/v1/pegawai/repo"

	"github.com/labstack/echo"
)

func HandleGetJenisNoRegisAll(a app.App) echo.HandlerFunc {
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

func HandleGetJenisNoRegisDummy(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		return c.JSONBlob(http.StatusOK, []byte(dummyJenisNoRegis))
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
