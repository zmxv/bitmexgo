package bitmexgo

import (
	"time"
)

// Best Bid/Offer Snapshots & Historical Bins
type Quote struct {
	Timestamp time.Time `json:"timestamp"`
	Symbol    string    `json:"symbol"`
	BidSize   int       `json:"bidSize,omitempty"`
	BidPrice  float64   `json:"bidPrice,omitempty"`
	AskPrice  float64   `json:"askPrice,omitempty"`
	AskSize   int       `json:"askSize,omitempty"`
}
