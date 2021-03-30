package usecase

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"svc-insani-go/app"
	"svc-insani-go/app/database"
	"svc-insani-go/modules/v1/sk-pengangkatan/model"
	"testing"

	"github.com/labstack/echo"
)

func TestUsecase(t *testing.T) {
	db, err := database.Connect()
	if err != nil {
		t.Skipf("error db connect: %s\n", err.Error())
	}
	err = db.Ping()
	if err != nil {
		t.Skipf("error db ping: %s\n", err.Error())
	}
	loc := app.GetFixedTimeZone()
	a := app.App{DB: db, TimeLocation: loc}

	e := echo.New()
	t.Run("create_sk_pengangkatan_tendik", func(t *testing.T) {
		urlQuery := make(url.Values)
		urlQuery.Set("uuid_pegawai", "e37046f9-1437-11eb-a014-7eb0d4a3c7a0") // ?uuid_pegawai=e37046f9-1437-11eb-a014-7eb0d4a3c7a0

		formData := make(url.Values)
		formData.Set("nomor_sk", "001/HARIS/AHF.I/2020")
		formData.Set("tmt", "2020-01-03")
		formData.Set("uuid_kelompok_sk_pengangkatan", "46994cd3-ec0a-11ea-8c77-7eb0d4a3c7a0")
		formData.Set("uuid_unit_kerja", "05577f32-e996-11e9-8f20-506b8da96a87")
		formData.Set("uuid_jenis_sk", "ebc9e2c0-ee60-11ea-8c77-7eb0d4a3c7a0")
		formData.Set("uuid_jabatan_fungsional", "aeb51169-2fc6-11eb-a014-7eb0d4a3c7a0")
		formData.Set("uuid_status_pengangkatan", "47dd67dc-0479-11eb-8c77-7eb0d4a3c7a0")
		formData.Set("uuid_pangkat_golongan_pegawai", "c6101e45-09e3-11eb-8c77-7eb0d4a3c7a0")
		formData.Set("uuid_unit_pengangkat", "05577756-e996-11e9-8f20-506b8da96a87")
		// formData.Set("uuid_pegawai_penetapan", "") // skip dulu
		formData.Set("uuid_jenis_ijazah", "74d1a731-ee86-11ea-8c77-7eb0d4a3c7a0")
		formData.Set("masa_kerja_ril_tahun", "1")
		formData.Set("masa_kerja_ril_bulan", "12")
		formData.Set("masa_kerja_gaji_tahun", "")
		formData.Set("masa_kerja_gaji_bulan", "11")
		formData.Set("masa_kerja_diakui_tahun", "3")
		formData.Set("masa_kerja_diakui_bulan", "10")
		formData.Set("gaji_pokok", "2600000")
		formData.Set("tanggal_ditetapkan", "2020-02-02")

		req := httptest.NewRequest(
			http.MethodPost,
			"/?"+urlQuery.Encode(),
			strings.NewReader(formData.Encode()),
		)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
		req.Header.Set("X-Member", "haris")

		res := httptest.NewRecorder()
		ctx := e.NewContext(req, res)
		HandleCreateSKPengangkatanTendik(a)(ctx)

		fmt.Printf("code: %d\n", res.Code)
		fmt.Printf("body: %s\n", res.Body.String())
	})

	t.Run("get_sk_pengangkatan_tendik_detail", func(t *testing.T) {
		urlQuery := make(url.Values)
		urlQuery.Set("uuid_sk_pengangkatan_tendik", "fe1ca162-2ffe-11eb-a014-7eb0d4a3c7a0")

		req := httptest.NewRequest(
			http.MethodGet,
			"/?"+urlQuery.Encode(),
			nil,
		)
		req.Header.Set("Accept", "application/json")

		res := httptest.NewRecorder()
		ctx := e.NewContext(req, res)
		HandleGetDetailSKPengangkatanTendik(a)(ctx)

		var skPengangkatanTendik map[string][]*model.SKPengangkatanTendikDetail
		rawJSONResponse := []byte(fmt.Sprintf(`%s`, res.Body.String()))
		err := json.Unmarshal(rawJSONResponse, &skPengangkatanTendik)
		if err != nil {
			t.Fatalf("err marshaling: %s\n", err.Error())
		}
		b, err := json.MarshalIndent(&skPengangkatanTendik, "", "\t")
		if err != nil {
			t.Fatalf("err marshaling: %s\n", err.Error())
		}
		fmt.Printf("code: %d\n", res.Code)
		fmt.Printf("body:\n%s\n", b)
	})

}

func TestJoinString(t *testing.T) {
	ss := []string{"a", "asd"}
	fmt.Println("joined:", strings.Join(ss, ", "))
	ss = []string{}
	fmt.Println("joined:", strings.Join(ss, ", "))
}
