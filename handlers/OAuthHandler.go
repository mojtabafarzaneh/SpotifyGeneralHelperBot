package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

var (
	clientID     = "712630d2b9424750a0ea6c9af45a2165"
	clientSecret = "2166d6c5eec547ddb79b069e2869ffd9"
	redirectURL  = "http://localhost:8000/callback"
	demoToken    *oauth2.Token
)

var spotifyOAuthConfig = &oauth2.Config{
	ClientID:     clientID,
	ClientSecret: clientSecret,
	Scopes:       []string{"user-library-read"},
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://accounts.spotify.com/authorize",
		TokenURL: "https://accounts.spotify.com/api/token",
	},
	RedirectURL: redirectURL,
}

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	if errParam := r.FormValue("error"); errParam != "" {
		http.Error(w, "authorization error: "+errParam, http.StatusBadRequest)
		return
	}

	code := r.FormValue("code")
	if code == "" {
		http.Error(w, "missing code in callback", http.StatusBadRequest)
		return
	}

	token, err := spotifyOAuthConfig.Exchange(ctx, code)
	if err != nil {
		http.Error(w, "token exchange failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	demoToken = token

	fmt.Fprintf(w, "Got token (expires in %d). Now visit /albums to fetch saved albums.\n", token.Extra("expires_in"))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	state := fmt.Sprintf("st%d", time.Now().UnixNano())
	url := spotifyOAuthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusFound)
}
