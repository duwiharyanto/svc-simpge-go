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

	// personal, err := srv.TestPersonal()
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// fmt.Printf("[DEBUG] personal: %+v\n", personal)
	// fmt.Println("len data:", len(personal["data"].([]interface{})))

	// srv.CreatePersonal(t)
}
