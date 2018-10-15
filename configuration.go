package bitmexgo

import (
	"context"
	"net/http"
)

// contextKeys are used to identify the type of value in the context.
type contextKey string

var ContextAPIKey = contextKey("apikey")

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Secret string
}

type Configuration struct {
	BasePath      string            `json:"basePath,omitempty"`
	DefaultHeader map[string]string `json:"defaultHeader,omitempty"`
	UserAgent     string            `json:"userAgent,omitempty"`
	HTTPClient    *http.Client
}

func NewConfiguration() *Configuration {
	return &Configuration{
		BasePath:      "https://www.bitmex.com/api/v1",
		DefaultHeader: make(map[string]string),
		UserAgent:     "bitmexgo/1.0.0",
	}
}

func NewTestnetConfiguration() *Configuration {
	return &Configuration{
		BasePath:      "https://testnet.bitmex.com/api/v1",
		DefaultHeader: make(map[string]string),
		UserAgent:     "bitmexgo/1.0.0",
	}
}

func NewAPIKeyContext(key, secret string) context.Context {
	return context.WithValue(context.TODO(), ContextAPIKey, APIKey{Key: key, Secret: secret})
}
