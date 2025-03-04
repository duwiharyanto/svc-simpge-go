package helper

import (
	"fmt"
	"io"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
)

const (
	dateFormat     = "2006-01-02 15:04:05 Z0700 MST"
	DateTimeFormat = "2006-01-02 15:04:05"
	DateFormat     = "2006-01-02"
)

var indonesianMonths = [...]string{
	"Januari",
	"Februari",
	"Maret",
	"April",
	"Mei",
	"Juni",
	"Juli",
	"Agustus",
	"September",
	"Oktober",
	"November",
	"Desember",
}

var indonesianDays = [...]string{
	"Ahad",
	"Senin",
	"Selasa",
	"Rabu",
	"Kamis",
	"Jumat",
	"Sabtu",
}

type Date struct {
	Date string
}

type DateInIndonesia struct {
	FullDate []Date
}

func GetIndonesianMonth(format, date string) (int, string) {
	t, _ := time.Parse(format, date)
	month := t.Month()
	var indonesianMonth string
	if time.January <= month && month <= time.December {
		indonesianMonth = indonesianMonths[month-1]
	}
	return int(month), indonesianMonth
}

func GetIndonesianDate(format, date string) (int, string) {
	dpDate, _ := time.Parse(format, date)
	dpMonth := dpDate.Month()
	var IDMonth string
	if time.January <= dpMonth && dpMonth <= time.December {
		IDMonth = indonesianMonths[dpMonth-1]
	}
	indonesianDate := fmt.Sprintf("%d %s %d", dpDate.Day(), IDMonth, dpDate.Year())

	return dpDate.Day(), indonesianDate
}

func GetIndonesianMonthName(date string) string {
	dpDate, _ := time.Parse(DateFormat, date)
	dpMonth := dpDate.Month()
	var IDMonth string
	if time.January <= dpMonth && dpMonth <= time.December {
		IDMonth = indonesianMonths[dpMonth-1]
	}
	indonesianDate := fmt.Sprintf("%d %s %d", dpDate.Day(), IDMonth, dpDate.Year())

	return indonesianDate
}

func GetIndonesianDayName(date string) string {
	dpDate, _ := time.Parse(DateFormat, date)
	dpDay := dpDate.Weekday()
	var IdDay string
	IdDay = indonesianDays[dpDay]
	return IdDay
}

func GetDateByYear(year string) []Date {
	var (
		date  Date
		dates []Date
	)
	yearInt, _ := strconv.Atoi(year)
	kabisat := yearInt % 4
	t := time.Date(yearInt, time.January, 1, 0, 0, 0, 0, time.Local)
	if kabisat == 0 {
		for i := 1; i <= 366; i++ {
			date.Date = t.Format(DateFormat)
			dates = append(dates, date)
			t = t.AddDate(0, 0, 1)
		}
	}
	if kabisat != 0 {
		for i := 1; i <= 365; i++ {
			date.Date = t.Format(DateFormat)
			dates = append(dates, date)
			t = t.AddDate(0, 0, 1)
		}
	}

	return dates
}

func GetDateByDayName(theDate []Date, dayName string) []Date {
	var (
		date  Date
		dates []Date
	)
	for i := 0; i < len(theDate); i++ {
		day := GetIndonesianDayName(theDate[i].Date)
		if day == dayName {
			date.Date = theDate[i].Date
			dates = append(dates, date)
		}
	}
	return dates
}

func SetNULLIfToEmptyString(s string) string {
	if len(s) != 0 {
		return fmt.Sprintf("'%s'", s)
	}

	return "NULL"
}

func FlatQuery(q string) string {
	r := strings.NewReplacer("\n", " ", "\t", "")
	return r.Replace(q)
}

func ToSnakeCase(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, " ", "-"))
}

func IsYearFormatValid(year string) bool {
	_, err := time.Parse("2006", year)
	if err != nil {
		return false
	}
	return true
}

func IsDateFormatValid(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		return false
	}
	return true
}

func IsYearRangeValid(startYearTime, endYearTime time.Time) bool {
	diff := endYearTime.Sub(startYearTime)
	if diff < 0 {
		return false
	}
	return true
}

func FillFormDataWriter(w *multipart.Writer, m map[string]string) error {
	for k, v := range m {
		formField, err := w.CreateFormField(k)
		if err != nil {
			return fmt.Errorf("failed create field %s: %w", k, err)
		}
		_, err = io.Copy(formField, strings.NewReader(v))
		if err != nil {
			return fmt.Errorf("failed copy %s value: %w", k, err)
		}
	}

	return nil
}
