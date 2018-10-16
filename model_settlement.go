package bitmexgo

import (
	"time"
)

// Historical Settlement Data
type Settlement struct {
	Timestamp             time.Time `json:"timestamp"`
	Symbol                string    `json:"symbol"`
	SettlementType        string    `json:"settlementType,omitempty"`
	SettledPrice          float64   `json:"settledPrice,omitempty"`
	OptionStrikePrice     float64   `json:"optionStrikePrice,omitempty"`
	OptionUnderlyingPrice float64   `json:"optionUnderlyingPrice,omitempty"`
	Bankrupt              int       `json:"bankrupt,omitempty"`
	TaxBase               int       `json:"taxBase,omitempty"`
	TaxRate               float64   `json:"taxRate,omitempty"`
}
