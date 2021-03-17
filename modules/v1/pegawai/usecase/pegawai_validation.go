package usecase

import (
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/pegawai/model"
	"svc-insani-go/modules/v1/pegawai/repo"

	indukKerjaRepo "svc-insani-go/modules/v1/master-induk-kerja/repo"
	jabatanFungsionalRepo "svc-insani-go/modules/v1/master-jabatan-fungsional/repo"
	jenisNoRegisRepo "svc-insani-go/modules/v1/master-jenis-nomor-registrasi/repo"
	jenisPTTRepo "svc-insani-go/modules/v1/master-jenis-pegawai-tidak-tetap/repo"
	jenisPegawaiRepo "svc-insani-go/modules/v1/master-jenis-pegawai/repo"
	kelompokPegawaiRepo "svc-insani-go/modules/v1/master-kelompok-pegawai/repo"
	lokasiKerjaRepo "svc-insani-go/modules/v1/master-lokasi-kerja/repo"
	pangkatPegawaiRepo "svc-insani-go/modules/v1/master-pangkat-golongan-pegawai/repo"
	statusPegawaiAktifRepo "svc-insani-go/modules/v1/master-status-pegawai-aktif/repo"
	statusPegawaiRepo "svc-insani-go/modules/v1/master-status-pegawai/repo"

	"github.com/labstack/echo"
)

// func ValidateGetPegawaiByUUID(a app.App, c echo.Context) (*model.Pegawai, error) {
// 	uuidPegawai := new(model.Pegawai)
// 	if uuidPegawai.UUIDPegawai != "" {
// 		Pegawai, err := repo.GetPegawaiByUUID(a, uuidPegawai.UUIDPegawai)
// 		if err != nil {
// 			return nil, fmt.Errorf("%w", fmt.Errorf("[ERROR] get jenis presensi by uuid, %w", err))
// 		}
// 		if Pegawai == nil {
// 			return nil, fmt.Errorf("uuid jenis presensi tidak ditemukan")
// 		}
// 		uuidPegawai.UUIDPegawai = Pegawai.ID
// 	}
// }

func ValidateUpdatePegawaiByUUID(a app.App, c echo.Context) (*model.PegawaiUpdate, error) {
	uuidPegawai := c.Param("uuidPegawai")
	if uuidPegawai == "" {
		return nil, fmt.Errorf("uuid pegawai tidak boleh kosong")
	}

	pegawai, err := repo.GetPegawaiByUUID(a, uuidPegawai)
	if err != nil {
		return nil, fmt.Errorf("error from repo get uuid pegawai, %w", err)
	}

	//Binding Form Input To Struct
	pegawaiRequest := &model.PegawaiUpdate{}
	err = c.Bind(pegawaiRequest)
	if err != nil {
		fmt.Printf("[ERROR] binding request, %s\n", err.Error())
	}
	pegawaiRequest.Uuid = uuidPegawai
	pegawaiRequest.Id = pegawai.ID

	//Pengecekan Jenis Pegawai
	if pegawaiRequest.UuidJenisPegawai != "" {

		jenisPegawai, err := jenisPegawaiRepo.GetJenisPegawaiByUUID(a, pegawaiRequest.UuidJenisPegawai)
		if err != nil {
			return nil, fmt.Errorf("error from repo jenis pegawai by uuid, %w", err)
		}
		pegawaiRequest.IdJenisPegawai = jenisPegawai.ID
		pegawaiRequest.KdJenisPegawai = jenisPegawai.KDJenisPegawai
	}

	// Pengecekan Status Pegawai
	if pegawaiRequest.UuidStatusPegawai != "" {
		statusPegawai, err := statusPegawaiRepo.GetStatusPegawaiByUUID(a, c.Request().Context(), pegawaiRequest.UuidStatusPegawai)
		if err != nil {
			return nil, fmt.Errorf("error from repo status pegawai by uuid, %w", err)
		}
		pegawaiRequest.IdStatusPegawai = statusPegawai.ID
		pegawaiRequest.KdStatusPegawai = statusPegawai.KDStatusPegawai
	}

	// Pengecekan Kelompok Pegawai
	if pegawaiRequest.UuidKelompokPegawai != "" {
		kelompokPegawai, err := kelompokPegawaiRepo.GetKelompokPegawaiByUUID(a, c.Request().Context(), pegawaiRequest.UuidKelompokPegawai)
		if err != nil {
			return nil, fmt.Errorf("error from repo kelompok pegawai by uuid, %w", err)
		}
		pegawaiRequest.IdKelompokPegawai = kelompokPegawai.ID
		pegawaiRequest.KdKelompokPegawai = kelompokPegawai.KdKelompokPegawai
	}

	// Pengecekan Pangkat Golongan Pegawai
	if pegawaiRequest.UuidGolongan != "" {
		pangkatPegawai, err := pangkatPegawaiRepo.GetPangkatPegawaiByUUID(a, pegawaiRequest.UuidGolongan)
		if err != nil {
			return nil, fmt.Errorf("error from repo pangkat golongan pegawai by uuid, %w", err)
		}
		pegawaiRequest.IdGolongan = pangkatPegawai.ID
		pegawaiRequest.KdGolongan = pangkatPegawai.KdPangkat
	}

	// Pengecekan Jabatan Fungsional Yayasan
	if pegawaiRequest.PegawaiFungsional.UuidJabatanFungsional != "" {
		jabatanFungsional, err := jabatanFungsionalRepo.GetJabatanFungsionalByUUID(a, c.Request().Context(), pegawaiRequest.PegawaiFungsional.UuidJabatanFungsional)
		if err != nil {
			return nil, fmt.Errorf("error from repo pangkat golongan pegawai by uuid, %w", err)
		}
		pegawaiRequest.PegawaiFungsional.IdJabatanFungsional = jabatanFungsional.ID
		pegawaiRequest.PegawaiFungsional.KdJabatanFungsional = jabatanFungsional.KdJabatanFungsional
	}

	// Pengecekan Jenis Nomor Resgistrasi
	if pegawaiRequest.PegawaiFungsional.UuidJenisNomorRegistrasi != "" {
		jenisNoRegis, err := jenisNoRegisRepo.GetJenisNoRegisByUUID(a, c.Request().Context(), pegawaiRequest.PegawaiFungsional.UuidJenisNomorRegistrasi)
		if err != nil {
			return nil, fmt.Errorf("error from repo jenis nomor registrasi by uuid, %w", err)
		}
		pegawaiRequest.PegawaiFungsional.IdJenisNomorRegistrasi = jenisNoRegis.ID
		pegawaiRequest.PegawaiFungsional.KdJenisNomorRegistrasi = jenisNoRegis.KdJenisRegis
	}

	// Pengecekan Induk Kerja
	if pegawaiRequest.UuidUnitKerja1 != "" {
		indukKerja, err := indukKerjaRepo.GetIndukKerjaByUUID(a, c.Request().Context(), pegawaiRequest.UuidUnitKerja1)
		if err != nil {
			return nil, fmt.Errorf("error from repo induk kerja by uuid, %w", err)
		}
		pegawaiRequest.IdUnitKerja1 = indukKerja.ID
		pegawaiRequest.KdUnit1 = indukKerja.KdUnit1
	}

	// Pengecekan Unit Kerja
	if pegawaiRequest.UuidUnitKerja2 != "" {
		unitKerja, err := indukKerjaRepo.GetUnitKerjaByUUID(a, c.Request().Context(), pegawaiRequest.UuidUnitKerja2)
		if err != nil {
			return nil, fmt.Errorf("error from repo unit kerja by uuid, %w", err)
		}
		pegawaiRequest.IdUnitKerja2 = unitKerja.ID
		pegawaiRequest.KdUnit2 = unitKerja.KdUnit2
	}

	// Pengecekan Bagian Kerja
	if pegawaiRequest.UuidUnitKerja3 != "" {
		bagianKerja, err := indukKerjaRepo.GetBagianKerjaByUUID(a, c.Request().Context(), pegawaiRequest.UuidUnitKerja3)
		if err != nil {
			return nil, fmt.Errorf("error from repo bagian kerja by uuid, %w", err)
		}
		pegawaiRequest.IdUnitKerja3 = bagianKerja.ID
		pegawaiRequest.KdUnit3 = bagianKerja.KdUnit3
	}

	// Pengecekan Lokasi Kerja
	if pegawaiRequest.UuidUnitKerjaLokasi != "" {
		lokasiKerja, err := lokasiKerjaRepo.GetLokasiKerjaByUUID(a, c.Request().Context(), pegawaiRequest.UuidUnitKerjaLokasi)
		if err != nil {
			return nil, fmt.Errorf("error from repo lokasi kerja by uuid, %w", err)
		}
		pegawaiRequest.IdUnitKerjaLokasi = lokasiKerja.ID
		pegawaiRequest.LokasiKerja = lokasiKerja.LokasiKerja
	}

	// Pengecekan Pangkat Golongan Ruang PNS
	if pegawaiRequest.PegawaiPNS.UuidPangkatGolongan != "" {
		pangkatPNS, err := pangkatPegawaiRepo.GetPangkatPegawaiByUUID(a, pegawaiRequest.PegawaiPNS.UuidPangkatGolongan)
		if err != nil {
			return nil, fmt.Errorf("error from repo pangkat golongan pegawai by uuid, %w", err)
		}
		pegawaiRequest.PegawaiPNS.IdPangkatGolongan = pangkatPNS.ID
		pegawaiRequest.PegawaiPNS.KdPangkatGolongan = pangkatPNS.KdPangkat
	}

	// Pengecekan Jabatan Fungsional Yayasan
	if pegawaiRequest.PegawaiPNS.UuidJabatanFungsional != "" {
		jabatanPNS, err := jabatanFungsionalRepo.GetJabatanFungsionalByUUID(a, c.Request().Context(), pegawaiRequest.PegawaiPNS.UuidJabatanFungsional)
		if err != nil {
			return nil, fmt.Errorf("error from repo pangkat golongan pegawai by uuid, %w", err)
		}
		pegawaiRequest.PegawaiPNS.IdJabatanFungsional = jabatanPNS.ID
		pegawaiRequest.PegawaiPNS.KdJabatanFungsional = jabatanPNS.KdJabatanFungsional
	}

	// Pengecekan Jenis PTT
	if pegawaiRequest.PegawaiPTT.UuidJenisPtt != "" {
		jenisPTT, err := jenisPTTRepo.GetJenisPTTByUUID(a, c.Request().Context(), pegawaiRequest.PegawaiPTT.UuidJenisPtt)
		if err != nil {
			return nil, fmt.Errorf("error from repo jenis pegawai tidak tetap by uuid, %w", err)
		}
		pegawaiRequest.PegawaiPTT.IdJenisPtt = jenisPTT.ID
		pegawaiRequest.PegawaiPTT.KdJenisPtt = jenisPTT.KdJenisPTT
	}

	// Pengecekan Status Pegawai
	if pegawaiRequest.PegawaiFungsional.UuidStatusPegawaiAktif != "" {
		statusPegawaiAktif, err := statusPegawaiAktifRepo.GetStatusPegawaiAktifByUUID(a, c.Request().Context(), pegawaiRequest.PegawaiFungsional.UuidStatusPegawaiAktif)
		if err != nil {
			return nil, fmt.Errorf("error from repo status pegawai aktif by uuid, %w", err)
		}
		pegawaiRequest.PegawaiFungsional.IdStatusPegawaiAktif = statusPegawaiAktif.ID
		pegawaiRequest.PegawaiFungsional.KdStatusPegawaiAktif = statusPegawaiAktif.KdStatusAktif
	}

	return pegawaiRequest, nil
}
