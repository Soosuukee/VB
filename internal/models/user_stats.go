package models

import "gorm.io/gorm"

type UserStats struct {
	gorm.Model

	// Relations
	ParticipantID uint             `gorm:"not null" json:"participant_id"`
	Participant   MatchParticipant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"participant"`

	UserID uint `gorm:"not null" json:"user_id"`
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`

	// Statistiques de gameplay (match unique)
	Eliminations int      `gorm:"not null;default:0" json:"eliminations"`
	Deaths       *int     `gorm:"default:0" json:"deaths,omitempty"`
	Assists      *int     `gorm:"default:0" json:"assists,omitempty"`
	Damage       *int     `gorm:"default:0" json:"damage,omitempty"`
	Headshots    *int     `gorm:"default:0" json:"headshots,omitempty"`
	PlayerScore  *float64 `gorm:"default:0" json:"player_score,omitempty"`

	// Tracking externe
	ExternalID *string `gorm:"size:100" json:"external_id,omitempty"` // ex: Riot/Epic ID
	Source     *string `gorm:"size:50" json:"source,omitempty"`       // ex: "api", "replay", "manual"
}
