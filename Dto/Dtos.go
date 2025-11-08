package dto

import "time"

type SavedAlbumsResponse struct {
	Href     string           `json:"href"`
	Limit    int              `json:"limit"`
	Next     string           `json:"next"`
	Offset   int              `json:"offset"`
	Previous string           `json:"previous"`
	Total    int              `json:"total"`
	Items    []SavedAlbumItem `json:"items"`
}

type SavedAlbumItem struct {
	AddedAt time.Time `json:"added_at"`
	Album   AlbumDTO  `json:"album"`
}

type AlbumDTO struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	AlbumType   string      `json:"album_type"`
	TotalTracks int         `json:"total_tracks"`
	ReleaseDate string      `json:"release_date"`
	Href        string      `json:"href"`
	Images      []ImageDTO  `json:"images"`
	Artists     []ArtistDTO `json:"artists"`
}

type ArtistDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Href string `json:"href"`
}

type ImageDTO struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type SpotifyTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}
