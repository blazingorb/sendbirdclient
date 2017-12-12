package sendbirdclient_test

import (
	. "sendbirdclient"
	"strings"
	"testing"
)

const (
	TestOpenChannelName                 = "TestOpenChannel"
	TestOpenChannelUpdateName           = "TestOpenChannelUpdated"
	TestOpenChannelURL                  = "TestOpenChannelURL"
	TestOpenChannelUserID1              = "TestOpenChannelUser1"
	TestOpenChannelUserID2              = "TestOpenChannelUser2"
	TestOpenChannelUpdatedBannedSeconds = 100
)

var TestOpenChannelOperators = []string{TestOpenChannelUserID1, TestOpenChannelUserID2}

var testOpenChannelClient = NewTestClient()

func TestOpenChannelsActions(t *testing.T) {
	//Init Delete
	testDeleteAnOpenChannel(t, TestOpenChannelURL)
	testDeleteAUser(t, TestOpenChannelUserID1)
	testDeleteAUser(t, TestOpenChannelUserID2)

	testCreateAUserWithURL(t, TestOpenChannelUserID1)
	testCreateAUserWithURL(t, TestOpenChannelUserID2)

	testCreateAnOpenChannelWithURL(t, TestOpenChannelName, TestOpenChannelURL, TestOpenChannelOperators)

	testListOpenChannels(t)

	testUpdateAnOpenChannelWithURL(t, TestOpenChannelURL, TestOpenChannelUpdateName)

	testViewAnOpenChannel(t, TestOpenChannelURL)

	testListOpenChannelParticipants(t, TestOpenChannelURL)

	testFreezeAnOpenChannel(t, TestOpenChannelURL, true)
	testFreezeAnOpenChannel(t, TestOpenChannelURL, false)

	testBanAUserInOpenChannel(t, TestOpenChannelURL, TestOpenChannelUserID1)
	testListBannedUsersInOpenChannel(t, TestOpenChannelURL)
	testUpdateBanInOpenChannel(t, TestOpenChannelURL, TestOpenChannelUserID1, TestOpenChannelUpdatedBannedSeconds)
	testViewBanInOpenChannel(t, TestOpenChannelURL, TestOpenChannelUserID1)
	testUnbanAUserInOpenChannel(t, TestOpenChannelURL, TestOpenChannelUserID1)

	testMuteAUserInOpenChannel(t, TestOpenChannelURL, TestOpenChannelUserID1)
	testListMutedUsersInOpenChannel(t, TestOpenChannelURL)
	testViewAMuteInOpenChannel(t, TestOpenChannelURL, TestOpenChannelUserID1)
	testUnmuteAUserInOpenChannel(t, TestOpenChannelURL, TestOpenChannelUserID1)

	testDeleteAnOpenChannel(t, TestOpenChannelURL)

	testDeleteAUser(t, TestOpenChannelUserID1)
	testDeleteAUser(t, TestOpenChannelUserID2)
}

func testCreateAnOpenChannelWithURL(t *testing.T, name string, channelURL string, operators []string) {
	r := &CreateAnOpenChannelWithURLRequest{
		Name:       name,
		ChannelURL: channelURL,
		Operators:  operators,
	}

	result, err := testOpenChannelClient.CreateAnOpenChannelWithURL(r)
	if err != nil && !strings.Contains(err.Error(), SendbirdAPIErrorUniqueIDConstraint) {
		t.Errorf("Fail in testCreateAnOpenChannelWithURL(): %+v", err)
	}

	t.Logf("testCreateAnOpenChannelWithURL() Result: %+v", result)
}

func testListOpenChannels(t *testing.T) {
	r := &ListOpenChannelsRequest{}

	result, err := testOpenChannelClient.ListOpenChannels(r)
	if err != nil {
		t.Errorf("Fail in testListOpenChannels(): %+v", err)
	}

	t.Logf("testListOpenChannels() Result: %+v", result)
}

func testUpdateAnOpenChannelWithURL(t *testing.T, channelURL string, updateName string) {
	r := &UpdateAnOpenChannelWithURLRequest{
		Name: updateName,
	}

	result, err := testOpenChannelClient.UpdateAnOpenChannelWithURL(channelURL, r)
	if err != nil {
		t.Errorf("Fail in testUpdateAnOpenChannelWithURL(): %+v", err)
	}

	t.Logf("testUpdateAnOpenChannelWithURL() Result: %+v", result)
}

func testViewAnOpenChannel(t *testing.T, channelURL string) {
	r := &ViewAnOpenChannelRequest{
		Participants: true,
	}

	result, err := testOpenChannelClient.ViewAnOpenChannel(channelURL, r)
	if err != nil {
		t.Errorf("Fail in testViewAnOpenChannel(): %+v", err)
	}

	t.Logf("testViewAnOpenChannel() Result: %+v", result)
}

func testDeleteAnOpenChannel(t *testing.T, channelURL string) {

	result, err := testOpenChannelClient.DeleteAnOpenChannel(channelURL)
	if err != nil && !strings.Contains(err.Error(), SendbirdAPIErrorChannelNotFound) {
		t.Errorf("Fail in testDeleteAnOpenChannel(): %+v", err)
	}

	t.Logf("testDeleteAnOpenChannel() Result: %+v", result)
}

func testListOpenChannelParticipants(t *testing.T, channelURL string) {
	r := &ListOpenChannelParticipantsRequest{}
	result, err := testOpenChannelClient.ListOpenChannelParticipants(channelURL, r)
	if err != nil {
		t.Errorf("Fail in testListOpenChannelParticipants(): %+v", err)
	}

	t.Logf("testListOpenChannelParticipants() Result: %+v", result)
}

func testFreezeAnOpenChannel(t *testing.T, channelURL string, freeze bool) {
	r := &FreezeAnOpenChannelRequest{
		Freeze: freeze,
	}
	result, err := testOpenChannelClient.FreezeAnOpenChannel(channelURL, r)
	if err != nil {
		t.Errorf("Fail in testFreezeAnOpenChannel(): %+v", err)
	}

	t.Logf("testFreezeAnOpenChannel() Result: %+v", result)
}

func testBanAUserInOpenChannel(t *testing.T, channelURL string, userID string) {
	r := &BanAUserInOpenChannelRequest{
		UserID: userID,
	}
	result, err := testOpenChannelClient.BanAUserInOpenChannel(channelURL, r)
	if err != nil {
		t.Errorf("Fail in testBanAUserInOpenChannel(): %+v", err)
	}

	t.Logf("testBanAUserInOpenChannel() Result: %+v", result)
}

func testListBannedUsersInOpenChannel(t *testing.T, channelURL string) {
	r := &ListBannedUsersInOpenChannelRequest{}
	result, err := testOpenChannelClient.ListBannedUsersInOpenChannel(channelURL, r)
	if err != nil {
		t.Errorf("Fail in testListBannedUsersInOpenChannel(): %+v", err)
	}

	t.Logf("testListBannedUsersInOpenChannel() Result: %+v", result)
}

func testUpdateBanInOpenChannel(t *testing.T, channelURL string, bannedUserID string, updateSeconds int) {
	r := &UpdateBanInOpenChannelRequest{
		Seconds: updateSeconds,
	}
	result, err := testOpenChannelClient.UpdateBanInOpenChannel(channelURL, bannedUserID, r)
	if err != nil {
		t.Errorf("Fail in testUpdateBanInOpenChannel(): %+v", err)
	}

	t.Logf("testUpdateBanInOpenChannel() Result: %+v", result)
}

func testViewBanInOpenChannel(t *testing.T, channelURL string, bannedUserID string) {
	result, err := testOpenChannelClient.ViewBanInOpenChannel(channelURL, bannedUserID)
	if err != nil {
		t.Errorf("Fail in testViewBanInOpenChannel(): %+v", err)
	}

	t.Logf("testViewBanInOpenChannel() Result: %+v", result)
}

func testUnbanAUserInOpenChannel(t *testing.T, channelURL string, bannedUserID string) {
	result, err := testOpenChannelClient.UnbanAUserInOpenChannel(channelURL, bannedUserID)
	if err != nil {
		t.Errorf("Fail in testUnbanAUserInOpenChannel(): %+v", err)
	}

	t.Logf("testUnbanAUserInOpenChannel() Result: %+v", result)
}

func testMuteAUserInOpenChannel(t *testing.T, channelURL string, userID string) {
	r := &MuteAUserInOpenChannelRequest{
		UserID: userID,
	}

	result, err := testOpenChannelClient.MuteAUserInOpenChannel(channelURL, r)
	if err != nil {
		t.Errorf("Fail in testMuteAUserInOpenChannel(): %+v", err)
	}

	t.Logf("testMuteAUserInOpenChannel() Result: %+v", result)

}

func testListMutedUsersInOpenChannel(t *testing.T, channelURL string) {
	r := &ListMutedUsersInOpenChannelRequest{}

	result, err := testOpenChannelClient.ListMutedUsersInOpenChannel(channelURL, r)
	if err != nil {
		t.Errorf("Fail in testListMutedUsersInOpenChannel(): %+v", err)
	}

	t.Logf("testListMutedUsersInOpenChannel() Result: %+v", result)
}

func testViewAMuteInOpenChannel(t *testing.T, channelURL string, mutedUserID string) {
	result, err := testOpenChannelClient.ViewAMuteInOpenChannel(channelURL, mutedUserID)
	if err != nil {
		t.Errorf("Fail in testViewAMuteInOpenChannel(): %+v", err)
	}

	t.Logf("testViewAMuteInOpenChannel() Result: %+v", result)
}

func testUnmuteAUserInOpenChannel(t *testing.T, channelURL string, mutedUserID string) {
	result, err := testOpenChannelClient.UnmuteAUserInOpenChannel(channelURL, mutedUserID)
	if err != nil {
		t.Errorf("Fail in testUnmuteAUserInOpenChannel(): %+v", err)
	}

	t.Logf("testUnmuteAUserInOpenChannel() Result: %+v", result)
}
