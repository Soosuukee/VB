package models

import (
	"time"

	"gorm.io/gorm"
)

type Group struct {
	gorm.Model

	Name      string `gorm:"size:100;not null"`
	PhaseID   uint   `gorm:"not null"`
	Phase     Phase  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time

	// Relations
	GroupUsers []GroupUser `gorm:"constraint:OnDelete:CASCADE;"`
	GroupTeams []GroupTeam `gorm:"constraint:OnDelete:CASCADE;"`
}
