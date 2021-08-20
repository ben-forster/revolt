package revoltgo

import "github.com/sacOO7/gowebsocket"

// Client struct
type Client struct {
	Token  string
	Socket gowebsocket.Socket
}
