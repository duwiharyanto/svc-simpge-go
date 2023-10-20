package delivery

import (
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/generate/delivery/api"
	"svc-insani-go/modules/v1/generate/usecase"

	"github.com/labstack/echo/v4"
)

type Route struct{}

func (r *Route) Init(e *echo.Echo, a *app.App) {
	request := api.NewGenerateNikRequest()
	usecase := usecase.NewGenerateUsecase()
	handler := NewGenerateHandler(request, usecase)

	insaniGroupingPath := e.Group("/public/api/v1")
	insaniPrivateGroupingPath := e.Group("/private/api/v1")
	_ = insaniPrivateGroupingPath

	insaniGroupingPath.GET("/generate-nik", handler.GenerateNik(a))

}
