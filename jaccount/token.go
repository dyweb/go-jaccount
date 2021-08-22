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
	"fmt"
	"time"

	"golang.org/x/oauth2"
	"gopkg.in/square/go-jose.v2/jwt"
)

// Endpoint is jAccount's OAuth 2.0 endpoint.
//
// See https://developer.sjtu.edu.cn
var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://jaccount.sjtu.edu.cn/oauth2/authorize",
	TokenURL: "https://jaccount.sjtu.edu.cn/oauth2/token",
}

const issuer = "https://jaccount.sjtu.edu.cn/oauth2/"

type idToken struct {
	// The URL of the server which issued this token.
	Issuer string `json:"iss"`

	// The client ID, or set of client IDs, that this token is issued for. For
	// common uses, this is the client that initialized the auth flow.
	Audience string `json:"aud"`

	// A unique string which identifies the end user.
	Subject string `json:"sub"`

	// Expiry of the token.
	Expiry int64 `json:"exp"`

	// When the token was issued by the provider.
	IssuedAt int64 `json:"iat"`

	// Initial nonce provided during the authentication redirect.
	Nonce string `json:"nonce"`

	Name string `json:"name"`

	Code string `json:"code"`

	Type string `json:"type"`
}

func VerifyToken(rawToken string) (*IDToken, error) {
	token, err := jwt.ParseSigned(rawToken)
	if err != nil {
		return nil, err
	}

	var idToken idToken
	err = token.UnsafeClaimsWithoutVerification(&idToken)
	if err != nil {
		return nil, err
	}

	t := &IDToken{
		Issuer:   idToken.Issuer,
		Audience: idToken.Audience,
		Subject:  idToken.Subject,
		Expiry:   time.Unix(idToken.Expiry, 0),
		IssuedAt: time.Unix(idToken.IssuedAt, 0),
		Nonce:    idToken.Nonce,
		Name:     idToken.Name,
		Code:     idToken.Code,
		Type:     Type(idToken.Type),
	}

	if t.Issuer != issuer {
		return nil, fmt.Errorf("ID token issued by a different provider, expected %q got %q", issuer, t.Issuer)
	}

	now := time.Now()
	if t.Expiry.Before(now) {
		return nil, fmt.Errorf("token is expired")
	}

	return t, nil
}

type IDToken struct {
	Issuer   string
	Audience string
	Subject  string
	Expiry   time.Time
	IssuedAt time.Time
	Nonce    string
	Name     string
	Code     string
	Type     Type
}

type Type string

const (
	// FACULTY represents "教职工".
	FACULTY Type = "faculty"
	// STUDENT represents "学生".
	STUDENT Type = "student"
	// MEDICAL_SCHOOL_FACULTY represents "医学院教职工".
	MEDICAL_SCHOOL_FACULTY Type = "yxy"
	// AFFILIATED_UNIT_FACULTY represents "附属单位职工".
	AFFILIATED_UNIT_FACULTY Type = "fsyyjzg"
	// VIP represents "贵宾".
	VIP Type = "vip"
	// POSTPHD represents "博士后".
	POSTPHD Type = "postphd"
	// EXTERNAL_TEACHER represents "外聘教师".
	EXTERNAL_TEACHER Type = "external_teacher"
	// SUMMER represents "暑期生".
	SUMMER Type = "summer"
	// TEAM represents "集体账号".
	TEAM Type = "team"
	// ALUMNI represents "校友".
	ALUMNI Type = "alumni"
	// GREEN represents "绿色通道".
	GREEN Type = "green"
	// OUTSIDE represents "合作交流".
	OUTSIDE Type = "outside"
)
