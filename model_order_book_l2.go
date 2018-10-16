package bitmexgo

type OrderBookL2 struct {
	Symbol string  `json:"symbol"`
	Id     int     `json:"id"`
	Side   string  `json:"side"`
	Size   int     `json:"size,omitempty"`
	Price  float64 `json:"price,omitempty"`
}
