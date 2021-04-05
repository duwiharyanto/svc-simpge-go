package model

import (
	pegawai "svc-insani-go/modules/v1/pegawai/model"
	skPengangkatanModel "svc-insani-go/modules/v1/sk-pengangkatan/model"
	// skPrajabatanModel "svc-decree-go/modules/v1/sk-prajabatan/model"
)

const (
	KdJenisSkPengangkatan = "1"
)

type SKPegawai struct {
	IDPegawai            string                                   `json:"-"`
	ID                   string                                   `json:"-"`
	NIK                  string                                   `json:"nik"`
	Nama                 string                                   `json:"nama"`
	GelarDepan           string                                   `json:"gelar_depan"`
	GelarBelakang        string                                   `json:"gelar_belakang"`
	IDJenisSK            string                                   `json:"id_jenis_sk"`
	NomorSK              string                                   `json:"nomor_sk"`
	TentangSK            string                                   `json:"tentang_sk"`
	TMT                  string                                   `json:"tmt"`
	TMTIDN               string                                   `json:"tmt_idn"`
	UUID                 string                                   `json:"uuid"`
	UserInput            string                                   `json:"user_input"`
	UserUpdate           string                                   `json:"user_update"`
	SKPengangkatanDosen  skPengangkatanModel.SKPengangkatanDosen  `json:"sk_pengangkatan_dosen,omitempty"`
	SKPengangkatanTendik skPengangkatanModel.SKPengangkatanTendik `json:"sk_pengangkatan_tendik,omitempty"`
	Pegawai              *pegawai.Pegawai                         `json:"-"`

	// SKPengangkatanPegawai skPengangkatanModel.SKPengangkatanPegawai `json:"sk_pengangkatan"`
	// SKPrajabatan skPrajabatanModel.SKPrajabatan `json:"sk_prajabatan"`
}

// func (sk *SKPegawai) SetSKPengangkatan(skPengangkatan skPengangkatanModel.SKPengangkatanPegawai) {
// 	sk.SKPengangkatanPegawai = skPengangkatan
// }
// func (sk *SKPegawai) SetSKPengangkatanDosen(skPengangkatanDosen skPengangkatanModel.SKPengangkatanDosen) {
// 	sk.SKPengangkatanDosen = skPengangkatanDosen
// }
func (sk *SKPegawai) SetSKPengangkatanTendik(skPengangkatanTendik skPengangkatanModel.SKPengangkatanTendik) {
	sk.SKPengangkatanTendik = skPengangkatanTendik
}

type SKPegawaiResponse struct {
	Data []SKPegawai `json:"data"`
}
