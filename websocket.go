package revoltgo

import (
	"fmt"
	"time"

	"github.com/sacOO7/gowebsocket"
)

func (c *Client) Start() {
	// Create new socket
	c.Socket = gowebsocket.New(WS_URL)

	// Send auth when connected
	c.Socket.OnConnected = func(_ gowebsocket.Socket) {
		c.Socket.SendText(fmt.Sprintf("{\"type\": \"Authenticate\", \"token\": \"%s\"}", c.Token))
	}

	c.Socket.OnTextMessage = func(message string, _ gowebsocket.Socket) {
		if message == "{\"type\":\"Authenticated\"}" {
			go c.ping()
		}

		fmt.Println(message)
	}

	// Start connection
	c.Socket.Connect()
}

// Destroy the websocket.
func (c *Client) Destroy() {
	c.Socket.Close()
}

// Ping websocket.
func (c *Client) ping() {
	for {
		time.Sleep(30 * time.Second)
		c.Socket.SendText(fmt.Sprintf("{\"type\":\"Ping\",\"time\":%d}", time.Now().Unix()))
	}
}
