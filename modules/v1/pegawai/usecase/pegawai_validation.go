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

	"github.com/cstockton/go-conv"
	"github.com/labstack/echo"
)

// func ValidateGetPegawaiByUUID(a app.App, c echo.Context) (*model.Pegawai, error) {
// 	uuidPegawai := new(model.Pegawai)
// 	if uuidPegawai.UUIDPegawai != "" {
// 		Pegawai, err := repo.GetPegawaiByUUID(a, uuidPegawai.UUIDPegawai)
// 		if err != nil {
// 			return fmt.Errorf("%w", fmt.Errorf("[ERROR] get jenis presensi by uuid, %w", err))
// 		}
// 		if Pegawai == nil {
// 			return fmt.Errorf("uuid jenis presensi tidak ditemukan")
// 		}
// 		uuidPegawai.UUIDPegawai = Pegawai.ID
// 	}
// }

func ValidateUpdatePegawaiByUUID(a app.App, c echo.Context) (model.PegawaiUpdate, error) {
	uuidPegawai := c.Param("uuidPegawai")
	if uuidPegawai == "" {
		return model.PegawaiUpdate{}, fmt.Errorf("uuid pegawai tidak boleh kosong")
	}

	pegawai, err := repo.GetPegawaiByUUID(a, uuidPegawai)
	if err != nil {
		return model.PegawaiUpdate{}, fmt.Errorf("error from repo get uuid pegawai, %w", err)
	}

	pegawaiOld, err := repo.GetOldPegawai(a, c.Request().Context(), uuidPegawai)
	if err != nil {
		return model.PegawaiUpdate{}, fmt.Errorf("error from get old pegawai, %w", err)
	}

	user := "ega.evan"

	pegawaiReq := &model.PegawaiUpdate{}

	err = c.Bind(pegawaiReq)
	if err != nil {
		fmt.Printf("[ERROR] binding requestpegawai , %s\n", err.Error())
	}

	pegawaiReq.Uuid = uuidPegawai
	pegawaiReq.Id, _ = conv.Int(pegawai.ID)

	//Pengecekan Jenis Pegawai
	if pegawaiReq.UuidJenisPegawai != "" {

		jenisPegawai, err := jenisPegawaiRepo.GetJenisPegawaiByUUID(a, pegawaiReq.UuidJenisPegawai)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo jenis pegawai by uuid, %w", err)
		}
		pegawaiOld.IdJenisPegawai, _ = conv.Int(jenisPegawai.ID)
		pegawaiOld.KdJenisPegawai = jenisPegawai.KDJenisPegawai
	}

	// Pengecekan Status Pegawai
	if pegawaiReq.UuidStatusPegawai != "" {
		statusPegawai, err := statusPegawaiRepo.GetStatusPegawaiByUUID(a, c.Request().Context(), pegawaiReq.UuidStatusPegawai)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo status pegawai by uuid, %w", err)
		}
		pegawaiOld.IdStatusPegawai, _ = conv.Int(statusPegawai.ID)
		pegawaiOld.KdStatusPegawai = statusPegawai.KDStatusPegawai
	}

	// Pengecekan Kelompok Pegawai
	if pegawaiReq.UuidKelompokPegawai != "" {
		kelompokPegawai, err := kelompokPegawaiRepo.GetKelompokPegawaiByUUID(a, c.Request().Context(), pegawaiReq.UuidKelompokPegawai)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo kelompok pegawai by uuid, %w", err)
		}
		pegawaiOld.IdKelompokPegawai, _ = conv.Int(kelompokPegawai.ID)
		pegawaiOld.KdKelompokPegawai = kelompokPegawai.KdKelompokPegawai
	}

	// Pengecekan Pangkat Golongan Pegawai
	if pegawaiReq.UuidGolongan != "" {
		pangkatPegawai, err := pangkatPegawaiRepo.GetPangkatGolonganPegawaiByUUID(a, pegawaiReq.UuidGolongan)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo pangkat golongan pegawai by uuid, %w", err)
		}
		pegawaiOld.IdGolongan, _ = conv.Int(pangkatPegawai.ID)
		pegawaiOld.KdGolongan = pangkatPegawai.KdPangkat
	}

	// Pengecekan Jabatan Fungsional Yayasan
	if pegawaiReq.PegawaiFungsional.UuidJabatanFungsional != "" {
		jabatanFungsional, err := jabatanFungsionalRepo.GetJabatanFungsionalByUUID(a, c.Request().Context(), pegawaiReq.PegawaiFungsional.UuidJabatanFungsional)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo pangkat golongan pegawai by uuid, %w", err)
		}
		pegawaiOld.PegawaiFungsional.IdJabatanFungsional, _ = conv.Int(jabatanFungsional.ID)
		pegawaiOld.PegawaiFungsional.KdJabatanFungsional = jabatanFungsional.KdJabatanFungsional
	}

	// Pengecekan Jenis Nomor Resgistrasi
	if pegawaiReq.PegawaiFungsional.UuidJenisNomorRegistrasi != "" {
		jenisNoRegis, err := jenisNoRegisRepo.GetJenisNoRegisByUUID(a, c.Request().Context(), pegawaiReq.PegawaiFungsional.UuidJenisNomorRegistrasi)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo jenis nomor registrasi by uuid, %w", err)
		}
		pegawaiOld.PegawaiFungsional.IdJenisNomorRegistrasi, _ = conv.Int(jenisNoRegis.ID)
		pegawaiOld.PegawaiFungsional.KdJenisNomorRegistrasi = jenisNoRegis.KdJenisRegis
	}

	// Pengecekan Induk Kerja
	if pegawaiReq.UuidUnitKerja1 != "" {
		indukKerja, err := indukKerjaRepo.GetIndukKerjaByUUID(a, c.Request().Context(), pegawaiReq.UuidUnitKerja1)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo induk kerja by uuid, %w", err)
		}
		pegawaiOld.IdUnitKerja1, _ = conv.Int(indukKerja.ID)
		pegawaiOld.KdUnit1 = indukKerja.KdUnit1
	}

	// Pengecekan Unit Kerja
	if pegawaiReq.UuidUnitKerja2 != "" {
		unitKerja, err := indukKerjaRepo.GetUnitKerjaByUUID(a, c.Request().Context(), pegawaiReq.UuidUnitKerja2)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo unit kerja by uuid, %w", err)
		}
		pegawaiOld.IdUnitKerja2, _ = conv.Int(unitKerja.ID)
		pegawaiOld.KdUnit2 = unitKerja.KdUnit2
	}

	// Pengecekan Bagian Kerja
	if pegawaiReq.UuidUnitKerja3 != "" {
		bagianKerja, err := indukKerjaRepo.GetBagianKerjaByUUID(a, c.Request().Context(), pegawaiReq.UuidUnitKerja3)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo bagian kerja by uuid, %w", err)
		}
		pegawaiOld.IdUnitKerja3, _ = conv.Int(bagianKerja.ID)
		pegawaiOld.KdUnit3 = bagianKerja.KdUnit3
	}

	// Pengecekan Lokasi Kerja
	if pegawaiReq.UuidUnitKerjaLokasi != "" {
		lokasiKerja, err := lokasiKerjaRepo.GetLokasiKerjaByUUID(a, c.Request().Context(), pegawaiReq.UuidUnitKerjaLokasi)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo lokasi kerja by uuid, %w", err)
		}
		pegawaiOld.IdUnitKerjaLokasi, _ = conv.Int(lokasiKerja.ID)
		pegawaiOld.LokasiKerja = lokasiKerja.LokasiKerja
	}

	// Pengecekan Homebase Pddikti
	if pegawaiReq.PegawaiFungsional.UuidHomebasePddikti != "" {
		homebasePddikti, err := indukKerjaRepo.GetUnitKerjaByUUID(a, c.Request().Context(), pegawaiReq.PegawaiFungsional.UuidHomebasePddikti)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo unit kerja by uuid, %w", err)
		}
		pegawaiOld.PegawaiFungsional.IdHomebasePddikti, _ = conv.Int(homebasePddikti.ID)
	}

	// Pengecekan Homebase UII
	if pegawaiReq.PegawaiFungsional.UuidHomebaseUii != "" {
		homebaseUUuidHomebaseUii, err := indukKerjaRepo.GetUnitKerjaByUUID(a, c.Request().Context(), pegawaiReq.PegawaiFungsional.UuidHomebaseUii)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo unit kerja by uuid, %w", err)
		}
		pegawaiOld.PegawaiFungsional.IdHomebaseUii, _ = conv.Int(homebaseUUuidHomebaseUii.ID)
	}

	// Pengecekan Pangkat Golongan Ruang PNS
	if pegawaiReq.PegawaiPNS.UuidPangkatGolongan != "" {
		pangkatPNS, err := pangkatPegawaiRepo.GetPangkatPegawaiByUUID(a, pegawaiReq.PegawaiPNS.UuidPangkatGolongan)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo pangkat golongan pegawai by uuid, %w", err)
		}
		pegawaiOld.PegawaiPNS.IdPangkatGolongan, _ = conv.Int(pangkatPNS.ID)
		pegawaiOld.PegawaiPNS.KdPangkatGolongan = pangkatPNS.KdPangkat
	}

	// Pengecekan Jabatan Fungsional Yayasan
	if pegawaiReq.PegawaiPNS.UuidJabatanFungsional != "" {
		jabatanPNS, err := jabatanFungsionalRepo.GetJabatanFungsionalByUUID(a, c.Request().Context(), pegawaiReq.PegawaiPNS.UuidJabatanFungsional)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo pangkat golongan pegawai by uuid, %w", err)
		}
		pegawaiOld.PegawaiPNS.IdJabatanFungsional, _ = conv.Int(jabatanPNS.ID)
		pegawaiOld.PegawaiPNS.KdJabatanFungsional = jabatanPNS.KdJabatanFungsional
	}

	// Pengecekan Jenis PTT
	if pegawaiReq.PegawaiPNS.UuidJenisPtt != "" {
		jenisPTT, err := jenisPTTRepo.GetJenisPTTByUUID(a, c.Request().Context(), pegawaiReq.PegawaiPNS.UuidJenisPtt)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo jenis pegawai tidak tetap by uuid, %w", err)
		}
		pegawaiOld.PegawaiPNS.IdJenisPtt, _ = conv.Int(jenisPTT.ID)
		pegawaiOld.PegawaiPNS.KdJenisPtt = jenisPTT.KdJenisPTT
	}

	// Pengecekan Status Pegawai
	if pegawaiReq.PegawaiFungsional.UuidStatusPegawaiAktif != "" {
		statusPegawaiAktif, err := statusPegawaiAktifRepo.GetStatusPegawaiAktifByUUID(a, c.Request().Context(), pegawaiReq.PegawaiFungsional.UuidStatusPegawaiAktif)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo status pegawai aktif by uuid, %w", err)
		}
		pegawaiOld.PegawaiFungsional.IdStatusPegawaiAktif, _ = conv.Int(statusPegawaiAktif.ID)
		pegawaiOld.PegawaiFungsional.KdStatusPegawaiAktif = statusPegawaiAktif.KdStatusAktif
	}

	// Binding nilai request ke struct

	if pegawaiReq.PegawaiFungsional.TmtPangkatGolongan != "" {
		pegawaiOld.PegawaiFungsional.TmtPangkatGolongan = pegawaiReq.PegawaiFungsional.TmtPangkatGolongan
	}
	if pegawaiReq.PegawaiFungsional.TmtJabatan != "" {
		pegawaiOld.PegawaiFungsional.TmtJabatan = pegawaiReq.PegawaiFungsional.TmtJabatan
	}
	if pegawaiReq.PegawaiFungsional.MasaKerjaBawaanTahun != "" {
		pegawaiOld.PegawaiFungsional.MasaKerjaBawaanTahun = pegawaiReq.PegawaiFungsional.MasaKerjaBawaanTahun
	}
	if pegawaiReq.PegawaiFungsional.MasaKerjaBawaanBulan != "" {
		pegawaiOld.PegawaiFungsional.MasaKerjaBawaanBulan = pegawaiReq.PegawaiFungsional.MasaKerjaBawaanBulan
	}
	if pegawaiReq.PegawaiFungsional.MasaKerjaGajiTahun != "" {
		pegawaiOld.PegawaiFungsional.MasaKerjaGajiTahun = pegawaiReq.PegawaiFungsional.MasaKerjaGajiTahun
	}
	if pegawaiReq.PegawaiFungsional.MasaKerjaGajiBulan != "" {
		pegawaiOld.PegawaiFungsional.MasaKerjaGajiBulan = pegawaiReq.PegawaiFungsional.MasaKerjaGajiBulan
	}
	if pegawaiReq.PegawaiFungsional.MasaKerjaTotalTahun != "" {
		pegawaiOld.PegawaiFungsional.MasaKerjaTotalTahun = pegawaiReq.PegawaiFungsional.MasaKerjaTotalTahun
	}
	if pegawaiReq.PegawaiFungsional.MasaKerjaTotalBulan != "" {
		pegawaiOld.PegawaiFungsional.MasaKerjaTotalBulan = pegawaiReq.PegawaiFungsional.MasaKerjaTotalBulan
	}
	if pegawaiReq.PegawaiFungsional.AngkaKredit != "" {
		pegawaiOld.PegawaiFungsional.AngkaKredit = pegawaiReq.PegawaiFungsional.AngkaKredit
	}
	if pegawaiReq.PegawaiFungsional.NomorSertifikasi != "" {
		pegawaiOld.PegawaiFungsional.NomorSertifikasi = pegawaiReq.PegawaiFungsional.NomorSertifikasi
	}
	if pegawaiReq.PegawaiFungsional.NomorRegistrasi != "" {
		pegawaiOld.PegawaiFungsional.NomorRegistrasi = pegawaiReq.PegawaiFungsional.NomorRegistrasi
	}
	if pegawaiReq.PegawaiFungsional.NomorSkPertama != "" {
		pegawaiOld.PegawaiFungsional.NomorSkPertama = pegawaiReq.PegawaiFungsional.NomorSkPertama
	}
	if pegawaiReq.PegawaiFungsional.TmtSkPertama != "" {
		pegawaiOld.PegawaiFungsional.TmtSkPertama = pegawaiReq.PegawaiFungsional.TmtSkPertama
	}
	if pegawaiReq.PegawaiPNS.InstansiAsal != "" {
		pegawaiOld.PegawaiPNS.InstansiAsal = pegawaiReq.PegawaiPNS.InstansiAsal
	}
	if pegawaiReq.PegawaiPNS.NipPns != "" {
		pegawaiOld.PegawaiPNS.NipPns = pegawaiReq.PegawaiPNS.NipPns
	}
	if pegawaiReq.PegawaiPNS.NoKartuPegawai != "" {
		pegawaiOld.PegawaiPNS.NoKartuPegawai = pegawaiReq.PegawaiPNS.NoKartuPegawai
	}
	if pegawaiReq.PegawaiPNS.TmtPangkatGolongan != "" {
		pegawaiOld.PegawaiPNS.TmtPangkatGolongan = pegawaiReq.PegawaiPNS.TmtPangkatGolongan
	}
	if pegawaiReq.PegawaiPNS.TmtJabatan != "" {
		pegawaiOld.PegawaiPNS.TmtJabatan = pegawaiReq.PegawaiPNS.TmtJabatan
	}
	if pegawaiReq.PegawaiPNS.MasaKerjaTahun != "" {
		pegawaiOld.PegawaiPNS.MasaKerjaTahun = pegawaiReq.PegawaiPNS.MasaKerjaTahun
	}
	if pegawaiReq.PegawaiPNS.MasaKerjaBulan != "" {
		pegawaiOld.PegawaiPNS.MasaKerjaBulan = pegawaiReq.PegawaiPNS.MasaKerjaBulan
	}
	if pegawaiReq.PegawaiPNS.AngkaKredit != "" {
		pegawaiOld.PegawaiPNS.AngkaKredit = pegawaiReq.PegawaiPNS.AngkaKredit
	}
	if pegawaiReq.PegawaiPNS.Keterangan != "" {
		pegawaiOld.PegawaiPNS.Keterangan = pegawaiReq.PegawaiPNS.Keterangan
	}
	pegawaiOld.UserUpdate = user
	pegawaiOld.PegawaiFungsional.UserUpdate = user

	return pegawaiOld, nil
}
