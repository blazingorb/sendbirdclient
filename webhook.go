package sendbirdclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WebhookHelper struct {
	OnOpenChannelMsgSend            func(payload map[string]string) error
	OnGroupChannelMsgSend           func(payload map[string]string) error
	OnOpenChannelMsgDeleted         func(payload map[string]string) error
	OnGroupChannelMsgDeleted        func(payload map[string]string) error
	OnGroupChannelMsgRead           func(payload map[string]string) error
	OnOpenChannelCreated            func(payload map[string]string) error
	OnGroupChannelCreated           func(payload map[string]string) error
	OnOpenChannelRemoved            func(payload map[string]string) error
	OnGroupChannelInvited           func(payload map[string]string) error
	OnGroupChannelJoined            func(payload map[string]string) error
	OnGroupChannelDeclineInvite     func(payload map[string]string) error
	OnUserBlocked                   func(payload map[string]string) error
	OnUserUnblocked                 func(payload map[string]string) error
	OnAlertUserMsgRateLimitExceeded func(payload map[string]string) error
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

	var pl map[string]string
	err = json.Unmarshal(body, &pl)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Payload: %+v \n", pl)

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
		OnOpenChannelMsgSend:            func(payload map[string]string) error { return nil },
		OnGroupChannelMsgSend:           func(payload map[string]string) error { return nil },
		OnOpenChannelMsgDeleted:         func(payload map[string]string) error { return nil },
		OnGroupChannelMsgDeleted:        func(payload map[string]string) error { return nil },
		OnGroupChannelMsgRead:           func(payload map[string]string) error { return nil },
		OnOpenChannelCreated:            func(payload map[string]string) error { return nil },
		OnGroupChannelCreated:           func(payload map[string]string) error { return nil },
		OnOpenChannelRemoved:            func(payload map[string]string) error { return nil },
		OnGroupChannelInvited:           func(payload map[string]string) error { return nil },
		OnGroupChannelJoined:            func(payload map[string]string) error { return nil },
		OnGroupChannelDeclineInvite:     func(payload map[string]string) error { return nil },
		OnUserBlocked:                   func(payload map[string]string) error { return nil },
		OnUserUnblocked:                 func(payload map[string]string) error { return nil },
		OnAlertUserMsgRateLimitExceeded: func(payload map[string]string) error { return nil },
	}

	return wh
}
