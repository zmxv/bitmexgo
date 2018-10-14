package bitmexgo

// Information on Top Users
type Leaderboard struct {
	Name       string  `json:"name"`
	IsRealName bool    `json:"isRealName,omitempty"`
	Profit     float64 `json:"profit,omitempty"`
}
