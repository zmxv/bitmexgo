package bitmexgo

import (
	"time"
)

// Individual & Bucketed Trades
type Trade struct {
	Timestamp       time.Time `json:"timestamp"`
	Symbol          string    `json:"symbol"`
	Side            string    `json:"side,omitempty"`
	Size            int       `json:"size,omitempty"`
	Price           float64   `json:"price,omitempty"`
	TickDirection   string    `json:"tickDirection,omitempty"`
	TrdMatchID      string    `json:"trdMatchID,omitempty"`
	GrossValue      int       `json:"grossValue,omitempty"`
	HomeNotional    float64   `json:"homeNotional,omitempty"`
	ForeignNotional float64   `json:"foreignNotional,omitempty"`
}
