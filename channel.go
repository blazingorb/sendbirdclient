package sendbirdclient

type channel struct {
	Name             string `json:"name"`
	ChannelURL       string `json:"channel_url"`
	CoverURL         string `json:"cover_url"`
	CoverFile        []byte `json:"cover_file"`
	Data             string `json:"data"`
	CustomType       string `json:"custom_type"`
	CreatedAt        int64  `json:"created_at"`
	MaxLengthMessage int    `json:"max_length_message"`
}
