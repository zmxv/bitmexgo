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
	Trades          float32   `json:"trades,omitempty"`
	Volume          float32   `json:"volume,omitempty"`
	Vwap            float64   `json:"vwap,omitempty"`
	LastSize        float32   `json:"lastSize,omitempty"`
	Turnover        float32   `json:"turnover,omitempty"`
	HomeNotional    float64   `json:"homeNotional,omitempty"`
	ForeignNotional float64   `json:"foreignNotional,omitempty"`
}
