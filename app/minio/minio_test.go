package minio

import (
	"os"
	"testing"
)

func TestFormFile(t *testing.T) {
	mc, err := Connect()
	if err != nil {
		t.Fatal(err)
	}

	file, err := os.Open("minio.go")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	ff := NewFormFile(&mc)
	// var buf bytes.Buffer
	ff.Append("insani", "test1", "plain/txt", 123, file)
	objName := ff.GenerateObjectName("test1", "sk", "pengangkatan", "uuid1")
	t.Log("objName:", objName)
	err = ff.Upload()
	if err != nil {
		t.Fatal(err)
	}
}
