package model

import (
	"fmt"
	"svc-insani-go/app/helper"
)

type PegawaiCreateRequest struct {
	Nik                 string `form:"nik"`
	UuidKelompokPegawai string `form:"uuid_kelompok_pegawai"`
	UuidUnitKerja       string `form:"uuid_unit_kerja"`
	UuidBagianKerja     string `form:"uuid_bagian_kerja"`
	UuidLokasiKerja     string `form:"uuid_lokasi_kerja"`
	UuidPersonal        string `form:"uuid_personal"`
	UserUpdate          string `form:"-"`
	Uuid                string `form:"-"`
}

func (req PegawaiCreateRequest) Validate() error {
	switch {
	case len(req.Nik) != 9 || !helper.IsNumber(req.Nik):
		return fmt.Errorf("nik wajib diisi berupa 9 digit angka")
	case req.UuidKelompokPegawai == "":
		return fmt.Errorf("uuid_kelompok_pegawai tidak boleh kosong")
	case req.UuidUnitKerja == "":
		return fmt.Errorf("uuid_unit_kerja tidak boleh kosong")
	case req.UuidBagianKerja == "":
		return fmt.Errorf("uuid_bagian_kerja tidak boleh kosong")
	case req.UuidLokasiKerja == "":
		return fmt.Errorf("uuid_lokasi_kerja tidak boleh kosong")
	case req.UuidPersonal == "":
		return fmt.Errorf("uuid_personal tidak boleh kosong")
	default:
		return nil
	}
}
