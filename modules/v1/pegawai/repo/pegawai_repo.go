package repo

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/pegawai/model"
	presensiHttp "svc-insani-go/modules/v1/presensi/http"
	presensiModel "svc-insani-go/modules/v1/presensi/model"
	"time"

	ptr "github.com/openlyinc/pointy"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	BucketIjazah    = "personal-ijazah"
	JenisFileIjazah = "ijazah"
)

func GetAllPegawai(a *app.App, req *model.PegawaiRequest) ([]model.Pegawai, error) {
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

func GetAllPegawaiPrivate(a *app.App, req *model.PegawaiPrivateRequest) ([]model.PegawaiPrivate, error) {
	sqlQuery := getListAllPegawaiPrivateQuery(req)
	// fmt.Println(sqlQuery)
	rows, err := a.DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying get pegawai, %w", err)
	}
	defer rows.Close()

	pp := []model.PegawaiPrivate{}
	for rows.Next() {
		var p model.PegawaiPrivate
		err := rows.Scan(
			&p.IdPegawai,
			&p.IdPersonal,
			&p.Nama,
			&p.NIK,
			&p.JenisPegawai,
			&p.IdJenisPegawai,
			&p.KdJenisPegawai,
			&p.KelompokPegawai,
			&p.IdKelompokPegawai,
			&p.KdKelompokPegawai,
			&p.IdKategoriKelompokPegawai,
			&p.KdKategoriKelompokPegawai,
			&p.Golongan,
			&p.IdGolongan,
			&p.KdGolongan,
			&p.GolonganNegara,
			&p.IdGolonganNegara,
			&p.KdGolonganNegara,
			&p.Ruang,
			&p.IdRuang,
			&p.KdRuang,
			&p.RuangNegara,
			&p.IdRuangNegara,
			&p.KdRuangNegara,
			&p.UnitKerja,
			&p.IdUnit,
			&p.KdUnit,
			&p.IndukKerja,
			&p.IdIndukKerja,
			&p.KdIndukKerja,
			&p.IdStatusPegawaiAktif,
			&p.StatusPegawaiAktif,
			&p.KdStatusPegawaiAktif,
			&p.StatusPegawai,
			&p.IdStatusPegawai,
			&p.KdStatusPegawai,
			&p.JenisKelamin,
			&p.JabatanFungsionalYayasan,
			&p.IdJabatanFungsionalYayasan,
			&p.KdJabatanFungsionalYayasan,
			&p.JabatanFungsionalNegara,
			&p.IdJabatanFungsionalNegara,
			&p.KdJabatanFungsionalNegara,
			&p.IdDetailProfesi,
			&p.DetailProfesi,
			&p.IdJenjangPendidikan,
			&p.KdJenjangPendidikan,
			&p.JenjangPendidikan,
			&p.TmtSkPertama,
			&p.MasaKerjaTahun,
			&p.MasaKerjaBulan,
			&p.JumlahKeluarga,
			&p.JumlahAnak,
			&p.Npwp,
			&p.IdStatusPernikahan,
			&p.KdStatusPernikahan,
			&p.StatusPernikahan,
			// &p.StatusPernikahanPtkp,
			&p.NikSuamiIstri,
			&p.NikKtp,
			// &p.JumlahTanggungan,
			// &p.JumlahTanggunganPtkp,
			&p.FlagKlaimTanggungan,
			&p.FlagPensiun,
			&p.FlagMeninggal,
			&p.FlagSuamiIstriSekantor,
			&p.IsFungsional,
			&p.IsStruktural,
		)
		if err != nil {
			return nil, fmt.Errorf("error scan pegawai row, %s", err.Error())
		}

		pp = append(pp, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error pegawai rows, %s", err.Error())
	}

	return pp, nil
}

func CountPegawai(a *app.App, req *model.PegawaiRequest) (int, error) {
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

func GetPegawaiByUUID(a *app.App, uuid string) (*model.Pegawai, error) {
	sqlQuery := getPegawaiByUUID(uuid)
	var pegawai model.Pegawai
	err := a.DB.QueryRow(sqlQuery).Scan(
		&pegawai.Id,
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

func GetPegawaiPribadi(a *app.App, uuid string) (*model.PegawaiPribadi, error) {
	sqlQuery := getPegawaiPribadiQuery(uuid)
	var pegawaiPribadi model.PegawaiPribadi

	err := a.DB.QueryRow(sqlQuery).Scan(
		&pegawaiPribadi.ID,
		&pegawaiPribadi.Nama,
		&pegawaiPribadi.NIK,
		&pegawaiPribadi.IdAgama,
		&pegawaiPribadi.KdAgama,
		&pegawaiPribadi.KdItemAgama,
		&pegawaiPribadi.IdGolonganDarah,
		&pegawaiPribadi.KdGolonganDarah,
		&pegawaiPribadi.GolonganDarah,
		&pegawaiPribadi.KdKelamin,
		&pegawaiPribadi.IdStatusPerkawinan,
		&pegawaiPribadi.KdNikah,
		&pegawaiPribadi.TempatLahir,
		&pegawaiPribadi.TanggalLahir,
		&pegawaiPribadi.FlagPensiun,
		&pegawaiPribadi.GelarDepan,
		&pegawaiPribadi.GelarBelakang,
		&pegawaiPribadi.NoKTP,
		&pegawaiPribadi.JenisPegawai,
		&pegawaiPribadi.KelompokPegawai,
		&pegawaiPribadi.UnitKerja,
		&pegawaiPribadi.KdStatusPendidikanMasuk,
		&pegawaiPribadi.KdJenisPendidikan,
		&pegawaiPribadi.UrlFileFoto,
		&pegawaiPribadi.UserInput,
		&pegawaiPribadi.UserUpdate,
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

func GetPegawaiByNIK(a *app.App, ctx context.Context, nik string) (*model.CreatePegawai, error) {

	var pegawaiAll model.CreatePegawai
	tx := a.GormDB.WithContext(ctx)

	res := tx.First(&pegawaiAll, "nik = ?", nik)
	if res.Error != nil {
		return nil, res.Error
	}

	return &pegawaiAll, nil
}

func GetOldPegawai(a *app.App, ctx context.Context, uuidPegawai string) (model.PegawaiUpdate, error) {
	var pegawai model.PegawaiUpdate

	db := a.GormDB.WithContext(ctx)
	res := db.Joins("PegawaiPNS").
		Joins("PegawaiFungsional").
		Find(&pegawai, "pegawai.uuid = ?", uuidPegawai)
	if res.Error != nil {
		return model.PegawaiUpdate{}, res.Error
	}

	return pegawai, nil
}

func UpdatePegawai(a *app.App, ctx context.Context, pegawaiUpdate model.PegawaiUpdate) error {
	db := a.GormDB.WithContext(ctx).Session(&gorm.Session{FullSaveAssociations: true})

	st := time.Now()
	res := db.Save(&pegawaiUpdate)
	if res.Error != nil {
		return res.Error
	}
	fmt.Printf("[DEBUG] update pegawai: %v ms\n", time.Now().Sub(st).Milliseconds())

	return nil
}

func CreatePegawai(a *app.App, ctx context.Context, pegawai model.PegawaiCreate) error {
	disableSyncPresence, _ := strconv.ParseBool(os.Getenv("DISABLE_SYNC_PRESENCE"))
	if !disableSyncPresence {
		user := &presensiModel.UserPresensi{
			Nip:             pegawai.Nik,
			KdJenisPegawai:  pegawai.KdJenisPegawai,
			KdUnitKerja:     pegawai.KdUnit2,
			KdLokasiKerja:   pegawai.LokasiKerja,
			Tmt:             ptr.StringValue(pegawai.PegawaiFungsional.TmtSkPertama, ""),
			KdJenisPresensi: pegawai.KdJenisPresensi,
			UserUpdate:      pegawai.UserUpdate,
		}
		err := presensiHttp.CreateUserPresensi(ctx, &http.Client{}, user)
		if err != nil {
			return fmt.Errorf("%w", fmt.Errorf("error create user presensi: %s", err.Error()))
		}
	}

	tx := a.GormDB.Session(&gorm.Session{
		Context: ctx,
		// FullSaveAssociations: true,
	})

	result := tx.Omit(clause.Associations).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "nik"}},
		UpdateAll: true}).
		Create(&pegawai)
	if result.Error != nil {
		return fmt.Errorf("error creating data simpeg : %w", result.Error)
	}

	result = tx.Find(&pegawai, "nik = ?", pegawai.Nik)
	if result.Error != nil {
		tx.Rollback()
		return fmt.Errorf("error find data simpeg nik : %w", result.Error)
	}

	pegawai.PegawaiFungsional.IdPegawai = pegawai.Id
	pegawai.PegawaiPNS.IdPegawai = pegawai.Id

	result = tx.Omit(clause.Associations).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id_pegawai"}},
		UpdateAll: true}).
		Create(&pegawai.PegawaiFungsional)
	if result.Error != nil {
		return fmt.Errorf("error creating data simpeg : %w", result.Error)
	}

	result = tx.Omit(clause.Associations).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id_pegawai"}},
		UpdateAll: true}).
		Create(&pegawai.PegawaiPNS)
	if result.Error != nil {
		return fmt.Errorf("error creating data simpeg : %w", result.Error)
	}

	return nil
}

func CheckNikPegawai(a *app.App, ctx context.Context, nik string) (*model.PegawaiCreate, bool, error) {

	var pegawaiOld model.PegawaiCreate
	var flagCheck bool

	db := a.GormDB.WithContext(ctx)
	res := db.Find(&pegawaiOld, "nik = ?", nik)
	if res.Error != nil {
		return nil, flagCheck, res.Error
	}

	if res.RowsAffected != 0 {
		flagCheck = true
		return &pegawaiOld, flagCheck, nil
	}

	return &pegawaiOld, flagCheck, nil
}

func GetPresignUrlFotoPegawai(a *app.App, pathFoto string) (string, error) {
	minioBucketNamePersonalFoto := "personal-foto"
	urlFileFoto, err := a.MinioClient.GetDownloadURL(minioBucketNamePersonalFoto, pathFoto, "")
	if err != nil {
		fmt.Printf("error get url file foto, %s", err.Error())
		urlFileFoto = ""
	}

	return urlFileFoto, nil
}

func GetPegawaiByNikPrivate(a *app.App, nik string) (*model.PegawaiByNik, error) {
	sqlQuery := getPegawaiByNik(nik)
	var pegawai model.PegawaiByNik
	err := a.DB.QueryRow(sqlQuery).Scan(
		&pegawai.Nama,
		&pegawai.GelarDepan,
		&pegawai.GelarBelakang,
		&pegawai.Nik,
		&pegawai.TempatLahir,
		&pegawai.TglLahir,
		&pegawai.JenisKelamin,
		&pegawai.KdPendidikanTerakhir,
		&pegawai.KdStatusPegawai,
		&pegawai.StatusPegawai,
		&pegawai.KdKelompokPegawai,
		&pegawai.KelompokPegawai,
		&pegawai.KdPangkatGolongan,
		&pegawai.Pangkat,
		&pegawai.KdGolongan,
		&pegawai.Golongan,
		&pegawai.KdRuang,
		&pegawai.TmtPangkatGolongan,
		// &pegawai.TmtPangkatGolonganIDN,
		&pegawai.KdJabatanFungsional,
		&pegawai.Fungsional,
		&pegawai.TmtJabatan,
		// &pegawai.TmtJabatanIDN,
		&pegawai.KdUnit1,
		&pegawai.Unit1,
		&pegawai.KdUnit2,
		&pegawai.Unit2,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error querying get pegawai by nik, %s", err.Error())
	}

	return &pegawai, nil
}
