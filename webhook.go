package sendbirdclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//Sendbird Webhook Payload Category
type webhookCategory string

const (
	WebhookCategoryOpenChannelMsgSend        webhookCategory = "open_channel:message_send"
	WebhookCategoryGroupChannelMsgSend       webhookCategory = "group_channel:message_send"
	WebhookCategoryOpenChannelMsgDeleted     webhookCategory = "open_channel:message_delete"
	WebhookCategoryGroupChannelMsgDeleted    webhookCategory = "group_channel:message_delete"
	WebhookCategoryGroupChannelMsgRead       webhookCategory = "group_channel:message_read"
	WebhookCategoryOpenChannelCreated        webhookCategory = "open_channel:create"
	WebhookCategoryGroupChannelCreated       webhookCategory = "group_channel:create"
	WebhookCategoryOpenChannelRemoved        webhookCategory = "open_channel:remove"
	WebhookCategoryGroupChannelInvited       webhookCategory = "group_channel:invite"
	WebhookCategoryGroupChannelJoined        webhookCategory = "group_channel:join"
	WebhookCategoryGroupChannelDeclineInvite webhookCategory = "group_channel:decline_invite"
	WebhookCategoryUserBlocked               webhookCategory = "user:block"
	WebhookCategoryUserUnblocked             webhookCategory = "user:unblock"
	WebhookCategoryUserMsgRateLimitExceeded  webhookCategory = "alert:user_message_rate_limit_exceeded"
)

type WebhookCallback func(message interface{}) error

// var WebhookStub = func(map[string]interface{}) error { return nil }

type WebhookHelper struct {
	routing map[webhookCategory][]WebhookCallback
}

func (wh *WebhookHelper) SendbirdWebhook(w http.ResponseWriter, req *http.Request) {
	fmt.Println("listJSON Endpoint: ", req.RemoteAddr)

	if req.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if req.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "We only support application/json format in POST.", http.StatusUnsupportedMediaType)
		return
	}

	if req.Body == nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	var pl map[string]interface{}
	err = json.Unmarshal(body, &pl)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Webhook Message: %+v \n", pl)

	//callbacks, _ := wh.routing[pl["category"].(webhookCategory)]
	callbacks, _ := wh.routing[webhookCategory(pl["category"].(string))]
	for _, callback := range callbacks {
		callback(pl)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (h *WebhookHelper) Subscribe(callbackTopic webhookCategory, callback WebhookCallback) {
	callbacks, ok := h.routing[callbackTopic]
	if !ok {
		h.routing[callbackTopic] = []WebhookCallback{callback}
	}

	for _, c := range callbacks {
		if &c == &callback {
			return
		}
	}

	h.routing[callbackTopic] = append(callbacks, callback)
}

func NewWebhookHelper() *WebhookHelper {
	wh := &WebhookHelper{
		routing: make(map[webhookCategory][]WebhookCallback),
	}

	return wh
}
