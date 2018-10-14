package bitmexgo

import (
	"time"
)

// Persistent API Keys for Developers
type ApiKey struct {
	Id          string    `json:"id"`
	Secret      string    `json:"secret"`
	Name        string    `json:"name"`
	Nonce       float32   `json:"nonce"`
	Cidr        string    `json:"cidr,omitempty"`
	Permissions []XAny    `json:"permissions,omitempty"`
	Enabled     bool      `json:"enabled,omitempty"`
	UserId      float32   `json:"userId"`
	Created     time.Time `json:"created,omitempty"`
}
