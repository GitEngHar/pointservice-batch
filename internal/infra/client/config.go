package client

import (
	"net/url"
	"os"
)

type Config struct {
	Data   string `json:"data"`
	url    url.URL
	method string
}

func NewConfig() *Config {
	serverUrlSchema := os.Getenv("REQUEST_TO_SERVER_URL_SCHEMA")
	serverUrlPath := os.Getenv("REQUEST_TO_SERVER_URL_PATH")
	method := os.Getenv("REQUEST_TO_SERVER_METHOD")
	// TODO: error handle
	serverUrl := url.URL{
		Scheme: serverUrlSchema,
		Path:   serverUrlPath,
	}
	return &Config{
		url:    serverUrl,
		method: method,
	}
}
