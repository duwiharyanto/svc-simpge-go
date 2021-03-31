package repo

import (
	"context"
	"svc-insani-go/app"
	"svc-insani-go/modules/v2/kepegawaian/model"
)

func GetAllPangkatGolonganRuang(a app.App, ctx context.Context) []model.PangkatGolonganRuang {
	var pp []model.PangkatGolonganRuang
	a.GormDB.
		WithContext(ctx).
		Where(&model.PangkatGolonganRuang{FlagAktif: 1}). // untuk sk pengangkatan tampilkan semua unit
		Find(&pp)
	return pp
}
