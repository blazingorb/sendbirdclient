package sendbirdclient_test

import (
	. "sendbirdclient"
)

const (
	apiKey         = "Please input your Sendbird ApiKey here."
	baseURL        = "api.sendbird.com"
	testGCMToken   = "TestGCMToken"
	testAPNToken   = "TestAPNToken"
	testChannelURL = "testChannelURL"
)

func NewTestClient() *Client {
	var testClient, err = NewClient(WithAPIKey(apiKey))
	if err != nil {
		panic(err)
	}
	return testClient
}
