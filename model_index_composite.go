package bitmexgo

import (
	"time"
)

type IndexComposite struct {
	Timestamp   time.Time `json:"timestamp"`
	Symbol      string    `json:"symbol,omitempty"`
	IndexSymbol string    `json:"indexSymbol,omitempty"`
	Reference   string    `json:"reference,omitempty"`
	LastPrice   float64   `json:"lastPrice,omitempty"`
	Weight      float64   `json:"weight,omitempty"`
	Logged      time.Time `json:"logged,omitempty"`
}
