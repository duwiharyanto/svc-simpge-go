package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/pegawai/model"
	"svc-insani-go/modules/v1/pegawai/repo"
	personalRepo "svc-insani-go/modules/v1/personal/repo"

	ptr "github.com/openlyinc/pointy"

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
	// Get presign URL file foto
	pegawaiPribadi.UrlFileFoto, err = repo.GetPresignUrlFotoPegawai(a, pegawaiPribadi.UrlFileFoto)
	if err != nil {
		return model.PegawaiDetail{}, fmt.Errorf("error repo get persign url foto pegawai, %w", err)
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
		pegawai, err := ValidateUpdatePegawaiByUUID(a, c)
		if err != nil {
			fmt.Printf("[ERROR]: %s\n", err.Error())
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}

		// Update Data
		err = repo.UpdatePegawai(a, c.Request().Context(), pegawai)
		if err != nil {
			fmt.Printf("[ERROR] update pegawai: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		// Set Flag Pendidikan
		uuidPendidikanDiakui := c.FormValue("uuid_tingkat_pdd_diakui")     // uuid dari pendidikan yang dipilih sbg ijazah tertinggi diakui
		uuidPendidikanTerakhir := c.FormValue("uuid_tingkat_pdd_terakhir") // uuid dari pendidikan yang dipilih sbg ijazah terakhir
		idPersonalPegawai := pegawai.IdPersonalDataPribadi

		err = repo.UpdatePendidikanPegawai(a, c.Request().Context(),
			model.PegawaiPendidikanRequest{
				UuidPendidikanDiakui:                 uuidPendidikanDiakui,
				UuidPendidikanTerakhir:               uuidPendidikanTerakhir,
				IdJenjangPendidikanDetailDiakui:      pegawai.IdStatusPendidikanMasuk,
				IdJenjangPendidikanDetailTerakhir:    pegawai.IdJenisPendidikan,
				UuidJenjangPendidikanTertinggiDiakui: ptr.StringValue(pegawai.UuidPendidikanMasuk, ""),
				IdPersonalPegawai:                    ptr.Uint64Value(idPersonalPegawai, 0),
				UserUpdate:                           pegawai.UserUpdate,
			})
		if err != nil {
			fmt.Printf("[ERROR] update pendidikan pegawai: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		// Menampilkan response
		pegawaiDetail, err := PrepareGetSimpegPegawaiByUUID(a, ptr.StringValue(pegawai.Uuid, ""))
		if err != nil {
			fmt.Printf("[ERROR] repo get kepegawaian: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		fmt.Printf("[DEBUG] update response end\n")
		// go func(a *app.App, ctx context.Context, errChan chan error) {
		// 	defer func(n time.Time) {
		// 		fmt.Printf("[DEBUG] send to simpeg: %v ms\n", time.Now().Sub(n).Milliseconds())
		// 	}(time.Now())
		// 	fmt.Println("[DEBUG] Go routine start after update")

		// 	flagSinkronSimpeg, err := pengaturan.LoadPengaturan(a, ctx, nil, pengaturanAtributFlagSinkronSimpeg)
		// 	if err != nil {
		// 		log.Println("error load pengaturan flag sinkron simpeg: %w", err)
		// 		errChan <- err
		// 		return
		// 	}

		// 	disableSyncSimpegOracle, _ := strconv.ParseBool(os.Getenv("DISABLE_SYNC_SIMPEG_ORACLE"))
		// 	if flagSinkronSimpeg != "1" || disableSyncSimpegOracle {
		// 		log.Printf("[DEBUG] flag sinkron simpeg 0\n")
		// 		return
		// 	}

		// 	dur, err := time.ParseDuration(os.Getenv("RESPONSE_TIMEOUT_MS" + "ms"))
		// 	if err != nil {
		// 		dur = time.Second * 40
		// 	}
		// 	ctx, cancel := context.WithTimeout(ctx, dur)
		// 	// ctx, cancel := context.WithTimeout(context.Background(), dur) // kalau ke cancel pake yang ini
		// 	defer cancel()

		// 	// fmt.Println("DEBUG : Go routin before sinkron simpeg")
		// 	pegawaiOra := newPegawaiOra(&pegawaiDetail)
		// 	err = pegawaiOraHttp.UpdateKepegawaianYayasan(ctx, &http.Client{}, pegawaiOra)
		// 	if err != nil {
		// 		errChan <- fmt.Errorf("[ERROR] repo update kepegawaian yayasan: %w\n", err)
		// 		return
		// 	}
		// }(a, ctx, errChan)

		return c.JSON(http.StatusOK, pegawaiDetail)
	}

	return echo.HandlerFunc(h)
}

func HandleCreatePegawai(a *app.App, ctx context.Context, errChan chan error) echo.HandlerFunc {
	h := func(c echo.Context) error {
		// Validasi Data
		pegawai, err := PrepareCreateSimpeg(a, c)
		if errors.Unwrap(err) != nil {
			fmt.Printf("[ERROR] prepare create simpeg: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}

		// Create Data
		err = repo.CreatePegawai(a, c.Request().Context(), pegawai)
		if errors.Unwrap(err) != nil && strings.Contains(err.Error(), "presensi") {
			fmt.Printf("[ERROR] prepare create simpeg: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Gagal simpan user presensi pegawai"})
		}
		if err != nil {
			fmt.Printf("[ERROR] create pegawai: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		// Menampilkan response
		pegawaiDetail, err := PrepareGetSimpegPegawaiByUUID(a, pegawai.Uuid)
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
			// err = personalRepo.PersonalActivation(c.FormValue("uuid_personal"))
			// if err != nil {
			// 	errChan <- err
			// 	return
			// }
		}(a, errChan, pegawai.Uuid)

		go func(uuidPersonal string) {
			err = personalRepo.PersonalActivation(c.FormValue("uuid_personal"))
		}(c.FormValue("uuid_personal"))

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
