package repo

import (
	"context"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-pendidikan/model"
)

func GetGelarDepan(a *app.App, ctx context.Context) ([]model.GelarDepan, error) {
	sqlQuery := getGelarDepanQuery()
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get gelar depan: %w", err)
	}
	defer rows.Close()

	list := []model.GelarDepan{}
	for rows.Next() {
		var row model.GelarDepan
		err := rows.Scan(&row.Gelar)
		if err != nil {
			return nil, fmt.Errorf("error scan gelar depan row: %s", err.Error())
		}
		list = append(list, row)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error gelar depan rows: %s", err.Error())
	}

	return list, nil
}

func GetGelarBelakang(a *app.App, ctx context.Context) ([]model.GelarBelakang, error) {
	sqlQuery := getGelarBelakangQuery()
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get gelar belakang: %w", err)
	}
	defer rows.Close()

	list := []model.GelarBelakang{}
	for rows.Next() {
		var row model.GelarBelakang
		err := rows.Scan(&row.Gelar)
		if err != nil {
			return nil, fmt.Errorf("error scan gelar belakang row: %s", err.Error())
		}
		list = append(list, row)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error gelar belakang rows: %s", err.Error())
	}

	return list, nil
}

func GetJenjangPendidikan(a *app.App, ctx context.Context) (model.JenjangPendidikanList, error) {
	sqlQuery := getJenjangPendidikanQuery()
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get jenjang pendidikan: %w", err)
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
			&p.KdPendidikanSimpeg,
			&p.NamaPendidikanSimpeg,
			&p.UUID,
		)
		if err != nil {
			return nil, fmt.Errorf("error scan jenjang pendidikan row: %s", err.Error())
		}
		pp = append(pp, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error jenjang pendidikan rows: %s", err.Error())
	}

	return pp, nil
}

func GetJenjangPendidikanDetail(a *app.App, ctx context.Context) ([]model.JenjangPendidikanDetail, error) {
	sqlQuery := getJenjangPendidikanDetailQuery()
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get jenjang pendidikan detail: %w", err)
	}
	defer rows.Close()

	pp := []model.JenjangPendidikanDetail{}
	for rows.Next() {
		var p model.JenjangPendidikanDetail
		err := rows.Scan(
			&p.ID,
			&p.KdJenjangPendidikan,
			&p.KdDetail,
			&p.NamaDetail,
			&p.Keterangan,
			&p.UUID,
		)
		if err != nil {
			return nil, fmt.Errorf("error scan jenjang pendidikan detail row: %s", err.Error())
		}
		pp = append(pp, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error jenjang pendidikan detail rows: %s", err.Error())
	}

	return pp, nil
}

func GetJenjangPendidikanByUUID(a *app.App, ctx context.Context, uuid string) (*model.JenjangPendidikan, error) {
	var jenjangPendidikan model.JenjangPendidikan

	tx := a.GormDB.WithContext(ctx)
	res := tx.Where("flag_aktif = 1 AND uuid = ?", uuid).First(&jenjangPendidikan)
	if res.Error != nil {
		return nil, fmt.Errorf("error querying jenjang pendidikan by uuid %s", res.Error)
	}
	return &jenjangPendidikan, nil
}
