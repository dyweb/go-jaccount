package jaccount

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Profile represents the profile of the user.
//
// See http://developer.sjtu.edu.cn/wiki/APIs#Profile for more information.
type Profile struct {
	ID       *string   `json:"id,omitempty"`
	Account  *string   `json:"account,omitempty"`
	Name     *string   `json:"name,omitempty"`
	Kind     *string   `json:"kind,omitempty"`
	Code     *string   `json:"code,omitempty"`
	UserType *string   `json:"userType,omitempty"`
	Organize *Organize `json:"organize,omitempty"`
	ClassNO  *string   `json:"classNo,omitempty"`
	TimeZone *int      `json:"timeZone,omitempty"`
	UnionID  *string   `json:"unionId,omitempty"`
}

// Organize represents the organization of the user.
type Organize struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// GetProfile gets the profile of the user.
func (c *Client) GetProfile(ctx context.Context) (*Profile, error) {
	req, err := c.NewRequest(http.MethodGet, "me/profile")
	resp, err := c.Do(req)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := new(Response)
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	profile := make([]Profile, 1)
	err = json.Unmarshal(response.Entities, &profile)
	if err != nil {
		return nil, err
	}

	return &profile[0], nil
}
