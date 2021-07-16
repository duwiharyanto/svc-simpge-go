package integrationtest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func (s *TestServer) CreatePersonal(t *testing.T) {
	formDataBuffer := &bytes.Buffer{}
	multipartWriter := multipart.NewWriter(formDataBuffer)
	err := FillFormDataFieldMap(multipartWriter, map[string]string{
		"nama_lengkap":  "Integration Test A",
		"tempat_lahir":  "Padang",
		"tanggal_lahir": "1996-06-01",
	})

	multipartWriter.Close()

	testUrl := fmt.Sprintf("%s%s", s.Server.URL, v1path)
	fmt.Printf("[DEBUG] testUrl: %s\n", testUrl)
	req, err := http.NewRequest(
		http.MethodPost,
		testUrl,
		bytes.NewReader((formDataBuffer.Bytes())),
	)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	req.Header.Set("X-Member", "testadmin")

	res, err := s.Client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	rawResBodyJSON, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	var resBody map[string]interface{}
	err = json.Unmarshal(rawResBodyJSON, &resBody)
	if err != nil {
		t.Fatal(err)
	}

	if msg, exist := resBody["message"]; !exist || !strings.Contains(strings.ToLower(msg.(string)), "Berhasil simpan personal") {
		fmt.Printf("[DEBUG] res body: %+v\n", resBody)
		t.Fatal("should return success message")
	}

}

func (s *TestServer) TestPersonal() (map[string]interface{}, error) {
	v := url.Values{}
	v.Set("status_personal", "PGW")
	v.Set("limit", "10")
	v.Set("offset", "0")
	v.Set("cari", "user")
	testUrl := fmt.Sprintf("%s%s?%s", s.Server.URL, v1path, v.Encode())

	fmt.Printf("[DEBUG] testUrl: %s\n", testUrl)
	req, err := http.NewRequest(
		http.MethodGet,
		testUrl,
		nil,
	)
	if err != nil {
		return nil, err
	}

	res, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}

	rawResBodyJSON, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	var resBody map[string]interface{}
	err = json.Unmarshal(rawResBodyJSON, &resBody)
	if err != nil {
		return nil, err
	}

	return resBody, nil

}
