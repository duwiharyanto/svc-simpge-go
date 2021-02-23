package repo

import (
	"database/sql"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/master-jabatan-fungsional/model"
)

func GetAllJabatanFungsional(a app.App, kdJenisPegawai string) ([]model.JabatanFungsional, error) {
	sqlQuery := getJabatanFungsionalQuery(kdJenisPegawai)
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get jabatan fungsional pegawai, %s", err.Error())
	}
	defer rows.Close()

	jabatanFungsional := []model.JabatanFungsional{}
	for rows.Next() {
		var s model.JabatanFungsional
		err := rows.Scan(&s.KdJabatanFungsional, &s.JabatanFungsional, &s.UUID)
		if err != nil {
			return nil, fmt.Errorf("error scan jabatan fungsional row, %s", err.Error())
		}
		jabatanFungsional = append(jabatanFungsional, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error jabatan fungsional rows, %s", err.Error())
	}

	return jabatanFungsional, nil
}
func GetJabatanFungsionalByUUID(a app.App, uuid string) (*model.JabatanFungsional, error) {
	sqlQuery := getJabatanFungsionalByUUID(uuid)
	//fmt.Printf("[DEBUG] jabatan fungsional pegawai by uuid:\n%s\n", sqlQuery)
	var jabatanFungsional model.JabatanFungsional
	err := a.DB.QueryRow(sqlQuery).Scan(&jabatanFungsional.ID, &jabatanFungsional.KdJabatanFungsional, &jabatanFungsional.JabatanFungsional, &jabatanFungsional.UUID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error querying jabatan fungsional sk by uuid %s", err.Error())
	}
	return &jabatanFungsional, nil
}
