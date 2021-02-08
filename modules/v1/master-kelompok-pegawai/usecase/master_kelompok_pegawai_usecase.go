package usecase

import (
	"net/http"
	"svc-insani-go/app"
	JenisPegawaiModel "svc-insani-go/modules/v1/master-jenis-pegawai/model"
	JenisPegawaiRepo "svc-insani-go/modules/v1/master-jenis-pegawai/repo"
	"svc-insani-go/modules/v1/master-kelompok-pegawai/model"
	"svc-insani-go/modules/v1/master-kelompok-pegawai/repo"

	"github.com/labstack/echo"
)

func HandleGetKelompokPegawai(a app.App) echo.HandlerFunc {
	h := func(c echo.Context) error {
		KdJenisPegawai := c.QueryParam("kd_jenis_pegawai")
		var JenisPegawai *JenisPegawaiModel.JenisPegawai
		var IDJenisPegawai string
		var err error
		if KdJenisPegawai != "" {
			JenisPegawai, err = JenisPegawaiRepo.GetJenisPegawaiByKD(a, KdJenisPegawai)
			if err != nil {
				//fmt.Printf("[ERROR] %s\n", err.Error())
				return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
			}
			if JenisPegawai != nil {
				IDJenisPegawai = JenisPegawai.ID
			}
		}
		KelompokPegawai, err := repo.GetAllKelompokPegawai(a, IDJenisPegawai)
		if err != nil {
			//fmt.Printf("[ERROR] %s\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Layanan sedang bermasalah"})
		}

		return c.JSON(http.StatusOK, model.KelompokPegawaiResponse{
			Data: KelompokPegawai,
		})
	}
	return echo.HandlerFunc(h)
}
