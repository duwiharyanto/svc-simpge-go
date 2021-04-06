package repo

import (
	"database/sql"
	"fmt"
	"svc-insani-go/modules/v1/sk-pegawai/model"

	guuid "github.com/google/uuid"
)

// func GetSKPegawai(a app.App, IDPegawai string) ([]model.SKPegawai, error) {
// 	sqlQuery := getAllSKPegawaibyUUIDQuery(IDPegawai)
// 	rows, err := a.DB.Query(sqlQuery)
// 	if err != nil {
// 		return nil, fmt.Errorf("error querying get sk pengangkatan pegawai, %s", err.Error())
// 	}
// 	defer rows.Close()
// 	SKPegawai := []model.SKPegawai{}
// 	for rows.Next() {
// 		var s model.SKPegawai
// 		err := rows.Scan( //&s.NIK,
// 			//&s.KdJenisSK,
// 			&s.IDJenisSK,
// 			&s.NomorSK,
// 			&s.TentangSK,
// 			&s.TMT,
// 			&s.UUID)
// 		if err != nil {
// 			return nil, fmt.Errorf("error scan sk pengangkatan pegawai row, %s", err.Error())
// 		}
// 		SKPegawai = append(SKPegawai, s)
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, fmt.Errorf("error status sk pengangkatan pegawai rows, %s", err.Error())
// 	}

// 	return SKPegawai, nil
// }
// func GetSKPengangkatan(a app.App, IDSKPegawai string) (skPengangkatanModel.SKPengangkatanPegawai, error) {
// 	// c.Param("kd_jenis_pegawai")
// 	sqlQuery := getSKPengangkatanQuery(IDSKPegawai)
// 	//fmt.Printf("log query : %s\n", sqlQuery)
// 	var s skPengangkatanModel.SKPengangkatanPegawai
// 	err := a.DB.QueryRow(sqlQuery).Scan(
// 		&s.IDUnitPengangkat,
// 		&s.IDUnitPegawai,
// 		&s.MasaRilBulan,
// 		&s.MasaRilTahun,
// 		&s.MasaGajiBulan,
// 		&s.MasaGajiTahun,
// 		&s.IDPangkatPegawai,
// 		&s.IDGolonganPegawai,
// 		&s.IDStatusPengangkatan,
// 		&s.IDJenisIjazah,
// 		&s.PathSKPengangkatan,
// 	)
// 	if err == sql.ErrNoRows {
// 		return skPengangkatanModel.SKPengangkatanPegawai{}, nil
// 	}
// 	if err != nil {
// 		return skPengangkatanModel.SKPengangkatanPegawai{}, fmt.Errorf("error querying or scanning get sk pengangkatan pegawai, %s", err.Error())
// 	}

// 	return s, nil
// }

func CreateSKPegawai(tx *sql.Tx, skPegawai model.SKPegawai) (string, error) {
	uuid := guuid.New().String()
	id := guuid.New().ID()
	skPegawai.UUID = uuid
	skPegawai.ID = fmt.Sprintf("%d", id)

	//fmt.Printf("\n log id sk pegawai : %+v \n", skPegawai.ID)
	sqlQuery := createSKPegawaiQuery(skPegawai)
	//fmt.Printf("log query sk pegawai %s\n", sqlQuery)
	res, err := tx.Exec(sqlQuery)
	if err != nil {
		return "", fmt.Errorf("error executing create sk pegawai query, %w", err)
	}
	affectedRows, err := res.RowsAffected()
	if err != nil {
		return "", fmt.Errorf("error executing get affected rows, %w", err)
	}
	if affectedRows == 0 {
		return "", sql.ErrNoRows
	}
	return skPegawai.ID, nil
}

func UpdateSKPegawai() {

}
