package model

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
	KdHomebasePddikti   string `json:"kd_homebase_pddikti"`
	UuidHomebasePddikti string `json:"uuid_homebase_pddikti"`
	KdHomebaseUii       string `json:"kd_homebase_uii"`
	UuidHomebaseUii     string `json:"uuid_homebase_uii"`
}
