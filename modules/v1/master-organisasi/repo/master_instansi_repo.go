package repo

import (
	"context"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-organisasi/model"
)

func SearchInstansi(a app.App, ctx context.Context, nama string) (*model.Instansi, error) {

	var instansi model.Instansi
	tx := a.GormDB.WithContext(ctx)

	if nama != "" {
		res := tx.Where("nama_instansi LIKE ? OR kd_instansi LIKE ?", "%"+nama+"%", "%"+nama+"%").
			First(&instansi)
		if res.Error != nil {
			return nil, res.Error
		}
	}

	return &instansi, nil
}
