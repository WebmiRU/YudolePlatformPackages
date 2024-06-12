package message

type StreamChat struct {
	//Id      string            `json:"id"`
	Type    string            `json:"type"`
	Service string            `json:"service"`
	Module  string            `json:"module"`
	Payload StreamChatPayload `json:"payload"`
}

type StreamChatPayload struct {
	Id      string            `json:"id"`
	User    StreamChatUser    `json:"user"`
	Channel string            `json:"channel"`
	Html    string            `json:"html"`
	Text    string            `json:"text"`
	Tags    map[string]string `json:"tags"`
}

type StreamChatUser struct {
	Login  string   `json:"login"`
	Nick   string   `json:"nick"`
	Badges []string `json:"badges"`
}
