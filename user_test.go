package sendbirdclient_test

import (
	. "sendbirdclient"
	"strings"
	"testing"
)

const (
	TestUserID1 = "UserTobeTested"
	TestUserID2 = "UserTobeBlocked"
	TestUserID3 = "UserTobeDeleted"

	TestUpdatedValue = "UpdatedValue"
)

var testUserClient = NewTestClient()

func TestUserActions(t *testing.T) {

	testCreateAUserWithURL(t, TestUserID1)
	testCreateAUserWithURL(t, TestUserID2)
	testCreateAUserWithURL(t, TestUserID3)

	testListUsers(t, TestUserID1, TestUserID2, TestUserID3)

	testUpdateAUserWithURL(t, TestUserID1)

	testViewAUser(t, TestUserID1)

	testGetUnreadMessageCount(t, TestUserID1)

	testBlockAUser(t, TestUserID1, TestUserID2)
	testListBlockedUsers(t, TestUserID1)
	testUnblockAUser(t, TestUserID1, TestUserID2)

	testListBannedChannels(t, TestUserID1)

	testMarkAllMessagesAsRead(t, TestUserID1)

	testListMyGroupChannels(t, TestUserID1)

	//testRegisterADeviceToken(t, "UserTobeTested", testGCMToken)
	//testUnregisterADeviceToken(t, "")
	//testUnregisterAllDeviceTokens(t, )
	//testListDeviceTokens(t, "UserTobeTested", "gcm")

	testUpdatePushPerferences(t, TestUserID1)
	testGetPushPerferences(t, TestUserID1)
	testResetPushPerferences(t, TestUserID1)

	//testUpdateChannelPushPerferences(t, "UserTobeTested", testChannelURL)
	//testGetChannelPushPerferences(t, "UserTobeTested", testChannelURL)

	testDeleteAUser(t, TestUserID1)
	testDeleteAUser(t, TestUserID2)
	testDeleteAUser(t, TestUserID3)
}

func testCreateAUserWithURL(t *testing.T, userID string) {
	r := &CreateAUserWithURLRequest{
		UserID:     userID,
		NickName:   userID,
		ProfileURL: "",
	}

	result, err := testUserClient.CreateAUserWithURL(r)
	if err != nil && !strings.Contains(err.Error(), SendbirdAPIErrorUniqueIDConstraint) {
		t.Errorf("Fail in testCreateAUserWithURL(): %+v", err)
	}

	t.Logf("testCreateAUserWithURL() Result: %+v", result)
}

func testListUsers(t *testing.T, userIDs ...string) {
	r := &ListUsersRequest{
		UserIDs: userIDs,
	}

	result, err := testUserClient.ListUsers(r)
	if err != nil {
		t.Errorf("Fail in testListUsers(): %+v", err)
	}

	t.Logf("testListUsers() Result: %+v", result)
}

func testUpdateAUserWithURL(t *testing.T, userID string) {
	r := &UpdateAUserWithURLRequest{
		NickName:   TestUpdatedValue,
		ProfileURL: "",
	}

	result, err := testUserClient.UpdateAUserWithURL(userID, r)
	if err != nil {
		t.Errorf("Fail in testUpdateAUserWithURL(): %+v", err)
	}

	t.Logf("testUpdateAUserWithURL() Result: %+v", result)
}

func testViewAUser(t *testing.T, userID string) {
	result, err := testUserClient.ViewAUser(userID)
	if err != nil {
		t.Errorf("Fail in testViewAUser(): %+v", err)
	}

	t.Logf("testViewAUser() Result: %+v", result)
}

func testGetUnreadMessageCount(t *testing.T, userID string) {
	result, err := testUserClient.GetUnreadMessageCount(userID)
	if err != nil {
		t.Errorf("Fail in testGetUnreadMessageCount(): %+v", err)
	}

	t.Logf("testGetUnreadMessageCount() Result: %+v", result)
}

func testBlockAUser(t *testing.T, userID string, targetID string) {
	r := &BlockAUserRequest{
		TargetID: targetID,
	}

	result, err := testUserClient.BlockAUser(userID, r)
	if err != nil {
		t.Errorf("Fail in testBlockAUser(): %+v", err)
	}

	t.Logf("testBlockAUser() Result: %+v", result)
}

func testListBlockedUsers(t *testing.T, userID string) {
	r := &ListBlockUsersRequest{}

	result, err := testUserClient.ListBlockedUsers(userID, r)
	if err != nil {
		t.Errorf("Fail in testListBlockedUsers(): %+v", err)
	}

	t.Logf("testListBlockedUsers() Result: %+v", result)
}

func testUnblockAUser(t *testing.T, userID string, targetID string) {
	result, err := testUserClient.UnblockAUser(userID, targetID)
	if err != nil {
		t.Errorf("Fail in testUnblockAUser(): %+v", err)
	}

	t.Logf("testUnblockAUser() Result: %+v", result)
}

func testDeleteAUser(t *testing.T, userID string) {
	result, err := testUserClient.DeleteAUser(userID)
	if err != nil && !strings.Contains(err.Error(), SendbirdAPIErrorUserNotFound) {
		t.Errorf("Fail in testDeleteAUser(): %+v", err)
	}

	t.Logf("testDeleteAUser() Result: %+v", result)
}

func testListBannedChannels(t *testing.T, userID string) {
	r := &ListBannedChannelsRequest{}

	result, err := testUserClient.ListBannedChannels(userID, r)
	if err != nil {
		t.Errorf("Fail in testListBannedChannels(): %+v", err)
	}

	t.Logf("testListBannedChannels() Result: %+v", result)
}

func testMutedChannels(t *testing.T, userID string) {
	r := &ListMutedChannelsRequest{}

	result, err := testUserClient.ListMutedChannels(userID, r)
	if err != nil {
		t.Errorf("Fail in testMutedChannels(): %+v", err)
	}

	t.Logf("testMutedChannels() Result: %+v", result)
}

func testMarkAllMessagesAsRead(t *testing.T, userID string) {
	result, err := testUserClient.MarkAllMessagesAsRead(userID)
	if err != nil {
		t.Errorf("Fail in testMarkAllMessagesAsRead(): %+v", err)
	}

	t.Logf("testMarkAllMessagesAsRead() Result: %+v", result)
}

func testListMyGroupChannels(t *testing.T, userID string) {
	r := &ListMyGroupChannelsRequest{}

	result, err := testUserClient.ListMyGroupChannels(userID, r)
	if err != nil {
		t.Errorf("Fail in testListMyGroupChannels(): %+v", err)
	}

	t.Logf("testListMyGroupChannels() Result: %+v", result)
}

func testRegisterADeviceToken(t *testing.T, userID string, tokenType string) {
	r := &RegisterADeviceTokenRequest{}

	result, err := testUserClient.RegisterADeviceToken(userID, tokenType, r)
	if err != nil {
		t.Errorf("Fail in testRegisterADeviceToken(): %+v", err)
	}

	t.Logf("testRegisterADeviceToken() Result: %+v", result)
}

func testUnregisterADeviceToken(t *testing.T, userID string, tokenType string, pushToken string) {
	result, err := testUserClient.UnregisterADeviceToken(userID, tokenType, pushToken)
	if err != nil {
		t.Errorf("Fail in testUnregisterADeviceToken(): %+v", err)
	}

	t.Logf("testUnregisterADeviceToken() Result: %+v", result)
}

func testUnregisterAllDeviceTokens(t *testing.T, userID string) {
	result, err := testUserClient.UnregisterAllDeviceTokens(userID)
	if err != nil {
		t.Errorf("Fail in testUnregisterAllDeviceTokens(): %+v", err)
	}

	t.Logf("testUnregisterAllDeviceTokens() Result: %+v", result)
}

func testListDeviceTokens(t *testing.T, userID string, tokenType string) {
	result, err := testUserClient.ListDeviceTokens(userID, tokenType)
	if err != nil {
		t.Errorf("Fail in testListDeviceTokens(): %+v", err)
	}

	t.Logf("testListDeviceTokens() Result: %+v", result)
}

func testUpdatePushPerferences(t *testing.T, userID string) {
	r := &UpdatePushPerferencesRequest{}

	result, err := testUserClient.UpdatePushPerferences(userID, r)
	if err != nil {
		t.Errorf("Fail in testUpdatePushPerferences(): %+v", err)
	}

	t.Logf("testUpdatePushPerferences() Result: %+v", result)
}

func testGetPushPerferences(t *testing.T, userID string) {
	result, err := testUserClient.GetPushPerferences(userID)
	if err != nil {
		t.Errorf("Fail in testGetPushPerferences(): %+v", err)
	}

	t.Logf("testGetPushPerferences() Result: %+v", result)
}

func testResetPushPerferences(t *testing.T, userID string) {
	result, err := testUserClient.ResetPushPerferences(userID)
	if err != nil {
		t.Errorf("Fail in testResetPushPerferences(): %+v", err)
	}

	t.Logf("testResetPushPerferences() Result: %+v", result)
}

func testUpdateChannelPushPerferences(t *testing.T, userID string, channelURL string) {
	r := &UpdateChannelPushPerferencesRequest{
		Enable: true,
	}

	result, err := testUserClient.UpdateChannelPushPerferences(userID, channelURL, r)
	if err != nil {
		t.Errorf("Fail in testUpdateChannelPushPerferences(): %+v", err)
	}

	t.Logf("testUpdateChannelPushPerferences() Result: %+v", result)

}

func testGetChannelPushPerferences(t *testing.T, userID string, channelURL string) {
	result, err := testUserClient.GetChannelPushPerferences(userID, channelURL)
	if err != nil {
		t.Errorf("Fail in testGetChannelPushPerferences(): %+v", err)
	}

	t.Logf("testGetChannelPushPerferences() Result: %+v", result)

}
