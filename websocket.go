package revoltgo

import (
	"fmt"

	"github.com/sacOO7/gowebsocket"
)

const WS_URL = "wss://ws.revolt.chat"

func (c *Client) Start() {
	// Create new socket
	c.Socket = gowebsocket.New(WS_URL)

	// Send auth when connected
	c.Socket.OnConnected = func(socket gowebsocket.Socket) {
		c.Socket.SendText(fmt.Sprintf("{\"type\": \"Authenticate\", \"token\": \"%s\"}", c.Token))
	}

	// Start connection
	c.Socket.Connect()
}
