package usecase

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/sk-pegawai/model"
	skPengangkatanDosen "svc-insani-go/modules/v1/sk-pengangkatan/model"
	"svc-insani-go/modules/v1/sk-pengangkatan/repo"

	"github.com/labstack/echo"
)

func HandleCreateSKPengangkatanDosen(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		kdKelompokPegawai := c.QueryParam("kd_kelompok_pegawai")
		var skPegawai model.SKPegawai
		var err error
		if kdKelompokPegawai != "ED" {
			skPegawai, err = ValidateCreateSKPengangkatanDosen(a, c)
			if errors.Unwrap(err) != nil {
				fmt.Printf("[ERROR] validate create sk pengangkatan, %s\n", err.Error())
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
			}
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
			}
			err = repo.CreateSKPengangkatanDosen(a, skPegawai)
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

func HandleGetDetailSKPengangkatanDosen(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		UUIDSKPengangkatanDosen := c.QueryParam("uuid_sk_pengangkatan_dosen")
		SKPengangkatanDosen, err := repo.GetDetailSKPengangkatanDosen(a, UUIDSKPengangkatanDosen)
		if err != nil {
			fmt.Printf("[ERROR] %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		return c.JSON(http.StatusOK, skPengangkatanDosen.SKPengangkatanDosenResponse{
			Data: SKPengangkatanDosen,
		})
	}
	return echo.HandlerFunc(h)
}

func HandleDeleteSKPengangkatanDosenByUUID(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		param := c.QueryParam("uuid_sk_pengangkatan_dosen")
		err := repo.DeleteSKPengangkatanDosenByUUID(a, param)
		if err != nil {
			res := echo.Map{"message": "Gagal menghapus data"}
			return c.JSON(http.StatusInternalServerError, res)
		}
		res := echo.Map{"message": "Berhasil menghapus data"}
		return c.JSON(http.StatusOK, res)
	}
	return echo.HandlerFunc(h)
}
