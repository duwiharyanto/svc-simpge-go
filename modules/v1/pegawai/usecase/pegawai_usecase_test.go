package usecase

import (
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/app/database"
	"svc-insani-go/app/minio"
	"testing"
)

func TestNewPegawaiOra(t *testing.T) {
	db, err := database.Connect()
	if err != nil {
		t.Skip("failed connect db:", err)
	}

	gormDb, err := database.InitGorm(db, true)
	if err != nil {
		t.Fatal(err)
	}
	mc, err := minio.Connect()
	if err != nil {
		t.Fatal(err)
	}

	timeLocation := app.GetFixedTimeZone()
	a := &app.App{DB: db, GormDB: gormDb, TimeLocation: timeLocation, MinioClient: mc, MinioBucketName: "insani"}

	uuid := "872e2c5a-eb92-11eb-8820-000c2977b907"
	pegawaiDetail, err := PrepareGetSimpegPegawaiByUUID(a, uuid)
	if err != nil {
		t.Fatal(err)
	}
	pegawaiOra := newPegawaiOra(&pegawaiDetail)
	if pegawaiOra == nil {
		t.Fatal("Should not be nil")
	}
	fmt.Printf("[DEBUG] pgw ora: %+v\n", pegawaiOra)

}

// func FillFormDataFieldMap(w *multipart.Writer, m map[string]string) error {
// 	for k, v := range m {
// 		formField, err := w.CreateFormField(k)
// 		if err != nil {
// 			return fmt.Errorf("failed create field %s: %w", k, err)
// 		}
// 		_, err = io.Copy(formField, strings.NewReader(v))
// 		if err != nil {
// 			return fmt.Errorf("failed copy %s value: %w", k, err)
// 		}
// 	}

// 	return nil
// }

// func FillFormDataField(w *multipart.Writer, formField io.Writer, key, value string) (io.Writer, error) {
// 	formField, err := w.CreateFormField(key)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed create field %s: %w", key, err)
// 	}

// 	_, err = io.Copy(formField, strings.NewReader(value))
// 	if err != nil {
// 		return nil, fmt.Errorf("failed copy %s value: %w", key, err)
// 	}

// 	return formField, nil
// }

// func TestHandleUpdateSimpeg(t *testing.T) {
// 	e := echo.New()
// 	e.Use(router.SetResponseTimeout(context.Background()))

// 	db, err := database.Connect()
// 	if err != nil {
// 		t.Skip("failed connect db:", err)
// 	}

// 	gormDb, err := database.InitGorm(db, true)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	mc, err := minio.Connect()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	a := &app.App{DB: db, GormDB: gormDb, MinioClient: mc, MinioBucketName: "insani"}

// 	appCtx := context.Background()

// 	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
// 		AllowOrigins: []string{"*"},
// 		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
// 	}))
// 	slackErrChan := app.NewSlackLogger(appCtx, a.HttpClient)

// 	router.InitRoute(a, appCtx, e, slackErrChan)

// 	server := httptest.NewServer(e)
// 	defer server.Close()

// 	// init form data
// 	wbuf := &bytes.Buffer{}
// 	wr := multipart.NewWriter(wbuf)
// 	err = FillFormDataFieldMap(wr, map[string]string{
// 		//"uuid_jenis_pegawai":,
// 		//"uuid_status_egawai":,
// 		//"uuid_kelompok_pegawai":,
// 		//"uuid_golongan":,
// 		//"uuid_ruang":,
// 		//"uuid_induk_kerja":,
// 		//"uuid_unit_kerja":,
// 		//"uuid_bagian_kerja":7f9952ef-1fd7-11eb-a014-7eb0d4a3c7a0,
// 		//"uuid_lokasi_unit_kerja":,
// 		//"uuid_pangkat_golongan":,
// 		//"uuid_jabatan_fungsional":,
// 		//"tmt_pangkat_golongan":2020-01-10,
// 		//"tmt_jabatan":2020-01-10,
// 		"masa_kerja_bawaan_tahun": "13",
// 		"masa_kerja_bawaan_bulan": "5",
// 		//"masa_kerja_gaji_tahun":,
// 		//"masa_kerja_gaji_bulan":,
// 		//"masa_kerja_total_tahun":,
// 		//"masa_kerja_total_bulan":,
// 		//"angka_kredit":221,
// 		//"nomor_sertifikasi":,
// 		"uuid_jenis_nomor_registrasi": "2c10a574-7594-11eb-9f3c-7eb0d4a3c7a0",
// 		//"nomor_registrasi":,
// 		//"nomor_sk_pertama":,
// 		"tmt_sk_pertama": "2020-01-10",
// 		//"uuid_status_pegawai_aktif":,
// 		//"nip_pns":,
// 		//"no_kartu_pegawai":,
// 		//"uuid_pangkat_gol_ruang_pns":,
// 		"tmt_pangkat_gol_ruang_pns": "2020-01-10",
// 		//"uuid_jabatan_pns":,
// 		"tmt_jabatan_pns": "2020-01-10",
// 		//"masa_kerja_pns_tahun":,
// 		//"masa_kerja_pns_bulan":,
// 		//"angka_kredit_pns":,
// 		//"uuid_jenis_ptt":,
// 		"instansi_asal_ptt": "Instansi Asal Test",
// 		"keterangan_pns":    "Ini Keterangan Test",
// 		//"uuid_tingkat_pdd_diakui":"e63bfcea-eb1d-44b7-8df1-4e84afea1714",
// 		"uuid_tingkat_pdd_terakhir": "e63bfcea-eb1d-44b7-8df1-4e84afea1714",
// 	})

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	wr.Close()

// 	// create request

// 	uuidPegawai := "e5762619-1437-11eb-a014-7eb0d4a3c7a0"
// 	baseURL := server.URL + "/public/api/v1/pegawai/" + uuidPegawai
// 	// fmt.Printf("[DEBUG] base url: %s\n", baseURL)
// 	req, err := http.NewRequest(http.MethodPut, baseURL, bytes.NewReader((wbuf.Bytes())))
// 	req.Header.Set("Content-Type", wr.FormDataContentType())
// 	req.Header.Set("X-Member", "admin 4")

// 	// send http request
// 	client := http.DefaultClient
// 	res, err := client.Do(req)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// read response body
// 	rawResBodyJSON, err := ioutil.ReadAll(res.Body)
// 	res.Body.Close()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// format res body indentation
// 	var buf bytes.Buffer
// 	json.Indent(&buf, rawResBodyJSON, "", "\t")
// 	fmt.Printf("[DEBUG] rec body: %s\n", buf.String())

// 	var resBody map[string]interface{}
// 	err = json.Unmarshal(rawResBodyJSON, &resBody)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// if msg, exist := resBody["message"]; !exist || !strings.Contains(strings.ToLower(msg.(string)), "berhasil") {
// 	// 	fmt.Printf("[DEBUG] name: %+v\n", resBody)
// 	// 	t.Fatal("should return success message")
// 	// }
// }

// func TestHandleCreateSimpeg(t *testing.T) {
// 	e := echo.New()
// 	e.Use(router.SetResponseTimeout(context.Background()))

// 	db, err := database.Connect()
// 	if err != nil {
// 		t.Skip("failed connect db:", err)
// 	}

// 	gormDb, err := database.InitGorm(db, true)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	mc, err := minio.Connect()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	a := &app.App{DB: db, GormDB: gormDb, MinioClient: mc, MinioBucketName: "insani"}

// 	appCtx := context.Background()

// 	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
// 		AllowOrigins: []string{"*"},
// 		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
// 	}))
// 	slackErrChan := app.NewSlackLogger(appCtx, a.HttpClient)

// 	router.InitRoute(a, appCtx, e, slackErrChan)

// 	server := httptest.NewServer(e)
// 	defer server.Close()

// 	// init form data
// 	wbuf := &bytes.Buffer{}
// 	wr := multipart.NewWriter(wbuf)
// 	err = FillFormDataFieldMap(wr, map[string]string{
// 		"uuid_jenis_pegawai":          "06e83088-0467-11eb-8c77-7eb0d4a3c7a0",
// 		"uuid_status_pegawai":         "aa14d46c-0871-11eb-8c77-7eb0d4a3c7a0",
// 		"uuid_kelompok_pegawai":       "741fc23c-4fea-11eb-bf95-a74048ab8082",
// 		"uuid_golongan":               "c6101e45-09e3-11eb-8c77-7eb0d4a3c7a0",
// 		"tmt_pangkat_golongan":        "2020-01-01",
// 		"uuid_jabatan_fungsional":     "aeb50dc8-2fc6-11eb-a014-7eb0d4a3c7a0",
// 		"tmt_jabatan":                 "2020-01-01",
// 		"masa_kerja_bawaan_tahun":     "1",
// 		"masa_kerja_bawaan_bulan":     "11",
// 		"masa_kerja_gaji_tahun":       "1",
// 		"masa_kerja_gaji_bulan":       "2",
// 		"masa_kerja_total_tahun":      "1",
// 		"masa_kerja_total_bulan":      "2",
// 		"angka_kredit":                "22",
// 		"nomor_sertifikasi":           "2322",
// 		"uuid_jenis_nomor_registrasi": "2c10a574-7594-11eb-9f3c-7eb0d4a3c7a0",
// 		"nomor_registrasi":            "1122334455",
// 		"uuid_induk_kerja":            "fb92f553-1fd2-11eb-a014-7eb0d4a3c7a0",
// 		"uuid_unit_kerja":             "798c791e-1fd3-11eb-a014-7eb0d4a3c7a0",
// 		"uuid_bagian_kerja":           "7f994497-1fd7-11eb-a014-7eb0d4a3c7a0",
// 		"uuid_lokasi_kerja":           "a818aed6-4fff-11eb-bf95-a74048ab8082",
// 		"nomor_sk_pertama":            "112233",
// 		"tmt_sk_pertama":              "2020-10-10",
// 		"uuid_homebase_pddikti":       "798c791e-1fd3-11eb-a014-7eb0d4a3c7a0",
// 		"uuid_homebase_uii":           "798c791e-1fd3-11eb-a014-7eb0d4a3c7a0",
// 		"uuid_jenis_ptt":              "a3e6f4b8-75a4-11eb-9f3c-7eb0d4a3c7a0",
// 		"instansi_asal_ptt":           "Instansi Asal Test",
// 		"no_kartu_pegawai":            "0102030405",
// 		"uuid_pangkat_gol_ruang_pns":  "c6101e45-09e3-11eb-8c77-7eb0d4a3c7a0",
// 		"tmt_pangkat_gol_ruang_pns":   "2020-01-01",
// 		"uuid_jabatan_pns":            "aeb50dc8-2fc6-11eb-a014-7eb0d4a3c7a0",
// 		"tmt_jabatan_pns":             "2020-01-11",
// 		"masa_kerja_pns_tahun":        "1",
// 		"masa_kerja_pns_bulan":        "2",
// 		"angka_kredit_pns":            "443",
// 		"keterangan_pns":              "Keterangan Ok Lagi Lagi",
// 		"uuid_status_pegawai_aktif":   "055dd61b-75a4-11eb-9f3c-7eb0d4a3c7a0",
// 	})

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	wr.Close()

// 	// create request

// 	nikPegawai := "201005104"
// 	baseURL := server.URL + "/public/api/v1/pegawai-simpeg/" + nikPegawai
// 	req, err := http.NewRequest(http.MethodPut, baseURL, bytes.NewReader((wbuf.Bytes())))
// 	req.Header.Set("Content-Type", wr.FormDataContentType())
// 	req.Header.Set("X-Member", "admin 4")

// 	// send http request
// 	client := http.DefaultClient
// 	res, err := client.Do(req)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// read response body
// 	rawResBodyJSON, err := ioutil.ReadAll(res.Body)
// 	res.Body.Close()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// format res body indentation
// 	var buf bytes.Buffer
// 	json.Indent(&buf, rawResBodyJSON, "", "\t")
// 	fmt.Printf("[DEBUG] rec body: %s\n", buf.String())

// 	var resBody map[string]interface{}
// 	err = json.Unmarshal(rawResBodyJSON, &resBody)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// if msg, exist := resBody["message"]; !exist || !strings.Contains(strings.ToLower(msg.(string)), "berhasil") {
// 	// 	fmt.Printf("[DEBUG] name: %+v\n", resBody)
// 	// 	t.Fatal("should return success message")
// 	// }
// }
