package repo

import (
	"context"
	"svc-insani-go/app"
	"svc-insani-go/modules/v2/sk/model"
)

func GetAllJenisSk(a app.App, ctx context.Context) []model.JenisSk {
	var jjs []model.JenisSk
	a.GormDB.
		WithContext(ctx).
		Where(&model.JenisSk{FlagAktif: 1}).
		Find(&jjs)
	return jjs
}
