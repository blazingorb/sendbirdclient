package sendbirdclient_test

import (
	"os"

	. "github.com/blazingorb/sendbirdclient"
)

const (
	//apiKey         = "Please input your Sendbird ApiKey here."
	apiKeyEnvVariableName = "SENDBIRD_API_KEY"
	baseURL               = "api.sendbird.com"
	testGCMToken          = "TestGCMToken"
	testAPNToken          = "TestAPNToken"
	testChannelURL        = "testChannelURL"
)

func NewTestClient() *Client {
	apiKey := os.Getenv(apiKeyEnvVariableName)
	var testClient, err = NewClient(WithAPIKey(apiKey))
	if err != nil {
		panic(err)
	}
	return testClient
}
