package bitmexgo

import (
	"time"
)

// Swap Funding History
type Funding struct {
	Timestamp        time.Time `json:"timestamp"`
	Symbol           string    `json:"symbol"`
	FundingInterval  time.Time `json:"fundingInterval,omitempty"`
	FundingRate      float64   `json:"fundingRate,omitempty"`
	FundingRateDaily float64   `json:"fundingRateDaily,omitempty"`
}
