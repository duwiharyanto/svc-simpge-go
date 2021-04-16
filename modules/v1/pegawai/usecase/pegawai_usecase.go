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

// func HandleUpdateSimpegPegawaiByUUID(a app.App) echo.HandlerFunc {
// 	h := func(c echo.Context) error {
// 		return c.JSON(http.StatusOK, map[string]string{"message": "Perubahan pegawai berhasil disimpan"})
// 	}
// 	return echo.HandlerFunc(h)
// }

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
		err = repo.UpdatePegawai(a, c.Request().Context(), pegawaiUpdate)
		if err != nil {
			fmt.Printf("[ERROR], %s\n", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}

		// Set Flag Pendidikan
		uuidPendidikanDiakui := c.FormValue("uuid_tingkat_pdd_diakui")
		uuidPendidikanTerakhir := c.FormValue("uuid_tingkat_pdd_terakhir")
		idPersonalPegawai := pegawaiUpdate.IdPersonalDataPribadi

		fmt.Println("Ini Id Pegawai : ", idPersonalPegawai)

		if uuidPendidikanDiakui != "" && uuidPendidikanTerakhir != "" {
			err = repo.UpdatePendidikanPegawai(a, c.Request().Context(), uuidPendidikanDiakui, uuidPendidikanTerakhir, idPersonalPegawai)
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
			// fmt.Println("DEBUG : Go routin before prepare sipeg")
			pegawaiDetail, err := PrepareGetSimpegPegawaiByUUID(a, pegawaiUpdate.Uuid)
			if err != nil {
				errChan <- err
				return
			}

			fmt.Println("DEBUG : Go routin before sinkron simpeg")
			// fmt.Printf("DEBUG Pegawai Detail %+v \n", &pegawaiDetail)
			err = prepareSinkronSimpeg(ctx, &pegawaiDetail)
			if err != nil {
				errChan <- err
				return
			}
		}(a, ctx, errChan)

		return c.JSON(http.StatusOK, pegawaiDetail)
	}

	return echo.HandlerFunc(h)
}

func prepareSinkronSimpeg(ctx context.Context, pegawaiInsani *model.PegawaiDetail) error {

	pegawaiOra := &pegawaiOraModel.KepegawaianYayasanSimpeg{}
	pegawaiOra.JenisPegawai = &pegawaiOraModel.JenisPegawai{}
	pegawaiOra.InstansiAsalPtt = &pegawaiOraModel.InstansiAsalPtt{}
	pegawaiOra.JenisPegawai = &pegawaiOraModel.JenisPegawai{}
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

	// fmt.Printf("DEBUG pegawai yayasan %+v \n: ", pegawaiInsani.PegawaiYayasan)
	// fmt.Printf("DEBUG pegawai yayasan %+v \n: ", pegawaiInsani.PegawaiYayasan.KDJenisPegawai)

	// Sinkron NIP Pegawai
	pegawaiOra.NIP = pegawaiInsani.PegawaiPribadi.NIK

	// Sinkron Kepegawaian Yayaysan - Status
	if pegawaiInsani.PegawaiYayasan.KDJenisPegawai != "" {
		pegawaiOra.KdJenisPegawai = pegawaiInsani.PegawaiYayasan.KDJenisPegawai
	}
	// fmt.Printf("DEBUG Kd Jenis %+v \n: \n ", pegawaiOra.KdJenisPegawai)

	if pegawaiInsani.PegawaiYayasan.StatusPegawai != "" {
		pegawaiOra.KdStatusPegawai = pegawaiInsani.PegawaiYayasan.KDStatusPegawai
	}

	// fmt.Printf("DEBUG fmt %+v \n: \n ", pegawaiOra.KdStatusPegawai)

	if pegawaiInsani.PegawaiYayasan.KdKelompokPegawai != "" {
		pegawaiOra.KdKelompokPegawai = pegawaiInsani.PegawaiYayasan.KdKelompokPegawai
	}
	// fmt.Printf("DEBUG pegawaiOra.KdKelompokPegawai %+v \n: \n ", pegawaiOra.KdKelompokPegawai)

	// Sinkron Kepegawaian Yayaysan - Pangkat / Jabatan

	if pegawaiInsani.PegawaiYayasan.KdGolongan != "" {
		pegawaiOra.PangkatYayasan.KdGolongan = pegawaiInsani.PegawaiYayasan.KdGolongan
	}
	// fmt.Printf("DEBUG pegawaiOra.PangkatYayasan %+v \n: \n ", pegawaiOra.PangkatYayasan.KdGolongan)

	if pegawaiInsani.PegawaiYayasan.KdRuang != "" {
		pegawaiOra.PangkatYayasan.KdRuang = pegawaiInsani.PegawaiYayasan.KdRuang
	}
	// fmt.Printf("DEBUG pegawaiOra.PangkatYayasan %+v \n: \n ", pegawaiOra.PangkatYayasan.KdRuang)

	if pegawaiInsani.PegawaiYayasan.TmtPangkatGolongan != "" {
		pegawaiOra.PangkatYayasan.TmtPangkat = pegawaiInsani.PegawaiYayasan.TmtPangkatGolongan
	}
	// fmt.Printf("DEBUG pegawaiOra.PangkatYayasan %+v \n: \n ", pegawaiOra.PangkatYayasan.TmtPangkat)

	if pegawaiInsani.PegawaiYayasan.KdJabatanFungsional != "" {
		pegawaiOra.JabatanFungsional.KdFungsional = pegawaiInsani.PegawaiYayasan.KdJabatanFungsional
	}
	// fmt.Printf("DEBUG pegawaiOra.KdFungsional %+v \n: \n ", pegawaiOra.JabatanFungsional.KdFungsional)

	if pegawaiInsani.PegawaiYayasan.TmtJabatan != "" {
		pegawaiOra.TmtFungsional = pegawaiInsani.PegawaiYayasan.TmtJabatan
	}
	// fmt.Printf("DEBUG pegawaiOra.TmtFungsional %+v \n: \n ", pegawaiOra.TmtFungsional)

	if pegawaiInsani.PegawaiYayasan.MasaKerjaGajiTahun != "" {
		pegawaiOra.MasaKerjaGajiTahun, _ = strconv.Atoi(pegawaiInsani.PegawaiYayasan.MasaKerjaGajiTahun)
	}
	// fmt.Printf("DEBUG pegawaiOra.MasaKerjaGajiTahun %+v \n: \n ", pegawaiOra.MasaKerjaGajiTahun)

	if pegawaiInsani.PegawaiYayasan.MasaKerjaGajiBulan != "" {
		pegawaiOra.MasaKerjaGajiBulan, _ = strconv.Atoi(pegawaiInsani.PegawaiYayasan.MasaKerjaGajiBulan)
	}
	// fmt.Printf("DEBUG pegawaiOra.MasaKerjaGajiBulan %+v \n: \n ", pegawaiOra.MasaKerjaGajiBulan)

	if pegawaiInsani.PegawaiYayasan.AngkaKredit != "" {
		pegawaiOra.AngkaKreditFungsional, _ = strconv.ParseFloat(pegawaiInsani.PegawaiYayasan.AngkaKredit, 64)
	}
	// fmt.Printf("DEBUG pegawaiOra.AngkaKreditFungsional %+v \n: \n ", pegawaiOra.AngkaKreditFungsional)

	// Sinkron Unit Kerja

	if pegawaiInsani.UnitKerjaPegawai.KdIndukKerja != "" {
		pegawaiOra.Unit1.KdUnit1 = pegawaiInsani.UnitKerjaPegawai.KdIndukKerja
	}
	// fmt.Printf("DEBUG pegawaiOra.Unit1 %+v \n: \n ", pegawaiOra.Unit1.KdUnit1)

	if pegawaiInsani.UnitKerjaPegawai.KdUnitKerja != "" {
		pegawaiOra.Unit2.KdUnit2 = pegawaiInsani.UnitKerjaPegawai.KdUnitKerja
	}
	// fmt.Printf("DEBUG pegawaiOra.Unit2 %+v \n: \n ", pegawaiOra.Unit2.KdUnit2)

	if pegawaiInsani.UnitKerjaPegawai.KdBagianKerja != "" {
		pegawaiOra.Unit3.KdUnit3 = pegawaiInsani.UnitKerjaPegawai.KdBagianKerja
	}
	// fmt.Printf("DEBUG pegawaiOra.Unit3 %+v \n: \n ", pegawaiOra.Unit3.KdUnit3)

	if pegawaiInsani.UnitKerjaPegawai.NoSkPertama != "" {
		pegawaiOra.NoSkPertama = pegawaiInsani.UnitKerjaPegawai.NoSkPertama
	}
	// fmt.Printf("DEBUG pegawaiOra.NoSkPertama %+v \n: \n ", pegawaiOra.NoSkPertama)

	if pegawaiInsani.UnitKerjaPegawai.TmtSkPertama != "" {
		pegawaiOra.TglSkPertama = pegawaiInsani.UnitKerjaPegawai.TmtSkPertama
	}
	// fmt.Printf("DEBUG pegawaiOra.TglSkPertama %+v \n: \n ", pegawaiOra.TglSkPertama)

	if pegawaiInsani.UnitKerjaPegawai.KdHomebasePddikti != "" {
		pegawaiOra.KdHomebasePddikti = pegawaiInsani.UnitKerjaPegawai.KdHomebasePddikti
	}
	// fmt.Printf("DEBUG pegawaiOra.KdHomebasePddikti %+v \n: \n ", pegawaiOra.KdHomebasePddikti)

	if pegawaiInsani.UnitKerjaPegawai.KdHomebaseUii != "" {
		pegawaiOra.KdHomebaseUii = pegawaiInsani.UnitKerjaPegawai.KdHomebaseUii
	}
	// fmt.Printf("DEBUG pegawaiOra.KdHomebaseUii %+v \n: \n ", pegawaiOra.KdHomebaseUii)

	// Sinkron Kepegawaian Negara / PTT

	if pegawaiInsani.PegawaiPNSPTT.InstansiAsalPtt != "" {
		pegawaiOra.Instansi = pegawaiInsani.PegawaiPNSPTT.InstansiAsalPtt
	}
	// fmt.Printf("DEBUG pegawaiOra.Instansi %+v \n: \n ", pegawaiOra.Instansi)

	if pegawaiInsani.PegawaiPNSPTT.NipPNS != "" {
		pegawaiOra.NipKopertis = pegawaiInsani.PegawaiPNSPTT.NipPNS
	}
	// fmt.Printf("DEBUG pegawaiOra.NipKopertis %+v \n: \n ", pegawaiOra.NipKopertis)

	if pegawaiInsani.PegawaiPNSPTT.NoKartuPegawai != "" {
		pegawaiOra.NoKarpeg = pegawaiInsani.PegawaiPNSPTT.NoKartuPegawai
	}
	// fmt.Printf("DEBUG pegawaiOra.NoKarpeg %+v \n: \n ", pegawaiOra.NoKarpeg)

	if pegawaiInsani.PegawaiPNSPTT.KdGolonganPNS != "" {
		pegawaiOra.PangkatKopertis.KdGolongan = pegawaiInsani.PegawaiPNSPTT.KdGolonganPNS
	}
	fmt.Printf("DEBUG pegawaiOra.PangkatKopertis.KdGolongan %+v \n: \n ", pegawaiOra.PangkatKopertis.KdGolongan)

	if pegawaiInsani.PegawaiPNSPTT.KdPangkatGolonganPns != "" {
		pegawaiOra.PangkatKopertis.KdRuang = pegawaiInsani.PegawaiPNSPTT.KdPangkatGolonganPns
	}
	// fmt.Printf("DEBUG pegawaiOra.PangkatKopertis %+v \n: \n ", pegawaiOra.PangkatKopertis.KdRuang)

	if pegawaiInsani.PegawaiPNSPTT.TmtPangkatGolongan != "" {
		pegawaiOra.PangkatKopertis.TmtPangkat = pegawaiInsani.PegawaiPNSPTT.TmtPangkatGolongan
	}
	// fmt.Printf("DEBUG pegawaiOra.PangkatKopertis %+v \n: \n ", pegawaiOra.PangkatKopertis.TmtPangkat)

	if pegawaiInsani.PegawaiPNSPTT.KdJabatanPns != "" {
		pegawaiOra.JabatanFungsionalKopertis.KdFungsional = pegawaiInsani.PegawaiPNSPTT.KdJabatanPns
	}
	// fmt.Printf("DEBUG pegawaiOra.KdFungsional %+v \n: \n ", pegawaiOra.JabatanFungsionalKopertis.KdFungsional)

	if pegawaiInsani.PegawaiPNSPTT.TmtPangkatGolongan != "" {
		pegawaiOra.TmtFungsional = pegawaiInsani.PegawaiPNSPTT.TmtPangkatGolongan
	}
	// fmt.Printf("DEBUG pegawaiOra.TmtFungsional %+v \n: \n ", pegawaiOra.TmtFungsional)

	if pegawaiInsani.PegawaiPNSPTT.MasaKerjaPnsTahun != "" {
		pegawaiOra.MasaKerjaKopertisTahun, _ = strconv.Atoi(pegawaiInsani.PegawaiPNSPTT.MasaKerjaPnsTahun)
	}
	// fmt.Printf("DEBUG pegawaiOra.MasaKerjaKopertisTahun %+v \n: \n ", pegawaiOra.MasaKerjaKopertisTahun)

	if pegawaiInsani.PegawaiPNSPTT.MasaKerjaPnsBulan != "" {
		pegawaiOra.MasaKerjaKopertisBulan, _ = strconv.Atoi(pegawaiInsani.PegawaiPNSPTT.MasaKerjaPnsBulan)
	}
	// fmt.Printf("DEBUG pegawaiOra.MasaKerjaKopertisBulan %+v \n: \n ", pegawaiOra.MasaKerjaKopertisBulan)

	if pegawaiInsani.PegawaiPNSPTT.AngkaKreditPns != "" {
		pegawaiOra.AngkaKreditKopertis, _ = strconv.ParseFloat(pegawaiInsani.PegawaiPNSPTT.AngkaKreditPns, 64)
	}
	// fmt.Printf("DEBUG pegawaiOra.AngkaKreditKopertis %+v \n: \n ", pegawaiOra.AngkaKreditKopertis)

	if pegawaiInsani.PegawaiPNSPTT.KeteranganPNS != "" {
		pegawaiOra.Keterangan = pegawaiInsani.PegawaiPNSPTT.KeteranganPNS
	}
	// fmt.Printf("DEBUG pegawaiOra.Keterangan %+v \n: \n ", pegawaiOra.Keterangan)

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
