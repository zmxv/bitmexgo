package bitmexgo

import (
	"time"
)

// Insurance Fund Data
type Insurance struct {
	Currency      string    `json:"currency"`
	Timestamp     time.Time `json:"timestamp"`
	WalletBalance int       `json:"walletBalance,omitempty"`
}
