package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"net/http"
	"svc-insani-go/app"
	pegawaiOraHttp "svc-insani-go/modules/v1/pegawai-oracle/http"
	pegawaiOraModel "svc-insani-go/modules/v1/pegawai-oracle/model"
	"svc-insani-go/modules/v1/pegawai/model"
	"svc-insani-go/modules/v1/pegawai/repo"
	pengaturan "svc-insani-go/modules/v1/pengaturan-insani/usecase"

	"github.com/labstack/echo/v4"
)

const (
	pengaturanAtributFlagSinkronSimpeg = "flag_sinkron_simpeg"
)

func HandleGetPegawai(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		req := &model.PegawaiRequest{}
		err := c.Bind(req)
		if err != nil {
			fmt.Printf("[WARNING] binding pegawai request: %s\n", err.Error())
		}
		res := model.PegawaiResponse{
			Data: []model.Pegawai{},
		}
		count, err := repo.CountPegawai(a, req)
		if err != nil {
			fmt.Printf("[ERROR] repo count pegawai: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		if count == 0 {
			return c.JSON(http.StatusOK, res)
		}
		pp, err := repo.GetAllPegawai(a, req)
		if err != nil {
			fmt.Printf("[ERROR] repo get all pegawai: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		res.Count = count
		res.Data = pp
		res.Limit = req.Limit
		res.Offset = req.Offset
		return c.JSON(http.StatusOK, res)
	}
	return echo.HandlerFunc(h)
}

func HandleGetSimpegPegawaiByUUID(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		uuidPegawai := c.Param("uuidPegawai")
		if uuidPegawai == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "parameter uuid pegawai wajib diisi"})
		}

		pegawaiDetail, err := PrepareGetSimpegPegawaiByUUID(a, uuidPegawai)
		if err != nil {
			log.Printf("[ERROR] repo get kepegawaian: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		return c.JSON(http.StatusOK, pegawaiDetail)
	}
	return echo.HandlerFunc(h)
}

func PrepareGetSimpegPegawaiByUUID(a *app.App, uuidPegawai string) (model.PegawaiDetail, error) {
	pegawaiDetail := model.PegawaiDetail{}

	pegawaiPribadi, err := repo.GetPegawaiPribadi(a, uuidPegawai)
	if err != nil {
		return model.PegawaiDetail{}, fmt.Errorf("error repo get pribadi pegawai uuid, %w", err)
	}

	kepegawaianYayasan, err := repo.GetKepegawaianYayasan(a, uuidPegawai)
	if err != nil {
		return model.PegawaiDetail{}, fmt.Errorf("error repo get kepegawaian yayasan uuid, %w", err)
	}

	unitKerjaPegawai, err := repo.GetUnitKerjaPegawai(a, uuidPegawai)
	if err != nil {
		return model.PegawaiDetail{}, fmt.Errorf("error repo get unit kerja pegawai by uuid, %w", err)
	}

	pegawaiPNS, err := repo.GetPegawaiPNS(a, uuidPegawai)
	if err != nil {
		return model.PegawaiDetail{}, fmt.Errorf("error repo get pegawai pns by uuid, %w", err)
	}

	statusPegawaiAktif, err := repo.GetStatusPegawaiAktif(a, uuidPegawai)
	if err != nil {
		return model.PegawaiDetail{}, fmt.Errorf("error repo get status aktif pegawai by uuid, %w", err)
	}

	pegawaiPendidikan, err := repo.GetPegawaiPendidikan(a, uuidPegawai, true)
	if err != nil {
		return model.PegawaiDetail{}, fmt.Errorf("error repo get pendidikan pegawai by uuid, %w", err)
	}
	kepegawaianYayasan.SetMasaKerjaTotal(unitKerjaPegawai)
	pegawaiDetail.PegawaiYayasan = kepegawaianYayasan
	pegawaiDetail.UnitKerjaPegawai = unitKerjaPegawai
	pegawaiDetail.StatusAktif = statusPegawaiAktif
	pegawaiDetail.PegawaiPribadi = pegawaiPribadi
	pegawaiDetail.JenjangPendidikan.Data = pegawaiPendidikan
	pegawaiDetail.JenjangPendidikan.UuidPendidikanMasuk = kepegawaianYayasan.UuidPendidikanMasuk
	pegawaiDetail.JenjangPendidikan.KdPendidikanMasuk = kepegawaianYayasan.KdPendidikanMasukSimpeg
	pegawaiDetail.JenjangPendidikan.PendidikanMasuk = kepegawaianYayasan.PendidikanMasuk

	pegawaiDetail.PegawaiPNSPTT = pegawaiPNS

	return pegawaiDetail, nil
}

func HandleUpdatePegawai(a *app.App, ctx context.Context, errChan chan error) echo.HandlerFunc {
	h := func(c echo.Context) error {
		// Validasi Data
		pegawaiUpdate, err := ValidateUpdatePegawaiByUUID(a, c)
		if err != nil {
			fmt.Printf("[ERROR]: %s\n", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}

		// Update Data
		err = repo.UpdatePegawai(a, c.Request().Context(), pegawaiUpdate)
		if err != nil {
			fmt.Printf("[ERROR]: %s\n", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}

		// Set Flag Pendidikan
		uuidPendidikanDiakui := c.FormValue("uuid_tingkat_pdd_diakui")
		uuidPendidikanTerakhir := c.FormValue("uuid_tingkat_pdd_terakhir")
		idPersonalPegawai := pegawaiUpdate.IdPersonalDataPribadi

		if uuidPendidikanDiakui != "" || uuidPendidikanTerakhir != "" {
			err = repo.UpdatePendidikanPegawai(a, c.Request().Context(), uuidPendidikanDiakui, uuidPendidikanTerakhir, idPersonalPegawai)
			if err != nil {
				fmt.Printf("[ERROR]: %s\n", err.Error())
				return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
			}
		}

		// Menampilkan response
		pegawaiDetail, err := PrepareGetSimpegPegawaiByUUID(a, pegawaiUpdate.Uuid)
		if err != nil {
			fmt.Printf("[ERROR] repo get kepegawaian: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		fmt.Printf("[DEBUG] update response end\n")
		go func(
			a *app.App,
			ctx context.Context,
			errChan chan error,
		) {
			defer func(n time.Time) {
				fmt.Printf("[DEBUG] send to simpeg: %v ms\n", time.Now().Sub(n).Milliseconds())
			}(time.Now())
			fmt.Println("[DEBUG] Go routine start after update")

			flagSinkronSimpeg, err := pengaturan.LoadPengaturan(a, ctx, nil, pengaturanAtributFlagSinkronSimpeg)
			if err != nil {
				log.Println("error load pengaturan flag sinkron simpeg: %w", err)
				errChan <- err
				return
			}

			if flagSinkronSimpeg != "1" {
				log.Printf("[DEBUG] flag sinkron simpeg 0\n")
				return
			}

			dur, err := time.ParseDuration(os.Getenv("RESPONSE_TIMEOUT_MS" + "ms"))
			if err != nil {
				dur = time.Second * 40
			}
			ctx, cancel := context.WithTimeout(ctx, dur)
			// ctx, cancel := context.WithTimeout(context.Background(), dur) // kalau ke cancel pake yang ini
			defer cancel()

			// fmt.Println("DEBUG : Go routin before sinkron simpeg")
			pegawaiOra := newPegawaiOra(&pegawaiDetail)
			err = pegawaiOraHttp.UpdateKepegawaianYayasan(ctx, &http.Client{}, pegawaiOra)
			if err != nil {
				errChan <- fmt.Errorf("[ERROR] repo update kepegawaian yayasan: %w\n", err)
				return
			}
		}(a, ctx, errChan)

		return c.JSON(http.StatusOK, pegawaiDetail)
	}

	return echo.HandlerFunc(h)
}

func HandleCreatePegawai(a *app.App, ctx context.Context, errChan chan error) echo.HandlerFunc {
	h := func(c echo.Context) error {
		// Validasi Data
		pegawaiCreate, err := PrepareCreateSimpeg(a, c)
		if errors.Unwrap(err) != nil {
			fmt.Printf("[ERROR] prepare create simpeg: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}

		fmt.Printf("\n[DEBUG] pegawai before create: %+v\n", pegawaiCreate)

		// Create Data
		err = repo.CreatePegawai(a, c.Request().Context(), pegawaiCreate)
		if err != nil {
			fmt.Printf("[ERROR]: %s\n", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}

		// Set Flag Pendidikan
		uuidPendidikanDiakui := c.FormValue("uuid_tingkat_pdd_diakui")
		uuidPendidikanTerakhir := c.FormValue("uuid_tingkat_pdd_terakhir")
		idPersonalPegawai := pegawaiCreate.IdPersonalDataPribadi

		err = repo.UpdatePendidikanPegawai(a, c.Request().Context(), uuidPendidikanDiakui, uuidPendidikanTerakhir, idPersonalPegawai)
		if err != nil {
			fmt.Printf("[ERROR]: %s\n", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}

		// Menampilkan response
		pegawaiDetail, err := PrepareGetSimpegPegawaiByUUID(a, pegawaiCreate.Uuid)
		if err != nil {
			fmt.Printf("[ERROR] repo get kepegawaian: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		go func(a *app.App, errChan chan error, uuidPegawai string) {
			dur, err := time.ParseDuration(os.Getenv("RESPONSE_TIMEOUT_MS" + "ms"))
			if err != nil {
				dur = time.Second * 40
			}
			ctx, cancel := context.WithTimeout(ctx, dur)
			// ctx, cancel := context.WithTimeout(context.Background(), dur) // kalau ke cancel pake yang ini
			defer cancel()
			err = SendPegawaiToOracle(a, ctx, uuidPegawai)
			if err != nil {
				errChan <- err
				return
			}
		}(a, errChan, pegawaiCreate.Uuid)

		return c.JSON(http.StatusOK, pegawaiDetail)
	}

	return echo.HandlerFunc(h)
}

func HandleGetPendidikanByUUIDPersonal(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		uuidPersonal := c.Param("uuidPersonal")
		if uuidPersonal == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "parameter uuid personal wajib diisi"})
		}

		pendidikanPegawai, err := repo.GetPegawaiPendidikanPersonal(a, uuidPersonal)
		if err != nil {
			log.Printf("[ERROR] repo get pendidikan: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		pendidikanDetail := model.PendidikanPersonal{
			Data: pendidikanPegawai,
		}

		return c.JSON(http.StatusOK, pendidikanDetail)
	}
	return echo.HandlerFunc(h)
}

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

	// Sinkron Status Aktif
	if pegawaiInsani.StatusAktif.FlagAktifPegawai == "1" {
		pegawaiOra.PegawaiStatus.FlagMengajar = "N"
		pegawaiOra.FlagPensiun = "N"
		pegawaiOra.KdStatusHidup = "Y"
		pegawaiOra.PegawaiStatus.FlagSekolah = "N"
		if pegawaiInsani.PegawaiYayasan.KDJenisPegawai == "ED" {
			pegawaiOra.PegawaiStatus.FlagMengajar = "Y"
			if pegawaiInsani.StatusAktif.KdStatusAktifPegawai == "IBL" {
				pegawaiOra.PegawaiStatus.FlagSekolah = "Y"
			}
		}
	}

	if pegawaiInsani.StatusAktif.FlagAktifPegawai == "0" {
		pegawaiOra.PegawaiStatus.FlagMengajar = "N"
		pegawaiOra.FlagPensiun = "N"
		pegawaiOra.KdStatusHidup = "Y"
		pegawaiOra.PegawaiStatus.FlagSekolah = "N"
		if pegawaiInsani.StatusAktif.KdStatusAktifPegawai == "PEN" {
			pegawaiOra.FlagPensiun = "Y"
		}
		if pegawaiInsani.StatusAktif.KdStatusAktifPegawai == "MNG" {
			pegawaiOra.KdStatusHidup = "N"
		}
	}
	pegawaiOra.UserUpdate = pegawaiInsani.PegawaiPribadi.UserUpdate

	return pegawaiOra
}

func HandleCheckNikPegawai(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		nikPegawai := c.QueryParam("nik")
		if nikPegawai == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "parameter nik pegawai wajib diisi"})
		}

		if len(nikPegawai) != 9 {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "panjang NIK pegawai tidak valid"})
		}

		checkNik, flagCheck, err := repo.CheckNikPegawai(a, c.Request().Context(), nikPegawai)
		if err != nil {
			log.Printf("[ERROR] check nik pegawai: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		if flagCheck == true {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "NIK " + checkNik.Nik + " sudah digunakan oleh " + checkNik.Nama})
		}
		return c.JSON(http.StatusOK, nil)
	}
	return echo.HandlerFunc(h)
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
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		if pegawai.IsEmpty() {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Pegawai tidak ditemukan"})
		}

		ctx := c.Request().Context()
		pegawaiOra, err := pegawaiOraHttp.GetKepegawaianYayasan(ctx, a.HttpClient, pegawai.PegawaiPribadi.NIK)
		if err != nil {
			fmt.Printf("[ERROR] get kepegawaian yayasan: %s\n", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Layanan sedang bermasalah"})
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
				return c.JSON(http.StatusBadRequest, map[string]string{"message": "Layanan sedang bermasalah"})
			}
		} else {
			pegawaiOraUpdate := newPegawaiOra(&pegawai)
			err = pegawaiOraHttp.UpdateKepegawaianYayasan(ctx, &http.Client{}, pegawaiOraUpdate)
			if err != nil {
				fmt.Printf("[ERROR] repo update kepegawaian yayasan: %s\n", err.Error())
				return c.JSON(http.StatusBadRequest, map[string]string{"message": "Layanan sedang bermasalah"})
			}

		}
		return c.JSON(http.StatusOK, map[string]string{"message": "Berhasil sinkron ulang pegawai"})
	}
	return echo.HandlerFunc(h)
}

func SendPegawaiToOracle(a *app.App, ctx context.Context, uuid string) error {
	flagSinkronSimpeg, err := pengaturan.LoadPengaturan(a, ctx, nil, pengaturanAtributFlagSinkronSimpeg)
	if err != nil {
		return fmt.Errorf("error load pengaturan flag sinkron simpeg: %w", err)
	}

	if flagSinkronSimpeg != "1" {
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
