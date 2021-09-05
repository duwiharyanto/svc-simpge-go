package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"svc-insani-go/app"
	"svc-insani-go/app/database"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestHandleGetIndukKerja(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	db, err := database.Connect()
	if err != nil {
		t.Skip("failed connect db:", err)
	}
	a := &app.App{DB: db}
	HandleGetIndukKerja(a)(c)

	var buf bytes.Buffer
	json.Indent(&buf, rec.Body.Bytes(), "", "\t")
	fmt.Printf("[DEBUG] rec body: %s\n", buf.String())

	var any interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &any)
	if err != nil {
		t.Fatal(err)
	}

	result, ok := any.(map[string]interface{})
	if !ok {
		t.Fatal()
	}

	data, ok := result["data"].([]interface{})
	if !ok {
		t.Fatal()
	}
	if len(data) == 0 {
		t.Fatal("should not be empty")
	}

}
func TestHandleGetUnitKerja(t *testing.T) {
	e := echo.New()

	q := make(url.Values)
	q.Set("kd_induk_kerja", "1")
	qParam := "?" + q.Encode()
	req := httptest.NewRequest(http.MethodGet, "/"+qParam, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	db, err := database.Connect()
	if err != nil {
		t.Skip("failed connect db:", err)
	}

	gormDb, err := database.InitGorm(db, true)
	if err != nil {
		t.Fatal(err)
	}

	a := &app.App{DB: db, GormDB: gormDb}

	HandleGetUnitKerja(a)(c)

	var buf bytes.Buffer
	json.Indent(&buf, rec.Body.Bytes(), "", "\t")
	fmt.Printf("[DEBUG] rec body: %s\n", buf.String())

	var any interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &any)
	if err != nil {
		t.Fatal(err)
	}

	result, ok := any.(map[string]interface{})
	if !ok {
		t.Fatal()
	}

	data, ok := result["data"].([]interface{})
	if !ok {
		t.Fatal()
	}
	if len(data) == 0 {
		t.Fatal("should not be empty")
	}

}
