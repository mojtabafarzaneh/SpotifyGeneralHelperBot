package models

type Artist struct {
	ID   string `gorm:"primaryKey;size:50"`
	Name string `gorm:"size:255"`
	Href string `gorm:"size:255"`
}
