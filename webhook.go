package sendbirdclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WebhookHelper struct {
	OnOpenChannelMsgSend            func(message map[string]interface{}) error
	OnGroupChannelMsgSend           func(message map[string]interface{}) error
	OnOpenChannelMsgDeleted         func(message map[string]interface{}) error
	OnGroupChannelMsgDeleted        func(message map[string]interface{}) error
	OnGroupChannelMsgRead           func(message map[string]interface{}) error
	OnOpenChannelCreated            func(message map[string]interface{}) error
	OnGroupChannelCreated           func(message map[string]interface{}) error
	OnOpenChannelRemoved            func(message map[string]interface{}) error
	OnGroupChannelInvited           func(message map[string]interface{}) error
	OnGroupChannelJoined            func(message map[string]interface{}) error
	OnGroupChannelDeclineInvite     func(message map[string]interface{}) error
	OnUserBlocked                   func(message map[string]interface{}) error
	OnUserUnblocked                 func(message map[string]interface{}) error
	OnAlertUserMsgRateLimitExceeded func(message map[string]interface{}) error
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

	switch pl["category"] {
	case WebhookCategoryOpenChannelMsgSend:
		err = wh.OnOpenChannelMsgSend(pl)
	case WebhookCategoryGroupChannelMsgSend:
		err = wh.OnGroupChannelMsgSend(pl)
	case WebhookCategoryOpenChannelMsgDeleted:
		err = wh.OnOpenChannelMsgDeleted(pl)
	case WebhookCategoryGroupChannelMsgDeleted:
		err = wh.OnGroupChannelMsgDeleted(pl)
	case WebhookCategoryGroupChannelMsgRead:
		err = wh.OnGroupChannelMsgRead(pl)
	case WebhookCategoryOpenChannelCreated:
		err = wh.OnOpenChannelCreated(pl)
	case WebhookCategoryGroupChannelCreated:
		err = wh.OnGroupChannelCreated(pl)
	case WebhookCategoryOpenChannelRemoved:
		err = wh.OnOpenChannelRemoved(pl)
	case WebhookCategoryGroupChannelInvited:
		err = wh.OnGroupChannelInvited(pl)
	case WebhookCategoryGroupChannelJoined:
		err = wh.OnGroupChannelJoined(pl)
	case WebhookCategoryGroupChannelDeclineInvite:
		err = wh.OnGroupChannelDeclineInvite(pl)
	case WebhookCategoryUserBlocked:
		err = wh.OnUserBlocked(pl)
	case WebhookCategoryUserUnblocked:
		err = wh.OnUserUnblocked(pl)
	case WebhookCategoryUserMsgRateLimitExceeded:
		err = wh.OnAlertUserMsgRateLimitExceeded(pl)
	default:
		http.Error(w, SendbirdClientErrorUnsuppoertedWebhookCategory, http.StatusInternalServerError)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func NewWebhookHelper() *WebhookHelper {
	wh := &WebhookHelper{
		OnOpenChannelMsgSend:            func(message map[string]interface{}) error { return nil },
		OnGroupChannelMsgSend:           func(message map[string]interface{}) error { return nil },
		OnOpenChannelMsgDeleted:         func(message map[string]interface{}) error { return nil },
		OnGroupChannelMsgDeleted:        func(message map[string]interface{}) error { return nil },
		OnGroupChannelMsgRead:           func(message map[string]interface{}) error { return nil },
		OnOpenChannelCreated:            func(message map[string]interface{}) error { return nil },
		OnGroupChannelCreated:           func(message map[string]interface{}) error { return nil },
		OnOpenChannelRemoved:            func(message map[string]interface{}) error { return nil },
		OnGroupChannelInvited:           func(message map[string]interface{}) error { return nil },
		OnGroupChannelJoined:            func(message map[string]interface{}) error { return nil },
		OnGroupChannelDeclineInvite:     func(message map[string]interface{}) error { return nil },
		OnUserBlocked:                   func(message map[string]interface{}) error { return nil },
		OnUserUnblocked:                 func(message map[string]interface{}) error { return nil },
		OnAlertUserMsgRateLimitExceeded: func(message map[string]interface{}) error { return nil },
	}

	return wh
}
