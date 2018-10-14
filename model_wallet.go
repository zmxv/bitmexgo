package bitmexgo

import (
	"time"
)

type Wallet struct {
	Account          float32   `json:"account"`
	Currency         string    `json:"currency"`
	PrevDeposited    float32   `json:"prevDeposited,omitempty"`
	PrevWithdrawn    float32   `json:"prevWithdrawn,omitempty"`
	PrevTransferIn   float32   `json:"prevTransferIn,omitempty"`
	PrevTransferOut  float32   `json:"prevTransferOut,omitempty"`
	PrevAmount       float32   `json:"prevAmount,omitempty"`
	PrevTimestamp    time.Time `json:"prevTimestamp,omitempty"`
	DeltaDeposited   float32   `json:"deltaDeposited,omitempty"`
	DeltaWithdrawn   float32   `json:"deltaWithdrawn,omitempty"`
	DeltaTransferIn  float32   `json:"deltaTransferIn,omitempty"`
	DeltaTransferOut float32   `json:"deltaTransferOut,omitempty"`
	DeltaAmount      float32   `json:"deltaAmount,omitempty"`
	Deposited        float32   `json:"deposited,omitempty"`
	Withdrawn        float32   `json:"withdrawn,omitempty"`
	TransferIn       float32   `json:"transferIn,omitempty"`
	TransferOut      float32   `json:"transferOut,omitempty"`
	Amount           float32   `json:"amount,omitempty"`
	PendingCredit    float32   `json:"pendingCredit,omitempty"`
	PendingDebit     float32   `json:"pendingDebit,omitempty"`
	ConfirmedDebit   float32   `json:"confirmedDebit,omitempty"`
	Timestamp        time.Time `json:"timestamp,omitempty"`
	Addr             string    `json:"addr,omitempty"`
	Script           string    `json:"script,omitempty"`
	WithdrawalLock   []string  `json:"withdrawalLock,omitempty"`
}
