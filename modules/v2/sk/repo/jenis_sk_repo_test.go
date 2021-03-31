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
	db, err := database.InitGorm(conn)
	if err != nil {
		t.Fatal(err)
	}
	a := app.App{GormDB: db}
	jjs := GetAllJenisSk(a, context.Background())
	if len(jjs) == 0 {
		t.Fatal("should not be empty")
	}
	for _, js := range jjs {
		fmt.Printf("[DEBUG] jenis sk: %+v\n", js)
	}
}
