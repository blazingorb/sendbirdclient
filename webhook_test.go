package sendbirdclient_test

import (
	"testing"

	. "github.com/blazingorb/sendbirdclient"
)

func TestWebhook(t *testing.T) {
	helper := NewWebhookHelper()

	testFunc := func(message map[string]interface{}) error {

		return nil
	}

	helper.OnAlertUserMsgRateLimitExceeded = testFunc
	//helper.SendbirdWebhook()
}
