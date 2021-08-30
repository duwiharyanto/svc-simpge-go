package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/app/database"
	"testing"

	"github.com/openlyinc/pointy"
)

func TestGetPegawaiPendidikan(t *testing.T) {
	db, err := database.Connect()
	if err != nil {
		t.Fatal("failed connect to db:", err)
	}
	a := &app.App{DB: db}
	uuid := "db3b4cea-1437-11eb-a014-7eb0d4a3c7a0"
	jenjangPendidikan, err := GetPegawaiPendidikan(a, uuid, false)
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
	a := &app.App{DB: db}
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
// 	a := &app.App{DB: db}

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
	a := &app.App{DB: db}

	gormDB, err := database.InitGorm(a.DB, true)
	if err != nil {
		t.Fatal("failed connect to gorm db:", err)
	}
	ax := &app.App{GormDB: gormDB}

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
	a := &app.App{DB: db}

	gormDB, err := database.InitGorm(a.DB, true)
	if err != nil {
		t.Fatal("failed connect to gorm db:", err)
	}
	ax := &app.App{GormDB: gormDB}

	uuidPendidikanDiakui := "822c2cd9-4d0f-47ac-91b8-80f4b5d42444"
	uuidPendidikanTerakhir := "822c2cd9-4d0f-47ac-91b8-80f4b5d42444"
	var idPegawai uint64 = 1231231

	err = UpdatePendidikanPegawai(ax, context.Background(), uuidPendidikanDiakui, uuidPendidikanTerakhir, idPegawai)

	if err != nil {
		t.Fatal("failed update flag pendidikan pegawai:", err)
	}
}

func TestModelUpdate(t *testing.T) {
	db, err := database.Connect()
	if err != nil {
		t.Fatal("failed connect to db:", err)
	}

	type Pegawai struct {
		Id              uint64  `json:"id,omitempty"`
		IdGolonganDarah uint64  `json:"id_golongan_darah,omitempty"`
		NikKtp          *string `json:"-,omitempty"`
		UserUpdate      string  `json:"user_update,omitempty" gorm:"default:null"`
		UserInput       string  `json:"user_input,omitempty"`
	}

	gormDB, err := database.InitGorm(db, true)
	if err != nil {
		t.Fatal("failed connect to gorm db:", err)
	}

	// s := Pegawai{19296139857813084, "ahmad", "haris"}
	s := Pegawai{}
	s.Id = 99221016355366834
	s.IdGolonganDarah = 8745490293271215460
	t.Logf("s.Id: %#v\n", s.Id)
	s.UserUpdate = ""
	s.UserInput = "hariszv"

	var gp Pegawai
	err = gormDB.Find(&gp, "id = ?", s.Id).Error
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("gp: %#v", gp)
	// return
	bs, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(bs, &m)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(m)
	t.Log("id_golongan_darah:", m["id_golongan_darah"])
	// return

	// res := gormDB.Model(&s).Updates(Pegawai{})
	// res := gormDB.Save(&s)
	res := gormDB.Model(&Pegawai{Id: s.Id}).Updates(m)
	if res.Error != nil {
		t.Fatal("err save:", res.Error)
	}
	t.Log("res aff:", res.RowsAffected)

}

func TestPointer(t *testing.T) {
	foo := pointy.Uint64(8745490293271215460)
	fmt.Println("foo is a pointer to:", *foo)

	bar := pointy.String("point to me")
	fmt.Println("bar is a pointer to:", *bar)

	// get the value back out (new in v1.1.0)
	bar = nil
	barVal := pointy.StringValue(bar, "empty!")
	fmt.Println("bar's value is:", barVal)
}
