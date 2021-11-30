package usecase

import (
	"fmt"
	"strconv"
	"svc-insani-go/app"
	"svc-insani-go/app/helper"
	detailProfesiRepo "svc-insani-go/modules/v1/master-detail-profesi/repo"
	jabatanFungsionalRepo "svc-insani-go/modules/v1/master-jabatan-fungsional/repo"
	jenisNoRegisRepo "svc-insani-go/modules/v1/master-jenis-nomor-registrasi/repo"
	jenisPTTRepo "svc-insani-go/modules/v1/master-jenis-pegawai-tidak-tetap/repo"
	jenisPegawaiRepo "svc-insani-go/modules/v1/master-jenis-pegawai/repo"
	kelompokPegawaiRepo "svc-insani-go/modules/v1/master-kelompok-pegawai/repo"
	lokasiKerjaRepo "svc-insani-go/modules/v1/master-lokasi-kerja/repo"
	indukKerjaRepo "svc-insani-go/modules/v1/master-organisasi/repo"
	pangkatPegawaiRepo "svc-insani-go/modules/v1/master-pangkat-golongan-pegawai/repo"
	masterPendidikan "svc-insani-go/modules/v1/master-pendidikan/repo"
	statusPegawaiAktifRepo "svc-insani-go/modules/v1/master-status-pegawai-aktif/repo"
	statusPegawaiRepo "svc-insani-go/modules/v1/master-status-pegawai/repo"
	"svc-insani-go/modules/v1/pegawai/model"
	"svc-insani-go/modules/v1/pegawai/repo"
	personalRepo "svc-insani-go/modules/v1/personal/repo"
	"time"

	guuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
	ptr "github.com/openlyinc/pointy"
)

func ValidateUpdatePegawaiByUUID(a *app.App, c echo.Context) (model.PegawaiUpdate, error) {
	uuidPegawai := c.Param("uuidPegawai")
	if uuidPegawai == "" {
		return model.PegawaiUpdate{}, fmt.Errorf("uuid pegawai tidak boleh kosong")
	}

	pegawai, err := repo.GetOldPegawai(a, c.Request().Context(), uuidPegawai)
	if err != nil {
		return model.PegawaiUpdate{}, fmt.Errorf("error from get old pegawai: %w", err)
	}

	user := c.Request().Header.Get("X-Member")

	pegawaiReq := &model.PegawaiUpdate{}

	err = c.Bind(pegawaiReq)
	if err != nil {
		fmt.Printf("[ERROR] binding request pegawai , %s\n", err.Error())
	}

	pegawaiReq.Uuid = ptr.String(uuidPegawai)
	pegawaiReq.Id = pegawai.Id

	if pegawaiReq.GelarDepan != nil {
		if ptr.StringValue(pegawaiReq.GelarDepan, "") == "" {
			pegawai.GelarDepan = nil
		} else {
			pegawai.GelarDepan = pegawaiReq.GelarDepan
		}
	}

	if pegawaiReq.GelarBelakang != nil {
		if ptr.StringValue(pegawaiReq.GelarBelakang, "") == "" {
			pegawai.GelarBelakang = nil
		} else {
			pegawai.GelarBelakang = pegawaiReq.GelarBelakang
		}
	}

	//Pengecekan Jenis Pegawai
	if ptr.StringValue(pegawaiReq.UuidJenisPegawai, "") != "" {
		jenisPegawai, err := jenisPegawaiRepo.GetJenisPegawaiByUUID(a, ptr.StringValue(pegawaiReq.UuidJenisPegawai, ""))
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo jenis pegawai by uuid: %w", err)
		}
		pegawai.IdJenisPegawai = ptr.Uint64(jenisPegawai.ID)
		pegawai.KdJenisPegawai = ptr.String(jenisPegawai.KDJenisPegawai)
	}

	// Pengecekan Status Pegawai
	if ptr.StringValue(pegawaiReq.UuidStatusPegawai, "") != "" {
		statusPegawai, err := statusPegawaiRepo.GetStatusPegawaiByUUID(a, c.Request().Context(), ptr.StringValue(pegawaiReq.UuidStatusPegawai, ""))
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo status pegawai by uuid: %w", err)
		}
		pegawai.IdStatusPegawai = ptr.Uint64(statusPegawai.ID)
		pegawai.KdStatusPegawai = ptr.String(statusPegawai.KDStatusPegawai)
	}

	// Pengecekan Kelompok Pegawai
	if ptr.StringValue(pegawaiReq.UuidKelompokPegawai, "") != "" {
		kelompokPegawai, err := kelompokPegawaiRepo.GetKelompokPegawaiByUUID(a, c.Request().Context(), ptr.StringValue(pegawaiReq.UuidKelompokPegawai, ""))
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo kelompok pegawai by uuid: %w", err)
		}
		pegawai.IdKelompokPegawai = ptr.Uint64(kelompokPegawai.ID)
		pegawai.KdKelompokPegawai = ptr.String(kelompokPegawai.KdKelompokPegawai)
	}

	// Pengecekan Kelompok Pegawai Payroll
	if ptr.StringValue(pegawaiReq.UuidKelompokPegawaiPayroll, "") != "" {
		kelompokPegawai, err := kelompokPegawaiRepo.GetKelompokPegawaiByUUID(a, c.Request().Context(), ptr.StringValue(pegawaiReq.UuidKelompokPegawaiPayroll, ""))
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo kelompok pegawai by uuid: %w", err)
		}
		pegawai.IdKelompokPegawaiPayroll = ptr.Uint64(kelompokPegawai.ID)
		pegawai.KdKelompokPegawaiPayroll = ptr.String(kelompokPegawai.KdKelompokPegawai)
	}

	// Pengecekan Detail Profesi
	if ptr.StringValue(pegawaiReq.UuidDetailProfesi, "") != "" {
		detailProfesi, err := detailProfesiRepo.GetDetailProfesiByUUID(a, c.Request().Context(), ptr.StringValue(pegawaiReq.UuidDetailProfesi, ""))
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo detail profesi by uuid: %w", err)
		}
		pegawai.IdDetailProfesi = ptr.Uint64(detailProfesi.ID)
	}

	// Pengecekan Jenjang dari pendidikan tertinggi diakui
	if ptr.StringValue(pegawaiReq.UuidPendidikanMasuk, "") != "" {
		pendidikanMasuk, err := masterPendidikan.GetJenjangPendidikanByUUID(a, c.Request().Context(), ptr.StringValue(pegawaiReq.UuidPendidikanMasuk, ""))
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo jenis ijazah pendidikan masuk by uuid: %w", err)
		}
		pegawai.IdPendidikanMasuk = ptr.Uint64(pendidikanMasuk.ID)
		pegawai.KdPendidikanMasuk = ptr.String(pendidikanMasuk.KdPendidikanSimpeg)
		pegawai.UuidPendidikanMasuk = pegawaiReq.UuidPendidikanMasuk
	}

	// Pengecekan jenjang dari Pendidikan Terakhir
	if ptr.StringValue(pegawaiReq.UuidPendidikanTerakhir, "") != "" {
		pendidikanTerakhir, err := masterPendidikan.GetJenjangPendidikanByUUID(a, c.Request().Context(), ptr.StringValue(pegawaiReq.UuidPendidikanTerakhir, ""))
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo jenis ijazah pendidikan terakhir by uuid: %w", err)
		}
		pegawai.IdPendidikanTerakhir = ptr.Uint64(pendidikanTerakhir.ID)
		pegawai.KdPendidikanTerakhir = ptr.String(pendidikanTerakhir.KdPendidikanSimpeg)
	}

	// Pengecekan jenis pendidikan tertinggi diakui
	if ptr.StringValue(pegawaiReq.UuidStatusPendidikanMasuk, "") != "" {
		data, err := masterPendidikan.GetJenjangPendidikanDetailByUUID(a, c.Request().Context(), ptr.StringValue(pegawaiReq.UuidStatusPendidikanMasuk, ""))
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo get jenjang pendidikan detail by uuid status pendidikan masuk: %w", err)
		}
		pegawai.IdStatusPendidikanMasuk = ptr.Uint64(data.ID)
		pegawai.KdStatusPendidikanMasuk = ptr.String(data.KdDetail)
	}

	// Pengecekan jenis pendidikan terakhir
	if ptr.StringValue(pegawaiReq.UuidJenisPendidikan, "") != "" {
		data, err := masterPendidikan.GetJenjangPendidikanDetailByUUID(a, c.Request().Context(), ptr.StringValue(pegawaiReq.UuidJenisPendidikan, ""))
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo get jenjang pendidikan detail by uuid jenis pendidikan: %w", err)
		}
		pegawai.IdJenisPendidikan = ptr.Uint64(data.ID)
		pegawai.KdJenisPendidikan = ptr.String(data.KdDetail)
	}

	// Pengecekan Pangkat Golongan Pegawai
	if ptr.StringValue(pegawaiReq.UuidGolongan, "") != "" {
		pangkatPegawai, err := pangkatPegawaiRepo.GetPangkatGolonganPegawaiByUUID(a, ptr.StringValue(pegawaiReq.UuidGolongan, ""))
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo pangkat golongan pegawai by uuid: %w", err)
		}
		pegawai.PegawaiFungsional.IdPangkatGolongan = ptr.Uint64(pangkatPegawai.ID)
		pegawai.PegawaiFungsional.KdPangkatGolongan = ptr.String(pangkatPegawai.KdPangkat)
		pegawai.IdGolongan = ptr.Uint64(pangkatPegawai.IdGolongan)
		pegawai.KdGolongan = ptr.String(pangkatPegawai.KdGolongan)
		pegawai.IdRuang = ptr.Uint64(pangkatPegawai.IdRuang)
		pegawai.KdRuang = ptr.String(pangkatPegawai.KdRuang)
	}

	// Pengecekan Jabatan Fungsional Yayasan
	if ptr.StringValue(pegawaiReq.PegawaiFungsional.UuidJabatanFungsional, "") != "" {
		jabatanFungsional, err := jabatanFungsionalRepo.GetJabatanFungsionalByUUID(a, c.Request().Context(), ptr.StringValue(pegawaiReq.PegawaiFungsional.UuidJabatanFungsional, ""))
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo pangkat golongan pegawai by uuid: %w", err)
		}
		pegawai.PegawaiFungsional.IdJabatanFungsional = ptr.Uint64(jabatanFungsional.ID)
		pegawai.PegawaiFungsional.KdJabatanFungsional = ptr.String(jabatanFungsional.KdJabatanFungsional)
	}

	// Pengecekan Jenis Nomor Registrasi
	if ptr.StringValue(pegawaiReq.PegawaiFungsional.UuidJenisNomorRegistrasi, "") != "" {
		jenisNoRegis, err := jenisNoRegisRepo.GetJenisNoRegisByUUID(a, c.Request().Context(), ptr.StringValue(pegawaiReq.PegawaiFungsional.UuidJenisNomorRegistrasi, ""))
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo jenis nomor registrasi by uuid: %w", err)
		}
		pegawai.PegawaiFungsional.IdJenisNomorRegistrasi = ptr.Uint64(jenisNoRegis.ID)
		pegawai.PegawaiFungsional.KdJenisNomorRegistrasi = ptr.String(jenisNoRegis.KdJenisRegis)
	}

	// Pengecekan Induk Kerja
	if ptr.StringValue(pegawaiReq.UuidUnitKerja1, "") != "" {
		indukKerja, err := indukKerjaRepo.GetIndukKerjaByUUID(a, c.Request().Context(), ptr.StringValue(pegawaiReq.UuidUnitKerja1, ""))
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo induk kerja by uuid: %w", err)
		}
		pegawai.IdUnitKerja1 = ptr.Uint64(indukKerja.ID)
		pegawai.KdUnit1 = ptr.String(indukKerja.KdUnit1)
	}

	// Pengecekan Unit Kerja
	if ptr.StringValue(pegawaiReq.UuidUnitKerja2, "") != "" {
		unitKerja, err := indukKerjaRepo.GetUnitKerjaByUUID(a, c.Request().Context(), ptr.StringValue(pegawaiReq.UuidUnitKerja2, ""))
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo unit kerja by uuid: %w", err)
		}
		pegawai.IdUnitKerja2 = ptr.Uint64(unitKerja.ID)
		pegawai.KdUnit2 = ptr.String(unitKerja.KdUnit2)
	}

	// Pengecekan Bagian Kerja
	if ptr.StringValue(pegawaiReq.UuidUnitKerja3, "") != "" {
		bagianKerja, err := indukKerjaRepo.GetBagianKerjaByUUID(a, c.Request().Context(), ptr.StringValue(pegawaiReq.UuidUnitKerja3, ""))
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo bagian kerja by uuid: %w", err)
		}
		pegawai.IdUnitKerja3 = ptr.Uint64(bagianKerja.ID)
		pegawai.KdUnit3 = ptr.String(bagianKerja.KdUnit3)
	}

	// Pengecekan Lokasi Kerja
	if ptr.StringValue(pegawaiReq.UuidLokasiKerja, "") != "" {
		lokasiKerja, err := lokasiKerjaRepo.GetLokasiKerjaByUUID(a, c.Request().Context(), ptr.StringValue(pegawaiReq.UuidLokasiKerja, ""))
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo lokasi kerja by uuid: %w", err)
		}
		pegawai.IdUnitKerjaLokasi = ptr.Uint64(lokasiKerja.ID)
		pegawai.LokasiKerja = ptr.String(lokasiKerja.LokasiKerja)
	}

	// Pengecekan Homebase Pddikti
	if ptr.StringValue(pegawaiReq.PegawaiFungsional.UuidHomebasePddikti, "") != "" {
		homebasePddikti, err := indukKerjaRepo.GetUnitKerjaByUUID(a, c.Request().Context(), ptr.StringValue(pegawaiReq.PegawaiFungsional.UuidHomebasePddikti, ""))
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo unit kerja by uuid: %w", err)
		}
		pegawai.PegawaiFungsional.IdHomebasePddikti = ptr.Uint64(homebasePddikti.ID)
		pegawai.PegawaiFungsional.KdHomebasePddikti = ptr.String(homebasePddikti.KdPddikti)
	}

	// Pengecekan Homebase UII
	if ptr.StringValue(pegawaiReq.PegawaiFungsional.UuidHomebaseUii, "") != "" {
		homebaseUUuidHomebaseUii, err := indukKerjaRepo.GetUnitKerjaByUUID(a, c.Request().Context(), ptr.StringValue(pegawaiReq.PegawaiFungsional.UuidHomebaseUii, ""))
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo unit kerja by uuid: %w", err)
		}
		pegawai.PegawaiFungsional.IdHomebaseUii = ptr.Uint64(homebaseUUuidHomebaseUii.ID)
		pegawai.PegawaiFungsional.KdHomebaseUii = ptr.String(homebaseUUuidHomebaseUii.KdUnit2)
	}

	// Pengecekan Pangkat Golongan Ruang PNS
	if ptr.StringValue(pegawaiReq.PegawaiPNS.UuidPangkatGolongan, "") != "" {
		pangkatPNS, err := pangkatPegawaiRepo.GetPangkatGolonganPegawaiByUUID(a, ptr.StringValue(pegawaiReq.PegawaiPNS.UuidPangkatGolongan, ""))
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo pangkat golongan pegawai by uuid: %w", err)
		}
		pegawai.PegawaiPNS.IdPangkatGolongan = ptr.Uint64(pangkatPNS.ID)
		pegawai.PegawaiPNS.KdPangkatGolongan = ptr.String(pangkatPNS.KdPangkat)
	}

	// Pengecekan Jabatan Fungsional Yayasan
	if ptr.StringValue(pegawaiReq.PegawaiPNS.UuidJabatanFungsional, "") != "" {
		jabatanPNS, err := jabatanFungsionalRepo.GetJabatanFungsionalByUUID(a, c.Request().Context(), ptr.StringValue(pegawaiReq.PegawaiPNS.UuidJabatanFungsional, ""))
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo pangkat golongan pegawai by uuid: %w", err)
		}
		pegawai.PegawaiPNS.IdJabatanFungsional = ptr.Uint64(jabatanPNS.ID)
		pegawai.PegawaiPNS.KdJabatanFungsional = ptr.String(jabatanPNS.KdJabatanFungsional)
	}

	// Pengecekan Jenis PTT
	if ptr.StringValue(pegawaiReq.PegawaiPNS.UuidJenisPtt, "") != "" {
		jenisPTT, err := jenisPTTRepo.GetJenisPTTByUUID(a, c.Request().Context(), ptr.StringValue(pegawaiReq.PegawaiPNS.UuidJenisPtt, ""))
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo jenis pegawai tidak tetap by uuid: %w", err)
		}
		pegawai.PegawaiPNS.IdJenisPtt = ptr.Uint64(jenisPTT.ID)
		pegawai.PegawaiPNS.KdJenisPtt = ptr.String(jenisPTT.KdJenisPTT)
	}

	// Pengecekan Status Pegawai
	if ptr.StringValue(pegawaiReq.PegawaiFungsional.UuidStatusPegawaiAktif, "") != "" {
		statusPegawaiAktif, err := statusPegawaiAktifRepo.GetStatusPegawaiAktifByUUID(a, c.Request().Context(), ptr.StringValue(pegawaiReq.PegawaiFungsional.UuidStatusPegawaiAktif, ""))
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("error from repo status pegawai aktif by uuid: %w", err)
		}
		pegawai.PegawaiFungsional.IdStatusPegawaiAktif = ptr.Uint64(statusPegawaiAktif.ID)
		pegawai.PegawaiFungsional.KdStatusPegawaiAktif = ptr.String(statusPegawaiAktif.KdStatusAktif)

		pegawai.FlagMengajar = ptr.String("0")
		if pegawai.IsLecturer() && statusPegawaiAktif.IsActive() && !statusPegawaiAktif.IsOnStudyingPermission() {
			pegawai.FlagMengajar = ptr.String("1")
		}
		pegawai.FlagSekolah = ptr.String("0")
		if statusPegawaiAktif.IsOnStudyingPermission() {
			pegawai.FlagSekolah = ptr.String("1")
		}
		pegawai.FlagPensiun = ptr.String("0")
		if statusPegawaiAktif.IsRetired() {
			pegawai.FlagPensiun = ptr.String("1")
		}
		pegawai.FlagMeninggal = ptr.String("0")
		if statusPegawaiAktif.IsDied() {
			pegawai.FlagMeninggal = ptr.String("1")
		}

	}

	tglStatusAktif := ptr.StringValue(pegawaiReq.PegawaiFungsional.TglStatusPegawaiAktif, "")
	_, err = time.Parse("2006-01-02", tglStatusAktif)
	if tglStatusAktif != "" {
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("tgl_status_aktif harus sesuai format tanggal yyyy-mm-dd")
		}
		pegawai.PegawaiFungsional.TglStatusPegawaiAktif = ptr.String(tglStatusAktif)
	}

	// Binding nilai request ke struct

	if pegawaiReq.PegawaiFungsional.TmtPangkatGolongan != nil {
		_, err = time.Parse("2006-01-02", ptr.StringValue(pegawaiReq.PegawaiFungsional.TmtPangkatGolongan, ""))
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("tmt_pangkat_golongan harus sesuai format tanggal yyyy-mm-dd")
		}
		pegawai.PegawaiFungsional.TmtPangkatGolongan = pegawaiReq.PegawaiFungsional.TmtPangkatGolongan
	}
	if pegawaiReq.PegawaiFungsional.TmtJabatan != nil {
		_, err = time.Parse("2006-01-02", ptr.StringValue(pegawaiReq.PegawaiFungsional.TmtJabatan, ""))
		if err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("tmt_jabatan harus sesuai format tanggal yyyy-mm-dd")
		}
		pegawai.PegawaiFungsional.TmtJabatan = pegawaiReq.PegawaiFungsional.TmtJabatan
	}
	if ptr.StringValue(pegawaiReq.PegawaiFungsional.MasaKerjaBawaanTahun, "") != "" {
		tahun, _ := strconv.Atoi(ptr.StringValue(pegawaiReq.PegawaiFungsional.MasaKerjaBawaanTahun, ""))
		if !(tahun >= 0 && tahun <= 99) {
			return model.PegawaiUpdate{}, fmt.Errorf("masa kerja bawaan tahun hanya dapat diisi antara 0-99")
		}
		pegawai.PegawaiFungsional.MasaKerjaBawaanTahun = pegawaiReq.PegawaiFungsional.MasaKerjaBawaanTahun
	}
	if ptr.StringValue(pegawaiReq.PegawaiFungsional.MasaKerjaBawaanBulan, "") != "" {
		bulan, _ := strconv.Atoi(ptr.StringValue(pegawaiReq.PegawaiFungsional.MasaKerjaBawaanBulan, ""))
		if !(bulan >= 0 && bulan <= 11) {
			return model.PegawaiUpdate{}, fmt.Errorf("masa kerja bawaan bulan hanya dapat diisi antara 0-11")
		}
		pegawai.PegawaiFungsional.MasaKerjaBawaanBulan = pegawaiReq.PegawaiFungsional.MasaKerjaBawaanBulan
	}
	if ptr.StringValue(pegawaiReq.PegawaiFungsional.MasaKerjaGajiTahun, "") != "" {
		tahun, _ := strconv.Atoi(ptr.StringValue(pegawaiReq.PegawaiFungsional.MasaKerjaGajiTahun, ""))
		if !(tahun >= 0 && tahun <= 99) {
			return model.PegawaiUpdate{}, fmt.Errorf("masa kerja gaji tahun hanya dapat diisi antara 0-99")
		}
		pegawai.PegawaiFungsional.MasaKerjaGajiTahun = pegawaiReq.PegawaiFungsional.MasaKerjaGajiTahun
	}
	if ptr.StringValue(pegawaiReq.PegawaiFungsional.MasaKerjaGajiBulan, "") != "" {
		bulan, _ := strconv.Atoi(ptr.StringValue(pegawaiReq.PegawaiFungsional.MasaKerjaGajiBulan, ""))
		if !(bulan >= 0 && bulan <= 11) {
			return model.PegawaiUpdate{}, fmt.Errorf("masa kerja gaji bulan hanya dapat diisi antara 0-11")
		}
		pegawai.PegawaiFungsional.MasaKerjaGajiBulan = pegawaiReq.PegawaiFungsional.MasaKerjaGajiBulan
	}
	if ptr.StringValue(pegawaiReq.PegawaiFungsional.MasaKerjaAwalKepegawaianTahun, "") != "" {
		tahun, _ := strconv.Atoi(ptr.StringValue(pegawaiReq.PegawaiFungsional.MasaKerjaAwalKepegawaianTahun, ""))
		if !(tahun >= 0 && tahun <= 99) {
			return model.PegawaiUpdate{}, fmt.Errorf("masa kerja awal kepegawaian tahun hanya dapat diisi antara 0-99")
		}
		pegawai.PegawaiFungsional.MasaKerjaAwalKepegawaianTahun = pegawaiReq.PegawaiFungsional.MasaKerjaAwalKepegawaianTahun
	}
	if ptr.StringValue(pegawaiReq.PegawaiFungsional.MasaKerjaAwalKepegawaianBulan, "") != "" {
		tahun, _ := strconv.Atoi(ptr.StringValue(pegawaiReq.PegawaiFungsional.MasaKerjaAwalKepegawaianBulan, ""))
		if !(tahun >= 0 && tahun <= 11) {
			return model.PegawaiUpdate{}, fmt.Errorf("masa kerja awal kepegawaian bulan hanya dapat diisi antara 0-11")
		}
		pegawai.PegawaiFungsional.MasaKerjaAwalKepegawaianBulan = pegawaiReq.PegawaiFungsional.MasaKerjaAwalKepegawaianBulan
	}
	if ptr.StringValue(pegawaiReq.PegawaiFungsional.MasaKerjaAwalPensiunTahun, "") != "" {
		tahun, _ := strconv.Atoi(ptr.StringValue(pegawaiReq.PegawaiFungsional.MasaKerjaAwalPensiunTahun, ""))
		if !(tahun >= 0 && tahun <= 99) {
			return model.PegawaiUpdate{}, fmt.Errorf("masa kerja awal pensiun tahun hanya dapat diisi antara 0-99")
		}
		pegawai.PegawaiFungsional.MasaKerjaAwalPensiunTahun = pegawaiReq.PegawaiFungsional.MasaKerjaAwalPensiunTahun
	}
	if ptr.StringValue(pegawaiReq.PegawaiFungsional.MasaKerjaAwalPensiunBulan, "") != "" {
		tahun, _ := strconv.Atoi(ptr.StringValue(pegawaiReq.PegawaiFungsional.MasaKerjaAwalPensiunBulan, ""))
		if !(tahun >= 0 && tahun <= 11) {
			return model.PegawaiUpdate{}, fmt.Errorf("masa kerja awal pensiun bulan hanya dapat diisi antara 0-11")
		}
		pegawai.PegawaiFungsional.MasaKerjaAwalPensiunBulan = pegawaiReq.PegawaiFungsional.MasaKerjaAwalPensiunBulan
	}

	angkaKreditFungsional := ptr.StringValue(pegawaiReq.PegawaiFungsional.AngkaKredit, "")
	if angkaKreditFungsional != "" {
		pegawai.PegawaiFungsional.AngkaKredit = pegawaiReq.PegawaiFungsional.AngkaKredit
	}
	if ptr.StringValue(pegawaiReq.PegawaiFungsional.NomorSertifikasi, "") != "" {
		if len(ptr.StringValue(pegawaiReq.PegawaiFungsional.NomorSertifikasi, "")) > 20 {
			return model.PegawaiUpdate{}, fmt.Errorf("panjang karakter nomor sertifikasi maksimal 20")
		}
		pegawai.PegawaiFungsional.NomorSertifikasi = pegawaiReq.PegawaiFungsional.NomorSertifikasi
	}
	if ptr.StringValue(pegawaiReq.PegawaiFungsional.NomorRegistrasi, "") != "" {
		if len(ptr.StringValue(pegawaiReq.PegawaiFungsional.NomorRegistrasi, "")) > 10 {
			return model.PegawaiUpdate{}, fmt.Errorf("panjang karakter nomor registrasi maksimal 10")
		}
		pegawai.PegawaiFungsional.NomorRegistrasi = pegawaiReq.PegawaiFungsional.NomorRegistrasi
	}
	if ptr.StringValue(pegawaiReq.PegawaiFungsional.NomorSkPertama, "") != "" {
		if len(ptr.StringValue(pegawaiReq.PegawaiFungsional.NomorSkPertama, "")) > 30 {
			return model.PegawaiUpdate{}, fmt.Errorf("panjang karakter nomor sk pertama maksimal 30")
		}
		pegawai.PegawaiFungsional.NomorSkPertama = pegawaiReq.PegawaiFungsional.NomorSkPertama
	}
	if pegawaiReq.PegawaiFungsional.TmtSkPertama != nil {
		pegawai.PegawaiFungsional.TmtSkPertama = pegawaiReq.PegawaiFungsional.TmtSkPertama
	}
	if ptr.StringValue(pegawaiReq.PegawaiPNS.InstansiAsal, "") != "" {
		pegawai.PegawaiPNS.InstansiAsal = pegawaiReq.PegawaiPNS.InstansiAsal
	}
	if ptr.StringValue(pegawaiReq.PegawaiPNS.NipPns, "") != "" {
		if len(ptr.StringValue(pegawaiReq.PegawaiPNS.NipPns, "")) != 18 {
			return model.PegawaiUpdate{}, fmt.Errorf("panjang karakter NIP PNS hanya boleh 18")
		}
		pegawai.PegawaiPNS.NipPns = pegawaiReq.PegawaiPNS.NipPns
	}

	Nira := ptr.StringValue(pegawaiReq.PegawaiPNS.Nira, "")
	if Nira != "" {
		if len(Nira) > 23 {
			return model.PegawaiUpdate{}, fmt.Errorf("panjang karakter NIRA maksimal 23")
		} else if _, err = strconv.ParseComplex(Nira, 128); err != nil {
			return model.PegawaiUpdate{}, fmt.Errorf("kolom NIRA hanya dapat diisi dengan angka")
		} else {
			pegawai.PegawaiPNS.Nira = pegawaiReq.PegawaiPNS.Nira
		}
	}

	if ptr.StringValue(pegawaiReq.PegawaiPNS.NoKartuPegawai, "") != "" {
		if len(ptr.StringValue(pegawaiReq.PegawaiPNS.NoKartuPegawai, "")) != 18 {
			return model.PegawaiUpdate{}, fmt.Errorf("panjang karakter nomor kartu pegawai hanya boleh 18")
		}
		pegawai.PegawaiPNS.NoKartuPegawai = pegawaiReq.PegawaiPNS.NoKartuPegawai
	}
	if pegawaiReq.PegawaiPNS.TmtPangkatGolongan != nil {
		pegawai.PegawaiPNS.TmtPangkatGolongan = pegawaiReq.PegawaiPNS.TmtPangkatGolongan
	}
	if pegawaiReq.PegawaiPNS.TmtJabatan != nil {
		pegawai.PegawaiPNS.TmtJabatan = pegawaiReq.PegawaiPNS.TmtJabatan
	}
	if ptr.StringValue(pegawaiReq.PegawaiPNS.MasaKerjaTahun, "") != "" {
		pegawai.PegawaiPNS.MasaKerjaTahun = pegawaiReq.PegawaiPNS.MasaKerjaTahun
	}
	if ptr.StringValue(pegawaiReq.PegawaiPNS.MasaKerjaBulan, "") != "" {
		a, _ := strconv.Atoi(ptr.StringValue(pegawaiReq.PegawaiPNS.MasaKerjaBulan, ""))
		if a > 12 {
			return model.PegawaiUpdate{}, fmt.Errorf("masa kerja bulan maksimal 12")
		}
		pegawai.PegawaiPNS.MasaKerjaBulan = pegawaiReq.PegawaiPNS.MasaKerjaBulan
	}
	if ptr.StringValue(pegawaiReq.PegawaiPNS.AngkaKredit, "") != "" {
		pegawai.PegawaiPNS.AngkaKredit = pegawaiReq.PegawaiPNS.AngkaKredit
	}
	keterangan := ptr.StringValue(pegawaiReq.PegawaiPNS.Keterangan, "")
	if len(keterangan) > 400 {
		return model.PegawaiUpdate{}, fmt.Errorf("panjang keterangan maksimal 400 karakter")
	}
	if keterangan != "" {
		pegawai.PegawaiPNS.Keterangan = pegawaiReq.PegawaiPNS.Keterangan
	}

	pegawai.UserUpdate = ptr.String(user)
	pegawai.PegawaiFungsional.UserUpdate = ptr.String(user)
	pegawai.PegawaiPNS.UserUpdate = ptr.String(user)
	return pegawai, nil
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

	pegawai.FlagMengajar = "0"
	if pegawai.IsLecturer() && statusPegawaiAktif.IsActive() && !statusPegawaiAktif.IsOnStudyingPermission() {
		pegawai.FlagMengajar = "1"
	}
	pegawai.FlagSekolah = "0"
	if statusPegawaiAktif.IsOnStudyingPermission() {
		pegawai.FlagSekolah = "1"
	}
	pegawai.FlagPensiun = "0"
	if statusPegawaiAktif.IsRetired() {
		pegawai.FlagPensiun = "1"
	}
	pegawai.FlagMeninggal = "0"
	if statusPegawaiAktif.IsDied() {
		pegawai.FlagMeninggal = "1"
	}

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
	case pegawaiReq.KdJenisPresensi == "":
		return model.PegawaiCreate{}, fmt.Errorf("kd_jenis_presensi tidak boleh kosong")
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
	pegawaiReq.IdKelompokPegawaiPayroll = kelompokPegawai.ID                //
	pegawaiReq.KdKelompokPegawaiPayroll = kelompokPegawai.KdKelompokPegawai //
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
