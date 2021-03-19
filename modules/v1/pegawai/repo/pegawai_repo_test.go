package repo

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/app/database"
	"svc-insani-go/modules/v1/pegawai/model"
	"testing"
)

func TestGetPegawaiPendidikan(t *testing.T) {
	db, err := database.Connect()
	if err != nil {
		t.Fatal("failed connect to db:", err)
	}
	a := app.App{DB: db}
	uuid := "db3b4cea-1437-11eb-a014-7eb0d4a3c7a0"
	jenjangPendidikan, err := GetPegawaiPendidikan(a, uuid)
	if err != nil {
		t.Fatal("failed get jenjang pendidikan:", err)
	}
	j, _ := json.MarshalIndent(jenjangPendidikan, "", "\t")

	t.Logf("jenjang:\n%s\n", j)
}

func TestGetPegawaiFilePendidikan(t *testing.T) {
	db, err := database.Connect()
	if err != nil {
		t.Fatal("failed connect to db:", err)
	}
	a := app.App{DB: db}
	idList := []string{"3887303357", "1101348424"}
	filePendidikan, err := GetPegawaiFilePendidikan(a, idList...)
	if err != nil {
		t.Fatal("failed get jenjang pendidikan:", err)
	}
	j, _ := json.MarshalIndent(filePendidikan, "", "\t")

	t.Logf("berkas:\n%s\n", j)
}

// func TestGetAllPegawai(t *testing.T) {
// 	db, err := database.Connect()
// 	if err != nil {
// 		t.Fatal("failed connect to db:", err)
// 	}
// 	a := app.App{DB: db}

// 	gormDB, err := database.InitGorm(a.DB)
// 	if err != nil {
// 		t.Fatal("failed connect to gorm db:", err)
// 	}
// 	ax := app.App{GormDB: gormDB}
// 	allPegawai, err := GetAllPegawaix(ax)

// 	if err != nil {
// 		t.Fatal("failed get all pegawai:", err)
// 	}

// 	j, _ := json.MarshalIndent(allPegawai, "", "\t")

// 	t.Logf("pegawai:\n%s\n", j)

// }

func TestUpdatePegawai(t *testing.T) {
	db, err := database.Connect()
	if err != nil {
		t.Fatal("failed connect to db:", err)
	}
	a := app.App{DB: db}

	gormDB, err := database.InitGorm(a.DB)
	if err != nil {
		t.Fatal("failed connect to gorm db:", err)
	}
	ax := app.App{GormDB: gormDB}

	pu := &model.PegawaiUpdate{}
	// pu.Id = "819533014919819842"
	pu.Uuid = "e65ce262-1437-11eb-a014-7eb0d4a3c7a0"
	// pu.UserUpdate = "ahmad"
	err = UpdatePegawaix(ax, context.Background(), pu)
	if err != nil {
		t.Fatal("failed update pegawai:", err)
	}

	// j, _ := json.MarshalIndent(allPegawai, "", "\t")
	// t.Logf("pegawai:\n%s\n", j)
	// 81929613986368762599
	// 18446744073709551615

}

func TestGetAllPegawai(t *testing.T) {
	db, err := database.Connect()
	if err != nil {
		t.Fatal("failed connect to db:", err)
	}
	a := app.App{DB: db}

	gormDB, err := database.InitGorm(a.DB)
	if err != nil {
		t.Fatal("failed connect to gorm db:", err)
	}
	ax := app.App{GormDB: gormDB}

	allPegawai, err := GetAllPegawaix(ax, context.Background(), 10, 0)

	if err != nil {
		t.Fatal("failed get all pegawai:", err)
	}

	j, _ := json.MarshalIndent(allPegawai, "", "\t")

	t.Logf("pegawai:\n%s\n", j)

}

type Makul struct {
	Id     int64
	IdTest int64
	Nama   string
	UUID   string
}

func getMakulByUUID(a app.App, ctx context.Context, uuid string) (*Makul, error) {
	q := fmt.Sprintf(`SELECT id, id_test, nama_makul, uuid FROM makul WHERE uuid = %q`, uuid)
	var mk Makul
	err := a.DB.QueryRowContext(ctx, q).Scan(
		&mk.Id,
		&mk.IdTest,
		&mk.Nama,
		&mk.UUID,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("error querying and scan get makul: %w", err)
	}
	return &mk, nil
}

func TestMakul(t *testing.T) {
	db, err := database.Connect()
	if err != nil {
		t.Fatal("failed connect to db:", err)
	}
	a := app.App{DB: db}

	uuidMk := "38174566-2ecb-11eb-a014-7eb0d4a3c7a0"
	ctx := context.Background()

	t.Run("get_by_uuid", func(t *testing.T) {
		mk, err := getMakulByUUID(a, ctx, uuidMk)
		if err != nil {
			t.Fatal("failed get makul by uuid:", err)
		}
		if mk == nil {
			t.Fatal("mk should not be nil")
		}
		t.Log("mk by uuid:", mk)

	})
}
