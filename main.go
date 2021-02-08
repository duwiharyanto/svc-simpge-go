package main

import (
	"context"
	"errors"
	"fmt"
	lg "log"
	"net/http"
	"os"
	"strings"
	"svc-insani-go/app"
	"svc-insani-go/app/database"
	"svc-insani-go/app/minio"
	"svc-insani-go/router"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	// Membuat koneksi ke database, koneksi hanya dibuat satu kali dan akan digunakan di seluruh proses service
	db, err := database.Connect()
	if err != nil {
		lg.Println("Can't connect to db:", err.Error())
	}

	minioClient, err := minio.Connect()
	if err != nil {
		lg.Println("Can't connect to minio:", err.Error())
	}

	timeLocation := app.GetFixedTimeZone()
	// Variabel a akan digunakan sepanjang proses service
	// berisi koneksi database
	// dan data lain yang memungkinkan untuk digunakan secara berulang
	a := app.App{
		DB:              db,
		HttpClient:      &http.Client{},
		MinioBucketName: os.Getenv("MINIO_BUCKET_PERSONAL"),
		MinioClient:     minioClient,
		Name:            "Personal Service",
		TimeLocation:    timeLocation,
	}
	if a.MinioBucketName == "" {
		a.MinioBucketName = "personal"
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	e.Use(router.SetResponseTimeout)

	// Memanggil fungsi yang mengelola routing
	router.InitRoute(a, e)

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if err == nil {
			return
		}
		if strings.Contains(err.Error(), context.DeadlineExceeded.Error()) {
			c.JSON(http.StatusGatewayTimeout, map[string]string{"message": "Server kehabisan waktu"})
			return
		}
		if strings.Contains(err.Error(), context.Canceled.Error()) {
			fmt.Printf("[INFO] user canceled request\n")
			return
		}
		var he *echo.HTTPError
		if errors.As(err, &he) {
			if he.Code == http.StatusInternalServerError {
				c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
			}
		}
	}

	// e.Use(router.PostRequest(a))
	// e.Use(router.HandleError)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"status":${status},"method":"${method}","uri":"${uri}","time":"${time_rfc3339_nano}","x-member":"${header:X-Member}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
			`"error":"${error}","latency":${latency},` +
			`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
			`"bytes_out":${bytes_out}}` + "\n",
	}))

	e.Use(middleware.Recover())
	if e.Debug {
		e.Logger.SetLevel(log.DEBUG)
	}

	// Menjalankan service di port 80
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	e.Logger.Fatal(e.Start(":" + port))

}
