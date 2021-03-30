package repo

import (
	"database/sql"
	"fmt"
	"mime/multipart"
	"strings"
	"svc-insani-go/app"
	skPegawaiModel "svc-insani-go/modules/v1/sk-pegawai/model"
	skPegawaiRepo "svc-insani-go/modules/v1/sk-pegawai/repo"
	"svc-insani-go/modules/v1/sk-pengangkatan/model"
	"time"
)

const (
	BucketSKDosen = "sk-pegawai"
	//BucketSKDosen   = "sk-dosen"
)

func CreateSKPengangkatanDosen(a app.App, skPegawai skPegawaiModel.SKPegawai) error {
	// skPengangkatan.PathSKPengangkatan = PrepareFileUploadPath(a, skPengangkatan.FileSKPengangkatan, skPengangkatan.UUIDSKUtama)
	tx, err := a.DB.Begin()
	if err != nil {
		return fmt.Errorf("error beginning transaction, %w", err)
	}
	idSKPegawai, err := skPegawaiRepo.CreateSKPegawai(tx, skPegawai)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing create sk pegawai, %w", err)
	}

	skPegawai.SKPengangkatanDosen.PathSKDosen = PrepareFileUploadPathDosen(a, skPegawai.SKPengangkatanDosen.FileSKDosen, idSKPegawai)
	//fmt.Printf("log skPegawai.SKPengangkatanDosen.PathSKDosen : %+v\n", skPegawai.SKPengangkatanDosen.PathSKDosen)
	sqlQuery := createSKPengangkatanDosenQuery(skPegawai.SKPengangkatanDosen, idSKPegawai)

	res, err := tx.Exec(sqlQuery)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing create sk pengangkatan Dosen query, %w", err)
	}
	affectedRows, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing get affected rows, %w", err)
	}
	if affectedRows == 0 {
		tx.Rollback()
		return sql.ErrNoRows
	}
	err = ExecMakulSkPengangkatanDosen(tx, idSKPegawai, skPegawai.UserUpdate, skPegawai.SKPengangkatanDosen.IDMataKuliah)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error insert table pivot makul sk pengangkatan dosen, %w", err)
	}
	err = UploadFIleSKDosen(a, BucketSKDosen, skPegawai.SKPengangkatanDosen.PathSKDosen, skPegawai.SKPengangkatanDosen.FileSKDosen)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error upload file sk pengangkatan Dosen, %w", err)
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error commiting query, %w", err)
	}
	return nil
}

func UploadFIleSKDosen(a app.App, bucketName, filePath string, fileSKPengangkatanDosen *multipart.FileHeader) error {
	if fileSKPengangkatanDosen == nil {
		return nil
	}
	file, err := fileSKPengangkatanDosen.Open()
	if err != nil {
		return fmt.Errorf("error opening file, %w", err)
	}
	defer file.Close()

	fileContentType := fileSKPengangkatanDosen.Header.Get("Content-Type")
	err = a.MinioClient.Upload(bucketName, filePath, file, fileSKPengangkatanDosen.Size, fileContentType)
	if err != nil {
		return fmt.Errorf("error uploading file, %w", err)
	}
	return nil
}

func PrepareFileUploadPathDosen(a app.App, fileSKPengangkatanDosen *multipart.FileHeader, idSKUtama string) string {
	if fileSKPengangkatanDosen == nil {
		fmt.Printf("PRINT DEBUG FILE SK Pengangkatan kosong : \n")
		return ""
	}
	file, err := fileSKPengangkatanDosen.Open()
	if err != nil {
		return ""
	}
	defer file.Close()
	unixTime := time.Now().In(a.TimeLocation).Unix()
	skPengangkatanFileContentType := fileSKPengangkatanDosen.Header.Get("Content-Type")
	skPengangkatanFileExtension := strings.Split(skPengangkatanFileContentType, "/")[1]
	skPengangkatanMinioFolderName := idSKUtama
	skPengangkatanMinioFileName := fmt.Sprintf("%d", unixTime)
	skPengangkatanMinioUploadPath := fmt.Sprintf("%s/%s.%s", skPengangkatanMinioFolderName, skPengangkatanMinioFileName, skPengangkatanFileExtension)
	return skPengangkatanMinioUploadPath
}

func GetDetailSKPengangkatanDosen(a app.App, UUIDSKPengangkatanDosen string) ([]model.SKPengangkatanDosen, error) {
	sqlQuery := getDetailSKPengangkatanDosenQuery(UUIDSKPengangkatanDosen)
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get detail sk pengangkatan Dosen. %s", err.Error())
	}
	defer rows.Close()

	SKPengangkatanDosen := []model.SKPengangkatanDosen{}
	for rows.Next() {
		var s model.SKPengangkatanDosen
		err := rows.Scan(
			&s.IDJenisSKPengangkatan,
			&s.MasaKerjaDiakuiBulanLama,
			&s.MasaKerjaDiakuiTahunBaru,
			&s.IDPangkatGolonganPegawaiLama,
			&s.IDJabatanFungsionalLama,
			&s.IDIndukKerjaBaru,
			&s.IDPangkatGolonganPegawaiBaru,
			&s.GajiPokok,
			&s.TunjanganBeras,
			&s.TunjanganKhusus,
			&s.SKSMengajar,
			&s.BantuanKomunikasi,
			&s.TunjanganTahunan,
			&s.MasaKerjaRilBulanBaru,
			&s.MasaKerjaRilTahunBaru,
			&s.MasaKerjaGajiBulanBaru,
			&s.MasaKerjaGajiTahunBaru,
			&s.IDJabatanPenetap,
			&s.IDPegawaiPenetap,
			&s.TanggalDitetapkan,
			&s.IDUnitKerja,
			&s.IDJabatanFungsionalBaru,
			&s.IDJenisIjazah,
			&s.InstansiKerja,
			&s.TglBerakhir,
			&s.JangkaWaktuEvaluasi,
			&s.PathSKDosen,
			&s.UserInput,
			&s.UserUpdate,
			&s.IDSKPegawai,
			&s.UUIDSKPengangkatanDosen)
		if err != nil {
			return nil, fmt.Errorf("error scan sk pengangkatan Dosen %s", err.Error())
		}
		SKPengangkatanDosen = append(SKPengangkatanDosen, s)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error sk pengangkatan Dosen pegawai rows, %s", err.Error())
	}
	return SKPengangkatanDosen, nil
}

func ExecMakulSkPengangkatanDosen(tx *sql.Tx, idSkPengangkatanDosen, userUpdate string, idMakul []string) error {
	for _, id := range idMakul {
		sqlQuery := createPivotMakul(idSkPengangkatanDosen, id, userUpdate)
		//fmt.Printf("[DEBUG] createPivotMakul : %s", createPivotMakul)
		res, err := tx.Exec(sqlQuery)
		if err != nil {
			return fmt.Errorf("error executing Create Pivot Makul SK Pengangkatan Dosen query, %w", err)
		}
		affectedRows, err := res.RowsAffected()
		if err != nil {
			return fmt.Errorf("error executing get affected rows, %w", err)
		}
		if affectedRows == 0 {
			return sql.ErrNoRows
		}
	}

	return nil
}

func DeleteSKPengangkatanDosenByUUID(a app.App, uuid string) error {
	sqlQuery := deleteSKPengangkatanDosenByUUID()
	stmt, err := a.DB.Prepare(sqlQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(uuid)
	if err != nil {
		return err
	}

	return nil
}
func UpdateSKPengangkatanDosen(a app.App, skPegawai skPegawaiModel.SKPegawai) error {
	tx, err := a.DB.Begin()
	if err != nil {
		return fmt.Errorf("error beginning transaction, %w", err)
	}
	idSKPegawai, err := skPegawaiRepo.CreateSKPegawai(tx, skPegawai)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing update sk pegawai, %w", err)
	}

	skPegawai.SKPengangkatanDosen.PathSKDosen = PrepareFileUploadPathDosen(a, skPegawai.SKPengangkatanDosen.FileSKDosen, idSKPegawai)
	sqlQuery := updateSKPengangkatanDosenQuery(skPegawai.SKPengangkatanDosen, idSKPegawai)

	res, err := tx.Exec(sqlQuery)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing create sk pengangkatan Dosen query, %w", err)
	}
	affectedRows, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing get affected rows, %w", err)
	}
	if affectedRows == 0 {
		tx.Rollback()
		return sql.ErrNoRows
	}
	err = ExecMakulSkPengangkatanDosen(tx, idSKPegawai, skPegawai.UserUpdate, skPegawai.SKPengangkatanDosen.IDMataKuliah)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error update table pivot makul sk pengangkatan dosen, %w", err)
	}
	err = UploadFIleSKDosen(a, BucketSKDosen, skPegawai.SKPengangkatanDosen.PathSKDosen, skPegawai.SKPengangkatanDosen.FileSKDosen)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error upload file sk pengangkatan Dosen, %w", err)
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error commiting query, %w", err)
	}
	return nil
}
