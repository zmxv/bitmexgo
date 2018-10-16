package bitmexgo

import (
	"time"
)

type StatsHistory struct {
	Date       time.Time `json:"date"`
	RootSymbol string    `json:"rootSymbol"`
	Currency   string    `json:"currency,omitempty"`
	Volume     int       `json:"volume,omitempty"`
	Turnover   int       `json:"turnover,omitempty"`
}
