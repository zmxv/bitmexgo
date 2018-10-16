package bitmexgo

import (
	"time"
)

// Account Operations
type User struct {
	Id           int              `json:"id,omitempty"`
	OwnerId      int              `json:"ownerId,omitempty"`
	Firstname    string           `json:"firstname,omitempty"`
	Lastname     string           `json:"lastname,omitempty"`
	Username     string           `json:"username"`
	Email        string           `json:"email"`
	Phone        string           `json:"phone,omitempty"`
	Created      time.Time        `json:"created,omitempty"`
	LastUpdated  time.Time        `json:"lastUpdated,omitempty"`
	Preferences  *UserPreferences `json:"preferences,omitempty"`
	TFAEnabled   string           `json:"TFAEnabled,omitempty"`
	AffiliateID  string           `json:"affiliateID,omitempty"`
	PgpPubKey    string           `json:"pgpPubKey,omitempty"`
	Country      string           `json:"country,omitempty"`
	GeoipCountry string           `json:"geoipCountry,omitempty"`
	GeoipRegion  string           `json:"geoipRegion,omitempty"`
	Typ          string           `json:"typ,omitempty"`
}
