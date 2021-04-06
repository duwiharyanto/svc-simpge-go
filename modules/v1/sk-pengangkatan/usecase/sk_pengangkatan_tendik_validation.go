package usecase

import (
	"fmt"
	"strings"
	"svc-insani-go/app"

	kelompokSKPengangkatanRepo "svc-insani-go/modules/v1/master-kelompok-pegawai/repo"

	jabatanFungsionalRepo "svc-insani-go/modules/v1/master-jabatan-fungsional/repo"
	pangkatPegawaiRepo "svc-insani-go/modules/v1/master-pangkat-golongan-pegawai/repo"
	unitKerjaRepo "svc-insani-go/modules/v1/master-unit-kerja/repo"
	pegawaiRepo "svc-insani-go/modules/v1/pegawai/repo"
	skPegawaiModel "svc-insani-go/modules/v1/sk-pegawai/model"
	"svc-insani-go/modules/v1/sk-pengangkatan/model"
	skRepo "svc-insani-go/modules/v1/sk/repo"
	statusPengangkatanRepo "svc-insani-go/modules/v1/sk/repo"

	"github.com/labstack/echo"
)

const ErrEchoEmptyRequestBody = "Request body can't be empty"

func ValidateCreateSKPengangkatanTendik(a app.App, c echo.Context) (skPegawai skPegawaiModel.SKPegawai, err error) {
	uuidPegawai := c.QueryParam("uuid_pegawai")
	if uuidPegawai == "" {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("uuid pegawai tidak boleh kosong")
	}

	skPengangkatanTendik := new(model.SKPengangkatanTendik)
	err = c.Bind(skPengangkatanTendik)
	if err != nil {
		if strings.Contains(err.Error(), ErrEchoEmptyRequestBody) {
			return skPegawaiModel.SKPegawai{}, fmt.Errorf("Input kosong, tidak ada data yang berubah")
		}
	}
	skPengangkatanTendik.FileSKTendik, _ = c.FormFile("file_sk")

	pegawai, err := pegawaiRepo.GetPegawaiByUUID(a, uuidPegawai)
	if err != nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("error from repo get pegawai by uuid, %w", err)
	}
	if pegawai == nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("pegawai tidak ditemukan")
	}

	kelompokSKPengangkatan, err := kelompokSKPengangkatanRepo.GetKelompokPegawaiByUUID(a, skPengangkatanTendik.UUIDKelompokSKPengangkatan)
	if err != nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("error from repo get kelompok sk pengangkatan by uuid, %w", err)
	}
	if kelompokSKPengangkatan == nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("kelompok sk pengangkatan tidak ditemukan")
	}

	unitPegawai, err := unitKerjaRepo.GetUnit2ByUUID(a, skPengangkatanTendik.UUIDUnitPegawai)
	if err != nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("error from repo get unit kerja by uuid, %w", err)
	}
	if unitPegawai == nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("unit kerja tidak ditemukan")
	}

	jabatanFungsional, err := jabatanFungsionalRepo.GetJabatanFungsionalByUUID(a, skPengangkatanTendik.UUIDJabatanFungsional)
	if err != nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("error from repo get jabatan fungsional by uuid, %w", err)
	}
	if jabatanFungsional == nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("jabatan fungsional tidak ditemukan")
	}

	jenisIjazah, err := skRepo.GetJenisIjazahByUUID(a, skPengangkatanTendik.UUIDJenisIjazah)
	if err != nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("error from repo get jenis ijazah by uuid %w", err)
	}
	if jenisIjazah == nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("jenis ijazah tidak ditemukan")
	}

	jenissk, err := skRepo.GetJenisSKByCode(a, skPegawaiModel.KdJenisSkPengangkatan)
	if err != nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("error from repo get jenis sk by code, %w", err)
	}
	if jenissk == nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("jenis sk tidak ditemukan")
	}

	pangkatGolonganPegawai, err := pangkatPegawaiRepo.GetPangkatGolonganPegawaiByUUID(a, skPengangkatanTendik.UUIDPangkatGolonganPegawai)
	if err != nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("error from repo get pangkat golongan pegawai, %w", err)
	}
	if pangkatGolonganPegawai == nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("pangkat golongan pegawai tidak ditemukan")
	}

	statusPengangkatan, err := statusPengangkatanRepo.GetStatusPengangkatanByUUID(a, skPengangkatanTendik.UUIDStatusPengangkatan)
	if err != nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("error from repo get status pengangkatan, %w", err)
	}
	if statusPengangkatan == nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("status pengangkatan tidak ditemukan")
	}

	unitPengangkat, err := unitKerjaRepo.GetUnit2ByUUID(a, skPengangkatanTendik.UUIDUnitPengangkat)
	if err != nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("error from unit pengangkat pegawai, %w", err)
	}
	if unitPengangkat == nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("unit pengangkat pegawai tidak ditemukan")
	}

	// pegawai penetap skip dulu, datanya gak jelas
	// pegawaiPenetap, err := pegawaiPenetapRepo.GetPegawaiPenetapSKByUUID(a, skPengangkatanTendik.UUIDPegawaiPenetap)
	// if err != nil {
	// 	return skPegawaiModel.SKPegawai{}, fmt.Errorf("error from unit pengangkat pegawai pegawai %w", err)
	// }

	// if pegawaiPenetap == nil {
	// 	return skPegawaiModel.SKPegawai{}, fmt.Errorf("pegawai penetap tidak ditemukan")
	// }

	skPengangkatanTendik.UserUpdate = c.Request().Header.Get("X-Member")
	skPengangkatanTendik.UserInput = c.Request().Header.Get("X-Member")
	skPengangkatanTendik.IDKelompokSKPengangkatan = kelompokSKPengangkatan.ID
	skPengangkatanTendik.IDUnitPengangkat = unitPengangkat.ID
	skPengangkatanTendik.IDJabatanFungsional = jabatanFungsional.ID
	skPengangkatanTendik.IDJenisIjazah = jenisIjazah.ID
	skPengangkatanTendik.IDUnitPegawai = unitPegawai.ID
	skPengangkatanTendik.IDJenisSK = jenissk.ID
	skPengangkatanTendik.IDPangkatGolonganPegawai = pangkatGolonganPegawai.ID
	skPengangkatanTendik.IDStatusPengangkatan = statusPengangkatan.ID
	// skPengangkatanTendik.IDPegawaiPenetap = pegawaiPenetap.ID // skip dulu, datanya belum jelas

	skPegawai = skPegawaiModel.SKPegawai{}
	skPegawai.SKPengangkatanTendik = *skPengangkatanTendik
	skPegawai.NomorSK = c.FormValue("nomor_sk")
	skPegawai.IDJenisSK = jenissk.ID
	// skPegawai.TentangSK = c.FormValue("tentang_sk") // gak ada di desain
	skPegawai.TMT = c.FormValue("tmt")
	skPegawai.UserUpdate = c.Request().Header.Get("X-Member")
	skPegawai.UserInput = c.Request().Header.Get("X-Member")
	skPegawai.IDPegawai = pegawai.ID
	skPegawai.Pegawai = pegawai
	return skPegawai, nil

}
