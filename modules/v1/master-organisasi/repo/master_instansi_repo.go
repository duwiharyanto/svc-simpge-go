package repo

import (
	"context"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-organisasi/model"
)

func SearchInstansi(a *app.App, ctx context.Context) ([]model.Instansi, error) {

	var instansi []model.Instansi
	tx := a.GormDB.WithContext(ctx)

	res := tx.Find(&instansi)
	if res.Error != nil {
		return nil, res.Error
	}

	return instansi, nil
}
