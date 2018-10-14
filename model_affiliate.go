package bitmexgo

import (
	"time"
)

type Affiliate struct {
	Account         float32   `json:"account"`
	Currency        string    `json:"currency"`
	PrevPayout      float32   `json:"prevPayout,omitempty"`
	PrevTurnover    float32   `json:"prevTurnover,omitempty"`
	PrevComm        float32   `json:"prevComm,omitempty"`
	PrevTimestamp   time.Time `json:"prevTimestamp,omitempty"`
	ExecTurnover    float32   `json:"execTurnover,omitempty"`
	ExecComm        float32   `json:"execComm,omitempty"`
	TotalReferrals  float32   `json:"totalReferrals,omitempty"`
	TotalTurnover   float32   `json:"totalTurnover,omitempty"`
	TotalComm       float32   `json:"totalComm,omitempty"`
	PayoutPcnt      float64   `json:"payoutPcnt,omitempty"`
	PendingPayout   float32   `json:"pendingPayout,omitempty"`
	Timestamp       time.Time `json:"timestamp,omitempty"`
	ReferrerAccount float64   `json:"referrerAccount,omitempty"`
}
