package minio

import (
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
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
		return "", err
	}

	return presignedURL.String(), nil
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
