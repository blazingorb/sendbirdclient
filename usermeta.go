package sendbirdclient

import (
	"net/url"
	"sendbirdclient/templates"
	"strings"
)

type userMetaTemplateData struct {
	UserID  string
	KeyName string
}

func (c *Client) CreateAnUserMetaData(userID string, r *CreateAnUserMetaDataRequest) (map[string]string, error) {
	pathString, err := templates.GetUsersTemplate(userMetaTemplateData{UserID: url.PathEscape(userID)}, templates.SendbirdURLUserMetadataWithUserID)
	if err != nil {
		return nil, err
	}

	parsedURL := c.PrepareUrl(pathString)
	result := make(map[string]string)

	if err := c.postAndReturnJSON(parsedURL, r, &result); err != nil {
		return nil, err
	}

	return result, nil
}

type CreateAnUserMetaDataRequest struct {
	MetaData map[string]string `json:"metadata"`
}

func (c *Client) ListUserMetaData(userID string, r *ListUserMetaDataRequest) (map[string]string, error) {
	pathString, err := templates.GetUsersTemplate(userMetaTemplateData{UserID: url.PathEscape(userID)}, templates.SendbirdURLUserMetadataWithUserID)
	if err != nil {
		return nil, err
	}

	parsedURL := c.PrepareUrl(pathString)
	result := make(map[string]string)

	raw := r.params().Encode()
	//#HACK special case handling for sendbird API
	raw = strings.Replace(raw, "%2C", ",", -1)

	err = c.getAndReturnJSON(parsedURL, raw, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *ListUserMetaDataRequest) params() url.Values {
	q := make(url.Values)

	if r.Keys != nil && len(r.Keys) > 0 {
		q.Set("keys", strings.Join(r.Keys, ","))
	}

	return q
}

type ListUserMetaDataRequest struct {
	Keys []string `json:"keys,omitempty"`
}

func (c *Client) ViewAnUserMetaData(userID string, keyName string) (map[string]string, error) {
	pathString, err := templates.GetUsersTemplate(userMetaTemplateData{
		UserID:  url.PathEscape(userID),
		KeyName: url.PathEscape(keyName),
	}, templates.SendbirdURLUserMetadataWithUserIDandKeyName)
	if err != nil {
		return nil, err
	}

	parsedURL := c.PrepareUrl(pathString)
	result := make(map[string]string)

	err = c.getAndReturnJSON(parsedURL, "", &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) UpdateUserMetaData(userID string, r *UpdateUserMetaDataRequest) (map[string]string, error) {
	pathString, err := templates.GetUsersTemplate(userMetaTemplateData{UserID: url.PathEscape(userID)}, templates.SendbirdURLUserMetadataWithUserID)
	if err != nil {
		return nil, err
	}

	parsedURL := c.PrepareUrl(pathString)
	result := make(map[string]string)

	if err := c.putAndReturnJSON(parsedURL, r, &result); err != nil {
		return nil, err
	}

	return result, nil
}

type UpdateUserMetaDataRequest struct {
	MetaData interface{} `json:"metadata"`
	Upsert   bool        `json:"upsert,omitempty"`
}

func (c *Client) UpdateUserMetaDataItem(userID string, keyName string, r *UpdateUserMetaDataItemRequest) (map[string]string, error) {
	pathString, err := templates.GetUsersTemplate(userMetaTemplateData{
		UserID:  url.PathEscape(userID),
		KeyName: url.PathEscape(keyName),
	}, templates.SendbirdURLUserMetadataWithUserIDandKeyName)
	if err != nil {
		return nil, err
	}

	parsedURL := c.PrepareUrl(pathString)
	result := make(map[string]string)

	if err := c.putAndReturnJSON(parsedURL, r, &result); err != nil {
		return nil, err
	}

	return result, nil
}

type UpdateUserMetaDataItemRequest struct {
	Value  string `json:"value"`
	Upsert bool   `json:"upsert,omitempty"`
}

func (c *Client) DeleteUserMetaData(userID string) (sendbirdErrorResponse, error) {
	pathString, err := templates.GetUsersTemplate(userMetaTemplateData{UserID: url.PathEscape(userID)}, templates.SendbirdURLUserMetadataWithUserID)
	if err != nil {
		return sendbirdErrorResponse{}, err
	}

	parsedURL := c.PrepareUrl(pathString)
	result := sendbirdErrorResponse{}

	err = c.deleteAndReturnJSON(parsedURL, "", &result)
	if err != nil {
		return sendbirdErrorResponse{}, err
	}

	return result, nil
}

func (c *Client) DeleteUserMetaDataItem(userID string, keyName string) (sendbirdErrorResponse, error) {
	pathString, err := templates.GetUsersTemplate(userMetaTemplateData{
		UserID:  url.PathEscape(userID),
		KeyName: url.PathEscape(keyName),
	}, templates.SendbirdURLUserMetadataWithUserIDandKeyName)
	if err != nil {
		return sendbirdErrorResponse{}, err
	}

	parsedURL := c.PrepareUrl(pathString)
	result := sendbirdErrorResponse{}

	err = c.deleteAndReturnJSON(parsedURL, "", &result)
	if err != nil {
		return sendbirdErrorResponse{}, err
	}

	return result, nil
}
