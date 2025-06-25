package client

import (
	"fmt"
	"net/http"
)

type Handler interface {
	Handle(*http.Request) (string, error)
}

type HttpClient struct {
	config  Config
	request *http.Request
	handler Handler
}

func NewHttpClient(config *Config, handler Handler) *HttpClient {
	return &HttpClient{
		config:  *config,
		handler: handler,
	}
}

func (h *HttpClient) Client() *HttpClient {
	url := fmt.Sprintf("%s://%s", h.config.url.Scheme, h.config.url.Path)
	req, err := http.NewRequest(h.config.method, url, nil)
	if err != nil {
		return nil
	}
	req.Header.Set("Content-Type", "application/json")
	h.request = req
	return h
}

func (h *HttpClient) Start() (string, error) {
	resp, err := h.handler.Handle(h.request)
	if err != nil {
		return "", err
	}
	return resp, nil
}
