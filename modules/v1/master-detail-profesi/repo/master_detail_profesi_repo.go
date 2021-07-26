package repo

import (
	"context"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-detail-profesi/model"
)

func GetDetailProfesi(a *app.App, ctx context.Context) ([]model.DetailProfesi, error) {
	sqlQuery := getDetailProfesiQuery()
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get detail profesi, %w", err)
	}
	defer rows.Close()

	pp := []model.DetailProfesi{}
	for rows.Next() {
		var p model.DetailProfesi
		err := rows.Scan(
			&p.ID,
			&p.DetailProfesi,
			&p.UUID,
		)
		if err != nil {
			return nil, fmt.Errorf("error scan detail profesi row, %s", err.Error())
		}
		pp = append(pp, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error detail profesi rows, %s", err.Error())
	}

	return pp, nil
}

func GetDetailProfesiByUUID(a *app.App, ctx context.Context, uuid string) (*model.DetailProfesi, error) {
	var detailProfesi model.DetailProfesi

	tx := a.GormDB.WithContext(ctx)
	res := tx.First(&detailProfesi, "uuid = ?", uuid)
	if res.Error != nil {
		return nil, fmt.Errorf("error querying detail profesi by uuid %s", res.Error)
	}
	return &detailProfesi, nil
}
