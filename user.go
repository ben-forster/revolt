package revoltgo

import (
	"time"

	"github.com/oklog/ulid/v2"
)

// User struct.
type User struct {
	Client    *Client
	CreatedAt time.Time

	Id             string           `json:"_id"`
	Username       string           `json:"username"`
	Avatar         *Attachment      `json:"avatar"`
	Relations      []*UserRelations `json:"relations"`
	Badges         int              `json:"badges"`
	Status         *UserStatus      `json:"status"`
	Relationship   string           `json:"relationship"`
	IsOnline       bool             `json:"online"`
	Flags          int              `json:"flags"`
	BotInformation *BotInformation  `json:"bot"`
}

// User relations struct.
type UserRelations struct {
	Id     string `json:"_id"`
	Status string `json:"status"`
}

// User status struct.
type UserStatus struct {
	Text     string `json:"text"`
	Presence string `json:"presence"`
}

// Bot information struct.
type BotInformation struct {
	Owner string `json:"owner"`
}

// Calculate creation date and edit the struct.
func (u *User) CalculateCreationDate() error {
	ulid, err := ulid.Parse(u.Id)

	if err != nil {
		return err
	}

	u.CreatedAt = time.UnixMilli(int64(ulid.Time()))
	return nil
}

// Create a mention format.
func (u User) FormatMention() string {
	return "<@" + u.Id + ">"
}
