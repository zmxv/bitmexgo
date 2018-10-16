package bitmexgo

import (
	"time"
)

// Persistent API Keys for Developers
type ApiKey struct {
	Id          string    `json:"id"`
	Secret      string    `json:"secret"`
	Name        string    `json:"name"`
	Nonce       int       `json:"nonce"`
	Cidr        string    `json:"cidr,omitempty"`
	Permissions []string  `json:"permissions,omitempty"`
	Enabled     bool      `json:"enabled,omitempty"`
	UserId      int       `json:"userId"`
	Created     time.Time `json:"created,omitempty"`
}
