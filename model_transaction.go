package bitmexgo

import (
	"time"
)

type Transaction struct {
	TransactID     string    `json:"transactID"`
	Account        int       `json:"account,omitempty"`
	Currency       string    `json:"currency,omitempty"`
	TransactType   string    `json:"transactType,omitempty"`
	Amount         int       `json:"amount,omitempty"`
	Fee            int       `json:"fee,omitempty"`
	TransactStatus string    `json:"transactStatus,omitempty"`
	Address        string    `json:"address,omitempty"`
	Tx             string    `json:"tx,omitempty"`
	Text           string    `json:"text,omitempty"`
	TransactTime   time.Time `json:"transactTime,omitempty"`
	Timestamp      time.Time `json:"timestamp,omitempty"`
}
