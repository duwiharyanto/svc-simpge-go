package repo

import (
	"context"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/app/database"
	"testing"
)

func TestGetAllSkPegawai(t *testing.T) {
	conn, err := database.Connect()
	if err != nil {
		t.Fatal(err)
	}
	db, err := database.InitGorm(conn, true)
	if err != nil {
		t.Fatal(err)
	}
	a := app.App{GormDB: db}
	ssp := GetAllSkPegawai(a, context.Background())
	if len(ssp) == 0 {
		t.Fatal("should not be empty")
	}
	for _, sp := range ssp {
		fmt.Printf("[DEBUG] sk pegawai: %+v\n", sp)
	}
}

func TestGetSkPegawaiByUUID(t *testing.T) {
	conn, err := database.Connect()
	if err != nil {
		t.Fatal(err)
	}
	db, err := database.InitGorm(conn, true)
	if err != nil {
		t.Fatal(err)
	}
	a := app.App{GormDB: db}
	uuid := "79fb7ae9-d436-4c2e-87ed-cd68c78c5b9e"
	skp := GetSkPegawai(a, context.Background(), uuid)
	if skp == nil {
		t.Fatal("should not be empty")
	}
	fmt.Printf("[DEBUG] sk pegawai: %+v\n", skp)
}

func TestGetAllJenisIjazah(t *testing.T) {
	conn, err := database.Connect()
	if err != nil {
		t.Fatal(err)
	}
	db, err := database.InitGorm(conn, true)
	if err != nil {
		t.Fatal(err)
	}
	a := app.App{GormDB: db}
	jj, err := GetAllJenisIjazah(a, context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(jj) == 0 {
		t.Fatal("should not be empty")
	}
	for _, j := range jj {
		fmt.Printf("[DEBUG] jenis ijazah: %+v\n", j)
	}
	fmt.Printf("[DEBUG] len jenis ijazah: %d\n", len(jj))
}
