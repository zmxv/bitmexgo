package bitmexgo

type UserWithdrawalFees struct {
	Currency string `json:"currency,omitempty"`
	Fee      int    `json:"fee,omitempty"`
	MinFee   int    `json:"minFee,omitempty"`
	MaxFee   int    `json:"maxFee,omitempty"`
}
