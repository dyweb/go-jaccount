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
	"net/http"
)

// EnterpriseService handles communications with the enterprise data related methods of the jAccount API.
//
// See https://developer.sjtu.edu.cn/api/enterprise.html for more information.
type EnterpriseService service

func (s *EnterpriseService) GetUserPositions(ctx context.Context) (*Positions, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/v1/enterprise/user/positions")
	if err != nil {
		return nil, err
	}

	positions := make([]Positions, 1)
	_, err = s.client.Do(ctx, req, &positions)
	if err != nil {
		return nil, err
	}

	return &positions[0], nil
}

type Positions struct {
	Account   string     `json:"account,omitempty"`
	Name      string     `json:"name,omitempty"`
	Positions []Position `json:"positions,omitempty"`
}

type Position struct {
	Post Post `json:"post,omitempty"`
	Dept Dept `json:"dept,omitempty"`
}

type Post struct {
	PostCode string `json:"postCode,omitempty"`
	PostName string `json:"postName,omitempty"`
	Formal   bool   `json:"formal,omitempty"`
}

type Dept struct {
	OrganizeID       string `json:"organizeId,omitempty"`
	OrganizeName     string `json:"organizeName,omitempty"`
	ParentOrganizeID string `json:"parentOrganizeId,omitempty"`
	Independent      bool   `json:"independent,omitempty"`
}
