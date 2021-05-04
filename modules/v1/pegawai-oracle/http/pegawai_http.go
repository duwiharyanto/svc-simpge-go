package pegawaihttp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/pegawai-oracle/model"
)

var pegawaiSimpegURL = fmt.Sprintf("%s/public/api/v1/pegawai", os.Getenv("URL_HCM_SIMPEG_SERVICE"))

// var pegawaiSimpegURL = fmt.Sprintf("%s/pegawai", os.Getenv("URL_HCM_SIMPEG_SERVICE"))
var pegawaiStatusSimpegURL = fmt.Sprintf("%s/public/api/v1/pegawai-status", os.Getenv("URL_HCM_SIMPEG_SERVICE"))
var kepegawaianYayasanSimpegURL = pegawaiSimpegURL + "/%s/kepegawaian-yayasan"

const (
	contentTypeJSON = "application/json"
)

func GetPegawai(ctx context.Context, client *http.Client, nip string) (*model.PegawaiSimpeg, error) {
	endpoint := fmt.Sprintf("%s/%s", pegawaiSimpegURL, nip)
	// fmt.Printf("[DEBUG] endpoint: %s\n", endpoint)
	res, err := app.SendHttpRequest(ctx, client, http.MethodGet, endpoint, contentTypeJSON, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("error send http request: %w", err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error read response body: %w", err)
	}
	// fmt.Printf("[DEBUG] raw res body: %s\n", resBody)
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error status not ok: %s", strings.Trim(fmt.Sprintf("%q", resBody), `"`))
	}

	var result model.GetPegawaiSimpegResult
	err = json.Unmarshal(resBody, &result)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response body: %w", err)
	}
	pegawai := result.Data
	// fmt.Printf("[DEBUG] pegawai after unmarshal: %+v\n", pegawai)

	if (pegawai == model.PegawaiSimpeg{}) {
		return nil, nil
	}

	fmt.Printf("[DEBUG] pegawai simpeg: %+v\n", pegawai)
	pegawai.NIP = nip
	pegawai.TglUpdate = ""
	pegawai.UserUpdate = ""

	return &pegawai, nil
}

func GetKepegawaianYayasan(ctx context.Context, client *http.Client, nip string) (*model.KepegawaianYayasanSimpeg, error) {
	endpoint := fmt.Sprintf(kepegawaianYayasanSimpegURL, nip)
	// fmt.Printf("[DEBUG] endpoint: %s\n", endpoint)
	res, err := app.SendHttpRequest(ctx, client, http.MethodGet, endpoint, "", nil, nil)
	if err != nil {
		return nil, fmt.Errorf("error send http request: %w", err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error read response body: %w", err)
	}

	var out bytes.Buffer
	json.Indent(&out, resBody, "", "\t")
	fmt.Printf("[DEBUG] raw res body: %s\n", out.Bytes())
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error status not ok: %s", strings.Trim(fmt.Sprintf("%q", resBody), `"`))
	}

	var result model.GetKepegawaianYayasanSimpegResult
	err = json.Unmarshal(resBody, &result)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response body: %w", err)
	}

	pegawai := result.Data
	// fmt.Printf("[DEBUG] pegawai after unmarshal: %+v\n", pegawai)

	if (pegawai == model.KepegawaianYayasanSimpeg{}) {
		return nil, nil
	}

	// fmt.Printf("[DEBUG] pegawai simpeg: %#v\n", pegawai)
	pegawai.NIP = nip
	// pegawai.TglUpdate = ""
	// pegawai.UserUpdate = ""

	return &pegawai, nil
}

func UpdateKepegawaianYayasan(ctx context.Context, client *http.Client, pegawai *model.KepegawaianYayasanSimpeg) error {
	endpoint := fmt.Sprintf(kepegawaianYayasanSimpegURL, pegawai.NIP)
	// fmt.Printf("DEBUG endpoint : %+v \n", endpoint)
	// fmt.Printf("DEBUG pegawaisimpegURL : %+v \n", pegawaiSimpegURL)
	// authToken := os.Getenv("AUTH_TOKEN")
	header := map[string]string{
		// "Authorization": authToken,
		"X-Member": pegawai.UserUpdate,
	}
	// fmt.Printf("DEBUG header : %+v \n", header)
	// fmt.Printf("DEBUG pegawai : %+v \n", pegawai)
	res, err := app.SendHttpRequest(ctx, client, http.MethodPut, endpoint, contentTypeJSON, header, pegawai)
	if err != nil {
		return fmt.Errorf("error send http request: %w", err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error read response body: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error status not ok: %s", resBody)
	}

	fmt.Printf("[DEBUG] response from update kepegawaian yayasan simpeg: %s\n", resBody)
	return nil
}

func GetPegawaiStatus(ctx context.Context, client *http.Client, nip string) (*model.PegawaiStatusSimpeg, error) {
	endpoint := fmt.Sprintf("%s/%s", pegawaiStatusSimpegURL, nip)
	res, err := app.SendHttpRequest(ctx, client, http.MethodGet, endpoint, contentTypeJSON, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("error send http request: %w", err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error read response body: %w", err)
	}
	// fmt.Printf("[DEBUG] raw res body: %s\n", resBody)

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error status not ok: %s", strings.Trim(fmt.Sprintf("%q", resBody), `"`))
	}

	var result model.GetPegawaiStatusSimpegResult
	err = json.Unmarshal(resBody, &result)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response body: %w", err)
	}
	pegawaiStatus := result.Data
	// fmt.Printf("[DEBUG] pegawai status after unmarshal: %+v\n", pegawaiStatus)

	if (pegawaiStatus == model.PegawaiStatusSimpeg{}) {
		return nil, nil
	}

	fmt.Printf("[DEBUG] pegawai status simpeg: %+v\n", pegawaiStatus)
	pegawaiStatus.NIP = nip
	pegawaiStatus.TglUpdate = ""
	pegawaiStatus.UserUpdate = ""

	return &pegawaiStatus, nil
}

func UpdatePegawaiStatus(ctx context.Context, client *http.Client, pegawaiStatus model.PegawaiStatusSimpeg) error {
	fmt.Printf("[DEBUG] reqbody: %+v\n", pegawaiStatus)
	endpoint := fmt.Sprintf("%s/%s", pegawaiStatusSimpegURL, pegawaiStatus.NIP)
	header := map[string]string{"X-Member": pegawaiStatus.UserUpdate}

	res, err := app.SendHttpRequest(ctx, client, http.MethodPut, endpoint, contentTypeJSON, header, pegawaiStatus)
	if err != nil {
		return fmt.Errorf("error send http request: %w", err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error read response body: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		// return fmt.Errorf("error status not ok: %s", strings.Trim(fmt.Sprintf("%q", resBody), `"`))
		// return fmt.Errorf("error status not ok: %q", resBody)
		return fmt.Errorf("error status not ok: %s", resBody)
	}

	fmt.Printf("[DEBUG] response from update pegawai status simpeg: %s\n", resBody)
	return nil
}
