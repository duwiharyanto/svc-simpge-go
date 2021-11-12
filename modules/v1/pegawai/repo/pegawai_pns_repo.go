package repo

import (
	"database/sql"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/pegawai/model"
)

func GetPegawaiPNS(a *app.App, uuid string) (*model.PegawaiPNSPTT, error) {
	sqlQuery := getPegawaiPNSQuery(uuid)
	var pegawaiPNSPTT model.PegawaiPNSPTT

	err := a.DB.QueryRow(sqlQuery).Scan(
		&pegawaiPNSPTT.NipPNS,
		&pegawaiPNSPTT.NoKartuPegawai,
		&pegawaiPNSPTT.UuidDetailProfesi,
		&pegawaiPNSPTT.DetailProfesi,
		&pegawaiPNSPTT.UuidJenisPTT,
		&pegawaiPNSPTT.KdJenisPTT,
		&pegawaiPNSPTT.JenisPTT,
		&pegawaiPNSPTT.InstansiAsalPtt,
		&pegawaiPNSPTT.UuidPangkatGolongan,
		&pegawaiPNSPTT.KdPangkatGolonganPns,
		&pegawaiPNSPTT.PangkatPNS,
		&pegawaiPNSPTT.GolonganPNS,
		&pegawaiPNSPTT.KdGolonganPNS,
		&pegawaiPNSPTT.KdRuangPNS,
		&pegawaiPNSPTT.TmtPangkatGolongan,
		&pegawaiPNSPTT.UuidJabatanPns,
		&pegawaiPNSPTT.KdJabatanPns,
		&pegawaiPNSPTT.JabatanPns,
		&pegawaiPNSPTT.TmtJabatanPns,
		&pegawaiPNSPTT.MasaKerjaPnsTahun,
		&pegawaiPNSPTT.MasaKerjaPnsBulan,
		&pegawaiPNSPTT.AngkaKreditPns,
		&pegawaiPNSPTT.KeteranganPNS,
		&pegawaiPNSPTT.Nira,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("error querying and scanning pegawai pns, %s", err.Error())
	}

	pegawaiPNSPTT.SetTanggalIDN()
	return &pegawaiPNSPTT, nil
}
