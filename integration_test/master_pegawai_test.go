package integrationtest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func GetKelompokPegawai(t *testing.T, s *TestServer, resp map[string]interface{}) func(*testing.T) {
	return func(t *testing.T) {
		targetUrl := fmt.Sprintf("%s/%s/%s", s.Server.URL, v1path, "master-kelompok-pegawai")
		hres, err := s.Client.SendRequest(http.MethodGet, targetUrl, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		if resp == nil {
			resp = make(map[string]interface{})
		}
		err = json.Unmarshal(hres, &resp)
		if err != nil {
			t.Fatal(err)
		}
		if len(resp["data"].([]interface{})) == 0 {
			t.Fatal("should not be empty")
		}
	}
}

func GetUnitKerja(t *testing.T, s *TestServer, resp map[string]interface{}) func(*testing.T) {
	return func(t *testing.T) {
		targetUrl := fmt.Sprintf("%s/%s/%s", s.Server.URL, v1path, "unit-kerja")
		hres, err := s.Client.SendRequest(http.MethodGet, targetUrl, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		if resp == nil {
			resp = make(map[string]interface{})
		}
		err = json.Unmarshal(hres, &resp)
		if err != nil {
			t.Fatal(err)
		}
		if len(resp["data"].([]interface{})) == 0 {
			t.Fatal("should not be empty")
		}
	}
}

func GetBagianKerja(t *testing.T, s *TestServer, resp map[string]interface{}) func(*testing.T) {
	return func(t *testing.T) {
		targetUrl := fmt.Sprintf("%s/%s/%s", s.Server.URL, v1path, "bagian-kerja")
		hres, err := s.Client.SendRequest(http.MethodGet, targetUrl, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		if resp == nil {
			resp = make(map[string]interface{})
		}
		err = json.Unmarshal(hres, &resp)
		if err != nil {
			t.Fatal(err)
		}
		if len(resp["data"].([]interface{})) == 0 {
			t.Fatal("should not be empty")
		}
	}
}

func GetLokasiKerja(t *testing.T, s *TestServer, resp map[string]interface{}) func(*testing.T) {
	return func(t *testing.T) {
		targetUrl := fmt.Sprintf("%s/%s/%s", s.Server.URL, v1path, "master-lokasi-kerja")
		hres, err := s.Client.SendRequest(http.MethodGet, targetUrl, nil, nil)
		if err != nil {
			t.Fatal(err)
		}
		if resp == nil {
			resp = make(map[string]interface{})
		}
		err = json.Unmarshal(hres, &resp)
		if err != nil {
			t.Fatal(err)
		}
		if len(resp["data"].([]interface{})) == 0 {
			t.Fatal("should not be empty")
		}
	}
}
