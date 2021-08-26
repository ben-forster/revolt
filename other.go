package revoltgo

import "fmt"

// Similar to message, but created for send message function.
type SendMessage struct {
	Content     string   `json:"content,omitempty"`
	Attachments []string `json:"attachments,omitempty"`
	Nonce       string   `json:"nonce,omitempty"`
	DeleteAfter uint     `json:"-"`
	Replies     []struct {
		Id      string `json:"id,omitempty"`
		Mention bool   `json:"mention,omitempty"`
	} `json:"replies,omitempty"`
}

// Set content.
func (sms *SendMessage) SetContent(content string) *SendMessage {
	sms.Content = content
	return sms
}

// Set and format content.
func (sms *SendMessage) SetContentf(format string, values ...interface{}) *SendMessage {
	sms.Content = fmt.Sprintf(format, values...)
	return sms
}

// Set delete after option.
func (sms *SendMessage) SetDeleteAfter(second uint) *SendMessage {
	sms.DeleteAfter = second
	return sms
}

// Add a new attachment.
func (sms *SendMessage) AddAttachment(attachment string) *SendMessage {
	sms.Attachments = append(sms.Attachments, attachment)
	return sms
}

// Add a new reply.
func (sms *SendMessage) AddReply(id string, mention bool) *SendMessage {
	sms.Replies = append(sms.Replies, struct {
		Id      string "json:\"id,omitempty\""
		Mention bool   "json:\"mention,omitempty\""
	}{
		Id:      id,
		Mention: mention,
	})

	return sms
}

// Create a unique nonce.
func (sms *SendMessage) CreateNonce() *SendMessage {
	sms.Nonce = genULID()
	return sms
}

// Edit channel struct.
// Please see: https://developers.revolt.chat/api/#tag/Channel-Information/paths/~1channels~1:channel/patch for more information.
type EditChannel struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Icon        string `json:"icon,omitempty"`
	Remove      string `json:"remove,omitempty"`
}

// Set name for struct.
func (ec *EditChannel) SetName(name string) *EditChannel {
	if len(name) < 1 || len(name) > 32 {
		return ec
	}

	ec.Name = name
	return ec
}

// Set description for struct.
func (ec *EditChannel) SetDescription(desc string) *EditChannel {
	if len(desc) > 1024 {
		return ec
	}

	ec.Description = desc
	return ec
}

// Set icon for struct.
func (ec *EditChannel) SetIcon(autumn_id string) *EditChannel {
	if len(autumn_id) < 1 || len(autumn_id) > 128 {
		return ec
	}

	ec.Icon = autumn_id
	return ec
}

// Set remove item.
func (ec *EditChannel) RemoveItem(item string) *EditChannel {
	if item != "Description" && item != "Icon" {
		return ec
	}

	ec.Remove = item
	return ec
}

// Edit server struct.
// Please see https://developers.revolt.chat/api/#tag/Server-Information/paths/~1servers~1:server/patch for more detail.
type EditServer struct {
	Name           string            `json:"name,omitempty"`
	Description    string            `json:"description,omitempty"`
	Icon           string            `json:"icon,omitempty"`
	Banner         string            `json:"banner,omitempty"`
	Categories     []*ServerCategory `json:"categories,omitempty"`
	SystemMessages *SystemMessages   `json:"system_messages,omitempty"`
	Remove         string            `json:"remove,omitempty"`
}

// Set name for struct
func (es *EditServer) SetName(name string) *EditServer {
	if len(name) < 1 || len(name) > 32 {
		return es
	}

	es.Name = name
	return es
}

// Set description for struct.
func (es *EditServer) SetDescription(desc string) *EditServer {
	if len(desc) > 1024 {
		return es
	}

	es.Description = desc
	return es
}

// Set icon for struct.
func (es *EditServer) SetIcon(autumn_id string) *EditServer {
	if len(autumn_id) < 1 || len(autumn_id) > 128 {
		return es
	}

	es.Icon = autumn_id
	return es
}

// Set banner for struct.
func (es *EditServer) SetBanner(autumn_id string) *EditServer {
	if len(autumn_id) < 1 || len(autumn_id) > 128 {
		return es
	}

	es.Banner = autumn_id
	return es
}

// Add a new category for struct.
func (es *EditServer) AddCategory(category *ServerCategory) *EditServer {
	es.Categories = append(es.Categories, category)
	return es
}

// Set system messages for struct.
func (es *EditServer) SetSystemMessages(sm *SystemMessages) *EditServer {
	es.SystemMessages = sm
	return es
}

// Set remove item.
func (es *EditServer) RemoveItem(item string) *EditServer {
	if item != "Description" && item != "Banner" && item != "Icon" {
		return es
	}

	es.Remove = item
	return es
}

// Edit member struct.
// Please see https://developers.revolt.chat/api/#tag/Server-Members/paths/~1servers~1:server~1members~1:member/patch for more information.
type EditMember struct {
	Nickname string   `json:"nickname,omitempty"`
	Avatar   string   `json:"avatar,omitempty"`
	Roles    []string `json:"roles,omitempty"`
	Remove   string   `json:"remove,omitempty"`
}

// Set nickname for struct.
func (em *EditMember) SetNickname(nick string) *EditMember {
	if len(nick) < 1 || len(nick) > 32 {
		return em
	}

	em.Nickname = nick
	return em
}

// Set avatar for struct.
func (em *EditMember) SetAvatar(autumn_id string) *EditMember {
	if len(autumn_id) < 1 || len(autumn_id) > 128 {
		return em
	}

	em.Avatar = autumn_id
	return em
}

// Add role for struct.
func (em *EditMember) AddRole(role_id string) *EditMember {
	em.Roles = append(em.Roles, role_id)
	return em
}

// Set remove item.
func (em *EditMember) RemoveItem(item string) *EditMember {
	if item != "Avatar" && item != "Nickname" {
		return em
	}

	em.Remove = item
	return em
}

// Edit role struct.
type EditRole struct {
	Name   string `json:"name"`
	Colour string `json:"colour"`
	Hoist  bool   `json:"hoist"`
	Rank   int    `json:"rank"`
	Remove string `json:"remove"`
}

// Set name for struct.
func (er *EditRole) SetName(name string) *EditRole {
	if len(name) < 1 || len(name) > 32 {
		return er
	}

	er.Name = name
	return er
}

// Set valid HTML color for struct.
func (er *EditRole) SetColour(color string) *EditRole {
	if len(color) < 1 || len(color) > 32 {
		return er
	}

	er.Name = color
	return er
}

// Set hoist boolean value for struct.
func (er *EditRole) IsHoist(hoist bool) *EditRole {
	er.Hoist = hoist
	return er
}

// Set role ranking for struct.
func (er *EditRole) SetRank(rank int) *EditRole {
	er.Rank = rank
	return er
}

// Set role ranking for struct.
func (er *EditRole) RemoveColour() *EditRole {
	er.Remove = "Colour"
	return er
}
