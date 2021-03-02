package router

import (
	"database/sql"
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/app/database"
	jabatanFungsional "svc-insani-go/modules/v1/master-jabatan-fungsional/usecase"
	jenisNoRegis "svc-insani-go/modules/v1/master-jenis-nomor-registrasi/usecase"
	jenisPTT "svc-insani-go/modules/v1/master-jenis-pegawai-tidak-tetap/usecase"
	jenisPegawai "svc-insani-go/modules/v1/master-jenis-pegawai/usecase"
	kelompokPegawai "svc-insani-go/modules/v1/master-kelompok-pegawai/usecase"
	lokasiKerja "svc-insani-go/modules/v1/master-lokasi-kerja/usecase"
	pangkatPegawai "svc-insani-go/modules/v1/master-pangkat-golongan-pegawai/usecase"
	statusPegawai "svc-insani-go/modules/v1/master-status-pegawai/usecase"
	unitKerja "svc-insani-go/modules/v1/master-unit-kerja/usecase"
	pegawai "svc-insani-go/modules/v1/pegawai/usecase"

	"github.com/labstack/echo"
)

func InitRoute(a app.App, e *echo.Echo) {
	insaniGroupingPath := e.Group("/public/api/v1")
	// Route di bawah akan dikelola oleh handler
	insaniGroupingPath.GET("/pegawai", pegawai.HandleGetPegawai(a))
	insaniGroupingPath.GET("/pegawai-simpeg/:uuidPegawai/detail", pegawai.HandleGetSimpegPegawaiByUUID(a))
	insaniGroupingPath.PUT("/pegawai-simpeg/:uuidPegawai", pegawai.HandleUpdateSimpegPegawaiByUUID(a))

	// Data Master
	insaniGroupingPath.GET("/master-jenis-pegawai", jenisPegawai.HandleGetAllJenisPegawai(a))
	insaniGroupingPath.GET("/master-kelompok-pegawai", kelompokPegawai.HandleGetKelompokPegawai(a))
	insaniGroupingPath.GET("/master-status-pegawai", statusPegawai.HandleGetAllStatusPegawai(a))
	insaniGroupingPath.GET("/master-unit-kerja", unitKerja.HandleGetUnitKerja(a))
	insaniGroupingPath.GET("/master-jabatan-fungsional", jabatanFungsional.HandleGetJabatanFungsional(a))
	insaniGroupingPath.GET("/master-pangkat-golongan", pangkatPegawai.HandleGetPangkatGolonganPegawai(a))
	insaniGroupingPath.GET("/master-lokasi-kerja", lokasiKerja.HandleGetLokasiKerja(a))
	insaniGroupingPath.GET("/master-jenis-nomor-registrasi", jenisNoRegis.HandleGetJenisNoRegis(a))
	insaniGroupingPath.GET("/master-jenis-pegawai-tidak-tetap", jenisPTT.HandleGetJenisPTT(a))

}

func healthCheck(db *sql.DB) echo.HandlerFunc {
	h := func(c echo.Context) error {
		err := database.Healthz(c.Request().Context(), db)
		if err != nil {
			return fmt.Errorf("[ERROR] health check, %w", err)
			// fmt.Printf("[ERROR] health check, %s\n", err.Error())
			// return echo.NewHTTPError(http.StatusGatewayTimeout)
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "Layanan ok"})
	}
	return echo.HandlerFunc(h)
}

func dummyError() echo.HandlerFunc {
	h := func(c echo.Context) error {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan dibuat masalah"})
	}
	return echo.HandlerFunc(h)
}
