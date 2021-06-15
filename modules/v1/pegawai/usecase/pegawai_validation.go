package usecase

import (
	"fmt"
	"strconv"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/pegawai/model"
	"svc-insani-go/modules/v1/pegawai/repo"

	jabatanFungsionalRepo "svc-insani-go/modules/v1/master-jabatan-fungsional/repo"
	jenisNoRegisRepo "svc-insani-go/modules/v1/master-jenis-nomor-registrasi/repo"
	jenisPTTRepo "svc-insani-go/modules/v1/master-jenis-pegawai-tidak-tetap/repo"
	jenisPegawaiRepo "svc-insani-go/modules/v1/master-jenis-pegawai/repo"
	jenjangPendidikan "svc-insani-go/modules/v1/master-jenjang-pendidikan/repo"
	kelompokPegawaiRepo "svc-insani-go/modules/v1/master-kelompok-pegawai/repo"
	lokasiKerjaRepo "svc-insani-go/modules/v1/master-lokasi-kerja/repo"
	indukKerjaRepo "svc-insani-go/modules/v1/master-organisasi/repo"
	pangkatPegawaiRepo "svc-insani-go/modules/v1/master-pangkat-golongan-pegawai/repo"
	statusPegawaiAktifRepo "svc-insani-go/modules/v1/master-status-pegawai-aktif/repo"
	statusPegawaiRepo "svc-insani-go/modules/v1/master-status-pegawai/repo"
	personalRepo "svc-insani-go/modules/v1/personal/repo"

	"github.com/cstockton/go-conv"
	"github.com/labstack/echo"
)

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

	user := c.Request().Header.Get("X-Member")

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

	// Pengecekan Ijazah Pendidikan Masuk
	fmt.Println("DEBUG : ", pegawaiReq.UuidPendidikanMasuk)
	if pegawaiReq.UuidPendidikanMasuk != "" {
		pendidikanMasuk, err := jenjangPendidikan.GetJenjangPendidikanByUUID(a, c.Request().Context(), pegawaiReq.UuidPendidikanMasuk)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo jenis ijazah by uuid, %w", err)
		}
		pegawaiOld.IdPendidikanMasuk, _ = conv.Int(pendidikanMasuk.ID)
		pegawaiOld.KdPendidikanMasuk = pendidikanMasuk.KdJenjang
	}

	// Pengecekan Pangkat Golongan Pegawai
	if pegawaiReq.UuidGolongan != "" {
		pangkatPegawai, err := pangkatPegawaiRepo.GetPangkatGolonganPegawaiByUUID(a, pegawaiReq.UuidGolongan)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo pangkat golongan pegawai by uuid, %w", err)
		}
		pegawaiOld.PegawaiFungsional.IdPangkatGolongan, _ = conv.Int(pangkatPegawai.ID)
		pegawaiOld.PegawaiFungsional.KdPangkatGolongan = pangkatPegawai.KdPangkat
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
	if pegawaiReq.UuidLokasiKerja != "" {
		lokasiKerja, err := lokasiKerjaRepo.GetLokasiKerjaByUUID(a, c.Request().Context(), pegawaiReq.UuidLokasiKerja)
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
		pangkatPNS, err := pangkatPegawaiRepo.GetPangkatGolonganPegawaiByUUID(a, pegawaiReq.PegawaiPNS.UuidPangkatGolongan)
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

	if pegawaiReq.PegawaiFungsional.TmtPangkatGolongan != nil {
		pegawaiOld.PegawaiFungsional.TmtPangkatGolongan = pegawaiReq.PegawaiFungsional.TmtPangkatGolongan
	}
	if pegawaiReq.PegawaiFungsional.TmtJabatan != nil {
		pegawaiOld.PegawaiFungsional.TmtJabatan = pegawaiReq.PegawaiFungsional.TmtJabatan
	}
	if pegawaiReq.PegawaiFungsional.MasaKerjaBawaanTahun != "" {
		pegawaiOld.PegawaiFungsional.MasaKerjaBawaanTahun = pegawaiReq.PegawaiFungsional.MasaKerjaBawaanTahun
	}
	if pegawaiReq.PegawaiFungsional.MasaKerjaBawaanBulan != "" {
		a, _ := strconv.Atoi(pegawaiReq.PegawaiFungsional.MasaKerjaBawaanBulan)
		if a > 12 {
			return model.PegawaiUpdate{}, fmt.Errorf("[ERROR] data bulan tidak valid")
		}
		pegawaiOld.PegawaiFungsional.MasaKerjaBawaanBulan = pegawaiReq.PegawaiFungsional.MasaKerjaBawaanBulan
	}
	if pegawaiReq.PegawaiFungsional.MasaKerjaGajiTahun != "" {
		pegawaiOld.PegawaiFungsional.MasaKerjaGajiTahun = pegawaiReq.PegawaiFungsional.MasaKerjaGajiTahun
	}
	if pegawaiReq.PegawaiFungsional.MasaKerjaGajiBulan != "" {
		a, _ := strconv.Atoi(pegawaiReq.PegawaiFungsional.MasaKerjaGajiBulan)
		if a > 12 {
			return model.PegawaiUpdate{}, fmt.Errorf("[ERROR] data bulan tidak valid")
		}
		pegawaiOld.PegawaiFungsional.MasaKerjaGajiBulan = pegawaiReq.PegawaiFungsional.MasaKerjaGajiBulan
	}
	// if pegawaiReq.PegawaiFungsional.MasaKerjaTotalTahun != "" {
	// 	pegawaiOld.PegawaiFungsional.MasaKerjaTotalTahun = pegawaiReq.PegawaiFungsional.MasaKerjaTotalTahun
	// }
	// if pegawaiReq.PegawaiFungsional.MasaKerjaTotalBulan != "" {
	// 	a, _ := strconv.Atoi(pegawaiReq.PegawaiFungsional.MasaKerjaTotalBulan)
	// 	if a > 12 {
	// 		return model.PegawaiUpdate{}, fmt.Errorf("error data bulan tidak valid")
	// 	}
	// 	pegawaiOld.PegawaiFungsional.MasaKerjaTotalBulan = pegawaiReq.PegawaiFungsional.MasaKerjaTotalBulan
	// }
	if pegawaiReq.PegawaiFungsional.AngkaKredit != "" {
		pegawaiOld.PegawaiFungsional.AngkaKredit = pegawaiReq.PegawaiFungsional.AngkaKredit
	}
	if pegawaiReq.PegawaiFungsional.NomorSertifikasi != "" {
		if len(pegawaiReq.PegawaiFungsional.NomorSertifikasi) > 20 {
			return model.PegawaiUpdate{}, fmt.Errorf("error nomor sertifikasi tidak valid")
		}
		pegawaiOld.PegawaiFungsional.NomorSertifikasi = pegawaiReq.PegawaiFungsional.NomorSertifikasi
	}
	if pegawaiReq.PegawaiFungsional.NomorRegistrasi != "" {
		if len(pegawaiReq.PegawaiFungsional.NomorRegistrasi) > 10 {
			return model.PegawaiUpdate{}, fmt.Errorf("error nomor registrasi tidak valid")
		}
		pegawaiOld.PegawaiFungsional.NomorRegistrasi = pegawaiReq.PegawaiFungsional.NomorRegistrasi
	}
	if pegawaiReq.PegawaiFungsional.NomorSkPertama != "" {
		if len(pegawaiReq.PegawaiFungsional.NomorSkPertama) > 30 {
			return model.PegawaiUpdate{}, fmt.Errorf("error nomor sk pertama tidak valid")
		}
		pegawaiOld.PegawaiFungsional.NomorSkPertama = pegawaiReq.PegawaiFungsional.NomorSkPertama
	}
	if pegawaiReq.PegawaiFungsional.TmtSkPertama != nil {
		pegawaiOld.PegawaiFungsional.TmtSkPertama = pegawaiReq.PegawaiFungsional.TmtSkPertama
	}
	if pegawaiReq.PegawaiPNS.InstansiAsal != "" {
		pegawaiOld.PegawaiPNS.InstansiAsal = pegawaiReq.PegawaiPNS.InstansiAsal
	}
	if pegawaiReq.PegawaiPNS.NipPns != "" {
		if len(pegawaiReq.PegawaiPNS.NipPns) > 18 {
			return model.PegawaiUpdate{}, fmt.Errorf("error nip pns tidak valid")
		}
		pegawaiOld.PegawaiPNS.NipPns = pegawaiReq.PegawaiPNS.NipPns
	}
	if pegawaiReq.PegawaiPNS.NoKartuPegawai != "" {
		if len(pegawaiReq.PegawaiPNS.NoKartuPegawai) > 18 {
			return model.PegawaiUpdate{}, fmt.Errorf("error nomor kartu pegawai tidak valid")
		}
		pegawaiOld.PegawaiPNS.NoKartuPegawai = pegawaiReq.PegawaiPNS.NoKartuPegawai
	}
	if pegawaiReq.PegawaiPNS.TmtPangkatGolongan != nil {
		pegawaiOld.PegawaiPNS.TmtPangkatGolongan = pegawaiReq.PegawaiPNS.TmtPangkatGolongan
	}
	if pegawaiReq.PegawaiPNS.TmtJabatan != nil {
		pegawaiOld.PegawaiPNS.TmtJabatan = pegawaiReq.PegawaiPNS.TmtJabatan
	}
	if pegawaiReq.PegawaiPNS.MasaKerjaTahun != "" {
		pegawaiOld.PegawaiPNS.MasaKerjaTahun = pegawaiReq.PegawaiPNS.MasaKerjaTahun
	}
	if pegawaiReq.PegawaiPNS.MasaKerjaBulan != "" {
		a, _ := strconv.Atoi(pegawaiReq.PegawaiPNS.MasaKerjaBulan)
		if a > 12 {
			return model.PegawaiUpdate{}, fmt.Errorf("error data bulan tidak valid")
		}
		pegawaiOld.PegawaiPNS.MasaKerjaBulan = pegawaiReq.PegawaiPNS.MasaKerjaBulan
	}
	if pegawaiReq.PegawaiPNS.AngkaKredit != "" {
		pegawaiOld.PegawaiPNS.AngkaKredit = pegawaiReq.PegawaiPNS.AngkaKredit
	}
	if pegawaiReq.PegawaiPNS.Keterangan != "" {
		pegawaiOld.PegawaiPNS.Keterangan = pegawaiReq.PegawaiPNS.Keterangan
	}
	if pegawaiReq.DetailProfesi != "" {
		pegawaiOld.DetailProfesi = pegawaiReq.DetailProfesi
	}
	pegawaiOld.UserUpdate = user
	pegawaiOld.PegawaiFungsional.UserUpdate = user

	return pegawaiOld, nil
}

func PrepareCreateSimpeg(a app.App, c echo.Context) (model.PegawaiCreate, error) {
	uuidPersonal := c.Param("uuidPegawai")
	if uuidPersonal == "" {
		return model.PegawaiCreate{}, fmt.Errorf("uuid personal tidak boleh kosong")
	}

	// fmt.Println("Uuid personal : ", uuidPersonal)

	pegawai, err := ValidateCreatePegawai(a, c)
	if err != nil {
		return model.PegawaiCreate{}, fmt.Errorf("error validate create pegawai, %s\n", err.Error())
	}

	personal, err := personalRepo.GetPersonalByUuid(a, c.Request().Context(), uuidPersonal)
	if err != nil {
		return model.PegawaiCreate{}, fmt.Errorf("error search personal")
	}

	pegawai.IdPersonalDataPribadi, _ = conv.String(personal.Id)
	pegawai.Nama = personal.NamaLengkap
	pegawai.NikKtp = personal.NikKtp
	// pegawai.Nik = personal.NikPegawai
	pegawai.TglLahir = personal.TglLahir
	pegawai.TempatLahir = personal.TempatLahir
	pegawai.IdAgama, _ = conv.String(personal.IdAgama)
	pegawai.KdAgama = personal.KdAgama
	pegawai.JenisKelamin = personal.JenisKelamin
	pegawai.IdGolonganDarah = personal.IdGolonganDarah
	pegawai.KdGolonganDarah = personal.KdGolonganDarah
	pegawai.IdStatusPerkawinan = personal.IdStatusPernikahan

	return pegawai, nil
}

func ValidateCreatePegawai(a app.App, c echo.Context) (model.PegawaiCreate, error) {

	user := c.Request().Header.Get("X-Member")

	pegawaiReq := model.PegawaiCreate{}

	fmt.Println("[ERROR] before binding")
	err := c.Bind(&pegawaiReq)
	if err != nil {
		fmt.Printf("[ERROR] binding requestpegawai , %s\n", err.Error())
	}

	fmt.Println("[ERROR] after binding")

	//Pengecekan Jenis Pegawai
	if pegawaiReq.UuidJenisPegawai != "" {
		jenisPegawai, err := jenisPegawaiRepo.GetJenisPegawaiByUUID(a, pegawaiReq.UuidJenisPegawai)
		if err != nil {
			return model.PegawaiCreate{}, fmt.Errorf("error from repo jenis pegawai by uuid, %w", err)
		}
		pegawaiReq.IdJenisPegawai, err = conv.Int(jenisPegawai.ID)
		if err != nil {
			fmt.Println(err)
			return model.PegawaiCreate{}, err
		}
		pegawaiReq.KdJenisPegawai = jenisPegawai.KDJenisPegawai
	}

	// Pengecekan Status Pegawai
	if pegawaiReq.UuidStatusPegawai != "" {
		statusPegawai, err := statusPegawaiRepo.GetStatusPegawaiByUUID(a, c.Request().Context(), pegawaiReq.UuidStatusPegawai)
		if err != nil {
			return model.PegawaiCreate{}, fmt.Errorf("error from repo status pegawai by uuid, %w", err)
		}
		pegawaiReq.IdStatusPegawai, _ = conv.Int(statusPegawai.ID)
		pegawaiReq.KdStatusPegawai = statusPegawai.KDStatusPegawai
	}

	// Pengecekan Kelompok Pegawai
	if pegawaiReq.UuidKelompokPegawai != "" {
		kelompokPegawai, err := kelompokPegawaiRepo.GetKelompokPegawaiByUUID(a, c.Request().Context(), pegawaiReq.UuidKelompokPegawai)
		if err != nil {
			return model.PegawaiCreate{}, fmt.Errorf("error from repo kelompok pegawai by uuid, %w", err)
		}
		pegawaiReq.IdKelompokPegawai, _ = conv.Int(kelompokPegawai.ID)
		pegawaiReq.KdKelompokPegawai = kelompokPegawai.KdKelompokPegawai
	}

	// Pengecekan Pendidikan Masuk
	fmt.Println("DEBUG : ", pegawaiReq.UuidPendidikanMasuk)
	if pegawaiReq.UuidPendidikanMasuk != "" {
		pendidikanMasuk, err := jenjangPendidikan.GetJenjangPendidikanByUUID(a, c.Request().Context(), pegawaiReq.UuidPendidikanMasuk)
		if err != nil {
			return model.PegawaiCreate{}, fmt.Errorf("error from repo jenis ijazah by uuid, %w", err)
		}
		pegawaiReq.IdPendidikanMasuk, _ = conv.Int(pendidikanMasuk.ID)
		pegawaiReq.KdPendidikanMasuk = pendidikanMasuk.KdJenjang
	}

	// Pengecekan Pangkat Golongan Pegawai
	if pegawaiReq.UuidGolongan != "" {
		pangkatPegawai, err := pangkatPegawaiRepo.GetPangkatGolonganPegawaiByUUID(a, pegawaiReq.UuidGolongan)
		if err != nil {
			return model.PegawaiCreate{}, fmt.Errorf("error from repo pangkat golongan pegawai by uuid, %w", err)
		}
		pegawaiReq.PegawaiFungsional.IdPangkatGolongan, _ = conv.Int(pangkatPegawai.ID)
		pegawaiReq.PegawaiFungsional.KdPangkatGolongan = pangkatPegawai.KdPangkat
	}

	// Pengecekan Jabatan Fungsional Yayasan
	if pegawaiReq.PegawaiFungsional.UuidJabatanFungsional != "" {
		jabatanFungsional, err := jabatanFungsionalRepo.GetJabatanFungsionalByUUID(a, c.Request().Context(), pegawaiReq.PegawaiFungsional.UuidJabatanFungsional)
		if err != nil {
			return model.PegawaiCreate{}, fmt.Errorf("error from repo pangkat golongan pegawai by uuid, %w", err)
		}
		pegawaiReq.PegawaiFungsional.IdJabatanFungsional, _ = conv.Int(jabatanFungsional.ID)
		pegawaiReq.PegawaiFungsional.KdJabatanFungsional = jabatanFungsional.KdJabatanFungsional
	}

	// Pengecekan Jenis Nomor Resgistrasi
	if pegawaiReq.PegawaiFungsional.UuidJenisNomorRegistrasi != "" {
		jenisNoRegis, err := jenisNoRegisRepo.GetJenisNoRegisByUUID(a, c.Request().Context(), pegawaiReq.PegawaiFungsional.UuidJenisNomorRegistrasi)
		if err != nil {
			return model.PegawaiCreate{}, fmt.Errorf("error from repo jenis nomor registrasi by uuid, %w", err)
		}
		pegawaiReq.PegawaiFungsional.IdJenisNomorRegistrasi, _ = conv.Int(jenisNoRegis.ID)
		pegawaiReq.PegawaiFungsional.KdJenisNomorRegistrasi = jenisNoRegis.KdJenisRegis
	}

	// Pengecekan Induk Kerja
	if pegawaiReq.UuidUnitKerja1 != "" {
		indukKerja, err := indukKerjaRepo.GetIndukKerjaByUUID(a, c.Request().Context(), pegawaiReq.UuidUnitKerja1)
		if err != nil {
			return model.PegawaiCreate{}, fmt.Errorf("error from repo induk kerja by uuid, %w", err)
		}
		pegawaiReq.IdUnitKerja1, _ = conv.Int(indukKerja.ID)
		pegawaiReq.KdUnit1 = indukKerja.KdUnit1
	}

	// Pengecekan Unit Kerja
	if pegawaiReq.UuidUnitKerja2 != "" {
		unitKerja, err := indukKerjaRepo.GetUnitKerjaByUUID(a, c.Request().Context(), pegawaiReq.UuidUnitKerja2)
		if err != nil {
			return model.PegawaiCreate{}, fmt.Errorf("error from repo unit kerja by uuid, %w", err)
		}
		pegawaiReq.IdUnitKerja2, _ = conv.Int(unitKerja.ID)
		pegawaiReq.KdUnit2 = unitKerja.KdUnit2
	}

	// Pengecekan Bagian Kerja
	if pegawaiReq.UuidUnitKerja3 != "" {
		bagianKerja, err := indukKerjaRepo.GetBagianKerjaByUUID(a, c.Request().Context(), pegawaiReq.UuidUnitKerja3)
		if err != nil {
			return model.PegawaiCreate{}, fmt.Errorf("error from repo bagian kerja by uuid, %w", err)
		}
		pegawaiReq.IdUnitKerja3, _ = conv.Int(bagianKerja.ID)
		pegawaiReq.KdUnit3 = bagianKerja.KdUnit3
	}

	// Pengecekan Lokasi Kerja
	if pegawaiReq.UuidLokasiKerja != "" {
		lokasiKerja, err := lokasiKerjaRepo.GetLokasiKerjaByUUID(a, c.Request().Context(), pegawaiReq.UuidLokasiKerja)
		if err != nil {
			return model.PegawaiCreate{}, fmt.Errorf("error from repo lokasi kerja by uuid, %w", err)
		}
		pegawaiReq.IdUnitKerjaLokasi, _ = conv.Int(lokasiKerja.ID)
		pegawaiReq.LokasiKerja = lokasiKerja.LokasiKerja
	}

	// Pengecekan Homebase Pddikti
	if pegawaiReq.PegawaiFungsional.UuidHomebasePddikti != "" {
		homebasePddikti, err := indukKerjaRepo.GetUnitKerjaByUUID(a, c.Request().Context(), pegawaiReq.PegawaiFungsional.UuidHomebasePddikti)
		if err != nil {
			return model.PegawaiCreate{}, fmt.Errorf("error from repo unit kerja by uuid, %w", err)
		}
		pegawaiReq.PegawaiFungsional.IdHomebasePddikti, _ = conv.Int(homebasePddikti.ID)
	}

	// Pengecekan Homebase UII
	if pegawaiReq.PegawaiFungsional.UuidHomebaseUii != "" {
		homebaseUUuidHomebaseUii, err := indukKerjaRepo.GetUnitKerjaByUUID(a, c.Request().Context(), pegawaiReq.PegawaiFungsional.UuidHomebaseUii)
		if err != nil {
			return model.PegawaiCreate{}, fmt.Errorf("error from repo unit kerja by uuid, %w", err)
		}
		pegawaiReq.PegawaiFungsional.IdHomebaseUii, _ = conv.Int(homebaseUUuidHomebaseUii.ID)
	}

	// Pengecekan Pangkat Golongan Ruang PNS
	if pegawaiReq.PegawaiPNS.UuidPangkatGolongan != "" {
		pangkatPNS, err := pangkatPegawaiRepo.GetPangkatGolonganPegawaiByUUID(a, pegawaiReq.PegawaiPNS.UuidPangkatGolongan)
		if err != nil {
			return model.PegawaiCreate{}, fmt.Errorf("error from repo pangkat golongan pegawai by uuid, %w", err)
		}
		pegawaiReq.PegawaiPNS.IdPangkatGolongan, _ = conv.Int(pangkatPNS.ID)
		pegawaiReq.PegawaiPNS.KdPangkatGolongan = pangkatPNS.KdPangkat
	}

	// Pengecekan Jabatan Fungsional Yayasan
	if pegawaiReq.PegawaiPNS.UuidJabatanFungsional != "" {
		jabatanPNS, err := jabatanFungsionalRepo.GetJabatanFungsionalByUUID(a, c.Request().Context(), pegawaiReq.PegawaiPNS.UuidJabatanFungsional)
		if err != nil {
			return model.PegawaiCreate{}, fmt.Errorf("error from repo pangkat golongan pegawai by uuid, %w", err)
		}
		pegawaiReq.PegawaiPNS.IdJabatanFungsional, _ = conv.Int(jabatanPNS.ID)
		pegawaiReq.PegawaiPNS.KdJabatanFungsional = jabatanPNS.KdJabatanFungsional
	}

	// Pengecekan Jenis PTT
	if pegawaiReq.PegawaiPNS.UuidJenisPtt != "" {
		jenisPTT, err := jenisPTTRepo.GetJenisPTTByUUID(a, c.Request().Context(), pegawaiReq.PegawaiPNS.UuidJenisPtt)
		if err != nil {
			return model.PegawaiCreate{}, fmt.Errorf("error from repo jenis pegawai tidak tetap by uuid, %w", err)
		}
		pegawaiReq.PegawaiPNS.IdJenisPtt, _ = conv.Int(jenisPTT.ID)
		pegawaiReq.PegawaiPNS.KdJenisPtt = jenisPTT.KdJenisPTT
	}

	// Pengecekan Status Pegawai
	if pegawaiReq.PegawaiFungsional.UuidStatusPegawaiAktif != "" {
		statusPegawaiAktif, err := statusPegawaiAktifRepo.GetStatusPegawaiAktifByUUID(a, c.Request().Context(), pegawaiReq.PegawaiFungsional.UuidStatusPegawaiAktif)
		if err != nil {
			return model.PegawaiCreate{}, fmt.Errorf("error from repo status pegawai aktif by uuid, %w", err)
		}
		pegawaiReq.PegawaiFungsional.IdStatusPegawaiAktif, _ = conv.Int(statusPegawaiAktif.ID)
		pegawaiReq.PegawaiFungsional.KdStatusPegawaiAktif = statusPegawaiAktif.KdStatusAktif
	}

	// Binding nilai request ke struct

	if pegawaiReq.PegawaiFungsional.MasaKerjaBawaanBulan != "" {
		a, _ := strconv.Atoi(pegawaiReq.PegawaiFungsional.MasaKerjaBawaanBulan)
		if a > 12 {
			return model.PegawaiCreate{}, fmt.Errorf("[ERROR] data bulan tidak valid")
		}
	}

	if pegawaiReq.PegawaiFungsional.MasaKerjaGajiBulan != "" {
		a, _ := strconv.Atoi(pegawaiReq.PegawaiFungsional.MasaKerjaGajiBulan)
		if a > 12 {
			return model.PegawaiCreate{}, fmt.Errorf("[ERROR] data bulan tidak valid")
		}
	}

	// if pegawaiReq.PegawaiFungsional.MasaKerjaTotalBulan != "" {
	// 	a, _ := strconv.Atoi(pegawaiReq.PegawaiFungsional.MasaKerjaTotalBulan)
	// 	if a > 12 {
	// 		return model.PegawaiCreate{}, fmt.Errorf("error data bulan tidak valid")
	// 	}
	// }
	if pegawaiReq.PegawaiFungsional.NomorSertifikasi != "" {
		if len(pegawaiReq.PegawaiFungsional.NomorSertifikasi) > 20 {
			return model.PegawaiCreate{}, fmt.Errorf("error nomor sertifikasi tidak valid")
		}
	}
	if pegawaiReq.PegawaiFungsional.NomorRegistrasi != "" {
		if len(pegawaiReq.PegawaiFungsional.NomorRegistrasi) > 10 {
			return model.PegawaiCreate{}, fmt.Errorf("error nomor registrasi tidak valid")
		}
	}
	if pegawaiReq.PegawaiFungsional.NomorSkPertama != "" {
		if len(pegawaiReq.PegawaiFungsional.NomorSkPertama) > 10 {
			return model.PegawaiCreate{}, fmt.Errorf("error nomor sk pertama tidak valid")
		}
	}
	if pegawaiReq.PegawaiPNS.NipPns != "" {
		if len(pegawaiReq.PegawaiPNS.NipPns) < 18 {
			return model.PegawaiCreate{}, fmt.Errorf("error nip pns tidak valid")
		}
	}
	if pegawaiReq.PegawaiPNS.NoKartuPegawai != "" {
		if len(pegawaiReq.PegawaiPNS.NoKartuPegawai) > 18 {
			return model.PegawaiCreate{}, fmt.Errorf("error nomor kartu pegawai tidak valid")
		}
	}
	if pegawaiReq.PegawaiPNS.MasaKerjaBulan != "" {
		a, _ := strconv.Atoi(pegawaiReq.PegawaiPNS.MasaKerjaBulan)
		if a > 12 {
			return model.PegawaiCreate{}, fmt.Errorf("error data bulan tidak valid")
		}
	}

	if pegawaiReq.DetailProfesi != "" {
		fmt.Println("Ini Detail Profese : ", pegawaiReq.DetailProfesi)
		if len(pegawaiReq.DetailProfesi) > 35 {
			return model.PegawaiCreate{}, fmt.Errorf("error detail terlalu panjang")
		}
	}

	if pegawaiReq.Nik != "" {
		if len(pegawaiReq.Nik) > 9 {
			return model.PegawaiCreate{}, fmt.Errorf("error NIK terlalu panjang")
		}
	}

	pegawaiReq.UserUpdate = user
	pegawaiReq.PegawaiFungsional.UserUpdate = user

	return pegawaiReq, nil
}
