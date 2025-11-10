package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	dto "github.com/mojtabafarzaneh/SpotifyGeneralHelperBot/Dto"
	"golang.org/x/oauth2"
)

func AlbumsHandler(w http.ResponseWriter, r *http.Request) {
	if demoToken == nil {
		http.Error(w, "no token available. Visit /login first", http.StatusUnauthorized)
		return
	}

	ctx := context.Background()
	client := oauth2.NewClient(ctx, oauth2.StaticTokenSource(demoToken))

	var allAlbums []dto.Album
	limit := 50
	offset := 0

	for {
		url := fmt.Sprintf("https://api.spotify.com/v1/me/albums?limit=%d&offset=%d", limit, offset)
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			http.Error(w, "failed to create request: "+err.Error(), http.StatusInternalServerError)
			return
		}
		req.Header.Set("Accept", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, "spotify request failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			http.Error(w, fmt.Sprintf("spotify returned %d: %s", resp.StatusCode, string(body)), resp.StatusCode)
			return
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "failed to read response: "+err.Error(), http.StatusInternalServerError)
			return
		}

		var raw struct {
			Items []struct {
				Album dto.Album `json:"album"`
			} `json:"items"`
			Total int `json:"total"`
		}

		if err := json.Unmarshal(body, &raw); err != nil {
			http.Error(w, "failed to unmarshal: "+err.Error(), http.StatusInternalServerError)
			return
		}

		if len(raw.Items) == 0 {
			break
		}

		for _, item := range raw.Items {
			allAlbums = append(allAlbums, item.Album)
		}

		offset += len(raw.Items)
		if offset >= raw.Total {
			break
		}
	}

	if len(allAlbums) == 0 {
		http.Error(w, "no albums found", http.StatusNotFound)
		return
	}

	result, _ := json.MarshalIndent(allAlbums, "", "  ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
