package repo

import (
	"fmt"
	"math"
	"strconv"
	"svc-insani-go/app"
	"svc-insani-go/app/database"
	"testing"
	"time"
)

func TestSetMasaKerjaTotal(t *testing.T) {
	db, err := database.Connect()
	if err != nil {
		t.Fatal("failed connect to db:", err)
	}
	a := &app.App{DB: db}

	uuidPegawai := "cef383e3-6475-11eb-92df-506b8db8fcca"
	kepegawaianYayasan, err := GetKepegawaianYayasan(a, uuidPegawai)
	if err != nil {
		t.Fatal(err)
	}

	unit, err := GetUnitKerjaPegawai(a, uuidPegawai)
	if err != nil {
		t.Fatal(err)
	}

	// kepegawaianYayasan.SetMasaKerjaTotal(unit)
	kepegawaianYayasan.SetMasaKerjaTotal(unit)
	fmt.Printf("\n\npgw: %+v\n\n", kepegawaianYayasan)
	fmt.Printf("pgw.MasaKerjaBawaanTahun: %+v\n", kepegawaianYayasan.MasaKerjaBawaanTahun)
	fmt.Printf("pgw.MasaKerjaBawaanBulan: %+v\n", kepegawaianYayasan.MasaKerjaBawaanBulan)
	fmt.Printf("pgw.TmtSkPertama: %+v\n", unit.TmtSkPertama)
	fmt.Printf("pgw.MasaKerjaTotalTahun: %+v\n", kepegawaianYayasan.MasaKerjaTotalTahun)
	fmt.Printf("pgw.MasaKerjaTotalBulan: %+v\n", kepegawaianYayasan.MasaKerjaTotalBulan)
}

func TestConvertStr(t *testing.T) {
	masaKerjaBawaanTahun := "1"
	masaKerjaBawaanTahun = ""
	masaKerjaBawaanTahunInt, _ := strconv.Atoi(masaKerjaBawaanTahun)

	masaKerjaBawaanBulan := ""
	masaKerjaBawaanBulanInt, _ := strconv.Atoi(masaKerjaBawaanBulan)

	// tmtSkPertama := "2009-10-01"
	tmtSkPertama := "1996-06-15"
	// tmtSkPertama = "2019-10-01"
	// tmtSkPertama = ""

	tmtSkPertamaTime, err := time.Parse("2006-01-02", tmtSkPertama)
	var tmtSkPertamaDuration time.Duration
	if err == nil {
		// t.Fatalf("tmtSkPertamaTime: %+v\nerr: %+v\n", tmtSkPertamaTime, err.Error())
		tmtSkPertamaDuration = time.Now().Sub(tmtSkPertamaTime)
	}
	// tmtSkPertamaDurationDays := int(tmtSkPertamaDuration.Hours() / 24)
	// tmtSkPertamaDurationYears := int(tmtSkPertamaDurationDays / 365)
	// tmtSkPertamaDurationRealMonths := int(tmtSkPertamaDurationDays / 30)
	// tmtSkPertamaDurationMonths := int(tmtSkPertamaDurationDays / 30 % 12)

	tmtSkPertamaDurationDays := tmtSkPertamaDuration.Hours() / 24
	tmtSkPertamaDurationYears := int(tmtSkPertamaDurationDays / 365)
	tmtSkPertamaDurationRealMonths := int(tmtSkPertamaDurationDays / 365 * 12)
	tmtSkPertamaDurationMonths := tmtSkPertamaDurationRealMonths % 12

	masaKerjaTotalRealBulan := ((masaKerjaBawaanTahunInt * 12) + masaKerjaBawaanBulanInt) + tmtSkPertamaDurationRealMonths
	masaKerjaTotalTahun := masaKerjaTotalRealBulan / 12
	masaKerjaTotalBulan := masaKerjaTotalRealBulan % 12
	// masaKerjaTotalTahunInt := masaKerjaBawaanTahunInt

	fmt.Printf("[DEBUG] mkbw th: %v\n", masaKerjaBawaanTahunInt)
	fmt.Printf("[DEBUG] mkbw bl: %v\n", masaKerjaBawaanBulanInt)

	fmt.Printf("\n[DEBUG] dur: %v\n", tmtSkPertamaDurationDays)
	fmt.Printf("[DEBUG] dur th: %v\n", tmtSkPertamaDurationYears)
	fmt.Printf("[DEBUG] dur real bl: %v\n", tmtSkPertamaDurationRealMonths)
	fmt.Printf("[DEBUG] dur bl: %v\n\n", tmtSkPertamaDurationMonths)

	fmt.Printf("[DEBUG] mkt real bl: %v\n", masaKerjaTotalRealBulan)
	fmt.Printf("[DEBUG] mkt th: %v\n", masaKerjaTotalTahun)
	fmt.Printf("[DEBUG] mkt bl: %v\n", masaKerjaTotalBulan)
}

func TestRound(t *testing.T) {
	f := 1.9
	res := math.Round(f)
	t.Log("res:", int(res))
	res = math.RoundToEven(f)
	t.Log("res:", int(res))
	t.Log("res:", int(f))
}
