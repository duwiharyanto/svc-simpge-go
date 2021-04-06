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
	BucketSKPegawai = "sk-pegawai"
	//BucketSKDosen   = "sk-dosen"
)

func CreateSKPengangkatanTendik(a app.App, skPegawai skPegawaiModel.SKPegawai) error {
	tx, err := a.DB.Begin()
	if err != nil {
		return fmt.Errorf("error beginning transaction, %w", err)
	}
	idSKUtama, err := skPegawaiRepo.CreateSKPegawai(tx, skPegawai)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing create sk pegawai, %w", err)
	}

	skPegawai.SKPengangkatanTendik.PathSKTendik = PrepareFileUploadPath(a, skPegawai.SKPengangkatanTendik.FileSKTendik, idSKUtama)
	sqlQuery := createSKPengangkatanTendikQuery(skPegawai.SKPengangkatanTendik, idSKUtama)
	// fmt.Printf("log query createSKPengangkatanQuery %+v \n", sqlQuery)
	res, err := tx.Exec(sqlQuery)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error executing create sk pengangkatan tendik query, %w", err)
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
	err = UploadFIleSK(a, BucketSKPegawai, skPegawai.SKPengangkatanTendik.PathSKTendik, skPegawai.SKPengangkatanTendik.FileSKTendik)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error upload file sk pengangkatan tendik, %w", err)
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error commiting query, %w", err)
	}
	return nil
}

func UploadFIleSK(a app.App, bucketName, filePath string, fileSKPengangkatanTendik *multipart.FileHeader) error {
	if fileSKPengangkatanTendik == nil {
		return nil
	}
	file, err := fileSKPengangkatanTendik.Open()
	if err != nil {
		return fmt.Errorf("error opening file, %w", err)
	}
	defer file.Close()

	fileContentType := fileSKPengangkatanTendik.Header.Get("Content-Type")
	err = a.MinioClient.Upload(bucketName, filePath, file, fileSKPengangkatanTendik.Size, fileContentType)
	if err != nil {
		return fmt.Errorf("error uploading file, %w", err)
	}
	return nil
}

func PrepareFileUploadPath(a app.App, fileSKPengangkatanTendik *multipart.FileHeader, idSKUtama string) string {
	if fileSKPengangkatanTendik == nil {
		return ""
	}
	file, err := fileSKPengangkatanTendik.Open()
	if err != nil {
		return ""
	}
	defer file.Close()
	unixTime := time.Now().In(a.TimeLocation).Unix()
	skPengangkatanFileContentType := fileSKPengangkatanTendik.Header.Get("Content-Type")
	skPengangkatanFileExtension := strings.Split(skPengangkatanFileContentType, "/")[1]
	skPengangkatanMinioFolderName := idSKUtama
	skPengangkatanMinioFileName := fmt.Sprintf("%d", unixTime)
	skPengangkatanMinioUploadPath := fmt.Sprintf("%s/%s.%s", skPengangkatanMinioFolderName, skPengangkatanMinioFileName, skPengangkatanFileExtension)
	return skPengangkatanMinioUploadPath
}

func GetDetailSKPengangkatanTendik(a app.App, UUIDSKPengangkatanTendik string) (*model.SKPengangkatanTendikDetail, error) {
	var sk model.SKPengangkatanTendikDetail
	sqlQuery := getDetailSKPengangkatanTendikQuery(UUIDSKPengangkatanTendik)
	// fmt.Printf("[DEBUG] q: %s\n", sqlQuery)
	err := a.DB.QueryRow(sqlQuery).Scan(
		&sk.NamaPegawai,
		&sk.NIKPegawai,
		&sk.TempatLahirPegawai,
		&sk.TanggalLahirPegawai,
		&sk.KDJenisSK,
		&sk.JenisSK,
		&sk.UUIDJenisSK,
		&sk.KDKelompokSKPengangkatan,
		&sk.KelompokSKPengangkatan,
		&sk.UUIDKelompokSKPengangkatan,
		&sk.NomorSK,
		&sk.KDUnitPengangkat,
		&sk.UnitPengangkat,
		&sk.UUIDUnitPengangkat,
		&sk.TMT,
		&sk.KDUnitPegawai,
		&sk.UnitPegawai,
		&sk.UUIDUnitPegawai,
		&sk.KDJabatanFungsional,
		&sk.JabatanFungsional,
		&sk.UUIDJabatanFungsional,
		&sk.KDPangkatGolonganPegawai,
		&sk.PangkatGolonganPegawai,
		&sk.UUIDPangkatGolonganPegawai,
		&sk.JabatanPenetap.NamaJenisJabatan,
		&sk.JabatanPenetap.NamaJenisUnit,
		&sk.JabatanPenetap.NamaUnit,
		&sk.JabatanPenetap.KdUnit,
		&sk.JabatanPenetap.Uuid,
		&sk.PejabatPenetap.Nama,
		&sk.PejabatPenetap.GelarDepan,
		&sk.PejabatPenetap.GelarBelakang,
		&sk.PejabatPenetap.Uuid,
		&sk.GajiPokok,
		&sk.MasaRilBulan,
		&sk.MasaRilTahun,
		&sk.MasaGajiBulan,
		&sk.MasaGajiTahun,
		&sk.MasaKerjaDiakuiBulan,
		&sk.MasaKerjaDiakuiTahun,
		&sk.KDStatusPengangkatan,
		&sk.StatusPengangkatan,
		&sk.UUIDStatusPengangkatan,
		&sk.KDJenisIjazah,
		&sk.JenisIjazah,
		&sk.UUIDJenisIjazah,
		&sk.TanggalDitetapkan,
		&sk.PathSKTendik,
		&sk.UUIDSKPengangkatanTendik,
	)
	// fmt.Printf("[DEBUG] sk pengangkatan tendik scanned: %+v\n", sk)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error querying get detail sk pengangkatan tendik. %s", err.Error())
	}
	sk.SetTTL()
	sk.SetTanggalSK()

	if sk.PathSKTendik != "" {
		sk.URLSKTendik, err = a.MinioClient.GetDownloadURL(a.MinioBucketName, sk.PathSKTendik, "")
		if err != nil {
			fmt.Printf("error getting file url, %s", err.Error())
		}
	}

	return &sk, nil
}

func DeleteSKPengangkatanTendikByUUID(a app.App, uuid string) error {
	sqlQuery := deleteSKPengangkatanTendikByUUID()
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

// func UpdateSKPengangkatanTendik(a app.App, skPegawai skPegawaiModel.SKPegawai) error {
// 	// if skPegawaiModel.PathSKTendik != nil {
// 	// 	skPegawaiModel.PathSKTendik = PrepareFileUploadPath(a, pendidikan.FileIjazah, pendidikan.UUIDPersonal)
// 	// }
// 	sqlQuery := updateSKPengangkatanTendikQuery(skPegawai)
// 	tx, err := a.DB.Begin()
// 	if err != nil {
// 		return fmt.Errorf("error beginning transaction, %s", err.Error())
// 	}
// 	_, err = tx.Exec(sqlQuery)
// 	if err != nil {
// 		return fmt.Errorf("error executing query, %s", err.Error())
// 	}

// 	err = UploadFilePendidikan(a, BucketIjazah, pendidikan.PathIjazah, pendidikan.FileIjazah)
// 	if err != nil {
// 		tx.Rollback()
// 		return fmt.Errorf("error upload file ktp, %s", err.Error())
// 	}

// 	// Commit untuk selesai transaksi
// 	// Semua proses (tambah kontak & tambah sosmed) yang dieksekusi akan disimpan di DB
// 	err = tx.Commit()
// 	if err != nil {
// 		tx.Rollback()
// 		return fmt.Errorf("error commiting query, %s", err.Error())
// 	}

// 	return nil
// }
