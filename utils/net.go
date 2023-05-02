package utils

import (
	"bytes"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/muhadyan/financial-planner/config"
)

func GetHTTPRequestJSON(method string, url string, body io.Reader, headers ...map[string]string) (res []byte, statusCode int, err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	// iterate optional data of headers
	for _, header := range headers {
		for key, value := range header {
			req.Header.Set(key, value)
		}
	}

	timeout, _ := strconv.Atoi(config.GetConfig().DefaultTimeout)
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
	r, err := client.Do(req)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	resp := StreamToByte(r.Body)

	defer func() {
		r.Body.Close()
	}()

	return resp, r.StatusCode, nil
}

func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}
