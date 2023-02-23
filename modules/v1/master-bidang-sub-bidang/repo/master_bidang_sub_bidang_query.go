package repo

import "fmt"

func getBidangQuery() string {
	return "SELECT kd_bidang,bidang, uuid FROM bidang WHERE flag_aktif=1"
}
func getBidangByUUIDQuery(uuid string) string {
	return fmt.Sprintf(`SELECT id, kd_bidang,bidang, uuid FROM bidang WHERE flag_aktif=1 AND uuid = %q`, uuid)
}

// sub bidang
func getSubBidangQuery() string {
	return "SELECT b.kd_bidang, b.bidang,sb.kd_sub_bidang,sb.sub_bidang, sb.uuid FROM sub_bidang sb JOIN bidang b ON sb.id_bidang = b.id WHERE sb.flag_aktif=1"
}
func getSubBidangByUUIDQuery(uuid string) string {
	return fmt.Sprintf(`SELECT sb.id,sb.kd_sub_bidang,sb.sub_bidang, sb.uuid FROM sub_bidang sb WHERE sb.flag_aktif=1 AND sb.uuid = %q`, uuid)
}
