package bitmexgo

type ErrorError struct {
	Message string `json:"message,omitempty"`
	Name    string `json:"name,omitempty"`
}
