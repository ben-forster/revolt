package revoltgo

// Bot struct.
type Bot struct {
	Id              string `json:"_id"`
	OwnerId         string `json:"owner"`
	Token           string `json:"token"`
	IsPublic        bool   `json:"public"`
	InteractionsUrl string `json:"interactionsURL"`
}
