package models

import (
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model

	Name           string   `gorm:"size:255;not null"`
	Slug           string   `gorm:"size:255;not null;unique"`
	Color          string   `gorm:"size:255;not null;unique"`
	Type           string   `gorm:"size:50;not null"`
	Icon           string   `gorm:"size:255;not null"`
	Banner         string   `gorm:"size:255;not null"`
	Heroes         string   `gorm:"size:255;not null"`
	IsAvailable    bool     `gorm:"not null"`
	AvailableModes []string `gorm:"type:json"` // PostgreSQL accepte JSON natif
	Platforms      []string `gorm:"type:json"` // idem

	// Relations
	Events []Event `gorm:"foreignKey:GameID"`
}
