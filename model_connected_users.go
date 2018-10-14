package bitmexgo

type ConnectedUsers struct {
	Users float32 `json:"users,omitempty"`
	Bots  float32 `json:"bots,omitempty"`
}
