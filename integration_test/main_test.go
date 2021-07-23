package integrationtest

import (
	"testing"
)

func TestMain(t *testing.T) {
	srv, err := UpServer()
	if err != nil {
		t.Fatal(err)
	}
	defer srv.Server.Close()

	// t.Run("kelompok", GetKelompokPegawai(t, srv, groups))
	t.Run("pegawai", Pegawai(t, srv))
}
