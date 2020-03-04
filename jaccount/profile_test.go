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
