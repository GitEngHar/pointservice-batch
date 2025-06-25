package client_handler

import (
	"io"
	"net/http"
)

type HttpHealthCheckHandler struct{}

func NewHttpHealthCheckHandler() *HttpHealthCheckHandler {
	return &HttpHealthCheckHandler{}
}

func (h *HttpHealthCheckHandler) Handle(r *http.Request) (string, error) {
	httpClient := &http.Client{}
	resp, err := httpClient.Do(r)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
