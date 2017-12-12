package sendbirdclient_test

import (
	. "sendbirdclient"
	"strings"
	"testing"
)

const (
	TestChannelMetaOpenChannelName      = "TestOpenChannelForMeta"
	TestChannelMetaOpenChannelURL       = "TestOpenChannelForMetaURL"
	TestChannelType1                    = "open_channels"
	TestChannelType2                    = "group_channels"
	TestChannelMetaDataKey1             = "key1"
	TestChannelMetaDataKey2             = "key2"
	TestChannelMetaDataKey3             = "key3"
	TestChannelMetaDataKey4             = "key4"
	TestChannelMetaDataValue1           = "value1"
	TestChannelMetaDataValue2           = "value2"
	TestChannelMetaDataValue3           = "value3"
	TestChannelMetaUpdatedValue1        = "updatedValue1"
	TestChannelMetaUpdatedValue2        = "updatedValue2"
	TestChannelMetaUpdatedValue3        = "updatedValue3"
	TestChannelMetaUpdatedValue4        = "updatedValue4"
	TestChannelMetaCounterValue1        = 1
	TestChannelMetaCounterValue2        = 2
	TestChannelMetaCounterValue3        = 3
	TestChannelMetaCounterUpdatedValue1 = 101
	TestChannelMetaCounterUpdatedValue2 = 102
	TestChannelMetaCounterUpdatedValue3 = 103
	TestChannelMetaCounterUpdatedValue4 = 104
)

var testChannelMetadata1 = map[string]string{
	TestChannelMetaDataKey1: TestChannelMetaDataValue1,
	TestChannelMetaDataKey2: TestChannelMetaDataValue2,
	TestChannelMetaDataKey3: TestChannelMetaDataValue3,
}

var testChannelMetadata2 = map[string]string{
	TestChannelMetaDataKey1: TestChannelMetaUpdatedValue1,
	TestChannelMetaDataKey2: TestChannelMetaUpdatedValue2,
	TestChannelMetaDataKey3: TestChannelMetaUpdatedValue3,
	TestChannelMetaDataKey4: TestChannelMetaUpdatedValue4,
}

var testChannelMetaCounterData1 = map[string]int{
	TestChannelMetaDataKey1: TestChannelMetaCounterValue1,
	TestChannelMetaDataKey2: TestChannelMetaCounterValue2,
	TestChannelMetaDataKey3: TestChannelMetaCounterValue3,
}

var testChannelMetaCounterData2 = map[string]int{
	TestChannelMetaDataKey1: TestChannelMetaCounterUpdatedValue1,
	TestChannelMetaDataKey2: TestChannelMetaCounterUpdatedValue2,
	TestChannelMetaDataKey3: TestChannelMetaCounterUpdatedValue3,
	TestChannelMetaDataKey4: TestChannelMetaCounterUpdatedValue4,
}

var testChannelMetaKeys = []string{TestChannelMetaDataKey1, TestChannelMetaDataKey2}

var testChannelMetaClient = NewTestClient()

func TestChannelMetadataActions(t *testing.T) {
	//Init Delete
	testDeleteChannelMetaDataByKeyName(t, TestChannelType1, TestChannelMetaOpenChannelURL, TestChannelMetaDataKey1)
	testDeleteChannelMetaData(t, TestChannelType1, TestChannelMetaOpenChannelURL)
	testDeleteChannelMetaCounterByKeyName(t, TestChannelType1, TestChannelMetaOpenChannelURL, TestChannelMetaDataKey1)
	testDeleteChannelMetaCounter(t, TestChannelType1, TestChannelMetaOpenChannelURL)
	testDeleteAnOpenChannel(t, TestChannelMetaOpenChannelURL)

	testCreateAnOpenChannelWithURL(t, TestChannelMetaOpenChannelName, TestChannelMetaOpenChannelURL, make([]string, 0))

	testCreateAChannelMetadata(t, TestChannelType1, TestChannelMetaOpenChannelURL, testChannelMetadata1)

	testViewChannelMetadata(t, TestChannelType1, TestChannelMetaOpenChannelURL, testChannelMetaKeys)
	testViewChannelMetadataByKeyName(t, TestChannelType1, TestChannelMetaOpenChannelURL, TestChannelMetaDataKey1)

	testUpdateChannelMetaData(t, TestChannelType1, TestChannelMetaOpenChannelURL, testChannelMetadata2, true)
	testUpdateChannelMetaDataByKeyName(t, TestChannelType1, TestChannelMetaOpenChannelURL, TestChannelMetaDataKey1, TestChannelMetaUpdatedValue1, true)

	testDeleteChannelMetaDataByKeyName(t, TestChannelType1, TestChannelMetaOpenChannelURL, TestChannelMetaDataKey1)
	testDeleteChannelMetaData(t, TestChannelType1, TestChannelMetaOpenChannelURL)

	testCreateChannelMetaCounter(t, TestChannelType1, TestChannelMetaOpenChannelURL, testChannelMetaCounterData1)

	testViewChannelMetaCounter(t, TestChannelType1, TestChannelMetaOpenChannelURL, testChannelMetaKeys)
	testViewChannelMetaCounterByKeyName(t, TestChannelType1, TestChannelMetaOpenChannelURL, TestChannelMetaDataKey1)

	testUpdateChannelMetaCounter(t, TestChannelType1, TestChannelMetaOpenChannelURL, testChannelMetaCounterData2, true)
	testUpdateChannelMetaCounterByKeyName(t, TestChannelType1, TestChannelMetaOpenChannelURL, TestChannelMetaDataKey1, TestChannelMetaCounterUpdatedValue1, true)

	testDeleteChannelMetaCounterByKeyName(t, TestChannelType1, TestChannelMetaOpenChannelURL, TestChannelMetaDataKey1)
	testDeleteChannelMetaCounter(t, TestChannelType1, TestChannelMetaOpenChannelURL)

	testDeleteAnOpenChannel(t, TestChannelMetaOpenChannelURL)

}

func testCreateAChannelMetadata(t *testing.T, channelType string, channelURL string, metaData map[string]string) {
	r := &CreateAChannelMetadataRequest{
		Metadata: metaData,
	}

	result, err := testChannelMetaClient.CreateAChannelMetadata(channelType, channelURL, r)
	if err != nil {
		t.Errorf("Fail in testCreateAChannelMetadata(): %+v", err)
	}

	t.Logf("testCreateAChannelMetadata() Result: %+v", result)
}

func testViewChannelMetadata(t *testing.T, channelType string, channelURL string, keys []string) {
	r := &ViewChannelMetadataRequest{
		Keys: keys,
	}

	result, err := testChannelMetaClient.ViewChannelMetadata(channelType, channelURL, r)
	if err != nil {
		t.Errorf("Fail in testViewChannelMetadata(): %+v", err)
	}

	t.Logf("testViewChannelMetadata() Result: %+v", result)

}

func testViewChannelMetadataByKeyName(t *testing.T, channelType string, channelURL string, keyName string) {
	result, err := testChannelMetaClient.ViewChannelMetadataByKeyName(channelType, channelURL, keyName)
	if err != nil {
		t.Errorf("Fail in testViewChannelMetadataByKeyName(): %+v", err)
	}

	t.Logf("testViewChannelMetadataByKeyName() Result: %+v", result)
}

func testUpdateChannelMetaData(t *testing.T, channelType string, channelURL string, metaData map[string]string, upsert bool) {
	r := &UpdateChannelMetaDataRequest{
		Metadata: metaData,
		Upsert:   upsert,
	}

	result, err := testChannelMetaClient.UpdateChannelMetaData(channelType, channelURL, r)
	if err != nil {
		t.Errorf("Fail in testUpdateChannelMetaData(): %+v", err)
	}

	t.Logf("testUpdateChannelMetaData() Result: %+v", result)

}

func testUpdateChannelMetaDataByKeyName(t *testing.T, channelType string, channelURL string, keyName string, value string, upsert bool) {
	r := &UpdateChannelMetaDataByKeyNameRequest{
		Value:  value,
		Upsert: upsert,
	}

	result, err := testChannelMetaClient.UpdateChannelMetaDataByKeyName(channelType, channelURL, keyName, r)
	if err != nil {
		t.Errorf("Fail in testUpdateChannelMetaDataByKeyName(): %+v", err)
	}

	t.Logf("testUpdateChannelMetaDataByKeyName() Result: %+v", result)
}

func testDeleteChannelMetaData(t *testing.T, channelType string, channelURL string) {
	result, err := testChannelMetaClient.DeleteChannelMetaData(channelType, channelURL)
	if err != nil && !strings.Contains(err.Error(), SendbirdAPIErrorChannelNotFound) {
		t.Errorf("Fail in testDeleteChannelMetaData(): %+v", err)
	}

	t.Logf("testDeleteChannelMetaData() Result: %+v", result)
}

func testDeleteChannelMetaDataByKeyName(t *testing.T, channelType string, channelURL string, keyName string) {
	result, err := testChannelMetaClient.DeleteChannelMetaDataByKeyName(channelType, channelURL, keyName)
	if err != nil && !strings.Contains(err.Error(), SendbirdAPIErrorChannelNotFound) {
		t.Errorf("Fail in testDeleteChannelMetaDataByKeyName(): %+v", err)
	}

	t.Logf("testDeleteChannelMetaDataByKeyName() Result: %+v", result)
}

func testCreateChannelMetaCounter(t *testing.T, channelType string, channelURL string, metaCounter map[string]int) {
	r := &CreateChannelMetaCounterRequest{
		Metacounter: metaCounter,
	}
	result, err := testChannelMetaClient.CreateChannelMetaCounter(channelType, channelURL, r)
	if err != nil {
		t.Errorf("Fail in testCreateChannelMetaCounter(): %+v", err)
	}

	t.Logf("testCreateChannelMetaCounter() Result: %+v", result)
}

func testViewChannelMetaCounter(t *testing.T, channelType string, channelURL string, keys []string) {
	r := &ViewChannelMetaCounterRequest{
		Keys: keys,
	}
	result, err := testChannelMetaClient.ViewChannelMetaCounter(channelType, channelURL, r)
	if err != nil {
		t.Errorf("Fail in testViewChannelMetaCounter(): %+v", err)
	}

	t.Logf("testViewChannelMetaCounter() Result: %+v", result)
}

func testViewChannelMetaCounterByKeyName(t *testing.T, channelType string, channelURL string, keyName string) {
	result, err := testChannelMetaClient.ViewChannelMetaCounterByKeyName(channelType, channelURL, keyName)
	if err != nil {
		t.Errorf("Fail in testViewChannelMetaCounterByKeyName(): %+v", err)
	}

	t.Logf("testViewChannelMetaCounterByKeyName() Result: %+v", result)
}

func testUpdateChannelMetaCounter(t *testing.T, channelType string, channelURL string, updatedMetaCounter map[string]int, upsert bool) {
	r := &UpdateChannelMetaCounterRequest{
		Metacounter: updatedMetaCounter,
		Upsert:      upsert,
	}
	result, err := testChannelMetaClient.UpdateChannelMetaCounter(channelType, channelURL, r)
	if err != nil {
		t.Errorf("Fail in testUpdateChannelMetaCounter(): %+v", err)
	}

	t.Logf("testUpdateChannelMetaCounter() Result: %+v", result)
}

func testUpdateChannelMetaCounterByKeyName(t *testing.T, channelType string, channelURL string, keyName string, updatedValue int, upsert bool) {
	r := &UpdateChannelMetaCounterByKeyNameRequest{
		Value:  updatedValue,
		Upsert: upsert,
	}
	result, err := testChannelMetaClient.UpdateChannelMetaCounterByKeyName(channelType, channelURL, keyName, r)
	if err != nil {
		t.Errorf("Fail in testUpdateChannelMetaCounterByKeyName(): %+v", err)
	}

	t.Logf("testUpdateChannelMetaCounterByKeyName() Result: %+v", result)
}

func testDeleteChannelMetaCounter(t *testing.T, channelType string, channelURL string) {
	result, err := testChannelMetaClient.DeleteChannelMetaCounter(channelType, channelURL)
	if err != nil && !strings.Contains(err.Error(), SendbirdAPIErrorChannelNotFound) {
		t.Errorf("Fail in testDeleteChannelMetaCounter(): %+v", err)
	}

	t.Logf("testDeleteChannelMetaCounter() Result: %+v", result)
}

func testDeleteChannelMetaCounterByKeyName(t *testing.T, channelType string, channelURL string, keyName string) {
	result, err := testChannelMetaClient.DeleteChannelMetaCounterByKeyName(channelType, channelURL, keyName)
	if err != nil && !strings.Contains(err.Error(), SendbirdAPIErrorChannelNotFound) {
		t.Errorf("Fail in testDeleteChannelMetaCounterByKeyName(): %+v", err)
	}

	t.Logf("testDeleteChannelMetaCounterByKeyName() Result: %+v", result)
}
