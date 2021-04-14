package repo

import (
	"context"
	"encoding/json"
	"svc-insani-go/app"
	"svc-insani-go/app/database"
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

func TestGetPegawaiByUUID(t *testing.T) {
	db, err := database.Connect()
	if err != nil {
		t.Fatal("failed connect to db:", err)
	}
	a := app.App{DB: db}

	gormDB, err := database.InitGorm(a.DB, true)
	if err != nil {
		t.Fatal("failed connect to gorm db:", err)
	}
	ax := app.App{GormDB: gormDB}

	// pu := model.PegawaiUpdate{}
	// pu.Id = 819533014919819842
	// pu.Uuid = "e65ce262-1437-11eb-a014-7eb0d4a3c7a0"
	uuid := "d95551e5-1437-11eb-a014-7eb0d4a3c7a0"
	// pu.UserUpdate = "ahmad h"
	// pu.IdUnitKerja3 = 819533014920016302
	allPegawai, err := GetOldPegawai(ax, context.Background(), uuid)

	if err != nil {
		t.Fatal("failed get pegawai old:", err)
	}

	j, _ := json.MarshalIndent(allPegawai, "", "\t")

	t.Logf("pegawai:\n%s\n", j)

}

func TestUpdatePendidikanPegawai(t *testing.T) {
	db, err := database.Connect()
	if err != nil {
		t.Fatal("failed connect to db:", err)
	}
	a := app.App{DB: db}

	gormDB, err := database.InitGorm(a.DB, true)
	if err != nil {
		t.Fatal("failed connect to gorm db:", err)
	}
	ax := app.App{GormDB: gormDB}

	uuidPendidikanDiakui := "822c2cd9-4d0f-47ac-91b8-80f4b5d42444"
	uuidPendidikanTerakhir := "822c2cd9-4d0f-47ac-91b8-80f4b5d42444"
	idPegawai := 1231231

	err = UpdatePendidikanPegawai(ax, context.Background(), uuidPendidikanDiakui, uuidPendidikanTerakhir, idPegawai)

	if err != nil {
		t.Fatal("failed update flag pendidikan pegawai:", err)
	}
}

// func TestUpdatePegawai(t *testing.T) {
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

// 	pu := &model.PegawaiUpdate{}
// 	// pu.Id = "819533014919819842"
// 	// pu.Uuid = "e65ce262-1437-11eb-a014-7eb0d4a3c7a0"
// 	pu.Uuid = "dc8d3a7e-1437-11eb-a014-7eb0d4a3c7a0"
// 	pu.UserUpdate = "ahmad h"
// 	// pu.IdUnitKerja3 = "819533014920016302"
// 	err = UpdatePegawaix(ax, context.Background(), pu)
// 	if err != nil {
// 		t.Fatal("failed update pegawai:", err)
// 	}

// 	// j, _ := json.MarshalIndent(allPegawai, "", "\t")
// 	// t.Logf("pegawai:\n%s\n", j)
// 	// 81929613986368762599
// 	// 18446744073709551615

// }
// func TestPegawaiOld(t *testing.T) {
// 	uuid := "e5762619-1437-11eb-a014-7eb0d4a3c7a0"
// 	pegawai, err := GetPegawaiByUUID(a, uuid)
// 	if err != nil {
// 		t.Fatal("failed get pegawai:", err)
// 	}
// 	j, _ := json.MarshalIndent(pegawai, "", "\t")

// 	t.Logf("pegawai:\n%s\n", j)
// }
