package router

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

func SetResponseTimeout(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defaultTimeoutMs := 13000
		timeoutMs, err := strconv.Atoi(os.Getenv("RESPONSE_TIMEOUT_MS"))
		if err != nil {
			timeoutMs = defaultTimeoutMs
		}
		ctx, cancel := context.WithTimeout(c.Request().Context(), time.Duration(timeoutMs)*time.Millisecond)
		defer cancel()
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}

func HandleError(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if he, ok := err.(*echo.HTTPError); ok {
			fmt.Printf("he: %+v\n", he)
		}
		return next(c)
		// return echo.NewHTTPError(404, "not found")
	}
}

// func PostRequest(a app.App) echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {

// 			// mengecek apakah yang mengupdate user sendiri atau admin
// 			// http.Request.Method == http.MethodGet
// 			if !(c.Request().Method == http.MethodPost || c.Request().Method == http.MethodPut) {
// 				return next(c)
// 			}
// 			c.Response().After(func() {

// 				isUser := c.FormValue("user")

// 				var KdStatusValidasi string
// 				var KeteranganValidasiInput string

// 				if isUser == "1" {
// 					KdStatusValidasi = "MNG"
// 					KeteranganValidasiInput = "Data sedang dalam proses pemeriksaan"
// 				}

// 				if isUser == "0" {
// 					KdStatusValidasi = "VAL"
// 					KeteranganValidasiInput = "Data Telah Divalidasi"
// 				}

// 				if err != nil {
// 					err = fmt.Errorf("error from repo get status by kd, %w", err)
// 					return
// 				}
// 				if status == nil {
// 					err = fmt.Errorf("Data Status tidak ditemukan")
// 					return
// 				}

// 				currentTime := time.Now()

// 					UUID:               c.Param("uuidPersonal"),
// 					IdStatusValidasi:   status.ID,
// 					KeteranganValidasi: KeteranganValidasiInput,
// 					TglValidasi:        currentTime.Format("2006-01-02"),
// 					UserValidasi:       c.Request().Header.Get("X-Member"),
// 				}

// 				path := c.Path()

// 				// fmt.Println("ini path: ", path)

// 				base := "/public/api/v1"

// 				switch path {

// 				// Validasi Data Pribadi
// 				case base + "/:uuidPersonal":
// 					// dataValidasi.KeteranganValidasi = "Data Personal Berhasil Divalidasi"
// 					if err != nil {
// 						err = c.JSON(http.StatusInternalServerError, map[string]string{
// 							"message": "Layanan sedang bermasalah",
// 							"error":   err.Error(),
// 						})
// 						return
// 					}

// 				// Validasi Data Identitas
// 				case base + "/:uuidPersonal/identitas":
// 					// dataValidasi.KeteranganValidasi = "Data Identitas Berhasil Divalidasi"
// 					if err != nil {
// 						err = c.JSON(http.StatusInternalServerError, map[string]string{
// 							"message": "Layanan sedang bermasalah",
// 							"error":   err.Error(),
// 						})
// 						return
// 					}

// 				// Validasi Data Alamat
// 				case base + "/:uuidPersonal/alamat":
// 					// dataValidasi.KeteranganValidasi = "Data Alamat Berhasil Divalidasi"
// 					if err != nil {
// 						err = c.JSON(http.StatusInternalServerError, map[string]string{
// 							"message": "Layanan sedang bermasalah",
// 							"error":   err.Error(),
// 						})
// 						return
// 					}

// 				// Validasi Data Kontak
// 				case base + "/:uuidPersonal/kontak":
// 					// dataValidasi.KeteranganValidasi = "Data Kontak Berhasil Divalidasi"
// 					if err != nil {
// 						err = c.JSON(http.StatusInternalServerError, map[string]string{
// 							"message": "Layanan sedang bermasalah",
// 							"error":   err.Error(),
// 						})
// 						return
// 					}

// 					if err != nil {
// 						err = c.JSON(http.StatusInternalServerError, map[string]string{
// 							"message": "Layanan sedang bermasalah",
// 							"error":   err.Error(),
// 						})
// 						return
// 					}

// 				// Validasi Data Pendidikan
// 				case base + "/:uuidPersonal/pendidikan/:uuidPendidikan":
// 					dataValidasi.UUID = c.Param("uuidPendidikan")
// 					// dataValidasi.KeteranganValidasi = "Data Pendidikan Berhasil Divalidasi"
// 					if err != nil {
// 						err = c.JSON(http.StatusInternalServerError, map[string]string{
// 							"message": "Layanan sedang bermasalah",
// 							"error":   err.Error(),
// 						})
// 						return
// 					}

// 				// Validasi Data Keluarga
// 				case base + "/:uuidPersonal/keluarga/:uuidKeluarga":
// 					dataValidasi.UUID = c.Param("uuidKeluarga")
// 					uuidParents := c.Param("uuidPersonal")
// 					// dataValidasi.KeteranganValidasi = "Data Keluarga Berhasil Divalidasi"
// 					if err != nil {
// 						err = c.JSON(http.StatusInternalServerError, map[string]string{
// 							"message": "Layanan sedang bermasalah",
// 							"error":   err.Error(),
// 						})
// 						return
// 					}

// 				// Validasi Data Kk Keluarga
// 				case base + "/:uuidPersonal/kk":
// 					dataValidasi.UUID = c.Param("uuidPersonal")
// 					// dataValidasi.KeteranganValidasi = "Data Kk Keluarga Berhasil Divalidasi"
// 					if err != nil {
// 						err = c.JSON(http.StatusInternalServerError, map[string]string{
// 							"message": "Layanan sedang bermasalah",
// 							"error":   err.Error(),
// 						})
// 						return
// 					}
// 				}

// 			})
// 			return next(c)
// 		}
// 	}
// }
