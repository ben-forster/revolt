package revoltgo

import (
	"encoding/json"
	"fmt"
	"reflect"
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

// Similar to message, but created for sendmessage function.
type SendMessageStruct struct {
	Content     string   `json:"content"`
	Attachments []string `json:"attachments"`
	Nonce       string   `json:"nonce"`
	Replies     []struct {
		Id      string `json:"id"`
		Mention bool   `json:"mention"`
	}
}

// Fetched messages struct.
type FetchedMessages struct {
	Messages []*Message `json:"messages"`
	Users    []*User    `json:"users"`
}

// Send a message to the channel.
func (c Channel) SendMessage(message *SendMessageStruct) (*Message, error) {
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

	respMessage.Client = c.Client
	return respMessage, nil
}

// Fetch messages from channel.
// Check: https://developers.revolt.chat/api/#tag/Messaging/paths/~1channels~1:channel~1messages/get for map parameters.
func (c Channel) FetchMessages(options map[string]interface{}) (*FetchedMessages, error) {
	// Format url
	url := "/channels/" + c.Id + "/messages?"

	for key, value := range options {
		if !reflect.ValueOf(value).IsZero() {
			url += fmt.Sprintf("%s=%v&", key, value)
		}
	}

	url = url[:len(url)-1]

	fetchedMsgs := &FetchedMessages{}

	// Send request
	resp, err := c.Client.Request("GET", url, []byte{})

	if err != nil {
		return fetchedMsgs, err
	}

	err = json.Unmarshal(resp, &fetchedMsgs)

	if err != nil {
		err = json.Unmarshal([]byte(fmt.Sprintf("{\"messages\": %s}", resp)), &fetchedMsgs)

		if err != nil {
			return fetchedMsgs, err
		}
	}

	// Add client to users & messages
	for _, msg := range fetchedMsgs.Messages {
		msg.Client = c.Client
	}

	if fetchedMsgs.Users != nil {
		for _, msg := range fetchedMsgs.Users {
			msg.Client = c.Client
		}
	}

	return fetchedMsgs, nil
}
