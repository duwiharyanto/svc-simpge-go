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
		// Where(&model.PangkatGolonganRuang{FlagAktif: 1}). // untuk sk pengangkatan tampilkan semua flag
		Find(&pp)
	return pp
}

func GetAllJabatanFungsional(a app.App, ctx context.Context) ([]model.JabatanFungsional, error) {
	var jj []model.JabatanFungsional
	err := a.GormDB.
		WithContext(ctx).
		// Where(&model.JabatanFungsional{FlagAktif: 1}). // untuk sk pengangkatan tampilkan semua flag
		Find(&jj).
		Error

	if err != nil {
		return nil, err
	}

	return jj, nil
}
