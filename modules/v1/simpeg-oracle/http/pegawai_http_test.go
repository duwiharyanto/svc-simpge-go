package http

import (
	"context"
	"fmt"
	"io/ioutil"
	netHttp "net/http"
	"os"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/simpeg-oracle/model"
	"testing"
	"time"
)

func TestKepegawaianYayasan(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	client := &netHttp.Client{Transport: app.DefaultHttpTransport()}
	// nip := "864210102" // kd_golongan != kd_golongan_kopertis
	// nip := "041002467" // kd_fungsional null
	// nip := "200000112" // kd_golongan null
	// nip := "974200410"
	// nip := "985240101" // tmt_jabatan null
	// nip := "785110201"
	nip := "145230403"
	// nip := "051002465"
	var pegawai *model.KepegawaianYayasanSimpeg
	var err error

	t.Run("get", func(t *testing.T) {
		pegawai, err = GetKepegawaianYayasan(ctx, client, nip)
		if err != nil {
			t.Fatalf("error get kepegawaian yayasan: %s", err.Error())
		}
		if pegawai == nil {
			t.Fatal("Should not be nil")
		}
		t.Logf("kepegawaian yayasan: %+v\n%+v\n%+v\n%+v\n", pegawai, pegawai.JenisPegawai, pegawai.KelompokPegawai, pegawai.StatusPegawai)
		if pegawai.PegawaiStatus != nil {
			t.Logf("pegawai status: %+v\n", pegawai.PegawaiStatus)
			// t.Log("no karpeg:", pegawai.PegawaiStatus.NoKarpeg)
		}
		if pegawai.InstansiAsalPtt != nil {
			t.Logf("instansi asal ptt: %+v\n", pegawai.InstansiAsalPtt)
		}
	})

	t.Run("update", func(t *testing.T) {
		pegawai.UserUpdate = "haris"
		// pegawai.JenisPegawai.KdJenisPegawai = "AD"
		// pegawai.KelompokPegawai.KdKelompokPegawai = "16"
		// pegawai.StatusPegawai.KdStatusPegawai = "TT"
		// pegawai.PegawaiStatus.PangkatYayasan.KdGolongan = "IV"
		// pegawai.PegawaiStatus.PangkatYayasan.KdRuang = "b"
		// pegawai.PegawaiStatus.PangkatYayasan.TmtPangkat = "2020-01-02 00:00:01"
		tmtFungsional := "2021-02-03 00:00:00"
		pegawaiStatus := &model.PegawaiStatus{
			FlagMengajar: "N",
			FlagSekolah:  "Y",
			JabatanFungsional: &model.JabatanFungsional{
				KdFungsional: "06",
				// TmtFungsional: "2021-02-03 00:00:00",

				TmtFungsional: tmtFungsional,
			},
			JabatanFungsionalKopertis: &model.JabatanFungsional{
				KdFungsional:  "",
				TmtFungsional: "",
			},
			AngkaKreditFungsional:  0,
			AngkaKreditKopertis:    0,
			KdHomebasePddikti:      "12341",
			KdHomebaseUii:          "525",
			MasaKerjaGajiTahun:     0,
			MasaKerjaGajiBulan:     3,
			MasaKerjaKopertisTahun: 3,
			MasaKerjaKopertisBulan: 0,
			NoKarpeg:               "",
			NoSkPertama:            "12345678901234567890123456789012345678901234567890",
			TglSkPertama:           "2021-02-03 00:00:00",
			PangkatYayasan: &model.Pangkat{
				TmtPangkat: "",
			},
			PangkatKopertis: &model.Pangkat{
				KdGolongan: "II",
				KdRuang:    "c",
				TmtPangkat: "",
			},
		}
		// pegawai.PegawaiStatus.JabatanFungsional.KdFungsional = "06"
		// pegawai.PegawaiStatus.JabatanFungsional.TmtFungsional = "2021-02-03 00:00:00"
		unit1 := &model.Unit1{
			KdUnit1: "520",
		}
		unit2 := &model.Unit2{
			KdUnit2: "123",
		}
		unit3 := &model.Unit3{
			KdUnit3: "5204",
		}
		lokasiKerja := &model.LokasiKerja{
			KdLokasi: "310",
		}
		pegawai.PegawaiStatus = pegawaiStatus
		// pegawai.FlagPensiun = "Y"
		// pegawai.KdStatusHidup = "N"
		pegawai.Unit1 = unit1
		pegawai.Unit2 = unit2
		pegawai.Unit3 = unit3
		pegawai.LokasiKerja = lokasiKerja
		pegawai.NipKopertis = ""
		pegawai.InstansiAsalPtt = &model.InstansiAsalPtt{
			Instansi:   "UGM",
			Keterangan: "tes",
		}
		err = UpdateKepegawaianYayasan(ctx, client, pegawai)
		if err != nil {
			t.Fatalf("error update pegawai status: %s", err.Error())
		}
	})

}

func TestSimpegOra(t *testing.T) {
	endpoint := getPegawaiSimpegUrl() + "/200000101"
	authToken := os.Getenv("AUTH_TOKEN")
	header := map[string]string{"Authorization": authToken}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	client := &netHttp.Client{Transport: app.DefaultHttpTransport()}
	res, err := app.SendHttpRequest(ctx, client, netHttp.MethodGet, endpoint, contentTypeJSON, header, nil)
	if err != nil {
		t.Fatalf("error send http request: %s", err.Error())
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("error read response body: %s", err)
	}

	if res.StatusCode != netHttp.StatusOK {
		t.Fatalf("error status not ok: %s", resBody)
	}

	fmt.Printf("[DEBUG] response from update kepegawaian yayasan simpeg: %s\n", resBody)

}
