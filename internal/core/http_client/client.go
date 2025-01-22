package httpClient

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func NewClient() httpClient {
	return HttpClient{}
}

func (h HttpClient) Request(ctx *context.Context, url string, method string, payload interface{}, headers map[string]string) (HttpResponse, error) {
	responseObj := HttpResponse{}
	requestPayload, err := h.preparePayload(payload)
	if err != nil {
		return responseObj, err
	}

	httpRequest, err := http.NewRequestWithContext(*ctx, method, url, requestPayload)
	if err != nil {
		return responseObj, err
	}

	h.prepareHeaders(headers, httpRequest)

	startTime := time.Now()
	result, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return responseObj, err
	}

	defer result.Body.Close()

	responseObj.StatusCode = result.StatusCode

	responseBytes, err := io.ReadAll(result.Body)
	if err != nil {
		return responseObj, err
	}
	latency := time.Since(startTime).Milliseconds()
	responseObj.Latency = latency
	responseObj.Response = responseBytes
	return responseObj, nil
}

func (h *HttpClient) preparePayload(payload interface{}) (*bytes.Reader, error) {
	payloadBytes, err := json.Marshal(&payload)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(payloadBytes), nil
}

func (h *HttpClient) prepareHeaders(headers map[string]string, httpRequest *http.Request) {
	for key, val := range headers {
		httpRequest.Header.Add(key, val)
	}
}
