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
}
