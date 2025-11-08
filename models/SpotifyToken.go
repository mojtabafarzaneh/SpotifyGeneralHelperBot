package models

import "time"

type SpotifyToken struct {
	AccessToken  string    `json:"access_token" gorm:"type:text"`
	TokenType    string    `json:"token_type" gorm:"size:50"`
	ExpiresIn    int       `json:"expires_in" gorm:"-"`
	ExpiresAt    time.Time `json:"-" gorm:"not null"`
	RefreshToken string    `json:"refresh_token,omitempty" gorm:"type:text"`
}
