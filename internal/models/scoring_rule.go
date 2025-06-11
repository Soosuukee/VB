package models

import (
	"gorm.io/gorm"
)

type ScoringRule struct {
	gorm.Model

	EventID uint
	Event   Event `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	PhaseID *uint
	Phase   *Phase `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Type     string `gorm:"size:20;not null"` // Exemple : "placement", "kill"
	Position int    `gorm:"not null"`         // Exemple : 1 = first place
	Points   int    `gorm:"not null"`         // Exemple : 10 points for first
}
