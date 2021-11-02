package repo

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"svc-insani-go/app"

	"svc-insani-go/modules/v1/pegawai/model"
)

func UpdatePendidikanPegawai(a *app.App, ctx context.Context, req model.PegawaiPendidikanRequest) error {
	db := a.GormDB.WithContext(ctx)
	var pegawaiPendidikanUpdate model.PegawaiPendidikanUpdate

	// Flag Ijazah Diakui
	if req.UuidPendidikanDiakui != "" {
		// Flag Ijazah Diakui ke Nul
		res := db.Model(&pegawaiPendidikanUpdate).
			Where("id_personal_data_pribadi = ? AND flag_aktif = 1", req.IdPersonalPegawai).
			Updates(map[string]interface{}{"flag_ijazah_diakui": 0, "user_update": req.UserUpdate})
		if res.Error != nil {
			return res.Error
		}

		res = db.Model(&pegawaiPendidikanUpdate).
			Where("uuid = ?", req.UuidPendidikanDiakui).
			Updates(map[string]interface{}{"flag_ijazah_diakui": 1, "user_update": req.UserUpdate})
		if res.Error != nil {
			return res.Error
		}

		if req.IdJenjangPendidikanDetailDiakui != nil {
			// Set all IdJenjangPendidikanDetailDiakui to null
			res := db.Model(&pegawaiPendidikanUpdate).
				Where("id_personal_data_pribadi = ? AND flag_aktif = 1", req.IdPersonalPegawai).
				Updates(map[string]interface{}{"id_jenjang_pdd_detail_diakui": nil, "user_update": req.UserUpdate})
			if res.Error != nil {
				return res.Error
			}

			res = db.Model(&pegawaiPendidikanUpdate).
				Where("uuid = ?", req.UuidPendidikanDiakui).
				Updates(map[string]interface{}{"id_jenjang_pdd_detail_diakui": req.IdJenjangPendidikanDetailDiakui, "user_update": req.UserUpdate})
			if res.Error != nil {
				return res.Error
			}
		}
	} else if req.UuidJenjangPendidikanTertinggiDiakui != "" {
		res := db.Model(&pegawaiPendidikanUpdate).
			Where("id_personal_data_pribadi = ? AND flag_aktif = 1", req.IdPersonalPegawai).
			Updates(map[string]interface{}{"flag_ijazah_diakui": 0, "id_jenjang_pdd_detail_diakui": nil, "user_update": req.UserUpdate})
		if res.Error != nil {
			return res.Error
		}
	}

	// Flag Ijazah Terakhir
	if req.UuidPendidikanTerakhir != "" {
		// Flag Ijazah Terakhir ke Nul
		res := db.Model(&pegawaiPendidikanUpdate).
			Where("id_personal_data_pribadi = ? AND flag_aktif = 1", req.IdPersonalPegawai).
			Updates(map[string]interface{}{"flag_ijazah_terakhir": 0, "user_update": req.UserUpdate})
		if res.Error != nil {
			return res.Error
		}
		res = db.Model(&pegawaiPendidikanUpdate).
			Where("uuid = ?", req.UuidPendidikanTerakhir).
			Updates(map[string]interface{}{"flag_ijazah_terakhir": 1, "user_update": req.UserUpdate})
		if res.Error != nil {
			return res.Error
		}

		if req.IdJenjangPendidikanDetailTerakhir != nil {
			// Set all IdJenjangPendidikanDetailTerakhir to null
			res := db.Model(&pegawaiPendidikanUpdate).
				Where("id_personal_data_pribadi = ? AND flag_aktif = 1", req.IdPersonalPegawai).
				Updates(map[string]interface{}{"id_jenjang_pdd_detail_terakhir": nil, "user_update": req.UserUpdate})
			if res.Error != nil {
				return res.Error
			}

			res = db.Model(&pegawaiPendidikanUpdate).Where("uuid = ?", req.UuidPendidikanTerakhir).
				Updates(map[string]interface{}{"id_jenjang_pdd_detail_terakhir": req.IdJenjangPendidikanDetailTerakhir, "user_update": req.UserUpdate})
			if res.Error != nil {
				return res.Error
			}
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
			&pegawaiPendidikan.IDJenjangPddDetailDiakui,
			&pegawaiPendidikan.KdJenjangPddDetailDiakui,
			&pegawaiPendidikan.NamaJenjangPddDetailDiakui,
			&pegawaiPendidikan.UuidJenjangPddDetailDiakui,
			&pegawaiPendidikan.IDJenjangPddDetailTerakhir,
			&pegawaiPendidikan.KdJenjangPddDetailTerakhir,
			&pegawaiPendidikan.NamaJenjangPddDetailTerakhir,
			&pegawaiPendidikan.UuidJenjangPddDetailTerakhir,
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
