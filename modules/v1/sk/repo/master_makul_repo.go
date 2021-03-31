package repo

import (
	"database/sql"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/sk/model"
)

func GetAllMataKuliah(a app.App) ([]model.MataKuliah, error) {
	// c.Param("kd_jenis_pegawai")
	sqlQuery := getMataKuliahQuery()
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get jabatan fungsional pegawai, %s", err.Error())
	}
	defer rows.Close()

	makul := []model.MataKuliah{}
	for rows.Next() {
		var s model.MataKuliah
		err := rows.Scan(&s.KdMakul, &s.NamaMakul, &s.UUID)
		if err != nil {
			return nil, fmt.Errorf("error scan jabatan fungsional row, %s", err.Error())
		}
		makul = append(makul, s)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error jabatan fungsional rows, %s", err.Error())
	}

	return makul, nil
}

func GetMataKuliahByUUID(a app.App, uuid string) (*model.MataKuliah, error) {
	sqlQuery := getMataKuliahByUUID(uuid)
	//fmt.Printf("[DEBUG] unit pengangkat by uuid:\n%s\n", sqlQuery)
	var makul model.MataKuliah
	err := a.DB.QueryRow(sqlQuery).Scan(&makul.ID, &makul.KdMakul, &makul.NamaMakul, &makul.UUID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error querying jabatan fungsional sk by uuid %s", err.Error())
	}
	return &makul, nil
}

func GetMataKuliahIDByUUID(a app.App, uuidMataKuliah []string) ([]string, error) {
	// c.Param("kd_jenis_pegawai")
	sqlQuery := getMataKuliahIDByUUID(uuidMataKuliah)
	fmt.Printf("[DEBUG] getMataKuliahIDByUUID:\n%s\n", sqlQuery)
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get mata kuliah id by uuid, %s", err.Error())
	}
	defer rows.Close()

	makul := []string{}
	for rows.Next() {
		var idMataKuliah string
		err := rows.Scan(&idMataKuliah)
		if err != nil {
			return nil, fmt.Errorf("error scan mata kuliah id by uuid, %s", err.Error())
		}
		makul = append(makul, idMataKuliah)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error mata kuliah id by uuid, %s", err.Error())
	}

	return makul, nil
}
