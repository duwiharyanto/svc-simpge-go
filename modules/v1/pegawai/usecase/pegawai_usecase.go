package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/pegawai/model"
	"svc-insani-go/modules/v1/pegawai/repo"
	pegawaiRepo "svc-insani-go/modules/v1/pegawai/repo"
	pengaturan "svc-insani-go/modules/v1/pengaturan-insani/usecase"
	personalRepo "svc-insani-go/modules/v1/personal/repo"
	pegawaiOraHttp "svc-insani-go/modules/v1/simpeg-oracle/http"
	_ "svc-insani-go/modules/v2/organisasi/model"
	organisaiPrivate "svc-insani-go/modules/v2/organisasi/model"
	_ "svc-insani-go/modules/v2/organisasi/repo"

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

func HandleGetSimpegPegawaiDetail(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		user := c.Request().Header.Get("X-Member")
		if user == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "pengguna tidak valid"})
		}

		appCtx := context.Background()
		pegawai, err := pegawaiRepo.GetPegawaiByNIK(a, appCtx, user)
		if err != nil {
			log.Printf("[ERROR] repo get kepegawaian: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		pegawaiDetail, err := PrepareGetSimpegPegawaiByUUID(a, pegawai.UUID)
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

		go func(a *app.App, ctx context.Context, errChan chan error) {
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

			disableSyncSimpegOracle, _ := strconv.ParseBool(os.Getenv("DISABLE_SYNC_SIMPEG_ORACLE"))
			if flagSinkronSimpeg != "1" || disableSyncSimpegOracle {
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

func HandleGetPegawaiByNik(a *app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		nik := c.Param("nik")
		if nik == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "parameter nik wajib diisi"})
		}

		pegawai, err := pegawaiRepo.GetPegawaiByNikPrivate(a, nik)
		if err != nil {
			log.Printf("[ERROR] repo get kepegawaian: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		data := model.PegawaiByNikResponse{
			Status:  200,
			Pegawai: pegawai,
		}

		return c.JSON(http.StatusOK, data)
	}
	return echo.HandlerFunc(h)
}

func HandleGetPegawaiPrivate(a *app.App, public bool) echo.HandlerFunc {
	h := func(c echo.Context) error {
		env := os.Getenv("ENV")
		fmt.Println(env)
		if public {
			if env == "staging" || env == "production" {
				return c.JSON(404, "layanan tidak ditemukan")
			}
		}

		reqNik := c.QueryParam("nik")
		var nik string
		if len(reqNik) > 0 {
			nik = reqNik[:len(reqNik)-1]
			nik = nik[1:]
		}

		req := &model.PegawaiPrivateRequest{}
		err := c.Bind(req)
		if err != nil {
			fmt.Printf("[WARNING] binding pegawai request: %s\n", err.Error())
		}

		res := model.PegawaiPrivateResponse{
			Data: []model.PegawaiPrivate{},
		}
		req.Nik = nik

		pp, err := repo.GetAllPegawaiPrivate(a, req)
		if err != nil {
			fmt.Printf("[ERROR] repo get all pegawai: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		res.Data = pp

		// get data jabatan fungsional
		var pegawaiAndFungsional []model.PegawaiPrivate
		var fullJabatanFungsionalDataItem = model.PegawaiFungsionalDataItemY{}
		for _, data := range res.Data {
			pegawaiFungsionalYayasan, err := repo.GetJabatanFungsionalYayasan(a, strconv.FormatInt(int64(data.IdPegawai), 10))
			if err != nil {
				fmt.Println(err)
				return c.JSON(500, nil)
			}

			pegawaiFungsionalNegara, err := repo.GetJabatanFungsionalNegara(a, strconv.FormatInt(int64(data.IdPegawai), 10))
			if err != nil {
				fmt.Println(err)
				return c.JSON(500, nil)
			}

			fullJabatanFungsionalDataItem.FullPegawaiFungsional.PegawaiFungsionalYayasan = pegawaiFungsionalYayasan
			if pegawaiFungsionalYayasan == nil {
				fullJabatanFungsionalDataItem.FullPegawaiFungsional.PegawaiFungsionalYayasan = &model.PegawaiFungsionalYayasan{}
			}

			fullJabatanFungsionalDataItem.FullPegawaiFungsional.PegawaiFungsionalNegara = pegawaiFungsionalNegara

			if pegawaiFungsionalNegara == nil {
				// fmt.Println("kosong")
				fullJabatanFungsionalDataItem.FullPegawaiFungsional.PegawaiFungsionalNegara = &model.PegawaiFungsionalNegara{}
			}

			data.JabatanFungsional = fullJabatanFungsionalDataItem
			pegawaiAndFungsional = append(pegawaiAndFungsional, data)
		}

		// get data jabatan struktural
		stmt, err := a.DB.Prepare(`SELECT COALESCE(p.id,0), COALESCE(po.id_jenis_jabatan,0),COALESCE(po.id_unit,0),COALESCE(u.id_jenis_unit,0) FROM pegawai p JOIN hcm_organisasi.pejabat_organisasi po ON po.id_pegawai = p.id JOIN hcm_organisasi.unit u ON u.id = po.id_unit`)

		var pejabat []organisaiPrivate.PejabatStrukturalPrivate
		rows, err := stmt.Query()
		if err != nil {
			fmt.Println(err)
			return c.JSON(500, nil)
		}
		defer rows.Close()
		// Loop through rows, using Scan to assign column data to struct fields.
		for rows.Next() {
			var ps organisaiPrivate.PejabatStrukturalPrivate
			if err := rows.Scan(&ps.IdPegawai, &ps.IdJenisUnit, &ps.IdJenisJabatan, &ps.IdUnit); err != nil {
				fmt.Println(err)
				return c.JSON(500, nil)
			}
			pejabat = append(pejabat, ps)
		}
		if err := rows.Err(); err != nil {
			fmt.Println(err)
			return c.JSON(500, nil)
		}

		var pegawaiAndFungsionalAndStruktural []model.PegawaiPrivate
		IsNotStruktural := true
		for _, data := range pegawaiAndFungsional {

			tmtSkPertamaTime, err := time.Parse("2006-01-02", data.TmtSkPertama)
			var tmtSkPertamaDuration time.Duration
			if err == nil {
				tmtSkPertamaDuration = time.Now().Sub(tmtSkPertamaTime)
			}
			tmtSkPertamaDurationDays := tmtSkPertamaDuration.Hours() / 24
			// tmtSkPertamaDurationDays := tmtSkPertamaDuration.Hours() / 24
			tmtSkPertamaDurationRealMonths := int(tmtSkPertamaDurationDays / 365 * 12)
			masaKerjaKepegawaianTahunInt, _ := strconv.Atoi(data.MasaKerjaTahun)
			masaKerjaKepegawaianBulanInt, _ := strconv.Atoi(data.MasaKerjaBulan)
			masaKerjaTotalKepegawaianRealBulan := ((masaKerjaKepegawaianTahunInt * 12) + masaKerjaKepegawaianBulanInt) + tmtSkPertamaDurationRealMonths
			data.MasaKerjaTahun = fmt.Sprintf("%d", masaKerjaTotalKepegawaianRealBulan/12)
			data.MasaKerjaBulan = fmt.Sprintf("%d", masaKerjaTotalKepegawaianRealBulan%12)

			for _, pejabat := range pejabat {
				if data.IdPegawai == pejabat.IdPegawai {
					data.JabatanStruktural = append(data.JabatanStruktural, pejabat)
					IsNotStruktural = false
				}
			}

			if IsNotStruktural {
				// data.JabatanStruktural = append(data.JabatanStruktural, organisaiPrivate.PejabatStrukturalPrivate{})
				data.JabatanStruktural = make([]organisaiPrivate.PejabatStrukturalPrivate, 0)
			}

			pegawaiAndFungsionalAndStruktural = append(pegawaiAndFungsionalAndStruktural, data)
		}

		// get data kontrak
		stmt3, err := a.DB.Prepare(`SELECT p.id, COALESCE(pf.nomor_sk,'') no_surat,
		COALESCE(pf.tmt_sk,'') tanggal_mulai,
		COALESCE(pf.tgl_sk,'') tanggal_surat,
		COALESCE(pf.tmt_awal_kontrak,'') awal_kontrak,
		COALESCE(pf.tmt_akhir_kontrak,'') akhir_kontrak
		FROM pegawai p
		JOIN pegawai_fungsional pf ON p.id = pf.id_pegawai`)

		var kontrakPegawai []model.PegawaiKontrakPrivate
		rows3, err := stmt3.Query()
		if err != nil {
			fmt.Println(err)
			return c.JSON(500, nil)
		}
		defer rows3.Close()
		for rows3.Next() {
			var pk model.PegawaiKontrakPrivate
			if err := rows3.Scan(&pk.IdPegawai, &pk.NoSurat, &pk.TglMulai, &pk.TglSurat, &pk.AwalKontrak, &pk.AkhirKontrak); err != nil {
				fmt.Println(err)
				return c.JSON(500, nil)
			}
			kontrakPegawai = append(kontrakPegawai, pk)
		}
		if err := rows3.Err(); err != nil {
			fmt.Println(err)
			return c.JSON(500, nil)
		}

		var pegawaiJabfungJabstrukAndKontrak []model.PegawaiPrivate
		IsNotKontrak := true
		for _, data := range pegawaiAndFungsionalAndStruktural {
			for _, kontrak := range kontrakPegawai {
				if data.IdPegawai == kontrak.IdPegawai {
					// fmt.Println(kontrak)
					// data.PegawaiKontrakPrivate = append(data.PegawaiKontrakPrivate, kontrak)
					// data.PegawaiKontrakPrivate = kontrak
					data.PegawaiKontrakPrivate = model.PegawaiKontrakPrivate{NoSurat: kontrak.NoSurat, TglMulai: kontrak.TglMulai, TglSurat: kontrak.TglSurat, AwalKontrak: kontrak.AwalKontrak, AkhirKontrak: kontrak.AkhirKontrak}
					IsNotKontrak = false
					// fmt.Println("cek")
				}
			}
			if IsNotKontrak {
				// data.PegawaiKontrakPrivate = make([]model.PegawaiKontrakPrivate, 0)
				// data.PegawaiKontrakPrivate =
				// data.PegawaiKontrakPrivate = append(data.PegawaiKontrakPrivate, model.PegawaiKontrakPrivate{})
			}

			pegawaiJabfungJabstrukAndKontrak = append(pegawaiJabfungJabstrukAndKontrak, data)
		}

		// res.Data = pegawaiAndFungsionalAndStruktural
		// res.Data = pegawaiJabfungJabstrukAndKontrak

		tanggunganResponse := GetDataTanggungan(public)

		var pegawaiJabfungJabstrukAndKontrakAndTangungan []model.PegawaiPrivate
		for _, data := range pegawaiJabfungJabstrukAndKontrak {
			for _, tanggungan := range tanggunganResponse.Data {
				// if data.IdPersonal == tanggungan.IdPersonal {
				if strconv.FormatInt(int64(data.IdPegawai), 10) == tanggungan.IdPersonal {
					data.IdStatusPernikahanPtkp = tanggungan.IdStatusPernikahanPtkp
					data.KdStatusPernikahanPtkp = tanggungan.KdStatusPernikahanPtkp
					data.StatusPernikahanPtkp = tanggungan.StatusPernikahanPtkp
					data.JumlahTanggungan = tanggungan.JumlahTanggungan
					data.JumlahTanggunganPtkp = tanggungan.JumlahTanggunganPtkp
				}
			}
			pegawaiJabfungJabstrukAndKontrakAndTangungan = append(pegawaiJabfungJabstrukAndKontrakAndTangungan, data)
		}

		res.Data = pegawaiJabfungJabstrukAndKontrakAndTangungan
		return c.JSON(http.StatusOK, res)
	}
	return echo.HandlerFunc(h)
}

func GetDataTanggungan(public bool) *model.TanggunganResponseBody {
	// fmt.Println(env)
	var baseURL string
	baseURL = os.Getenv("URL_HCM_TANGGUNGAN")
	if public {
		// baseURL = "http://localhost:81/public/api/v1/tanggungan-private"
		baseURL = "http://svc-dependents-go.hcm-dev.svc.cluster.local/public/api/v1/tanggungan-private"
	}
	var client = &http.Client{}
	request, err := http.NewRequest(http.MethodGet, baseURL, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer response.Body.Close()
	b, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	data := &model.TanggunganResponseBody{}
	err = json.Unmarshal(b, data)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return data
}
