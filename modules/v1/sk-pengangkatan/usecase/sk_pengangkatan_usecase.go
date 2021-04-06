package usecase

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/sk-pegawai/model"
	"svc-insani-go/modules/v1/sk-pengangkatan/repo"

	"github.com/labstack/echo"
)

func HandleCreateSKPengangkatan(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		kdKelompokPegawai := c.QueryParam("kd_kelompok_pegawai")
		var skPegawai model.SKPegawai
		var err error
		if kdKelompokPegawai != "ED" {
			skPegawai, err = ValidateCreateSKPengangkatanTendik(a, c)
			if errors.Unwrap(err) != nil {
				fmt.Printf("[ERROR] validate create sk pengangkatan tendik, %s\n", err.Error())
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
			}
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
			}
			err = repo.CreateSKPengangkatanTendik(a, skPegawai)
			if err == sql.ErrNoRows {
				return c.JSON(http.StatusBadRequest, map[string]string{"message": "layanan sedang bermasalah"})
			}
			if err != nil {
				fmt.Printf("[ERROR] create sk pengangkatan, %s\n", err.Error())
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
			}
		} else {
			skPegawai, err = ValidateCreateSKPengangkatanDosen(a, c)
			if errors.Unwrap(err) != nil {
				fmt.Printf("[ERROR] validate create sk pengangkatan dosen, %s\n", err.Error())
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
			}
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
			}
			// err = repo.CreateSKPengangkatanDosen(a, skPegawai)
			// if err == sql.ErrNoRows {
			// 	return c.JSON(http.StatusBadRequest, map[string]string{"message": "layanan sedang bermasalah"})
			// }
			// if err != nil {
			// 	fmt.Printf("[ERROR] create sk pengangkatan, %s\n", err.Error())
			// 	return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
			// }
		}
		message := fmt.Sprintf("Berhasil tambah data sk pengangkatan pegawai")
		return c.JSON(http.StatusOK, map[string]string{"message": message})
	}
	return echo.HandlerFunc(h)
}
