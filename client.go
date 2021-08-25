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
	OnReadyFunctions         []func()
	OnMessageFunctions       []func(message *Message)
	OnMessageUpdateFunctions []func(channel_id, message_id string, payload map[string]interface{})
	OnMessageDeleteFunctions []func(channel_id, message_id string)
	OnChannelCreateFunctions []func(channel *Channel)
	OnChannelUpdateFunctions []func(channel_id, clear string, payload map[string]interface{})
	OnChannelDeleteFunctions []func(channel_id string)
}

// Client cache struct.
type Cache struct {
	Users    []*User    `json:"users"`
	Servers  []*Server  `json:"servers"`
	Channels []*Channel `json:"channels"`
	Members  []*Member  `json:"members"`
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

// Fetch a channel by Id.
func (c *Client) FetchChannel(id string) (*Channel, error) {
	channel := &Channel{}
	channel.Client = c

	data, err := c.Request("GET", "/channels/"+id, []byte{})

	if err != nil {
		return channel, err
	}

	err = json.Unmarshal(data, channel)

	if err != nil {
		return channel, err
	}
	return channel, nil
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

	if err != nil {
		return user, err
	}

	return user, nil
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

	if err != nil {
		return server, err
	}

	return server, nil
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

	if err != nil {
		return server, err
	}

	return server, nil
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

	if err != nil {
		return err
	}

	return nil
}

// Get a channel from cache by Id.
// Will return an empty channel struct if not found.
func (c *Cache) GetChannel(id string) *Channel {
	for _, i := range c.Channels {
		if i.Id == id {
			return i
		}
	}

	return &Channel{}
}

// Get a server from cache by Id.
// Will return an empty server struct if not found.
func (c *Cache) GetServer(id string) *Server {
	for _, i := range c.Servers {
		if i.Id == id {
			return i
		}
	}

	return &Server{}
}

// Get an user from cache by Id.
// Will return an empty user struct if not found.
func (c *Cache) GetUser(id string) *User {
	for _, i := range c.Users {
		if i.Id == id {
			return i
		}
	}

	return &User{}
}

// Get a member from cache by Id.
// Will return an empty member struct if not found.
func (c *Cache) GetMember(id string) *Member {
	for _, i := range c.Members {
		if i.Informations.UserId == id {
			return i
		}
	}

	return &Member{}
}

// Remove a channel from cache by Id.
// Will not delete the channel, just deletes the channel from cache.
// Will change the entire channel cache order!
func (c *Cache) RemoveChannel(id string) error {
	for i, v := range c.Channels {
		if v.Id == id {
			c.Channels[i] = c.Channels[len(c.Channels)-1]
			c.Channels = c.Channels[:len(c.Channels)-1]

			return nil
		}
	}

	return fmt.Errorf("channel not found")
}

// Remove a server from cache by Id.
// Will not delete the server, just deletes the server from cache.
// Will change the entire server cache order!
func (c *Cache) RemoveServer(id string) error {
	for i, v := range c.Servers {
		if v.Id == id {
			c.Servers[i] = c.Servers[len(c.Servers)-1]
			c.Servers = c.Servers[:len(c.Servers)-1]

			return nil
		}
	}

	return fmt.Errorf("server not found")
}

// Remove an user from cache by Id.
// Will not delete the user, just deletes the user from cache.
// Will change the entire user cache order!
func (c *Cache) RemoveUser(id string) error {
	for i, v := range c.Users {
		if v.Id == id {
			c.Users[i] = c.Users[len(c.Users)-1]
			c.Users = c.Users[:len(c.Users)-1]

			return nil
		}
	}

	return fmt.Errorf("user not found")
}

// Remove a member from cache by Id.
// Will not delete the member, just deletes the member from cache.
// Will change the entire member cache order!
func (c *Cache) RemoveMember(id string) error {
	for i, v := range c.Members {
		if v.Informations.UserId == id {
			c.Members[i] = c.Members[len(c.Members)-1]
			c.Members = c.Members[:len(c.Members)-1]

			return nil
		}
	}

	return fmt.Errorf("member not found")
}
