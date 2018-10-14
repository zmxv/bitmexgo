package bitmexgo

type InstrumentInterval struct {
	Intervals []string `json:"intervals"`
	Symbols   []string `json:"symbols"`
}
