package repo

import (
	"fmt"
)

func getLokasiKerjaQuery() string {
	return fmt.Sprintf(`SELECT id, lokasi_kerja, lokasi_desc, uuid FROM lokasi_kerja WHERE flag_aktif=1`)
}
