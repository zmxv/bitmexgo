package bitmexgo

import (
	"time"
)

// Trollbox Data
type Chat struct {
	Id        float32   `json:"id,omitempty"`
	Date      time.Time `json:"date"`
	User      string    `json:"user"`
	Message   string    `json:"message"`
	Html      string    `json:"html"`
	FromBot   bool      `json:"fromBot,omitempty"`
	ChannelID float64   `json:"channelID,omitempty"`
}
