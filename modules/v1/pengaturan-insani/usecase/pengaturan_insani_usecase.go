package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/pengaturan-insani/model"
	"svc-insani-go/modules/v1/pengaturan-insani/repo"

	"github.com/labstack/echo"
)

const (
	err500Msg = "Layanan sedang bermasalah"
)

// TODO: get all pengaturan
func HandleGetPengaturan(a *app.App, pengaturanCache *model.Pengaturan) echo.HandlerFunc {
	h := func(c echo.Context) error {
		atribut := c.QueryParam("atribut")
		data := make(map[string]interface{})
		res := model.PengaturanResponse{Data: data}
		if atribut == "" {
			return c.JSON(http.StatusOK, res)
		}

		nilai, err := LoadPengaturan(a, c.Request().Context(), pengaturanCache, atribut)
		if err != nil {
			fmt.Printf("[ERROR] load pengaturan: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err500Msg})
		}

		if nilai == "" {
			return c.JSON(http.StatusOK, res)
		}

		var any interface{}
		// fmt.Printf("[DEBUG] nilai hasil load: %s\n", nilai)
		// fmt.Printf("[DEBUG] nilai == \"1\": %t\n", nilai == "1")
		buf := bytes.NewBufferString(nilai)
		err = json.Unmarshal(buf.Bytes(), &any)
		if err != nil {
			fmt.Printf("[ERROR] unmarshaling nilai pengaturan: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err500Msg})
		}

		res.Data[atribut] = any
		return c.JSON(http.StatusOK, res)
	}
	return echo.HandlerFunc(h)
}

func HandleUpdatePengaturan(a *app.App, pengaturanCache *model.Pengaturan) echo.HandlerFunc {
	h := func(c echo.Context) error {
		data := make(map[string]interface{})
		res := model.PengaturanResponse{Data: data}
		req := make(map[string]interface{})
		if len(req) > 1 {
			res.Message = "Hanya dapat menyimpan satu pengaturan dalam sekali proses"
			return c.JSON(http.StatusBadRequest, res)
		}
		err := c.Bind(&req)
		if err != nil {
			fmt.Printf("[DEBUG] error binding update pengaturan request: %s\n", err.Error())
		}

		var atribut, nilai string
		var any interface{}
		for k, v := range req {
			any = v
			atribut = k
		}

		b, err := json.Marshal(any)
		if err != nil {
			fmt.Printf("[ERROR] marshaling pengaturan request body: %s\n", err.Error())
			res.Message = err500Msg
			return c.JSON(http.StatusInternalServerError, res)
		}
		nilai = fmt.Sprintf("%s", b)
		// fmt.Printf("[DEBUG] atribut: %s\n", atribut)
		// fmt.Printf("[DEBUG] nilai: %s\n", nilai)

		existingNilai, err := LoadPengaturan(a, c.Request().Context(), pengaturanCache, atribut)
		if err != nil {
			fmt.Printf("[ERROR] load pengaturan: %s\n", err.Error())
			res.Message = err500Msg
			return c.JSON(http.StatusInternalServerError, res)
		}

		userUpdate := c.Request().Header.Get("X-Member")
		ctx := c.Request().Context()
		failedMsg := "Gagal simpan pengaturan"
		if existingNilai == "" {
			// fmt.Printf("[DEBUG] existingNilai before insert: %s\n", existingNilai)
			err := repo.InsertPengaturan(a, ctx, atribut, nilai, userUpdate)
			if err != nil {
				fmt.Printf("[ERROR] repo insert pengaturan: %s\n", err.Error())
				res.Message = failedMsg
			}
		} else {
			fmt.Printf("[DEBUG] nilai before update: %s\n", nilai)
			err := repo.UpdatePengaturan(a, ctx, atribut, nilai, userUpdate)
			if err != nil {
				fmt.Printf("[ERROR] repo update pengaturan: %s\n", err.Error())
				res.Message = failedMsg
			}
		}
		if res.Message == failedMsg {
			res.Message = err500Msg
			return c.JSON(http.StatusInternalServerError, res)
		}

		res.Data[atribut] = any
		if pengaturanCache != nil {
			pengaturanCache.Set(atribut, nilai)
		}
		return c.JSON(http.StatusOK, res)
	}
	return echo.HandlerFunc(h)
}

func LoadPengaturan(a *app.App, ctx context.Context, pengaturanCache *model.Pengaturan, atribut string) (string, error) {
	if pengaturanCache != nil {
		nilai := pengaturanCache.Get(atribut)
		if nilai != "" {
			return nilai, nil
		}
	}

	nilai, err := repo.GetPengaturan(a, ctx, atribut)
	if err != nil {
		return "", fmt.Errorf("[ERROR] load pengaturan: %w", err)
	}

	if nilai == "" {
		return "", nil
	}

	if pengaturanCache != nil {
		pengaturanCache.Set(atribut, nilai)
	}

	return nilai, nil
}
