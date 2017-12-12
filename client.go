package sendbirdclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

const (
	SendbirdAPITokenHeaderKey      = "Api-Token"
	SendbirdContentTypeHeaderValue = "application/json, charset=utf8"
)

// Client may be used to make requests to Sendbird Platform APIs
type Client struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
}

// ClientOption is the type of constructor options for NewClient(...).
type ClientOption func(*Client) error

// NewClient is the factory function for Client struct
func NewClient(options ...ClientOption) (*Client, error) {
	c := &Client{}
	WithHTTPClient(&http.Client{})(c)
	for _, option := range options {
		err := option(c)
		if err != nil {
			return nil, err
		}
	}

	if c.apiKey == "" {
		return nil, errors.New("Missing API Key")
	}
	return c, nil
}

// WithHTTPClient configures a Maps API client with a http.Client to make requests over.
func WithHTTPClient(c *http.Client) ClientOption {
	return func(client *Client) error {
		client.httpClient = c
		return nil
	}
}

// WithAPIKey configures a Maps API client with an API Key
func WithAPIKey(apiKey string) ClientOption {
	return func(c *Client) error {
		c.apiKey = apiKey
		return nil
	}
}

type apiRequest interface {
	params() url.Values
}

func (c *Client) prepareHeader(req *http.Request) {
	req.Header.Set("Content-Type", SendbirdContentTypeHeaderValue)
	req.Header.Set(SendbirdAPITokenHeaderKey, c.apiKey)
}

func (c *Client) get(config *url.URL, rawQueryString string) (*http.Response, error) {

	req, err := http.NewRequest("GET", config.String(), nil)
	if err != nil {
		return nil, err
	}

	c.prepareHeader(req)
	req.URL.RawQuery = rawQueryString

	fmt.Println("URL:", req.URL.String())

	return c.httpClient.Do(req)
}

func (c *Client) post(config *url.URL, apiReq interface{}) (*http.Response, error) {

	body, err := json.Marshal(apiReq)
	if err != nil {
		return nil, err
	}

	fmt.Println("post(): requestBody=", string(body))

	req, err := http.NewRequest("POST", config.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	c.prepareHeader(req)

	req.URL.RawQuery = url.Values{}.Encode()
	return c.httpClient.Do(req)
}

func (c *Client) delete(config *url.URL, rawQueryString string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", config.String(), nil)
	if err != nil {
		return nil, err
	}

	c.prepareHeader(req)
	req.URL.RawQuery = rawQueryString
	return c.httpClient.Do(req)
}

func (c *Client) put(config *url.URL, apiReq interface{}) (*http.Response, error) {

	body, err := json.Marshal(apiReq)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", config.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	c.prepareHeader(req)

	req.URL.RawQuery = url.Values{}.Encode()
	return c.httpClient.Do(req)
}

func (c *Client) getAndReturnJSON(config *url.URL, rawQuery string, resp interface{}) error {
	httpResp, err := c.get(config, rawQuery)
	if err != nil {
		return err
	}

	defer httpResp.Body.Close()

	err = CheckSendbirdError(httpResp)
	if err != nil {
		return err
	}

	return json.NewDecoder(httpResp.Body).Decode(resp)
}

func (c *Client) postAndReturnJSON(config *url.URL, apiReq interface{}, resp interface{}) error {
	httpResp, err := c.post(config, apiReq)
	if err != nil {
		return err
	}

	defer httpResp.Body.Close()

	err = CheckSendbirdError(httpResp)
	if err != nil {
		return err
	}

	return json.NewDecoder(httpResp.Body).Decode(resp)
}

func (c *Client) deleteAndReturnJSON(config *url.URL, rawQueryString string, resp interface{}) error {
	httpResp, err := c.delete(config, rawQueryString)
	if err != nil {
		return err
	}

	defer httpResp.Body.Close()

	err = CheckSendbirdError(httpResp)
	if err != nil {
		return err
	}

	return json.NewDecoder(httpResp.Body).Decode(resp)
}

func (c *Client) putAndReturnJSON(config *url.URL, apiReq interface{}, resp interface{}) error {
	httpResp, err := c.put(config, apiReq)
	if err != nil {
		return err
	}

	defer httpResp.Body.Close()

	err = CheckSendbirdError(httpResp)
	if err != nil {
		return err
	}

	return json.NewDecoder(httpResp.Body).Decode(resp)
}

func (c *Client) PrepareUrl(pathEncodedUrl string) *url.URL {
	urlVal := &url.URL{
		Scheme:  constScheme,
		Host:    constHost,
		Path:    constVersion + pathEncodedUrl,
		RawPath: constVersion + pathEncodedUrl,
	}

	//fmt.Println("Prepare URL:", urlVal.String())

	return urlVal
}
