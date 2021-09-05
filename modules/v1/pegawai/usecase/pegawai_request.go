package usecase

import (
	"fmt"
	"svc-insani-go/app/helper"
)

type PegawaiCreateRequest struct {
	UuidKelompokPegawai string `form:"uuid_kelompok_pegawai"`
	Nik                 string `form:"nik"`
	UuidUnitKerja       string `form:"uuid_unit_kerja"`
	TmtSkPertama        string `form:"tmt_sk_pertama"`
	UuidLokasiKerja     string `form:"uuid_lokasi_kerja"`
	UuidPersonal        string `form:"uuid_personal"`
}

func (req PegawaiCreateRequest) Validate() error {
	switch {
	case len(req.Nik) != 9 || !helper.IsNumber(req.Nik):
		return fmt.Errorf("nik wajib diisi berupa 9 digit angka")
	case req.UuidKelompokPegawai == "":
		return fmt.Errorf("uuid_kelompok_pegawai tidak boleh kosong")
	case req.UuidUnitKerja == "":
		return fmt.Errorf("uuid_unit_kerja tidak boleh kosong")
	case !helper.IsDateFormatValid("2006-01-02", req.TmtSkPertama):
		return fmt.Errorf("tmt_sk_pertama wajib diisi dengan format yyyy-mm-dd")
	case req.UuidLokasiKerja == "":
		return fmt.Errorf("uuid_lokasi_kerja tidak boleh kosong")
	default:
		return nil
	}
}
