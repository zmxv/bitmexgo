package bitmexgo

import (
	"time"
)

// Raw Order and Balance Data
type Execution struct {
	ExecID                string    `json:"execID"`
	OrderID               string    `json:"orderID,omitempty"`
	ClOrdID               string    `json:"clOrdID,omitempty"`
	ClOrdLinkID           string    `json:"clOrdLinkID,omitempty"`
	Account               int       `json:"account,omitempty"`
	Symbol                string    `json:"symbol,omitempty"`
	Side                  string    `json:"side,omitempty"`
	LastQty               int       `json:"lastQty,omitempty"`
	LastPx                float64   `json:"lastPx,omitempty"`
	UnderlyingLastPx      float64   `json:"underlyingLastPx,omitempty"`
	LastMkt               string    `json:"lastMkt,omitempty"`
	LastLiquidityInd      string    `json:"lastLiquidityInd,omitempty"`
	SimpleOrderQty        float64   `json:"simpleOrderQty,omitempty"`
	OrderQty              int       `json:"orderQty,omitempty"`
	Price                 float64   `json:"price,omitempty"`
	DisplayQty            int       `json:"displayQty,omitempty"`
	StopPx                float64   `json:"stopPx,omitempty"`
	PegOffsetValue        float64   `json:"pegOffsetValue,omitempty"`
	PegPriceType          string    `json:"pegPriceType,omitempty"`
	Currency              string    `json:"currency,omitempty"`
	SettlCurrency         string    `json:"settlCurrency,omitempty"`
	ExecType              string    `json:"execType,omitempty"`
	OrdType               string    `json:"ordType,omitempty"`
	TimeInForce           string    `json:"timeInForce,omitempty"`
	ExecInst              string    `json:"execInst,omitempty"`
	ContingencyType       string    `json:"contingencyType,omitempty"`
	ExDestination         string    `json:"exDestination,omitempty"`
	OrdStatus             string    `json:"ordStatus,omitempty"`
	Triggered             string    `json:"triggered,omitempty"`
	WorkingIndicator      bool      `json:"workingIndicator,omitempty"`
	OrdRejReason          string    `json:"ordRejReason,omitempty"`
	SimpleLeavesQty       float64   `json:"simpleLeavesQty,omitempty"`
	LeavesQty             int       `json:"leavesQty,omitempty"`
	SimpleCumQty          float64   `json:"simpleCumQty,omitempty"`
	CumQty                int       `json:"cumQty,omitempty"`
	AvgPx                 float64   `json:"avgPx,omitempty"`
	Commission            float64   `json:"commission,omitempty"`
	TradePublishIndicator string    `json:"tradePublishIndicator,omitempty"`
	MultiLegReportingType string    `json:"multiLegReportingType,omitempty"`
	Text                  string    `json:"text,omitempty"`
	TrdMatchID            string    `json:"trdMatchID,omitempty"`
	ExecCost              int       `json:"execCost,omitempty"`
	ExecComm              int       `json:"execComm,omitempty"`
	HomeNotional          float64   `json:"homeNotional,omitempty"`
	ForeignNotional       float64   `json:"foreignNotional,omitempty"`
	TransactTime          time.Time `json:"transactTime,omitempty"`
	Timestamp             time.Time `json:"timestamp,omitempty"`
}
