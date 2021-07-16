package repo

import (
	"database/sql"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/sk/model"
)

func GetAllJabatanPenetap(a *app.App) ([]model.JabatanPenetap, error) {
	// c.Param("kd_jenis_pegawai")
	sqlQuery := getJabatanPenetapQuery()
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get jabatan penetap pegawai, %s", err.Error())
	}
	defer rows.Close()

	jabatanFungsional := []model.JabatanPenetap{}
	for rows.Next() {
		var s model.JabatanPenetap
		err := rows.Scan(&s.KdJabatanPenetap, &s.JabatanPenetap, &s.UUID)
		if err != nil {
			return nil, fmt.Errorf("error scan jabatan penetap row, %s", err.Error())
		}
		jabatanFungsional = append(jabatanFungsional, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error jabatan penetap rows, %s", err.Error())
	}

	return jabatanFungsional, nil
}
func GetAllJabatanPenetapByUUID(a *app.App, uuid string) (*model.JabatanPenetap, error) {
	sqlQuery := getJabatanPenetapQueryByUUID(uuid)
	//fmt.Printf("[DEBUG] jabatan fungsional pegawai by uuid:\n%s\n", sqlQuery)
	var jabatanPenetap model.JabatanPenetap
	err := a.DB.QueryRow(sqlQuery).Scan(&jabatanPenetap.ID, &jabatanPenetap.KdJabatanPenetap, &jabatanPenetap.JabatanPenetap, &jabatanPenetap.UUID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error querying jabatan fungsional sk by uuid %s", err.Error())
	}
	return &jabatanPenetap, nil
}
