package bitmexgo

type StatsUsd struct {
	RootSymbol   string `json:"rootSymbol"`
	Currency     string `json:"currency,omitempty"`
	Turnover24h  int    `json:"turnover24h,omitempty"`
	Turnover30d  int    `json:"turnover30d,omitempty"`
	Turnover365d int    `json:"turnover365d,omitempty"`
	Turnover     int    `json:"turnover,omitempty"`
}
