package bitmexgo

// Exchange Statistics
type Stats struct {
	RootSymbol   string `json:"rootSymbol"`
	Currency     string `json:"currency,omitempty"`
	Volume24h    int    `json:"volume24h,omitempty"`
	Turnover24h  int    `json:"turnover24h,omitempty"`
	OpenInterest int    `json:"openInterest,omitempty"`
	OpenValue    int    `json:"openValue,omitempty"`
}
