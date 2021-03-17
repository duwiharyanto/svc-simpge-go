package repo

import (
	"fmt"
)

func getLokasiKerjaQuery() string {
	return fmt.Sprintf(`SELECT id, COALESCE(lokasi_kerja,''), COALESCE(lokasi_desc,''), uuid FROM lokasi_kerja WHERE flag_aktif=1`)
}
