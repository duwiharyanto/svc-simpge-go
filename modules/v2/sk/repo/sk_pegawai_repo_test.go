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
	db, err := database.InitGorm(conn)
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
