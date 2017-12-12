package sendbirdclient

import (
	"errors"
	"net/url"
	"sendbirdclient/templates"
	"strings"
)

type User struct {
	UserID      string `json:"user_id"`
	NickName    string `json:"nickname"`
	ProfileURL  string `json:"profile_url"`
	ProfileFile []byte `json:"profile_file"`
	AccessToken string `json:"access_token"`
	IsActive    bool   `json:"is_active"`
	IsOnline    bool   `json:"is_online"`
	LastSeenAt  int64  `json:"last_seen_at"`
}

type usersTemplateData struct {
	UserID     string
	TargetID   string
	TokenType  string
	PushToken  string
	ChannelURL string
}

func (c *Client) CreateAUserWithURL(r *CreateAUserWithURLRequest) (User, error) {
	if r.UserID == "" {
		return User{}, errors.New("user: UserID missing")
	}

	result := User{}

	if err := c.postAndReturnJSON(c.PrepareUrl(SendbirdURLUsers), r, &result); err != nil {
		return User{}, err
	}

	return result, nil
}

type CreateAUserWithURLRequest struct {
	UserID           string `json:"user_id"`
	NickName         string `json:"nickname"`
	ProfileURL       string `json:"profile_url"`
	IssueAccessToken bool   `json:"issue_access_token,omitempty"`
}

func (c *Client) CreateAUserWithFile(r *CreateAUserWithFileRequest) (User, error) {
	return User{}, errors.New(SendbirdClientErrorNotImplemented)
}

type CreateAUserWithFileRequest struct {
	UserID           string `json:"user_id"`
	NickName         string `json:"nickname"`
	ProfileFile      string `json:"profile_file"`
	IssueAccessToken bool   `json:"issue_access_token,omitempty"`
}

func (c *Client) ListUsers(r *ListUsersRequest) (ListUsersResponse, error) {
	result := ListUsersResponse{}

	raw := r.params().Encode()
	//#HACK special case handling for sendbird API
	raw = strings.Replace(raw, "%2C", ",", -1)

	//err := c.getJSONWithRawQueryString(c.PrepareUrl(SendbirdURLUsers), raw, &result)
	err := c.getAndReturnJSON(c.PrepareUrl(SendbirdURLUsers), raw, &result)
	if err != nil {
		return ListUsersResponse{}, err
	}

	return result, nil
}

func (r *ListUsersRequest) params() url.Values {
	q := make(url.Values)

	if r.Token != "" {
		q.Set("token", r.Token)
	}

	if r.Limit >= ListLimitLowerBound && r.Limit <= ListLimitUpperBound {
		q.Set("limit", string(r.Limit))
	}

	if r.ActiveMode != "" {
		q.Set("active_mode", r.ActiveMode)
	}

	if r.ShowBot {
		q.Set("show_bot", "true")
	}

	if r.UserIDs != nil && len(r.UserIDs) > 0 {
		q.Set("user_ids", strings.Join(r.UserIDs, ","))
	}

	return q
}

type ListUsersRequest struct {
	Token      string   `json:"token,omitempty"`
	Limit      int      `json:"limit,omitempty"`
	ActiveMode string   `json:"active_mode,omitempty"`
	ShowBot    bool     `json:"show_bot,omitempty"`
	UserIDs    []string `json:"user_ids,omitempty"`
}

type ListUsersResponse struct {
	Users []User `json:"users"`
	Next  string `json:"next"`
}

func (c *Client) UpdateAUserWithURL(userID string, r *UpdateAUserWithURLRequest) (User, error) {
	pathString, err := templates.GetUsersTemplate(usersTemplateData{UserID: url.PathEscape(userID)}, templates.SendbirdURLUserswithUserID)
	if err != nil {
		return User{}, err
	}

	parsedURL := c.PrepareUrl(pathString)

	result := User{}

	if err := c.putAndReturnJSON(parsedURL, r, &result); err != nil {
		return User{}, err
	}

	return result, nil
}

type UpdateAUserWithURLRequest struct {
	NickName                string `json:"nickname,omitempty"`
	ProfileURL              string `json:"profile_url,omitempty"`
	IssueAccessToken        bool   `json:"issue_access_token,omitempty"`
	IsActive                bool   `json:"is_active,omitempty"`
	LeaveAllWhenDeactivated bool   `json:"leave_all_when_deactivated,omitempty"`
}

func (c *Client) UpdateAUserWithFile(userID string, r *UpdateAUserWithFileRequest) (User, error) {
	return User{}, errors.New(SendbirdClientErrorNotImplemented)
}

type UpdateAUserWithFileRequest struct {
	NickName                string `json:"nickname,omitempty"`
	ProfileFile             string `json:"profile_file,omitempty"`
	IssueAccessToken        bool   `json:"issue_access_token,omitempty"`
	IsActive                bool   `json:"is_active,omitempty"`
	LeaveAllWhenDeactivated bool   `json:"leave_all_when_deactivated,omitempty"`
}

func (c *Client) ViewAUser(userID string) (User, error) {
	pathString, err := templates.GetUsersTemplate(usersTemplateData{UserID: url.PathEscape(userID)}, templates.SendbirdURLUserswithUserID)
	if err != nil {
		return User{}, err
	}

	parsedURL := c.PrepareUrl(pathString)
	result := User{}

	err = c.getAndReturnJSON(parsedURL, "", &result)
	if err != nil {
		return User{}, err
	}

	return result, nil
}

func (c *Client) DeleteAUser(userID string) (sendbirdErrorResponse, error) {
	pathString, err := templates.GetUsersTemplate(usersTemplateData{UserID: url.PathEscape(userID)}, templates.SendbirdURLUserswithUserID)
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

func (c *Client) GetUnreadMessageCount(userID string) (GetUnreadMessageCountResponse, error) {
	pathString, err := templates.GetUsersTemplate(usersTemplateData{UserID: url.PathEscape(userID)}, templates.SendbirdURLUsersUnreadCountWithUserID)
	if err != nil {
		return GetUnreadMessageCountResponse{}, err
	}

	parsedURL := c.PrepareUrl(pathString)
	result := GetUnreadMessageCountResponse{}

	err = c.getAndReturnJSON(parsedURL, "", &result)
	if err != nil {
		return GetUnreadMessageCountResponse{}, err
	}

	return result, nil
}

type GetUnreadMessageCountResponse struct {
	UnreadCount int64 `json:"unread_count"`
	//commonResponse
}

func (c *Client) BlockAUser(userID string, r *BlockAUserRequest) (User, error) {
	pathString, err := templates.GetUsersTemplate(usersTemplateData{UserID: url.PathEscape(userID)}, templates.SendbirdURLUsersBlockWithUserID)
	if err != nil {
		return User{}, err
	}

	parsedURL := c.PrepareUrl(pathString)
	result := User{}

	if err := c.postAndReturnJSON(parsedURL, r, &result); err != nil {
		return User{}, err
	}

	return result, nil
}

type BlockAUserRequest struct {
	TargetID string `json:"target_id"`
}

func (c *Client) ListBlockedUsers(userID string, r *ListBlockUsersRequest) (ListBlockUsersResponse, error) {
	pathString, err := templates.GetUsersTemplate(usersTemplateData{UserID: url.PathEscape(userID)}, templates.SendbirdURLUsersBlockWithUserID)
	if err != nil {
		return ListBlockUsersResponse{}, err
	}

	parsedURL := c.PrepareUrl(pathString)
	result := ListBlockUsersResponse{}

	raw := r.params().Encode()
	if err := c.getAndReturnJSON(parsedURL, raw, &result); err != nil {
		return ListBlockUsersResponse{}, err
	}

	return result, nil
}

func (r *ListBlockUsersRequest) params() url.Values {
	q := make(url.Values)

	if r.Token != "" {
		q.Set("token", r.Token)
	}

	if r.Limit >= ListLimitLowerBound && r.Limit <= ListLimitUpperBound {
		q.Set("limit", string(r.Limit))
	}

	return q
}

type ListBlockUsersRequest struct {
	Token string `json:"token,omitempty"`
	Limit int    `json:"limit,omitempty"`
}

type ListBlockUsersResponse struct {
	Users []User `json:"users"`
	Next  string `json:"next"`
}

func (c *Client) UnblockAUser(userID string, targetID string) (sendbirdErrorResponse, error) {
	pathString, err := templates.GetUsersTemplate(usersTemplateData{
		UserID:   url.PathEscape(userID),
		TargetID: url.PathEscape(targetID),
	}, templates.SendbirdURLUsersBlockWithUserIDandTargetID)
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

func (c *Client) ListBannedChannels(userID string, r *ListBannedChannelsRequest) (ListBannedChannelsResponse, error) {
	pathString, err := templates.GetUsersTemplate(usersTemplateData{UserID: url.PathEscape(userID)}, templates.SendbirdURLUsersBanWithUserID)
	if err != nil {
		return ListBannedChannelsResponse{}, err
	}

	parsedURL := c.PrepareUrl(pathString)
	result := ListBannedChannelsResponse{}

	raw := r.params().Encode()
	if err := c.getAndReturnJSON(parsedURL, raw, &result); err != nil {
		return ListBannedChannelsResponse{}, err
	}

	return result, nil
}

func (r *ListBannedChannelsRequest) params() url.Values {
	q := make(url.Values)

	if r.Token != "" {
		q.Set("token", r.Token)
	}

	if r.Limit >= ListLimitLowerBound && r.Limit <= ListLimitUpperBound {
		q.Set("limit", string(r.Limit))
	}

	return q
}

type ListBannedChannelsRequest struct {
	Token string `json:"token,omitempty"`
	Limit int    `json:"limit,omitempty"`
}

type ListBannedChannelsResponse struct {
	BannedChannels []BannedChannelResult `json:"banned_channels"`
	Next           string                `json:"next"`
}

type BannedChannelResult struct {
	StartAt     int64       `json:"start_at"`
	EndAt       int64       `json:"end_at"`
	Description string      `json:"description"`
	Channel     OpenChannel `json:"channel"`
}

func (c *Client) ListMutedChannels(userID string, r *ListMutedChannelsRequest) (ListMutedChannelsResponse, error) {
	pathString, err := templates.GetUsersTemplate(usersTemplateData{UserID: url.PathEscape(userID)}, templates.SendbirdURLUsersMuteWithUserID)
	if err != nil {
		return ListMutedChannelsResponse{}, err
	}

	parsedURL := c.PrepareUrl(pathString)
	result := ListMutedChannelsResponse{}

	raw := r.params().Encode()
	if err := c.getAndReturnJSON(parsedURL, raw, &result); err != nil {
		return ListMutedChannelsResponse{}, err
	}

	return result, nil
}

func (r *ListMutedChannelsRequest) params() url.Values {
	q := make(url.Values)

	if r.Token != "" {
		q.Set("token", r.Token)
	}

	if r.Limit >= ListLimitLowerBound && r.Limit <= ListLimitUpperBound {
		q.Set("limit", string(r.Limit))
	}

	return q
}

type ListMutedChannelsRequest struct {
	Token string `json:"token,omitempty"`
	Limit int    `json:"limit,omitempty"`
}

type ListMutedChannelsResponse struct {
	MutedChannels []OpenChannel `json:"muted_channels"`
	Next          string        `json:"next"`
	//commonResponse
}

func (c *Client) MarkAllMessagesAsRead(userID string) (sendbirdErrorResponse, error) {
	pathString, err := templates.GetUsersTemplate(usersTemplateData{UserID: url.PathEscape(userID)}, templates.SendbirdURLUsersMarkReadAllWithUserID)
	if err != nil {
		return sendbirdErrorResponse{}, err
	}

	parsedURL := c.PrepareUrl(pathString)
	result := sendbirdErrorResponse{}

	err = c.putAndReturnJSON(parsedURL, nil, &result)
	if err != nil {
		return sendbirdErrorResponse{}, err
	}

	return result, nil
}

func (c *Client) ListMyGroupChannels(userID string, r *ListMyGroupChannelsRequest) (ListMyGroupChannelsResponse, error) {
	pathString, err := templates.GetUsersTemplate(usersTemplateData{UserID: url.PathEscape(userID)}, templates.SendbirdURLUsersListGroupChannelsWithUserID)
	if err != nil {
		return ListMyGroupChannelsResponse{}, err
	}

	parsedURL := c.PrepareUrl(pathString)

	raw := r.params().Encode()
	//#HACK special case handling for sendbird API
	raw = strings.Replace(raw, "%2C", ",", -1)

	result := ListMyGroupChannelsResponse{}
	err = c.getAndReturnJSON(parsedURL, raw, &result)
	if err != nil {
		return ListMyGroupChannelsResponse{}, err
	}

	return result, nil
}

func (r *ListMyGroupChannelsRequest) params() url.Values {
	q := make(url.Values)

	if r.Token != "" {
		q.Set("token", r.Token)
	}

	if r.Limit >= ListLimitLowerBound && r.Limit <= ListLimitUpperBound {
		q.Set("limit", string(r.Limit))
	}

	if r.ShowEmpty {
		q.Set("show_empty", "true")
	}

	if r.ShowMember {
		q.Set("show_member", "true")
	}

	if r.ShowReadReceipt {
		q.Set("show_read_receipt", "true")
	}

	switch r.DistinctMode {
	case "distinct":
		q.Set("distinct_mode", "distinct")
	case "nondistinct":
		q.Set("distinct_mode", "nondistinct")
	case "all":
		q.Set("distinct_mode", "all")
	default:
		q.Set("distinct_mode", "all")
	}

	switch r.Order {
	case "latest_last_message":
		q.Set("order", "latest_last_message")
	case "chronological":
		q.Set("order", "chronological")
	default:
		q.Set("order", "chronological")
	}

	if r.MembersExactlyIn != nil && len(r.MembersExactlyIn) > 0 {
		q.Set("members_exactly_in", strings.Join(r.MembersExactlyIn, ","))
	}

	if r.MembersNicknameContains != nil && len(r.MembersIncludeIn) > 0 {
		q.Set("members_nickname_contains", strings.Join(r.MembersNicknameContains, ","))
	}

	if r.MembersIncludeIn != nil && len(r.MembersIncludeIn) > 0 {
		q.Set("members_include_in", strings.Join(r.MembersIncludeIn, ","))
	}

	if r.QueryType != "" {
		q.Set("query_type", r.QueryType)
	}

	if r.CustomType != "" {
		q.Set("custom_type", r.CustomType)
	}

	if r.ChannelURLs != nil && len(r.ChannelURLs) > 0 {
		q.Set("channel_urls", strings.Join(r.ChannelURLs, ","))
	}

	if r.CreatedAfter != 0 {
		q.Set("created_after", string(r.CreatedAfter))
	}

	if r.CreatedBefore != 0 {
		q.Set("created_before", string(r.CreatedBefore))
	}

	return q
}

type ListMyGroupChannelsRequest struct {
	Token                   string   `json:"token,omitempty"`
	Limit                   int      `json:"limit,omitempty"`
	ShowEmpty               bool     `json:"show_empty,omitempty"`
	ShowMember              bool     `json:"show_member,omitempty"`
	ShowReadReceipt         bool     `json:"show_read_receipt,omitempty"`
	DistinctMode            string   `json:"distinct_mode,omitempty"`
	Order                   string   `json:"order,omitempty"`
	MembersExactlyIn        []string `json:"members_exactly_in,omitempty"`
	MembersNicknameContains []string `json:"members_nickname_contains,omitempty"`
	MembersIncludeIn        []string `json:"members_include_in,omitempty"`
	QueryType               string   `json:"query_type,omitempty"`
	CustomType              string   `json:"custom_type,omitempty"`
	ChannelURLs             []string `json:"channel_urls,omitempty"`
	CreatedAfter            int64    `json:"created_after,omitempty"`
	CreatedBefore           int64    `json:"created_before,omitempty"`
}

type ListMyGroupChannelsResponse struct {
	Channels []GroupChannel `json:"channels"`
	Next     string         `json:"next"`
	//commonResponse
}

func (c *Client) RegisterADeviceToken(userID string, tokenType string, r *RegisterADeviceTokenRequest) (RegisterADeviceTokenResponse, error) {
	pathString, err := templates.GetUsersTemplate(usersTemplateData{
		UserID:    url.PathEscape(userID),
		TokenType: url.PathEscape(tokenType),
	}, templates.SendbirdURLUsersDeviceTokenWithUserIDandTokenType)
	if err != nil {
		return RegisterADeviceTokenResponse{}, err
	}

	parsedURL := c.PrepareUrl(pathString)
	result := RegisterADeviceTokenResponse{}

	if err := c.postAndReturnJSON(parsedURL, r, &result); err != nil {
		return RegisterADeviceTokenResponse{}, err
	}

	return result, nil
}

type RegisterADeviceTokenRequest struct {
	GcmRegToken     string `json:"gcm_reg_token,omitempty"`
	ApnsDeviceToken string `json:"apns_device_token,omitempty"`
}

type RegisterADeviceTokenResponse struct {
	Token []string `json:"token"`
	Type  string   `json:"type"`
	User  User     `json:"user"`
}

func (c *Client) UnregisterADeviceToken(userID string, tokenType string, pushToken string) (UnregisterADeviceTokenResponse, error) {
	pathString, err := templates.GetUsersTemplate(usersTemplateData{
		UserID:    url.PathEscape(userID),
		TokenType: url.PathEscape(tokenType),
		PushToken: url.PathEscape(pushToken),
	}, templates.SendbirdURLUsersDeviceTokenWithUserIDandTokenTypeandPushToken)
	if err != nil {
		return UnregisterADeviceTokenResponse{}, err
	}

	parsedURL := c.PrepareUrl(pathString)
	result := UnregisterADeviceTokenResponse{}

	err = c.deleteAndReturnJSON(parsedURL, "", &result)
	if err != nil {
		return UnregisterADeviceTokenResponse{}, err
	}

	return result, nil
}

type UnregisterADeviceTokenResponse struct {
	Token []string `json:"token"`
	User  User     `json:"user"`
}

func (c *Client) UnregisterAllDeviceTokens(userID string) (User, error) {
	pathString, err := templates.GetUsersTemplate(usersTemplateData{UserID: url.PathEscape(userID)}, templates.SendbirdURLUsersDeviceTokenWithUserID)
	if err != nil {
		return User{}, err
	}

	parsedURL := c.PrepareUrl(pathString)
	result := User{}

	err = c.deleteAndReturnJSON(parsedURL, "", &result)
	if err != nil {
		return User{}, err
	}

	return result, nil
}

func (c *Client) ListDeviceTokens(userID string, tokenType string) (ListDeviceTokensResponse, error) {
	pathString, err := templates.GetUsersTemplate(usersTemplateData{
		UserID:    url.PathEscape(userID),
		TokenType: url.PathEscape(tokenType),
	}, templates.SendbirdURLUsersDeviceTokenWithUserIDandTokenType)
	if err != nil {
		return ListDeviceTokensResponse{}, err
	}

	parsedURL := c.PrepareUrl(pathString)
	result := ListDeviceTokensResponse{}

	err = c.getAndReturnJSON(parsedURL, "", &result)
	if err != nil {
		return ListDeviceTokensResponse{}, err
	}

	return result, nil
}

type ListDeviceTokensResponse struct {
	Tokens []string `json:"tokens"`
	Type   string   `json:"type"`
	User   User     `json:"user"`
}

func (c *Client) UpdatePushPerferences(userID string, r *UpdatePushPerferencesRequest) (UpdatePushPerferencesResponse, error) {
	pathString, err := templates.GetUsersTemplate(usersTemplateData{UserID: url.PathEscape(userID)}, templates.SendbirdURLUsersPushPreferenceWithUserID)
	if err != nil {
		return UpdatePushPerferencesResponse{}, err
	}

	parsedURL := c.PrepareUrl(pathString)
	result := UpdatePushPerferencesResponse{}

	if err := c.putAndReturnJSON(parsedURL, r, &result); err != nil {
		return UpdatePushPerferencesResponse{}, err
	}

	return result, nil
}

type UpdatePushPerferencesRequest struct {
	DoNotDisturb bool   `json:"do_not_disturb,omitempty"`
	StartHour    int    `json:"start_hour,omitempty"`
	StartMin     int    `json:"start_min,omitempty"`
	EndHour      int    `json:"end_hour,omitempty"`
	EndMin       int    `json:"end_min,omitempty"`
	TimeZone     string `json:"timezone,omitempty"`
}

type UpdatePushPerferencesResponse struct {
	DoNotDisturb bool   `json:"do_not_disturb,omitempty"`
	StartHour    int    `json:"start_hour,omitempty"`
	StartMin     int    `json:"start_min,omitempty"`
	EndHour      int    `json:"end_hour,omitempty"`
	EndMin       int    `json:"end_min,omitempty"`
	TimeZone     string `json:"timezone,omitempty"`
	//commonResponse
}

func (c *Client) GetPushPerferences(userID string) (GetPushPerferencesResponse, error) {
	pathString, err := templates.GetUsersTemplate(usersTemplateData{UserID: url.PathEscape(userID)}, templates.SendbirdURLUsersPushPreferenceWithUserID)
	if err != nil {
		return GetPushPerferencesResponse{}, err
	}

	parsedURL := c.PrepareUrl(pathString)
	result := GetPushPerferencesResponse{}

	err = c.getAndReturnJSON(parsedURL, "", &result)
	if err != nil {
		return GetPushPerferencesResponse{}, err
	}

	return result, nil
}

type GetPushPerferencesResponse struct {
	DoNotDisturb bool   `json:"do_not_disturb,omitempty"`
	StartHour    int    `json:"start_hour,omitempty"`
	StartMin     int    `json:"start_min,omitempty"`
	EndHour      int    `json:"end_hour,omitempty"`
	EndMin       int    `json:"end_min,omitempty"`
	TimeZone     string `json:"timezone,omitempty"`
}

func (c *Client) ResetPushPerferences(userID string) (sendbirdErrorResponse, error) {
	pathString, err := templates.GetUsersTemplate(usersTemplateData{UserID: url.PathEscape(userID)}, templates.SendbirdURLUsersPushPreferenceWithUserID)
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

func (c *Client) UpdateChannelPushPerferences(userID string, channelURL string, r *UpdateChannelPushPerferencesRequest) (UpdateChannelPushPerferencesResponse, error) {
	pathString, err := templates.GetUsersTemplate(usersTemplateData{
		UserID:     url.PathEscape(userID),
		ChannelURL: url.PathEscape(channelURL),
	}, templates.SendbirdURLUsersPushPreferenceWithUserIDandChannelURL)
	if err != nil {
		return UpdateChannelPushPerferencesResponse{}, err
	}

	parsedURL := c.PrepareUrl(pathString)
	result := UpdateChannelPushPerferencesResponse{}

	if err := c.putAndReturnJSON(parsedURL, r, &result); err != nil {
		return UpdateChannelPushPerferencesResponse{}, err
	}

	return result, nil
}

type UpdateChannelPushPerferencesRequest struct {
	Enable bool `json:"enable"`
}

type UpdateChannelPushPerferencesResponse struct {
	Enable bool `json:"enable"`
}

func (c *Client) GetChannelPushPerferences(userID string, channelURL string) (GetChannelPushPerferencesResponse, error) {
	pathString, err := templates.GetUsersTemplate(usersTemplateData{
		UserID:     url.PathEscape(userID),
		ChannelURL: url.PathEscape(channelURL),
	}, templates.SendbirdURLUsersPushPreferenceWithUserIDandChannelURL)
	if err != nil {
		return GetChannelPushPerferencesResponse{}, err
	}

	parsedURL := c.PrepareUrl(pathString)
	result := GetChannelPushPerferencesResponse{}

	err = c.getAndReturnJSON(parsedURL, "", &result)
	if err != nil {
		return GetChannelPushPerferencesResponse{}, err
	}

	return result, nil
}

type GetChannelPushPerferencesResponse struct {
	Enable bool `json:"enable"`
}
