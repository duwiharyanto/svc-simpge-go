package repo

import (
	"context"
	"encoding/json"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/app/database"
	pegawai "svc-insani-go/modules/v1/pegawai/model"
	"svc-insani-go/modules/v1/personal/model"
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
		// uuid := "e421fc4e-4e8e-4dd3-8d0c-f41b7a38df5c"
		uuid := "1a6c87ca-e39d-11eb-8820-000c2977b907"
		v, err := GetPersonalByUuid(a, context.Background(), uuid)
		if err != nil {
			t.Fatal(err)
		}
		if v == nil {
			t.Fatal("should not be empty")
		}
		fmt.Printf("[DEBUG] per: %+v\n", v)
		fmt.Printf("[DEBUG] v.GolonganDarah == GolonganDarah{}: %t\n", v.GolonganDarah == model.GolonganDarah{})
		if v.Pegawai.PegawaiFungsional.StatusPegawaiAktif.IsActive() {
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
