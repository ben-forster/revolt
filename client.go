package revoltgo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sacOO7/gowebsocket"
)

const (
	WS_URL  = "wss://ws.revolt.chat"
	API_URL = "https://api.revolt.chat"
)

// Client struct.
type Client struct {
	SelfBot *SelfBot
	Token   string
	Socket  gowebsocket.Socket
	HTTP    *http.Client
	Cache   *Cache

	// Event Functions
	OnReadyFunctions              []func()
	OnMessageFunctions            []func(message *Message)
	OnMessageUpdateFunctions      []func(channel_id, message_id string, payload map[string]interface{})
	OnMessageDeleteFunctions      []func(channel_id, message_id string)
	OnChannelCreateFunctions      []func(channel *Channel)
	OnChannelUpdateFunctions      []func(channel_id, clear string, payload map[string]interface{})
	OnChannelDeleteFunctions      []func(channel_id string)
	OnUnknownEventFunctions       []func(message string)
	OnChannelStartTypingFunctions []func(channel_id, user_id string)
	OnChannelStopTypingFunctions  []func(channel_id, user_id string)
	OnServerUpdateFunctions       []func(server_id, clear string, payload map[string]interface{})
	OnServerDeleteFunctions       []func(server_id string)
}

// Self bot struct.
type SelfBot struct {
	Email        string `json:"-"`
	Password     string `json:"-"`
	Id           string `json:"id"`
	UserId       string `json:"user_id"`
	SessionToken string `json:"session_token"`
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

// On channel create event will run when someone creates a channel.
func (c *Client) OnChannelCreate(fn func(channel *Channel)) {
	c.OnChannelCreateFunctions = append(c.OnChannelCreateFunctions, fn)
}

// On channel update event will run when someone updates a channel.
func (c *Client) OnChannelUpdate(fn func(channel_id, clear string, payload map[string]interface{})) {
	c.OnChannelUpdateFunctions = append(c.OnChannelUpdateFunctions, fn)
}

// On channel delete event will run when someone deletes a channel.
func (c *Client) OnChannelDelete(fn func(channel_id string)) {
	c.OnChannelDeleteFunctions = append(c.OnChannelDeleteFunctions, fn)
}

// On unknown event will run when client gets a unknown event.
func (c *Client) OnUnknownEvent(fn func(message string)) {
	c.OnUnknownEventFunctions = append(c.OnUnknownEventFunctions, fn)
}

// On channel start typing will run when someone starts to type a message.
func (c *Client) OnChannelStartTyping(fn func(channel_id, user_id string)) {
	c.OnChannelStartTypingFunctions = append(c.OnChannelStartTypingFunctions, fn)
}

// On channel stop typing will run when someone stops the typing status.
func (c *Client) OnChannelStopTyping(fn func(channel_id, user_id string)) {
	c.OnChannelStopTypingFunctions = append(c.OnChannelStopTypingFunctions, fn)
}

// On server update will run when someone updates a server.
func (c *Client) OnServerUpdate(fn func(server_id, clear string, payload map[string]interface{})) {
	c.OnServerUpdateFunctions = append(c.OnServerUpdateFunctions, fn)
}

// On server delete will run when someone deletes a server.
func (c *Client) OnServerDelete(fn func(server_id string)) {
	c.OnServerDeleteFunctions = append(c.OnServerDeleteFunctions, fn)
}

// Fetch a channel by Id.
func (c *Client) FetchChannel(id string) (*Channel, error) {
	channel := &Channel{}
	channel.Client = c

	data, err := c.Request("GET", "/channels/"+id, []byte{})

	if err != nil {
		return channel, err
	}

	err = json.Unmarshal(data, channel)
	return channel, err
}

// Fetch an user by Id.
func (c *Client) FetchUser(id string) (*User, error) {
	user := &User{}
	user.Client = c

	data, err := c.Request("GET", "/users/"+id, []byte{})

	if err != nil {
		return user, err
	}

	err = json.Unmarshal(data, user)
	return user, err
}

// Fetch a server by Id.
func (c *Client) FetchServer(id string) (*Server, error) {
	server := &Server{}
	server.Client = c

	data, err := c.Request("GET", "/servers/"+id, []byte{})

	if err != nil {
		return server, err
	}

	err = json.Unmarshal(data, server)
	return server, err
}

// Create a server.
func (c *Client) CreateServer(name, description string) (*Server, error) {
	server := &Server{}
	server.Client = c

	data, err := c.Request("POST", "/servers/create", []byte("{\"name\":\""+name+"\",\"description\":\""+description+"\",\"nonce\":\""+genULID()+"\"}"))

	if err != nil {
		return server, err
	}

	err = json.Unmarshal(data, server)
	return server, err
}

// Auth client user.
func (c *Client) Auth() error {
	if c.SelfBot == nil {
		return fmt.Errorf("can't auth user (not a self-bot.)")
	}

	resp, err := c.Request("POST", "/auth/login", []byte("{\"email\":\""+c.SelfBot.Email+"\",\"password\":\""+c.SelfBot.Password+"\",\"captcha\": \"\"}"))

	if err != nil {
		return err
	}

	err = json.Unmarshal(resp, c.SelfBot)
	return err
}

// Fetch all of the DMs.
func (c *Client) FetchDirectMessages() ([]*Channel, error) {
	var dmChannels []*Channel

	resp, err := c.Request("GET", "/users/dms", []byte{})

	if err != nil {
		return dmChannels, err
	}

	err = json.Unmarshal(resp, &dmChannels)

	if err != nil {
		return dmChannels, err
	}

	// Prepare channels.
	for _, i := range dmChannels {
		i.Client = c
	}

	return dmChannels, nil
}
