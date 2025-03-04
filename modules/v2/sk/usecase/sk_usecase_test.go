package usecase_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"svc-insani-go/app"
	"svc-insani-go/app/database"
	"svc-insani-go/router"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestHandleGetAllJenisIjazah(t *testing.T) {
	// init server
	e := echo.New()
	e.Use(router.SetResponseTimeout(context.Background()))

	db, err := database.Connect()
	if err != nil {
		t.Skip("failed connect db:", err)
	}
	gormDb, err := database.InitGorm(db, true)
	if err != nil {
		t.Fatal(err)
	}
	a := &app.App{DB: db, GormDB: gormDb}
	fmt.Print(a)
	// router.InitRoute(a, e)
	server := httptest.NewServer(e)
	defer server.Close()

	// create request
	baseURL := server.URL + "/public/api/v1/jenis-ijazah"
	// fmt.Printf("[DEBUG] base url: %s\n", baseURL)
	req, err := http.NewRequest(http.MethodGet, baseURL, nil)
	if err != nil {
		t.Fatal(err)
	}

	// send http request
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	// read response body
	rawResBodyJSON, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	// format res body indentation
	// var buf bytes.Buffer
	// json.Indent(&buf, rawResBodyJSON, "", "\t")
	// fmt.Printf("[DEBUG] rec body: %s\n", buf.String())

	var result map[string][]interface{}
	err = json.Unmarshal(rawResBodyJSON, &result)
	if err != nil {
		t.Log(err)
		t.Log(string(rawResBodyJSON))
		t.Fail()
	}

	if len(result) == 0 {
		t.Fatal("should not be empty")
	}

	// for _, v := range result {
	// 	fmt.Printf("[DEBUG]: %+v\n", v)
	// }
}

func TestHandleGetAllJenisSk(t *testing.T) {
	// init server
	e := echo.New()
	e.Use(router.SetResponseTimeout(context.Background()))

	db, err := database.Connect()
	if err != nil {
		t.Skip("failed connect db:", err)
	}
	gormDb, err := database.InitGorm(db, true)
	if err != nil {
		t.Fatal(err)
	}
	a := &app.App{DB: db, GormDB: gormDb}
	fmt.Print(a)
	// router.InitRoute(a, e)
	server := httptest.NewServer(e)
	defer server.Close()

	// create request
	baseURL := server.URL + "/public/api/v1/jenis-sk"
	// fmt.Printf("[DEBUG] base url: %s\n", baseURL)
	req, err := http.NewRequest(http.MethodGet, baseURL, nil)
	if err != nil {
		t.Fatal(err)
	}

	// send http request
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	// read response body
	rawResBodyJSON, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	// format res body indentation
	var buf bytes.Buffer
	json.Indent(&buf, rawResBodyJSON, "", "\t")
	fmt.Printf("[DEBUG] rec body: %s\n", buf.String())

	var result map[string][]interface{}
	err = json.Unmarshal(rawResBodyJSON, &result)
	if err != nil {
		t.Log(err)
		t.Log(string(rawResBodyJSON))
		t.Fail()
	}

	if len(result) == 0 {
		t.Fatal("should not be empty")
	}

}

func TestHandleGetAllKelompokSkPengangkatan(t *testing.T) {
	// init server
	e := echo.New()
	e.Use(router.SetResponseTimeout(context.Background()))

	db, err := database.Connect()
	if err != nil {
		t.Skip("failed connect db:", err)
	}
	gormDb, err := database.InitGorm(db, true)
	if err != nil {
		t.Fatal(err)
	}
	a := &app.App{DB: db, GormDB: gormDb}
	fmt.Print(a)
	// router.InitRoute(a, e)
	server := httptest.NewServer(e)
	defer server.Close()

	// create request
	baseURL := server.URL + "/public/api/v1/kelompok-sk-pengangkatan"
	// fmt.Printf("[DEBUG] base url: %s\n", baseURL)
	req, err := http.NewRequest(http.MethodGet, baseURL, nil)
	if err != nil {
		t.Fatal(err)
	}

	// send http request
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	// read response body
	rawResBodyJSON, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	// format res body indentation
	var buf bytes.Buffer
	json.Indent(&buf, rawResBodyJSON, "", "\t")
	// fmt.Printf("[DEBUG] rec body: %s\n", buf.String())

	var result map[string][]interface{}
	err = json.Unmarshal(rawResBodyJSON, &result)
	if err != nil {
		t.Log(err)
		t.Log(string(rawResBodyJSON))
		t.Fail()
	}

	if len(result) == 0 {
		t.Fatal("should not be empty")
	}

}

func TestHandleGetAllMataKuliah(t *testing.T) {
	// init server
	e := echo.New()
	e.Use(router.SetResponseTimeout(context.Background()))

	db, err := database.Connect()
	if err != nil {
		t.Skip("failed connect db:", err)
	}
	gormDb, err := database.InitGorm(db, true)
	if err != nil {
		t.Fatal(err)
	}
	a := &app.App{DB: db, GormDB: gormDb}
	fmt.Print(a)
	// router.InitRoute(a, e)
	server := httptest.NewServer(e)
	defer server.Close()

	// create request
	baseURL := server.URL + "/public/api/v1/mata-kuliah"
	// fmt.Printf("[DEBUG] base url: %s\n", baseURL)
	req, err := http.NewRequest(http.MethodGet, baseURL, nil)
	if err != nil {
		t.Fatal(err)
	}

	// send http request
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	// read response body
	rawResBodyJSON, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	// format res body indentation
	var buf bytes.Buffer
	json.Indent(&buf, rawResBodyJSON, "", "\t")
	fmt.Printf("[DEBUG] rec body: %s\n", buf.String())

	var result map[string][]interface{}
	err = json.Unmarshal(rawResBodyJSON, &result)
	if err != nil {
		t.Log(err)
		t.Log(string(rawResBodyJSON))
		t.Fail()
	}

	if len(result) == 0 {
		t.Fatal("should not be empty")
	}

}
