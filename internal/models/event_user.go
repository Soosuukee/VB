package models

import "gorm.io/gorm"

type EventUser struct {
	gorm.Model

	// Relations
	EventID uint  `gorm:"not null" json:"event_id"`
	Event   Event `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"event"`

	UserID uint `gorm:"not null" json:"user_id"`
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`

	// Données
	PointsTotal *int    `gorm:"default:0" json:"points_total,omitempty"`          // Total de points attribués dans l'event
	Ranking     *int    `json:"ranking,omitempty"`                                // Position globale dans l'événement
	Note        *string `gorm:"type:text" json:"note,omitempty"`                  // Note éventuelle laissée par un admin
	IsValidated bool    `gorm:"not null;default:false;index" json:"is_validated"` // Statut de validation
}
