package repo

import (
	"context"
	"svc-insani-go/app"
	"svc-insani-go/modules/v2/sk/model"
)

func GetAllKelompokSkPengangkatan(a app.App, ctx context.Context) []model.KelompokSkPengangkatan {
	var ksk []model.KelompokSkPengangkatan
	// a.GormDB.WithContext(ctx).Find(&ksk)
	a.GormDB.
		WithContext(ctx).
		Where(&model.KelompokSkPengangkatan{FlagAktif: 1}).
		Find(&ksk)
	return ksk
}
