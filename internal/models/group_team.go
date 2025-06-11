package models

import (
	"gorm.io/gorm"
)

type GroupTeam struct {
	gorm.Model

	GroupID uint  `json:"group_id"`
	Group   Group `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"group"`

	TeamID uint `json:"team_id"`
	Team   Team `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"team"`

	// Données contextuelles
	Ranking *int    `json:"ranking,omitempty"` // Classement dans le groupe
	Note    *string `json:"note,omitempty"`    // Commentaire spécifique
}
