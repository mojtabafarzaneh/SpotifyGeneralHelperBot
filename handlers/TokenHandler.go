package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"time"
)

type SpotifyTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func GetSpotifyToken(db *sql.DB, clientID, clientSecret, grant_type, redirectURI string, userID uint64) error {
	spotify_url := "https://accounts.spotify.com/api/token"
	data := url.Values{}

	data.Set("grant_type", grant_type)
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	req, err := http.NewRequest("POST", spotify_url, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("spotify token request failed: %s", string(body))
	}

	var tokenResp SpotifyTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return err
	}

	expiresAt := time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second)

	query := `
		INSERT INTO spotify_tokens (user_id, access_token, token_type, expires_at)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (user_id) DO UPDATE
		SET access_token = EXCLUDED.access_token,
		    token_type = EXCLUDED.token_type,
		    expires_at = EXCLUDED.expires_at;
	`

	_, err = db.Exec(query, userID, tokenResp.AccessToken, tokenResp.TokenType, expiresAt)
	return err
}

func CallbackHandler(db *sql.DB, clientID, grant_type, clientSecret, redirectURI string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userID := uint64(1)

		err := GetSpotifyToken(db, clientID, clientSecret, grant_type, redirectURI, userID)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to get token: %v", err), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Spotify token saved successfully!"))
	}
}
