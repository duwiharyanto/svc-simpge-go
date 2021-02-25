package repo

import (
	"fmt"
)

func getJenisNomorRegistrasiQuery() string {
	return fmt.Sprintf(`SELECT id, kd_jenis_regis, jenis_no_regis, uuid FROM jenis_nomor_registrasi WHERE flag_aktif=1`)
}
