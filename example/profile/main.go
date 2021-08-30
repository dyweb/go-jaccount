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
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dyweb/go-jaccount/example/util"
	"github.com/dyweb/go-jaccount/jaccount"
	"golang.org/x/oauth2"
)

var (
	ClientID     = os.Getenv("CLIENT_ID")
	ClientSecret = os.Getenv("CLIENT_SECRET")
)

func main() {
	var config = &oauth2.Config{
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		Endpoint:     jaccount.Endpoint,
		RedirectURL:  "http://localhost:8000/callback",
		Scopes:       []string{jaccount.ScopeEssential},
	}

	var client *jaccount.Client

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		state := base64.RawStdEncoding.EncodeToString(util.RandBytes(16))

		util.SetCookie(w, r, "state", state)

		url := config.AuthCodeURL(state)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	})

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		state, err := r.Cookie("state")
		if err != nil {
			http.Error(w, "state not found", http.StatusBadRequest)
			return
		}

		util.DeleteCookie(w, r, "state")
		if r.URL.Query().Get("state") != state.Value {
			http.Error(w, "state mismatch", http.StatusBadRequest)
			return
		}

		code := r.URL.Query().Get("code")
		oauth2Token, err := config.Exchange(r.Context(), code)
		if err != nil {
			http.Error(w, "failed to exchange token", http.StatusInternalServerError)
			return
		}

		c := config.Client(context.Background(), oauth2Token)

		client = jaccount.NewClient(c)
		http.Redirect(w, r, "/profile", http.StatusTemporaryRedirect)
	})

	http.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		profile, err := client.Profile.Get(context.Background())
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to fetch profile: %s", err), http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(profile)
		if err != nil {
			http.Error(w, "failed to marshal profile", http.StatusInternalServerError)
			return
		}
		w.Write(data)
	})

	http.HandleFunc("/card/info", func(w http.ResponseWriter, r *http.Request) {
		cardInfo, err := client.Card.GetCardInfo(context.Background())
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to fetch card information: %s", err), http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(cardInfo)
		if err != nil {
			http.Error(w, "failed to marshal card information", http.StatusInternalServerError)
			return
		}
		w.Write(data)
	})

	http.HandleFunc("/card/transactions", func(w http.ResponseWriter, r *http.Request) {
		cardInfo, err := client.Card.ListTransactions(context.Background(), &jaccount.CardListTransactionsOptions{BeginDate: 1619881405000})
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to fetch card transactions: %s", err), http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(cardInfo)
		if err != nil {
			http.Error(w, "failed to marshal card transactions", http.StatusInternalServerError)
			return
		}
		w.Write(data)
	})

	log.Println("listening on :8000")
	http.ListenAndServe(":8000", nil)
}
