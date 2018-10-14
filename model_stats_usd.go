package bitmexgo

type StatsUsd struct {
	RootSymbol   string  `json:"rootSymbol"`
	Currency     string  `json:"currency,omitempty"`
	Turnover24h  float32 `json:"turnover24h,omitempty"`
	Turnover30d  float32 `json:"turnover30d,omitempty"`
	Turnover365d float32 `json:"turnover365d,omitempty"`
	Turnover     float32 `json:"turnover,omitempty"`
}
