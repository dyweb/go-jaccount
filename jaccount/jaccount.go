package jaccount

import (
	"encoding/json"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
)

// Endpoint is jAccount's OAuth 2.0 endpoint.
//
// See http://developer.sjtu.edu.cn
var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://jaccount.sjtu.edu.cn/oauth2/authorize",
	TokenURL: "https://jaccount.sjtu.edu.cn/oauth2/token",
}

const (
	defaultBaseURL = "https://api.sjtu.edu.cn/v1/"
)

// Client manages communication with the jAccount API.
type Client struct {
	client *http.Client

	BaseURL *url.URL
}

// NewClient returns a new jAccount API client.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{client: httpClient, BaseURL: baseURL}

	return c
}

// NewRequest creates an API request.
func (c *Client) NewRequest(method string, path string) (*http.Request, error) {
	url, err := c.BaseURL.Parse(path)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// Response is a jAccount API response.
type Response struct {
	ErrNO    *int            `json:"errno,omitempty"`
	Error    *string         `json:"error,omitempty"`
	Total    *int            `json:"total,omitempty"`
	Entities json.RawMessage `json:"entities,omitempty"`
}

// Do sends an API request and returns the API response.
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ErrorResponse reports one or more errors caused by an API request.
type ErrorResponse struct {
	ErrNO *int    `json:"errno,omitempty"`
	Error *string `json:"error,omitempty"`
	Total *int    `json:"total,omitempty"`
}
