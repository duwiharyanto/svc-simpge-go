package repo

import (
	"context"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/app/database"
	"testing"
)

func TestKelompok(t *testing.T) {
	conn, err := database.Connect()
	if err != nil {
		t.Fatal(err)
	}
	db, err := database.InitGorm(conn, true)
	if err != nil {
		t.Fatal(err)
	}
	a := &app.App{GormDB: db}
	t.Run("is_empty", func(t *testing.T) {
		uuid := ""
		v, err := GetKelompokPegawaiByUUID(a, context.Background(), uuid)
		if err != nil {
			t.Fatal(err)
		}
		if v != nil {
			fmt.Printf("[DEBUG] ekk: %+v\n", v)
			t.Fatal("should be empty")
		}
	})
	t.Run("is_exist", func(t *testing.T) {
		uuid := "c2cff014-5156-11eb-abec-000c29d8230c"
		v, err := GetKelompokPegawaiByUUID(a, context.Background(), uuid)
		if err != nil {
			t.Fatal(err)
		}
		if v == nil {
			t.Fatal("should not be empty")
		}
		fmt.Printf("[DEBUG] result: %+v\n", v)
	})
}
