package http

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	netHttp "net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/presensi/model"
	"testing"
	"time"
)

func TestUserPresensi(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	client := &netHttp.Client{Transport: app.DefaultHttpTransport()}
	var user *model.UserPresensi
	var err error

	// t.Run("get", func(t *testing.T) {
	// 	pegawai, err = GetKepegawaianYayasan(ctx, client, nip)
	// 	if err != nil {
	// 		t.Fatalf("error get kepegawaian yayasan: %s", err.Error())
	// 	}
	// 	if pegawai == nil {
	// 		t.Fatal("Should not be nil")
	// 	}
	// 	t.Logf("kepegawaian yayasan: %+v\n%+v\n%+v\n%+v\n", pegawai, pegawai.JenisPegawai, pegawai.KelompokPegawai, pegawai.StatusPegawai)
	// 	if pegawai.PegawaiStatus != nil {
	// 		t.Logf("pegawai status: %+v\n", pegawai.PegawaiStatus)
	// 		// t.Log("no karpeg:", pegawai.PegawaiStatus.NoKarpeg)
	// 	}
	// 	if pegawai.InstansiAsalPtt != nil {
	// 		t.Logf("instansi asal ptt: %+v\n", pegawai.InstansiAsalPtt)
	// 	}
	// })

	t.Run("create", func(t *testing.T) {
		user = &model.UserPresensi{
			Nip:             "251232654",
			KdJenisPegawai:  "ED",
			KdUnitKerja:     "123",
			KdLokasiKerja:   "100",
			Tmt:             "2019-05-04",
			KdJenisPresensi: "21",
			UserUpdate:      "fahmi",
		}
		err = CreateUserPresensi(ctx, client, user)
		if err != nil {
			t.Fatalf("error create user presensi: %s", err.Error())
		}
	})
	t.Run("raw_request", func(t *testing.T) {
		user = &model.UserPresensi{
			Nip:             "231232654",
			KdJenisPegawai:  "AD",
			KdUnitKerja:     "123",
			KdLokasiKerja:   "100",
			Tmt:             "2019-05-02",
			KdJenisPresensi: "12",
			UserUpdate:      "harisf",
		}
		formDataBuffer := &bytes.Buffer{}
		formDataWriter := multipart.NewWriter(formDataBuffer)
		formDataWriter.WriteField("nip", user.Nip)
		formDataWriter.WriteField("kd_jenis_pegawai", user.KdJenisPegawai)
		formDataWriter.WriteField("kd_unit_kerja", user.KdUnitKerja)
		formDataWriter.WriteField("kd_lokasi_kerja", user.KdLokasiKerja)
		formDataWriter.WriteField("tmt", user.Tmt)
		formDataWriter.WriteField("kd_jenis_presensi", user.KdJenisPresensi)

		contentType := formDataWriter.FormDataContentType()
		// reqBody := bytes.NewReader(formDataBuffer.Bytes())
		formDataWriter.Close()

		fmt.Printf("[DEBUG] ctn type in http req: %+v\n", contentType)
		req, err := netHttp.NewRequestWithContext(ctx, netHttp.MethodPost, createPresenceUserURL, formDataBuffer)
		if err != nil {
			t.Fatal(err)
		}

		if contentType != "" {
			req.Header.Set("Content-Type", contentType)
		}

		req.Header.Set("X-Member", user.UserUpdate)
		res, err := client.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}

		if res.StatusCode != netHttp.StatusOK {
			fmt.Printf("[DEBUG] res stat code: %+v\n", res.StatusCode)
			fmt.Printf("[DEBUG] res body: %s\n", resBody)
			t.Fatal(err)
		}

		fmt.Printf("[DEBUG] response from create presence user: %s\n", resBody)
	})
}
