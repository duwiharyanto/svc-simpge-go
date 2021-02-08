package e2e_test

import (
	"net/http"
	"testing"
)

func getMasterJenisFilePendidikan(client *http.Client, url string, printResponse bool) func(*testing.T) {
	url = url + "/public/api/v1/master-jenis-file-pendidikan"
	return func(t *testing.T) {
		resBody, err := execHttpRequest(client, http.MethodGet, url, "", nil, printResponse)
		if err != nil {
			t.Fatal(err.Error())
		}
		if len(resBody.Data) == 0 {
			t.Fatal("data is empty")
		}
	}
}

func getMasterJenjangPendidikan(client *http.Client, url string, printResponse bool) func(*testing.T) {
	url = url + "/public/api/v1/master-jenjang-pendidikan"
	return func(t *testing.T) {
		resBody, err := execHttpRequest(client, http.MethodGet, url, "", nil, printResponse)
		if err != nil {
			t.Fatal(err.Error())
		}
		if len(resBody.Data) == 0 {
			t.Fatal("data is empty")
		}
	}
}

func getMasterPT(client *http.Client, url string, printResponse bool) func(*testing.T) {
	url = url + "/public/api/v1/master-perguruan-tinggi"
	return func(t *testing.T) {
		resBody, err := execHttpRequest(client, http.MethodGet, url, "", nil, printResponse)
		if err != nil {
			t.Fatal(err.Error())
		}
		if len(resBody.Data) == 0 {
			t.Fatal("data is empty")
		}
	}
}

func getMasterPTWithFlagLuarNegeri(client *http.Client, url string, printResponse bool) func(*testing.T) {
	url = url + "/public/api/v1/master-perguruan-tinggi?flag_institusi_luar_negeri=1"
	return func(t *testing.T) {
		resBody, err := execHttpRequest(client, http.MethodGet, url, "", nil, printResponse)
		if err != nil {
			t.Fatal(err.Error())
		}
		if len(resBody.Data) == 0 {
			t.Fatal("data is empty")
		}
	}
}

func getMasterPTWithNama(client *http.Client, url string, printResponse bool) func(*testing.T) {
	url = url + "/public/api/v1/master-perguruan-tinggi?nama=nsti"
	return func(t *testing.T) {
		resBody, err := execHttpRequest(client, http.MethodGet, url, "", nil, printResponse)
		if err != nil {
			t.Fatal(err.Error())
		}
		if len(resBody.Data) == 0 {
			t.Fatal("data is empty")
		}
	}
}
