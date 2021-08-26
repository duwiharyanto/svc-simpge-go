package usecase

import (
	"fmt"
	"strconv"
	"svc-insani-go/app"
	"svc-insani-go/app/helper"
	"svc-insani-go/modules/v1/pegawai/model"
	"svc-insani-go/modules/v1/pegawai/repo"

	detailProfesiRepo "svc-insani-go/modules/v1/master-detail-profesi/repo"
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

	"github.com/labstack/echo/v4"

	guuid "github.com/google/uuid"
)

func ValidateUpdatePegawaiByUUID(a *app.App, c echo.Context) (model.PegawaiUpdate, error) {
	uuidPegawai := c.Param("uuidPegawai")
	if uuidPegawai == "" {
		return model.PegawaiUpdate{}, fmt.Errorf("uuid pegawai tidak boleh kosong")
	}

	pegawai, err := repo.GetPegawaiByUUID(a, uuidPegawai)
	if err != nil {
		return model.PegawaiUpdate{}, fmt.Errorf("error from repo get uuid pegawai: %w", err)
	}

	pegawaiOld, err := repo.GetOldPegawai(a, c.Request().Context(), uuidPegawai)
	if err != nil {
		return model.PegawaiUpdate{}, fmt.Errorf("error from get old pegawai: %w", err)
	}

	user := c.Request().Header.Get("X-Member")

	pegawaiReq := &model.PegawaiUpdate{}

	err = c.Bind(pegawaiReq)
	if err != nil {
		fmt.Printf("[ERROR] binding requestpegawai , %s\n", err.Error())
	}

	pegawaiReq.Uuid = uuidPegawai
	pegawaiReq.Id = pegawai.Id

	//Pengecekan Jenis Pegawai
	if pegawaiReq.UuidJenisPegawai != "" {
		jenisPegawai, err := jenisPegawaiRepo.GetJenisPegawaiByUUID(a, pegawaiReq.UuidJenisPegawai)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo jenis pegawai by uuid: %w", err)
		}
		pegawaiOld.IdJenisPegawai = jenisPegawai.ID
		pegawaiOld.KdJenisPegawai = jenisPegawai.KDJenisPegawai
	}

	// Pengecekan Status Pegawai
	if pegawaiReq.UuidStatusPegawai != "" {
		statusPegawai, err := statusPegawaiRepo.GetStatusPegawaiByUUID(a, c.Request().Context(), pegawaiReq.UuidStatusPegawai)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo status pegawai by uuid: %w", err)
		}
		pegawaiOld.IdStatusPegawai = statusPegawai.ID
		pegawaiOld.KdStatusPegawai = statusPegawai.KDStatusPegawai
	}

	// Pengecekan Kelompok Pegawai
	if pegawaiReq.UuidKelompokPegawai != "" {
		kelompokPegawai, err := kelompokPegawaiRepo.GetKelompokPegawaiByUUID(a, c.Request().Context(), pegawaiReq.UuidKelompokPegawai)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo kelompok pegawai by uuid: %w", err)
		}
		pegawaiOld.IdKelompokPegawai = kelompokPegawai.ID
		pegawaiOld.KdKelompokPegawai = kelompokPegawai.KdKelompokPegawai
	}

	// Pengecekan Detail Profesi
	if pegawaiReq.UuidDetailProfesi != "" {
		detailProfesi, err := detailProfesiRepo.GetDetailProfesiByUUID(a, c.Request().Context(), pegawaiReq.UuidDetailProfesi)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo detail profesi by uuid: %w", err)
		}
		pegawaiOld.IdDetailProfesi = detailProfesi.ID
	}

	// Pengecekan Ijazah Pendidikan Masuk
	if pegawaiReq.UuidPendidikanMasuk != "" {
		pendidikanMasuk, err := jenjangPendidikan.GetJenjangPendidikanByUUID(a, c.Request().Context(), pegawaiReq.UuidPendidikanMasuk)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo jenis ijazah pendidikan masuk by uuid: %w", err)
		}
		pegawaiOld.IdPendidikanMasuk = pendidikanMasuk.ID
		pegawaiOld.KdPendidikanMasuk = pendidikanMasuk.KdPendidikanSimpeg
	}

	// Pengecekan Ijazah Pendidikan Terakhir
	if pegawaiReq.UuidPendidikanTerakhir != "" {
		pendidikanTerakhir, err := jenjangPendidikan.GetJenjangPendidikanByUUID(a, c.Request().Context(), pegawaiReq.UuidPendidikanTerakhir)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo jenis ijazah pendidikan terakhir by uuid: %w", err)
		}
		pegawaiOld.IdPendidikanTerakhir = pendidikanTerakhir.ID
		pegawaiOld.KdPendidikanTerakhir = pendidikanTerakhir.KdPendidikanSimpeg
	}

	// Pengecekan Pangkat Golongan Pegawai
	if pegawaiReq.UuidGolongan != "" {
		pangkatPegawai, err := pangkatPegawaiRepo.GetPangkatGolonganPegawaiByUUID(a, pegawaiReq.UuidGolongan)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo pangkat golongan pegawai by uuid: %w", err)
		}
		pegawaiOld.PegawaiFungsional.IdPangkatGolongan = pangkatPegawai.ID
		pegawaiOld.PegawaiFungsional.KdPangkatGolongan = pangkatPegawai.KdPangkat
	}

	// Pengecekan Jabatan Fungsional Yayasan
	if pegawaiReq.PegawaiFungsional.UuidJabatanFungsional != "" {
		jabatanFungsional, err := jabatanFungsionalRepo.GetJabatanFungsionalByUUID(a, c.Request().Context(), pegawaiReq.PegawaiFungsional.UuidJabatanFungsional)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo pangkat golongan pegawai by uuid: %w", err)
		}
		pegawaiOld.PegawaiFungsional.IdJabatanFungsional = jabatanFungsional.ID
		pegawaiOld.PegawaiFungsional.KdJabatanFungsional = jabatanFungsional.KdJabatanFungsional
	}

	// Pengecekan Jenis Nomor Resgistrasi
	if pegawaiReq.PegawaiFungsional.UuidJenisNomorRegistrasi != "" {
		jenisNoRegis, err := jenisNoRegisRepo.GetJenisNoRegisByUUID(a, c.Request().Context(), pegawaiReq.PegawaiFungsional.UuidJenisNomorRegistrasi)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo jenis nomor registrasi by uuid: %w", err)
		}
		pegawaiOld.PegawaiFungsional.IdJenisNomorRegistrasi = jenisNoRegis.ID
		pegawaiOld.PegawaiFungsional.KdJenisNomorRegistrasi = jenisNoRegis.KdJenisRegis
	}

	// Pengecekan Induk Kerja
	if pegawaiReq.UuidUnitKerja1 != "" {
		indukKerja, err := indukKerjaRepo.GetIndukKerjaByUUID(a, c.Request().Context(), pegawaiReq.UuidUnitKerja1)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo induk kerja by uuid: %w", err)
		}
		pegawaiOld.IdUnitKerja1 = indukKerja.ID
		pegawaiOld.KdUnit1 = indukKerja.KdUnit1
	}

	// Pengecekan Unit Kerja
	if pegawaiReq.UuidUnitKerja2 != "" {
		unitKerja, err := indukKerjaRepo.GetUnitKerjaByUUID(a, c.Request().Context(), pegawaiReq.UuidUnitKerja2)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo unit kerja by uuid: %w", err)
		}
		pegawaiOld.IdUnitKerja2 = unitKerja.ID
		pegawaiOld.KdUnit2 = unitKerja.KdUnit2
	}

	// Pengecekan Bagian Kerja
	if pegawaiReq.UuidUnitKerja3 != "" {
		bagianKerja, err := indukKerjaRepo.GetBagianKerjaByUUID(a, c.Request().Context(), pegawaiReq.UuidUnitKerja3)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo bagian kerja by uuid: %w", err)
		}
		pegawaiOld.IdUnitKerja3 = bagianKerja.ID
		pegawaiOld.KdUnit3 = bagianKerja.KdUnit3
	}

	// Pengecekan Lokasi Kerja
	if pegawaiReq.UuidLokasiKerja != "" {
		lokasiKerja, err := lokasiKerjaRepo.GetLokasiKerjaByUUID(a, c.Request().Context(), pegawaiReq.UuidLokasiKerja)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo lokasi kerja by uuid: %w", err)
		}
		pegawaiOld.IdUnitKerjaLokasi = lokasiKerja.ID
		pegawaiOld.LokasiKerja = lokasiKerja.LokasiKerja
	}

	// Pengecekan Homebase Pddikti
	if pegawaiReq.PegawaiFungsional.UuidHomebasePddikti != "" {
		homebasePddikti, err := indukKerjaRepo.GetUnitKerjaByUUID(a, c.Request().Context(), pegawaiReq.PegawaiFungsional.UuidHomebasePddikti)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo unit kerja by uuid: %w", err)
		}
		pegawaiOld.PegawaiFungsional.IdHomebasePddikti = homebasePddikti.ID
	}

	// Pengecekan Homebase UII
	if pegawaiReq.PegawaiFungsional.UuidHomebaseUii != "" {
		homebaseUUuidHomebaseUii, err := indukKerjaRepo.GetUnitKerjaByUUID(a, c.Request().Context(), pegawaiReq.PegawaiFungsional.UuidHomebaseUii)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo unit kerja by uuid: %w", err)
		}
		pegawaiOld.PegawaiFungsional.IdHomebaseUii = homebaseUUuidHomebaseUii.ID
	}

	// Pengecekan Pangkat Golongan Ruang PNS
	if pegawaiReq.PegawaiPNS.UuidPangkatGolongan != "" {
		pangkatPNS, err := pangkatPegawaiRepo.GetPangkatGolonganPegawaiByUUID(a, pegawaiReq.PegawaiPNS.UuidPangkatGolongan)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo pangkat golongan pegawai by uuid: %w", err)
		}
		pegawaiOld.PegawaiPNS.IdPangkatGolongan = pangkatPNS.ID
		pegawaiOld.PegawaiPNS.KdPangkatGolongan = pangkatPNS.KdPangkat
	}

	// Pengecekan Jabatan Fungsional Yayasan
	if pegawaiReq.PegawaiPNS.UuidJabatanFungsional != "" {
		jabatanPNS, err := jabatanFungsionalRepo.GetJabatanFungsionalByUUID(a, c.Request().Context(), pegawaiReq.PegawaiPNS.UuidJabatanFungsional)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo pangkat golongan pegawai by uuid: %w", err)
		}
		pegawaiOld.PegawaiPNS.IdJabatanFungsional = jabatanPNS.ID
		pegawaiOld.PegawaiPNS.KdJabatanFungsional = jabatanPNS.KdJabatanFungsional
	}

	// Pengecekan Jenis PTT
	if pegawaiReq.PegawaiPNS.UuidJenisPtt != "" {
		jenisPTT, err := jenisPTTRepo.GetJenisPTTByUUID(a, c.Request().Context(), pegawaiReq.PegawaiPNS.UuidJenisPtt)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo jenis pegawai tidak tetap by uuid: %w", err)
		}
		pegawaiOld.PegawaiPNS.IdJenisPtt = jenisPTT.ID
		pegawaiOld.PegawaiPNS.KdJenisPtt = jenisPTT.KdJenisPTT
	}

	// Pengecekan Status Pegawai
	if pegawaiReq.PegawaiFungsional.UuidStatusPegawaiAktif != "" {
		statusPegawaiAktif, err := statusPegawaiAktifRepo.GetStatusPegawaiAktifByUUID(a, c.Request().Context(), pegawaiReq.PegawaiFungsional.UuidStatusPegawaiAktif)
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo status pegawai aktif by uuid: %w", err)
		}
		pegawaiOld.PegawaiFungsional.IdStatusPegawaiAktif = statusPegawaiAktif.ID
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
			return model.PegawaiUpdate{}, fmt.Errorf("Bulan masa kerja bawaan maksimal 12")
		}
		pegawaiOld.PegawaiFungsional.MasaKerjaBawaanBulan = pegawaiReq.PegawaiFungsional.MasaKerjaBawaanBulan
	}
	if pegawaiReq.PegawaiFungsional.MasaKerjaGajiTahun != "" {
		pegawaiOld.PegawaiFungsional.MasaKerjaGajiTahun = pegawaiReq.PegawaiFungsional.MasaKerjaGajiTahun
	}
	if pegawaiReq.PegawaiFungsional.MasaKerjaGajiBulan != "" {
		a, _ := strconv.Atoi(pegawaiReq.PegawaiFungsional.MasaKerjaGajiBulan)
		if a > 12 {
			return model.PegawaiUpdate{}, fmt.Errorf("Bulan masa kerja gaji maksimal 12")
		}
		pegawaiOld.PegawaiFungsional.MasaKerjaGajiBulan = pegawaiReq.PegawaiFungsional.MasaKerjaGajiBulan
	}

	if pegawaiReq.PegawaiFungsional.AngkaKredit != "" {
		pegawaiOld.PegawaiFungsional.AngkaKredit = pegawaiReq.PegawaiFungsional.AngkaKredit
	}
	if pegawaiReq.PegawaiFungsional.NomorSertifikasi != "" {
		if len(pegawaiReq.PegawaiFungsional.NomorSertifikasi) > 20 {
			return model.PegawaiUpdate{}, fmt.Errorf("Panjang karakter nomor sertifikasi maksimal 20")
		}
		pegawaiOld.PegawaiFungsional.NomorSertifikasi = pegawaiReq.PegawaiFungsional.NomorSertifikasi
	}
	if pegawaiReq.PegawaiFungsional.NomorRegistrasi != "" {
		if len(pegawaiReq.PegawaiFungsional.NomorRegistrasi) > 10 {
			return model.PegawaiUpdate{}, fmt.Errorf("Panjang karakter nomor registrasi maksimal 10")
		}
		pegawaiOld.PegawaiFungsional.NomorRegistrasi = pegawaiReq.PegawaiFungsional.NomorRegistrasi
	}
	if pegawaiReq.PegawaiFungsional.NomorSkPertama != "" {
		if len(pegawaiReq.PegawaiFungsional.NomorSkPertama) > 30 {
			return model.PegawaiUpdate{}, fmt.Errorf("Panjang karakter nomor sk pertama maksimal 30")
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
		if len(pegawaiReq.PegawaiPNS.NipPns) != 18 {
			return model.PegawaiUpdate{}, fmt.Errorf("Panjang karakter NIP PNS hanya boleh 18")
		}
		pegawaiOld.PegawaiPNS.NipPns = pegawaiReq.PegawaiPNS.NipPns
	}
	if pegawaiReq.PegawaiPNS.NoKartuPegawai != "" {
		if len(pegawaiReq.PegawaiPNS.NoKartuPegawai) != 18 {
			return model.PegawaiUpdate{}, fmt.Errorf("Panjang karakter nomor kartu pegawai hanya boleh 18")
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
			return model.PegawaiUpdate{}, fmt.Errorf("Masa kerja bulan maksimal 12")
		}
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
	pegawaiOld.PegawaiPNS.UserUpdate = user
	return pegawaiOld, nil
}

const (
	statusActive = "AKT"
)

func PrepareCreateSimpeg(a *app.App, c echo.Context) (model.PegawaiCreate, error) {
	pegawai, err := ValidateCreatePegawai(a, c)
	if err != nil {
		return model.PegawaiCreate{}, err
	}
	userUpdate := c.Request().Header.Get("X-Member")
	statusPegawaiAktif, err := statusPegawaiAktifRepo.GetStatusPegawaiAktifByCode(a, c.Request().Context(), statusActive)
	if err != nil {
		return model.PegawaiCreate{}, fmt.Errorf("error from repo status pegawai aktif by uuid: %w", err)
	}
	if statusPegawaiAktif == nil {
		return model.PegawaiCreate{}, fmt.Errorf("%w", fmt.Errorf("status pegawai aktif tidak ditemukan"))
	}
	pegawai.PegawaiFungsional.IdStatusPegawaiAktif = statusPegawaiAktif.ID
	pegawai.PegawaiFungsional.KdStatusPegawaiAktif = statusPegawaiAktif.KdStatusAktif

	pegawai.UserUpdate = userUpdate
	pegawai.UserInput = userUpdate
	pegawai.PegawaiFungsional.UserInput = userUpdate
	pegawai.PegawaiFungsional.UserUpdate = userUpdate
	pegawai.PegawaiPNS.UserInput = userUpdate
	pegawai.PegawaiPNS.UserUpdate = userUpdate

	uuidPersonal := c.Param("uuidPegawai")
	if pegawai.UuidPersonal != "" {
		uuidPersonal = pegawai.UuidPersonal
	}
	if uuidPersonal == "" {
		return model.PegawaiCreate{}, fmt.Errorf("uuid personal tidak boleh kosong")
	}

	personal, err := personalRepo.GetPersonalByUuid(a, c.Request().Context(), uuidPersonal)
	if err != nil {
		return model.PegawaiCreate{}, fmt.Errorf("error search personal")
	}
	if personal == nil {
		return model.PegawaiCreate{}, fmt.Errorf("personal tidak ditemukan")
	}
	if personal.Pegawai.PegawaiFungsional.StatusPegawaiAktif.IsActive() {
		return model.PegawaiCreate{}, fmt.Errorf("tidak dapat menambah data dari pegawai yang sudah aktif")
	}

	pegawai.IdPersonalDataPribadi = personal.Id
	pegawai.Nama = personal.NamaLengkap
	pegawai.NikKtp = personal.NikKtp
	pegawai.TglLahir = personal.TglLahir
	pegawai.TempatLahir = personal.TempatLahir
	pegawai.IdAgama = personal.Agama.Id
	pegawai.KdAgama = personal.Agama.KdItem
	pegawai.JenisKelamin = personal.JenisKelamin
	pegawai.IdGolonganDarah = personal.GolonganDarah.Id
	pegawai.KdGolonganDarah = personal.GolonganDarah.GolonganDarah
	pegawai.IdStatusPerkawinan = personal.StatusPernikahan.Id
	pegawai.KdStatusPerkawinan = personal.StatusPernikahan.KdStatus
	pegawai.GelarDepan = personal.GelarDepan
	pegawai.GelarBelakang = personal.GelarBelakang
	pegawai.Uuid = guuid.New().String()

	return pegawai, nil
}

func ValidateCreatePegawai(a *app.App, c echo.Context) (model.PegawaiCreate, error) {
	pegawaiReq := model.PegawaiCreate{}
	err := c.Bind(&pegawaiReq)
	if err != nil {
		fmt.Printf("[WARNING] binding create pegawai request: %s\n", err.Error())
	}

	switch {
	case len(pegawaiReq.Nik) != 9 || !helper.IsNumber(pegawaiReq.Nik):
		return model.PegawaiCreate{}, fmt.Errorf("nik wajib diisi berupa 9 digit angka")
	case pegawaiReq.UuidKelompokPegawai == "":
		return model.PegawaiCreate{}, fmt.Errorf("uuid_kelompok_pegawai tidak boleh kosong")
	case pegawaiReq.UuidUnitKerja2 == "":
		return model.PegawaiCreate{}, fmt.Errorf("uuid_unit_kerja tidak boleh kosong")
	case pegawaiReq.PegawaiFungsional.TmtSkPertama == nil,
		pegawaiReq.PegawaiFungsional.TmtSkPertama != nil &&
			!helper.IsDateFormatValid("2006-01-02", *pegawaiReq.PegawaiFungsional.TmtSkPertama):
		return model.PegawaiCreate{}, fmt.Errorf("tmt_sk_pertama wajib diisi dengan format yyyy-mm-dd")
	case pegawaiReq.UuidLokasiKerja == "":
		return model.PegawaiCreate{}, fmt.Errorf("uuid_lokasi_kerja tidak boleh kosong")
	}

	// Pengecekan Kelompok Pegawai
	kelompokPegawai, err := kelompokPegawaiRepo.GetKelompokPegawaiByUUID(a, c.Request().Context(), pegawaiReq.UuidKelompokPegawai)
	if err != nil {
		return model.PegawaiCreate{}, fmt.Errorf("error from repo kelompok pegawai by uuid: %w", err)
	}
	if kelompokPegawai == nil {
		return model.PegawaiCreate{}, fmt.Errorf("uuid_kelompok_pegawai tidak ditemukan")
	}
	pegawaiReq.IdKelompokPegawai = kelompokPegawai.ID
	pegawaiReq.KdKelompokPegawai = kelompokPegawai.KdKelompokPegawai
	pegawaiReq.IdJenisPegawai = kelompokPegawai.JenisPegawai.ID
	pegawaiReq.KdJenisPegawai = kelompokPegawai.JenisPegawai.KDJenisPegawai
	pegawaiReq.IdStatusPegawai = kelompokPegawai.StatusPegawai.ID
	pegawaiReq.KdStatusPegawai = kelompokPegawai.StatusPegawai.KDStatusPegawai
	// Pengecekan Unit Kerja
	unitKerja, err := indukKerjaRepo.GetUnitKerjaByUUID(a, c.Request().Context(), pegawaiReq.UuidUnitKerja2)
	if err != nil {
		return model.PegawaiCreate{}, fmt.Errorf("error from repo unit kerja by uuid: %w", err)
	}
	if unitKerja == nil {
		return model.PegawaiCreate{}, fmt.Errorf("uuid_unit_kerja tidak ditemukan")
	}
	pegawaiReq.IdUnitKerja2 = unitKerja.ID
	pegawaiReq.KdUnit2 = unitKerja.KdUnit2
	pegawaiReq.IdUnitKerja1 = unitKerja.Unit1.ID
	pegawaiReq.KdUnit1 = unitKerja.Unit1.KdUnit1

	// Pengecekan Bagian Kerja
	if pegawaiReq.UuidUnitKerja3 != "" {
		bagianKerja, err := indukKerjaRepo.GetBagianKerjaByUUID(a, c.Request().Context(), pegawaiReq.UuidUnitKerja3)
		if err != nil {
			return model.PegawaiCreate{}, fmt.Errorf("error from repo bagian kerja by uuid: %w", err)
		}
		if bagianKerja == nil {
			return model.PegawaiCreate{}, fmt.Errorf("uuid_bagian_kerja tidak ditemukan")
		}
		pegawaiReq.IdUnitKerja3 = bagianKerja.ID
		pegawaiReq.KdUnit3 = bagianKerja.KdUnit3
	}

	// Pengecekan Lokasi Kerja
	lokasiKerja, err := lokasiKerjaRepo.GetLokasiKerjaByUUID(a, c.Request().Context(), pegawaiReq.UuidLokasiKerja)
	if err != nil {
		return model.PegawaiCreate{}, fmt.Errorf("error from repo lokasi kerja by uuid: %w", err)
	}
	if lokasiKerja == nil {
		return model.PegawaiCreate{}, fmt.Errorf("uuid_lokasi_kerja tidak ditemukan")
	}
	pegawaiReq.IdUnitKerjaLokasi = lokasiKerja.ID
	pegawaiReq.LokasiKerja = lokasiKerja.LokasiKerja

	return pegawaiReq, nil
}
