package handlers

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"math/rand/v2"
// 	"net/http"
// 	"net/url"
// 	"strings"

// 	"github.com/jackc/pgx/v5"

// 	"time"
// )

// type SpotifyTokenResponse struct {
// 	AccessToken string `json:"access_token"`
// 	TokenType   string `json:"token_type"`
// 	ExpiresIn   int    `json:"expires_in"`
// }

// func GetSpotifyToken(db *pgx.Conn, clientID, clientSecret, grant_type, redirectURI string, userID uint64) error {
// 	spotify_url := "https://accounts.spotify.com/api/token"
// 	data := url.Values{}

// 	data.Set("grant_type", grant_type)
// 	data.Set("client_id", clientID)
// 	data.Set("client_secret", clientSecret)
// 	req, err := http.NewRequest("POST", spotify_url, strings.NewReader(data.Encode()))
// 	if err != nil {
// 		return err
// 	}

// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		body, _ := io.ReadAll(resp.Body)
// 		return fmt.Errorf("spotify token request failed: %s", string(body))
// 	}
// 	var tokenResp SpotifyTokenResponse
// 	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
// 		return err
// 	}

// 	var newUserID int64
// 	err = db.QueryRow(
// 		context.Background(),
// 		"SELECT create_user($1, $2, $3)",
// 		clientID, "Moj", "john@example.com",
// 	).Scan(&newUserID)

// 	if err != nil {
// 		return fmt.Errorf("failed to create user: %w", err)
// 	}

// 	expiresAt := time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second)

// 	var token string
// 	err = db.QueryRow(
// 		context.Background(),
// 		"CALL upsert_spotify_token($1, $2, $3, $4)",
// 		newUserID, tokenResp.AccessToken, tokenResp.TokenType, expiresAt,
// 	).Scan(&token)
// 	if err != nil {
// 		return fmt.Errorf("failed to upsert Spotify token: %w", err)
// 	}

// 	return err
// }

// func CallbackHandler(db *pgx.Conn, clientID, grant_type, clientSecret, redirectURI string) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		userID := rand.Uint64()

// 		err := GetSpotifyToken(db, clientID, clientSecret, grant_type, redirectURI, userID)
// 		if err != nil {
// 			http.Error(w, fmt.Sprintf("failed to get token: %v", err), http.StatusInternalServerError)
// 			return
// 		}

// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte("Spotify token saved successfully!"))
// 	}
// }
