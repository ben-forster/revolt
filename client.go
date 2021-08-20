package revoltgo

import (
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
