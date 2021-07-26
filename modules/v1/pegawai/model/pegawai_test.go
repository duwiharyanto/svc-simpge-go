package model

import (
	"context"
	"errors"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/app/database"
	"testing"

	"gorm.io/gorm"
)

func TestPegawaiJoinPegawaiFungsional(t *testing.T) {
	conn, err := database.Connect()
	if err != nil {
		t.Fatal(err)
	}
	db, err := database.InitGorm(conn, true)
	if err != nil {
		t.Fatal(err)
	}
	a := &app.App{GormDB: db}
	t.Run("employee_exist", func(t *testing.T) {
		var id uint64 = 99221016355099875
		var pegawai Pegawai
		err := a.GormDB.WithContext(context.Background()).
			Preload("PegawaiFungsional.StatusPegawaiAktif").
			Joins("PegawaiFungsional").
			// Joins("PegawaiFungsional.StatusPegawaiAktif").
			Where(&Pegawai{Id: id}).
			First(&pegawai).
			Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			t.Fatal("should not be null")
		}

		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("[DEBUG] e: %+v\n", pegawai)
		t.SkipNow()

		var pf PegawaiFungsional
		err = a.GormDB.WithContext(context.Background()).
			Where(&PegawaiFungsional{IdPegawai: id}).
			First(&pf).
			Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			t.Fatal("pf should not be null")
		}

		if err != nil {
			t.Fatal(err)
		}

		pegawai.PegawaiFungsional = pf

		fmt.Printf("[DEBUG] p: %+v\n", pegawai)
	})

}
