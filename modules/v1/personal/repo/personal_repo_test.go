package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/app/database"
	pegawai "svc-insani-go/modules/v1/pegawai/model"
	"testing"
)

func TestSearchPersonal(t *testing.T) {
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

	nama := "qwe"
	// nikPegawai := "795110101"
	personal, err := SearchPersonal(ax, context.Background(), nama)

	if err != nil {
		t.Fatal("failed get personal:", err)
	}

	j, _ := json.MarshalIndent(personal, "", "\t")

	t.Logf("personal:\n%s\n", j)
}

func TestPersonal(t *testing.T) {
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
		v, err := GetPersonalByUuid(a, context.Background(), uuid)
		if err != nil {
			t.Fatal(err)
		}
		if v != nil {
			fmt.Printf("[DEBUG] data: %+v\n", v)
			t.Fatal("should be empty")
		}
	})
	t.Run("personal_is_not_employee", func(t *testing.T) {
		uuid := "d57ec3c4-1593-4592-9abf-60a2c10e48a8"
		v, err := GetPersonalByUuid(a, context.Background(), uuid)
		if err != nil {
			t.Fatal(err)
		}
		if v == nil {
			t.Fatal("should not be empty")
		}
		// if v.Pegawai != nil {
		if (v.Pegawai == pegawai.Pegawai{}) {
			t.Fatal("should not be employee")
		}
	})
	t.Run("personal_is_employee", func(t *testing.T) {
		uuid := "c10b9568-697d-4d1f-be11-000cee0d55bc"
		v, err := GetPersonalByUuid(a, context.Background(), uuid)
		if err != nil {
			t.Fatal(err)
		}
		if v == nil {
			t.Fatal("should not be empty")
		}
		// if v.Pegawai == nil {
		if (v.Pegawai == pegawai.Pegawai{}) {
			t.Fatal("should be employee")
		}
		fmt.Printf("[DEBUG] data: %+v\n", v)
		fmt.Printf("[DEBUG] pgw: %+v\n", v.Pegawai)
	})
}
