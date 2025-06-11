package models

import (
	"time"

	"gorm.io/gorm"
)

type TeamMember struct {
	gorm.Model

	// Relations
	UserID uint
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	TeamID uint
	Team   Team `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// Données
	IsCaptain *bool      `gorm:"not null" json:"is_captain"`
	JoinedAt  *time.Time `json:"joined_at"`
}
