package repo

import (
	"context"
	"database/sql"
	"fmt"
	"sort"
	"strings"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/pegawai/model"
)

const (
	BucketIjazah    = "personal-ijazah"
	JenisFileIjazah = "ijazah"
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
	err := a.DB.QueryRow(sqlQuery).Scan(&pegawai.ID,
		&pegawai.NIK,
		&pegawai.Nama,
		&pegawai.GelarDepan,
		&pegawai.GelarBelakang,
		&pegawai.JenisPegawai.KDJenisPegawai,
		&pegawai.JenisPegawai.JenisPegawai,
		&pegawai.JenisPegawai.UUID,
		&pegawai.KelompokPegawai.KdKelompokPegawai,
		&pegawai.KelompokPegawai.KelompokPegawai,
		&pegawai.KelompokPegawai.UUID,
		&pegawai.UnitKerja.KdUnitKerja,
		&pegawai.UnitKerja.NamaUnitKerja,
		&pegawai.UnitKerja.UUID,
		&pegawai.UUID,
	)
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
		&pegawaiYayasan.UuidJenisPegawai,
		&pegawaiYayasan.KDJenisPegawai,
		&pegawaiYayasan.JenisPegawai,
		&pegawaiYayasan.UuidStatusPegawai,
		&pegawaiYayasan.KDStatusPegawai,
		&pegawaiYayasan.StatusPegawai,
		&pegawaiYayasan.UuidKelompokPegawai,
		&pegawaiYayasan.KdKelompokPegawai,
		&pegawaiYayasan.KelompokPegawai,
		&pegawaiYayasan.UuidPangkatGolongan,
		&pegawaiYayasan.KdPangkat,
		&pegawaiYayasan.Pangkat,
		&pegawaiYayasan.Golongan,
		&pegawaiYayasan.KdRuang,
		&pegawaiYayasan.UuidJabatanFungsional,
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
		&pegawaiYayasan.UuidJenisRegis,
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

	pegawaiYayasan.SetTanggalIDN()

	return &pegawaiYayasan, nil
}

func GetUnitKerjaPegawai(a app.App, uuid string) (*model.UnitKerjaPegawai, error) {
	sqlQuery := getUnitKerjaPegawaiQuery(uuid)
	var unitKerjaPegawai model.UnitKerjaPegawai

	err := a.DB.QueryRow(sqlQuery).Scan(
		&unitKerjaPegawai.UuidIndukKerja,
		&unitKerjaPegawai.KdIndukKerja,
		&unitKerjaPegawai.IndukKerja,
		&unitKerjaPegawai.UuidUnitKerja,
		&unitKerjaPegawai.KdUnitKerja,
		&unitKerjaPegawai.UnitKerja,
		&unitKerjaPegawai.UuidBagianKerja,
		&unitKerjaPegawai.KdBagianKerja,
		&unitKerjaPegawai.BagianKerja,
		&unitKerjaPegawai.UuidLokasiKerja,
		&unitKerjaPegawai.LokasiKerja,
		&unitKerjaPegawai.LokasiDesc,
		&unitKerjaPegawai.NoSkPertama,
		&unitKerjaPegawai.TmtSkPertama,
		&unitKerjaPegawai.KdHomebasePddikti,
		&unitKerjaPegawai.KdHomebaseUii,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("error querying and scanning unit kerja pegawai, %s", err.Error())
	}

	unitKerjaPegawai.SetTanggalIDN()

	return &unitKerjaPegawai, nil
}

func GetPegawaiPNS(a app.App, uuid string) (*model.PegawaiPNSPTT, error) {
	sqlQuery := getPegawaiPNSQuery(uuid)
	var pegawaiPNSPTT model.PegawaiPNSPTT

	err := a.DB.QueryRow(sqlQuery).Scan(
		&pegawaiPNSPTT.NipPNS,
		&pegawaiPNSPTT.NoKartuPegawai,
		&pegawaiPNSPTT.UuidJenisPTT,
		&pegawaiPNSPTT.KdJenisPTT,
		&pegawaiPNSPTT.JenisPTT,
		&pegawaiPNSPTT.InstansiAsalPtt,
		&pegawaiPNSPTT.UuidPangkatGolongan,
		&pegawaiPNSPTT.KdPangkatGolonganPns,
		&pegawaiPNSPTT.PangkatPNS,
		&pegawaiPNSPTT.GolonganPNS,
		&pegawaiPNSPTT.TmtPangkatGolongan,
		&pegawaiPNSPTT.UuidJabatanPns,
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

	pegawaiPNSPTT.SetTanggalIDN()
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

func GetPegawaiFilePendidikan(a app.App, id ...string) (model.BerkasPendukungList, error) {

	sqlQuery := getPegawaiFilePendidikanQuery(id...)

	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get pendidikan file, %w", err)
	}
	defer rows.Close()

	berkasPendukungList := model.BerkasPendukungList{}
	for rows.Next() {
		var pegawaiFilePendidikan model.BerkasPendukung
		err := rows.Scan(
			&pegawaiFilePendidikan.KdJenisFile,
			&pegawaiFilePendidikan.JenisFile,
			&pegawaiFilePendidikan.PathFile,
			&pegawaiFilePendidikan.IDPendidikan,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning berkas pendukung, %w", err)
		}
		pegawaiFilePendidikan.SetDownloadFileName(a.TimeLocation)
		berkasPendukungList = append(berkasPendukungList, pegawaiFilePendidikan)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error from berkas pendukung rows, %w", err)
	}

	return berkasPendukungList, nil
}

func GetPegawaiPendidikan(a app.App, uuid string) ([]model.JenjangPendidikan, error) {
	sqlQuery := getPegawaiPendidikanQuery(uuid)

	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get pendidikan file, %w", err)
	}
	defer rows.Close()

	m := make(map[string][]model.PegawaiPendidikan)
	idPendidikanList := []string{}
	for rows.Next() {
		var pegawaiPendidikan model.PegawaiPendidikan
		err := rows.Scan(
			&pegawaiPendidikan.UuidPendidikan,
			&pegawaiPendidikan.IdPendidikan,
			&pegawaiPendidikan.KdJenjang,
			&pegawaiPendidikan.UrutanJenjang,
			&pegawaiPendidikan.NamaInstitusi,
			&pegawaiPendidikan.Jurusan,
			&pegawaiPendidikan.TglKelulusan,
			&pegawaiPendidikan.FlagIjazahDiakui,
			&pegawaiPendidikan.FlagIjazahTerakhir,
			&pegawaiPendidikan.Akreditasi,
			&pegawaiPendidikan.KonsentrasiBidang,
			&pegawaiPendidikan.Gelar,
			&pegawaiPendidikan.NomorInduk,
			&pegawaiPendidikan.TahunMasuk,
			&pegawaiPendidikan.JudulTugasAkhir,
			&pegawaiPendidikan.FlagInstitusiLuarNegeri,
			&pegawaiPendidikan.NomorIjazah,
			&pegawaiPendidikan.TglIjazah,
			&pegawaiPendidikan.PathIjazah,
			&pegawaiPendidikan.FlagIjazahTerverifikasi,
			&pegawaiPendidikan.Nilai,
			&pegawaiPendidikan.JumlahPelajaran,
			&pegawaiPendidikan.PathSKPenyetaraan,
			&pegawaiPendidikan.NomorSKPenyetaraan,
			&pegawaiPendidikan.TglSKPenyetaraan,
			&pegawaiPendidikan.UUIDPersonal,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning pendidikan pegawai, %w", err)
		}
		pegawaiPendidikan.SetTanggalIDN()
		pegawaiPendidikan.SetNamaFileIjazah()
		pegawaiPendidikan.SetNamaFilePenyetaraan()
		pegawaiPendidikan.SetDownloadFileNamePendidikan(a.TimeLocation)
		idPendidikanList = append(idPendidikanList, pegawaiPendidikan.IdPendidikan)
		m[fmt.Sprint(pegawaiPendidikan.KdJenjang, ".", pegawaiPendidikan.UrutanJenjang)] = append(m[fmt.Sprint(pegawaiPendidikan.KdJenjang, ".", pegawaiPendidikan.UrutanJenjang)], pegawaiPendidikan)

	}

	filePegawaiList, err := GetPegawaiFilePendidikan(a, idPendidikanList...)
	if err != nil {
		return nil, fmt.Errorf("error get file pendidikan: %w", err)
	}

	setBerkasPendukungWithURL(a, filePegawaiList)

	filePegawaiMap := filePegawaiList.MapByIdPendidikan()

	jenjangPendidikan := []model.JenjangPendidikan{}
	for kdJenjangUrut, pendidikanList := range m {
		for n, pendidikan := range pendidikanList {
			pendidikanList[n].BerkasPendukungList = filePegawaiMap[pendidikan.IdPendidikan]
		}
		splittedKdJenjangUrut := strings.Split(kdJenjangUrut, ".")
		kdJenjang := splittedKdJenjangUrut[0]
		urutanJenjang := splittedKdJenjangUrut[1]
		jenjangPendidikan = append(jenjangPendidikan, model.JenjangPendidikan{
			JenjangPendidikan: kdJenjang,
			UrutanJenjang:     urutanJenjang,
			Data:              pendidikanList,
		})
	}

	sort.SliceStable(jenjangPendidikan, func(i, j int) bool {
		return jenjangPendidikan[i].UrutanJenjang < jenjangPendidikan[j].UrutanJenjang
	})

	if err != nil {
		return nil, fmt.Errorf("error querying and scanning data pendidikan pegawai, %s", err.Error())
	}

	return jenjangPendidikan, nil
}

func GetURLFilePendidikan(a app.App, pendidikanList []*model.PegawaiPendidikan) []*model.PegawaiPendidikan {
	for _, pendidikan := range pendidikanList {
		if pendidikan.PathIjazah != "" {
			var err error
			pendidikan.URLIjazah, err = a.MinioClient.GetDownloadURL(a.MinioBucketName, pendidikan.PathIjazah, "")
			if err != nil {
				fmt.Printf("error getting file url, %s", err.Error())
				pendidikan.URLIjazah = ""
			}
		}
	}
	return pendidikanList
}

func setBerkasPendukungWithURL(a app.App, list model.BerkasPendukungList) {
	for i, berkas := range list {
		if berkas.PathFile == "" {
			continue
		}
		var err error
		list[i].URLFile, err = a.MinioClient.GetDownloadURL(a.MinioBucketName, berkas.PathFile, berkas.NamaFile)
		if err != nil {
			fmt.Printf("error get url berkas pendukung, %s", err.Error())
			list[i].URLFile = ""
		}
	}
}

func GetAllPegawaix(a app.App, ctx context.Context, limit int, offset int) (*model.PegawaiResponseTest, error) {

	var pegawaiAll []model.Pegawai2
	var count int64

	tx := a.GormDB.WithContext(ctx)
	results := model.PegawaiResponseTest{}
	res := tx.Limit(limit).
		Offset(offset).
		Find(&pegawaiAll).Count(&count)
	if res.Error != nil {
		return nil, res.Error
	}

	countConv := int(count)

	results.Data = pegawaiAll
	results.Count = countConv

	return &results, nil
}

func GetPegawaiByUUIDx(a app.App, ctx context.Context, uuid string) (*model.Pegawai, error) {

	var pegawaiAll model.Pegawai
	tx := a.GormDB.WithContext(ctx)

	res := tx.First(&pegawaiAll, "uuid = ?", uuid)
	if res.Error != nil {
		return nil, res.Error
	}

	return &pegawaiAll, nil
}

func GetOldPegawai(a app.App, ctx context.Context, uuidPegawai string) (model.PegawaiUpdate, error) {

	var pegawaiOld model.PegawaiUpdate

	db := a.GormDB.WithContext(ctx)
	res := db.Joins("PegawaiPNS").
		Joins("PegawaiFungsional").
		Find(&pegawaiOld, "pegawai.uuid = ?", uuidPegawai)
	if res.Error != nil {
		return model.PegawaiUpdate{}, res.Error
	}

	return pegawaiOld, nil
}

func UpdatePegawai(a app.App, ctx context.Context, pegawaiUpdate model.PegawaiUpdate) error {
	db := a.GormDB.WithContext(ctx)

	res := db.Save(&pegawaiUpdate)
	if res.Error != nil {
		return res.Error
	}

	res = db.Save(&pegawaiUpdate.PegawaiPNS)
	if res.Error != nil {
		return res.Error
	}

	res = db.Save(&pegawaiUpdate.PegawaiFungsional)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func UpdatePendidikanPegawai(a app.App, ctx context.Context, uuidPendidikanDiakui string, uuidPendidikanTerakhir string, idPersonalPegawai string) error {
	db := a.GormDB.WithContext(ctx)

	var pegawaiPendidikanUpdate model.PegawaiPendidikanUpdate

	// Flag Ijazah Diakui
	if uuidPendidikanDiakui != "" {
		res := db.Model(&pegawaiPendidikanUpdate).
			Where("uuid = ?", uuidPendidikanDiakui).
			Update("flag_ijazah_diakui", "1")
		if res.Error != nil {
			return res.Error
		}

		// Flag Ijazah ke Nul
		res = db.Model(&pegawaiPendidikanUpdate).
			Where("id_personal_data_pribadi = ? AND uuid != ?", idPersonalPegawai, uuidPendidikanDiakui).
			Update("flag_ijazah_diakui", "0")
		if res.Error != nil {
			return res.Error
		}
	}

	// Flag Ijazah Terakhir
	if uuidPendidikanTerakhir != "" {
		res := db.Model(&pegawaiPendidikanUpdate).
			Where("uuid = ?", uuidPendidikanTerakhir).
			Update("flag_ijazah_terakhir", "1")
		if res.Error != nil {
			return res.Error
		}

		// Flag Ijazah ke Nul
		res = db.Model(&pegawaiPendidikanUpdate).
			Where("id_personal_data_pribadi = ? AND uuid != ?", idPersonalPegawai, uuidPendidikanTerakhir).
			Update("flag_ijazah_diakui", "0")
		if res.Error != nil {
			return res.Error
		}
	}

	return nil
}
