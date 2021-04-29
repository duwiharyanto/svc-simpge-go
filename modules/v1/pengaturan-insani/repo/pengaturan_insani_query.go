package repo

import "fmt"

func getPengaturanQuery(atribut string) string {
	return fmt.Sprintf("SELECT COALESCE(nilai_atribut, '') FROM pengaturan_insani WHERE atribut_pengaturan = %q AND flag_aktif = 1", atribut)
}

func insertPengaturanQuery() string {
	return fmt.Sprintf("INSERT INTO pengaturan_insani (atribut_pengaturan, nilai_atribut, user_input, user_update) VALUES (?, ?, ?, ?)")
}

func updatePengaturanQuery() string {
	return fmt.Sprintf("UPDATE pengaturan_insani SET nilai_atribut = ?, user_update = ? WHERE atribut_pengaturan = ? AND flag_aktif = 1")
}
