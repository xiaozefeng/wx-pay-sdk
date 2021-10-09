package http

import (
	"bytes"
	"context"
	"io"
	"net/http"
)

func PostJSON(ctx *context.Context, url string, body []byte) ([]byte, error) {
	request, err := newRequest(url, http.MethodPost, body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	return executeRequest(request)
}

func executeRequest(request *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func newRequest(url, method string, body []byte) (*http.Request, error) {
	return http.NewRequest(method, url, bytes.NewBuffer(body))
}

func Get(ctx *context.Context, url string) ([]byte, error) {
	request, err := newRequest(url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	return executeRequest(request)
}
