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
		oauthClient := config.Client(oauth2.NoContext, token)

		client = jaccount.NewClient(oauthClient)
		http.Redirect(w, r, "/profile", http.StatusTemporaryRedirect)
	})

	http.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		profile, _ := client.Profile.Get(context.Background())
		data, _ := json.Marshal(profile)
		w.Write(data)
	})

	http.ListenAndServe(":8000", nil)
}
