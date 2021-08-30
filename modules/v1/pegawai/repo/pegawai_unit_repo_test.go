package repo

import (
	"svc-insani-go/app"
	"svc-insani-go/app/database"
	"testing"
)

func TestUnitKerjaPegawai(t *testing.T) {
	db, err := database.Connect()
	if err != nil {
		t.Fatal("failed connect to db:", err)
	}
	a := &app.App{DB: db}

	uuidPegawai := "cef383e3-6475-11eb-92df-506b8db8fcca"
	unit, err := GetUnitKerjaPegawai(a, uuidPegawai)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("unit pegawai: %+v\n", unit)
}
