package bitmexgo

// Exchange Statistics
type Stats struct {
	RootSymbol   string  `json:"rootSymbol"`
	Currency     string  `json:"currency,omitempty"`
	Volume24h    float32 `json:"volume24h,omitempty"`
	Turnover24h  float32 `json:"turnover24h,omitempty"`
	OpenInterest float32 `json:"openInterest,omitempty"`
	OpenValue    float32 `json:"openValue,omitempty"`
}
