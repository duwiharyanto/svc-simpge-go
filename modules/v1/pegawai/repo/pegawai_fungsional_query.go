package repo

import (
	"fmt"
	"svc-insani-go/helper"
)

func getPegawaiFungsionalYayasanByIdPegawai(idPegawai string) string {
	q := fmt.Sprintf(`SELECT COALESCE(p.id,0), COALESCE(jf.fungsional,''),
		COALESCE(jf.id,0),
		COALESCE(jf.kd_fungsional,'')
		FROM pegawai p
		JOIN pegawai_fungsional pf ON p.id = pf.id_pegawai
		JOIN jabatan_fungsional jf ON pf.id_jabatan_fungsional = jf.id 
		WHERE p.id = %q`, idPegawai)
	return helper.FlatQuery(q)
}

func getPegawaiFungsionalNegaraByIdPegawai(idPegawai string) string {
	q := fmt.Sprintf(`SELECT COALESCE(p.id,0), COALESCE(jf.fungsional,''),
		COALESCE(jf.id,0),
		COALESCE(jf.kd_fungsional,'')
		FROM pegawai p
		JOIN pegawai_pns pfn ON p.id = pfn.id_pegawai
		JOIN jabatan_fungsional jf ON pfn.id_jabatan_fungsional = jf.id 
		WHERE p.id = %q`, idPegawai)
	return helper.FlatQuery(q)
}
