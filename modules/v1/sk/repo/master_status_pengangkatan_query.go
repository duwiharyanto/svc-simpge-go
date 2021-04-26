package repo

import "fmt"

func getStatusPengangkatanQuery() string {
	return "SELECT kd_status_pengangkatan,status_pengangkatan, uuid FROM status_pengangkatan WHERE flag_aktif=1"
}
func getStatusPengangkatanQueryByUUID(uuid string) string {
	return fmt.Sprintf(`SELECT id, kd_status_pengangkatan, status_pengangkatan, uuid FROM status_pengangkatan WHERE uuid = %q`, uuid)
}
