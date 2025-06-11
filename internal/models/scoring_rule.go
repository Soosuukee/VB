package models

import (
	"gorm.io/gorm"
)

type ScoringRule struct {
	gorm.Model

	EventID uint  `json:"event_id"`
	Event   Event `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"event"`

	PhaseID *uint  `json:"phase_id,omitempty"`
	Phase   *Phase `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"phase,omitempty"`

	Type     string `gorm:"size:20;not null" json:"type"` // Exemple : "placement", "kill"
	Position int    `gorm:"not null" json:"position"`     // Exemple : 1 = première place
	Points   int    `gorm:"not null" json:"points"`       // Exemple : 10 points pour la première place
}
