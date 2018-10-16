package bitmexgo

import (
	"time"
)

type Margin struct {
	Account            int       `json:"account"`
	Currency           string    `json:"currency"`
	RiskLimit          int       `json:"riskLimit,omitempty"`
	PrevState          string    `json:"prevState,omitempty"`
	State              string    `json:"state,omitempty"`
	Action             string    `json:"action,omitempty"`
	Amount             int       `json:"amount,omitempty"`
	PendingCredit      int       `json:"pendingCredit,omitempty"`
	PendingDebit       int       `json:"pendingDebit,omitempty"`
	ConfirmedDebit     int       `json:"confirmedDebit,omitempty"`
	PrevRealisedPnl    int       `json:"prevRealisedPnl,omitempty"`
	PrevUnrealisedPnl  int       `json:"prevUnrealisedPnl,omitempty"`
	GrossComm          int       `json:"grossComm,omitempty"`
	GrossOpenCost      int       `json:"grossOpenCost,omitempty"`
	GrossOpenPremium   int       `json:"grossOpenPremium,omitempty"`
	GrossExecCost      int       `json:"grossExecCost,omitempty"`
	GrossMarkValue     int       `json:"grossMarkValue,omitempty"`
	RiskValue          int       `json:"riskValue,omitempty"`
	TaxableMargin      int       `json:"taxableMargin,omitempty"`
	InitMargin         int       `json:"initMargin,omitempty"`
	MaintMargin        int       `json:"maintMargin,omitempty"`
	SessionMargin      int       `json:"sessionMargin,omitempty"`
	TargetExcessMargin int       `json:"targetExcessMargin,omitempty"`
	VarMargin          int       `json:"varMargin,omitempty"`
	RealisedPnl        int       `json:"realisedPnl,omitempty"`
	UnrealisedPnl      int       `json:"unrealisedPnl,omitempty"`
	IndicativeTax      int       `json:"indicativeTax,omitempty"`
	UnrealisedProfit   int       `json:"unrealisedProfit,omitempty"`
	SyntheticMargin    int       `json:"syntheticMargin,omitempty"`
	WalletBalance      int       `json:"walletBalance,omitempty"`
	MarginBalance      int       `json:"marginBalance,omitempty"`
	MarginBalancePcnt  float64   `json:"marginBalancePcnt,omitempty"`
	MarginLeverage     float64   `json:"marginLeverage,omitempty"`
	MarginUsedPcnt     float64   `json:"marginUsedPcnt,omitempty"`
	ExcessMargin       int       `json:"excessMargin,omitempty"`
	ExcessMarginPcnt   float64   `json:"excessMarginPcnt,omitempty"`
	AvailableMargin    int       `json:"availableMargin,omitempty"`
	WithdrawableMargin int       `json:"withdrawableMargin,omitempty"`
	Timestamp          time.Time `json:"timestamp,omitempty"`
	GrossLastValue     int       `json:"grossLastValue,omitempty"`
	Commission         float64   `json:"commission,omitempty"`
}
