package repo

import (
	"context"
	"encoding/json"
	"svc-insani-go/app"
	"svc-insani-go/app/database"
	"testing"
)

func TestSearchPersonal(t *testing.T) {
	db, err := database.Connect()
	if err != nil {
		t.Fatal("failed connect to db:", err)
	}
	a := app.App{DB: db}

	gormDB, err := database.InitGorm(a.DB, true)
	if err != nil {
		t.Fatal("failed connect to gorm db:", err)
	}
	ax := app.App{GormDB: gormDB}

	nama := "Bobo"
	nikPegawai := "795110101"
	personal, err := SearchPersonal(ax, context.Background(), nama, nikPegawai)

	if err != nil {
		t.Fatal("failed get personal:", err)
	}

	j, _ := json.MarshalIndent(personal, "", "\t")

	t.Logf("personal:\n%s\n", j)
}
