package bitmexgo

import (
	"time"
)

type Transaction struct {
	TransactID     string    `json:"transactID"`
	Account        float32   `json:"account,omitempty"`
	Currency       string    `json:"currency,omitempty"`
	TransactType   string    `json:"transactType,omitempty"`
	Amount         float32   `json:"amount,omitempty"`
	Fee            float32   `json:"fee,omitempty"`
	TransactStatus string    `json:"transactStatus,omitempty"`
	Address        string    `json:"address,omitempty"`
	Tx             string    `json:"tx,omitempty"`
	Text           string    `json:"text,omitempty"`
	TransactTime   time.Time `json:"transactTime,omitempty"`
	Timestamp      time.Time `json:"timestamp,omitempty"`
}
