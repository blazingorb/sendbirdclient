package templates

type templateKeysUsers string
type templateKeysOpenChannels string
type templateKeysGroupChannels string
type templateKeysChannelMetadata string
type templateKeysMessages string

//Nested Template String
const (
	SendbirdURLUsersTemplate = `
	{{define "base"}}/users/{{.UserID}}{{end}}
	{{define "unreadCount"}}{{template "base" .}}/unread_count{{end}}
	{{define "block"}}{{template "base" .}}/block{{end}}
	{{define "blockWithTargetID"}}{{template "block" .}}/{{.TargetID}}{{end}}
	{{define "ban"}}{{template "base" .}}/ban{{end}}
	{{define "mute"}}{{template "base" .}}/mute{{end}}
	{{define "markReadAll"}}{{template "base" .}}/mark_as_read_all{{end}}
	{{define "listGroupChannel"}}{{template "base" .}}/my_group_channels{{end}}
	{{define "deviceToken"}}{{template "base" .}}/push{{end}}
	{{define "deviceTokenWithType"}}{{template "deviceToken" .}}/{{.TokenType}}{{end}}
	{{define "deviceTokenWithTypeAndPushToken"}}{{template "deviceTokenWithType" .}}/{{.PushToken}}{{end}}
	{{define "pushPreference"}}{{template "base" .}}/push_preference{{end}}
	{{define "pushPreferenceWithChannelURL"}}{{template "pushPreference" .}}/{{.ChannelURL}}{{end}}
	{{define "userMeta"}}{{template "base" .}}/metadata{{end}}
	{{define "userMetaAndKeyName"}}{{template "userMeta" .}}/{{.KeyName}}{{end}}

	{{template "base"}}
	{{template "unreadCount"}}
	{{template "block"}}
	{{template "blockWithTargetID"}}
	{{template "ban"}}
	{{template "mute"}}
	{{template "markReadAll"}}
	{{template "listGroupChannel"}}
	{{template "deviceToken"}}
	{{template "deviceTokenWithType"}}
	{{template "deviceTokenWithTypeAndPushToken"}}
	{{template "pushPreference"}}
	{{template "pushPreferenceWithChannelURL"}}
	{{template "userMeta"}}
	{{template "userMetaAndKeyName"}}
	`

	SendbirdURLOpenChannelsTemplate = `
	{{define "base"}}/open_channels/{{.ChannelURL}}{{end}}
	{{define "participants"}}{{template "base" .}}/participants{{end}}
	{{define "freeze"}}{{template "base" .}}/freeze{{end}}
	{{define "ban"}}{{template "base" .}}/ban{{end}}
	{{define "banWithUserID"}}{{template "ban" .}}/{{.BannedUserID}}{{end}}
	{{define "mute"}}{{template "base" .}}/mute{{end}}
	{{define "muteWithUserID"}}{{template "mute" .}}/{{.MutedUserID}}{{end}}

	{{template "base"}}
	{{template "participants"}}
	{{template "freeze"}}
	{{template "ban"}}
	{{template "banWithUserID"}}
	{{template "mute"}}
	{{template "muteWithUserID"}}
	`

	SendbirdURLGroupChannelsTemplate = `
	{{define "base"}}/group_channels/{{.ChannelURL}}{{end}}
	{{define "members"}}{{template "base" .}}/members{{end}}
	{{define "membersWithUserID"}}{{template "members" .}}/{{.UserID}}{{end}}
	{{define "invite"}}{{template "base" .}}/invite{{end}}
	{{define "hide"}}{{template "base" .}}/hide{{end}}
	{{define "leave"}}{{template "base" .}}/leave{{end}}
	{{define "freeze"}}{{template "base" .}}/freeze{{end}}

	{{template "base"}}
	{{template "members"}}
	{{template "membersWithUserID"}}
	{{template "invite"}}
	{{template "hide"}}
	{{template "leave"}}
	{{template "freeze"}}
	`

	SendbirdURLChannelMetadataTemplate = `
	{{define "base"}}/{{.ChannelType}}/{{.ChannelURL}}{{end}}
	{{define "chMeta"}}{{template "base" .}}/metadata{{end}}
	{{define "chMetaWithKeyName"}}{{template "chMeta" .}}/{{.KeyName}}{{end}}
	{{define "metacounter"}}{{template "base" .}}/metacounter{{end}}
	{{define "metacounterWithKeyName"}}{{template "metacounter" .}}/{{.KeyName}}{{end}}

	{{template "base"}}
	{{template "chMeta"}}
	{{template "chMetaWithKeyName"}}
	{{template "metacounter"}}
	{{template "metacounterWithKeyName"}}
	`

	SendbirdURLMessagesTemplate = `
	{{define "base"}}/{{.ChannelType}}/{{.ChannelURL}}{{end}}
	{{define "messages"}}{{template "base" .}}/messages{{end}}
	{{define "markAsRead"}}{{template "messages" .}}/mark_as_read{{end}}
	{{define "totalCount"}}{{template "messages" .}}/total_count{{end}}
	{{define "unreadCount"}}{{template "messages" .}}/unread_count{{end}}
	{{define "singleMessage"}}{{template "messages" .}}/{{.MessageID}}{{end}}

	{{template "base"}}
	{{template "messages"}}
	{{template "markAsRead"}}
	{{template "totalCount"}}
	{{template "unreadCount"}}
	{{template "singleMessage"}}
	`
)

//TemplateKeys
const (
	//Users
	SendbirdURLUserswithUserID                                    templateKeysUsers = `base`
	SendbirdURLUsersUnreadCountWithUserID                         templateKeysUsers = `unreadCount`
	SendbirdURLUsersBlockWithUserID                               templateKeysUsers = `block`
	SendbirdURLUsersBlockWithUserIDandTargetID                    templateKeysUsers = `blockWithTargetID`
	SendbirdURLUsersBanWithUserID                                 templateKeysUsers = `ban`
	SendbirdURLUsersMuteWithUserID                                templateKeysUsers = `mute`
	SendbirdURLUsersMarkReadAllWithUserID                         templateKeysUsers = `markReadAll`
	SendbirdURLUsersListGroupChannelsWithUserID                   templateKeysUsers = `listGroupChannel`
	SendbirdURLUsersDeviceTokenWithUserID                         templateKeysUsers = `deviceToken`
	SendbirdURLUsersDeviceTokenWithUserIDandTokenType             templateKeysUsers = `deviceTokenWithType`
	SendbirdURLUsersDeviceTokenWithUserIDandTokenTypeandPushToken templateKeysUsers = `deviceTokenWithTypeAndPushToken`
	SendbirdURLUsersPushPreferenceWithUserID                      templateKeysUsers = `pushPreference`
	SendbirdURLUsersPushPreferenceWithUserIDandChannelURL         templateKeysUsers = `pushPreferenceWithChannelURL`

	//User Metadata
	SendbirdURLUserMetadataWithUserID           templateKeysUsers = `userMeta`
	SendbirdURLUserMetadataWithUserIDandKeyName templateKeysUsers = `userMetaAndKeyName`

	//Open Channels
	SendbirdURLOpenChannelsWithChannelURL                   templateKeysOpenChannels = `base`
	SendbirdURLOpenChannelsParticipantsWithChannelURL       templateKeysOpenChannels = `participants`
	SendbirdURLOpenChannelsFreezeWithChannelURL             templateKeysOpenChannels = `freeze`
	SendbirdURLOpenChannelsBanWithChannelURL                templateKeysOpenChannels = `ban`
	SendbirdURLOpenChannelsBanWithChannelURLandBannedUserID templateKeysOpenChannels = `banWithUserID`
	SendbirdURLOpenChannelsMuteWithChannelURL               templateKeysOpenChannels = `mute`
	SendbirdURLOpenChannelsMuteWithChannelURLandMutedUserID templateKeysOpenChannels = `muteWithUserID`

	//Group Channels
	SendbirdURLGroupChannelsWithChannelURL                 templateKeysGroupChannels = `base`
	SendbirdURLGroupChannelsMembersWithChannelURL          templateKeysGroupChannels = `members`
	SendbirdURLGroupChannelsMembersWithChannelURLAndUserID templateKeysGroupChannels = `membersWithUserID`
	SendbirdURLGroupChannelsInviteWithChannelURL           templateKeysGroupChannels = `invite`
	SendbirdURLGroupChannelsHideWithChannelURL             templateKeysGroupChannels = `hide`
	SendbirdURLGroupChannelsLeaveWithChannelURL            templateKeysGroupChannels = `leave`
	SendbirdURLGroupChannelsFreezeWithChannelURL           templateKeysGroupChannels = `freeze`

	//Channel Metadata
	SendbirdURLChannelMetadataWithChannelTypeAndChannelURL              templateKeysChannelMetadata = `chMeta`
	SendbirdURLChannelMetadataWithChannelTypeAndChannelURLAndKeyName    templateKeysChannelMetadata = `chMetaWithKeyName`
	SendbirdURLChannelMetaCounterWithChannelTypeAndChannelURL           templateKeysChannelMetadata = `metacounter`
	SendbirdURLChannelMetaCounterWithChannelTypeAndChannelURLAndKeyName templateKeysChannelMetadata = `metacounterWithKeyName`

	//Messages
	SendbirdURLMessagesWithChannelTypeAndChannelURL             templateKeysMessages = `messages`
	SendbirdURLMessagesMarkAsReadWithChannelTypeAndChannelURL   templateKeysMessages = `markAsRead`
	SendbirdURLMessagesTotalCountWithChannelTypeAndChannelURL   templateKeysMessages = `totalCount`
	SendbirdURLMessagesUnreadCountWithChannelTypeAndChannelURL  templateKeysMessages = `unreadCount`
	SendbirdURLMessagesWithChannelTypeAndChannelURLAndMessageID templateKeysMessages = `singleMessage`
)
