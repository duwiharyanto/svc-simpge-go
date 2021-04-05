package usecase

import (
	"encoding/json"
	"fmt"
	"strings"
	"svc-insani-go/app"

	//jenisskPengangkatanRepo "svc-insani-go/mdoules/v1/master-kelompok_pegawai/repo"

	indukKerjaRepo "svc-insani-go/modules/v1/master-induk-kerja/repo"
	jabatanFungsionalRepo "svc-insani-go/modules/v1/master-jabatan-fungsional/repo"
	pangkatGolonganPegawaiRepo "svc-insani-go/modules/v1/master-pangkat-golongan-pegawai/repo"
	unitKerjaRepo "svc-insani-go/modules/v1/master-unit-kerja/repo" //cek
	pegawaiRepo "svc-insani-go/modules/v1/pegawai/repo"
	skPegawaiModel "svc-insani-go/modules/v1/sk-pegawai/model"
	"svc-insani-go/modules/v1/sk-pengangkatan/model"
	skRepo "svc-insani-go/modules/v1/sk/repo" //cek

	"github.com/labstack/echo"
)

func ValidateCreateSKPengangkatanDosen(a app.App, c echo.Context) (skPegawai skPegawaiModel.SKPegawai, err error) {
	uuidPegawai := c.QueryParam("uuid_pegawai")

	skPengangkatanDosen := &model.SKPengangkatanDosen{}
	err = c.Bind(skPengangkatanDosen)
	if err != nil {
		if strings.Contains(err.Error(), ErrEchoEmptyRequestBody) {
			return skPegawaiModel.SKPegawai{}, fmt.Errorf("Input kosong, tidak ada data yang berubah")
		} else {
			fmt.Printf("[DEBUG] error echo: %s\n", err.Error())
		}
	}

	skPengangkatanDosen.FileSKDosen, _ = c.FormFile("file_sk")

	if uuidPegawai == "" {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("uuid pegawai tidak boleh kosong")

	}
	pegawai, err := pegawaiRepo.GetPegawaiByUUID(a, uuidPegawai)
	if err != nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("error from repo get pegawai by uuid, %w", err)
	}
	if pegawai == nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("pegawai tidak ditemukan")

	}

	jenissk, err := skRepo.GetJenisSKByCode(a, skPegawaiModel.KdJenisSkPengangkatan)
	if err != nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("error from repo get jenis sk pegawai by code, %w", err)
	}
	if jenissk == nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("jenis sk pengangkatan tidak ditemukan")
	}

	jabatanFungsional, err := jabatanFungsionalRepo.GetJabatanFungsionalByUUID(a, skPengangkatanDosen.UUIDJabatanFungsionalLama)
	if err != nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("error from repo get jabatan fungsional pegawai by uuid, %w", err)
	}
	if jabatanFungsional == nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("jabatan fungsional pegawai tidak ditemukan")
	}

	pangkatGolonganPegawai, err := pangkatGolonganPegawaiRepo.GetPangkatGolonganPegawaiByUUID(a, skPengangkatanDosen.UUIDPangkatGolonganPegawaiLama)
	if err != nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("error from repo get pangkat golongan lama pegawai by uuid, %w", err)
	}
	if pangkatGolonganPegawai == nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("pangkat golongan lama pegawai tidak ditemukan")
	}

	indukKerja, err := indukKerjaRepo.GetIndukKerjaByUUID(a, skPengangkatanDosen.UUIDIndukKerjaBaru)
	if err != nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("error from repo induk kerja by uuid, %w", err)
	}
	//fmt.Printf("\n\n[DEBUG] uuid induk kerja ada tidak : \n%+v\n", indukKerja)
	if indukKerja == nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("induk kerja tidak ditemukan")
	}
	pangkatGolonganPegawaiBaru, err := pangkatGolonganPegawaiRepo.GetPangkatGolonganPegawaiByUUID(a, skPengangkatanDosen.UUIDPangkatGolonganPegawaiBaru)
	if err != nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("error from repo pangkat golongan baru by uuid, %w", err)
	}
	if pangkatGolonganPegawaiBaru == nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("pangkat golongan baru tidak ditemukan")
	}
	unitKerja, err := unitKerjaRepo.GetUnitKerjaByUUID(a, skPengangkatanDosen.UUIDUnitKerja)
	if err != nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("error from repo get unit kerja by uuid, %w", err)
	}
	if unitKerja == nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("unit kerja tidak ditemukan")
	}

	jabatanFungsionalBaru, err := jabatanFungsionalRepo.GetJabatanFungsionalByUUID(a, skPengangkatanDosen.UUIDJabatanFungsionalBaru)
	if err != nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("error from repo get jabatan fungsional baru by uuid, %w", err)
	}
	if jabatanFungsionalBaru == nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("jabatan fungsional baru tidak ditemukan")
	}

	jenisIjazah, err := skRepo.GetJenisIjazahByUUID(a, skPengangkatanDosen.UUIDJenisIjazah)
	if err != nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("error from repo jenis ijazah pegawai %w", err)
	}
	if jenisIjazah == nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("ijazah tidak ditemukan")
	}
	jabatanPenetap, err := skRepo.GetAllJabatanPenetapByUUID(a, skPengangkatanDosen.UUIDJabatanPenetap)
	if err != nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("error from repo jabatan penetap pegawai %w", err)
	}
	if jabatanPenetap == nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("jabatan penetap tidak ditemukan")
	}
	pegawaiPenetap, err := skRepo.GetPegawaiPenetapSKByUUID(a, skPengangkatanDosen.UUIDPegawaiPenetap)
	if err != nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("error from repo pegawai penetap pegawai %w", err)
	}
	if pegawaiPenetap == nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("pegawai penetap tidak ditemukan")
	}

	// fmt.Printf("\n\n[DEBUG] uuid mata kuliah str : %s\n", skPengangkatanDosen.UUIDMataKuliahStr)
	err = json.Unmarshal([]byte(skPengangkatanDosen.UUIDMataKuliahStr), &skPengangkatanDosen.UUIDMataKuliah)
	if err != nil {
		fmt.Printf("\n\n[WARNING] error unmarshaling uuid mata kuliah str: %s\n", err.Error())
	}

	// fmt.Printf("\n\n[DEBUG] log uuid mata kuliah ID : %+v\n", skPengangkatanDosen.UUIDMataKuliah)
	idMataKuliah, err := skRepo.GetMataKuliahIDByUUID(a, skPengangkatanDosen.UUIDMataKuliah)
	if err != nil {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("error from repo makul %w", err)
	}
	if len(idMataKuliah) == 0 {
		return skPegawaiModel.SKPegawai{}, fmt.Errorf("mata kuliah tidak boleh kosong")
	}

	// ini untuk sk pengangkatan tendiknya
	skPengangkatanDosen.UserUpdate = c.Request().Header.Get("X-Member")
	skPengangkatanDosen.UserInput = c.Request().Header.Get("X-Member")
	skPengangkatanDosen.IDJabatanPenetap = jabatanPenetap.ID
	skPengangkatanDosen.IDJenisIjazah = jenisIjazah.ID
	skPengangkatanDosen.IDUnitKerja = unitKerja.ID
	skPengangkatanDosen.IDJenisSKPengangkatan = jenissk.ID
	skPengangkatanDosen.IDPangkatGolonganPegawaiLama = pangkatGolonganPegawai.ID
	skPengangkatanDosen.IDPangkatGolonganPegawaiBaru = pangkatGolonganPegawaiBaru.ID
	skPengangkatanDosen.IDJabatanFungsionalLama = jabatanFungsional.ID
	skPengangkatanDosen.IDJabatanFungsionalBaru = jabatanFungsionalBaru.ID
	//fmt.Printf("\n\n[DEBUG] log induk kerja ID : %+v\n", indukKerja.ID)
	skPengangkatanDosen.IDIndukKerjaBaru = indukKerja.ID
	//fmt.Printf("\n\n[DEBUG] log skPengangkatanDosen.IDIndukKerjaBaru : %+v\n", skPengangkatanDosen.IDIndukKerjaBaru)
	skPengangkatanDosen.IDPegawaiPenetap = pegawaiPenetap.ID
	//fmt.Printf("\n\n[DEBUG] sk pengangkatan: \n%+v\n", skPengangkatanDosen)
	skPengangkatanDosen.IDMataKuliah = idMataKuliah

	// ini buat sk pegawainya
	skPegawai = skPegawaiModel.SKPegawai{}
	skPegawai.SKPengangkatanDosen = *skPengangkatanDosen
	skPegawai.NomorSK = c.FormValue("nomor_sk")
	skPegawai.IDJenisSK = jenissk.ID
	skPegawai.TentangSK = c.FormValue("tentang_sk")
	skPegawai.TMT = c.FormValue("tmt")
	skPegawai.UserUpdate = c.Request().Header.Get("X-Member")
	skPegawai.UserInput = c.Request().Header.Get("X-Member")
	skPegawai.IDPegawai = pegawai.ID
	skPegawai.Pegawai = pegawai
	//fmt.Printf("log data : %+v \n", skPegawai.SKPengangkatanPegawai)
	return skPegawai, nil

}

// uuid_mata_kuliah:["uuid-makul-1","uuid-makul-2","uuid-makul-2"]
