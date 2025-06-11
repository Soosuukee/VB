package models

import (
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model

	Name           string   `gorm:"size:255;not null" json:"name"`
	Slug           string   `gorm:"size:255;not null;unique" json:"slug"`
	Color          string   `gorm:"size:255;not null;unique" json:"color"`
	Type           string   `gorm:"size:50;not null" json:"type"`
	Icon           string   `gorm:"size:255;not null" json:"icon"`
	Banner         string   `gorm:"size:255;not null" json:"banner"`
	Heroes         string   `gorm:"size:255;not null" json:"heroes"`
	IsAvailable    bool     `gorm:"not null" json:"is_available"`
	AvailableModes []string `gorm:"type:json" json:"available_modes"`
	Platforms      []string `gorm:"type:json" json:"platforms"`

	// Relations
	Events []Event `gorm:"foreignKey:GameID" json:"events,omitempty"`
}
