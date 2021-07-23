package repo

import (
	"context"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/app/database"
	"testing"
)

func TestGetAllJabatanStruktural(t *testing.T) {
	conn, err := database.Connect()
	if err != nil {
		t.Fatal(err)
	}
	db, err := database.InitGorm(conn, true)
	if err != nil {
		t.Fatal(err)
	}
	a := &app.App{GormDB: db}
	jj, err := GetAllJabatanStruktural(a, context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(jj) == 0 {
		t.Fatal("should not be empty")
	}
	for _, js := range jj {
		fmt.Printf("[DEBUG] JabatanStruktural: %+v\n", js)
	}
	fmt.Printf("[DEBUG] len JabatanStruktural: %d\n", len(jj))
}

func TestGetPejabatStruktural(t *testing.T) {
	conn, err := database.Connect()
	if err != nil {
		t.Fatal(err)
	}
	db, err := database.InitGorm(conn, true)
	if err != nil {
		t.Fatal(err)
	}
	a := &app.App{GormDB: db}
	uuidJabatanStruktural := "6c1d68ef-9461-11eb-b06a-000c2977b907"
	jj, err := GetPejabatStruktural(a, context.Background(), uuidJabatanStruktural)
	if err != nil {
		t.Fatal(err)
	}
	if len(jj) == 0 {
		t.Fatal("should not be empty")
	}
	for _, js := range jj {
		fmt.Printf("[DEBUG] Pejabat Struktural: %+v\n", js)
	}
	fmt.Printf("[DEBUG] len Pejabat Struktural: %d\n", len(jj))
}
