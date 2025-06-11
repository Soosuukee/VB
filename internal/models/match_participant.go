package models

import (
	"gorm.io/gorm"
)

type MatchParticipant struct {
	gorm.Model

	// Relations
	MatchID uint
	Match   Match `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	TeamID *uint
	Team   *Team `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	PlayerID *uint
	Player   *User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	UserID uint
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	// Statistiques globales
	Score             *int    `json:"score"`
	Placement         int     `gorm:"not null" json:"placement"`
	IsWinner          *bool   `json:"is_winner"`
	IsDisqualified    *bool   `gorm:"default:false" json:"is_disqualified"`
	EliminationsTotal *int    `json:"eliminations_total"`
	Status            *string `gorm:"size:50" json:"status"` // ex: Forfeit, Absent, No-show
	Note              *string `gorm:"size:255" json:"note"`

	// Statistiques individuelles (liées à ce participant)
	UserStats []UserStats `gorm:"foreignKey:ParticipantID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"player_stats"`
}
