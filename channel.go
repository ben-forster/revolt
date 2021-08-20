package revoltgo

import (
	"encoding/json"
)

// Channel struct.
type Channel struct {
	Client *Client

	Id                 string      `json:"_id"`
	ChannelType        string      `json:"channel_type"`
	UserId             string      `json:"user"`
	Nonce              string      `json:"nonce"`
	Active             bool        `json:"active"`
	Recipients         []string    `json:"recipients"`
	LastMessage        string      `json:"last_message"`
	Name               string      `json:"name"`
	OwnerId            string      `json:"owner"`
	Description        string      `json:"description"`
	Icon               *Attachment `json:"icon"`
	DefaultPermissions int         `json:"default_permissions"`
	RolePermissions    int         `json:"role_permissions"`
	Permissions        int         `json:"permissions"`
}

// Send a message to the channel.
func (c Channel) SendMessage(message *Message) (*Message, error) {
	if message.Nonce == "" {
		message.Nonce = genULID()
	}

	respMessage := &Message{}
	msgData, err := json.Marshal(message)

	if err != nil {
		return respMessage, err
	}

	resp, err := c.Client.Request("POST", "/channels/"+c.Id+"/messages", msgData)

	if err != nil {
		return respMessage, err
	}

	err = json.Unmarshal(resp, respMessage)

	if err != nil {
		return respMessage, err
	}

	message.Client = c.Client
	return message, nil
}
