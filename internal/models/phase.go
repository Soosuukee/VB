package models

import (
	"time"

	"gorm.io/gorm"
)

type Phase struct {
	gorm.Model

	Name              string `gorm:"size:255;not null" json:"name"`
	Type              string `gorm:"size:255;not null" json:"type"`                // ex: "group", "bracket", "gsl"
	RoundLabelingMode string `gorm:"size:255;not null" json:"round_labeling_mode"` // ex: "Round 1", "Winner Round", etc.

	StartDate  time.Time `gorm:"not null" json:"start_date"`
	EndDate    time.Time `gorm:"not null" json:"end_date"`
	IsFinished *bool     `gorm:"not null" json:"is_finished"`
	PhaseOrder *int      `json:"phase_order,omitempty"`

	// Relations
	EventID uint  `json:"event_id"`
	Event   Event `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"event"`

	Teams        []Team        `gorm:"many2many:phase_teams;" json:"teams"`
	Matches      []Match       `gorm:"foreignKey:PhaseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"matches"`
	ScoringRules []ScoringRule `gorm:"foreignKey:PhaseID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"scoring_rules"`
}
