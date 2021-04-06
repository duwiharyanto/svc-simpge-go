package minio

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/minio/minio-go"
)

func Connect() (MinioClient, error) {
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ACCESSKEYID")
	secretAccessKey := os.Getenv("MINIO_SECRETACCESSKEY")
	useSSL, _ := strconv.ParseBool(os.Getenv("MINIO_SSL"))

	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)

	if err != nil {
		return MinioClient{}, err
	}

	minioClient.SetCustomTransport(&http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		DialContext: (&net.Dialer{
			Timeout: 3 * time.Second,
		}).DialContext,
	})

	return MinioClient{minioClient}, nil
}

type MinioClient struct {
	*minio.Client
}

func (mc MinioClient) CreateOrGetBucket(bucketName string) (string, error) {
	exists, err := mc.BucketExists(bucketName)
	if err != nil {
		return "", err
	}

	if exists {
		return bucketName, nil
	}

	location := os.Getenv("MINIO_LOCATION")
	err = mc.MakeBucket(bucketName, location)
	if err != nil {
		return "", err
	}

	return bucketName, nil
}

func (mc MinioClient) Upload(bucketName, objectName string, reader io.Reader, objectSize int64, contentType string) error {
	_, err := mc.PutObject(bucketName, objectName, reader, objectSize, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return err
	}

	return nil
}

func (mc MinioClient) Delete(bucketName, objectName string) error {
	err := mc.RemoveObject(bucketName, objectName)
	if err != nil {
		return err
	}
	return nil
}

func (mc MinioClient) GetDownloadURL(bucketName, objectName, downloadName string) (string, error) {
	reqParams := make(url.Values)
	if downloadName != "" {
		reqParams.Set("response-content-disposition", "filename=\""+downloadName+"\"")
	}

	presignedURL, err := mc.PresignedGetObject(bucketName, objectName, expirationTime(), reqParams)
	if err != nil {
		return "", fmt.Errorf("error get presigned get object: %w", err)
	}

	// b, err := presignedURL.MarshalBinary()
	// if err != nil {
	// 	return "", fmt.Errorf("error marshaling url: %w", err)
	// }

	return presignedURL.String(), nil
	// return string(b), nil
}

func (mc MinioClient) Get(bucketName, objectName string) (*minio.Object, error) {
	object, err := mc.GetObject(bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return object, nil
}

func expirationTime() time.Duration {
	defaultSecond := 600
	expSecond, err := strconv.Atoi(os.Getenv("MINIO_EXPIRATION_SECOND"))
	if err != nil {
		expSecond = defaultSecond
	}
	return time.Duration(expSecond) * time.Second
}

type FormFile struct {
	mc      *MinioClient
	fileMap map[string]formFile
}

func NewFormFile(minioClient *MinioClient) FormFile {
	return FormFile{
		mc:      minioClient,
		fileMap: make(map[string]formFile),
	}
}

func (ff FormFile) Append(bucket, name, path, contentType string, size int64, file io.Reader) {
	ff.fileMap[name] = formFile{
		bucket:      bucket,
		contentType: contentType,
		path:        path,
		size:        size,
		file:        file,
	}
}

func (ff FormFile) GetUrl(name string) string {
	f, exist := ff.fileMap[name]
	if !exist {
		return ""
	}
	return f.url
}

func (ff FormFile) GenerateUrl() error {
	for key := range ff.fileMap {
		var err error
		f := ff.fileMap[key]
		f.url, err = ff.mc.GetDownloadURL(f.bucket, f.path, "")
		if err != nil {
			return fmt.Errorf("error get download url: %w", err)
		}
		ff.fileMap[key] = f
	}

	return nil
}

func (ff FormFile) Upload() error {
	for _, f := range ff.fileMap {
		err := ff.mc.Upload(f.bucket, f.path, f.file, f.size, f.contentType)
		if err != nil {
			return fmt.Errorf("error form file upload: %w", err)
		}
	}
	return nil
}

func (ff FormFile) GenerateObjectName(name string, dirs ...string) string {
	f := ff.fileMap[name]
	f.path = strings.Join(dirs, "/")
	splitted := strings.Split(f.contentType, "/")
	var extension string
	if len(splitted) == 2 {
		extension = splitted[1]
	} else {
		extension = splitted[0]
	}
	f.path += fmt.Sprintf("/%d.%s", time.Now().Unix(), extension)
	ff.fileMap[name] = f
	return f.path
}

type formFile struct {
	bucket      string
	contentType string
	path        string
	url         string
	size        int64
	file        io.Reader
}
