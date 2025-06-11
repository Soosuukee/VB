package models

import "gorm.io/gorm"

type UserStats struct {
	gorm.Model

	// Relations
	ParticipantID uint
	Participant   MatchParticipant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	UserID uint
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// Statistiques de gameplay
	Eliminations int      `gorm:"not null" json:"eliminations"`
	Deaths       *int     `json:"deaths"`
	Assists      *int     `json:"assists"`
	Damage       *int     `json:"damage"`
	Headshots    *int     `json:"headshots"`
	PlayerScore  *float64 `json:"player_score"`

	// Tracking externe
	ExternalID *string `gorm:"size:100" json:"external_id"` // ID API externe (Riot, Epic, etc.)
	Source     *string `gorm:"size:50" json:"source"`       // ex: "api", "replay", "manual"
}
