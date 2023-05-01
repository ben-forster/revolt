package revgo

import (
	"time"
)

// Group channel struct.
type Group struct {
	Client    *Client
	CreatedAt time.Time

	Id                 string      `json:"_id"`
	Nonce              string      `json:"nonce"`
	OwnerId            string      `json:"owner"`	
	Name               string      `json:"name"`
	Description        string      `json:"description"`		
	Active             bool        `json:"active"`
	Recipients         []string    `json:"recipients"`
	LastMessage        interface{} `json:"last_message"`
	Icon               *Attachment `json:"icon"`
	Banner             *Attachment            `json:"banner"`
}