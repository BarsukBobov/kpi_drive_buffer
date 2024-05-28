package misc

import (
	"net/http"
	"time"
)

type HttpAdapter struct {
	DefaultUrl string
	Client     *http.Client
}

func NewHttpAdapter(defaultUrl string) *HttpAdapter {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}
	return &HttpAdapter{
		DefaultUrl: defaultUrl,
		Client:     &httpClient,
	}
}
