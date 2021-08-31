package http

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	netHttp "net/http"
	"os"
	"svc-insani-go/modules/v1/presensi/model"
)

var createPresenceUserURL = fmt.Sprintf("%s/public/api/v2/user", os.Getenv("URL_V2_PRESENCE_SERVICE"))

func CreateUserPresensi(ctx context.Context, client *netHttp.Client, user *model.UserPresensi) error {
	formDataBuffer := &bytes.Buffer{}
	formDataWriter := multipart.NewWriter(formDataBuffer)
	formDataWriter.WriteField("nip", user.Nip)
	formDataWriter.WriteField("kd_jenis_pegawai", user.KdJenisPegawai)
	formDataWriter.WriteField("kd_unit_kerja", user.KdUnitKerja)
	formDataWriter.WriteField("kd_lokasi_kerja", user.KdLokasiKerja)
	formDataWriter.WriteField("tmt", user.Tmt)
	formDataWriter.WriteField("kd_jenis_presensi", user.KdJenisPresensi)
	contentType := formDataWriter.FormDataContentType()
	formDataWriter.Close()

	req, err := netHttp.NewRequestWithContext(ctx, netHttp.MethodPost, createPresenceUserURL, formDataBuffer)
	if err != nil {
		return fmt.Errorf("error new request: %w", err)
	}

	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	req.Header.Set("X-Member", user.UserUpdate)
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error read response body: %w", err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error read response body: %w", err)
	}

	if res.StatusCode != netHttp.StatusOK {
		fmt.Printf("[DEBUG] res stat code: %+v\n", res.StatusCode)
		fmt.Printf("[DEBUG] res body: %s\n", resBody)
		return fmt.Errorf("error status not ok: %s", resBody)
	}

	fmt.Printf("[DEBUG] response from create presence user: %s\n", resBody)
	return nil
}
