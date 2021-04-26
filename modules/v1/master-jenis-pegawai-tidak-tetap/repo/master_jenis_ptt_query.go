package repo

import (
	"fmt"
)

func getJenisNomorRegistrasiQuery() string {
	return fmt.Sprintf(`SELECT id, kd_jenis_ptt, jenis_ptt, uuid FROM jenis_pegawai_tidak_tetap WHERE flag_aktif=1`)
}
