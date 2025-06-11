package models

import (
	"gorm.io/gorm"
)

type MatchParticipant struct {
	gorm.Model

	// Relations
	MatchID uint  `json:"match_id"`
	Match   Match `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"match"`

	TeamID *uint `json:"team_id,omitempty"`
	Team   *Team `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"team,omitempty"`

	UserID uint `json:"user_id"`
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`

	// Statistiques globales
	Score             *int    `json:"score,omitempty"`
	Placement         int     `gorm:"not null" json:"placement"`
	IsWinner          *bool   `json:"is_winner,omitempty"`
	IsDisqualified    *bool   `gorm:"default:false" json:"is_disqualified,omitempty"`
	EliminationsTotal *int    `json:"eliminations_total,omitempty"`
	Status            *string `gorm:"size:50" json:"status,omitempty"` // ex: Forfeit, Absent, No-show
	Note              *string `gorm:"size:255" json:"note,omitempty"`

	// Statistiques individuelles
	UserStats []UserStats `gorm:"foreignKey:ParticipantID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user_stats"`
}
