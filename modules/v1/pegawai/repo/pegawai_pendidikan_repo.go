package repo

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"svc-insani-go/app"

	"svc-insani-go/modules/v1/pegawai/model"
)

func UpdatePendidikanPegawai(a *app.App, ctx context.Context, uuidPendidikanDiakui, uuidPendidikanTerakhir string, idPersonalPegawai uint64) error {
	db := a.GormDB.WithContext(ctx)
	var pegawaiPendidikanUpdate model.PegawaiPendidikanUpdate

	// Flag Ijazah Diakui ke Nul
	res := db.Model(&pegawaiPendidikanUpdate).
		Where("id_personal_data_pribadi = ?", idPersonalPegawai).
		Update("flag_ijazah_diakui", "0")
	if res.Error != nil {
		return res.Error
	}

	// Flag Ijazah Diakui
	if uuidPendidikanDiakui != "" {
		res := db.Model(&pegawaiPendidikanUpdate).
			Where("uuid = ?", uuidPendidikanDiakui).
			Update("flag_ijazah_diakui", "1")
		if res.Error != nil {
			return res.Error
		}
	}

	// Flag Ijazah Terakhir ke Nul
	res = db.Model(&pegawaiPendidikanUpdate).
		Where("id_personal_data_pribadi = ?", idPersonalPegawai).
		Update("flag_ijazah_terakhir", "0")
	if res.Error != nil {
		return res.Error
	}

	// Flag Ijazah Terakhir
	if uuidPendidikanTerakhir != "" {
		res := db.Model(&pegawaiPendidikanUpdate).
			Where("uuid = ?", uuidPendidikanTerakhir).
			Update("flag_ijazah_terakhir", "1")
		if res.Error != nil {
			return res.Error
		}
	}

	return nil
}

func UpdatePendidikanDiakui(a *app.App, ctx context.Context, uuidPendidikanDiakui string, idPersonalPegawai string) error {
	db := a.GormDB.WithContext(ctx)

	var pegawaiPendidikanUpdate model.PegawaiPendidikanUpdate

	// Flag Ijazah Diakui ke Nul
	res := db.Model(&pegawaiPendidikanUpdate).
		Where("id_personal_data_pribadi = ?", idPersonalPegawai).
		Update("flag_ijazah_diakui", "0")
	if res.Error != nil {
		return res.Error
	}

	// Flag Ijazah Diakui
	if uuidPendidikanDiakui != "" {
		res := db.Model(&pegawaiPendidikanUpdate).
			Where("uuid = ?", uuidPendidikanDiakui).
			Update("flag_ijazah_diakui", "1")
		if res.Error != nil {
			return res.Error
		}

	}

	return nil
}

func GetPegawaiPendidikanPersonal(a *app.App, uuid string) ([]model.JenjangPendidikan, error) {
	sqlQuery := getPegawaiPendidikanPersonalQuery(uuid)

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
		setIjazahWithURL(a, &pegawaiPendidikan)
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

func GetPegawaiPendidikan(a *app.App, uuid string, withFile bool) (model.JenjangPendidikanList, error) {
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
		setIjazahWithURL(a, &pegawaiPendidikan)
		idPendidikanList = append(idPendidikanList, pegawaiPendidikan.IdPendidikan)
		m[fmt.Sprint(pegawaiPendidikan.KdJenjang, ".", pegawaiPendidikan.UrutanJenjang)] = append(m[fmt.Sprint(pegawaiPendidikan.KdJenjang, ".", pegawaiPendidikan.UrutanJenjang)], pegawaiPendidikan)

	}

	var filePegawaiList model.BerkasPendukungList
	if withFile {
		filePegawaiList, err = GetPegawaiFilePendidikan(a, idPendidikanList...)
		if err != nil {
			return nil, fmt.Errorf("error get file pendidikan: %w", err)
		}
		setBerkasPendukungWithURL(a, filePegawaiList)
	}

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
