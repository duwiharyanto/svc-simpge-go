package repo

import (
	"context"
	"svc-insani-go/app"
	"svc-insani-go/modules/v2/sk/model"
)

func GetAllSkPegawai(a app.App, ctx context.Context) []model.SkPegawai {
	var ssp []model.SkPegawai
	a.GormDB.
		WithContext(ctx).
		Where(&model.SkPegawai{FlagAktif: 1}).
		Find(&ssp)
	return ssp
}
