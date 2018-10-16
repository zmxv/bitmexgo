package bitmexgo

import (
	"time"
)

type Wallet struct {
	Account          int       `json:"account"`
	Currency         string    `json:"currency"`
	PrevDeposited    int       `json:"prevDeposited,omitempty"`
	PrevWithdrawn    int       `json:"prevWithdrawn,omitempty"`
	PrevTransferIn   int       `json:"prevTransferIn,omitempty"`
	PrevTransferOut  int       `json:"prevTransferOut,omitempty"`
	PrevAmount       int       `json:"prevAmount,omitempty"`
	PrevTimestamp    time.Time `json:"prevTimestamp,omitempty"`
	DeltaDeposited   int       `json:"deltaDeposited,omitempty"`
	DeltaWithdrawn   int       `json:"deltaWithdrawn,omitempty"`
	DeltaTransferIn  int       `json:"deltaTransferIn,omitempty"`
	DeltaTransferOut int       `json:"deltaTransferOut,omitempty"`
	DeltaAmount      int       `json:"deltaAmount,omitempty"`
	Deposited        int       `json:"deposited,omitempty"`
	Withdrawn        int       `json:"withdrawn,omitempty"`
	TransferIn       int       `json:"transferIn,omitempty"`
	TransferOut      int       `json:"transferOut,omitempty"`
	Amount           int       `json:"amount,omitempty"`
	PendingCredit    int       `json:"pendingCredit,omitempty"`
	PendingDebit     int       `json:"pendingDebit,omitempty"`
	ConfirmedDebit   int       `json:"confirmedDebit,omitempty"`
	Timestamp        time.Time `json:"timestamp,omitempty"`
	Addr             string    `json:"addr,omitempty"`
	Script           string    `json:"script,omitempty"`
	WithdrawalLock   []string  `json:"withdrawalLock,omitempty"`
}
