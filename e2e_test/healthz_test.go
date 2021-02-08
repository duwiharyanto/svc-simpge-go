package e2e_test

import (
	"net/http"
	"testing"
)

func healthCheck(client *http.Client, baseURL string, printResponse bool) func(*testing.T) {
	tf := func(t *testing.T) {
		url := baseURL + "/public/api/v1/healthz"
		resBody, err := execHttpRequest(client, http.MethodGet, url, "", nil, printResponse)
		if err != nil {
			t.Fatal(err.Error())
		}
		if resBody.StatusCode != http.StatusOK {
			t.Fatalf("got status code: %d, want: %d", resBody.StatusCode, http.StatusOK)
		}
	}
	return tf
}
