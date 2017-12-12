package sendbirdclient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Sendbird API Error Response

const (
	SendbirdAPIErrorUniqueIDConstraint = "violates unique constraint"
	SendbirdAPIErrorChannelNotFound    = "Channel not found"
	SendbirdAPIErrorUserNotFound       = "User not found"

	SendbirdClientErrorKeyValueMismatch = "Key and value are mismatched."
	SendbirdClientErrorUnsupportMethod  = "Unsupport http method."
	SendbirdClientErrorNotImplemented   = "Not implemented."
)

type sendbirdErrorResponse struct {
	HasError bool   `json:"error"`
	Message  string `json:"message"`
	Code     int    `json:"code"`
}

//implement error interface
func (s sendbirdErrorResponse) Error() string {
	if s.Code != 200 && s.Code != 0 {
		return fmt.Sprintf("SendbirdError: %d - %s", s.Code, s.Message) // or s.message or some kind of format
	}
	return "{}"
}

func CheckSendbirdError(httpResp *http.Response) error {
	if httpResp.StatusCode != 200 {
		errorMessageBody := sendbirdErrorResponse{}
		err := json.NewDecoder(httpResp.Body).Decode(&errorMessageBody)
		if err != nil {
			return fmt.Errorf("CheckSendbirdError(): %s", err)
		}

		return fmt.Errorf("CheckSendbirdError(): %d - %s", errorMessageBody.Code, errorMessageBody.Message)
	}
	return nil
}
