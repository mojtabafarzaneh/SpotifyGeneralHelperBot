package models

import "time"

type Album struct {
	ID          string `gorm:"primaryKey;size:50"`
	Name        string `gorm:"size:255"`
	AlbumType   string `gorm:"size:50"`
	TotalTracks int
	ReleaseDate string `gorm:"size:20"`
	Href        string `gorm:"size:255"`
	ImageURL    string `gorm:"size:255"`
}

type AlbumArtist struct {
	AlbumID  string `gorm:"size:50;primaryKey"`
	ArtistID string `gorm:"size:50;primaryKey"`
}

type UserSavedAlbum struct {
	UserID  uint64 `gorm:"primaryKey"`
	AlbumID string `gorm:"size:50;primaryKey"`
	AddedAt time.Time
}
