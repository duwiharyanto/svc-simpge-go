package api

import (
	"log"

	"github.com/labstack/echo/v4"
)

type GenerateNikReadRequest struct {
	UuidKelompokPegawai string `json:"uuid_kelompok_pegawai" form:"uuid_kelompok_pegawai"`
	UuidUnitPegawai     string `json:"uuid_unit_pegawai" form:"uuid_unit_pegawai"`
}

type GenerateRequest interface {
	GetGenerateNikReadRequest(ctx echo.Context) (*GenerateNikReadRequest, error)
}

type GenerateNikRequestImpl struct{}

func (request *GenerateNikRequestImpl) GetGenerateNikReadRequest(ctx echo.Context) (*GenerateNikReadRequest, error) {
	generateNikReadRequest := &GenerateNikReadRequest{}
	err := ctx.Bind(generateNikReadRequest)
	if err != nil {
		log.Fatal(err.Error())
	}

	return generateNikReadRequest, nil
}

func NewGenerateNikRequest() GenerateRequest {
	return &GenerateNikRequestImpl{}
}
