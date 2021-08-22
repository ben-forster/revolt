package revoltgo

// Similar to message, but created for send message function.
type SendMessage struct {
	Content     string   `json:"content"`
	Attachments []string `json:"attachments"`
	Nonce       string   `json:"nonce"`
	Replies     []struct {
		Id      string `json:"id"`
		Mention bool   `json:"mention"`
	}
}

// Set content.
func (sms *SendMessage) SetContent(content string) *SendMessage {
	sms.Content = content
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
		Id      string "json:\"id\""
		Mention bool   "json:\"mention\""
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
	ec.Name = name
	return ec
}

// Set description for struct.
func (ec *EditChannel) SetDescription(desc string) *EditChannel {
	ec.Description = desc
	return ec
}

// Set icon for struct.
func (ec *EditChannel) SetIcon(autumn_id string) *EditChannel {
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
