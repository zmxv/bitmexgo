package bitmexgo

import (
	"time"
)

// Public Announcements
type Announcement struct {
	Id      float32   `json:"id"`
	Link    string    `json:"link,omitempty"`
	Title   string    `json:"title,omitempty"`
	Content string    `json:"content,omitempty"`
	Date    time.Time `json:"date,omitempty"`
}
