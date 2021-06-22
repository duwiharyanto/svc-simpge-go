package repo

import (
	"fmt"
)

func getDetailProfesiQuery() string {
	return fmt.Sprintf(`SELECT id, COALESCE(detail_profesi,''), uuid FROM detail_profesi WHERE flag_aktif=1`)
}
