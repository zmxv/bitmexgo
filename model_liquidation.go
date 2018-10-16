package bitmexgo

// Active Liquidations
type Liquidation struct {
	OrderID   string  `json:"orderID"`
	Symbol    string  `json:"symbol,omitempty"`
	Side      string  `json:"side,omitempty"`
	Price     float64 `json:"price,omitempty"`
	LeavesQty int     `json:"leavesQty,omitempty"`
}
