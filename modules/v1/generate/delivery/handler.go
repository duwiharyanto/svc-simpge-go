package delivery

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/generate/delivery/api"
	"svc-insani-go/modules/v1/generate/usecase"
	"svc-insani-go/modules/v1/master-kelompok-pegawai/repo"
	unitKerjaRepo "svc-insani-go/modules/v1/master-unit-kerja/repo"

	"github.com/labstack/echo/v4"
)

type GenerateHandler interface {
	GenerateNik(a *app.App) echo.HandlerFunc
}

type GenerateHandlerImpl struct {
	Request api.GenerateRequest
	Usecase usecase.GenerateUsecase
}

func (h *GenerateHandlerImpl) GenerateNik(a *app.App) echo.HandlerFunc {
	f := func(ctx echo.Context) error {
		payload, err := h.Request.GetGenerateNikReadRequest(ctx)
		if err != nil {
			if errors.Is(err, echo.ErrBadRequest) {
				return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "masukan yang diberikan tidak valid"})
			}
			fmt.Printf("[ERROR] error get generate nik, %s\n", err.Error())
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "terjadi kesalahan pada server"})
		}

		err = validateGenerateNikReadRequest(a, ctx.Request().Context(), payload)
		if err != nil {
			if errors.Is(err, echo.ErrBadRequest) {
				return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "masukan yang diberikan tidak valid"})
			}
			fmt.Printf("[ERROR] error validate generate nik, %s\n", err.Error())
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "terjadi kesalahan pada server"})
		}

		nik, err := h.Usecase.GenerateNik(a, ctx.Request().Context(), payload)
		if err != nil {
			if errors.Is(err, echo.ErrBadRequest) {
				return ctx.JSON(http.StatusBadRequest, map[string]string{"message": "masukan yang diberikan tidak valid"})
			}

			fmt.Printf("[ERROR] error generate nik, %s\n", err.Error())
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "terjadi kesalahan pada server"})
		}

		return ctx.JSON(http.StatusOK, map[string]string{
			"message": "berhasil mendapatkan nik",
			"data":    nik,
		})
	}
	return echo.HandlerFunc(f)
}

func validateGenerateNikReadRequest(a *app.App, ctx context.Context, payload *api.GenerateNikReadRequest) error {
	if payload.UuidKelompokPegawai == "" {
		return fmt.Errorf("uuid kelompok pegawai wajib diisi")
	}
	kelompokPegawai, err := repo.GetKelompokPegawaiByUUID(a, ctx, payload.UuidKelompokPegawai)
	if err != nil {
		return fmt.Errorf("error validate generate nik read request %s", err.Error())
	}
	if kelompokPegawai == nil {
		return echo.ErrBadRequest
	}

	if payload.UuidUnitPegawai == "" {
		return fmt.Errorf("uuid unit kerja wajib diisi")
	}
	unitKerja, err := unitKerjaRepo.GetUnitKerjaByUUID(a, payload.UuidUnitPegawai)
	if err != nil {
		return fmt.Errorf("error validate generate nik read request %s", err.Error())
	}
	if unitKerja == nil {
		return echo.ErrBadRequest
	}

	return nil
}

func NewGenerateHandler(request api.GenerateRequest, usecase usecase.GenerateUsecase) GenerateHandler {
	return &GenerateHandlerImpl{
		Request: request,
		Usecase: usecase,
	}
}
