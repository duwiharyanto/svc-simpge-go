package repo

import "fmt"

func getBidangSubBidangQuery() string {
	return "SELECT kd_bidang_sub_bidang,bidang_sub_bidang, uuid FROM bidang_sub_bidang WHERE flag_aktif=1"
}
func getBidangSubBidangByUUIDQuery(uuid string) string {
	return fmt.Sprintf(`SELECT kd_bidang_sub_bidang,bidang_sub_bidang, uuid FROM bidang_sub_bidang WHERE flag_aktif=1 AND uuid = %q`, uuid)
}
