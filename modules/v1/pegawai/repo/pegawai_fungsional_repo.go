package repo

import (
	"database/sql"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/pegawai/model"
)

func GetJabatanFungsionalYayasan(a *app.App, idPegawai string) (*model.PegawaiFungsionalYayasan, error) {
	sqlQuery := getPegawaiFungsionalYayasanByIdPegawai(idPegawai)
	var pegawaiFungsinalYayasan model.PegawaiFungsionalYayasan
	err := a.DB.QueryRow(sqlQuery).Scan(
		&pegawaiFungsinalYayasan.IdPegawai,
		&pegawaiFungsinalYayasan.JabatanFungsional,
		&pegawaiFungsinalYayasan.IdJabatanFungsional,
		&pegawaiFungsinalYayasan.KdJabatanFungsional,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error querying and scanning jabatan fungsional yayasan, %w", err)
	}

	return &pegawaiFungsinalYayasan, nil
}

func GetJabatanFungsionalNegara(a *app.App, idPegawai string) (*model.PegawaiFungsionalNegara, error) {
	sqlQuery := getPegawaiFungsionalNegaraByIdPegawai(idPegawai)
	var pegawaiFungsinalNegara model.PegawaiFungsionalNegara
	err := a.DB.QueryRow(sqlQuery).Scan(
		&pegawaiFungsinalNegara.IdPegawai,
		&pegawaiFungsinalNegara.JabatanFungsional,
		&pegawaiFungsinalNegara.IdJabatanFungsional,
		&pegawaiFungsinalNegara.KdJabatanFungsional,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error querying and scanning jabatan fungsional yayasan, %w", err)
	}

	return &pegawaiFungsinalNegara, nil
}
