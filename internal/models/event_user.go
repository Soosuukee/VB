package models

import "gorm.io/gorm"

type EventUser struct {
	gorm.Model

	// Relations
	EventID uint
	Event   Event `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	UserID uint
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// Données
	PointsTotal *int    `json:"points_total"`
	Ranking     *int    `json:"ranking"`
	Note        *string `gorm:"type:text" json:"note"`
	IsValidated bool    `gorm:"not null;default:false" json:"is_validated"`
}
