package revoltgo

import (
	"time"

	"github.com/oklog/ulid/v2"
)

// Server struct.
type Server struct {
	Client    *Client
	CreatedAt time.Time

	Id                 string                 `json:"_id"`
	Nonce              string                 `json:"nonce"`
	OwnerId            string                 `json:"owner"`
	Name               string                 `json:"name"`
	Description        string                 `json:"description"`
	ChannelIds         []string               `json:"channels"`
	Categories         []*ServerCategories    `json:"categories"`
	SystemMessages     *SystemMessages        `json:"system_messages"`
	Roles              map[string]interface{} `json:"roles"`
	DefaultPermissions []interface{}          `json:"default_permissions"`
	Icon               *Attachment            `json:"icon"`
	Banner             *Attachment            `json:"banner"`
}

// Server categories struct.
type ServerCategories struct {
	Id         string   `json:"id"`
	Title      string   `json:"title"`
	ChannelIds []string `json:"channels"`
}

// System messages struct.
type SystemMessages struct {
	UserJoined string `json:"user_joined"`
	UserLeft   string `json:"user_left"`
	UserKicked string `json:"user_kicker"`
	UserBanned string `json:"user_banned"`
}

// Calculate creation date and edit the struct.
func (c *Server) CalculateCreationDate() error {
	ulid, err := ulid.Parse(c.Id)

	if err != nil {
		return err
	}

	c.CreatedAt = time.UnixMilli(int64(ulid.Time()))
	return nil
}