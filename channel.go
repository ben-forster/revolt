package revoltgo

// Channel struct.
type Channel struct {
	Client *Client

	Id                 string            `json:"_id"`
	ChannelType        string            `json:"channel_type"`
	UserId             string            `json:"user"`
	Nonce              string            `json:"nonce"`
	Active             bool              `json:"active"`
	Recipients         []string          `json:"recipients"`
	LastMessage        map[string]string `json:"last_message"`
	Name               string            `json:"name"`
	OwnerId            string            `json:"owner"`
	Description        string            `json:"description"`
	Icon               *Attachment       `json:"icon"`
	DefaultPermissions int               `json:"default_permissions"`
	RolePermissions    int               `json:"role_permissions"`
	Permissions        int               `json:"permissions"`
}
