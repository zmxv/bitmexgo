package bitmexgo

import (
	"time"
)

// Account Notifications
type Notification struct {
	Id                float32   `json:"id,omitempty"`
	Date              time.Time `json:"date"`
	Title             string    `json:"title"`
	Body              string    `json:"body"`
	Ttl               float32   `json:"ttl"`
	Type_             string    `json:"type,omitempty"`
	Closable          bool      `json:"closable,omitempty"`
	Persist           bool      `json:"persist,omitempty"`
	WaitForVisibility bool      `json:"waitForVisibility,omitempty"`
	Sound             string    `json:"sound,omitempty"`
}
