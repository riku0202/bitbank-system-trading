package bitbank

import (
	"net/http"
)

type APIClient struct {
	key        string
	secret     string
	httpClient *http.Client
}

func New(key, secret string) *APIClient {
	apiClient := &APIClient{
		key:        key,
		secret:     secret,
		httpClient: &http.Client{},
	}

	return apiClient
}