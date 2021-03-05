package repo

import (
	"fmt"
)

func getStatusPegawaiAktifQuery(flagStatus string) string {
	return fmt.Sprintf(`SELECT id, kd_status, status, uuid FROM status_pegawai_aktif WHERE flag_aktif=1 AND flag_status_aktif = %q`, flagStatus)
}
