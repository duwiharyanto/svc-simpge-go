package repo

import (
	"context"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/app/database"
	"testing"
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
	ksk := GetAllKelompokSkPengangkatan(a, context.Background())
	if len(ksk) == 0 {
		t.Fatal("should not be empty")
	}
	for _, ks := range ksk {
		fmt.Printf("[DEBUG] kelompok sk pengangkatan: %+v\n", ks)
	}
	fmt.Printf("[DEBUG] len kelompok sk pengangkatan: %d\n", len(ksk))
}
