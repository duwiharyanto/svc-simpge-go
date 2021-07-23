package integrationtest

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"svc-insani-go/app"
	"svc-insani-go/app/database"
	"svc-insani-go/app/minio"
	"svc-insani-go/router"

	"github.com/labstack/echo/v4"
)

const (
	v1path = "public/api/v1/insani"
)

type TestServer struct {
	Server *httptest.Server
	Client *TestClient
}

func UpServer() (*TestServer, error) {
	db, err := database.Connect()
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Can't connect to db: %s", err.Error())
	}

	gormLog := true
	gormDb, err := database.InitGorm(db, gormLog)
	if err != nil {
		return nil, fmt.Errorf("Can't init db conn with gorm: %s", err.Error())
	}

	minioClient, err := minio.Connect()
	if err != nil {
		return nil, fmt.Errorf("Can't connect to minio: %s", err.Error())
	}

	timeLocation := app.GetFixedTimeZone()
	a := &app.App{
		DB:              db,
		GormDB:          gormDb,
		HttpClient:      &http.Client{},
		MinioBucketName: "insani",
		MinioClient:     minioClient,
		Name:            "Insani Service",
		TimeLocation:    timeLocation,
	}
	if a.MinioBucketName == "" {
		a.MinioBucketName = "insani"
	}

	e := echo.New()
	appCtx := context.Background()
	e.Use(router.SetResponseTimeout(appCtx))
	errChan := make(chan error)
	router.InitRoute(a, appCtx, e, errChan)

	srv := httptest.NewServer(e)
	return &TestServer{
		Server: srv,
		Client: &TestClient{Client: a.HttpClient},
	}, nil

}

type TestClient struct {
	Client *http.Client
}

func (c *TestClient) SendRequest(method, targetUrl string, body io.Reader, header http.Header) ([]byte, error) {
	req, err := http.NewRequest(method, targetUrl, body)
	if err != nil {
		return nil, err
	}
	req.Header = header.Clone()
	hres, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	rawResBodyJSON, err := ioutil.ReadAll(hres.Body)
	hres.Body.Close()
	if err != nil {
		return nil, err
	}
	return rawResBodyJSON, nil
}

func FillFormDataFieldMap(w *multipart.Writer, m map[string]string) error {
	for k, v := range m {
		formField, err := w.CreateFormField(k)
		if err != nil {
			return fmt.Errorf("failed create field %s: %w", k, err)
		}
		_, err = io.Copy(formField, strings.NewReader(v))
		if err != nil {
			return fmt.Errorf("failed copy %s value: %w", k, err)
		}
	}

	return nil
}
