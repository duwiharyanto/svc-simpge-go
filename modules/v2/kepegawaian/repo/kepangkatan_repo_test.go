package repo

import (
	"context"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/app/database"
	"testing"
)

func TestGetAllPangkatGolonganRuang(t *testing.T) {
	conn, err := database.Connect()
	if err != nil {
		t.Fatal(err)
	}
	db, err := database.InitGorm(conn)
	if err != nil {
		t.Fatal(err)
	}
	a := app.App{GormDB: db}
	pp := GetAllPangkatGolonganRuang(a, context.Background())
	if len(pp) == 0 {
		t.Fatal("should not be empty")
	}
	for _, p := range pp {
		fmt.Printf("[DEBUG] pgr: %+v\n", p)
	}
	fmt.Printf("[DEBUG] len pgr: %d\n", len(pp))
}
