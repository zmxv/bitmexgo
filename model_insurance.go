package bitmexgo

import (
	"time"
)

// Insurance Fund Data
type Insurance struct {
	Currency      string    `json:"currency"`
	Timestamp     time.Time `json:"timestamp"`
	WalletBalance float32   `json:"walletBalance,omitempty"`
}
