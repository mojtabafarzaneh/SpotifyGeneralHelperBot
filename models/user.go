package models

type User struct {
	ID            uint64 `gorm:"primaryKey"`
	SpotifyUserID string `gorm:"uniqueIndex;size:100"`
	DisplayName   string `gorm:"size:255"`
	Email         string `gorm:"size:255"`
}
