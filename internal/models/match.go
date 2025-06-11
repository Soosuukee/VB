package models

import (
	"time"

	"gorm.io/gorm"
)

type Match struct {
	gorm.Model

	// Relations
	PhaseID uint
	Phase   Phase `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// Données de planification
	StartDate time.Time `gorm:"not null" json:"start_date"`
	EndDate   time.Time `gorm:"not null" json:"end_date"`

	// Statut du match
	IsFinished *bool   `gorm:"not null" json:"is_finished"`
	Status     string  `gorm:"size:20;default:'pending'" json:"status"` // ex: pending, active, finished, canceled
	MatchType  *string `gorm:"size:50" json:"match_type"`               // ex: "Bo1", "Bo3", "Battle Royale", "Custom"

	// Infos de lobby
	RoundLabel string  `gorm:"size:255;not null" json:"round_label"` // ex: "Finale", "Round 3"
	LobbyCode  *string `gorm:"size:20" json:"lobby_code"`            // ex: code pour rejoindre le lobby (si applicable)

	// Participants
	Participants []MatchParticipant `gorm:"foreignKey:MatchID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"participants"`
}

// TableName permet d'éviter un conflit éventuel avec le mot "match" en SQL
func (Match) TableName() string {
	return "matches"
}
