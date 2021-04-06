package repo

import "fmt"

func getJenisSKQuery() string {
	return "SELECT kd_jenis_sk,nama_sk, uuid FROM jenis_sk WHERE flag_aktif=1"
}

func getJenisSKByUUID(uuid string) string {
	return fmt.Sprintf(`SELECT id, kd_jenis_sk, nama_sk, uuid FROM jenis_sk WHERE flag_aktif=1 AND uuid = %q`, uuid)
}

func getJenisSKByCode(code string) string {
	return fmt.Sprintf(`SELECT id, kd_jenis_sk, nama_sk, uuid FROM jenis_sk WHERE flag_aktif=1 AND kd_jenis_sk = %q`, code)
}
