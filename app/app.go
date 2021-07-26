package app

import (
	"bytes"
	"context"
	"crypto/tls"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"svc-insani-go/app/minio"
	"time"

	"gorm.io/gorm"
)

type App struct {
	DB              *sql.DB
	GormDB          *gorm.DB
	HttpClient      *http.Client
	MinioBucketName string
	MinioClient     minio.MinioClient
	Name            string
	TimeLocation    *time.Location
}

func GetFixedTimeZone() *time.Location {
	timeOffset := os.Getenv("TIME_OFFSET")
	if timeOffset == "" {
		timeOffset = "7"
	}
	nTimeOffset, _ := strconv.Atoi(timeOffset)
	return time.FixedZone("", nTimeOffset*60*60)
}

func DefaultHttpTransport() *http.Transport {
	return &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		MaxIdleConns:        30,
		MaxIdleConnsPerHost: 30,
		MaxConnsPerHost:     50,
		IdleConnTimeout:     5 * time.Minute,
		DialContext: (&net.Dialer{
			Timeout: 3 * time.Second,
		}).DialContext,
	}
}

const (
	contentTypeJSON            = "application/json"
	contentTypeAppilcationForm = "application/form"
)

func SendHttpRequest(ctx context.Context, client *http.Client, method, baseURL, contentType string, header map[string]string, body interface{}) (*http.Response, error) {
	var reqBody io.Reader
	// mengecek http request content type

	fmt.Printf("[DEBUG] endpoint: %s %s\n", method, baseURL)
	switch contentType {
	case contentTypeJSON:
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error marshaling request body: %w", err)
		}
		// fmt.Printf("[DEBUG] req body in send http req: %s\n", jsonBody)
		reqBody = bytes.NewBuffer(jsonBody)
	case contentTypeAppilcationForm:
		formValues := body.(url.Values)
		reqBody = strings.NewReader(formValues.Encode())
	default:
		reqBody = nil
	}

	// reqTimeoutDur, err := time.ParseDuration(os.Getenv("RESPONSE_TIMEOUT_MS") + "ms")
	// if err != nil {
	// 	// fmt.Printf("[DEBUG] error parsing duration for http request timeout: %s\n", err.Error())
	// 	reqTimeoutDur = 50 * time.Second // default
	// }

	// membuat http request baru
	// ctx, cancel := context.WithTimeout(context.Background(), reqTimeoutDur)
	// ctx, cancel := context.WithTimeout(ctx, reqTimeoutDur)
	// defer cancel()
	req, err := http.NewRequestWithContext(ctx, method, baseURL, reqBody)
	// req, err := http.NewRequest(method, baseURL, reqBody)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	// req = req.WithContext(ctx)

	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	if len(header) != 0 {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}
	// mengirim request
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	return res, nil
}

const (
	ErrInternalServer = "Layanan sedang bermasalah"
)

func NewSlackLogger(ctx context.Context, client *http.Client) chan error {
	c := make(chan error)
	go watchSlackError(ctx, client, c)
	return c
}

func watchSlackError(ctx context.Context, client *http.Client, c chan error) {
	fmt.Printf("[DEBUG] slack logger is watching...\n")
	for err := range c {
		// send to slack
		go func(err error) {
			errSlack := sendErrorToSlack(ctx, client, err)
			if errSlack != nil {
				fmt.Printf("[ERROR] send error to slack: %s\n", err.Error())
			}
		}(err)
	}
}

const (
	slackMsgTemplate = `{
		"blocks": [
			{
				"type": "section",
				"text": {
					"type": "mrkdwn",
					"text": "Prod | svc-insani-go | 500 | %s"
				}
			}
		]
	}`
)

var (
	slackHookEndpoint = os.Getenv("SLACK_WEBHOOK_URL")
)

func sendErrorToSlack(ctx context.Context, client *http.Client, err error) error {
	// msg := fmt.Sprintf(slackMsgTemplate, strings.Trim(fmt.Sprintf("%q", err), `"`))
	msg := strings.Trim(fmt.Sprintf("%q", err), `"`)
	header := map[string]string{"Accept": "application/json"}
	slackMsgTmpl := map[string]interface{}{
		"blocks": []map[string]interface{}{
			map[string]interface{}{
				"type": "section",
				"text": map[string]interface{}{
					"type": "mrkdwn",
					"text": fmt.Sprintf("Prod | svc-insani-go | 500\n```%s```", msg),
				},
			},
		},
	}
	res, err := SendHttpRequest(ctx, client, http.MethodPost, slackHookEndpoint, contentTypeJSON, header, slackMsgTmpl)
	if err != nil {
		return fmt.Errorf("error send http request: %w", err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error read response body: %w", err)
	}
	fmt.Printf("[DEBUG] raw res body: %s\n", resBody)

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error status not ok: %s", resBody)
	}
	return nil
}
