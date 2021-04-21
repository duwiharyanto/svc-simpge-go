package usecase_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"svc-insani-go/app"
	"svc-insani-go/app/database"

	"svc-insani-go/router"
	"testing"

	"github.com/labstack/echo"
)

func TestHandleGetAllUnit2(t *testing.T) {
	// init server
	e := echo.New()
	e.Use(router.SetResponseTimeout)

	db, err := database.Connect()
	if err != nil {
		t.Skip("failed connect db:", err)
	}
	gormDb, err := database.InitGorm(db, true)
	if err != nil {
		t.Fatal(err)
	}
	a := app.App{DB: db, GormDB: gormDb}
	fmt.Print(a)
	// router.InitRoute(a, e)
	server := httptest.NewServer(e)
	defer server.Close()

	// create request
	baseURL := server.URL + "/public/api/v1/unit2"
	// fmt.Printf("[DEBUG] base url: %s\n", baseURL)
	req, err := http.NewRequest(http.MethodGet, baseURL, nil)

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
		t.Fatal(err)
	}

	if len(result) == 0 {
		t.Fatal("should not be empty")
	}

	// for _, v := range result {
	// 	fmt.Printf("[DEBUG]: %+v\n", v)
	// }
}
