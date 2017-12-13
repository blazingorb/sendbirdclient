package sendbirdclient_test

import (
	"strings"
	"testing"

	. "github.com/blazingorb/sendbirdclient"
)

const (
	TestUserID        = "TestMetaUser"
	TestDataKey1      = "key1"
	TestDataKey2      = "key2"
	TestDataKey3      = "key3"
	TestDataKey4      = "key4"
	TestDataValue1    = "value1"
	TestDataValue2    = "value2"
	TestDataValue3    = "value3"
	TestUpdatedValue1 = "updatedValue1"
	TestUpdatedValue2 = "updatedValue2"
	TestUpdatedValue3 = "updatedValue3"
	TestUpdatedValue4 = "updatedValue4"
)

var testUserMetadata1 = map[string]string{
	TestDataKey1: TestDataValue1,
	TestDataKey2: TestDataValue2,
	TestDataKey3: TestDataValue3,
}

var testUserMetadata2 = map[string]string{
	TestDataKey1: TestUpdatedValue1,
	TestDataKey2: TestUpdatedValue2,
	TestDataKey3: TestUpdatedValue3,
	TestDataKey4: TestUpdatedValue4,
}

var testMetaClient = NewTestClient()

func TestUserMetaActions(t *testing.T) {
	//Init Delete
	testDeleteUserMetaDataItem(t, TestUserID, TestDataKey1)
	testDeleteUserMetaData(t, TestUserID)
	testDeleteAUser(t, TestUserID)

	testCreateAUserWithURL(t, TestUserID)

	testCreateAnUserMetaData(t, TestUserID, testUserMetadata1)

	testListUserMetaData(t, TestUserID, testUserMetadata1, TestDataKey1, TestDataKey2)

	testViewAnUserMetaData(t, TestUserID, testUserMetadata1, TestDataKey1)

	testUpdateUserMetaData(t, TestUserID, testUserMetadata2)

	testUpdateUserMetaDataItem(t, TestUserID, TestDataKey1)

	testDeleteUserMetaDataItem(t, TestUserID, TestDataKey1)
	testDeleteUserMetaData(t, TestUserID)

	testDeleteAUser(t, TestUserID)
}

func testCreateAnUserMetaData(t *testing.T, userID string, metaData map[string]string) {

	r := &CreateAnUserMetaDataRequest{
		MetaData: metaData,
	}

	result, err := testMetaClient.CreateAnUserMetaData(userID, r)
	if err != nil {
		t.Errorf("Fail in testCreateAnUserMetaData(): %+v", err)
	}

	t.Logf("testCreateAnUserMetaData() Result: %+v", result)
}

func testListUserMetaData(t *testing.T, userID string, testMetaData map[string]string, metaKeys ...string) {
	r := &ListUserMetaDataRequest{
		Keys: metaKeys,
	}

	result, err := testMetaClient.ListUserMetaData(userID, r)
	if err != nil {
		t.Errorf("Fail in testListUserMetaData(): %+v", err)
	}

	for _, testKey := range metaKeys {
		mapValue, ok := result[testKey]
		if !ok || mapValue != testMetaData[testKey] {
			t.Errorf("Fail in testListUserMetaData(): %s", "Key and value dismatched.")
		}
	}

	t.Logf("testListUserMetaData() Result: %+v", result)
}

func testViewAnUserMetaData(t *testing.T, userID string, testMetaData map[string]string, keyName string) {
	result, err := testMetaClient.ViewAnUserMetaData(userID, keyName)
	if err != nil {
		t.Errorf("Fail in testViewAnUserMetaData(): %+v", err)
	}

	if result[keyName] != testMetaData[keyName] {
		t.Errorf("Fail in testViewAnUserMetaData(): %s", SendbirdClientErrorKeyValueMismatch)
	}

	t.Logf("testViewAnUserMetaData() Result: %+v", result)
}

func testUpdateUserMetaData(t *testing.T, userID string, updatedMetaData map[string]string) {
	r := &UpdateUserMetaDataRequest{
		MetaData: updatedMetaData,
		Upsert:   true,
	}

	result, err := testMetaClient.UpdateUserMetaData(userID, r)
	if err != nil {
		t.Errorf("Fail in testUpdateUserMetaData(): %+v", err)
	}

	t.Logf("testUpdateUserMetaData() Result: %+v", result)
}

func testUpdateUserMetaDataItem(t *testing.T, userID string, keyName string) {

	r := &UpdateUserMetaDataItemRequest{
		Value:  TestUpdatedValue1,
		Upsert: true,
	}

	result, err := testMetaClient.UpdateUserMetaDataItem(userID, keyName, r)
	if err != nil {
		t.Errorf("Fail in testUpdateUserMetaDataItem(): %+v", err)
	}

	t.Logf("testUpdateUserMetaDataItem() Result: %+v", result)
}

func testDeleteUserMetaData(t *testing.T, userID string) {
	result, err := testMetaClient.DeleteUserMetaData(userID)
	if err != nil && !strings.Contains(err.Error(), SendbirdAPIErrorUserNotFound) {
		t.Errorf("Fail in testUpdateUserMetaDataItem(): %+v", err)
	}

	t.Logf("testUpdateUserMetaDataItem() Result: %+v", result)
}

func testDeleteUserMetaDataItem(t *testing.T, userID string, keyName string) {
	result, err := testMetaClient.DeleteUserMetaDataItem(userID, keyName)
	if err != nil && !strings.Contains(err.Error(), SendbirdAPIErrorUserNotFound) {
		t.Errorf("Fail in testDeleteUserMetaItemData(): %+v", err)
	}

	t.Logf("testDeleteUserMetaItemData() Result: %+v", result)
}
