package repo

import (
	"context"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/app/database"
	"testing"
	"time"
)

func TestGetAllKelompokSkPengangkatan(t *testing.T) {
	conn, err := database.Connect()
	if err != nil {
		t.Fatal(err)
	}
	db, err := database.InitGorm(conn)
	if err != nil {
		t.Fatal(err)
	}
	a := app.App{GormDB: db}
	ksk, err := GetAllKelompokSkPengangkatan(a, context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(ksk) == 0 {
		t.Fatal("should not be empty")
	}
	for _, ks := range ksk {
		fmt.Printf("[DEBUG] kelompok sk pengangkatan: %+v\n", ks)
	}
	fmt.Printf("[DEBUG] len kelompok sk pengangkatan: %d\n", len(ksk))
}

func TestGetKelompokSkPengangkatan(t *testing.T) {
	conn, err := database.Connect()
	if err != nil {
		t.Fatal(err)
	}
	db, err := database.InitGorm(conn)
	if err != nil {
		t.Fatal(err)
	}
	a := app.App{GormDB: db}
	// uuid := "zzz"
	uuid := "f9a2a6e4-ec0a-11ea-8c77-7eb0d4a3c7a0"
	ksk, err := GetKelompokSkPengangkatan(a, context.Background(), uuid)
	if err != nil {
		t.Fatal(err)
	}
	if ksk == nil {
		t.Fatal("should not be empty")
	}
	fmt.Printf("[DEBUG] kelompok sk pengangkatan: %+v\n", ksk)
}

func TestUpdateSkPengangkatanTendik(t *testing.T) {
	conn, err := database.Connect()
	if err != nil {
		t.Fatal(err)
	}
	db, err := database.InitGorm(conn)
	if err != nil {
		t.Fatal(err)
	}
	a := app.App{GormDB: db}

	uuid := "dfef3d4d-2ffe-11eb-a014-7eb0d4a3c7a0"
	ctx := context.Background()
	skpt, err := GetSkPengangkatanTendik(a, ctx, uuid)
	if err != nil {
		t.Fatal(err)
	}
	if skpt == nil {
		t.Fatal("should not be empty")
	}
	fmt.Printf("[DEBUG] sk pengangkatan tendik: %+v\n", skpt)
	skpt.SkPegawai.TentangSk = "ini tentang v"
	skpt.SkPegawai.UserUpdate = "ahmad v"
	skpt.GajiPokok = 1500014
	skpt.UserUpdate = "ahmad v"
	skpt.JabatanFungsional.Id = 819533014920414656
	skpt.JenisIjazah.Id = 819296139862351087
	skpt.KelompokSkPengangkatan.Id = 819296139862340422
	skpt.PangkatGolonganRuang.Id = 819296139864116955
	// skpt.StatusPengangkatan.Id = 819296139862788372
	skpt.StatusPengangkatan.Id = 0
	skpt.UnitPengangkat.Id = 819533014920015782
	// skpt.IdUnitPengangkat = 819533014920015761
	skpt.UnitKerja.Id = 819533014920015785
	// skpt.IdUnitPegawai = 819533014920015775
	fmt.Printf("[DEBUG] skpt before exec update: %+v\n", skpt)
	ctx, _ = context.WithTimeout(ctx, time.Millisecond*5000)
	err = UpdateSkPengangkatanTendik(a, ctx, skpt)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetSkPengangkatanTendik(t *testing.T) {
	conn, err := database.Connect()
	if err != nil {
		t.Fatal(err)
	}
	db, err := database.InitGorm(conn)
	if err != nil {
		t.Fatal(err)
	}
	a := app.App{GormDB: db}
	uuid := "dfef3d4d-2ffe-11eb-a014-7eb0d4a3c7a0"
	// uuid := "6215c058-1e3d-11eb-a014-7eb0d4a3c7a0"
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*1000)
	defer cancel()
	skpt, err := GetSkPengangkatanTendik(a, ctx, uuid)
	if err != nil {
		t.Fatal(err)
	}
	if skpt == nil {
		t.Fatal("should not be empty")
	}
	fmt.Printf("[DEBUG] sk pengangkatan tendik: %+v\n", skpt)
}
