package models

import (
	"gorm.io/gorm"
)

type GroupTeam struct {
	gorm.Model

	GroupID uint
	Group   Group `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	TeamID uint
	Team   Team `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// Données contextuelles
	Ranking *int    `json:"ranking,omitempty"` // Classement dans le groupe
	Note    *string `json:"note,omitempty"`    // Commentaire spécifique
}
