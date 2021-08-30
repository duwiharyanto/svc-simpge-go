package repo

import (
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/pegawai/model"
)

func setBerkasPendukungWithURL(a *app.App, list model.BerkasPendukungList) {
	for i, berkas := range list {
		if berkas.PathFile == "" {
			continue
		}
		var err error
		minioBucketNamePersonal := "personal"
		list[i].URLFile, err = a.MinioClient.GetDownloadURL(minioBucketNamePersonal, berkas.PathFile, berkas.NamaFile)
		if err != nil {
			fmt.Printf("error get url berkas pendukung, %s", err.Error())
			list[i].URLFile = ""
		}
	}
}

func GetPegawaiFilePendidikan(a *app.App, id ...string) (model.BerkasPendukungList, error) {

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

func setIjazahWithURL(a *app.App, pendidikan *model.PegawaiPendidikan) {
	var err error
	minioBucketNamePersonal := "personal"

	if pendidikan.PathIjazah != "" {
		pendidikan.URLIjazah, err = a.MinioClient.GetDownloadURL(minioBucketNamePersonal, pendidikan.PathIjazah, pendidikan.NamaFileIjazah)
		if err != nil {
			fmt.Printf("error get url ijazah, %s", err.Error())
		}
	}

	if pendidikan.PathSKPenyetaraan != "" {
		pendidikan.URLSKPenyetaraan, err = a.MinioClient.GetDownloadURL(minioBucketNamePersonal, pendidikan.PathSKPenyetaraan, pendidikan.NamaFileSKPenyetaraan)
		if err != nil {
			fmt.Printf("error get url sk penyetaraan, %s", err.Error())
		}
	}
	// fmt.Println("URL Sk Penyetaraan : ", pendidikan.URLSKPenyetaraan)

}
