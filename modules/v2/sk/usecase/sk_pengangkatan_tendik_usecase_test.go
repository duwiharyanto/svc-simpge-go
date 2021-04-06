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
	"os"
	"reflect"
	"strings"
	"svc-insani-go/app"
	"svc-insani-go/app/database"
	"svc-insani-go/app/minio"
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

func TestHandleCreateSkPengangkatanTendik(t *testing.T) {
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

	mc, err := minio.Connect()
	if err != nil {
		t.Fatal(err)
	}

	a := app.App{DB: db, GormDB: gormDb, MinioClient: mc, MinioBucketName: "insani"}
	router.InitRoute(a, e)
	server := httptest.NewServer(e)
	defer server.Close()

	// init form data
	wbuf := &bytes.Buffer{}
	wr := multipart.NewWriter(wbuf)
	err = FillFormDataFieldMap(wr, map[string]string{
		"uuid_jenis_sk":                 "ebc9e2c0-ee60-11ea-8c77-7eb0d4a3c7a0",
		"gaji_pokok":                    "200000",
		"tanggal_ditetapkan":            "2021-04-03",
		"nomor_sk":                      "no/sk-haris/005/1",
		"tentang_sk":                    "tentang sk tendik 5",
		"tmt":                           "2021-09-11",
		"uuid_kelompok_sk_pengangkatan": "742024ac-4fea-11eb-bf95-a74048ab8082",
		"uuid_jabatan_fungsional":       "aeb51718-2fc6-11eb-a014-7eb0d4a3c7a0",
		"uuid_jabatan_penetap":          "6bd8793c-9461-11eb-b06a-000c2977b907",
		"uuid_unit_pengangkat":          "798c80c4-1fd3-11eb-a014-7eb0d4a3c7a0",
		"uuid_unit_kerja":               "798c8162-1fd3-11eb-a014-7eb0d4a3c7a0",
		"masa_kerja_diakui_bulan":       "55",
		"masa_kerja_diakui_tahun":       "55",
		"masa_kerja_ril_bulan":          "55",
		"masa_kerja_ril_tahun":          "55",
		"masa_kerja_gaji_bulan":         "55",
		"masa_kerja_gaji_tahun":         "55",
		"uuid_pangkat_golongan_pegawai": "c6101e45-09e3-11eb-8c77-7eb0d4a3c7a0",
		"uuid_pejabat_penetap":          "0e6047fd-9463-11eb-b06a-000c2977b907",
		"uuid_status_pengangkatan":      "47dd67dc-0479-11eb-8c77-7eb0d4a3c7a0",
		"uuid_jenis_ijazah":             "74cb40d3-ee86-11ea-8c77-7eb0d4a3c7a0",

		// "uuid_pegawai_penetap":        "c19bec78-4275-11ea-b751-7eb0d4a3c7a0",
	})

	if err != nil {
		t.Fatal(err)
	}

	// add file to form
	file, err := os.Open("./test.pdf")
	if err != nil {
		t.Fatal(err)
	}

	part, err := wr.CreateFormFile("file_sk", "./test.pdf")
	if err != nil {
		t.Fatal(err)
	}

	io.Copy(part, file)
	file.Close()

	wr.Close()

	// create request
	kdKelompokPegawai := "AD"
	// kdKelompokPegawai := "ED"
	uuidPegawai := "d8c26983-1437-11eb-a014-7eb0d4a3c7a0"
	baseURL := server.URL + "/public/api/v1/sk-pengangkatan-tendik?kd_kelompok_pegawai=" + kdKelompokPegawai + "&uuid_pegawai=" + uuidPegawai
	// fmt.Printf("[DEBUG] base url: %s\n", baseURL)
	req, err := http.NewRequest(http.MethodPost, baseURL, bytes.NewReader((wbuf.Bytes())))
	req.Header.Set("Content-Type", wr.FormDataContentType())
	req.Header.Set("X-Member", "admin 4")

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

	var resBody map[string]interface{}
	err = json.Unmarshal(rawResBodyJSON, &resBody)
	if err != nil {
		t.Fatal(err)
	}

	if msg, exist := resBody["message"]; !exist || !strings.Contains(strings.ToLower(msg.(string)), "berhasil") {
		fmt.Printf("[DEBUG] name: %+v\n", resBody)
		t.Fatal("should return success message")
	}

}

func TestHandleUpdateSkPengangkatanTendik(t *testing.T) {
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

	mc, err := minio.Connect()
	if err != nil {
		t.Fatal(err)
	}

	a := app.App{DB: db, GormDB: gormDb, MinioClient: mc, MinioBucketName: "insani"}
	router.InitRoute(a, e)
	server := httptest.NewServer(e)
	defer server.Close()

	// init form data
	wbuf := &bytes.Buffer{}
	wr := multipart.NewWriter(wbuf)
	err = FillFormDataFieldMap(wr, map[string]string{
		"uuid_jenis_sk":                 "ebc9e2c0-ee60-11ea-8c77-7eb0d4a3c7a0",
		"gaji_pokok":                    "999999",
		"tanggal_ditetapkan":            "2021-01-09",
		"nomor_sk":                      "no/sk-ega/9/5",
		"tentang_sk":                    "tentang sk tendik 9",
		"tmt":                           "2021-09-09",
		"uuid_kelompok_sk_pengangkatan": "742024ac-4fea-11eb-bf95-a74048ab8082",
		"uuid_jabatan_fungsional":       "aeb51718-2fc6-11eb-a014-7eb0d4a3c7a0",
		// "uuid_jabatan_penetap":          "6bd8793c-9461-11eb-b06a-000c2977b907",
		"uuid_jabatan_penetap":          "352be8f4-96de-11eb-86fa-0f2381201b27", // local
		"uuid_unit_pengangkat":          "798c80c4-1fd3-11eb-a014-7eb0d4a3c7a0",
		"uuid_unit_kerja":               "798c8162-1fd3-11eb-a014-7eb0d4a3c7a0",
		"masa_kerja_diakui_bulan":       "10",
		"masa_kerja_diakui_tahun":       "10",
		"masa_kerja_ril_bulan":          "10",
		"masa_kerja_ril_tahun":          "10",
		"masa_kerja_gaji_bulan":         "10",
		"masa_kerja_gaji_tahun":         "10",
		"uuid_pangkat_golongan_pegawai": "c6101e45-09e3-11eb-8c77-7eb0d4a3c7a0",
		// "uuid_pejabat_penetap":          "0e6047fd-9463-11eb-b06a-000c2977b907",
		"uuid_pejabat_penetap":     "4c835e92-96de-11eb-86fa-0f2381201b27", // local
		"uuid_status_pengangkatan": "47dd67dc-0479-11eb-8c77-7eb0d4a3c7a0",
		"uuid_jenis_ijazah":        "74cb40d3-ee86-11ea-8c77-7eb0d4a3c7a0",

		// "uuid_pegawai_penetap":        "c19bec78-4275-11ea-b751-7eb0d4a3c7a0",
	})

	if err != nil {
		t.Fatal(err)
	}

	// add file to form
	// file, err := os.Open("./test.pdf")
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// part, err := wr.CreateFormFile("file_sk", "./test.pdf")
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// io.Copy(part, file)
	// file.Close()

	wr.Close()

	// create request
	uuid := "0cbc769f-2fb7-11eb-a014-7eb0d4a3c7a0" // local
	// uuid := "ac34398c-5453-493c-b1bd-b6bb426dd4ae"
	// q := make(url.Values)
	// q.Set("uuid_sk_pengangkatan_tendik", uuid)
	// baseURL := server.URL + "/public/api/v1/sk-pengangkatan-tendik?" + q.Encode()
	baseURL := server.URL + "/public/api/v1/sk-pengangkatan-tendik?uuid_sk_pengangkatan_tendik=" + uuid
	// fmt.Printf("[DEBUG] base url: %s\n", baseURL)
	req, err := http.NewRequest(http.MethodPut, baseURL, bytes.NewReader((wbuf.Bytes())))
	req.Header.Set("Content-Type", wr.FormDataContentType())
	req.Header.Set("X-Member", "admin 7")

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

	var resBody map[string]interface{}
	err = json.Unmarshal(rawResBodyJSON, &resBody)
	if err != nil {
		t.Fatal(err)
	}

	if msg, exist := resBody["message"]; !exist || !strings.Contains(strings.ToLower(msg.(string)), "berhasil") {
		fmt.Printf("[DEBUG] name: %+v\n", resBody)
		t.Fatal("should return success message")
	}

}

func TestHandleGetSkPengangkatanTendik(t *testing.T) {
	// init server
	e := echo.New()
	e.Use(router.SetResponseTimeout)

	db, err := database.Connect()
	if err != nil {
		t.Skip("failed connect db:", err)
	}

	gormDb, err := database.InitGorm(db, false)
	if err != nil {
		t.Fatal(err)
	}

	mc, err := minio.Connect()
	if err != nil {
		t.Fatal(err)
	}

	a := app.App{DB: db, GormDB: gormDb, MinioBucketName: "insani", MinioClient: mc}
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

func TestHandleDeleteSkPengangkatanTendik(t *testing.T) {
	// init server
	e := echo.New()
	e.Use(router.SetResponseTimeout)

	db, err := database.Connect()
	if err != nil {
		t.Skip("failed connect db:", err)
	}

	gormDb, err := database.InitGorm(db, false)
	if err != nil {
		t.Fatal(err)
	}

	a := app.App{DB: db, GormDB: gormDb}
	router.InitRoute(a, e)
	server := httptest.NewServer(e)
	defer server.Close()

	// create request
	// uuidSkPengangkatanTendik := "dfef3d4d-2ffe-11eb-a014-7eb0d4a3c7a0"
	uuidSkPengangkatanTendik := "e62d4941-304e-11eb-a014-7eb0d4a3c7a0" // local
	baseURL := server.URL + "/public/api/v1/sk-pengangkatan-tendik?uuid_sk_pengangkatan_tendik=" + uuidSkPengangkatanTendik
	// fmt.Printf("[DEBUG] base url: %s\n", baseURL)
	req, err := http.NewRequest(http.MethodDelete, baseURL, nil)
	req.Header.Set("X-Member", "admin-1")
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

	var resBody map[string]interface{}
	err = json.Unmarshal(rawResBodyJSON, &resBody)
	if err != nil {
		t.Fatal(err)
	}

	if msg, exist := resBody["message"]; !exist || !strings.Contains(strings.ToLower(msg.(string)), "berhasil") {
		fmt.Printf("[DEBUG] name: %+v\n", resBody)
		t.Fatal("should return success message")
	}
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
