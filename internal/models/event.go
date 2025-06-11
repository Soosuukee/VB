package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model

	Name            string    `gorm:"size:255;not null" json:"name"`
	Date            time.Time `gorm:"not null" json:"date"`
	CreatedAt       time.Time `gorm:"not null" json:"created_at"`
	Image           *string   `gorm:"size:255" json:"image,omitempty"`
	RequiredPlayers int       `gorm:"not null" json:"required_players"`
	Description     string    `gorm:"type:text" json:"description"`
	Mode            *string   `gorm:"size:100" json:"mode,omitempty"`
	ScoringMode     string    `gorm:"size:50;default:standard" json:"scoring_mode"`
	Category        *string   `gorm:"size:100" json:"category,omitempty"`
	Format          *string   `gorm:"size:20" json:"format,omitempty"`
	IsSolo          *bool     `gorm:"not null" json:"is_solo"`

	// Relations
	GameID uint `json:"game_id"`
	Game   Game `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"game"`

	CreatedByID uint `json:"created_by_id"`
	CreatedBy   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"created_by"`

	ScoringRules []ScoringRule `gorm:"foreignKey:EventID" json:"scoring_rules,omitempty"`
	Participants []EventTeam   `gorm:"foreignKey:EventID" json:"participants,omitempty"`
	Phases       []Phase       `gorm:"foreignKey:EventID" json:"phases,omitempty"`
	EventUsers   []EventUser   `gorm:"foreignKey:EventID" json:"event_users,omitempty"`
}
