package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model

	Name            string    `gorm:"size:255;not null"`
	Date            time.Time `gorm:"not null"`
	CreatedAt       time.Time `gorm:"not null"`
	Image           *string   `gorm:"size:255"`
	RequiredPlayers int       `gorm:"not null"`
	Description     string    `gorm:"type:text"`
	Mode            *string   `gorm:"size:100"`
	ScoringMode     string    `gorm:"size:50;default:standard"`
	Category        *string   `gorm:"size:100"`
	Format          *string   `gorm:"size:20"`
	IsSolo          *bool     `gorm:"not null"`

	// Relations
	GameID uint
	Game   Game `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	CreatedByID uint
	CreatedBy   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	ScoringRules []ScoringRule `gorm:"foreignKey:EventID"`
	Participants []EventTeam   `gorm:"foreignKey:EventID"`
	Phases       []Phase       `gorm:"foreignKey:EventID"`
	EventUsers   []EventUser   `gorm:"foreignKey:EventID"`
}
