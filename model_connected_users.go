package bitmexgo

type ConnectedUsers struct {
	Users int `json:"users,omitempty"`
	Bots  int `json:"bots,omitempty"`
}
