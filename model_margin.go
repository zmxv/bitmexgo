package bitmexgo

import (
	"time"
)

type Margin struct {
	Account            float32   `json:"account"`
	Currency           string    `json:"currency"`
	RiskLimit          float32   `json:"riskLimit,omitempty"`
	PrevState          string    `json:"prevState,omitempty"`
	State              string    `json:"state,omitempty"`
	Action             string    `json:"action,omitempty"`
	Amount             float32   `json:"amount,omitempty"`
	PendingCredit      float32   `json:"pendingCredit,omitempty"`
	PendingDebit       float32   `json:"pendingDebit,omitempty"`
	ConfirmedDebit     float32   `json:"confirmedDebit,omitempty"`
	PrevRealisedPnl    float32   `json:"prevRealisedPnl,omitempty"`
	PrevUnrealisedPnl  float32   `json:"prevUnrealisedPnl,omitempty"`
	GrossComm          float32   `json:"grossComm,omitempty"`
	GrossOpenCost      float32   `json:"grossOpenCost,omitempty"`
	GrossOpenPremium   float32   `json:"grossOpenPremium,omitempty"`
	GrossExecCost      float32   `json:"grossExecCost,omitempty"`
	GrossMarkValue     float32   `json:"grossMarkValue,omitempty"`
	RiskValue          float32   `json:"riskValue,omitempty"`
	TaxableMargin      float32   `json:"taxableMargin,omitempty"`
	InitMargin         float32   `json:"initMargin,omitempty"`
	MaintMargin        float32   `json:"maintMargin,omitempty"`
	SessionMargin      float32   `json:"sessionMargin,omitempty"`
	TargetExcessMargin float32   `json:"targetExcessMargin,omitempty"`
	VarMargin          float32   `json:"varMargin,omitempty"`
	RealisedPnl        float32   `json:"realisedPnl,omitempty"`
	UnrealisedPnl      float32   `json:"unrealisedPnl,omitempty"`
	IndicativeTax      float32   `json:"indicativeTax,omitempty"`
	UnrealisedProfit   float32   `json:"unrealisedProfit,omitempty"`
	SyntheticMargin    float32   `json:"syntheticMargin,omitempty"`
	WalletBalance      float32   `json:"walletBalance,omitempty"`
	MarginBalance      float32   `json:"marginBalance,omitempty"`
	MarginBalancePcnt  float64   `json:"marginBalancePcnt,omitempty"`
	MarginLeverage     float64   `json:"marginLeverage,omitempty"`
	MarginUsedPcnt     float64   `json:"marginUsedPcnt,omitempty"`
	ExcessMargin       float32   `json:"excessMargin,omitempty"`
	ExcessMarginPcnt   float64   `json:"excessMarginPcnt,omitempty"`
	AvailableMargin    float32   `json:"availableMargin,omitempty"`
	WithdrawableMargin float32   `json:"withdrawableMargin,omitempty"`
	Timestamp          time.Time `json:"timestamp,omitempty"`
	GrossLastValue     float32   `json:"grossLastValue,omitempty"`
	Commission         float64   `json:"commission,omitempty"`
}
