package repo

import (
	"context"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/app/database"
	"testing"
)

func TestGetAllJenisSk(t *testing.T) {
	conn, err := database.Connect()
	if err != nil {
		t.Fatal(err)
	}
	db, err := database.InitGorm(conn, true)
	if err != nil {
		t.Fatal(err)
	}
	a := app.App{GormDB: db}
	jjs, err := GetAllJenisSk(a, context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(jjs) == 0 {
		t.Fatal("should not be empty")
	}
	for _, js := range jjs {
		fmt.Printf("[DEBUG] jenis sk: %+v\n", js)
	}
}

func TestGetJenisSk(t *testing.T) {
	conn, err := database.Connect()
	if err != nil {
		t.Fatal(err)
	}
	db, err := database.InitGorm(conn, true)
	if err != nil {
		t.Fatal(err)
	}
	a := app.App{GormDB: db}
	js, err := GetJenisSk(a, context.Background(), "1")
	if err != nil {
		t.Fatal(err)
	}
	if js == nil {
		t.Fatal("should not be empty")
	}
	fmt.Printf("[DEBUG] jenis sk: %+v\n", js)
}
