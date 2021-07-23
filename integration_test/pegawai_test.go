package integrationtest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"testing"
)

func Pegawai(t *testing.T, s *TestServer) func(t *testing.T) {
	return func(t *testing.T) {
		createPegawaiResp := make(map[string]interface{})
		t.Run("create_pegawai", CreatePegawai(t, s, createPegawaiResp))
		// fmt.Printf("[DEBUG] cp res: %+v\n", createPegawaiResp)
	}
}

const (
	falseNikMsg                    = "nik wajib diisi berupa 9 digit angka"
	emptyUuidKelompokPegawaiMsg    = "uuid_kelompok_pegawai tidak boleh kosong"
	notFoundUuidKelompokPegawaiMsg = "uuid_kelompok_pegawai tidak ditemukan"
	emptyUuidUnitKerjaMsg          = "uuid_unit_kerja tidak boleh kosong"
	notFoundUuidUnitKerjaMsg       = "uuid_unit_kerja tidak ditemukan"
	emptyUuidBagianKerjaMsg        = "uuid_bagian_kerja tidak boleh kosong"
	notFoundUuidBagianKerjaMsg     = "uuid_bagian_kerja tidak ditemukan"
	emptyUuidLokasiKerjaMsg        = "uuid_lokasi_kerja tidak boleh kosong"
	notFoundUuidLokasiKerjaMsg     = "uuid_lokasi_kerja tidak ditemukan"
)

func CreatePegawai(t *testing.T, s *TestServer, resp map[string]interface{}) func(*testing.T) {
	uuidPersonal := "7ab3f8d0-e6f7-433f-84c4-ef80270e6fca"

	groupsResp := make(map[string]interface{})
	t.Run("get_kelompok_pegawai", GetKelompokPegawai(t, s, groupsResp))
	groups := groupsResp["data"].([]interface{})

	unitsResp := make(map[string]interface{})
	t.Run("get_unit_kerja", GetUnitKerja(t, s, unitsResp))
	units := unitsResp["data"].([]interface{})

	divisionsResp := make(map[string]interface{})
	t.Run("get_bagian_kerja", GetBagianKerja(t, s, divisionsResp))
	divisions := divisionsResp["data"].([]interface{})

	locationsResp := make(map[string]interface{})
	t.Run("get_lokasi_kerja", GetLokasiKerja(t, s, locationsResp))
	locations := locationsResp["data"].([]interface{})

	testCases := []struct {
		title        string
		body         map[string]string
		checkResFunc func(code int, body map[string]interface{}) bool
		failMsg      func(caseTitle string, reqBody, resBody interface{}) string
	}{
		{
			title: "input_nik_empty_should_fail",
			body:  map[string]string{"nik": ""},
			checkResFunc: func(code int, body map[string]interface{}) bool {
				return code == http.StatusBadRequest && body["message"] == falseNikMsg
			},
			failMsg: func(caseTitle string, reqBody, resBody interface{}) string {
				return fmt.Sprintf(`%s, input: %+v, want: status should be bad request and return "%s", got: %+v`, caseTitle, reqBody, falseNikMsg, resBody)
			},
		},
		{
			title: "input_nik_12312312_should_fail",
			body:  map[string]string{"nik": "12312312"},
			checkResFunc: func(code int, body map[string]interface{}) bool {
				return code == http.StatusBadRequest && body["message"] == falseNikMsg
			},
			failMsg: func(caseTitle string, reqBody, resBody interface{}) string {
				return fmt.Sprintf(`%s, input: %+v, want: status should be bad request and return "%s", got: %+v`, caseTitle, reqBody, falseNikMsg, resBody)
			},
		},
		{
			title: "input_nik_-1_should_fail",
			body:  map[string]string{"nik": "-1"},
			checkResFunc: func(code int, body map[string]interface{}) bool {
				return code == http.StatusBadRequest && body["message"] == falseNikMsg
			},
			failMsg: func(caseTitle string, reqBody, resBody interface{}) string {
				return fmt.Sprintf(`%s, input: %+v, want: status should be bad request and return "%s", got: %+v`, caseTitle, reqBody, falseNikMsg, resBody)
			},
		},
		{
			title: "input_nik_abc_should_fail",
			body:  map[string]string{"nik": "abc"},
			checkResFunc: func(code int, body map[string]interface{}) bool {
				return code == http.StatusBadRequest && body["message"] == falseNikMsg
			},
			failMsg: func(caseTitle string, reqBody, resBody interface{}) string {
				return fmt.Sprintf(`%s, input: %+v, want: status should be bad request and return "%s", got: %+v`, caseTitle, reqBody, falseNikMsg, resBody)
			},
		},
		{
			title: "empty_uuid_kelompok_pegawai_should_fail",
			body:  map[string]string{"nik": "200000310"},
			checkResFunc: func(code int, body map[string]interface{}) bool {
				return code == http.StatusBadRequest && body["message"] == emptyUuidKelompokPegawaiMsg
			},
			failMsg: func(caseTitle string, reqBody, resBody interface{}) string {
				return fmt.Sprintf(`%s, input: %+v, want: status should be bad request and return "%s", got: %+v`, caseTitle, reqBody, emptyUuidKelompokPegawaiMsg, resBody)
			},
		},
		{
			title: "empty_uuid_unit_kerja_should_fail",
			body: map[string]string{
				"nik":                   "200000310",
				"uuid_kelompok_pegawai": "abc",
				"uuid_unit_kerja":       "",
			},
			checkResFunc: func(code int, body map[string]interface{}) bool {
				return code == http.StatusBadRequest && body["message"] == emptyUuidUnitKerjaMsg
			},
			failMsg: func(caseTitle string, reqBody, resBody interface{}) string {
				return fmt.Sprintf(`%s, input: %+v, want: status should be bad request and return "%s", got: %+v`, caseTitle, reqBody, emptyUuidUnitKerjaMsg, resBody)
			},
		},
		{
			title: "empty_uuid_bagian_kerja_should_fail",
			body: map[string]string{
				"nik":                   "200000310",
				"uuid_kelompok_pegawai": "abc",
				"uuid_unit_kerja":       "x",
				"uuid_bagian_kerja":     "",
			},
			checkResFunc: func(code int, body map[string]interface{}) bool {
				return code == http.StatusBadRequest && body["message"] == emptyUuidBagianKerjaMsg
			},
			failMsg: func(caseTitle string, reqBody, resBody interface{}) string {
				return fmt.Sprintf(`%s, input: %+v, want: status should be bad request and return "%s", got: %+v`, caseTitle, reqBody, emptyUuidBagianKerjaMsg, resBody)
			},
		},
		{
			title: "empty_uuid_lokasi_kerja_should_fail",
			body: map[string]string{
				"nik":                   "200000310",
				"uuid_kelompok_pegawai": "abc",
				"uuid_unit_kerja":       "x",
				"uuid_bagian_kerja":     "x",
				"uuid_lokasi_kerja":     "",
			},
			checkResFunc: func(code int, body map[string]interface{}) bool {
				return code == http.StatusBadRequest && body["message"] == emptyUuidLokasiKerjaMsg
			},
			failMsg: func(caseTitle string, reqBody, resBody interface{}) string {
				return fmt.Sprintf(`%s, input: %+v, want: status should be bad request and return "%s", got: %+v`, caseTitle, reqBody, emptyUuidLokasiKerjaMsg, resBody)
			},
		},
		{
			title: "random_uuid_kelompok_pegawai_should_fail",
			body: map[string]string{
				"nik":                   "200000310",
				"uuid_kelompok_pegawai": "abc",
				"uuid_unit_kerja":       "x",
				"uuid_bagian_kerja":     "x",
				"uuid_lokasi_kerja":     "x",
			},
			checkResFunc: func(code int, body map[string]interface{}) bool {
				return code == http.StatusBadRequest && body["message"] == notFoundUuidKelompokPegawaiMsg
			},
			failMsg: func(caseTitle string, reqBody, resBody interface{}) string {
				return fmt.Sprintf(`%s, input: %+v, want: status should be bad request and return "%s", got: %+v`, caseTitle, reqBody, notFoundUuidKelompokPegawaiMsg, resBody)
			},
		},
		{
			title: "random_uuid_unit_kerja_should_fail",
			body: map[string]string{
				"nik":                   "200000310",
				"uuid_kelompok_pegawai": groups[0].(map[string]interface{})["uuid"].(string),
				"uuid_unit_kerja":       "x",
				"uuid_bagian_kerja":     "x",
				"uuid_lokasi_kerja":     "x",
			},
			checkResFunc: func(code int, body map[string]interface{}) bool {
				return code == http.StatusBadRequest && body["message"] == notFoundUuidUnitKerjaMsg
			},
			failMsg: func(caseTitle string, reqBody, resBody interface{}) string {
				return fmt.Sprintf(`%s, input: %+v, want: status should be bad request and return "%s", got: %+v`, caseTitle, reqBody, notFoundUuidUnitKerjaMsg, resBody)
			},
		},
		{
			title: "random_uuid_bagian_kerja_should_fail",
			body: map[string]string{
				"nik":                   "200000310",
				"uuid_kelompok_pegawai": groups[0].(map[string]interface{})["uuid"].(string),
				"uuid_unit_kerja":       units[0].(map[string]interface{})["uuid"].(string),
				"uuid_bagian_kerja":     "x",
				"uuid_lokasi_kerja":     "x",
			},
			checkResFunc: func(code int, body map[string]interface{}) bool {
				return code == http.StatusBadRequest && body["message"] == notFoundUuidBagianKerjaMsg
			},
			failMsg: func(caseTitle string, reqBody, resBody interface{}) string {
				return fmt.Sprintf(`%s, input: %+v, want: status should be bad request and return "%s", got: %+v`, caseTitle, reqBody, notFoundUuidBagianKerjaMsg, resBody)
			},
		},
		{
			title: "random_uuid_lokasi_kerja_should_fail",
			body: map[string]string{
				"nik":                   "200000310",
				"uuid_kelompok_pegawai": groups[0].(map[string]interface{})["uuid"].(string),
				"uuid_unit_kerja":       units[0].(map[string]interface{})["uuid"].(string),
				"uuid_bagian_kerja":     divisions[0].(map[string]interface{})["uuid"].(string),
				"uuid_lokasi_kerja":     "x",
			},
			checkResFunc: func(code int, body map[string]interface{}) bool {
				fmt.Printf("[DEBUG] g: %+v\n", groups[1].(map[string]interface{}))
				return code == http.StatusBadRequest && body["message"] == notFoundUuidLokasiKerjaMsg
			},
			failMsg: func(caseTitle string, reqBody, resBody interface{}) string {
				return fmt.Sprintf(`%s, input: %+v, want: status should be bad request and return "%s", got: %+v`, caseTitle, reqBody, notFoundUuidLokasiKerjaMsg, resBody)
			},
		},
		{
			title: "success_created",
			body: map[string]string{
				"nik":                   "200000311",
				"uuid_kelompok_pegawai": groups[0].(map[string]interface{})["uuid"].(string),
				"uuid_unit_kerja":       units[0].(map[string]interface{})["uuid"].(string),
				"uuid_bagian_kerja":     divisions[0].(map[string]interface{})["uuid"].(string),
				"uuid_lokasi_kerja":     locations[0].(map[string]interface{})["uuid"].(string),
				"uuid_personal":         "0e3dbf08-d59a-4a5c-a55b-6b15a7799642",
			},
			checkResFunc: func(code int, body map[string]interface{}) bool {
				return code == http.StatusOK // && body["message"] == notFoundUuidLokasiKerjaMsg
			},
			failMsg: func(caseTitle string, reqBody, resBody interface{}) string {
				return fmt.Sprintf(`%s, input: %+v, want: status should be ok, got: %+v`, caseTitle, reqBody, resBody)
			},
		},
	}

	return func(t *testing.T) {
		for _, tc := range testCases {
			formDataBuffer := &bytes.Buffer{}
			formDataWriter := multipart.NewWriter(formDataBuffer)
			err := FillFormDataFieldMap(formDataWriter, tc.body)
			if err != nil {
				t.Fatal(err)
			}
			formDataWriter.Close()
			targetUrl := fmt.Sprintf("%s/%s/%s/%s", s.Server.URL, v1path, "pegawai-simpeg", uuidPersonal)
			req, err := http.NewRequest(http.MethodPost, targetUrl, bytes.NewReader(formDataBuffer.Bytes()))
			req.Header.Set("Content-Type", formDataWriter.FormDataContentType())
			req.Header.Set("X-Member", "testadmin")
			res, err := s.Client.Client.Do(req)
			if err != nil {
				t.Fatal(err)
			}

			rawResBodyJSON, err := ioutil.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				t.Fatal(err)
			}
			resp = make(map[string]interface{})
			err = json.Unmarshal(rawResBodyJSON, &resp)
			if err != nil {
				t.Fatal(err)
			}

			if !tc.checkResFunc(res.StatusCode, resp) {
				t.Fatal(tc.failMsg(tc.title, tc.body, resp))
			}
		}
	}
}
