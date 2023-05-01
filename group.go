package revgo

import (
	"time"
)

// Group channel struct.
type Group struct {
	Client    *Client
	CreatedAt time.Time
	
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Users       []string `json:"users"`
	Nonce       string   `json:"nonce"`
}