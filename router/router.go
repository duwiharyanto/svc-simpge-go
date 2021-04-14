package router

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"svc-insani-go/app"
	"svc-insani-go/app/database"
	indukKerja "svc-insani-go/modules/v1/master-induk-kerja/usecase"
	jabatanFungsional "svc-insani-go/modules/v1/master-jabatan-fungsional/usecase"
	jenisNoRegis "svc-insani-go/modules/v1/master-jenis-nomor-registrasi/usecase"
	jenisPTT "svc-insani-go/modules/v1/master-jenis-pegawai-tidak-tetap/usecase"
	jenisPegawai "svc-insani-go/modules/v1/master-jenis-pegawai/usecase"
	kelompokPegawai "svc-insani-go/modules/v1/master-kelompok-pegawai/usecase"
	lokasiKerja "svc-insani-go/modules/v1/master-lokasi-kerja/usecase"
	pangkatPegawai "svc-insani-go/modules/v1/master-pangkat-golongan-pegawai/usecase"
	statusPegawaiAktif "svc-insani-go/modules/v1/master-status-pegawai-aktif/usecase"
	statusPegawai "svc-insani-go/modules/v1/master-status-pegawai/usecase"
	unitKerja "svc-insani-go/modules/v1/master-unit-kerja/usecase"
	pegawai "svc-insani-go/modules/v1/pegawai/usecase"
	personal "svc-insani-go/modules/v1/personal/usecase"
	skPengangkatan "svc-insani-go/modules/v1/sk-pengangkatan/usecase"
	sk "svc-insani-go/modules/v1/sk/usecase"
	organisasiV2 "svc-insani-go/modules/v2/organisasi/usecase"
	skV2 "svc-insani-go/modules/v2/sk/usecase"

	"github.com/labstack/echo"
)

func InitRoute(a app.App, appCtx context.Context, e *echo.Echo, slackErrChan chan error) {
	insaniGroupingPath := e.Group("/public/api/v1")
	// Route di bawah akan dikelola oleh handler
	insaniGroupingPath.GET("/pegawai", pegawai.HandleGetPegawai(a))
	insaniGroupingPath.GET("/pegawai-simpeg/:uuidPegawai/detail", pegawai.HandleGetSimpegPegawaiByUUID(a))
	insaniGroupingPath.PUT("/pegawai-simpeg/:uuidPegawai", pegawai.HandleUpdateSimpegPegawaiByUUID(a))
	// insaniGroupingPath.POST("/pegawai-simpeg/:uuidPegawai", pegawai.HandleUpdateSimpegPegawaiByUUID(a))
	// 	FE
	// data
	// uuid_personal
	// nik_pegawai
	// (nik_ktp)
	// data-data UI SIMPEG
	// kirim ke POST /pegawai-simpeg
	// BE
	// terima request dari POST /pegawai-simpeg
	// persiapan tambah pegawai baru
	// ambil data pribadi berdasarkan uuid_personal dari personal_data_pribadi
	// buat id pegawai
	// simpan
	// data pribadi + data pegawai simpeg ke tabel pegawai
	// data pns jika dia pns ke tabel pegawai_pns
	// dll sesuai data ketika update pegawai

	// Data Master
	insaniGroupingPath.GET("/jabatan-struktural", organisasiV2.HandleGetAllJabatanStruktural(a))
	insaniGroupingPath.GET("/jenis-ijazah", skV2.HandleGetAllJenisIjazah(a))
	insaniGroupingPath.GET("/jenis-sk", skV2.HandleGetAllJenisSk(a))
	insaniGroupingPath.GET("/kelompok-sk-pengangkatan", skV2.HandleGetAllKelompokSkPengangkatan(a))
	insaniGroupingPath.GET("/mata-kuliah", skV2.HandleGetAllMataKuliah(a))
	insaniGroupingPath.GET("/pejabat-struktural", organisasiV2.HandleGetPejabatStruktural(a))
	insaniGroupingPath.GET("/unit2", organisasiV2.HandleGetAllUnit2(a))

	insaniGroupingPath.GET("/master-ijazah-pegawai", sk.HandleGetAllJenisIjazah(a))
	insaniGroupingPath.GET("/master-induk-kerja", indukKerja.HandleGetIndukKerja(a)) // bentrok dengan unit
	insaniGroupingPath.GET("/master-jabatan-fungsional", jabatanFungsional.HandleGetJabatanFungsional(a))
	insaniGroupingPath.GET("/master-jabatan-penetap", sk.HandleGetJabatanPenetap(a))
	insaniGroupingPath.GET("/master-jenis-nomor-registrasi", jenisNoRegis.HandleGetJenisNoRegis(a))
	insaniGroupingPath.GET("/master-jenis-pegawai", jenisPegawai.HandleGetAllJenisPegawai(a))
	insaniGroupingPath.GET("/master-jenis-pegawai-tidak-tetap", jenisPTT.HandleGetJenisPTT(a))
	insaniGroupingPath.GET("/master-jenis-sk", sk.HandleGetAllJenisSK(a))
	insaniGroupingPath.GET("/master-jenis-sk-pengangkatan", sk.HandleGetAllJenisSKPengangkatan(a))
	insaniGroupingPath.GET("/master-kelompok-pegawai", kelompokPegawai.HandleGetKelompokPegawai(a))
	insaniGroupingPath.GET("/master-lokasi-kerja", lokasiKerja.HandleGetLokasiKerja(a))
	insaniGroupingPath.GET("/master-makul", sk.HandleGetAllMataKuliah(a)) // still dummy
	insaniGroupingPath.GET("/master-pangkat-golongan-pegawai", pangkatPegawai.HandleGetPangkatGolonganPegawai(a))
	insaniGroupingPath.GET("/master-pegawai-penetap", sk.HandleGetPegawaiPenetap(a)) // still error
	insaniGroupingPath.GET("/master-status-pegawai", statusPegawai.HandleGetAllStatusPegawai(a))
	insaniGroupingPath.GET("/master-status-pegawai-aktif", statusPegawaiAktif.HandleGetStatusPegawaiAktif(a))
	insaniGroupingPath.GET("/master-induk-kerja", indukKerja.HandleGetIndukKerja(a))
	insaniGroupingPath.GET("/master-homebase", indukKerja.HandleHomebase(a))

	insaniGroupingPath.GET("/pegawai/personal", personal.HandleSearchPersonal(a))
	insaniGroupingPath.GET("/master-status-pengangkatan", sk.HandleGetAllStatusPengangkat(a))
	insaniGroupingPath.GET("/master-unit-kerja", unitKerja.HandleGetUnitKerja(a))
	insaniGroupingPath.GET("/master-unit-pengangkat", unitKerja.HandleGetUnitPengangkat(a))

	// insaniGroupingPath.POST("/sk-pengangkatan-tendik", skPengangkatan.HandleCreateSKPengangkatanTendik(a))
	insaniGroupingPath.POST("/sk-pengangkatan-tendik", skV2.HandleCreateSkPengangkatanTendik(a))
	insaniGroupingPath.GET("/sk-pengangkatan-tendik", skPengangkatan.HandleGetDetailSKPengangkatanTendik(a))
	insaniGroupingPath.PUT("/sk-pengangkatan-tendik", skV2.HandleUpdateSkPengangkatanTendik(a))
	insaniGroupingPath.DELETE("/sk-pengangkatan-tendik", skV2.HandleDeleteSKPengangkatanTendik(a))

	insaniGroupingPath.POST("/sk-pengangkatan-dosen", skV2.HandleCreateSkPengangkatanDosen(a))
	// insaniGroupingPath.POST("/sk-pengangkatan-dosen", skPengangkatan.HandleCreateSKPengangkatanDosen(a))
	insaniGroupingPath.PUT("/sk-pengangkatan-dosen", skV2.HandleUpdateSkPengangkatanDosen(a))
	// insaniGroupingPath.GET("/sk-pengangkatan-dosen", skPengangkatan.HandleGetDetailSKPengangkatanDosen(a))
	insaniGroupingPath.GET("/sk-pengangkatan-dosen", skV2.HandleGetSkPengangkatanDosen(a))
	insaniGroupingPath.DELETE("/sk-pengangkatan-dosen", skV2.HandleDeleteSkPengangkatanDosen(a))
	// insaniGroupingPath.DELETE("/sk-pengangkatan-dosen", skPengangkatan.HandleDeleteSKPengangkatanDosenByUUID(a))
	// insaniGroupingPath.DELETE("/sk-pengangkatan-tendik", skPengangkatan.HandleDeleteSKPengangkatanTendikByUUID(a))

	insaniGroupingPath.GET("/sk-pengangkatan-tendik-v2", skV2.HandleGetSkPengangkatanTendik(a))

	// SK list
	insaniGroupingPath.GET("/sk-pegawai", sk.HandleGetAllSKPegawai(a))

	// SK Kenaikan Gaji Berkala
	insaniGroupingPath.GET("/sk-kgb/:uuidSk/detail", sk.HandleGetSkKenaikanGajiDummy(a))
	insaniGroupingPath.PUT("/sk-kgb/:uuidSk", sk.HandleUpdateSkKenaikanGaji(a))

	// Testing Gorm
	insaniGroupingPath.GET("/pegawai2", pegawai.HandleGetPegawaix(a))
	insaniGroupingPath.GET("/pegawai2/:uuidPersonal", pegawai.HandleGetPegawaiByUUIDx(a))

	// Update Simpeg
	insaniGroupingPath.PUT("/pegawai/:uuidPegawai", pegawai.HandleUpdatePegawai(a, appCtx, slackErrChan))

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
