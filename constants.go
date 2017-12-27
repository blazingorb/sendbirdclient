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

//Sendbird Webhook Payload Category
const (
	WebhookCategoryOpenChannelMsgSend        = "open_channel:message_send"
	WebhookCategoryGroupChannelMsgSend       = "group_channel:message_send"
	WebhookCategoryOpenChannelMsgDeleted     = "open_channel:message_delete"
	WebhookCategoryGroupChannelMsgDeleted    = "group_channel:message_delete"
	WebhookCategoryGroupChannelMsgRead       = "group_channel:message_read"
	WebhookCategoryOpenChannelCreated        = "open_channel:create"
	WebhookCategoryGroupChannelCreated       = "group_channel:create"
	WebhookCategoryOpenChannelRemoved        = "open_channel:remove"
	WebhookCategoryGroupChannelInvited       = "group_channel:invite"
	WebhookCategoryGroupChannelJoined        = "group_channel:join"
	WebhookCategoryGroupChannelDeclineInvite = "group_channel:decline_invite"
	WebhookCategoryUserBlocked               = "user:block"
	WebhookCategoryUserUnblocked             = "user:unblock"
	WebhookCategoryUserMsgRateLimitExceeded  = "alert:user_message_rate_limit_exceeded	"
)
