package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	netHttp "net/http"
	"os"
	"strconv"
	"strings"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/simpeg-oracle/model"
)

var pegawaiSimpegURL = getPegawaiSimpegUrl()
var kepegawaianYayasanSimpegURL = pegawaiSimpegURL + "/%s/kepegawaian-yayasan"

const (
	contentTypeJSON = "application/json"
)

func GetKepegawaianYayasan(ctx context.Context, client *netHttp.Client, nip string) (*model.KepegawaianYayasanSimpeg, error) {
	endpoint := fmt.Sprintf(kepegawaianYayasanSimpegURL, nip)
	header := make(map[string]string)
	header = addDevBearer(header)
	res, err := app.SendHttpRequest(ctx, client, netHttp.MethodGet, endpoint, "", header, nil)
	if err != nil {
		return nil, fmt.Errorf("error send http request: %w", err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error read response body: %w", err)
	}

	if res.StatusCode != netHttp.StatusOK {
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

func UpdateKepegawaianYayasan(ctx context.Context, client *netHttp.Client, pegawai *model.KepegawaianYayasanSimpeg) error {
	endpoint := fmt.Sprintf(kepegawaianYayasanSimpegURL, pegawai.NIP)
	header := map[string]string{
		"X-Member": pegawai.UserUpdate,
	}
	header = addDevBearer(header)
	res, err := app.SendHttpRequest(ctx, client, netHttp.MethodPut, endpoint, contentTypeJSON, header, pegawai)
	if err != nil {
		return fmt.Errorf("error send http request: %w", err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error read response body: %w", err)
	}

	if res.StatusCode != netHttp.StatusOK {
		return fmt.Errorf("error status not ok: %s", resBody)
	}

	fmt.Printf("[DEBUG] response from update kepegawaian yayasan simpeg: %s\n", resBody)

	// log variable pegawai
	b, _ := json.MarshalIndent(pegawai, "", "\t")
	fmt.Println(string(b))
	return nil
}

func CreateKepegawaianYayasan(ctx context.Context, client *netHttp.Client, pegawai *model.KepegawaianYayasanSimpeg) error {
	endpoint := fmt.Sprint(pegawaiSimpegURL)
	header := map[string]string{
		"X-Member": pegawai.UserInput,
	}
	header = addDevBearer(header)

	j, _ := json.Marshal(pegawai)
	fmt.Printf("DEBUG reqbody pegawai : \n%s\n", j)

	res, err := app.SendHttpRequest(ctx, client, netHttp.MethodPost, endpoint, contentTypeJSON, header, pegawai)
	if err != nil {
		return fmt.Errorf("error send http request: %w", err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error read response body: %w", err)
	}

	if res.StatusCode != netHttp.StatusOK {
		return fmt.Errorf("error status not ok: %s", resBody)
	}

	fmt.Printf("[DEBUG] response from create kepegawaian yayasan simpeg: %s\n", resBody)
	return nil
}

func addDevBearer(h map[string]string) map[string]string {
	if useDevSimpegOraSvc, _ := strconv.ParseBool(os.Getenv("USE_DEV_SIMPEG_ORA_SVC")); !useDevSimpegOraSvc {
		return h
	}
	h["Authorization"] = os.Getenv("AUTH_TOKEN")
	return h
}

func getPegawaiSimpegUrl() string {
	urlHcmSimpegService := os.Getenv("URL_HCM_SIMPEG_SERVICE")
	pegawaiSimpegURL := fmt.Sprintf("%s/public/api/v1/pegawai", urlHcmSimpegService)
	if useDevSimpegOraSvc, _ := strconv.ParseBool(os.Getenv("USE_DEV_SIMPEG_ORA_SVC")); useDevSimpegOraSvc {
		pegawaiSimpegURL = fmt.Sprintf("%s/v1/simpeg-ora/pegawai", urlHcmSimpegService)
	}
	return pegawaiSimpegURL
}
