package repo

import (
	"context"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-jenjang-pendidikan/model"
)

func GetJenjangPendidikan(a *app.App, ctx context.Context) ([]model.JenjangPendidikan, error) {
	sqlQuery := getJenjangPendidikanQuery()
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get jenjang pendidikan, %w", err)
	}
	defer rows.Close()

	pp := []model.JenjangPendidikan{}
	for rows.Next() {
		var p model.JenjangPendidikan
		err := rows.Scan(
			&p.ID,
			&p.KdJenjang,
			&p.Jenjang,
			&p.NamaJenjang,
			&p.UUID,
		)
		if err != nil {
			return nil, fmt.Errorf("error scan jenjang pendidikan row, %s", err.Error())
		}
		pp = append(pp, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error jenjang pendidikan rows, %s", err.Error())
	}

	return pp, nil
}

func GetJenjangPendidikanByUUID(a *app.App, ctx context.Context, uuid string) (*model.JenjangPendidikan, error) {
	var lokasiKerja model.JenjangPendidikan

	tx := a.GormDB.WithContext(ctx)
	res := tx.First(&lokasiKerja, "uuid = ?", uuid)
	if res.Error != nil {
		return nil, fmt.Errorf("error querying jenjang pendidikan by uuid %s", res.Error)
	}
	return &lokasiKerja, nil
}
