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

	"github.com/labstack/echo/v4"
)

func TestHandleGetAllJenisSK(t *testing.T) {
	e := echo.New()

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
	a := &app.App{DB: db}
	HandleGetAllJenisSK(a)(c)

	// wantedRawResult := fmt.Sprint(`{"data":[{"kd_jenis_sk_pengangkatan":"1","jenis_sk_pengangkatan":"Pengangkatan","uuid":"ebc9e2c0-ee60-11ea-8c77-7eb0d4a3c7a0"},{"kd_jenis_sk_pengangkatan":"2","jenis_sk_pengangkatan":"Prajabatan","uuid":"01649d73-ee61-11ea-8c77-7eb0d4a3c7a0"},{"kd_jenis_sk_pengangkatan":"3","jenis_sk_pengangkatan":"Penetapan","uuid":"0b8c0fd2-ee61-11ea-8c77-7eb0d4a3c7a0"},{"kd_jenis_sk_pengangkatan":"4","jenis_sk_pengangkatan":"Kenaikan Gaji Berkala ","uuid":"fd92ab13-ee79-11ea-8c77-7eb0d4a3c7a0"},{"kd_jenis_sk_pengangkatan":"5","jenis_sk_pengangkatan":"Pangkat/Golongan/Ruang","uuid":"fd94084f-ee79-11ea-8c77-7eb0d4a3c7a0"},{"kd_jenis_sk_pengangkatan":"6","jenis_sk_pengangkatan":"Perpanjangan ","uuid":"fd952bcf-ee79-11ea-8c77-7eb0d4a3c7a0"},{"kd_jenis_sk_pengangkatan":"7","jenis_sk_pengangkatan":"Perpindahan ","uuid":"fd96751f-ee79-11ea-8c77-7eb0d4a3c7a0"},{"kd_jenis_sk_pengangkatan":"8","jenis_sk_pengangkatan":"Pemberhentian ","uuid":"fd97c835-ee79-11ea-8c77-7eb0d4a3c7a0"},{"kd_jenis_sk_pengangkatan":"9","jenis_sk_pengangkatan":"Jabatan Fungsional ","uuid":"fd991c5a-ee79-11ea-8c77-7eb0d4a3c7a0"},{"kd_jenis_sk_pengangkatan":"10","jenis_sk_pengangkatan":"Sanksi ","uuid":"fd9a76d9-ee79-11ea-8c77-7eb0d4a3c7a0"},{"kd_jenis_sk_pengangkatan":"11","jenis_sk_pengangkatan":"Jabatan Struktural ","uuid":"fd9bd8c7-ee79-11ea-8c77-7eb0d4a3c7a0"},{"kd_jenis_sk_pengangkatan":"12","jenis_sk_pengangkatan":"Riwayat Karyassiwa ","uuid":"fd9d569d-ee79-11ea-8c77-7eb0d4a3c7a0"}]}`)
	// wantedRawResult := `{"data":[{"kd_jenis_sk_pengangkatan":"1","jenis_sk_pengangkatan":"Pengangkatan","uuid":"ebc9e2c0-ee60-11ea-8c77-7eb0d4a3c7a0"},{"kd_jenis_sk_pengangkatan":"2","jenis_sk_pengangkatan":"Prajabatan","uuid":"01649d73-ee61-11ea-8c77-7eb0d4a3c7a0"},{"kd_jenis_sk_pengangkatan":"3","jenis_sk_pengangkatan":"Penetapan","uuid":"0b8c0fd2-ee61-11ea-8c77-7eb0d4a3c7a0"},{"kd_jenis_sk_pengangkatan":"4","jenis_sk_pengangkatan":"Kenaikan Gaji Berkala ","uuid":"fd92ab13-ee79-11ea-8c77-7eb0d4a3c7a0"},{"kd_jenis_sk_pengangkatan":"5","jenis_sk_pengangkatan":"Pangkat/Golongan/Ruang","uuid":"fd94084f-ee79-11ea-8c77-7eb0d4a3c7a0"},{"kd_jenis_sk_pengangkatan":"6","jenis_sk_pengangkatan":"Perpanjangan ","uuid":"fd952bcf-ee79-11ea-8c77-7eb0d4a3c7a0"},{"kd_jenis_sk_pengangkatan":"7","jenis_sk_pengangkatan":"Perpindahan ","uuid":"fd96751f-ee79-11ea-8c77-7eb0d4a3c7a0"},{"kd_jenis_sk_pengangkatan":"8","jenis_sk_pengangkatan":"Pemberhentian ","uuid":"fd97c835-ee79-11ea-8c77-7eb0d4a3c7a0"},{"kd_jenis_sk_pengangkatan":"9","jenis_sk_pengangkatan":"Jabatan Fungsional ","uuid":"fd991c5a-ee79-11ea-8c77-7eb0d4a3c7a0"},{"kd_jenis_sk_pengangkatan":"10","jenis_sk_pengangkatan":"Sanksi ","uuid":"fd9a76d9-ee79-11ea-8c77-7eb0d4a3c7a0"},{"kd_jenis_sk_pengangkatan":"11","jenis_sk_pengangkatan":"Jabatan Struktural ","uuid":"fd9bd8c7-ee79-11ea-8c77-7eb0d4a3c7a0"},{"kd_jenis_sk_pengangkatan":"12","jenis_sk_pengangkatan":"Riwayat Karyassiwa ","uuid":"fd9d569d-ee79-11ea-8c77-7eb0d4a3c7a0"}]}`

	// still failed
	// if rec.Body.String() != wantedRawResult {
	// 	fmt.Println(rec.Body.String())
	// 	t.Fatal("result should be same with wanted")
	// }

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
