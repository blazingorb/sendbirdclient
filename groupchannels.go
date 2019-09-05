package sendbirdclient

import (
	"errors"
	"net/url"
	"strings"

	"github.com/blazingorb/sendbirdclient/templates"
)

type GroupChannel struct {
	channel

	IsDistinct         bool             `json:"is_distinct"`
	MemberCount        int              `json:"member_count"`
	Members            []User           `json:"members"`
	ReadReceipt        map[string]int64 `json:"read_receipt"`
	UnreadMessageCount int              `json:"unread_message_count"`
	LastMessage        LastMessage      `json:"last_message"`
}

type LastMessage struct {
	CreatedAt int64 `json:"created_at"`
	User      User  `json:"user"`
}

type groupChannelsTemplateData struct {
	ChannelURL string
	UserID     string
}

func (c *Client) CreateAGroupChannelWithURL(r *CreateAGroupChannelWithURLRequest) (GroupChannel, error) {
	result := GroupChannel{}

	if err := c.postAndReturnJSON(c.PrepareUrl(SendbirdURLGroupChannels), r, &result); err != nil {
		return GroupChannel{}, err
	}

	return result, nil
}

type CreateAGroupChannelWithURLRequest struct {
	Name       string   `json:"name,omitempty"`
	CoverURL   string   `json:"cover_url,omitempty"`
	CustomType string   `json:"custom_type,omitempty"`
	Data       string   `json:"data,omitempty"`
	UserIDs    []string `json:"user_ids,omitempty"`
	IsDistinct bool     `json:"is_distinct,omitempty"`
}

func (c *Client) CreateAGroupChannelWithFile(r *CreateAGroupChannelWithFileRequest) (GroupChannel, error) {
	return GroupChannel{}, errors.New(SendbirdClientErrorNotImplemented)
}

type CreateAGroupChannelWithFileRequest struct {
	Name       string   `json:"name,omitempty"`
	CoverFile  []byte   `json:"cover_file,omitempty"`
	CustomType string   `json:"custom_type,omitempty"`
	Data       string   `json:"data,omitempty"`
	UserIDs    []string `json:"user_ids,omitempty"`
	IsDistinct bool     `json:"is_distinct,omitempty"`
}

func (c *Client) ListGroupChannels(r *ListGroupChannelsRequest) (ListGroupChannelsResponse, error) {
	raw := r.params().Encode()
	//#HACK special case handling for sendbird API
	raw = strings.Replace(raw, "%2C", ",", -1)

	result := ListGroupChannelsResponse{}
	err := c.getAndReturnJSON(c.PrepareUrl(SendbirdURLGroupChannels), raw, &result)
	if err != nil {
		return ListGroupChannelsResponse{}, err
	}

	return result, nil
}

func (r *ListGroupChannelsRequest) params() url.Values {
	q := make(url.Values)

	if r.Token != "" {
		q.Set("token", r.Token)
	}

	if r.Limit >= ListLimitLowerBound && r.Limit <= ListLimitUpperBound {
		q.Set("limit", string(r.Limit))
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

	if r.CreatedAfter > 0 {
		q.Set("created_after", string(r.CreatedAfter))
	}

	if r.CreatedBefore > 0 {
		q.Set("created_before", string(r.CreatedBefore))
	}

	return q
}

type ListGroupChannelsRequest struct {
	Token                   string   `json:"token,omitempty"`
	Limit                   int      `json:"limit,omitempty"`
	ShowMember              bool     `json:"show_member,omitempty"`
	ShowReadReceipt         bool     `json:"show_read_receipt,omitempty"`
	DistinctMode            string   `json:"distinct_mode,omitempty"`
	MembersExactlyIn        []string `json:"members_exactly_in,omitempty"`
	MembersIncludeIn        []string `json:"members_include_in,omitempty"`
	MembersNicknameContains []string `json:"members_nickname_contains,omitempty"`
	QueryType               string   `json:"query_type,omitempty"`
	CustomType              string   `json:"custom_type,omitempty"`
	ChannelURLs             []string `json:"channel_urls,omitempty"`
	CreatedAfter            int64    `json:"created_after,omitempty"`
	CreatedBefore           int64    `json:"created_before,omitempty"`
}

type ListGroupChannelsResponse struct {
	Channels []GroupChannel `json:"channels"`
	Next     string         `json:"next"`
}

func (c *Client) UpdateAGroupChannel(channelURL string, r *UpdateAGroupChannelRequest) (GroupChannel, error) {
	pathString, err := templates.GetGroupChannelTemplate(groupChannelsTemplateData{ChannelURL: url.PathEscape(channelURL)}, templates.SendbirdURLGroupChannelsWithChannelURL)

	if err != nil {
		return GroupChannel{}, err
	}

	parsedURL := c.PrepareUrl(pathString)

	result := GroupChannel{}

	if err := c.putAndReturnJSON(parsedURL, r, &result); err != nil {
		return GroupChannel{}, err
	}

	return result, nil
}

type UpdateAGroupChannelRequest struct {
	Name       string `json:"name,omitempty"`
	CoverURL   string `json:"cover_url,omitempty"`
	CustomType string `json:"custom_type,omitempty"`
	Data       string `json:"data,omitempty"`
	IsDistinct bool   `json:"is_distinct,omitempty"`
}

func (c *Client) ViewAGroupChannel(channelURL string, r *ViewAGroupChannelRequest) (GroupChannel, error) {
	pathString, err := templates.GetGroupChannelTemplate(groupChannelsTemplateData{ChannelURL: url.PathEscape(channelURL)}, templates.SendbirdURLGroupChannelsWithChannelURL)

	if err != nil {
		return GroupChannel{}, err
	}

	parsedURL := c.PrepareUrl(pathString)

	result := GroupChannel{}

	raw := r.params().Encode()
	if err := c.getAndReturnJSON(parsedURL, raw, &result); err != nil {
		return GroupChannel{}, err
	}

	return result, nil
}

func (r *ViewAGroupChannelRequest) params() url.Values {
	q := make(url.Values)

	if r.ShowReadReceipt {
		q.Set("show_read_receipt", "true")
	}

	if r.ShowMember {
		q.Set("show_member", "true")
	}

	return q
}

type ViewAGroupChannelRequest struct {
	ShowReadReceipt bool `json:"show_read_receipt,omitempty"`
	ShowMember      bool `json:"show_member,omitempty"`
}

func (c *Client) DeleteAGroupChannel(channelURL string) (sendbirdErrorResponse, error) {
	pathString, err := templates.GetGroupChannelTemplate(groupChannelsTemplateData{ChannelURL: url.PathEscape(channelURL)}, templates.SendbirdURLGroupChannelsWithChannelURL)

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

func (c *Client) ListMembersInGroupChannel(channelURL string, r *ListMembersInGroupChannelRequest) (ListMembersInGroupChannelResponse, error) {
	pathString, err := templates.GetGroupChannelTemplate(groupChannelsTemplateData{ChannelURL: url.PathEscape(channelURL)}, templates.SendbirdURLGroupChannelsMembersWithChannelURL)

	if err != nil {
		return ListMembersInGroupChannelResponse{}, err
	}

	parsedURL := c.PrepareUrl(pathString)

	result := ListMembersInGroupChannelResponse{}

	raw := r.params().Encode()
	if err := c.getAndReturnJSON(parsedURL, raw, &result); err != nil {
		return ListMembersInGroupChannelResponse{}, err
	}

	return result, nil
}

func (r *ListMembersInGroupChannelRequest) params() url.Values {
	q := make(url.Values)

	if r.Token != "" {
		q.Set("token", r.Token)
	}

	if r.Limit >= 1 && r.Limit <= 100 {
		q.Set("limit", string(r.Limit))
	}

	return q
}

type ListMembersInGroupChannelRequest struct {
	Token string `json:"token,omitempty"`
	Limit int    `json:"limit,omitempty"`
}

type ListMembersInGroupChannelResponse struct {
	Members []User `json:"members"`
	Next    string `json:"next"`
}

func (c *Client) CheckIfMemberInGroupChannel(channelURL string, userID string) (CheckIfMemberInGroupChannelResponse, error) {
	pathString, err := templates.GetGroupChannelTemplate(groupChannelsTemplateData{
		ChannelURL: url.PathEscape(channelURL),
		UserID:     url.PathEscape(userID),
	}, templates.SendbirdURLGroupChannelsMembersWithChannelURLAndUserID)

	if err != nil {
		return CheckIfMemberInGroupChannelResponse{}, err
	}

	parsedURL := c.PrepareUrl(pathString)

	result := CheckIfMemberInGroupChannelResponse{}

	if err := c.getAndReturnJSON(parsedURL, "", &result); err != nil {
		return CheckIfMemberInGroupChannelResponse{}, err
	}

	return result, nil
}

type CheckIfMemberInGroupChannelResponse struct {
	IsMember bool `json:"is_member"`
}

func (c *Client) InviteMembersToGroupChannel(channelURL string, r *InviteMembersToGroupChannelRequest) (GroupChannel, error) {
	pathString, err := templates.GetGroupChannelTemplate(groupChannelsTemplateData{ChannelURL: url.PathEscape(channelURL)}, templates.SendbirdURLGroupChannelsInviteWithChannelURL)

	if err != nil {
		return GroupChannel{}, err
	}

	parsedURL := c.PrepareUrl(pathString)

	result := GroupChannel{}

	if err := c.postAndReturnJSON(parsedURL, r, &result); err != nil {
		return GroupChannel{}, err
	}

	return result, nil
}

type InviteMembersToGroupChannelRequest struct {
	UserIDs []string `json:"user_ids"`
}

func (c *Client) HideFromAGroupChannel(channelURL string, r *HideFromAGroupChannelRequest) (sendbirdErrorResponse, error) {
	pathString, err := templates.GetGroupChannelTemplate(groupChannelsTemplateData{ChannelURL: url.PathEscape(channelURL)}, templates.SendbirdURLGroupChannelsHideWithChannelURL)

	if err != nil {
		return sendbirdErrorResponse{}, err
	}

	parsedURL := c.PrepareUrl(pathString)

	result := sendbirdErrorResponse{}

	if err := c.putAndReturnJSON(parsedURL, r, &result); err != nil {
		return sendbirdErrorResponse{}, err
	}

	return result, nil
}

type HideFromAGroupChannelRequest struct {
	UserID string `json:"user_id"`
}

func (c *Client) LeaveFromAGroupChannel(channelURL string, r *LeaveFromAGroupChannelRequest) (sendbirdErrorResponse, error) {
	pathString, err := templates.GetGroupChannelTemplate(groupChannelsTemplateData{ChannelURL: url.PathEscape(channelURL)}, templates.SendbirdURLGroupChannelsLeaveWithChannelURL)

	if err != nil {
		return sendbirdErrorResponse{}, err
	}

	parsedURL := c.PrepareUrl(pathString)

	result := sendbirdErrorResponse{}

	if err := c.putAndReturnJSON(parsedURL, r, &result); err != nil {
		return sendbirdErrorResponse{}, err
	}

	return result, nil
}

type LeaveFromAGroupChannelRequest struct {
	UserIDs []string `json:"user_ids"`
}

func (c *Client) FreezeAnGroupChannel(channelURL string, r *FreezeAnGroupChannelRequest) (GroupChannel, error) {
	pathString, err := templates.GetGroupChannelTemplate(groupChannelsTemplateData{ChannelURL: url.PathEscape(channelURL)}, templates.SendbirdURLGroupChannelsFreezeWithChannelURL)
	if err != nil {
		return GroupChannel{}, err
	}

	parsedURL := c.PrepareUrl(pathString)

	result := GroupChannel{}

	if err := c.putAndReturnJSON(parsedURL, r, &result); err != nil {
		return GroupChannel{}, err
	}

	return result, nil
}

type FreezeAnGroupChannelRequest struct {
	Freeze bool `json:"freeze"`
}
