package helper

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

type HttpHelper struct {
	client *http.Client
}

func NewHttpHelper(timeout time.Duration) *HttpHelper {
	return &HttpHelper{client: &http.Client{
		Timeout: timeout,
	}}
}

func (c *HttpHelper) Send(method, url string, body []byte, headers map[string]string) ([]byte, int, error) {
	var payload io.Reader

	if body != nil {
		payload = bytes.NewBuffer(body)
	}

	request, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to create request: %v", err)
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	response, err := c.client.Do(request)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to send request: %v", err)
	}
	defer response.Body.Close()

	message, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, response.StatusCode, fmt.Errorf("failed to read response: %v", err)
	}
	return message, response.StatusCode, nil
}
