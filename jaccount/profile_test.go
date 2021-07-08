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
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func TestProfileService_Get(t *testing.T) {
	profile := &Profile{
		ID:       String("00000000-0000-0000-0000-000000000000"),
		Account:  String("test"),
		Name:     String("test"),
		Kind:     String("canvas.profile"),
		Code:     String("000000000000"),
		UserType: String("student"),
		Organize: &Organize{
			Name: String("软件学院"),
			ID:   String("03700"),
		},
		ClassNO:  String("B0000000"),
		TimeZone: Int(0),
		UnionID:  String("union_id"),
		Birthday: &Birthday{
			BirthYear:  String("1970"),
			BirthMonth: String("01"),
			BirthDay:   String("01"),
		},
		Gender:   String("male"),
		Email:    String("example@example.com"),
		CardNO:   String("31010119700101000X"),
		CardType: String("01"),
	}
	profiles := [1]*Profile{profile}
	rawProfiles, err := json.Marshal(profiles)
	if err != nil {
		t.Errorf("error = %v", err)
	}

	response := &Response{
		ErrNO:    Int(0),
		Error:    String("success"),
		Total:    Int(0),
		Entities: rawProfiles,
	}
	rawResponse, err := json.Marshal(response)
	if err != nil {
		t.Errorf("error = %v", err)
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(rawResponse)
	}))
	defer ts.Close()

	testURL, err := url.Parse(ts.URL)
	if err != nil {
		t.Errorf("error = %v", err)
	}

	client := NewClient(nil)
	client.BaseURL = testURL

	tests := []struct {
		name    string
		want    *Profile
		wantErr bool
	}{
		{"success", profile, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.Profile.Get(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("ProfileService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProfileService.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
