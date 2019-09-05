package sendbirdclient_test

import (
	. "github.com/blazingorb/sendbirdclient"
	"strings"
	"testing"
)

const (
	TestGroupnChannelName                 = "TestGroupChannel"
	TestGroupnChannelUpdateName           = "TestGroupChannelUpdated"
	TestGroupnCoverURL                    = "TestGroupChannelCoverURL"
	TestGroupnChannelUserID1              = "TestGroupChannelUser1"
	TestGroupnChannelUserID2              = "TestGroupChannelUser2"
	TestGroupnChannelUserID3              = "TestGroupChannelUser3"
	TestGroupnChannelUserID4              = "TestGroupChannelUser4"
	TestGroupnChannelUpdatedBannedSeconds = 100
)

var TestGroupChannelUsers = []string{TestGroupnChannelUserID1, TestGroupnChannelUserID2}
var TestGroupChannelAdditionalUsers = []string{TestGroupnChannelUserID3, TestGroupnChannelUserID4}
var testGroupChannelClient = NewTestClient()

func TestGroupChannelsActions(t *testing.T) {
	//Init Delete
	testDeleteAUser(t, TestGroupnChannelUserID1)
	testDeleteAUser(t, TestGroupnChannelUserID2)
	testDeleteAUser(t, TestGroupnChannelUserID3)
	testDeleteAUser(t, TestGroupnChannelUserID4)

	testCreateAUserWithURL(t, TestGroupnChannelUserID1)
	testCreateAUserWithURL(t, TestGroupnChannelUserID2)
	testCreateAUserWithURL(t, TestGroupnChannelUserID3)
	testCreateAUserWithURL(t, TestGroupnChannelUserID4)

	groupChannel := testCreateAGroupChannelWithURL(t, TestGroupnChannelName, TestGroupnCoverURL, TestGroupChannelUsers)

	testListGroupChannels(t, TestGroupChannelUsers)

	testUpdateAGroupChannel(t, groupChannel.ChannelURL, TestGroupnChannelUpdateName)

	testViewAGroupChannel(t, groupChannel.ChannelURL, true, true)

	testListMembersInGroupChannel(t, groupChannel.ChannelURL)

	testCheckIfMemberInGroupChannel(t, groupChannel.ChannelURL, TestGroupnChannelUserID1)

	testInviteMembersToGroupChannel(t, groupChannel.ChannelURL, TestGroupChannelAdditionalUsers)

	testFreezeAnGroupChannel(t, groupChannel.ChannelURL, true)
	testFreezeAnGroupChannel(t, groupChannel.ChannelURL, false)

	testHideFromAGroupChannel(t, groupChannel.ChannelURL, TestGroupnChannelUserID1)

	testLeaveFromAGroupChannel(t, groupChannel.ChannelURL, TestGroupChannelAdditionalUsers)

	testDeleteAGroupChannel(t, groupChannel.ChannelURL)

	testDeleteAUser(t, TestGroupnChannelUserID1)
	testDeleteAUser(t, TestGroupnChannelUserID2)
	testDeleteAUser(t, TestGroupnChannelUserID3)
	testDeleteAUser(t, TestGroupnChannelUserID4)
}

func testCreateAGroupChannelWithURL(t *testing.T, name string, coverURL string, userIDs []string) GroupChannel {
	r := &CreateAGroupChannelWithURLRequest{
		Name:     name,
		CoverURL: coverURL,
		UserIDs:  userIDs,
	}

	result, err := testGroupChannelClient.CreateAGroupChannelWithURL(r)
	if err != nil && !strings.Contains(err.Error(), SendbirdAPIErrorUniqueIDConstraint) {
		t.Errorf("Fail in testCreateAGroupChannelWithURL(): %+v", err)
	}

	t.Logf("testCreateAGroupChannelWithURL() Result: %+v", result)

	return result
}

func testListGroupChannels(t *testing.T, membersIncludeIn []string) {
	r := &ListGroupChannelsRequest{
		MembersIncludeIn: membersIncludeIn,
	}

	result, err := testGroupChannelClient.ListGroupChannels(r)
	if err != nil {
		t.Errorf("Fail in testListGroupChannels(): %+v", err)
	}

	t.Logf("testListGroupChannels() Result: %+v", result)
}

func testUpdateAGroupChannel(t *testing.T, channelURL string, updatedValue string) {
	r := &UpdateAGroupChannelRequest{
		Name: updatedValue,
	}

	result, err := testGroupChannelClient.UpdateAGroupChannel(channelURL, r)
	if err != nil {
		t.Errorf("Fail in testUpdateAGroupChannel(): %+v", err)
	}

	t.Logf("testUpdateAGroupChannel() Result: %+v", result)
}

func testViewAGroupChannel(t *testing.T, channelURL string, showReadReceipt bool, showMember bool) {
	r := &ViewAGroupChannelRequest{
		ShowReadReceipt: showReadReceipt,
		ShowMember:      showMember,
	}

	result, err := testGroupChannelClient.ViewAGroupChannel(channelURL, r)
	if err != nil {
		t.Errorf("Fail in testViewAGroupChannel(): %+v", err)
	}

	t.Logf("testViewAGroupChannel() Result: %+v", result)
}

func testDeleteAGroupChannel(t *testing.T, channelURL string) {
	result, err := testGroupChannelClient.DeleteAGroupChannel(channelURL)
	if err != nil && !strings.Contains(err.Error(), SendbirdAPIErrorChannelNotFound) {
		t.Errorf("Fail in testDeleteAGroupChannel(): %+v", err)
	}

	t.Logf("testDeleteAGroupChannel() Result: %+v", result)
}

func testListMembersInGroupChannel(t *testing.T, channelURL string) {
	r := &ListMembersInGroupChannelRequest{}

	result, err := testGroupChannelClient.ListMembersInGroupChannel(channelURL, r)
	if err != nil {
		t.Errorf("Fail in testListMembersInGroupChannel(): %+v", err)
	}

	t.Logf("testListMembersInGroupChannel() Result: %+v", result)
}

func testCheckIfMemberInGroupChannel(t *testing.T, channelURL string, userID string) {
	result, err := testGroupChannelClient.CheckIfMemberInGroupChannel(channelURL, userID)
	if err != nil {
		t.Errorf("Fail in testCheckIfMemberInGroupChannel(): %+v", err)
	}

	t.Logf("testCheckIfMemberInGroupChannel() Result: %+v", result)
}

func testInviteMembersToGroupChannel(t *testing.T, channelURL string, addUserIDs []string) {
	r := &InviteMembersToGroupChannelRequest{
		UserIDs: addUserIDs,
	}

	result, err := testGroupChannelClient.InviteMembersToGroupChannel(channelURL, r)
	if err != nil {
		t.Errorf("Fail in testInviteMembersToGroupChannel(): %+v", err)
	}

	t.Logf("testInviteMembersToGroupChannel() Result: %+v", result)
}

func testHideFromAGroupChannel(t *testing.T, channelURL string, userID string) {
	r := &HideFromAGroupChannelRequest{
		UserID: userID,
	}

	result, err := testGroupChannelClient.HideFromAGroupChannel(channelURL, r)
	if err != nil {
		t.Errorf("Fail in testHideFromAGroupChannel(): %+v", err)
	}

	t.Logf("testHideFromAGroupChannel() Result: %+v", result)
}

func testLeaveFromAGroupChannel(t *testing.T, channelURL string, leaveUserIDs []string) {
	r := &LeaveFromAGroupChannelRequest{
		UserIDs: leaveUserIDs,
	}

	result, err := testGroupChannelClient.LeaveFromAGroupChannel(channelURL, r)
	if err != nil {
		t.Errorf("Fail in testLeaveFromAGroupChannel(): %+v", err)
	}

	t.Logf("testLeaveFromAGroupChannel() Result: %+v", result)
}

func testFreezeAnGroupChannel(t *testing.T, channelURL string, freeze bool) {
	r := &FreezeAnGroupChannelRequest{
		Freeze: freeze,
	}
	result, err := testGroupChannelClient.FreezeAnGroupChannel(channelURL, r)
	if err != nil {
		t.Errorf("Fail in testFreezeAnGroupChannel(): %+v", err)
	}

	t.Logf("testFreezeAnGroupChannel() Result: %+v", result)
}
