package repo

import (
	"database/sql"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/pegawai/model"
)

func GetAllPegawai(a app.App, req *model.PegawaiRequest) ([]model.Pegawai, error) {
	sqlQuery := getListAllPegawaiQuery(req)
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get pegawai, %w", err)
	}
	defer rows.Close()

	pp := []model.Pegawai{}
	for rows.Next() {
		var p model.Pegawai
		err := rows.Scan(
			&p.NIK,
			&p.Nama,
			&p.GelarDepan,
			&p.GelarBelakang,
			&p.KelompokPegawai.KdKelompokPegawai,
			&p.KelompokPegawai.KelompokPegawai,
			&p.KelompokPegawai.UUID,
			&p.UnitKerja.KdUnitKerja,
			&p.UnitKerja.NamaUnitKerja,
			&p.UnitKerja.UUID,
			&p.JenisPegawai.KDJenisPegawai,
			&p.JenisPegawai.JenisPegawai,
			&p.JenisPegawai.UUID,
			&p.StatusPegawai.KDStatusPegawai,
			&p.StatusPegawai.StatusPegawai,
			&p.StatusPegawai.UUID,
			&p.UUID,
		)
		if err != nil {
			return nil, fmt.Errorf("error scan pegawai row, %s", err.Error())
		}
		p.SetFlagDosen()
		pp = append(pp, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error pegawai rows, %s", err.Error())
	}

	return pp, nil
}

func CountPegawai(a app.App, req *model.PegawaiRequest) (int, error) {
	sqlQuery := countPegawaiQuery(req)
	var count int
	err := a.DB.QueryRow(sqlQuery).Scan(&count)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	if err != nil {
		return 0, fmt.Errorf("error querying count pegawai, %w", err)
	}

	return count, nil

}
func GetPegawaiByUUID(a app.App, uuid string) (*model.Pegawai, error) {
	sqlQuery := getPegawaiByUUID(uuid)
	var pegawai model.Pegawai
	err := a.DB.QueryRow(sqlQuery).Scan(&pegawai.ID, &pegawai.NIK, &pegawai.Nama, &pegawai.GelarDepan, &pegawai.GelarBelakang, &pegawai.UnitKerja.KdUnitKerja, &pegawai.KelompokPegawai.KdKelompokPegawai, &pegawai.UUID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error querying get pegawai by uuid, %s", err.Error())
	}

	return &pegawai, nil

}

func GetKepegawaianYayasan(a app.App, uuid string) (*model.PegawaiYayasan, error) {
	sqlQuery := getPegawaiYayasanQuery(uuid)
	var pegawaiYayasan model.PegawaiYayasan

	err := a.DB.QueryRow(sqlQuery).Scan(
		&pegawaiYayasan.KDJenisPegawai,
		&pegawaiYayasan.JenisPegawai,
		&pegawaiYayasan.KDStatusPegawai,
		&pegawaiYayasan.StatusPegawai,
		&pegawaiYayasan.KdKelompokPegawai,
		&pegawaiYayasan.KelompokPegawai,
		&pegawaiYayasan.KdPangkat,
		&pegawaiYayasan.Pangkat,
		&pegawaiYayasan.Golongan,
		&pegawaiYayasan.TmtPangkatGolongan,
		&pegawaiYayasan.KdJabatanFungsional,
		&pegawaiYayasan.JabatanFungsional,
		&pegawaiYayasan.TmtJabatan,
		&pegawaiYayasan.MasaKerjaBawaanTahun,
		&pegawaiYayasan.MasaKerjaBawaanBulan,
		&pegawaiYayasan.MasaKerjaGajiTahun,
		&pegawaiYayasan.MasaKerjaGajiBulan,
		&pegawaiYayasan.MasaKerjaTotalahun,
		&pegawaiYayasan.MasaKerjaTotalBulan,
		&pegawaiYayasan.AngkaKredit,
		&pegawaiYayasan.NoSertifikasi,
		&pegawaiYayasan.KdJenisRegis,
		&pegawaiYayasan.JenisNomorRegis,
		&pegawaiYayasan.NomorRegis,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("error querying and scanning kepegawaian yayasan, %s", err.Error())
	}

	return &pegawaiYayasan, nil
}

func GetUnitKerjaPegawai(a app.App, uuid string) (*model.UnitKerjaPegawai, error) {
	sqlQuery := getUnitKerjaPegawaiQuery(uuid)
	var unitKerjaPegawai model.UnitKerjaPegawai

	err := a.DB.QueryRow(sqlQuery).Scan(
		&unitKerjaPegawai.KdIndukKerja,
		&unitKerjaPegawai.IndukKerja,
		&unitKerjaPegawai.KdUnitKerja,
		&unitKerjaPegawai.UnitKerja,
		&unitKerjaPegawai.KdBagianKerja,
		&unitKerjaPegawai.BagianKerja,
		&unitKerjaPegawai.LokasiKerja,
		&unitKerjaPegawai.LokasiDesc,
		&unitKerjaPegawai.NoSkPertama,
		&unitKerjaPegawai.TmtSkPertama,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("error querying and scanning unit kerja pegawai, %s", err.Error())
	}

	return &unitKerjaPegawai, nil
}

func GetPegawaiPNS(a app.App, uuid string) (*model.PegawaiPNSPTT, error) {
	sqlQuery := getPegawaiPNSQuery(uuid)
	var pegawaiPNSPTT model.PegawaiPNSPTT

	err := a.DB.QueryRow(sqlQuery).Scan(
		&pegawaiPNSPTT.NipPNS,
		&pegawaiPNSPTT.NoKartuPegawai,
		&pegawaiPNSPTT.KdPangkatGolongan,
		&pegawaiPNSPTT.PangkatPNS,
		&pegawaiPNSPTT.GolonganPNS,
		&pegawaiPNSPTT.TmtPangkatGolongan,
		&pegawaiPNSPTT.KdJabatanPns,
		&pegawaiPNSPTT.JabatanPns,
		&pegawaiPNSPTT.TmtJabatanPns,
		&pegawaiPNSPTT.MasaKerjaPnsTahun,
		&pegawaiPNSPTT.MasaKerjaPnsBulan,
		&pegawaiPNSPTT.AngkaKreditPns,
		&pegawaiPNSPTT.KeteranganPNS,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("error querying and scanning pegawai pns, %s", err.Error())
	}

	return &pegawaiPNSPTT, nil
}

func GetPegawaiPTT(a app.App, uuid string) (*model.PegawaiPNSPTT, error) {
	sqlQuery := getPegawaiPTTQuery(uuid)
	var pegawaiPNSPTT model.PegawaiPNSPTT

	err := a.DB.QueryRow(sqlQuery).Scan(
		&pegawaiPNSPTT.KdJenisPTT,
		&pegawaiPNSPTT.JenisPTT,
		&pegawaiPNSPTT.InstansiAsalPtt,
		&pegawaiPNSPTT.KeteranganPtt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("error querying and scanning pegawai tidak tetap, %s", err.Error())
	}

	return &pegawaiPNSPTT, nil
}

func GetStatusPegawaiAktif(a app.App, uuid string) (*model.StatusAktif, error) {
	sqlQuery := getStatusPegawaiAktifQuery(uuid)
	var statusAktif model.StatusAktif

	err := a.DB.QueryRow(sqlQuery).Scan(
		&statusAktif.FlagAktifPegawai,
		&statusAktif.KdStatusAktifPegawai,
		&statusAktif.StatusAktifPegawai,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("error querying and scanning pegawai tidak tetap, %s", err.Error())
	}

	return &statusAktif, nil
}

func GetPegawaiPribadi(a app.App, uuid string) (*model.PegawaiPribadi, error) {
	sqlQuery := getPegawaiPribadiQuery(uuid)
	var pegawaiPribadi model.PegawaiPribadi

	err := a.DB.QueryRow(sqlQuery).Scan(
		&pegawaiPribadi.Nama,
		&pegawaiPribadi.NIK,
		&pegawaiPribadi.JenisPegawai,
		&pegawaiPribadi.KelompokPegawai,
		&pegawaiPribadi.UnitKerja,
		&pegawaiPribadi.UUID,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("error querying and scanning data pribadi pegawai, %s", err.Error())
	}

	return &pegawaiPribadi, nil
}
