package usecase

import (
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/app/database"
	"svc-insani-go/app/minio"
	"testing"
)

func TestNewPegawaiOra(t *testing.T) {
	db, err := database.Connect()
	if err != nil {
		t.Skip("failed connect db:", err)
	}

	gormDb, err := database.InitGorm(db, true)
	if err != nil {
		t.Fatal(err)
	}
	mc, err := minio.Connect()
	if err != nil {
		t.Fatal(err)
	}

	timeLocation := app.GetFixedTimeZone()
	a := &app.App{DB: db, GormDB: gormDb, TimeLocation: timeLocation, MinioClient: mc, MinioBucketName: "insani"}

	uuid := "872e2c5a-eb92-11eb-8820-000c2977b907"
	pegawaiDetail, err := PrepareGetSimpegPegawaiByUUID(a, uuid)
	if err != nil {
		t.Fatal(err)
	}
	pegawaiOra := newPegawaiOra(&pegawaiDetail)
	if pegawaiOra == nil {
		t.Fatal("Should not be nil")
	}
	fmt.Printf("[DEBUG] pgw ora: %+v\n", pegawaiOra)

}

func TestPrepareGetSimpegPegawaiByUUID(t *testing.T) {
	db, err := database.Connect()
	if err != nil {
		t.Skip("failed connect db:", err)
	}

	gormDb, err := database.InitGorm(db, true)
	if err != nil {
		t.Fatal(err)
	}
	mc, err := minio.Connect()
	if err != nil {
		t.Fatal(err)
	}

	timeLocation := app.GetFixedTimeZone()
	a := &app.App{DB: db, GormDB: gormDb, TimeLocation: timeLocation, MinioClient: mc, MinioBucketName: "insani"}

	uuid := "d8c77396-1437-11eb-a014-7eb0d4a3c7a0"
	res, err := PrepareGetSimpegPegawaiByUUID(a, uuid)
	if err != nil {
		t.Fatal(err)
	}
	if res.PegawaiPribadi != nil {
		fmt.Printf("[DEBUG] pgw pribadi: %+v\n\n", res.PegawaiPribadi)
	}
	if res.PegawaiYayasan != nil {
		fmt.Printf("[DEBUG] pgw yayasan: %+v\n\n", res.PegawaiYayasan)
	}
	if res.UnitKerjaPegawai != nil {
		fmt.Printf("[DEBUG] pgw unit kerja: %+v\n\n", res.UnitKerjaPegawai)
	}
	if res.PegawaiPNSPTT != nil {
		fmt.Printf("[DEBUG] pgw pns ptt: %+v\n\n", res.PegawaiPNSPTT)
	}
	if res.StatusAktif != nil {
		fmt.Printf("[DEBUG] pgw status aktif: %+v\n\n", res.StatusAktif)
	}
	if !res.JenjangPendidikan.IsEmpty() {
		fmt.Printf("[DEBUG] pgw pdd: %+v\n\n", res.JenjangPendidikan)
	}
	// j, _ := json.MarshalIndent(res, "", "\t")
	// fmt.Printf("[DEBUG] res: %s\n", j)
}
