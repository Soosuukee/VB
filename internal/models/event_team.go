package models

import "gorm.io/gorm"

type EventTeam struct {
	gorm.Model

	// Relations
	TeamID uint
	Team   Team `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	EventID uint
	Event   Event `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// Données
	IsValidated *bool   `gorm:"not null" json:"is_validated"`
	PointsTotal *int    `json:"points_total"`          // nullable
	Ranking     *int    `json:"ranking"`               // nullable
	Note        *string `gorm:"type:text" json:"note"` // nullable
}
