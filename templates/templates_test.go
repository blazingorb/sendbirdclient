package templates

import (
	"html/template"
	"testing"
)

const (
	TestUserID       = "TestUserID"
	TestTargetUserID = "TestTargetUserID"
	TestTokenType    = "GCM"
	TestPushToken    = "TestPushToken"
	TestChannelURL   = "TestChannelURL"

	TestKeyName = "Key1"

	TestChannelType1 = "open_channels"
	TestChannelType2 = "group_channels"

	TestMessageID = "TestMessageID"
)

var TestUsersTemplateData = struct {
	UserID     string
	TargetID   string
	TokenType  string
	PushToken  string
	ChannelURL string
}{
	UserID:     TestUserID,
	TargetID:   TestTargetUserID,
	TokenType:  TestTokenType,
	PushToken:  TestPushToken,
	ChannelURL: TestChannelURL,
}

var TestUserMetaTemplateData = struct {
	UserID  string
	KeyName string
}{
	UserID:  TestUserID,
	KeyName: TestKeyName,
}

var TestOpenChannelsTemplateData = struct {
	ChannelURL   string
	BannedUserID string
	MutedUserID  string
}{
	ChannelURL:   TestChannelURL,
	BannedUserID: TestUserID,
	MutedUserID:  TestUserID,
}

var TestGroupChannelsTemplateData = struct {
	ChannelURL string
	UserID     string
}{
	ChannelURL: TestChannelURL,
	UserID:     TestUserID,
}

var TestChannelMetadataTemplateData = struct {
	ChannelType string
	ChannelURL  string
	KeyName     string
}{
	ChannelType: TestChannelType1,
	ChannelURL:  TestChannelURL,
	KeyName:     TestKeyName,
}

var TestMessagesTemplateData = struct {
	ChannelType string
	ChannelURL  string
	MessageID   string
}{
	ChannelType: TestChannelType1,
	ChannelURL:  TestChannelURL,
	MessageID:   TestMessageID,
}

//Expected Results
const (
	SendbirdURLUserswithUserID_Result = `/users/12345`

	//Users
	SendbirdURLResultUserswithUserID                                    = `/users/TestUserID`
	SendbirdURLResultUsersUnreadCountWithUserID                         = `/users/TestUserID/unread_count`
	SendbirdURLResultUsersBlockWithUserID                               = `/users/TestUserID/block`
	SendbirdURLResultUsersBlockWithUserIDandTargetID                    = `/users/TestUserID/block/TestTargetUserID`
	SendbirdURLResultUsersBanWithUserID                                 = `/users/TestUserID/ban`
	SendbirdURLResultUsersMuteWithUserID                                = `/users/TestUserID/mute`
	SendbirdURLResultUsersMarkReadAllWithUserID                         = `/users/TestUserID/mark_as_read_all`
	SendbirdURLResultUsersListGroupChannelsWithUserID                   = `/users/TestUserID/my_group_channels`
	SendbirdURLResultUsersDeviceTokenWithUserID                         = `/users/TestUserID/push`
	SendbirdURLResultUsersDeviceTokenWithUserIDandTokenType             = `/users/TestUserID/push/GCM`
	SendbirdURLResultUsersDeviceTokenWithUserIDandTokenTypeandPushToken = `/users/TestUserID/push/GCM/TestPushToken`
	SendbirdURLResultUsersPushPreferenceWithUserID                      = `/users/TestUserID/push_preference`
	SendbirdURLResultUsersPushPreferenceWithUserIDandChannelURL         = `/users/TestUserID/push_preference/TestChannelURL`

	//User Metadata
	SendbirdURLResultUserMetadataWithUserID           = `/users/TestUserID/metadata`
	SendbirdURLResultUserMetadataWithUserIDandKeyName = `/users/TestUserID/metadata/Key1`

	//Open Channels
	SendbirdURLResultOpenChannelsWithChannelURL                   = `/open_channels/TestChannelURL`
	SendbirdURLResultOpenChannelsParticipantsWithChannelURL       = `/open_channels/TestChannelURL/participants`
	SendbirdURLResultOpenChannelsFreezeWithChannelURL             = `/open_channels/TestChannelURL/freeze`
	SendbirdURLResultOpenChannelsBanWithChannelURL                = `/open_channels/TestChannelURL/ban`
	SendbirdURLResultOpenChannelsBanWithChannelURLandBannedUserID = `/open_channels/TestChannelURL/ban/TestUserID`
	SendbirdURLResultOpenChannelsMuteWithChannelURL               = `/open_channels/TestChannelURL/mute`
	SendbirdURLResultOpenChannelsMuteWithChannelURLandMutedUserID = `/open_channels/TestChannelURL/mute/TestUserID`

	//Group Channels
	SendbirdURLResultGroupChannelsWithChannelURL                 = `/group_channels/TestChannelURL`
	SendbirdURLResultGroupChannelsMembersWithChannelURL          = `/group_channels/TestChannelURL/members`
	SendbirdURLResultGroupChannelsMembersWithChannelURLAndUserID = `/group_channels/TestChannelURL/members/TestUserID`
	SendbirdURLResultGroupChannelsInviteWithChannelURL           = `/group_channels/TestChannelURL/invite`
	SendbirdURLResultGroupChannelsHideWithChannelURL             = `/group_channels/TestChannelURL/hide`
	SendbirdURLResultGroupChannelsLeaveWithChannelURL            = `/group_channels/TestChannelURL/leave`
	SendbirdURLResultGroupChannelsFreezeWithChannelURL           = `/group_channels/TestChannelURL/freeze`

	//Group Metadata
	SendbirdURLResultChannelMetadataWithChannelTypeAndChannelURL              = `/open_channels/TestChannelURL/metadata`
	SendbirdURLResultChannelMetadataWithChannelTypeAndChannelURLAndKeyName    = `/open_channels/TestChannelURL/metadata/Key1`
	SendbirdURLResultChannelMetaCounterWithChannelTypeAndChannelURL           = `/open_channels/TestChannelURL/metacounter`
	SendbirdURLResultChannelMetaCounterWithChannelTypeAndChannelURLAndKeyName = `/open_channels/TestChannelURL/metacounter/Key1`

	//Messages
	SendbirdURLResultMessagesWithChannelTypeAndChannelURL             = `/open_channels/TestChannelURL/messages`
	SendbirdURLResultMessagesMarkAsReadWithChannelTypeAndChannelURL   = `/open_channels/TestChannelURL/messages/mark_as_read`
	SendbirdURLResultMessagesTotalCountWithChannelTypeAndChannelURL   = `/open_channels/TestChannelURL/messages/total_count`
	SendbirdURLResultMessagesUnreadCountWithChannelTypeAndChannelURL  = `/open_channels/TestChannelURL/messages/unread_count`
	SendbirdURLResultMessagesWithChannelTypeAndChannelURLAndMessageID = `/open_channels/TestChannelURL/messages/TestMessageID`
)

var testEngine = createEngine()

func TestTemplateActions(t *testing.T) {
	testCreateEngine(t)

	testEngineExecute(t, testEngine.templateUsers, string(SendbirdURLUsersDeviceTokenWithUserIDandTokenTypeandPushToken), TestUsersTemplateData, SendbirdURLResultUsersDeviceTokenWithUserIDandTokenTypeandPushToken)
	testEngineExecute(t, testEngine.templateUsers, string(SendbirdURLUserMetadataWithUserIDandKeyName), TestUserMetaTemplateData, SendbirdURLResultUserMetadataWithUserIDandKeyName)
	testEngineExecute(t, testEngine.templateOpenChannels, string(SendbirdURLOpenChannelsBanWithChannelURLandBannedUserID), TestOpenChannelsTemplateData, SendbirdURLResultOpenChannelsBanWithChannelURLandBannedUserID)
	testEngineExecute(t, testEngine.templateGroupChannels, string(SendbirdURLGroupChannelsMembersWithChannelURLAndUserID), TestGroupChannelsTemplateData, SendbirdURLResultGroupChannelsMembersWithChannelURLAndUserID)
	testEngineExecute(t, testEngine.templateChannelMetadata, string(SendbirdURLChannelMetaCounterWithChannelTypeAndChannelURLAndKeyName), TestChannelMetadataTemplateData, SendbirdURLResultChannelMetaCounterWithChannelTypeAndChannelURLAndKeyName)
	testEngineExecute(t, testEngine.templateMessages, string(SendbirdURLMessagesWithChannelTypeAndChannelURLAndMessageID), TestMessagesTemplateData, SendbirdURLResultMessagesWithChannelTypeAndChannelURLAndMessageID)
}

func testCreateEngine(t *testing.T) {

	if testEngine.templateUsers == nil {
		t.Errorf(SendbirdClientErrorFailedTemplateEngineCreation, "testEngine.templateUsers")
	}

	if testEngine.templateOpenChannels == nil {
		t.Errorf(SendbirdClientErrorFailedTemplateEngineCreation, "testEngine.templateOpenChannels")
	}

	if testEngine.templateGroupChannels == nil {
		t.Errorf(SendbirdClientErrorFailedTemplateEngineCreation, "testEngine.templateGroupChannels")
	}

	if testEngine.templateChannelMetadata == nil {
		t.Errorf(SendbirdClientErrorFailedTemplateEngineCreation, "testEngine.templateChannelMetadata")
	}

	if testEngine.templateMessages == nil {
		t.Errorf(SendbirdClientErrorFailedTemplateEngineCreation, "testEngine.templateMessages")
	}

}

func testEngineExecute(t *testing.T, tmpl *template.Template, key string, templateData interface{}, expectedAnswer string) {
	if tmpl == nil {
		t.Error(SendbirdClientErrorTemplateNil)
	}

	str, err := testEngine.execute(tmpl, key, templateData)
	if err != nil {
		t.Error(err)
	}

	if str != expectedAnswer {
		t.Errorf(SendbirdClientErrorTemplateResultMismatch, expectedAnswer, str)
	}
}
