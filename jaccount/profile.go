package jaccount

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ProfileService handles communications with the profile data related methods of the jAccount API.
//
// See http://developer.sjtu.edu.cn/wiki/APIs#Profile for more information.
type ProfileService service

// Profile represents the profile of the user.
type Profile struct {
	ID         *string     `json:"id,omitempty"`
	Account    *string     `json:"account,omitempty"`
	Name       *string     `json:"name,omitempty"`
	Kind       *string     `json:"kind,omitempty"`
	Code       *string     `json:"code,omitempty"`
	UserType   *string     `json:"userType,omitempty"`
	Organize   *Organize   `json:"organize,omitempty"`
	ClassNO    *string     `json:"classNo,omitempty"`
	Birthday   *Birthday   `json:"birthday,omitempty"`
	Gender     *string      `json:"gender,omitempty"`
	Email      *string      `json:"email,omitempty"`
	TimeZone   *int        `json:"timeZone,omitempty"`
	Identities []*Identity `json:"identities,omitempty"`
	CardNO     *string     `json:"cardNo,omitempty"`
	CardType   *string     `json:"cardType,omitempty"`
	UnionID    *string     `json:"unionId,omitempty"`
}

// Birthday represents the birthday of the user.
type Birthday struct {
	BirthYear  *string `json:"birthYear,omitempty"`
	BirthMonth *string `json:"birthMonth,omitempty"`
	BirthDay   *string `json:"birthDay,omitempty"`
}

// Organize represents the organization of the user.
type Organize struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// Identity represents the identification of the user.
type Identity struct {
	Kind          *string   `json:"kind,omitempty"`
	IsDefault     *bool     `json:"isDefault,omitempty"`
	Code          *string   `json:"code,omitempty"`
	UserType      *string   `json:"userType,omitempty"`
	Organize      *Organize `json:"organize,omitempty"`
	MgtOrganize   *Organize `json:"mgtOrganize,omitempty"`
	Status        *string   `json:"status,omitempty"`
	ExpireDate    *string   `json:"expireDate,omitempty"`
	CreateDate    *int64    `json:"createDate,omitempty"`
	UpdateDate    *int64    `json:"updateDate,omitempty"`
	ClassNO       *string   `json:"classNo,omitempty"`
	Gjm           *string   `json:"gjm,omitempty"`
	Major         *Major    `json:"major,omitempty"`
	AdmissionDate *string   `json:"admissionDate,omitempty"`
	TrainLevel    *string   `json:"trainLevel,omitempty"`
	GraduateDate  *string   `json:"graduateDate,omitempty"`
}

// Major represents the major of the user.
type Major struct {
	Name *string `json:"name,omitempty"`
	ID   *string `json:"id,omitempty"`
}

// Get gets the profile of the user.
func (s *ProfileService) Get(ctx context.Context) (*Profile, error) {
	req, err := s.client.NewRequest(http.MethodGet, "me/profile")
	resp, err := s.client.Do(req)

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
