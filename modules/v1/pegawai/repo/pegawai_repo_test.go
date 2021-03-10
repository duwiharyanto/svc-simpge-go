package repo

import (
	"encoding/json"
	"svc-insani-go/app"
	"svc-insani-go/app/database"
	"testing"
)

func TestGetPegawaiPendidikan(t *testing.T) {
	db, err := database.Connect()
	if err != nil {
		t.Fatal("failed connect to db:", err)
	}
	a := app.App{DB: db}
	uuid := "db3b4cea-1437-11eb-a014-7eb0d4a3c7a0"
	jenjangPendidikan, err := GetPegawaiPendidikan(a, uuid)
	if err != nil {
		t.Fatal("failed get jenjang pendidikan:", err)
	}
	j, _ := json.MarshalIndent(jenjangPendidikan, "", "\t")

	t.Logf("jenjang:\n%s\n", j)
}

func TestGetPegawaiFilePendidikan(t *testing.T) {
	db, err := database.Connect()
	if err != nil {
		t.Fatal("failed connect to db:", err)
	}
	a := app.App{DB: db}
	idList := []string{"3887303357", "1101348424"}
	filePendidikan, err := GetPegawaiFilePendidikan(a, idList...)
	if err != nil {
		t.Fatal("failed get jenjang pendidikan:", err)
	}
	j, _ := json.MarshalIndent(filePendidikan, "", "\t")

	t.Logf("berkas:\n%s\n", j)
}
