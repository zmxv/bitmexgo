package bitmexgo

import (
	"time"
)

type StatsHistory struct {
	Date       time.Time `json:"date"`
	RootSymbol string    `json:"rootSymbol"`
	Currency   string    `json:"currency,omitempty"`
	Volume     float32   `json:"volume,omitempty"`
	Turnover   float32   `json:"turnover,omitempty"`
}
