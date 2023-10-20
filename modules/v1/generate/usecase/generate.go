package usecase

import (
	"fmt"
	"strconv"
	"svc-insani-go/app"
	"svc-insani-go/app/helper"
	"svc-insani-go/modules/v1/generate/delivery/api"
	"svc-insani-go/modules/v1/master-kelompok-pegawai/repo"
	unitKerjaRepo "svc-insani-go/modules/v1/master-unit-kerja/repo"
	pegawaiRepo "svc-insani-go/modules/v1/pegawai/repo"

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

	unitKerja, err := unitKerjaRepo.GetUnit2ByUUID(a, payload.UuidUnitPegawai)
	if err != nil {
		return "", fmt.Errorf("error validate generate nik read request %s", err.Error())
	}
	if unitKerja == nil {
		return "", echo.ErrBadRequest
	}

	now, err := helper.DateNow()
	if err != nil {
		return "", fmt.Errorf("error validate generate nik read request: %s; ", err.Error())
	}

	tempNik := now.Format("06") + unitKerja.KdUnit + kelompokPegawai.KdKelompokPegawai
	counter, err := generateCounter(a, tempNik)
	if err != nil {
		return "", fmt.Errorf("error validate generate nik read request: %s; ", err.Error())
	}

	return tempNik + counter, nil
}

func generateCounter(a *app.App, tempNik string) (string, error) {
	count, err := pegawaiRepo.CountNikPegawai(a, tempNik)
	if err != nil {
		return "", fmt.Errorf("error generate counter: %s; ", err.Error())
	}

	return doGenerateCounter(count), nil
}

func doGenerateCounter(count int) string {
	if count < 10 {
		return "0" + strconv.Itoa(count+1)
	}

	return strconv.Itoa(count)
}

func NewGenerateUsecase() GenerateUsecase {
	return &GenerateUsecaseImpl{}
}
