package bitmexgo

import (
	"time"
)

type AccessToken struct {
	Id string `json:"id"`
	// time to live in seconds (2 weeks by default)
	Ttl     float64   `json:"ttl,omitempty"`
	Created time.Time `json:"created,omitempty"`
	UserId  float64   `json:"userId,omitempty"`
}
