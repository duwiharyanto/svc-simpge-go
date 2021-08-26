package repo

import (
	"context"
	"database/sql"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/pegawai/model"
	"time"

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

func UpdatePegawai(a *app.App, ctx context.Context, pegawaiUpdate model.PegawaiUpdate) error {
	db := a.GormDB.WithContext(ctx)

	st := time.Now()
	res := db.Save(&pegawaiUpdate)
	if res.Error != nil {
		return res.Error
	}
	fmt.Printf("[DEBUG] update pegawai: %v ms\n", time.Now().Sub(st).Milliseconds())

	st = time.Now()
	res = db.Save(&pegawaiUpdate.PegawaiPNS)
	if res.Error != nil {
		return res.Error
	}
	fmt.Printf("[DEBUG] update pegawai pns: %v ms\n", time.Now().Sub(st).Milliseconds())

	st = time.Now()
	res = db.Save(&pegawaiUpdate.PegawaiFungsional)
	if res.Error != nil {
		return res.Error
	}
	fmt.Printf("[DEBUG] update pegawai fung: %v ms\n", time.Now().Sub(st).Milliseconds())

	return nil
}

func CreatePegawai(a *app.App, ctx context.Context, pegawaiCreate model.PegawaiCreate) error {
	tx := a.GormDB.Session(&gorm.Session{
		Context: ctx,
		// FullSaveAssociations: true,
	})

	result := tx.Omit(clause.Associations).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "nik"}},
		UpdateAll: true}).
		Create(&pegawaiCreate)
	if result.Error != nil {
		return fmt.Errorf("error creating data simpeg : %w", result.Error)
	}

	result = tx.Find(&pegawaiCreate, "nik = ?", pegawaiCreate.Nik)
	if result.Error != nil {
		tx.Rollback()
		return fmt.Errorf("error find data simpeg nik : %w", result.Error)
	}

	pegawaiCreate.PegawaiFungsional.IdPegawai = pegawaiCreate.Id
	pegawaiCreate.PegawaiPNS.IdPegawai = pegawaiCreate.Id

	result = tx.Omit(clause.Associations).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id_pegawai"}},
		UpdateAll: true}).
		Create(&pegawaiCreate.PegawaiFungsional)
	if result.Error != nil {
		return fmt.Errorf("error creating data simpeg : %w", result.Error)
	}

	result = tx.Omit(clause.Associations).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id_pegawai"}},
		UpdateAll: true}).
		Create(&pegawaiCreate.PegawaiPNS)
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
