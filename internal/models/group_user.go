package models

import (
	"gorm.io/gorm"
)

type GroupUser struct {
	gorm.Model

	GroupID uint
	Group   Group `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	UserID uint
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	// Données contextuelles
	Ranking *int    `json:"ranking,omitempty"` // Position dans le groupe
	Note    *string `json:"note,omitempty"`    // Commentaire admin
}
