package sendbirdclient

import (
	"net/url"
	"strings"

	"github.com/blazingorb/sendbirdclient/templates"
)

type channelMetaTemplateData struct {
	ChannelType string
	ChannelURL  string
	KeyName     string
}

func (c *Client) CreateAChannelMetadata(channelType string, channelURL string, r *CreateAChannelMetadataRequest) (map[string]string, error) {
	pathString, err := templates.GetChannelMetadataTemplate(channelMetaTemplateData{
		ChannelType: url.PathEscape(channelType),
		ChannelURL:  url.PathEscape(channelURL),
	}, templates.SendbirdURLChannelMetadataWithChannelTypeAndChannelURL)
	if err != nil {
		return nil, err
	}

	parsedURL := c.PrepareUrl(pathString)
	result := make(map[string]string)

	err = c.postAndReturnJSON(parsedURL, r, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type CreateAChannelMetadataRequest struct {
	Metadata map[string]string `json:"metadata"`
}

func (c *Client) ViewChannelMetadata(channelType string, channelURL string, r *ViewChannelMetadataRequest) (map[string]string, error) {
	pathString, err := templates.GetChannelMetadataTemplate(channelMetaTemplateData{
		ChannelType: url.PathEscape(channelType),
		ChannelURL:  url.PathEscape(channelURL),
	}, templates.SendbirdURLChannelMetadataWithChannelTypeAndChannelURL)

	raw := r.params().Encode()
	//#HACK special case handling for sendbird API
	raw = strings.Replace(raw, "%2C", ",", -1)

	result := make(map[string]string)

	err = c.getAndReturnJSON(c.PrepareUrl(pathString), raw, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *ViewChannelMetadataRequest) params() url.Values {
	q := make(url.Values)

	if r.Keys != nil && len(r.Keys) > 0 {
		q.Set("keys", strings.Join(r.Keys, ","))
	}

	return q
}

type ViewChannelMetadataRequest struct {
	Keys []string `json:"keys"`
}

func (c *Client) ViewChannelMetadataByKeyName(channelType string, channelURL string, keyName string) (map[string]string, error) {
	pathString, err := templates.GetChannelMetadataTemplate(channelMetaTemplateData{
		ChannelType: url.PathEscape(channelType),
		ChannelURL:  url.PathEscape(channelURL),
		KeyName:     url.PathEscape(keyName),
	}, templates.SendbirdURLChannelMetadataWithChannelTypeAndChannelURLAndKeyName)

	result := make(map[string]string)

	err = c.getAndReturnJSON(c.PrepareUrl(pathString), "", &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) UpdateChannelMetaData(channelType string, channelURL string, r *UpdateChannelMetaDataRequest) (map[string]string, error) {
	pathString, err := templates.GetChannelMetadataTemplate(channelMetaTemplateData{
		ChannelType: url.PathEscape(channelType),
		ChannelURL:  url.PathEscape(channelURL),
	}, templates.SendbirdURLChannelMetadataWithChannelTypeAndChannelURL)
	if err != nil {
		return nil, err
	}

	parsedURL := c.PrepareUrl(pathString)
	result := make(map[string]string)

	err = c.putAndReturnJSON(parsedURL, r, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type UpdateChannelMetaDataRequest struct {
	Metadata map[string]string `json:"metadata"`
	Upsert   bool              `json:"upsert,omitempty"`
}

func (c *Client) UpdateChannelMetaDataByKeyName(channelType string, channelURL string, keyName string, r *UpdateChannelMetaDataByKeyNameRequest) (map[string]string, error) {
	pathString, err := templates.GetChannelMetadataTemplate(channelMetaTemplateData{
		ChannelType: url.PathEscape(channelType),
		ChannelURL:  url.PathEscape(channelURL),
		KeyName:     url.PathEscape(keyName),
	}, templates.SendbirdURLChannelMetadataWithChannelTypeAndChannelURLAndKeyName)
	if err != nil {
		return nil, err
	}

	parsedURL := c.PrepareUrl(pathString)
	result := make(map[string]string)

	err = c.putAndReturnJSON(parsedURL, r, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type UpdateChannelMetaDataByKeyNameRequest struct {
	Value  string `json:"value"`
	Upsert bool   `json:"upsert,omitempty"`
}

func (c *Client) DeleteChannelMetaData(channelType string, channelURL string) (sendbirdErrorResponse, error) {
	pathString, err := templates.GetChannelMetadataTemplate(channelMetaTemplateData{
		ChannelType: url.PathEscape(channelType),
		ChannelURL:  url.PathEscape(channelURL),
	}, templates.SendbirdURLChannelMetadataWithChannelTypeAndChannelURL)
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

func (c *Client) DeleteChannelMetaDataByKeyName(channelType string, channelURL string, keyName string) (sendbirdErrorResponse, error) {
	pathString, err := templates.GetChannelMetadataTemplate(channelMetaTemplateData{
		ChannelType: url.PathEscape(channelType),
		ChannelURL:  url.PathEscape(channelURL),
		KeyName:     url.PathEscape(keyName),
	}, templates.SendbirdURLChannelMetadataWithChannelTypeAndChannelURLAndKeyName)
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

func (c *Client) CreateChannelMetaCounter(channelType string, channelURL string, r *CreateChannelMetaCounterRequest) (map[string]int, error) {
	pathString, err := templates.GetChannelMetadataTemplate(channelMetaTemplateData{
		ChannelType: url.PathEscape(channelType),
		ChannelURL:  url.PathEscape(channelURL),
	}, templates.SendbirdURLChannelMetaCounterWithChannelTypeAndChannelURL)
	if err != nil {
		return nil, err
	}

	parsedURL := c.PrepareUrl(pathString)
	result := make(map[string]int)

	err = c.postAndReturnJSON(parsedURL, r, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type CreateChannelMetaCounterRequest struct {
	Metacounter map[string]int `json:"metacounter"`
}

func (c *Client) ViewChannelMetaCounter(channelType string, channelURL string, r *ViewChannelMetaCounterRequest) (map[string]int, error) {
	pathString, err := templates.GetChannelMetadataTemplate(channelMetaTemplateData{
		ChannelType: url.PathEscape(channelType),
		ChannelURL:  url.PathEscape(channelURL),
	}, templates.SendbirdURLChannelMetaCounterWithChannelTypeAndChannelURL)

	raw := r.params().Encode()
	//#HACK special case handling for sendbird API
	raw = strings.Replace(raw, "%2C", ",", -1)

	result := make(map[string]int)

	err = c.getAndReturnJSON(c.PrepareUrl(pathString), raw, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *ViewChannelMetaCounterRequest) params() url.Values {
	q := make(url.Values)

	if r.Keys != nil && len(r.Keys) > 0 {
		q.Set("keys", strings.Join(r.Keys, ","))
	}

	return q
}

type ViewChannelMetaCounterRequest struct {
	Keys []string `json:"keys"`
}

func (c *Client) ViewChannelMetaCounterByKeyName(channelType string, channelURL string, keyName string) (map[string]int, error) {
	pathString, err := templates.GetChannelMetadataTemplate(channelMetaTemplateData{
		ChannelType: url.PathEscape(channelType),
		ChannelURL:  url.PathEscape(channelURL),
		KeyName:     url.PathEscape(keyName),
	}, templates.SendbirdURLChannelMetaCounterWithChannelTypeAndChannelURLAndKeyName)

	result := make(map[string]int)

	err = c.getAndReturnJSON(c.PrepareUrl(pathString), "", &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) UpdateChannelMetaCounter(channelType string, channelURL string, r *UpdateChannelMetaCounterRequest) (map[string]int, error) {
	pathString, err := templates.GetChannelMetadataTemplate(channelMetaTemplateData{
		ChannelType: url.PathEscape(channelType),
		ChannelURL:  url.PathEscape(channelURL),
	}, templates.SendbirdURLChannelMetaCounterWithChannelTypeAndChannelURL)

	result := make(map[string]int)

	err = c.putAndReturnJSON(c.PrepareUrl(pathString), r, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type UpdateChannelMetaCounterRequest struct {
	Metacounter map[string]int `json:"metacounter"`
	Mode        string         `json:"mode,omitempty"`
	Upsert      bool           `json:"upsert,omitempty"`
}

func (c *Client) UpdateChannelMetaCounterByKeyName(channelType string, channelURL string, keyName string, r *UpdateChannelMetaCounterByKeyNameRequest) (map[string]int, error) {
	pathString, err := templates.GetChannelMetadataTemplate(channelMetaTemplateData{
		ChannelType: url.PathEscape(channelType),
		ChannelURL:  url.PathEscape(channelURL),
		KeyName:     url.PathEscape(keyName),
	}, templates.SendbirdURLChannelMetaCounterWithChannelTypeAndChannelURLAndKeyName)

	result := make(map[string]int)

	err = c.putAndReturnJSON(c.PrepareUrl(pathString), r, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type UpdateChannelMetaCounterByKeyNameRequest struct {
	Value  int    `json:"value"`
	Mode   string `json:"mode,omitempty"`
	Upsert bool   `json:"upsert,omitempty"`
}

func (c *Client) DeleteChannelMetaCounter(channelType string, channelURL string) (sendbirdErrorResponse, error) {
	pathString, err := templates.GetChannelMetadataTemplate(channelMetaTemplateData{
		ChannelType: url.PathEscape(channelType),
		ChannelURL:  url.PathEscape(channelURL),
	}, templates.SendbirdURLChannelMetaCounterWithChannelTypeAndChannelURL)

	result := sendbirdErrorResponse{}

	err = c.deleteAndReturnJSON(c.PrepareUrl(pathString), "", &result)
	if err != nil {
		return sendbirdErrorResponse{}, err
	}

	return result, nil
}

func (c *Client) DeleteChannelMetaCounterByKeyName(channelType string, channelURL string, keyName string) (sendbirdErrorResponse, error) {
	pathString, err := templates.GetChannelMetadataTemplate(channelMetaTemplateData{
		ChannelType: url.PathEscape(channelType),
		ChannelURL:  url.PathEscape(channelURL),
		KeyName:     url.PathEscape(keyName),
	}, templates.SendbirdURLChannelMetaCounterWithChannelTypeAndChannelURLAndKeyName)

	result := sendbirdErrorResponse{}

	err = c.deleteAndReturnJSON(c.PrepareUrl(pathString), "", &result)
	if err != nil {
		return sendbirdErrorResponse{}, err
	}

	return result, nil
}
