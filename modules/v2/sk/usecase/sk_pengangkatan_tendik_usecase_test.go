package usecase_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"svc-insani-go/app"
	"svc-insani-go/app/database"
	"svc-insani-go/modules/v2/sk/model"

	"svc-insani-go/router"
	"testing"

	"github.com/labstack/echo"
)

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

func FillFormDataField(w *multipart.Writer, formField io.Writer, key, value string) (io.Writer, error) {
	formField, err := w.CreateFormField(key)
	if err != nil {
		return nil, fmt.Errorf("failed create field %s: %w", key, err)
	}

	_, err = io.Copy(formField, strings.NewReader(value))
	if err != nil {
		return nil, fmt.Errorf("failed copy %s value: %w", key, err)
	}

	return formField, nil
}

func TestHandleUpdateSkPengangkatanTendik(t *testing.T) {
	// init server
	e := echo.New()
	e.Use(router.SetResponseTimeout)

	db, err := database.Connect()
	if err != nil {
		t.Skip("failed connect db:", err)
	}
	gormDb, err := database.InitGorm(db)
	if err != nil {
		t.Fatal(err)
	}
	a := app.App{DB: db, GormDB: gormDb}
	router.InitRoute(a, e)
	server := httptest.NewServer(e)
	defer server.Close()

	// init form data
	wbuf := &bytes.Buffer{}
	wr := multipart.NewWriter(wbuf)
	err = FillFormDataFieldMap(wr, map[string]string{
		"uuid_jenis_sk":      "ebc9e2c0-ee60-11ea-8c77-7eb0d4a3c7a0",
		"gaji_pokok":         "1124",
		"tanggal_ditetapkan": "333",
		// "nomor_sk":                      "no/sk/1/4",
		"nomor_sk":                      "",
		"tentang_sk":                    "tentang sk tendik 1",
		"tmt":                           "2021-09-08",
		"uuid_kelompok_sk_pengangkatan": "742024ac-4fea-11eb-bf95-a74048ab8082",
		"uuid_jabatan_fungsional":       "aeb51718-2fc6-11eb-a014-7eb0d4a3c7a0",
		"uuid_jabatan_penetap":          "6bd8793c-9461-11eb-b06a-000c2977b907",
		"uuid_unit_pengangkat":          "798c80c4-1fd3-11eb-a014-7eb0d4a3c7a0",
		"uuid_unit_kerja":               "798c8162-1fd3-11eb-a014-7eb0d4a3c7a0",
		"masa_kerja_diakui_bulan":       "7",
		"masa_kerja_diakui_tahun":       "1",
		"masa_kerja_ril_bulan":          "11",
		"masa_kerja_ril_tahun":          "2",
		"masa_kerja_gaji_bulan":         "0",
		"masa_kerja_gaji_tahun":         "3",
		"uuid_pangkat_golongan_pegawai": "c6101e45-09e3-11eb-8c77-7eb0d4a3c7a0",
		"uuid_pejabat_penetap":          "0e6047fd-9463-11eb-b06a-000c2977b907",
		"uuid_status_pengangkatan":      "47dd67dc-0479-11eb-8c77-7eb0d4a3c7a0",
		"uuid_jenis_ijazah":             "74cb40d3-ee86-11ea-8c77-7eb0d4a3c7a0",

		// "uuid_pegawai_penetap":        "c19bec78-4275-11ea-b751-7eb0d4a3c7a0",
	})

	if err != nil {
		t.Fatal(err)
	}
	wr.Close()

	// create request
	// uuid := "6215c058-1e3d-11eb-a014-7eb0d4a3c7a0"
	uuid := "dfef3d4d-2ffe-11eb-a014-7eb0d4a3c7a0"
	// q := make(url.Values)
	// q.Set("uuid_sk_pengangkatan_tendik", uuid)
	// baseURL := server.URL + "/public/api/v1/sk-pengangkatan-tendik?" + q.Encode()
	baseURL := server.URL + "/public/api/v1/sk-pengangkatan-tendik?uuid_sk_pengangkatan_tendik=" + uuid
	fmt.Printf("[DEBUG] base url: %s\n", baseURL)
	req, err := http.NewRequest(http.MethodPut, baseURL, bytes.NewReader((wbuf.Bytes())))
	req.Header.Set("Content-Type", wr.FormDataContentType())
	req.Header.Set("X-Member", "admin 1")

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

	// var any interface{}
	// err = json.Unmarshal(rec.Body.Bytes(), &any)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// fmt.Printf("[DEBUG] raw result: %s\n", rec.Body.String())

	// result, ok := any.(map[string]interface{})
	// if !ok {
	// 	t.Fatal()
	// }
	// // fmt.Printf("[DEBUG] any: %#v\n", any)

	// data, ok := result["data"].([]interface{})
	// if !ok {
	// 	t.Fatal()
	// }
	// if len(data) == 0 {
	// 	t.Fatal("should not be empty")
	// }

	// for _, v := range data {
	// 	fmt.Printf("[DEBUG]: %+v\n", v)
	// }
	// fmt.Printf("[DEBUG] data: %+v\n", result["data"])
}

func TestHandleGetSkPengangkatanTendik(t *testing.T) {
	// init server
	e := echo.New()
	e.Use(router.SetResponseTimeout)

	db, err := database.Connect()
	if err != nil {
		t.Skip("failed connect db:", err)
	}
	gormDb, err := database.InitGorm(db)
	if err != nil {
		t.Fatal(err)
	}
	a := app.App{DB: db, GormDB: gormDb}
	router.InitRoute(a, e)
	server := httptest.NewServer(e)
	defer server.Close()

	// create request
	uuidSkPengangkatanTendik := "dfef3d4d-2ffe-11eb-a014-7eb0d4a3c7a0"
	baseURL := server.URL + "/public/api/v1/sk-pengangkatan-tendik-v2?uuid_sk_pengangkatan_tendik=" + uuidSkPengangkatanTendik
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

type Mhs struct {
	Name string
}

func TestReflect(t *testing.T) {
	var s string = "abc"
	t.Log("type:", reflect.TypeOf(s))

	m := Mhs{"haris"}
	t.Log("type:", reflect.TypeOf(m))

	sk := model.SkPengangkatanTendik{GajiPokok: 123}
	t.Log("type:", reflect.TypeOf(sk))

}

func TestNew(t *testing.T) {
	m := new(Mhs)
	t.Log("m:", m)

	o := &Mhs{"ahmad"}
	n := *o
	n.Name = "fahmi"
	t.Log("o:", o)
	t.Log("n:", &n)
}
