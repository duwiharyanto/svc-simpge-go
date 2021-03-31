package repo

import (
	"fmt"
	"strings"
)

func getMataKuliahQuery() string {
	return "SELECT kd_makul,nama_makul, uuid FROM makul WHERE flag_aktif=1"
}
func getMataKuliahByUUID(uuid string) string {
	return fmt.Sprintf(`SELECT id, kd_makul, nama_makul, uuid FROM makul WHERE flag_aktif=1 AND uuid = %q`, uuid)
}
func getMataKuliahIDByUUID(uuid []string) string {
	uuidJoined := strings.Join(uuid, "', '")
	uuidStr := fmt.Sprintf("'%s'", uuidJoined) // outputnya: '123', '123, '123'
	return fmt.Sprintf(`SELECT id FROM makul WHERE flag_aktif=1 AND uuid IN (%s)`, uuidStr)
}
