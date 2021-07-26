package usecase

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/sk-pegawai/model"
	skPengangkatanModel "svc-insani-go/modules/v1/sk-pengangkatan/model"
	"svc-insani-go/modules/v1/sk-pengangkatan/repo"

	"bytes"

	"github.com/labstack/echo/v4"
)

func HandleCreateSKPengangkatanTendik(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		kdKelompokPegawai := c.QueryParam("kd_kelompok_pegawai")
		var skPegawai model.SKPegawai
		var err error
		if kdKelompokPegawai != "ED" {
			skPegawai, err = ValidateCreateSKPengangkatanTendik(a, c)
			if errors.Unwrap(err) != nil {
				fmt.Printf("[ERROR] validate create sk pengangkatan, %s\n", err.Error())
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
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": message,
			"data":    skPegawai.Pegawai,
		})
	}
	return echo.HandlerFunc(h)
}

func HandleGetDetailSKPengangkatanTendik(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		uuidSKPengangkatanTendik := c.QueryParam("uuid_sk_pengangkatan_tendik")
		sk, err := repo.GetDetailSKPengangkatanTendik(a, uuidSKPengangkatanTendik)
		if err != nil {
			fmt.Printf("[ERROR] repo get detail sk pengangkatan tendik, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		res := map[string][]*skPengangkatanModel.SKPengangkatanTendikDetail{
			"data": []*skPengangkatanModel.SKPengangkatanTendikDetail{},
		}

		if sk == nil {
			return c.JSON(http.StatusOK, res)
		}

		res["data"] = append(res["data"], sk)

		// fmt.Printf("[DEBUG] sk url: %s\n", sk.URLSKTendik)

		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		enc.SetEscapeHTML(false)
		err = enc.Encode(res)
		if err != nil {
			fmt.Printf("[ERROR] encode result: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		return c.JSONBlob(http.StatusOK, buf.Bytes())
		// return c.JSON(http.StatusOK, res)
	}
	return echo.HandlerFunc(h)
}

func HandleDeleteSKPengangkatanTendikByUUID(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		params := c.QueryParam("uuid_sk_pengangkatan_tendik")
		err := repo.DeleteSKPengangkatanTendikByUUID(a, params)
		if err != nil {
			res := echo.Map{"message": "Gagal menghapus data"}
			return c.JSON(http.StatusInternalServerError, res)
		}
		res := echo.Map{"message": "Berhasil menghapus data"}
		return c.JSON(http.StatusOK, res)
	}
	return echo.HandlerFunc(h)
}
