package repo

import (
	"context"
	"database/sql"
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
	sqlQuery := getBidangByUUIDQuery(uuid)
	var bidang model.Bidang
	err := a.DB.QueryRow(sqlQuery).Scan(&bidang.ID, &bidang.KdBidang, &bidang.Bidang, &bidang.UUID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error querying get bidang by uuid %s", err.Error())
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
	sqlQuery := getSubBidangByUUIDQuery(uuid)
	var subBidang model.SubBidang
	err := a.DB.QueryRow(sqlQuery).Scan(&subBidang.ID, &subBidang.KdSubBidang, &subBidang.SubBidang, &subBidang.UUID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error querying get sub bidang by uuid %s", err.Error())
	}
	return &subBidang, nil
}
