package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"svc-insani-go/app"
	"svc-insani-go/app/database"
	"testing"

	"github.com/labstack/echo"
)

func TestHandleGetAllJenisSKPengangkatan(t *testing.T) {
	e := echo.New()

	// wbuf := &bytes.Buffer{}
	// wr := multipart.NewWriter(wbuf)
	// fwr, err := wr.CreateFormField("nama_lengkap")
	// if err != nil {
	// 	t.Fatal("failed create field nama lengkap:", err)
	// }

	// _, err = io.Copy(fwr, strings.NewReader("harisf"))
	// if err != nil {
	// 	t.Fatal("failed copy nama lengkap value:", err)
	// }

	// wr.Close()

	// req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader((wbuf.Bytes())))
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	// req.Header.Set("Content-Type", wr.FormDataContentType())
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// c.SetParamNames("uuidPersonal")
	// c.SetParamValues("c39b7c23-b5cb-11ea-af8b-000c29d8230c")
	db, err := database.Connect()
	if err != nil {
		t.Skip("failed connect db:", err)
	}
	a := app.App{DB: db}
	HandleGetAllJenisSKPengangkatan(a)(c)

	var buf bytes.Buffer
	json.Indent(&buf, rec.Body.Bytes(), "", "\t")
	fmt.Printf("[DEBUG] rec body: %s\n", buf.String())

	var any interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &any)
	if err != nil {
		t.Fatal(err)
	}
	// fmt.Printf("[DEBUG] raw result: %s\n", rec.Body.String())

	result, ok := any.(map[string]interface{})
	if !ok {
		t.Fatal()
	}
	// fmt.Printf("[DEBUG] any: %#v\n", any)

	data, ok := result["data"].([]interface{})
	if !ok {
		t.Fatal()
	}
	if len(data) == 0 {
		t.Fatal("should not be empty")
	}

	// for _, v := range data {
	// 	fmt.Printf("[DEBUG]: %+v\n", v)
	// }
	// fmt.Printf("[DEBUG] data: %+v\n", result["data"])
}
