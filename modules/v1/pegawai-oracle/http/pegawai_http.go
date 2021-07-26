package pegawaihttp

import (
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

var kepegawaianYayasanSimpegURL = pegawaiSimpegURL + "/%s/kepegawaian-yayasan"

const (
	contentTypeJSON = "application/json"
)

func GetKepegawaianYayasan(ctx context.Context, client *http.Client, nip string) (*model.KepegawaianYayasanSimpeg, error) {
	endpoint := fmt.Sprintf(kepegawaianYayasanSimpegURL, nip)
	res, err := app.SendHttpRequest(ctx, client, http.MethodGet, endpoint, "", nil, nil)
	if err != nil {
		return nil, fmt.Errorf("error send http request: %w", err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error read response body: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error status not ok: %s", strings.Trim(fmt.Sprintf("%q", resBody), `"`))
	}

	var result model.GetKepegawaianYayasanSimpegResult
	err = json.Unmarshal(resBody, &result)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response body: %w", err)
	}

	pegawai := result.Data

	if (pegawai == model.KepegawaianYayasanSimpeg{}) {
		return nil, nil
	}

	pegawai.NIP = nip

	return &pegawai, nil
}

func UpdateKepegawaianYayasan(ctx context.Context, client *http.Client, pegawai *model.KepegawaianYayasanSimpeg) error {
	endpoint := fmt.Sprintf(kepegawaianYayasanSimpegURL, pegawai.NIP)
	header := map[string]string{
		"X-Member": pegawai.UserUpdate,
	}
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

func CreateKepegawaianYayasan(ctx context.Context, client *http.Client, pegawai *model.KepegawaianYayasanSimpeg) error {
	endpoint := fmt.Sprintf(pegawaiSimpegURL)
	header := map[string]string{
		"X-Member": pegawai.UserInput,
	}

	j, _ := json.MarshalIndent(pegawai, "", "\t")
	fmt.Printf("DEBUG reqbody pegawai : \n%s\n", j)

	res, err := app.SendHttpRequest(ctx, client, http.MethodPost, endpoint, contentTypeJSON, header, pegawai)
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

	fmt.Printf("[DEBUG] response from create kepegawaian yayasan simpeg: %s\n", resBody)
	return nil
}
