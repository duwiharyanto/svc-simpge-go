package repo

import (
	"context"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-bidang-sub-bidang/model"
)

func GetBidang(a *app.App, ctx context.Context) ([]model.Bidang, error) {
	sqlQuery := getBidangQuery()
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get bidang, %s", err.Error())
	}
	defer rows.Close()

	Bidang := []model.Bidang{}
	for rows.Next() {
		var b model.Bidang
		err := rows.Scan(&b.KdBidang, &b.Bidang, &b.UUID)
		if err != nil {
			return nil, fmt.Errorf("error scan bidang row, %s", err.Error())
		}
		Bidang = append(Bidang, b)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error bidang rows, %s", err.Error())
	}

	return Bidang, nil
}

func GetBidangByUUID(a *app.App, ctx context.Context, uuid string) (*model.Bidang, error) {
	var bidang model.Bidang
	tx := a.GormDB.WithContext(ctx)
	res := tx.First(&bidang, "uuid = ?", uuid)
	if res.Error != nil {
		return nil, fmt.Errorf("error querying bidang by uuid %s", res.Error)
	}
	return &bidang, nil
}

func GetSubBidang(a *app.App, ctx context.Context) ([]model.SubBidang, error) {
	sqlQuery := getSubBidangQuery()
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get bidang, %s", err.Error())
	}
	defer rows.Close()

	SubBidang := []model.SubBidang{}
	for rows.Next() {
		var sb model.SubBidang
		err := rows.Scan(&sb.KdBidang, &sb.Bidang, &sb.KdSubBidang, &sb.SubBidang, &sb.UUID)
		if err != nil {
			return nil, fmt.Errorf("error scan bidang row, %s", err.Error())
		}
		SubBidang = append(SubBidang, sb)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error bidang rows, %s", err.Error())
	}

	return SubBidang, nil
}

func GetSubBidangByUUID(a *app.App, ctx context.Context, uuid string) (*model.SubBidang, error) {
	var subBidang model.SubBidang
	tx := a.GormDB.WithContext(ctx)
	res := tx.First(&subBidang, "uuid = ?", uuid)
	if res.Error != nil {
		return nil, fmt.Errorf("error querying sub bidang by uuid %s", res.Error)
	}
	return &subBidang, nil
}
