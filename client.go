package revoltgo

import (
	"encoding/json"
	"net/http"

	"github.com/sacOO7/gowebsocket"
)

const (
	WS_URL  = "wss://ws.revolt.chat"
	API_URL = "https://api.revolt.chat"
)

// Client struct
type Client struct {
	Token  string
	Socket gowebsocket.Socket
	HTTP   *http.Client

	// Event Functions
	OnReadyFunctions         []func()
	OnMessageFunctions       []func(message *Message)
	OnMessageUpdateFunctions []func(channel_id, message_id string, payload map[string]interface{})
	OnMessageDeleteFunctions []func(channel_id, message_id string)
}

// On ready event will run when websocket connection is started and bot is ready to work.
func (c *Client) OnReady(fn func()) {
	c.OnReadyFunctions = append(c.OnReadyFunctions, fn)
}

// On message event will run when someone sends a message.
func (c *Client) OnMessage(fn func(message *Message)) {
	c.OnMessageFunctions = append(c.OnMessageFunctions, fn)
}

// On message update event will run when someone updates a message.
func (c *Client) OnMessageUpdate(fn func(channel_id, message_id string, payload map[string]interface{})) {
	c.OnMessageUpdateFunctions = append(c.OnMessageUpdateFunctions, fn)
}

// On message delete event will run when someone deletes a message.
func (c *Client) OnMessageDelete(fn func(channel_id, message_id string)) {
	c.OnMessageDeleteFunctions = append(c.OnMessageDeleteFunctions, fn)
}

// Fetch a channel by Id.
func (c Client) FetchChannel(id string) (*Channel, error) {
	channel := &Channel{}

	data, err := c.Request("GET", "/channels/"+id, []byte{})

	if err != nil {
		return channel, err
	}

	err = json.Unmarshal(data, channel)

	if err != nil {
		return channel, err
	}

	channel.Client = &c
	return channel, nil
}

// Fetch an user by Id.
func (c Client) FetchUser(id string) (*User, error) {
	user := &User{}

	data, err := c.Request("GET", "/users/"+id, []byte{})

	if err != nil {
		return user, err
	}

	err = json.Unmarshal(data, user)

	if err != nil {
		return user, err
	}

	user.Client = &c
	return user, nil
}
