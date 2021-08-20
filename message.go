package revoltgo

// Message struct
type Message struct {
	Client *Client

	Id          string        `json:"_id"`
	Nonce       string        `json:"nonce"`
	ChannelId   string        `json:"channel"`
	AuthorId    string        `json:"author"`
	Content     string        `json:"content"`
	Attachments []*Attachment `json:"attachments"`
	Mentions    []string      `json:"mentions"`
	Replies     []string      `json:"replies"`
}

// Attachment struct.
type Attachment struct {
	Id          string `json:"_id"`
	Tag         string `json:"tag"`
	Size        int    `json:"size"`
	FileName    string `json:"filename"`
	Metadata    *AttachmentMetadata
	ContentType string `json:"content_type"`
	Embeds      []*MessageEmbed
}

// Attachment metadata struct.
type AttachmentMetadata struct {
	Type   string `json:"type"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// Message edited struct.
type MessageEdited struct {
	Date int `json:"$date"`
}

// Message embed struct.
type MessageEmbed struct {
	Type        string `json:"type"`
	Url         string `json:"url"`
	Special     *MessageSpecialEmbed
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Image       *MessageEmbeddedImage `json:"image"`
	Video       *MessageEmbeddedVideo `json:"video"`
	IconUrl     string                `json:"icon_url"`
	Color       string                `json:"color"`
}

// Message special embed struct.
type MessageSpecialEmbed struct {
	Type        string `json:"type"`
	Id          string `json:"id"`
	ContentType string `json:"content_type"`
}

// Message embedded image struct
type MessageEmbeddedImage struct {
	Size   string `json:"size"`
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// Message embedded video struct
type MessageEmbeddedVideo struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
