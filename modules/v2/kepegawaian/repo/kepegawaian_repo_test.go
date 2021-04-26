package repo

import (
	"context"
	"svc-insani-go/app"
	"svc-insani-go/app/database"
	"testing"
)

func TestGetPegawai(t *testing.T) {
	conn, err := database.Connect()
	if err != nil {
		t.Fatal(err)
	}
	db, err := database.InitGorm(conn, true)
	if err != nil {
		t.Fatal(err)
	}
	a := app.App{GormDB: db}
	uuid := "d8c26983-1437-11eb-a014-7eb0d4a3c7a0x"
	pgw, err := GetPegawai(a, context.Background(), uuid)
	if err != nil {
		t.Fatal(err)
	}
	if pgw == nil {
		t.Fatal("should not be empty")
	}
	// fmt.Printf("[DEBUG] pegawai: %+v\n", pgw)
}
