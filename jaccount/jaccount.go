/*
Copyright 2021 The Go jAccount Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package jaccount

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://api.sjtu.edu.cn"
)

// Client manages communication with the jAccount API.
type Client struct {
	client *http.Client

	BaseURL *url.URL

	common service

	Profile    *ProfileService
	Enterprise *EnterpriseService
}

type service struct {
	client *Client
}

// NewClient returns a new jAccount API client.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL}
	c.common.client = c

	c.Profile = (*ProfileService)(&c.common)
	c.Enterprise = (*EnterpriseService)(&c.common)

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
	ErrNO    int             `json:"errno,omitempty"`
	Error    string          `json:"error,omitempty"`
	Total    int             `json:"total,omitempty"`
	Entities json.RawMessage `json:"entities,omitempty"`
}

// Do sends an API request and returns the API response.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	req = req.WithContext(ctx)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := new(Response)
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(response.Entities, v)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
