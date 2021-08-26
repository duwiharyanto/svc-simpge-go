package handler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v3/model"
	"svc-insani-go/modules/v3/repo"

	"github.com/labstack/echo/v4"
)

const (
	msgErrorServer = "Layanan sedang bermasalah"
)

func HandleCreatePegawai(a *app.App, ctx context.Context) echo.HandlerFunc {
	h := func(c echo.Context) error {
		req := &PegawaiCreateRequest{}
		err := c.Bind(req)
		if err != nil {
			fmt.Printf("[WARNING] binding pegawai create request: %s\n", err.Error())
		}
		uuidPersonal := c.Param("uuidPersonal")
		if req.UuidPersonal == "" {
			req.UuidPersonal = uuidPersonal
		}
		if err := req.Validate(); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}
		req.UserUpdate = c.Request().Header.Get("X-Member")
		ctx := c.Request().Context()
		pegawai, err := prepareCreatePegawai(a, ctx, req)
		if errors.Unwrap(err) != nil {
			fmt.Printf("[ERROR] prepare create pegawai: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": msgErrorServer})
		}
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}
		err = repo.CreatePegawai(a, ctx, pegawai)
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Tidak ada data yang berhasil disimpan"})
		}
		if err != nil {
			fmt.Printf("[ERROR] create pegawai: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": msgErrorServer})
		}
		return c.JSON(http.StatusOK, map[string]string{"message": ""})
	}
	return h
}

const (
	statusActiveEmployee = "AKT"
)

func prepareCreatePegawai(a *app.App, ctx context.Context, req *PegawaiCreateRequest) (*model.Pegawai, error) {
	statusPegawaiAktif, err := repo.GetStatusPegawaiAktifByCode(a, ctx, statusActiveEmployee)
	if err != nil {
		return nil, fmt.Errorf("error from repo status pegawai aktif by uuid: %w", err)
	}
	if statusPegawaiAktif == nil {
		return nil, fmt.Errorf("%w", fmt.Errorf("status pegawai aktif tidak ditemukan"))
	}
	var pegawai *model.Pegawai
	pegawai.StatusPegawaiAktif = statusPegawaiAktif
	pegawai.UserInput = pegawai.UserUpdate

	personal, err := repo.GetPersonalByUuid(a, ctx, req.UuidPersonal)
	if err != nil {
		return nil, fmt.Errorf("%w", fmt.Errorf("error get personal by uuid: %w", err))
	}
	if personal == nil {
		return nil, fmt.Errorf("personal tidak ditemukan")
	}
	// pegawai, err =
	// if personal.Pegawai.PegawaiFungsional.StatusPegawaiAktif.IsActive() {
	// 	return nil, fmt.Errorf("tidak dapat menambah data dari pegawai yang sudah aktif")
	// }

	pegawai.IdPersonalDataPribadi = personal.Id
	pegawai.Nama = personal.NamaLengkap
	pegawai.NikKtp = personal.NikKtp
	pegawai.TglLahir = personal.TglLahir
	pegawai.TempatLahir = personal.TempatLahir
	pegawai.IdAgama = personal.IdAgama
	pegawai.KdAgama = personal.KdAgama
	pegawai.JenisKelamin = personal.JenisKelamin
	pegawai.IdGolonganDarah = personal.IdGolonganDarah
	pegawai.KdGolonganDarah = personal.KdGolonganDarah
	pegawai.IdStatusPerkawinan = personal.IdStatusPernikahan
	pegawai.KdStatusPerkawinan = personal.KdStatusPerkawinan
	pegawai.GelarDepan = personal.GelarDepan
	pegawai.GelarBelakang = personal.GelarBelakang

	return pegawai, nil
}

func HandleUpdatePegawai(a *app.App, ctx context.Context) echo.HandlerFunc {
	h := func(c echo.Context) error {
		req := &PegawaiCreateRequest{}
		err := c.Bind(req)
		if err != nil {
			fmt.Printf("[WARNING] binding pegawai create request: %s\n", err.Error())
		}
		uuidPersonal := c.Param("uuidPersonal")
		if req.UuidPersonal == "" {
			req.UuidPersonal = uuidPersonal
		}
		if err := req.Validate(); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}
		req.UserUpdate = c.Request().Header.Get("X-Member")
		ctx := c.Request().Context()
		pegawai, err := prepareCreatePegawai(a, ctx, req)
		if errors.Unwrap(err) != nil {
			fmt.Printf("[ERROR] prepare create pegawai: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": msgErrorServer})
		}
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}
		err = repo.CreatePegawai(a, ctx, pegawai)
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Tidak ada data yang berhasil disimpan"})
		}
		if err != nil {
			fmt.Printf("[ERROR] create pegawai: %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": msgErrorServer})
		}
		return c.JSON(http.StatusOK, map[string]string{"message": ""})
	}
	return h
}
