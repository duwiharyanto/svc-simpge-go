package usecase

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/pegawai/model"
	pengaturan "svc-insani-go/modules/v1/pengaturan-insani/usecase"
	pegawaiOraHttp "svc-insani-go/modules/v1/simpeg-oracle/http"
	pegawaiOraModel "svc-insani-go/modules/v1/simpeg-oracle/model"

	"github.com/labstack/echo/v4"
)

func newPegawaiOra(pegawaiInsani *model.PegawaiDetail) *pegawaiOraModel.KepegawaianYayasanSimpeg {
	pegawaiOra := &pegawaiOraModel.KepegawaianYayasanSimpeg{}
	pegawaiOra.JenisPegawai = &pegawaiOraModel.JenisPegawai{}
	pegawaiOra.InstansiAsalPtt = &pegawaiOraModel.InstansiAsalPtt{}
	pegawaiOra.KelompokPegawai = &pegawaiOraModel.KelompokPegawai{}
	pegawaiOra.LokasiKerja = &pegawaiOraModel.LokasiKerja{}
	pegawaiOra.StatusPegawai = &pegawaiOraModel.StatusPegawai{}
	pegawaiOra.PegawaiStatus = &pegawaiOraModel.PegawaiStatus{}
	pegawaiOra.Unit1 = &pegawaiOraModel.Unit1{}
	pegawaiOra.Unit2 = &pegawaiOraModel.Unit2{}
	pegawaiOra.Unit3 = &pegawaiOraModel.Unit3{}
	pegawaiOra.PegawaiStatus.PangkatKopertis = &pegawaiOraModel.Pangkat{}
	pegawaiOra.PegawaiStatus.PangkatYayasan = &pegawaiOraModel.Pangkat{}
	pegawaiOra.PegawaiStatus.JabatanFungsionalKopertis = &pegawaiOraModel.JabatanFungsional{}
	pegawaiOra.PegawaiStatus.JabatanFungsional = &pegawaiOraModel.JabatanFungsional{}

	pegawaiOra.NIP = pegawaiInsani.PegawaiPribadi.NIK
	pegawaiOra.GelarDepan = pegawaiInsani.PegawaiPribadi.GelarDepan
	pegawaiOra.GelarBelakang = pegawaiInsani.PegawaiPribadi.GelarBelakang
	pegawaiOra.KdStatusPendidikanMasuk = pegawaiInsani.PegawaiPribadi.KdStatusPendidikanMasuk
	pegawaiOra.KdJenisPendidikan = pegawaiInsani.PegawaiPribadi.KdJenisPendidikan

	// Sinkron Kepegawaian Yayaysan - Status
	if pegawaiInsani.PegawaiYayasan.KDJenisPegawai != "" {
		// pegawaiOra.KdJenisPegawai = pegawaiInsani.PegawaiYayasan.KDJenisPegawai
		pegawaiOra.JenisPegawai.KdJenisPegawai = pegawaiInsani.PegawaiYayasan.KDJenisPegawai
	}

	if pegawaiInsani.PegawaiYayasan.StatusPegawai != "" {
		pegawaiOra.StatusPegawai.KdStatusPegawai = pegawaiInsani.PegawaiYayasan.KDStatusPegawai
	}

	if pegawaiInsani.PegawaiYayasan.KdPendidikanMasuk != "" {
		pegawaiOra.KdPendidikanMasuk = pegawaiInsani.PegawaiYayasan.KdPendidikanMasukSimpeg
	}

	if pegawaiInsani.PegawaiYayasan.KdPendidikanTerakhir != "" {
		pegawaiOra.KdPendidikan = pegawaiInsani.PegawaiYayasan.KdPendidikanTerakhirSimpeg
	}

	if pegawaiInsani.PegawaiYayasan.KdKelompokPegawai != "" {
		pegawaiOra.KelompokPegawai.KdKelompokPegawai = pegawaiInsani.PegawaiYayasan.KdKelompokPegawai
	}

	// Sinkron Kepegawaian Yayaysan - Pangkat / Jabatan
	if pegawaiInsani.PegawaiYayasan.KdGolongan != "" {
		pegawaiOra.PegawaiStatus.PangkatYayasan.KdGolongan = pegawaiInsani.PegawaiYayasan.KdGolongan
	}

	if pegawaiInsani.PegawaiYayasan.KdRuang != "" {
		pegawaiOra.PegawaiStatus.PangkatYayasan.KdRuang = pegawaiInsani.PegawaiYayasan.KdRuang
	}

	if pegawaiInsani.PegawaiYayasan.TmtPangkatGolongan != "" {
		pegawaiOra.PegawaiStatus.PangkatYayasan.TmtPangkat = pegawaiInsani.PegawaiYayasan.TmtPangkatGolongan
	}

	if pegawaiInsani.PegawaiYayasan.KdJabatanFungsional != "" {
		pegawaiOra.PegawaiStatus.JabatanFungsional.KdFungsional = pegawaiInsani.PegawaiYayasan.KdJabatanFungsional
	}

	if pegawaiInsani.PegawaiYayasan.TmtJabatan != "" {
		pegawaiOra.PegawaiStatus.JabatanFungsional.TmtFungsional = pegawaiInsani.PegawaiYayasan.TmtJabatan
	}

	if pegawaiInsani.PegawaiYayasan.MasaKerjaGajiTahun != "" {
		pegawaiOra.PegawaiStatus.MasaKerjaGajiTahun, _ = strconv.Atoi(pegawaiInsani.PegawaiYayasan.MasaKerjaGajiTahun)
	}

	if pegawaiInsani.PegawaiYayasan.MasaKerjaGajiBulan != "" {
		pegawaiOra.PegawaiStatus.MasaKerjaGajiBulan, _ = strconv.Atoi(pegawaiInsani.PegawaiYayasan.MasaKerjaGajiBulan)
	}

	if pegawaiInsani.PegawaiYayasan.AngkaKredit != "" {
		pegawaiOra.PegawaiStatus.AngkaKreditFungsional, _ = strconv.ParseFloat(pegawaiInsani.PegawaiYayasan.AngkaKredit, 64)
	}

	// Sinkron Unit Kerja
	if pegawaiInsani.UnitKerjaPegawai.KdIndukKerja != "" {
		pegawaiOra.Unit1.KdUnit1 = pegawaiInsani.UnitKerjaPegawai.KdIndukKerja
	}

	if pegawaiInsani.UnitKerjaPegawai.KdUnitKerja != "" {
		pegawaiOra.Unit2.KdUnit2 = pegawaiInsani.UnitKerjaPegawai.KdUnitKerja
	}

	if pegawaiInsani.UnitKerjaPegawai.KdBagianKerja != "" {
		pegawaiOra.Unit3.KdUnit3 = pegawaiInsani.UnitKerjaPegawai.KdBagianKerja
	}

	if pegawaiInsani.UnitKerjaPegawai.LokasiKerja != "" {
		pegawaiOra.LokasiKerja.KdLokasi = pegawaiInsani.UnitKerjaPegawai.LokasiKerja
	}

	if pegawaiInsani.UnitKerjaPegawai.NoSkPertama != "" {
		pegawaiOra.PegawaiStatus.NoSkPertama = pegawaiInsani.UnitKerjaPegawai.NoSkPertama
	}

	if pegawaiInsani.UnitKerjaPegawai.TmtSkPertama != "" {
		pegawaiOra.PegawaiStatus.TglSkPertama = pegawaiInsani.UnitKerjaPegawai.TmtSkPertama
	}

	if pegawaiInsani.UnitKerjaPegawai.KdHomebasePddikti != "" {
		pegawaiOra.PegawaiStatus.KdHomebasePddikti = pegawaiInsani.UnitKerjaPegawai.KdHomebasePddikti
	}

	if pegawaiInsani.UnitKerjaPegawai.KdHomebaseUii != "" {
		pegawaiOra.PegawaiStatus.KdHomebaseUii = pegawaiInsani.UnitKerjaPegawai.KdHomebaseUii
	}

	// Sinkron Kepegawaian Negara / PTT
	if pegawaiInsani.PegawaiPNSPTT.InstansiAsalPtt != "" {
		pegawaiOra.InstansiAsalPtt.Instansi = pegawaiInsani.PegawaiPNSPTT.InstansiAsalPtt
	}

	if pegawaiInsani.PegawaiPNSPTT.NipPNS != "" {
		pegawaiOra.NipKopertis = pegawaiInsani.PegawaiPNSPTT.NipPNS
	}

	if pegawaiInsani.PegawaiPNSPTT.NoKartuPegawai != "" {
		pegawaiOra.PegawaiStatus.NoKarpeg = pegawaiInsani.PegawaiPNSPTT.NoKartuPegawai
	}

	if pegawaiInsani.PegawaiPNSPTT.KdGolonganPNS != "" {
		pegawaiOra.PegawaiStatus.PangkatKopertis.KdGolongan = pegawaiInsani.PegawaiPNSPTT.KdGolonganPNS
	}

	if pegawaiInsani.PegawaiPNSPTT.KdRuangPNS != "" {
		pegawaiOra.PegawaiStatus.PangkatKopertis.KdRuang = pegawaiInsani.PegawaiPNSPTT.KdRuangPNS
	}

	if pegawaiInsani.PegawaiPNSPTT.TmtPangkatGolongan != "" {
		pegawaiOra.PegawaiStatus.PangkatKopertis.TmtPangkat = pegawaiInsani.PegawaiPNSPTT.TmtPangkatGolongan
	}

	if pegawaiInsani.PegawaiPNSPTT.KdJabatanPns != "" {
		pegawaiOra.PegawaiStatus.JabatanFungsionalKopertis.KdFungsional = pegawaiInsani.PegawaiPNSPTT.KdJabatanPns
	}

	if pegawaiInsani.PegawaiPNSPTT.TmtPangkatGolongan != "" {
		pegawaiOra.PegawaiStatus.JabatanFungsionalKopertis.TmtFungsional = pegawaiInsani.PegawaiPNSPTT.TmtJabatanPns
	}

	if pegawaiInsani.PegawaiPNSPTT.MasaKerjaPnsTahun != "" {
		pegawaiOra.PegawaiStatus.MasaKerjaKopertisTahun, _ = strconv.Atoi(pegawaiInsani.PegawaiPNSPTT.MasaKerjaPnsTahun)
	}

	if pegawaiInsani.PegawaiPNSPTT.MasaKerjaPnsBulan != "" {
		pegawaiOra.PegawaiStatus.MasaKerjaKopertisBulan, _ = strconv.Atoi(pegawaiInsani.PegawaiPNSPTT.MasaKerjaPnsBulan)
	}

	if pegawaiInsani.PegawaiPNSPTT.AngkaKreditPns != "" {
		pegawaiOra.PegawaiStatus.AngkaKreditKopertis, _ = strconv.ParseFloat(pegawaiInsani.PegawaiPNSPTT.AngkaKreditPns, 64)
	}

	if pegawaiInsani.PegawaiPNSPTT.KeteranganPNS != "" {
		pegawaiOra.InstansiAsalPtt.Keterangan = pegawaiInsani.PegawaiPNSPTT.KeteranganPNS
	}

	// status aktif
	pegawaiOra.PegawaiStatus.FlagMengajar = "N"
	if pegawaiInsani.PegawaiYayasan.KDJenisPegawai == "ED" &&
		pegawaiInsani.StatusAktif.FlagAktifPegawai == "1" &&
		pegawaiInsani.StatusAktif.KdStatusAktifPegawai != "IBL" {
		pegawaiOra.PegawaiStatus.FlagMengajar = "Y"
	}

	pegawaiOra.PegawaiStatus.FlagSekolah = "N"
	if pegawaiInsani.StatusAktif.KdStatusAktifPegawai == "IBL" {
		pegawaiOra.PegawaiStatus.FlagSekolah = "Y"
	}

	pegawaiOra.FlagPensiun = "N"
	if pegawaiInsani.StatusAktif.KdStatusAktifPegawai == "PEN" {
		pegawaiOra.FlagPensiun = "Y"
	}

	pegawaiOra.KdStatusHidup = "Y"
	if pegawaiInsani.StatusAktif.KdStatusAktifPegawai == "MNG" {
		pegawaiOra.KdStatusHidup = "N"
	}

	pegawaiOra.UserUpdate = pegawaiInsani.PegawaiPribadi.UserUpdate

	return pegawaiOra
}

func SendPegawaiToOracle(a *app.App, ctx context.Context, uuid string) error {
	flagSinkronSimpeg, err := pengaturan.LoadPengaturan(a, ctx, nil, pengaturanAtributFlagSinkronSimpeg)
	if err != nil {
		return fmt.Errorf("error load pengaturan flag sinkron simpeg: %w", err)
	}

	disableSyncSimpegOracle, _ := strconv.ParseBool(os.Getenv("DISABLE_SYNC_SIMPEG_ORACLE"))
	if flagSinkronSimpeg != "1" || disableSyncSimpegOracle {
		log.Printf("[DEBUG] flag sinkron simpeg 0\n")
		return nil
	}

	pegawaiDetail, err := PrepareGetSimpegPegawaiByUUID(a, uuid)
	if err != nil {
		return fmt.Errorf("error prepare get simpeg pegawai by uuid: %w", err)
	}

	pegawaiOra := newPegawaiOra(&pegawaiDetail)
	pegawaiOra.Nama = pegawaiDetail.PegawaiPribadi.Nama
	pegawaiOra.KdAgama = pegawaiDetail.PegawaiPribadi.KdItemAgama
	pegawaiOra.KdGolonganDarah = pegawaiDetail.PegawaiPribadi.GolonganDarah
	pegawaiOra.KdKelamin = pegawaiDetail.PegawaiPribadi.KdKelamin
	pegawaiOra.KdNikah = pegawaiDetail.PegawaiPribadi.KdNikah
	pegawaiOra.TempatLahir = pegawaiDetail.PegawaiPribadi.TempatLahir
	pegawaiOra.TanggalLahir = pegawaiDetail.PegawaiPribadi.TanggalLahir
	pegawaiOra.GelarDepan = pegawaiDetail.PegawaiPribadi.GelarDepan
	pegawaiOra.GelarBelakang = pegawaiDetail.PegawaiPribadi.GelarBelakang
	// pegawaiOra.JumlahAnak = pegawaiDetail.PegawaiPribadi.JumlahAnak
	// pegawaiOra.JumlahDitanggung = pegawaiDetail.PegawaiPribadi.JumlahDitanggung
	// pegawaiOra.JumlahKeluarga = pegawaiDetail.PegawaiPribadi.JumlahKeluarga
	pegawaiOra.NoKTP = pegawaiDetail.PegawaiPribadi.NoKTP
	// pegawaiOra.NoTelepon = pegawaiDetail.PegawaiPribadi.NoTelepon
	pegawaiOra.UserInput = pegawaiDetail.PegawaiPribadi.UserInput

	err = pegawaiOraHttp.CreateKepegawaianYayasan(ctx, a.HttpClient, pegawaiOra)
	if err != nil {
		return fmt.Errorf("[ERROR] create kepegawaian yayasan: %w", err)
	}

	return nil
}

func HandleResyncOracle(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		uuidPegawai := c.Param("uuidPegawai")
		if uuidPegawai == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "uuid pegawai tidak boleh kosong"})
		}

		pegawai, err := PrepareGetSimpegPegawaiByUUID(a, uuidPegawai)
		if err != nil {
			fmt.Printf("[ERROR] prepare get simpeg pegawai by uuid: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		if pegawai.IsEmpty() {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Pegawai tidak ditemukan"})
		}

		ctx := c.Request().Context()
		pegawaiOra, err := pegawaiOraHttp.GetKepegawaianYayasan(ctx, a.HttpClient, pegawai.PegawaiPribadi.NIK)
		if err != nil {
			fmt.Printf("[ERROR] get kepegawaian yayasan: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		if pegawaiOra == nil {
			pegawaiOraCreate := newPegawaiOra(&pegawai)
			pegawaiOraCreate.Nama = pegawai.PegawaiPribadi.Nama
			pegawaiOraCreate.KdAgama = pegawai.PegawaiPribadi.KdItemAgama
			pegawaiOraCreate.KdGolonganDarah = pegawai.PegawaiPribadi.GolonganDarah
			pegawaiOraCreate.KdKelamin = pegawai.PegawaiPribadi.KdKelamin
			pegawaiOraCreate.KdNikah = pegawai.PegawaiPribadi.KdNikah
			pegawaiOraCreate.TempatLahir = pegawai.PegawaiPribadi.TempatLahir
			pegawaiOraCreate.TanggalLahir = pegawai.PegawaiPribadi.TanggalLahir
			pegawaiOraCreate.GelarDepan = pegawai.PegawaiPribadi.GelarDepan
			pegawaiOraCreate.GelarBelakang = pegawai.PegawaiPribadi.GelarBelakang
			// pegawaiOraCreate.JumlahAnak = pegawai.PegawaiPribadi.JumlahAnak
			// pegawaiOraCreate.JumlahDitanggung = pegawai.PegawaiPribadi.JumlahDitanggung
			// pegawaiOraCreate.JumlahKeluarga = pegawai.PegawaiPribadi.JumlahKeluarga
			pegawaiOraCreate.NoKTP = pegawai.PegawaiPribadi.NoKTP
			// pegawaiOraCreate.NoTelepon = pegawai.PegawaiPribadi.NoTelepon
			pegawaiOraCreate.UserInput = pegawai.PegawaiPribadi.UserInput
			err = pegawaiOraHttp.CreateKepegawaianYayasan(ctx, a.HttpClient, pegawaiOraCreate)
			if err != nil {
				fmt.Printf("[ERROR] create kepegawaian yayasan: %s\n", err.Error())
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
			}
		} else {
			pegawaiOraUpdate := newPegawaiOra(&pegawai)
			err = pegawaiOraHttp.UpdateKepegawaianYayasan(ctx, &http.Client{}, pegawaiOraUpdate)
			if err != nil {
				fmt.Printf("[ERROR] repo update kepegawaian yayasan: %s\n", err.Error())
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
			}

		}
		return c.JSON(http.StatusOK, map[string]string{"message": "Berhasil sinkron ulang pegawai"})
	}
	return echo.HandlerFunc(h)
}
