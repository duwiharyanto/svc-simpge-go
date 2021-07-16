package e2e_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"svc-insani-go/app"
	"svc-insani-go/app/database"
	"svc-insani-go/app/minio"
	"testing"

	"github.com/labstack/echo/v4"
)

const (
	contentTypeJSON = "application/json"
)

// struktur response body
type resBody struct {
	Count      int           `json:"count"`
	Data       []interface{} `json:"data"`
	Limit      int           `json:"limit"`
	Offset     int           `json:"offset"`
	Message    string        `json:"message"`
	StatusCode int           `json:"status_code"`
}

// TestServer adalah fungsi yang dijalankan untuk menguji seluruh proses mulai service dijalankan
// hingga menguji semua endpoint yang didaftarkan
func TestServer(t *testing.T) {
	// Persiapan service
	// Jika terdapat masalah pada persiapan service, misalnya: variabel konfigurasi tidak terdaftar pada env variabel
	// maka tes akan dilewati (skip)
	db, err := database.Connect()
	err = db.Ping()
	if err != nil {
		t.Skip("Can't connect to db:", err.Error())
	}

	minioClient, err := minio.Connect()
	if err != nil {
		t.Skip("Can't connect to minio:", err.Error())
	}

	timeLocation := app.GetFixedTimeZone()
	a := &app.App{
		DB:              db,
		HttpClient:      &http.Client{},
		Name:            "Personal Service",
		TimeLocation:    timeLocation,
		MinioClient:     minioClient,
		MinioBucketName: "personal",
	}

	// Persiapan handler service
	e := echo.New()
	// router.InitRoute(a, e)

	// Mendaftarkan handler pada service yang akan diuji
	srv := httptest.NewServer(e)
	// Service dimatikan selesai tes selesai
	defer srv.Close()

	// Daftar testcase yang akan diuji
	// Cara menambahkan testcase yaitu dengan menambahkan item baru pada slice struct berikut
	testcases := []struct {
		// name adalah nama dari testcase yang akan diuji
		name string
		// f adalah fungsi yang akan diuji
		// fungsi dapat dibuat difile berbeda
		f func(*testing.T)
	}{
		{"health_check", healthCheck(a.HttpClient, srv.URL, false)},
	}

	// Eksekusi testcase dilakukan secara berurutan
	for _, tc := range testcases {
		t.Run(tc.name, tc.f)
	}

}

// execHttpRequest berisi pemanggilan-pemanggilan fungsi untuk mengeksekusi http request
func execHttpRequest(client *http.Client, method, url, contentType string, body interface{}, printResponse bool) (*resBody, error) {
	var reqBody io.Reader
	// mengecek http request content type
	switch contentType {
	case contentTypeJSON:
		b := []byte(fmt.Sprint(body))
		reqBody = bytes.NewBuffer(b)
	default:
		reqBody = nil
	}

	// membuat http request baru
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	// mengirim request
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// membaca isi respon mentah
	rawResBodyJSON, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("err read res body: %w", err)
	}

	var rb resBody
	// mengkonversi isi respons ke variabel rb
	err = json.Unmarshal(rawResBodyJSON, &rb)
	if err != nil {
		return nil, fmt.Errorf("err unmarshaling: %w", err)
	}

	// mengkonversi balik rb agar lebih rapi dengan indent
	b, err := json.MarshalIndent(&rb, "", "\t")
	if err != nil {
		return nil, fmt.Errorf("err marshaling: %w", err)
	}

	rb.StatusCode = res.StatusCode

	// if res.StatusCode != expectedStatusCode {
	// 	return nil, fmt.Errorf("response status is not ok!\nresponse:\n%s", b)
	// }

	if printResponse {
		fmt.Printf("\nresponse body:\n%s\n", b)
	}
	return &rb, nil
}
