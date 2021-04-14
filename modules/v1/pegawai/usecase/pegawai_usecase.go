package usecase

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"net/http"
	"svc-insani-go/app"
	pegawaiOraHttp "svc-insani-go/modules/v1/pegawai-oracle/http"
	pegawaiOraModel "svc-insani-go/modules/v1/pegawai-oracle/model"
	"svc-insani-go/modules/v1/pegawai/model"
	"svc-insani-go/modules/v1/pegawai/repo"

	"github.com/labstack/echo"
)

func HandleGetPegawai(a app.App) echo.HandlerFunc {
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
			fmt.Printf("[ERROR] repo count pegawai, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		if count == 0 {
			return c.JSON(http.StatusOK, res)
		}
		pp, err := repo.GetAllPegawai(a, req)
		if err != nil {
			fmt.Printf("[ERROR] repo get all pegawai, %s\n", err.Error())
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

func HandleGetSimpegPegawaiByUUID(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		uuidPegawai := c.Param("uuidPegawai")
		if uuidPegawai == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "parameter uuid pegawai wajib diisi"})
		}

		pegawaiDetail, err := PrepareGetSimpegPegawaiByUUID(a, uuidPegawai)
		if err != nil {
			fmt.Printf("[ERROR] repo get kepegawaian, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		return c.JSON(http.StatusOK, pegawaiDetail)
	}
	return echo.HandlerFunc(h)
}

func PrepareGetSimpegPegawaiByUUID(a app.App, uuidPegawai string) (model.PegawaiDetail, error) {
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

	pegawaiPendidikan, err := repo.GetPegawaiPendidikan(a, uuidPegawai)
	if err != nil {
		return model.PegawaiDetail{}, fmt.Errorf("error repo get pendidikan pegawai by uuid, %w", err)
	}

	pegawaiDetail.PegawaiYayasan = kepegawaianYayasan
	pegawaiDetail.UnitKerjaPegawai = unitKerjaPegawai
	pegawaiDetail.StatusAktif = statusPegawaiAktif
	pegawaiDetail.PegawaiPribadi = pegawaiPribadi
	pegawaiDetail.JenjangPendidikan = pegawaiPendidikan

	pegawaiDetail.PegawaiPNSPTT = pegawaiPNS

	return pegawaiDetail, nil
}

func HandleUpdateSimpegPegawaiByUUID(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Perubahan pegawai berhasil disimpan"})
	}
	return echo.HandlerFunc(h)
}

// Get All Pegawai With GORM
func HandleGetPegawaix(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		limit, err := strconv.Atoi(c.QueryParam("limit"))
		if err != nil {
			fmt.Printf("[ERROR] convert string to int, %s\n", err.Error())
		}

		offset, err := strconv.Atoi(c.QueryParam("offset"))
		if err != nil {
			fmt.Printf("[ERROR] convert string to int, %s\n", err.Error())
		}

		pp, err := repo.GetAllPegawaix(a, c.Request().Context(), limit, offset)
		if err != nil {
			fmt.Printf("[ERROR] repo get all pegawai, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		return c.JSON(http.StatusOK, pp)
	}
	return echo.HandlerFunc(h)
}

// Get Pegawai With GORM
func HandleGetPegawaiByUUIDx(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		uuidPersonal := c.Param("uuidPersonal")
		pp, err := repo.GetPegawaiByUUIDx(a, c.Request().Context(), uuidPersonal)
		if err != nil {
			fmt.Printf("[ERROR] repo get pegawai by uuid, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		return c.JSON(http.StatusOK, pp)
	}
	return echo.HandlerFunc(h)
}

func HandleUpdatePegawai(a app.App, ctx context.Context, errChan chan error) echo.HandlerFunc {
	h := func(c echo.Context) error {

		// Validasi Data
		pegawaiUpdate, err := ValidateUpdatePegawaiByUUID(a, c)
		if err != nil {
			fmt.Printf("[ERROR], %s\n", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}

		// Update Data
		err = repo.UpdatePegawaix(a, c.Request().Context(), pegawaiUpdate)
		if err != nil {
			fmt.Printf("[ERROR], %s\n", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}

		// Set Flag Pendidikan
		uuidPendidikanDiakui := c.FormValue("uuid_tingkat_pdd_diakui")
		uuidPendidikanTerakhir := c.FormValue("uuid_tingkat_pdd_terakhir")

		if uuidPendidikanDiakui != "" && uuidPendidikanTerakhir != "" {
			err = repo.UpdatePendidikanPegawai(a, c.Request().Context(), uuidPendidikanDiakui, uuidPendidikanTerakhir)
			if err != nil {
				fmt.Printf("[ERROR], %s\n", err.Error())
				return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
			}
		}

		// Menampilkan response
		pegawaiDetail, err := PrepareGetSimpegPegawaiByUUID(a, pegawaiUpdate.Uuid)
		if err != nil {
			fmt.Printf("[ERROR] repo get kepegawaian, %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		go func(
			a app.App,
			ctx context.Context,
			errChan chan error,
		) {
			fmt.Println("DEBUG : Go routin")
			dur, err := time.ParseDuration(os.Getenv("RESPONSE_TIMEOUT_MS" + "ms"))
			if err != nil {
				dur = time.Second * 40
			}
			ctx, cancel := context.WithTimeout(ctx, dur)
			// ctx, cancel := context.WithTimeout(context.Background(), dur) // kalau ke cancel pake yang ini
			defer cancel()
			fmt.Println("DEBUG : Go routin before prepare sipeg")
			pegawaiDetail, err := PrepareGetSimpegPegawaiByUUID(a, pegawaiUpdate.Uuid)
			if err != nil {
				errChan <- err
				return
			}

			fmt.Println("DEBUG : Go routin before sinkron simpeg")

			err = prepareSinkronSimpeg(ctx, pegawaiDetail)
			if err != nil {
				errChan <- err
				return
			}
		}(a, ctx, errChan)

		return c.JSON(http.StatusOK, pegawaiDetail)
	}

	return echo.HandlerFunc(h)
}

func prepareSinkronSimpeg(ctx context.Context, pegawaiInsani model.PegawaiDetail) error {

	pegawaiOra := pegawaiOraModel.KepegawaianYayasanSimpeg{}

	// Sinkron Kepegawaian Yayaysan - Status

	pegawaiOra.KdJenisPegawai = pegawaiInsani.PegawaiYayasan.KDJenisPegawai
	pegawaiOra.KdStatusPegawai = pegawaiInsani.PegawaiYayasan.StatusPegawai
	pegawaiOra.KdKelompokPegawai = pegawaiInsani.PegawaiYayasan.KdKelompokPegawai

	// Sinkron Kepegawaian Yayaysan - Pangkat / Jabatan

	pegawaiOra.PangkatYayasan.KdGolongan = pegawaiInsani.PegawaiYayasan.Golongan
	pegawaiOra.PangkatYayasan.KdRuang = pegawaiInsani.PegawaiYayasan.KdRuang
	pegawaiOra.PangkatYayasan.TmtPangkat = pegawaiInsani.PegawaiYayasan.TmtPangkatGolongan
	pegawaiOra.KdFungsional = pegawaiInsani.PegawaiYayasan.KdJabatanFungsional
	pegawaiOra.TmtFungsional = pegawaiInsani.PegawaiYayasan.TmtJabatan
	pegawaiOra.MasaKerjaGajiTahun, _ = strconv.Atoi(pegawaiInsani.PegawaiYayasan.MasaKerjaBawaanTahun)
	pegawaiOra.MasaKerjaGajiBulan, _ = strconv.Atoi(pegawaiInsani.PegawaiYayasan.MasaKerjaBawaanBulan)
	pegawaiOra.AngkaKreditFungsional, _ = strconv.ParseFloat(pegawaiInsani.PegawaiYayasan.AngkaKredit, 64)

	// Sinkron Unit Kerja

	pegawaiOra.Unit1.KdUnit1 = pegawaiInsani.UnitKerjaPegawai.KdIndukKerja
	pegawaiOra.Unit2.KdUnit2 = pegawaiInsani.UnitKerjaPegawai.KdUnitKerja
	pegawaiOra.Unit3.KdUnit3 = pegawaiInsani.UnitKerjaPegawai.KdBagianKerja
	pegawaiOra.NoSkPertama = pegawaiInsani.UnitKerjaPegawai.NoSkPertama
	pegawaiOra.TglSkPertama = pegawaiInsani.UnitKerjaPegawai.TmtSkPertama
	pegawaiOra.KdHomebasePddikti = pegawaiInsani.UnitKerjaPegawai.KdHomebasePddikti
	pegawaiOra.KdHomebaseUii = pegawaiInsani.UnitKerjaPegawai.KdHomebaseUii

	// Sinkron Kepegawaian Negara / PTT

	pegawaiOra.Instansi = pegawaiInsani.PegawaiPNSPTT.InstansiAsalPtt
	pegawaiOra.NipKopertis = pegawaiInsani.PegawaiPNSPTT.NipPNS
	pegawaiOra.NoKarpeg = pegawaiInsani.PegawaiPNSPTT.NoKartuPegawai
	pegawaiOra.PangkatKopertis.KdGolongan = pegawaiInsani.PegawaiPNSPTT.KdPangkatGolongan
	pegawaiOra.PangkatKopertis.KdRuang = pegawaiInsani.PegawaiPNSPTT.KdPangkatGolongan
	pegawaiOra.PangkatKopertis.TmtPangkat = pegawaiInsani.PegawaiPNSPTT.TmtPangkatGolongan
	pegawaiOra.KdFungsional = pegawaiInsani.PegawaiPNSPTT.KdJabatanPns
	pegawaiOra.TmtFungsional = pegawaiInsani.PegawaiPNSPTT.TmtPangkatGolongan
	pegawaiOra.MasaKerjaKopertisTahun, _ = strconv.Atoi(pegawaiInsani.PegawaiPNSPTT.MasaKerjaPnsTahun)
	pegawaiOra.MasaKerjaKopertisBulan, _ = strconv.Atoi(pegawaiInsani.PegawaiPNSPTT.MasaKerjaPnsBulan)
	pegawaiOra.AngkaKreditKopertis, _ = strconv.ParseFloat(pegawaiInsani.PegawaiPNSPTT.AngkaKreditPns, 64)
	pegawaiOra.Keterangan = pegawaiInsani.PegawaiPNSPTT.KeteranganPNS

	// Sinkron Status Aktif

	if pegawaiInsani.StatusAktif.FlagAktifPegawai == "1" {
		if pegawaiInsani.PegawaiYayasan.KDJenisPegawai == "ED" {
			pegawaiOra.FlagMengajar = "Y"
			if pegawaiInsani.StatusAktif.KdStatusAktifPegawai == "IBL" {
				pegawaiOra.FlagSekolah = "Y"
				pegawaiOra.FlagPensiun = "N"
			}
			if pegawaiInsani.StatusAktif.KdStatusAktifPegawai != "IBL" {
				pegawaiOra.FlagSekolah = "N"
				pegawaiOra.FlagPensiun = "Y"
			}

		}
	}

	if pegawaiInsani.StatusAktif.FlagAktifPegawai == "0" {
		if pegawaiInsani.PegawaiYayasan.KDJenisPegawai == "ED" {
			pegawaiOra.FlagMengajar = "N"
			if pegawaiInsani.StatusAktif.KdStatusAktifPegawai == "PEN" {
				pegawaiOra.FlagSekolah = "N"
				pegawaiOra.FlagPensiun = "Y"
			}
			if pegawaiInsani.StatusAktif.KdStatusAktifPegawai != "PEN" {
				pegawaiOra.FlagSekolah = "N"
				pegawaiOra.FlagPensiun = "N"
			}
			if pegawaiInsani.StatusAktif.KdStatusAktifPegawai != "MNG" {
				pegawaiOra.KdStatusHidup = "N"
			}
			pegawaiOra.KdStatusHidup = "Y"
		}
	}

	fmt.Println("DEBUG : Update Kepegawaian Yayasan")

	err := pegawaiOraHttp.UpdateKepegawaianYayasan(ctx, &http.Client{}, pegawaiOra)
	if err != nil {
		return fmt.Errorf("[ERROR] repo get kepegawaian yayasan update, %s\n", err.Error())
		// return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
	}

	return nil
}
