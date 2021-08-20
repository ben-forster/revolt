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
	OnReadyFunction   func()
	OnMessageFunction func(message *Message)
}

// On ready event will run when websocket connection is started and bot is ready to work.
func (c *Client) OnReady(fn func()) {
	c.OnReadyFunction = fn
}

// On message event will run when someone sends a message.
func (c *Client) OnMessage(fn func(message *Message)) {
	c.OnMessageFunction = fn
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
