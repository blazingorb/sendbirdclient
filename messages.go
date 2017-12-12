package sendbirdclient

type baseMessage struct {
	MessageID  string `json:"message_id"`
	Type       string `json:"type"`
	User       User   `json:"user"`
	CustomType string `json:"custom_type"`
	ChannelURL string `json:"channel_url"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}

type TextMessage struct {
	baseMessage

	Message string `json:"message"`
	Data    string `json:"data"`
	File    File   `json:"file"`
}

type FileMessage struct {
	baseMessage

	File     File   `json:"file"`
	FileName string `json:"file_name"`
	FileSize string `json:"file_size"`
	FileType string `json:"file_type"`
}

type File struct {
	URL  string `json:"url"`
	Data string `json:"data"`
}

type AdminMessage struct {
	baseMessage

	Message string `json:"message"`
	Data    string `json:"data"`
}

type messagesTemplateData struct {
	ChannelType string
	ChannelURL  string
	MessageID   string
}
