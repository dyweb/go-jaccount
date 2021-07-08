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

package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/dyweb/go-jaccount/jaccount"
	"golang.org/x/oauth2"
)

func main() {
	var config = &oauth2.Config{
		ClientID:     os.Getenv("clientid"),
		ClientSecret: os.Getenv("secretkey"),
		Endpoint:     jaccount.Endpoint,
		RedirectURL:  "http://localhost:8000/callback",
		Scopes:       []string{"essential"},
	}

	var client *jaccount.Client

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		url := config.AuthCodeURL("state")
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	})

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query()["code"][0]
		token, _ := config.Exchange(r.Context(), code)
		c := config.Client(oauth2.NoContext, token)

		client = jaccount.NewClient(c)
		http.Redirect(w, r, "/profile", http.StatusTemporaryRedirect)
	})

	http.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		profile, _ := client.Profile.Get(context.Background())
		data, _ := json.Marshal(profile)
		w.Write(data)
	})

	http.ListenAndServe(":8000", nil)
}
