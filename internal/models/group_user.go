package models

import (
	"gorm.io/gorm"
)

type GroupUser struct {
	gorm.Model

	GroupID uint  `json:"group_id"`
	Group   Group `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"group"`

	UserID uint `json:"user_id"`
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"user"`

	// Données contextuelles
	Ranking *int    `json:"ranking,omitempty"` // Position dans le groupe
	Note    *string `json:"note,omitempty"`    // Commentaire admin
}
