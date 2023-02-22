package repo

import (
	"context"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-bidang-sub-bidang/model"
)

func GetBidangSubBidang(a *app.App, ctx context.Context) ([]model.BidangSubBidang, error) {
	sqlQuery := getBidangSubBidangQuery()
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get bidang sub bidang, %s", err.Error())
	}
	defer rows.Close()

	BidangSubBidang := []model.BidangSubBidang{}
	for rows.Next() {
		var b model.BidangSubBidang
		err := rows.Scan(&b.KdBidangSubBidang, &b.BidangSubBidang, &b.UUID)
		if err != nil {
			return nil, fmt.Errorf("error scan bidang sub bidang row, %s", err.Error())
		}
		BidangSubBidang = append(BidangSubBidang, b)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error bidang sub bidang rows, %s", err.Error())
	}

	return BidangSubBidang, nil
}

func GetBidangSubBidangByUUID(a *app.App, ctx context.Context, uuid string) (*model.BidangSubBidang, error) {
	var bidangSubBidang model.BidangSubBidang
	tx := a.GormDB.WithContext(ctx)
	res := tx.First(&bidangSubBidang, "uuid = ?", uuid)
	if res.Error != nil {
		return nil, fmt.Errorf("error querying bidang sub bidang by uuid %s", res.Error)
	}
	return &bidangSubBidang, nil
}
