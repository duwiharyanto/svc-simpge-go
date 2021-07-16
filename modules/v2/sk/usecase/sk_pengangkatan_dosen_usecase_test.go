package usecase_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"svc-insani-go/app"
	"svc-insani-go/app/database"
	"svc-insani-go/app/minio"

	"svc-insani-go/router"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestHandleCreateSkPengangkatanDosen(t *testing.T) {
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

	mc, err := minio.Connect()
	if err != nil {
		t.Fatal(err)
	}

	a := &app.App{DB: db, GormDB: gormDb, MinioClient: mc, MinioBucketName: "insani"}
	fmt.Print(a)
	// router.InitRoute(a, e)
	server := httptest.NewServer(e)
	defer server.Close()

	// init form data
	wbuf := &bytes.Buffer{}
	wr := multipart.NewWriter(wbuf)
	err = FillFormDataFieldMap(wr, map[string]string{
		"uuid_jenis_sk":                 "ebc9e2c0-ee60-11ea-8c77-7eb0d4a3c7a0",
		"gaji_pokok":                    "200000",
		"tanggal_ditetapkan":            "2021-04-06",
		"nomor_sk":                      "no/sk-haris/006/1",
		"tentang_sk":                    "tentang sk tendik 6",
		"tmt":                           "2021-09-11",
		"uuid_kelompok_sk_pengangkatan": "742024ac-4fea-11eb-bf95-a74048ab8082",
		"uuid_jabatan_fungsional":       "aeb51718-2fc6-11eb-a014-7eb0d4a3c7a0",
		"uuid_jabatan_penetap":          "352be8f4-96de-11eb-86fa-0f2381201b27", // local
		// "uuid_jabatan_penetap":          "6bd8793c-9461-11eb-b06a-000c2977b907",
		"uuid_unit_pengangkat":          "798c80c4-1fd3-11eb-a014-7eb0d4a3c7a0",
		"uuid_unit_kerja":               "798c8162-1fd3-11eb-a014-7eb0d4a3c7a0",
		"masa_kerja_diakui_bulan":       "6",
		"masa_kerja_diakui_tahun":       "6",
		"masa_kerja_ril_bulan":          "6",
		"masa_kerja_ril_tahun":          "6",
		"masa_kerja_gaji_bulan":         "6",
		"masa_kerja_gaji_tahun":         "6",
		"uuid_pangkat_golongan_pegawai": "c6101e45-09e3-11eb-8c77-7eb0d4a3c7a0",
		"uuid_pejabat_penetap":          "4c835e92-96de-11eb-86fa-0f2381201b27", // local
		// "uuid_pejabat_penetap":          "0e6047fd-9463-11eb-b06a-000c2977b907",
		"uuid_status_pengangkatan": "47dd67dc-0479-11eb-8c77-7eb0d4a3c7a0",
		"uuid_jenis_ijazah":        "74cb40d3-ee86-11ea-8c77-7eb0d4a3c7a0",
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
	kdKelompokPegawai := "AD"
	// kdKelompokPegawai := "ED"
	uuidPegawai := "d8c26983-1437-11eb-a014-7eb0d4a3c7a0"
	baseURL := server.URL + "/public/api/v1/sk-pengangkatan-dosen?kd_kelompok_pegawai=" + kdKelompokPegawai + "&uuid_pegawai=" + uuidPegawai
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

func TestHandleUpdateSkPengangkatanDosen(t *testing.T) {
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

	mc, err := minio.Connect()
	if err != nil {
		t.Fatal(err)
	}

	a := &app.App{DB: db, GormDB: gormDb, MinioClient: mc, MinioBucketName: "insani"}
	fmt.Print(a)
	// router.InitRoute(a, e)
	server := httptest.NewServer(e)
	defer server.Close()

	// init form data
	wbuf := &bytes.Buffer{}
	wr := multipart.NewWriter(wbuf)
	err = FillFormDataFieldMap(wr, map[string]string{
		"uuid_jenis_sk":                 "ebc9e2c0-ee60-11ea-8c77-7eb0d4a3c7a0",
		"gaji_pokok":                    "999999",
		"tanggal_ditetapkan":            "2021-01-09",
		"nomor_sk":                      "no/sk-edit/8/5",
		"tentang_sk":                    "tentang sk tendik 8",
		"tmt":                           "2021-09-09",
		"uuid_kelompok_sk_pengangkatan": "742024ac-4fea-11eb-bf95-a74048ab8082",
		"uuid_jabatan_fungsional":       "aeb51718-2fc6-11eb-a014-7eb0d4a3c7a0",
		"uuid_jabatan_penetap":          "353ea3c2-96de-11eb-86fa-0f2381201b27", // local
		// "uuid_jabatan_penetap":          "6bd8793c-9461-11eb-b06a-000c2977b907",
		"uuid_mata_kuliah":              `["a", "b", "c"]`,
		"uuid_unit_pengangkat":          "798c80c4-1fd3-11eb-a014-7eb0d4a3c7a0",
		"uuid_unit_kerja":               "798c8162-1fd3-11eb-a014-7eb0d4a3c7a0",
		"masa_kerja_diakui_bulan":       "8",
		"masa_kerja_diakui_tahun":       "8",
		"masa_kerja_ril_bulan":          "8",
		"masa_kerja_ril_tahun":          "8",
		"masa_kerja_gaji_bulan":         "8",
		"masa_kerja_gaji_tahun":         "8",
		"uuid_pangkat_golongan_pegawai": "c6101e45-09e3-11eb-8c77-7eb0d4a3c7a0",
		"uuid_pejabat_penetap":          "4c832fbc-96de-11eb-86fa-0f2381201b27", // local
		// "uuid_pejabat_penetap":          "0e6047fd-9463-11eb-b06a-000c2977b907",
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
	uuid := "cfba6500-7974-4a3c-ad4c-268999c5ff5d" // local
	// uuid := "ac34398c-5453-493c-b1bd-b6bb426dd4ae"
	// q := make(url.Values)
	// q.Set("uuid_sk_pengangkatan_tendik", uuid)
	// baseURL := server.URL + "/public/api/v1/sk-pengangkatan-dosen?" + q.Encode()
	baseURL := server.URL + "/public/api/v1/sk-pengangkatan-dosen?uuid_sk_pengangkatan_dosen=" + uuid
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

func TestHandleGetSkPengangkatanDosen(t *testing.T) {
	// init server
	e := echo.New()
	e.Use(router.SetResponseTimeout(context.Background()))

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

	a := &app.App{DB: db, GormDB: gormDb, MinioBucketName: "insani", MinioClient: mc}
	fmt.Print(a)
	// router.InitRoute(a, e)
	server := httptest.NewServer(e)
	defer server.Close()

	// create request
	uuidSkPengangkatanDosen := "cfba6500-7974-4a3c-ad4c-268999c5ff5d" // local
	// uuidSkPengangkatanDosen := "dfef3d4d-2ffe-11eb-a014-7eb0d4a3c7a0"
	baseURL := server.URL + "/public/api/v1/sk-pengangkatan-dosen?uuid_sk_pengangkatan_dosen=" + uuidSkPengangkatanDosen
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

func TestHandleDeleteSkPengangkatanDosen(t *testing.T) {
	// init server
	e := echo.New()
	e.Use(router.SetResponseTimeout(context.Background()))

	db, err := database.Connect()
	if err != nil {
		t.Skip("failed connect db:", err)
	}

	gormDb, err := database.InitGorm(db, false)
	if err != nil {
		t.Fatal(err)
	}

	a := &app.App{DB: db, GormDB: gormDb}
	fmt.Print(a)
	// router.InitRoute(a, e)
	server := httptest.NewServer(e)
	defer server.Close()

	// create request
	// uuidSkPengangkatanDosen := "dfef3d4d-2ffe-11eb-a014-7eb0d4a3c7a0"
	uuidSkPengangkatanDosen := "e62d4941-304e-11eb-a014-7eb0d4a3c7a0" // local
	baseURL := server.URL + "/public/api/v1/sk-pengangkatan-dosen?uuid_sk_pengangkatan_dosen=" + uuidSkPengangkatanDosen
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
