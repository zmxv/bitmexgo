package bitmexgo

import (
	"time"
)

type TradeBin struct {
	Timestamp       time.Time `json:"timestamp"`
	Symbol          string    `json:"symbol"`
	Open            float64   `json:"open,omitempty"`
	High            float64   `json:"high,omitempty"`
	Low             float64   `json:"low,omitempty"`
	Close           float64   `json:"close,omitempty"`
	Trades          int       `json:"trades,omitempty"`
	Volume          int       `json:"volume,omitempty"`
	Vwap            float64   `json:"vwap,omitempty"`
	LastSize        int       `json:"lastSize,omitempty"`
	Turnover        int       `json:"turnover,omitempty"`
	HomeNotional    float64   `json:"homeNotional,omitempty"`
	ForeignNotional float64   `json:"foreignNotional,omitempty"`
}
