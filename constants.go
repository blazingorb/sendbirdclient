package sendbirdclient

//Sendbird list limits
const (
	ListLimitUpperBound = 100
	ListLimitLowerBound = 1
)

//Sendbird url prefix
const (
	constScheme  = "https"
	constHost    = "api.sendbird.com"
	constVersion = "/v3"
)

//Sendbird urls template
const (
	//Users
	SendbirdURLUsers = `/users`

	//Open Channels
	SendbirdURLOpenChannels = `/open_channels`

	//Group Channels
	SendbirdURLGroupChannels = `/group_channels`
)

//Sendbird token type for push service
const (
	TokenTypeGCM  = "GCM"
	TokenTypeAPNS = "APNS"
)

//Sendbird channel type
const (
	ChannelTypeOpenChannel  = "open_channels"
	ChannelTypeGroupChannel = "group_channels"
)
