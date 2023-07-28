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

// ini digunakan kalo uiitanggungan sudah digunakan
func HandleGetPegawaiPrivateSatu(a *app.App, public bool) echo.HandlerFunc {
	h := func(c echo.Context) error {
		env := os.Getenv("ENV")
		fmt.Println(env)
		if public {
			if env == "production" {
				// return c.JSON(404, "layanan tidak ditemukan")
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
			log.Printf("[ERROR] repo get all pegawai private: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		res.Data = pp
		// get data jabatan struktural
		stmt, err := a.DB.Prepare(`SELECT COALESCE(p.id,0), 
		COALESCE(u.id_jenis_unit,0),
		COALESCE(po.id_jenis_jabatan,0),
		COALESCE(po.id_unit,0),
		CASE WHEN DATE_FORMAT(sk.tst_surat_keputusan, '%Y-%m') >= DATE_FORMAT(now(), '%Y-%m') THEN 1
		ELSE 0 END AS flag_aktif
		FROM pegawai p 
		JOIN pejabat_organisasi po ON po.id_pegawai = p.id 
		JOIN unit u ON u.id = po.id_unit
		JOIN surat_keputusan sk ON po.id_surat_keputusan = sk.id 
		WHERE po.flag_aktif =1`)
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
			if err := rows.Scan(&ps.IdPegawai, &ps.IdJenisUnit, &ps.IdJenisJabatan, &ps.IdUnit, &ps.FlagAktif); err != nil {
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
		for _, data := range res.Data {
			masaKerjaBawaanTahunInt, _ := strconv.Atoi(data.MasaKerjaBawaanTahun)
			masaKerjaBawaanBulanInt, _ := strconv.Atoi(data.MasaKerjaBawaanBulan)

			tmtSkPertamaTime, err := time.Parse("2006-01-02", data.TmtSkPertama)
			var tmtSkPertamaDuration time.Duration
			if err == nil {
				tmtSkPertamaDuration = time.Now().Sub(tmtSkPertamaTime)
			}
			tmtSkPertamaDurationDays := tmtSkPertamaDuration.Hours() / 24
			// tmtSkPertamaDurationDays := tmtSkPertamaDuration.Hours() / 24
			tmtSkPertamaDurationRealMonths := int(tmtSkPertamaDurationDays / 365 * 12)

			// masa kerja gaji total
			masaKerjaTotalRealBulan := ((masaKerjaBawaanTahunInt * 12) + masaKerjaBawaanBulanInt) + tmtSkPertamaDurationRealMonths
			data.MasaKerjaTotalTahun = fmt.Sprintf("%d", masaKerjaTotalRealBulan/12)
			data.MasaKerjaTotalBulan = fmt.Sprintf("%d", masaKerjaTotalRealBulan%12)

			// masa kerja kepegawaian total
			masaKerjaKepegawaianTahunInt, _ := strconv.Atoi(data.MasaKerjaAwalKepegawaianTahun)
			masaKerjaKepegawaianBulanInt, _ := strconv.Atoi(data.MasaKerjaAwalKepegawaianBulan)
			masaKerjaTotalKepegawaianRealBulan := ((masaKerjaKepegawaianTahunInt * 12) + masaKerjaKepegawaianBulanInt) + tmtSkPertamaDurationRealMonths
			data.MasaKerjaAwalKepegawaianTahun = fmt.Sprintf("%d", masaKerjaTotalKepegawaianRealBulan/12)
			data.MasaKerjaAwalKepegawaianBulan = fmt.Sprintf("%d", masaKerjaTotalKepegawaianRealBulan%12)

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
		stmt3, err := a.DB.Prepare(`SELECT p.id, COALESCE(pf.nomor_surat_kontrak,'') no_surat,
		COALESCE(pf.tmt_surat_kontrak,'') tanggal_mulai,
		COALESCE(pf.tgl_surat_kontrak,'') tanggal_surat,
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
		stmt4, err := a.DB.Prepare(`SELECT
		COALESCE(pp.id_personal_data_pribadi,0),
		COALESCE(pp.kd_jenjang,''),
		COALESCE(pp.flag_ijazah_diakui,''),
		COALESCE(pp.nama_institusi,'')
		FROM pegawai_pendidikan pp
		LEFT JOIN pegawai p ON pp.id_personal_data_pribadi = p.id_personal_data_pribadi
		WHERE pp.flag_aktif = 1`)

		var pendidikanPegawai []model.PegawaiPendidikanPrivate
		rows4, err := stmt4.Query()
		if err != nil {
			fmt.Println(err)
			return c.JSON(500, nil)
		}
		defer rows4.Close()
		for rows4.Next() {
			var pp model.PegawaiPendidikanPrivate
			if err := rows4.Scan(&pp.IdPersonal, &pp.KdJenjang, &pp.FlagIjazahDiakui, &pp.NamaInstitusi); err != nil {
				fmt.Println(err)
				return c.JSON(500, nil)
			}
			pendidikanPegawai = append(pendidikanPegawai, pp)
		}

		var pegawaiJabfungJabstrukAndKontrakAndPendidikan []model.PegawaiPrivate
		IsPendidikanNull := true
		for _, data := range pegawaiJabfungJabstrukAndKontrak {
			for _, pendidikan := range pendidikanPegawai {
				if data.IdPersonal == pendidikan.IdPersonal {
					data.Pendidikan = append(data.Pendidikan, pendidikan)
					IsPendidikanNull = false
				}
			}

			if IsPendidikanNull {
				data.Pendidikan = make([]model.PegawaiPendidikanPrivate, 0)
			}

			pegawaiJabfungJabstrukAndKontrakAndPendidikan = append(pegawaiJabfungJabstrukAndKontrakAndPendidikan, data)
		}

		tanggunganResponse := GetDataTanggunganSatu(public)

		var pegawaiJabfungJabstrukAndKontrakAndPendidikanAndTangungan []model.PegawaiPrivate
		for _, data := range pegawaiJabfungJabstrukAndKontrakAndPendidikan {
			for _, tanggungan := range tanggunganResponse.Data {
				// if data.IdPersonal == tanggungan.IdPersonal {
				if strconv.FormatInt(int64(data.IdPersonal), 10) == tanggungan.IdPersonal {
					data.IdStatusPernikahanPtkp = tanggungan.IdStatusPernikahanPtkp
					data.KdStatusPernikahanPtkp = tanggungan.KdStatusPernikahanPtkp
					data.StatusPernikahanPtkp = tanggungan.StatusPernikahanPtkp
					data.JumlahKeluargaDitanggung = tanggungan.JumlahKeluargaDitanggung
					data.JumlahAnakDitanggung = tanggungan.JumlahAnakDitanggung
					// data.JumlahKeluargaDitanggungPtkp = tanggungan.JumlahKeluargaDitanggungPtkp
					// data.JumlahAnakDitanggungPtkp = tanggungan.JumlahAnakDitanggungPtkp
					// data.DetailTanggunganKeluarga = tanggungan.DetailTanggunganKeluarga
					// data.DetailTanggunganPtkp = tanggungan.DetailTanggunganPtkp
				}
			}
			pegawaiJabfungJabstrukAndKontrakAndPendidikanAndTangungan = append(pegawaiJabfungJabstrukAndKontrakAndPendidikanAndTangungan, data)
		}

		res.Data = pegawaiJabfungJabstrukAndKontrakAndPendidikanAndTangungan

		return c.JSON(http.StatusOK, res)
	}
	return echo.HandlerFunc(h)
}

// ini digunakan kalo uiitanggungan sudah digunakan
func GetDataTanggunganSatu(public bool) *model.TanggunganResponseBody {
	// fmt.Println(env)
	var baseURL string
	baseURL = os.Getenv("URL_HCM_TANGGUNGAN")
	fmt.Println("baseUrl = " + baseURL)
	// if public {
	// baseURL = "http://localhost:81/public/api/v1/tanggungan-private"
	// baseURL = "http://svc-dependents-go.hcm-dev.svc.cluster.local/public/api/v1/tanggungan-private"
	// }

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

func HandleGetPegawaiPrivate(a *app.App, public bool) echo.HandlerFunc {
	h := func(c echo.Context) error {
		env := os.Getenv("ENV")
		fmt.Println(env)
		if public {
			if env == "production" {
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
			log.Printf("[ERROR] repo get all pegawai private: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}
		res.Data = pp
		// get data jabatan struktural
		stmt, err := a.DB.Prepare(`SELECT COALESCE(p.id,0), 
		COALESCE(u.id_jenis_unit,0),
		COALESCE(po.id_jenis_jabatan,0),
		COALESCE(po.id_unit,0),
		CASE WHEN DATE_FORMAT(sk.tst_surat_keputusan, '%Y-%m') >= DATE_FORMAT(now(), '%Y-%m') THEN 1
		ELSE 0 END AS flag_aktif
		FROM pegawai p 
		JOIN pejabat_organisasi po ON po.id_pegawai = p.id 
		JOIN unit u ON u.id = po.id_unit
		JOIN surat_keputusan sk ON po.id_surat_keputusan = sk.id 
		WHERE po.flag_aktif =1`)
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
			if err := rows.Scan(&ps.IdPegawai, &ps.IdJenisUnit, &ps.IdJenisJabatan, &ps.IdUnit, &ps.FlagAktif); err != nil {
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
		for _, data := range res.Data {
			masaKerjaBawaanTahunInt, _ := strconv.Atoi(data.MasaKerjaBawaanTahun)
			masaKerjaBawaanBulanInt, _ := strconv.Atoi(data.MasaKerjaBawaanBulan)

			tmtSkPertamaTime, err := time.Parse("2006-01-02", data.TmtSkPertama)
			var tmtSkPertamaDuration time.Duration
			if err == nil {
				tmtSkPertamaDuration = time.Now().Sub(tmtSkPertamaTime)
			}
			tmtSkPertamaDurationDays := tmtSkPertamaDuration.Hours() / 24
			// tmtSkPertamaDurationDays := tmtSkPertamaDuration.Hours() / 24
			tmtSkPertamaDurationRealMonths := int(tmtSkPertamaDurationDays / 365 * 12)

			// masa kerja gaji total
			masaKerjaTotalRealBulan := ((masaKerjaBawaanTahunInt * 12) + masaKerjaBawaanBulanInt) + tmtSkPertamaDurationRealMonths
			data.MasaKerjaTotalTahun = fmt.Sprintf("%d", masaKerjaTotalRealBulan/12)
			data.MasaKerjaTotalBulan = fmt.Sprintf("%d", masaKerjaTotalRealBulan%12)

			// masa kerja kepegawaian total
			masaKerjaKepegawaianTahunInt, _ := strconv.Atoi(data.MasaKerjaAwalKepegawaianTahun)
			masaKerjaKepegawaianBulanInt, _ := strconv.Atoi(data.MasaKerjaAwalKepegawaianBulan)
			masaKerjaTotalKepegawaianRealBulan := ((masaKerjaKepegawaianTahunInt * 12) + masaKerjaKepegawaianBulanInt) + tmtSkPertamaDurationRealMonths
			data.MasaKerjaAwalKepegawaianTahun = fmt.Sprintf("%d", masaKerjaTotalKepegawaianRealBulan/12)
			data.MasaKerjaAwalKepegawaianBulan = fmt.Sprintf("%d", masaKerjaTotalKepegawaianRealBulan%12)

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
		stmt3, err := a.DB.Prepare(`SELECT p.id, COALESCE(pf.nomor_surat_kontrak,'') no_surat,
		COALESCE(pf.tmt_surat_kontrak,'') tanggal_mulai,
		COALESCE(pf.tgl_surat_kontrak,'') tanggal_surat,
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
		stmt4, err := a.DB.Prepare(`SELECT
		COALESCE(pp.id_personal_data_pribadi,0),
		COALESCE(pp.kd_jenjang,''),
		COALESCE(pp.flag_ijazah_diakui,''),
		COALESCE(pp.nama_institusi,'')
		FROM pegawai_pendidikan pp
		LEFT JOIN pegawai p ON pp.id_personal_data_pribadi = p.id_personal_data_pribadi
		WHERE pp.flag_aktif = 1`)

		var pendidikanPegawai []model.PegawaiPendidikanPrivate
		rows4, err := stmt4.Query()
		if err != nil {
			fmt.Println(err)
			return c.JSON(500, nil)
		}
		defer rows4.Close()
		for rows4.Next() {
			var pp model.PegawaiPendidikanPrivate
			if err := rows4.Scan(&pp.IdPersonal, &pp.KdJenjang, &pp.FlagIjazahDiakui, &pp.NamaInstitusi); err != nil {
				fmt.Println(err)
				return c.JSON(500, nil)
			}
			pendidikanPegawai = append(pendidikanPegawai, pp)
		}

		var pegawaiJabfungJabstrukAndKontrakAndPendidikan []model.PegawaiPrivate
		IsPendidikanNull := true
		for _, data := range pegawaiJabfungJabstrukAndKontrak {
			for _, pendidikan := range pendidikanPegawai {
				if data.IdPersonal == pendidikan.IdPersonal {
					data.Pendidikan = append(data.Pendidikan, pendidikan)
					IsPendidikanNull = false
				}
			}

			if IsPendidikanNull {
				data.Pendidikan = make([]model.PegawaiPendidikanPrivate, 0)
			}

			pegawaiJabfungJabstrukAndKontrakAndPendidikan = append(pegawaiJabfungJabstrukAndKontrakAndPendidikan, data)
		}

		tanggunganResponse, err := GetDataTanggungan(public)
		if err != nil {
			// fmt.Println(err)
			// return nil
			return c.JSON(http.StatusInternalServerError, nil)
		}

		var pegawaiJabfungJabstrukAndKontrakAndPendidikanAndTangungan []model.PegawaiPrivate
		for _, data := range pegawaiJabfungJabstrukAndKontrakAndPendidikan {
			for _, tanggungan := range tanggunganResponse.Data {
				// if data.IdPersonal == tanggungan.IdPersonal {
				if strconv.FormatInt(int64(data.IdPersonal), 10) == tanggungan.IdPersonal {
					data.IdStatusPernikahanPtkp = tanggungan.IdStatusPernikahanPtkp
					data.KdStatusPernikahanPtkp = tanggungan.KdStatusPernikahanPtkp
					data.StatusPernikahanPtkp = tanggungan.StatusPernikahanPtkp
					data.JumlahKeluargaDitanggung = tanggungan.JumlahKeluargaDitanggung
					data.JumlahAnakDitanggung = tanggungan.JumlahAnakDitanggung
					data.JumlahKeluargaDitanggungPtkp = tanggungan.JumlahKeluargaDitanggungPtkp
					data.JumlahAnakDitanggungPtkp = tanggungan.JumlahAnakDitanggungPtkp
					// 	data.DetailTanggunganKeluarga = tanggungan.DetailTanggunganKeluarga
					// 	data.DetailTanggunganPtkp = tanggungan.DetailTanggunganPtkp
				}
			}
			pegawaiJabfungJabstrukAndKontrakAndPendidikanAndTangungan = append(pegawaiJabfungJabstrukAndKontrakAndPendidikanAndTangungan, data)
		}

		res.Data = pegawaiJabfungJabstrukAndKontrakAndPendidikanAndTangungan

		return c.JSON(http.StatusOK, res)
	}
	return echo.HandlerFunc(h)
}

//
func GetDataTanggungan(public bool) (*model.TanggunganResponseBody, error) {
	// fmt.Println(env)
	var baseURL string
	baseURL = os.Getenv("URL_ACTIVATION_PERSONAL")
	endpoint := "/public/api/v1/tanggungan-private"
	destinationURL := baseURL + endpoint

	// fmt.Println(destinationURL)
	fmt.Println("baseUrl = " + destinationURL)

	var client = &http.Client{}
	request, err := http.NewRequest("GET", destinationURL, nil)
	if err != nil {
		// fmt.Println(err)
		// return nil
		fmt.Println(err)
		fmt.Printf("[ERROR] %s - at NewRequest\n", err)
		return nil, err
	}
	fmt.Println(request)
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("[ERROR] error get data tanggungan", err)
		return nil, err
	}
	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("[ERROR] %s - at readAll\n", err)
		return nil, err
	}

	data := &model.TanggunganResponseBody{}
	err = json.Unmarshal(b, data)
	fmt.Println("data: ", data)
	if err != nil {
		fmt.Printf("[ERROR] %s - at unmarshal\n", err)
		// return nil
		return nil, err
	}

	return data, nil
}
