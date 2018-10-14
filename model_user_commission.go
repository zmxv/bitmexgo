package bitmexgo

type UserCommission struct {
	MakerFee      float64 `json:"makerFee,omitempty"`
	TakerFee      float64 `json:"takerFee,omitempty"`
	SettlementFee float64 `json:"settlementFee,omitempty"`
	MaxFee        float64 `json:"maxFee,omitempty"`
}
