package bitmexgo

import (
	"time"
)

type Affiliate struct {
	Account         int       `json:"account"`
	Currency        string    `json:"currency"`
	PrevPayout      int       `json:"prevPayout,omitempty"`
	PrevTurnover    int       `json:"prevTurnover,omitempty"`
	PrevComm        int       `json:"prevComm,omitempty"`
	PrevTimestamp   time.Time `json:"prevTimestamp,omitempty"`
	ExecTurnover    int       `json:"execTurnover,omitempty"`
	ExecComm        int       `json:"execComm,omitempty"`
	TotalReferrals  int       `json:"totalReferrals,omitempty"`
	TotalTurnover   int       `json:"totalTurnover,omitempty"`
	TotalComm       int       `json:"totalComm,omitempty"`
	PayoutPcnt      float64   `json:"payoutPcnt,omitempty"`
	PendingPayout   int       `json:"pendingPayout,omitempty"`
	Timestamp       time.Time `json:"timestamp,omitempty"`
	ReferrerAccount int       `json:"referrerAccount,omitempty"`
}
