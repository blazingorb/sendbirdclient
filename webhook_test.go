package sendbirdclient_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/blazingorb/sendbirdclient"
)

var TestRow1 = []byte(`{"sdk": "JavaScript", "type": "MESG", "app_id": "11D3F84B-6C8E-4F39-B96B-9743C7820C53", "sender": {"user_id": "TestAdmin1", "metadata": {}, "nickname": "TestAdmin1", "profile_url": ""}, "channel": {"data": "", "name": "private_TestAdmin1_Test1", "channel_url": "sendbird_group_channel_55246620_2e5e5d73799da93c9cbd1df51e4a6d446c2c659c", "custom_type": "", "is_distinct": true}, "payload": {"data": "", "message": "Hello World", "created_at": 1514431258467, "message_id": 1287670462, "custom_type": "", "translations": {}}, "category": "group_channel:message_send", "custom_type": ""}`)
var TestRow2 = []byte(`{"sdk": "JavaScript", "type": "MESG", "app_id": "11D3F84B-6C8E-4F39-B96B-9743C7820C53", "sender": {"user_id": "TestAdmin1", "metadata": {}, "nickname": "TestAdmin1", "profile_url": ""}, "channel": {"data": "", "name": "private_TestAdmin1_Test1", "channel_url": "sendbird_group_channel_55246620_2e5e5d73799da93c9cbd1df51e4a6d446c2c659c", "custom_type": "", "is_distinct": true}, "payload": {"data": "", "message": "TEST", "created_at": 1514431295237, "message_id": 1287671774, "custom_type": "", "translations": {}}, "category": "group_channel:message_send", "custom_type": ""}`)
var TestRow3 = []byte(`{"sdk": "JavaScript", "type": "MESG", "app_id": "11D3F84B-6C8E-4F39-B96B-9743C7820C53", "sender": {"user_id": "TestAdmin1", "metadata": {}, "nickname": "TestAdmin1", "profile_url": ""}, "channel": {"data": "", "name": "private_TestAdmin1_Test1", "channel_url": "sendbird_group_channel_55246620_2e5e5d73799da93c9cbd1df51e4a6d446c2c659c", "custom_type": "", "is_distinct": true}, "payload": {"data": "", "message": "fgsgdsf", "created_at": 1514441822266, "message_id": 1288026128, "custom_type": "", "translations": {}}, "category": "group_channel:message_send", "custom_type": ""}`)

var testHelper = prepareTestHelper()

func prepareTestHelper() *WebhookHelper {
	helper := NewWebhookHelper()
	testFunc := func(message interface{}) error {
		return nil
	}

	helper.Subscribe(WebhookCategoryGroupChannelMsgSend, testFunc)
	return helper
}

func TestWebhookWithoutPostMethod(t *testing.T) {
	req, err := http.NewRequest("GET", "/sendbirdwebhook", nil)
	if err != nil {
		t.Error("Request Creation Failed: ", err)
	}

	reqr := httptest.NewRecorder()

	http.HandlerFunc(testHelper.SendbirdWebhook).ServeHTTP(reqr, req)
	if status := reqr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("Status code differs. Expected %d \n Got %d", http.StatusMethodNotAllowed, status)
	}
}

func TestWebhookWithoutHeader(t *testing.T) {
	req, err := http.NewRequest("POST", "/sendbirdwebhook", nil)
	if err != nil {
		t.Error("Request Creation Failed: ", err)
	}

	reqr := httptest.NewRecorder()

	http.HandlerFunc(testHelper.SendbirdWebhook).ServeHTTP(reqr, req)
	if status := reqr.Code; status != http.StatusUnsupportedMediaType {
		t.Errorf("Status code differs. Expected %d \n Got %d", http.StatusUnsupportedMediaType, status)
	}
}

func TestWebhookWithEmptyBody(t *testing.T) {
	req, err := http.NewRequest("POST", "/sendbirdwebhook", nil)
	if err != nil {
		t.Error("Request Creation Failed: ", err)
	}
	req.Header.Set("Content-Type", "application/json")

	reqr := httptest.NewRecorder()

	http.HandlerFunc(testHelper.SendbirdWebhook).ServeHTTP(reqr, req)
	if status := reqr.Code; status != http.StatusBadRequest {
		t.Errorf("Status code differs. Expected %d \n Got %d", http.StatusBadRequest, status)
	}
}

func TestWebhookSuccess(t *testing.T) {

	req, err := http.NewRequest("POST", "/sendbirdwebhook", bytes.NewBuffer(TestRow1))
	if err != nil {
		t.Error("Request Creation Failed: ", err)
	}
	req.Header.Set("Content-Type", "application/json")

	reqr := httptest.NewRecorder()

	http.HandlerFunc(testHelper.SendbirdWebhook).ServeHTTP(reqr, req)
	if status := reqr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d \n Got %d", http.StatusOK, status)
	}
}
