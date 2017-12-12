package sendbirdclient

import (
	"errors"
	"net/url"
	"sendbirdclient/templates"
)

type OpenChannel struct {
	channel

	Operators        []User `json:"operators"`
	Participants     []User `json:"participants"`
	ParticipantCount int    `json:"participant_count"`
	Freeze           bool   `json:"freeze"`
}

type openChannelsTemplateData struct {
	ChannelURL   string
	BannedUserID string
	MutedUserID  string
}

func (c *Client) CreateAnOpenChannelWithURL(r *CreateAnOpenChannelWithURLRequest) (OpenChannel, error) {
	result := OpenChannel{}

	if err := c.postAndReturnJSON(c.PrepareUrl(SendbirdURLOpenChannels), r, &result); err != nil {
		return OpenChannel{}, err
	}

	return result, nil
}

type CreateAnOpenChannelWithURLRequest struct {
	Name       string   `json:"name,omitempty"`
	ChannelURL string   `json:"channel_url,omitempty"`
	CoverURL   string   `json:"cover_url,omitempty"`
	CustomType string   `json:"custom_type,omitempty"`
	Data       string   `json:"data,omitempty"`
	Operators  []string `json:"operators,omitempty"`
}

func (c *Client) CreateAnOpenChannelWithFile(r *CreateAnOpenChannelWithFileRequest) (OpenChannel, error) {
	return OpenChannel{}, errors.New(SendbirdClientErrorNotImplemented)
}

type CreateAnOpenChannelWithFileRequest struct {
	Name       string   `json:"name,omitempty"`
	ChannelURL string   `json:"channel_url,omitempty"`
	CoverFile  []byte   `json:"cover_file,omitempty"`
	CustomType string   `json:"custom_type,omitempty"`
	Data       string   `json:"data,omitempty"`
	Operators  []string `json:"operators,omitempty"`
}

func (c *Client) ListOpenChannels(r *ListOpenChannelsRequest) (ListOpenChannelsResponse, error) {
	result := ListOpenChannelsResponse{}

	raw := r.params().Encode()
	if err := c.getAndReturnJSON(c.PrepareUrl(SendbirdURLOpenChannels), raw, &result); err != nil {
		return ListOpenChannelsResponse{}, err
	}

	return result, nil
}

func (r *ListOpenChannelsRequest) params() url.Values {
	q := make(url.Values)

	if r.Token != "" {
		q.Set("token", r.Token)
	}

	if r.Limit >= ListLimitLowerBound && r.Limit <= ListLimitUpperBound {
		q.Set("limit", string(r.Limit))
	}

	if r.CustomType != "" {
		q.Set("custom_type", r.CustomType)
	}

	return q
}

type ListOpenChannelsRequest struct {
	Token      string `json:"token,omitempty"`
	Limit      int    `json:"limit,omitempty"`
	CustomType string `json:"custom_type,omitempty"`
}

type ListOpenChannelsResponse struct {
	Channels []OpenChannel `json:"channels"`
	Next     string        `json:"next"`
}

func (c *Client) UpdateAnOpenChannelWithURL(channelURL string, r *UpdateAnOpenChannelWithURLRequest) (OpenChannel, error) {
	pathString, err := templates.GetOpenChannelTemplate(openChannelsTemplateData{ChannelURL: url.PathEscape(channelURL)}, templates.SendbirdURLOpenChannelsWithChannelURL)
	if err != nil {
		return OpenChannel{}, err
	}

	parsedURL := c.PrepareUrl(pathString)

	result := OpenChannel{}

	if err := c.putAndReturnJSON(parsedURL, r, &result); err != nil {
		return OpenChannel{}, err
	}

	return result, nil
}

type UpdateAnOpenChannelWithURLRequest struct {
	Name       string   `json:"name,omitempty"`
	CoverURL   string   `json:"cover_url,omitempty"`
	CustomType string   `json:"custom_type,omitempty"`
	Data       string   `json:"data,omitempty"`
	Operators  []string `json:"operators,omitempty"`
}

func (c *Client) UpdateAnOpenChannelWithFile(channelURL string, r *UpdateAnOpenChannelWithFileRequest) (OpenChannel, error) {
	return OpenChannel{}, errors.New(SendbirdClientErrorNotImplemented)
}

type UpdateAnOpenChannelWithFileRequest struct {
	Name       string   `json:"name,omitempty"`
	CustomType string   `json:"custom_type,omitempty"`
	CoverFile  string   `json:"cover_file,omitempty"`
	Data       string   `json:"data,omitempty"`
	Operators  []string `json:"operators,omitempty"`
}

func (c *Client) ViewAnOpenChannel(channelURL string, r *ViewAnOpenChannelRequest) (OpenChannel, error) {
	pathString, err := templates.GetOpenChannelTemplate(openChannelsTemplateData{ChannelURL: url.PathEscape(channelURL)}, templates.SendbirdURLOpenChannelsWithChannelURL)
	if err != nil {
		return OpenChannel{}, err
	}

	parsedURL := c.PrepareUrl(pathString)

	result := OpenChannel{}

	raw := r.params().Encode()
	if err := c.getAndReturnJSON(parsedURL, raw, &result); err != nil {
		return OpenChannel{}, err
	}

	return result, nil
}

func (r *ViewAnOpenChannelRequest) params() url.Values {
	q := make(url.Values)

	if r.Participants {
		q.Set("participants", "true")
	}

	return q
}

type ViewAnOpenChannelRequest struct {
	Participants bool `json:"participants,omitempty"`
}

func (c *Client) DeleteAnOpenChannel(channelURL string) (sendbirdErrorResponse, error) {
	pathString, err := templates.GetOpenChannelTemplate(openChannelsTemplateData{ChannelURL: url.PathEscape(channelURL)}, templates.SendbirdURLOpenChannelsWithChannelURL)
	if err != nil {
		return sendbirdErrorResponse{}, err
	}

	parsedURL := c.PrepareUrl(pathString)

	result := sendbirdErrorResponse{}

	if err := c.deleteAndReturnJSON(parsedURL, "", &result); err != nil {
		return sendbirdErrorResponse{}, err
	}

	return result, nil
}

func (c *Client) ListOpenChannelParticipants(channelURL string, r *ListOpenChannelParticipantsRequest) (ListOpenChannelParticipantsResponse, error) {
	pathString, err := templates.GetOpenChannelTemplate(openChannelsTemplateData{ChannelURL: url.PathEscape(channelURL)}, templates.SendbirdURLOpenChannelsParticipantsWithChannelURL)
	if err != nil {
		return ListOpenChannelParticipantsResponse{}, err
	}

	parsedURL := c.PrepareUrl(pathString)

	result := ListOpenChannelParticipantsResponse{}

	raw := r.params().Encode()
	if err := c.getAndReturnJSON(parsedURL, raw, &result); err != nil {
		return ListOpenChannelParticipantsResponse{}, err
	}

	return result, nil
}

func (r *ListOpenChannelParticipantsRequest) params() url.Values {
	q := make(url.Values)

	if r.Token != "" {
		q.Set("token", r.Token)
	}

	if r.Limit >= 1 && r.Limit <= 100 {
		q.Set("limit", string(r.Limit))
	}

	return q
}

type ListOpenChannelParticipantsRequest struct {
	Token string `json:"token,omitempty"`
	Limit int    `json:"limit,omitempty"`
}

type ListOpenChannelParticipantsResponse struct {
	Participants []User `json:"participants"`
	Next         string `json:"next"`
}

func (c *Client) FreezeAnOpenChannel(channelURL string, r *FreezeAnOpenChannelRequest) (OpenChannel, error) {
	pathString, err := templates.GetOpenChannelTemplate(openChannelsTemplateData{ChannelURL: url.PathEscape(channelURL)}, templates.SendbirdURLOpenChannelsFreezeWithChannelURL)
	if err != nil {
		return OpenChannel{}, err
	}

	parsedURL := c.PrepareUrl(pathString)

	result := OpenChannel{}

	if err := c.putAndReturnJSON(parsedURL, r, &result); err != nil {
		return OpenChannel{}, err
	}

	return result, nil
}

type FreezeAnOpenChannelRequest struct {
	Freeze bool `json:"freeze"`
}

func (c *Client) BanAUserInOpenChannel(channelURL string, r *BanAUserInOpenChannelRequest) (BanAUserInOpenChannelResponse, error) {
	pathString, err := templates.GetOpenChannelTemplate(openChannelsTemplateData{ChannelURL: url.PathEscape(channelURL)}, templates.SendbirdURLOpenChannelsBanWithChannelURL)
	if err != nil {
		return BanAUserInOpenChannelResponse{}, err
	}

	parsedURL := c.PrepareUrl(pathString)

	result := BanAUserInOpenChannelResponse{}

	if err := c.postAndReturnJSON(parsedURL, r, &result); err != nil {
		return BanAUserInOpenChannelResponse{}, err
	}

	return result, nil
}

type BanAUserInOpenChannelRequest struct {
	UserID      string `json:"user_id"`
	Seconds     int    `json:"seconds,omitempty"`
	Description string `json:"description,omitempty"`
	NextURL     string `json:"next_url,omitempty"`
}

type BanAUserInOpenChannelResponse struct {
	Description string `json:"description,omitempty"`
	User        User   `json:"user,omitempty"`
	StartAt     int64  `json:"start_at,omitempty"`
	EndAt       int64  `json:"end_at,omitempty"`
	NextURL     string `json:"next_url,omitempty"`
}

func (c *Client) ListBannedUsersInOpenChannel(channelURL string, r *ListBannedUsersInOpenChannelRequest) (ListBannedUsersInOpenChannelResponse, error) {
	pathString, err := templates.GetOpenChannelTemplate(openChannelsTemplateData{ChannelURL: url.PathEscape(channelURL)}, templates.SendbirdURLOpenChannelsBanWithChannelURL)
	if err != nil {
		return ListBannedUsersInOpenChannelResponse{}, err
	}

	parsedURL := c.PrepareUrl(pathString)

	result := ListBannedUsersInOpenChannelResponse{}

	raw := r.params().Encode()
	if err := c.getAndReturnJSON(parsedURL, raw, &result); err != nil {
		return ListBannedUsersInOpenChannelResponse{}, err
	}

	return result, nil
}

func (r *ListBannedUsersInOpenChannelRequest) params() url.Values {
	q := make(url.Values)

	if r.Token != "" {
		q.Set("token", r.Token)
	}

	if r.Limit >= 1 && r.Limit <= 100 {
		q.Set("limit", string(r.Limit))
	}

	return q
}

type ListBannedUsersInOpenChannelRequest struct {
	Token string `json:"token,omitempty"`
	Limit int    `json:"limit,omitempty"`
}

type ListBannedUsersInOpenChannelResponse struct {
	BannedList []BannedResult `json:"banned_list"`
	Next       string         `json:"next"`
}

type BannedResult struct {
	Description string `json:"description,omitempty"`
	User        User   `json:"user,omitempty"`
	StartAt     int64  `json:"start_at,omitempty"`
	EndAt       int64  `json:"end_at,omitempty"`
}

func (c *Client) UpdateBanInOpenChannel(channelURL string, bannedUserID string, r *UpdateBanInOpenChannelRequest) (BannedResult, error) {
	pathString, err := templates.GetOpenChannelTemplate(openChannelsTemplateData{
		ChannelURL:   url.PathEscape(channelURL),
		BannedUserID: url.PathEscape(bannedUserID),
	}, templates.SendbirdURLOpenChannelsBanWithChannelURLandBannedUserID)
	if err != nil {
		return BannedResult{}, err
	}

	parsedURL := c.PrepareUrl(pathString)

	result := BannedResult{}

	if err := c.putAndReturnJSON(parsedURL, r, &result); err != nil {
		return BannedResult{}, err
	}

	return result, nil
}

type UpdateBanInOpenChannelRequest struct {
	Description string `json:"description,omitempty"`
	Seconds     int    `json:"seconds"`
}

func (c *Client) ViewBanInOpenChannel(channelURL string, bannedUserID string) (BannedResult, error) {
	pathString, err := templates.GetOpenChannelTemplate(openChannelsTemplateData{
		ChannelURL:   url.PathEscape(channelURL),
		BannedUserID: url.PathEscape(bannedUserID),
	}, templates.SendbirdURLOpenChannelsBanWithChannelURLandBannedUserID)
	if err != nil {
		return BannedResult{}, err
	}

	parsedURL := c.PrepareUrl(pathString)

	result := BannedResult{}

	if err := c.getAndReturnJSON(parsedURL, "", &result); err != nil {
		return BannedResult{}, err
	}

	return result, nil
}

func (c *Client) UnbanAUserInOpenChannel(channelURL string, bannedUserID string) (sendbirdErrorResponse, error) {
	pathString, err := templates.GetOpenChannelTemplate(openChannelsTemplateData{
		ChannelURL:   url.PathEscape(channelURL),
		BannedUserID: url.PathEscape(bannedUserID),
	}, templates.SendbirdURLOpenChannelsBanWithChannelURLandBannedUserID)
	if err != nil {
		return sendbirdErrorResponse{}, err
	}

	parsedURL := c.PrepareUrl(pathString)

	result := sendbirdErrorResponse{}

	if err := c.deleteAndReturnJSON(parsedURL, "", &result); err != nil {
		return sendbirdErrorResponse{}, err
	}

	return result, nil
}

func (c *Client) MuteAUserInOpenChannel(channelURL string, r *MuteAUserInOpenChannelRequest) (OpenChannel, error) {
	pathString, err := templates.GetOpenChannelTemplate(openChannelsTemplateData{ChannelURL: url.PathEscape(channelURL)}, templates.SendbirdURLOpenChannelsMuteWithChannelURL)
	if err != nil {
		return OpenChannel{}, err
	}
	parsedURL := c.PrepareUrl(pathString)

	result := OpenChannel{}

	if err := c.postAndReturnJSON(parsedURL, r, &result); err != nil {
		return OpenChannel{}, err
	}

	return result, nil
}

type MuteAUserInOpenChannelRequest struct {
	UserID string `json:"user_id"`
}

func (c *Client) ListMutedUsersInOpenChannel(channelURL string, r *ListMutedUsersInOpenChannelRequest) (ListMutedUsersInOpenChannelResponse, error) {
	pathString, err := templates.GetOpenChannelTemplate(openChannelsTemplateData{ChannelURL: url.PathEscape(channelURL)}, templates.SendbirdURLOpenChannelsMuteWithChannelURL)
	if err != nil {
		return ListMutedUsersInOpenChannelResponse{}, err
	}
	parsedURL := c.PrepareUrl(pathString)

	result := ListMutedUsersInOpenChannelResponse{}

	raw := r.params().Encode()
	if err := c.getAndReturnJSON(parsedURL, raw, &result); err != nil {
		return ListMutedUsersInOpenChannelResponse{}, err
	}

	return result, nil
}

func (r *ListMutedUsersInOpenChannelRequest) params() url.Values {
	q := make(url.Values)

	if r.Token != "" {
		q.Set("token", r.Token)
	}

	if r.Limit >= ListLimitLowerBound && r.Limit <= ListLimitUpperBound {
		q.Set("limit", string(r.Limit))
	}

	return q
}

type ListMutedUsersInOpenChannelRequest struct {
	Token string `json:"token,omitempty"`
	Limit int    `json:"limit,omitempty"`
}

type ListMutedUsersInOpenChannelResponse struct {
	MutedList []User `json:"muted_list"`
	Next      string `json:"next"`
}

func (c *Client) ViewAMuteInOpenChannel(channelURL string, mutedUserID string) (ViewAMuteInOpenChannelResponse, error) {
	pathString, err := templates.GetOpenChannelTemplate(openChannelsTemplateData{
		ChannelURL:  url.PathEscape(channelURL),
		MutedUserID: url.PathEscape(mutedUserID),
	}, templates.SendbirdURLOpenChannelsMuteWithChannelURLandMutedUserID)
	if err != nil {
		return ViewAMuteInOpenChannelResponse{}, err
	}
	parsedURL := c.PrepareUrl(pathString)

	result := ViewAMuteInOpenChannelResponse{}

	if err := c.getAndReturnJSON(parsedURL, "", &result); err != nil {
		return ViewAMuteInOpenChannelResponse{}, err
	}

	return result, nil
}

type ViewAMuteInOpenChannelResponse struct {
	IsMuted bool `json:"is_muted"`
}

func (c *Client) UnmuteAUserInOpenChannel(channelURL string, mutedUserID string) (sendbirdErrorResponse, error) {
	pathString, err := templates.GetOpenChannelTemplate(openChannelsTemplateData{
		ChannelURL:  url.PathEscape(channelURL),
		MutedUserID: url.PathEscape(mutedUserID),
	}, templates.SendbirdURLOpenChannelsMuteWithChannelURLandMutedUserID)
	if err != nil {
		return sendbirdErrorResponse{}, err
	}
	parsedURL := c.PrepareUrl(pathString)

	result := sendbirdErrorResponse{}

	if err := c.deleteAndReturnJSON(parsedURL, "", &result); err != nil {
		return sendbirdErrorResponse{}, err
	}

	return result, nil
}
