package main

import (
	"context"
	"errors"
	"fmt"
	lg "log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"svc-insani-go/app"
	"svc-insani-go/app/database"
	"svc-insani-go/app/minio"
	"svc-insani-go/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

// test redepoloy

func main() {
	// Membuat koneksi ke database, koneksi hanya dibuat satu kali dan akan digunakan di seluruh proses service
	db, err := database.Connect()
	if err != nil {
		lg.Println("Can't connect to db:", err.Error())
	}

	isDebug, err := strconv.ParseBool(os.Getenv("APP_DEBUG"))
	if err != nil {
		lg.Println("Error parse APP_DEBUG:", err.Error())
	}
	gormDB, err := database.InitGorm(db, isDebug)
	if err != nil {
		lg.Println("Can't connect to gorm db:", err.Error())
	}

	minioClient, err := minio.Connect()
	if err != nil {
		lg.Println("Can't connect to minio:", err.Error())
	}

	timeLocation := app.GetFixedTimeZone()
	// Variabel a akan digunakan sepanjang proses service
	// berisi koneksi database
	// dan data lain yang memungkinkan untuk digunakan secara berulang

	a := &app.App{
		DB:              db,
		GormDB:          gormDB,
		HttpClient:      &http.Client{},
		MinioBucketName: os.Getenv("MINIO_BUCKETNAME"),
		MinioClient:     minioClient,
		Name:            os.Getenv("SERVICE_NAME"),
		TimeLocation:    timeLocation,
	}
	if a.MinioBucketName == "" {
		a.MinioBucketName = "insani"
	}

	appCtx := context.Background()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	// e.Use(router.SetResponseTimeout(appCtx))
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Skipper: router.SkipMasterEndpoint,
	// 	Format:  `${status} ${method} ${uri} ${time_rfc3339_nano} ${header:X-Member} ${latency} ${latency_human} ${bytes_in} ${bytes_out} ${error}` + "\n",
	// }))

	// pengaturan := pengaturanModel.InitPengaturan()

	slackErrChan := app.NewSlackLogger(appCtx, a.HttpClient)
	// Memanggil fungsi yang mengelola routing
	e.Use(router.SetResponseTimeout(appCtx))
	router.InitRoute(a, appCtx, e, slackErrChan)

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
		Format: `${status} ${method} ${uri} ${time_rfc3339_nano} ${header:X-Member} ${latency} ${latency_human} ${bytes_in} ${bytes_out} ${error}` + "\n",
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
