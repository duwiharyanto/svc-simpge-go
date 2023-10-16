package usecase

import (
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/generate/delivery/api"
	"svc-insani-go/modules/v1/master-kelompok-pegawai/repo"
	unitKerjaRepo "svc-insani-go/modules/v1/master-unit-kerja/repo"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/context"
)

type GenerateUsecase interface {
	GenerateNik(a *app.App, ctx context.Context, payload *api.GenerateNikReadRequest) (string, error)
}

type GenerateUsecaseImpl struct{}

func (u *GenerateUsecaseImpl) GenerateNik(a *app.App, ctx context.Context, payload *api.GenerateNikReadRequest) (string, error) {
	kelompokPegawai, err := repo.GetKelompokPegawaiByUUID(a, ctx, payload.UuidKelompokPegawai)
	if err != nil {
		return "", fmt.Errorf("error validate generate nik read request %s", err.Error())
	}
	if kelompokPegawai == nil {
		return "", echo.ErrBadRequest
	}

	unitKerja, err := unitKerjaRepo.GetUnitKerjaByUUID(a, payload.UuidUnitPegawai)
	if err != nil {
		return "", fmt.Errorf("error validate generate nik read request %s", err.Error())
	}
	if unitKerja == nil {
		return "", echo.ErrBadRequest
	}

	return "23" + unitKerja.KdUnitKerja + kelompokPegawai.KdKelompokPegawai, nil
}

func NewGenerateUsecase() GenerateUsecase {
	return &GenerateUsecaseImpl{}
}
