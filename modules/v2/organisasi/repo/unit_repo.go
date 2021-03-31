package repo

import (
	"context"
	"svc-insani-go/app"
	"svc-insani-go/modules/v2/organisasi/model"
)

func GetAllUnit2(a app.App, ctx context.Context) []model.Unit2 {
	var uu []model.Unit2
	a.GormDB.
		WithContext(ctx).
		// Where(&model.Unit2{FlagAktif: 1}). // untuk sk pengangkatan tampilkan semua unit
		Find(&uu)
	return uu
}
