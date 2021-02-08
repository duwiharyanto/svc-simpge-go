package app

import (
	"database/sql"
	"net/http"
	"os"
	"strconv"
	"svc-insani-go/app/minio"
	"time"
)

type App struct {
	DB              *sql.DB
	HttpClient      *http.Client
	MinioBucketName string
	MinioClient     minio.MinioClient
	Name            string
	TimeLocation    *time.Location
}

func GetFixedTimeZone() *time.Location {
	timeOffset := os.Getenv("TIME_OFFSET")
	if timeOffset == "" {
		timeOffset = "7"
	}
	nTimeOffset, _ := strconv.Atoi(timeOffset)
	return time.FixedZone("", nTimeOffset*60*60)
}
