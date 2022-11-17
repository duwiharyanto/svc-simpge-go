package model

import "encoding/json"

type UnitKerjaPegawai struct {
	UuidIndukKerja      string `json:"uuid_induk_kerja"`
	KdIndukKerja        string `json:"kd_induk_kerja"`
	IndukKerja          string `json:"induk_kerja"`
	UuidUnitKerja       string `json:"uuid_unit_kerja"`
	KdUnitKerja         string `json:"kd_unit_kerja"`
	UnitKerja           string `json:"unit_kerja"`
	UuidBagianKerja     string `json:"uuid_bagian_kerja"`
	KdBagianKerja       string `json:"kd_bagian_kerja"`
	BagianKerja         string `json:"bagian_kerja"`
	UuidLokasiKerja     string `json:"uuid_lokasi_kerja"`
	LokasiKerja         string `json:"kd_lokasi_kerja"`
	LokasiDesc          string `json:"lokasi_kerja"`
	NoSkPertama         string `json:"nomor_sk_pertama_unit_kerja"`
	TmtSkPertama        string `json:"tmt_sk_pertama_unit_kerja"`
	TmtSkPertamaIdn     string `json:"tmt_sk_pertama_unit_kerja_idn"`
	NoSk                string `json:"nomor_sk"`
	TmtSk               string `json:"tmt_sk"`
	TglSk               string `json:"tgl_sk"`
	TmtAwalKontrak      string `json:"tmt_awal_kontrak"`
	TmtAkhirKontrak     string `json:"tmt_akhir_kontrak"`
	KdHomebasePddikti   string `json:"kd_homebase_pddikti"`
	UuidHomebasePddikti string `json:"uuid_homebase_pddikti"`
	KdHomebaseUii       string `json:"kd_homebase_uii"`
	UuidHomebaseUii     string `json:"uuid_homebase_uii"`
}

func (u *UnitKerjaPegawai) IsOnlyTwoCharacter(property string) bool {
	var x map[string]interface{}
	m, _ := json.Marshal(u)
	json.Unmarshal(m, &x)
	return len(x[property].(string)) == 2
}

func (u *UnitKerjaPegawai) AddingNewCharacter(kdUnit string) string {
	return kdUnit + "0"
}
