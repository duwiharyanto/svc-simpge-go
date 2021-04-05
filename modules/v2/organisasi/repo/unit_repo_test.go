package repo

import (
	"context"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/app/database"
	"testing"
)

func TestGetAllUnit2(t *testing.T) {
	conn, err := database.Connect()
	if err != nil {
		t.Fatal(err)
	}
	db, err := database.InitGorm(conn)
	if err != nil {
		t.Fatal(err)
	}
	a := app.App{GormDB: db}
	uu, err := GetAllUnit2(a, context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(uu) == 0 {
		t.Fatal("should not be empty")
	}
	for _, u := range uu {
		fmt.Printf("[DEBUG] unit2: %+v\n", u)
	}
	fmt.Printf("[DEBUG] len unit2: %d\n", len(uu))
}
