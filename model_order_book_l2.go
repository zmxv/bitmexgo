package bitmexgo

type OrderBookL2 struct {
	Symbol string  `json:"symbol"`
	Id     float32 `json:"id"`
	Side   string  `json:"side"`
	Size   float32 `json:"size,omitempty"`
	Price  float64 `json:"price,omitempty"`
}
