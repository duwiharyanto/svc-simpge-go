package repo

import (
	"database/sql"
	"fmt"
	"svc-insani-go/app"
	"svc-insani-go/modules/v1/pegawai/model"
)

func GetKepegawaianYayasan(a *app.App, uuid string) (*model.PegawaiYayasan, error) {
	sqlQuery := getPegawaiYayasanQuery(uuid)
	var pegawaiYayasan model.PegawaiYayasan

	err := a.DB.QueryRow(sqlQuery).Scan(
		&pegawaiYayasan.ID,
		&pegawaiYayasan.UuidJenisPegawai,
		&pegawaiYayasan.IDJenisPegawai,
		&pegawaiYayasan.KDJenisPegawai,
		&pegawaiYayasan.JenisPegawai,
		&pegawaiYayasan.UuidStatusPegawai,
		&pegawaiYayasan.IDStatusPegawai,
		&pegawaiYayasan.KDStatusPegawai,
		&pegawaiYayasan.StatusPegawai,
		&pegawaiYayasan.UuidKelompokPegawai,
		&pegawaiYayasan.IdKelompokPegawai,
		&pegawaiYayasan.KdKelompokPegawai,
		&pegawaiYayasan.KelompokPegawai,
		&pegawaiYayasan.UuidKelompokPegawaiPayroll,
		&pegawaiYayasan.IdKelompokPegawaiPayroll,
		&pegawaiYayasan.KdKelompokPegawaiPayroll,
		&pegawaiYayasan.KelompokPegawaiPayroll,
		&pegawaiYayasan.IdPendidikanMasuk,
		&pegawaiYayasan.UuidPendidikanMasuk,
		&pegawaiYayasan.KdPendidikanMasuk,
		&pegawaiYayasan.KdPendidikanMasukSimpeg,
		&pegawaiYayasan.PendidikanMasuk,
		&pegawaiYayasan.UuidPendidikanTerakhir,
		&pegawaiYayasan.IdPendidikanTerakhir,
		&pegawaiYayasan.KdPendidikanTerakhir,
		&pegawaiYayasan.KdPendidikanTerakhirSimpeg,
		&pegawaiYayasan.PendidikanTerakhir,
		&pegawaiYayasan.UuidPangkatGolongan,
		&pegawaiYayasan.KdPangkat,
		&pegawaiYayasan.Pangkat,
		&pegawaiYayasan.KdGolongan,
		&pegawaiYayasan.Golongan,
		&pegawaiYayasan.KdRuang,
		&pegawaiYayasan.TmtPangkatGolongan,
		&pegawaiYayasan.UuidJabatanFungsional,
		&pegawaiYayasan.KdJabatanFungsional,
		&pegawaiYayasan.JabatanFungsional,
		&pegawaiYayasan.TmtJabatan,
		&pegawaiYayasan.MasaKerjaBawaanTahun,
		&pegawaiYayasan.MasaKerjaBawaanBulan,
		&pegawaiYayasan.MasaKerjaGajiTahun,
		&pegawaiYayasan.MasaKerjaGajiBulan,
		&pegawaiYayasan.MasaKerjaAwalKepegawaianTahun,
		&pegawaiYayasan.MasaKerjaAwalKepegawaianBulan,
		&pegawaiYayasan.MasaKerjaAwalPensiunTahun,
		&pegawaiYayasan.MasaKerjaAwalPensiunBulan,
		&pegawaiYayasan.AngkaKredit,
		&pegawaiYayasan.NoSertifikasi,
		&pegawaiYayasan.Nidn,
		&pegawaiYayasan.UuidJenisRegis,
		&pegawaiYayasan.IdJenisRegis,
		&pegawaiYayasan.KdJenisRegis,
		&pegawaiYayasan.JenisNomorRegis,
		&pegawaiYayasan.NomorRegis,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("error querying and scanning kepegawaian yayasan, %s", err.Error())
	}

	pegawaiYayasan.SetTanggalIDN()
	// pegawaiYayasan.SetMasaKerjaTotal()

	return &pegawaiYayasan, nil
}
