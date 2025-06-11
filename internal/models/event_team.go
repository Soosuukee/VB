package models

import "gorm.io/gorm"

type EventTeam struct {
	gorm.Model

	// Relations
	TeamID uint `gorm:"not null" json:"team_id"`
	Team   Team `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"team"`

	EventID uint  `gorm:"not null" json:"event_id"`
	Event   Event `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"event"`

	// Données
	IsValidated *bool   `gorm:"not null;default:false" json:"is_validated"` // Validation de participation
	PointsTotal *int    `gorm:"default:0" json:"points_total,omitempty"`    // Total de points attribués à l'équipe
	Ranking     *int    `json:"ranking,omitempty"`                          // Position globale
	Note        *string `gorm:"type:text" json:"note,omitempty"`            // Commentaire éventuel
}
