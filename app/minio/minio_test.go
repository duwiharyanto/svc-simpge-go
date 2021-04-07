package minio

import (
	"os"
	"testing"
)

func TestUploadFormFile(t *testing.T) {
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
	ff.Append("insani", "test1", "", "plain/txt", 123, file)
	objName := ff.GenerateObjectName("test1", "sk", "pengangkatan", "uuid1")
	t.Log("objName:", objName)
	err = ff.Upload()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetUrlFormFile(t *testing.T) {
	mc, err := Connect()
	if err != nil {
		t.Fatal(err)
	}

	ff := NewFormFile(&mc)
	// var buf bytes.Buffer
	// ff.Append("insani", "test1", "sk/pengangkatan/uuid1/1617522314.txt", "", 0, nil)
	ff.Append("insani", "test1", "sk/pengangkatan/e37046f9-1437-11eb-a014-7eb0d4a3c7a0/a-0card.jpg", "", 0, nil)
	err = ff.GenerateUrl()
	if err != nil {
		t.Fatal(err)
	}

	url, _ := ff.GetUrl("test1")
	if url == "" {
		t.Fatal("should not be empty")
	}

	t.Log("url:", url)
}
