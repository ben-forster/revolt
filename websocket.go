package revoltgo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sacOO7/gowebsocket"
)

func (c *Client) Start() {
	// Create new socket
	c.Socket = gowebsocket.New(WS_URL)
	c.HTTP = &http.Client{}

	// Auth the user if self-bot.
	// if c.SelfBot != nil {
	// 	c.Auth()
	// }

	// Send auth when connected
	c.Socket.OnConnected = func(_ gowebsocket.Socket) {
		c.handleWebsocketAuth()
	}

	c.Socket.OnTextMessage = func(message string, _ gowebsocket.Socket) {
		// Parse data
		rawData := &struct {
			Type string `json:"type"`
		}{}
		err := json.Unmarshal([]byte(message), rawData)

		if err != nil {
			c.Destroy()
			panic(err)
		}

		if rawData.Type == "Authenticated" {
			go c.ping()
		}

		// Handle events
		c.handleEvents(rawData, message)
		// fmt.Println(message)
	}

	// Start connection
	c.Socket.Connect()
}

// Handle on connected.
func (c *Client) handleWebsocketAuth() {
	if c.SelfBot == nil {
		c.Socket.SendText(fmt.Sprintf("{\"type\":\"Authenticate\",\"token\":\"%s\"}", c.Token))
	} else {
		c.Socket.SendText(fmt.Sprintf("{\"type\":\"Authenticate\",\"id\":\"%s\",\"session_token\":\"%s\",\"user_id\":\"%s\"}", c.SelfBot.Id, c.SelfBot.SessionToken, c.SelfBot.UserId))
	}
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

// Handle events.
func (c *Client) handleEvents(rawData *struct {
	Type string `json:"type"`
}, message string) {
	if rawData.Type == "Ready" {
		// Add cache
		c.handleCache(message)

		// Ready Event
		if c.OnReadyFunctions != nil {
			for _, i := range c.OnReadyFunctions {
				i()
			}
		}
	} else if rawData.Type == "Message" && c.OnMessageFunctions != nil {
		// Message Event
		msgData := &Message{}
		err := json.Unmarshal([]byte(message), msgData)

		if err != nil {
			fmt.Printf("Unexcepted Error: %s", err)
		}

		msgData.Client = c

		for _, i := range c.OnMessageFunctions {
			i(msgData)
		}
	} else if rawData.Type == "MessageUpdate" && c.OnMessageUpdateFunctions != nil {
		// Message Update Event
		data := &struct {
			ChannelId string                 `json:"channel"`
			MessageId string                 `json:"id"`
			Payload   map[string]interface{} `json:"data"`
		}{}

		err := json.Unmarshal([]byte(message), data)

		if err != nil {
			fmt.Printf("Unexcepted Error: %s", err)
		}

		for _, i := range c.OnMessageUpdateFunctions {
			i(data.ChannelId, data.MessageId, data.Payload)
		}
	} else if rawData.Type == "MessageDelete" && c.OnMessageDeleteFunctions != nil {
		// Message Delete Event
		data := &struct {
			ChannelId string `json:"channel"`
			MessageId string `json:"id"`
		}{}

		err := json.Unmarshal([]byte(message), data)

		if err != nil {
			fmt.Printf("Unexcepted Error: %s", err)
		}

		for _, i := range c.OnMessageDeleteFunctions {
			i(data.ChannelId, data.MessageId)
		}
	} else if rawData.Type == "ChannelCreate" && c.OnChannelCreateFunctions != nil {
		// Channel create event.
		channelData := &Channel{}

		err := json.Unmarshal([]byte(message), channelData)

		if err != nil {
			fmt.Printf("Unexcepted Error: %s", err)
		}

		channelData.Client = c

		for _, i := range c.OnChannelCreateFunctions {
			i(channelData)
		}
	} else if rawData.Type == "ChannelUpdate" && c.OnChannelUpdateFunctions != nil {
		// Channel update event.
		data := &struct {
			ChannelId string                 `json:"id"`
			Clear     string                 `json:"clear"`
			Payload   map[string]interface{} `json:"data"`
		}{}

		err := json.Unmarshal([]byte(message), data)

		if err != nil {
			fmt.Printf("Unexcepted Error: %s", err)
		}

		for _, i := range c.OnChannelUpdateFunctions {
			i(data.ChannelId, data.Clear, data.Payload)
		}
	} else if rawData.Type == "ChannelDelete" && c.OnChannelDeleteFunctions != nil {
		// Channel delete event.
		data := &struct {
			ChannelId string `json:"id"`
		}{}

		err := json.Unmarshal([]byte(message), data)

		if err != nil {
			fmt.Printf("Unexcepted Error: %s", err)
		}

		for _, i := range c.OnChannelDeleteFunctions {
			i(data.ChannelId)
		}
	}
}

func (c *Client) handleCache(data string) {
	cache := &Cache{}

	err := json.Unmarshal([]byte(data), cache)

	if err != nil {
		fmt.Printf("Unexcepted Error: %s", err)
	}

	// Add client to users.
	for _, i := range cache.Users {
		i.Client = c
	}

	for _, i := range cache.Servers {
		i.Client = c
	}

	for _, i := range cache.Channels {
		i.Client = c
	}

	c.Cache = cache
}
