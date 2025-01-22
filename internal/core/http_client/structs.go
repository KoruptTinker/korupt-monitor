package httpClient

import "context"

type httpClient interface {
	Request(ctx *context.Context, url string, method string, payload interface{}, headers map[string]string) (HttpResponse, error)
}

type HttpResponse struct {
	StatusCode      int
	Response        []byte
	ResponseHeaders map[string]string
	Latency         int64
}

type HttpClient struct{}
