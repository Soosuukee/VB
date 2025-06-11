package models

import (
	"time"

	"gorm.io/gorm"
)

type Group struct {
	gorm.Model

	Name      string    `gorm:"size:100;not null" json:"name"`
	PhaseID   uint      `gorm:"not null" json:"phase_id"`
	Phase     Phase     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"phase"`
	CreatedAt time.Time `json:"created_at"`

	// Relations
	GroupUsers []GroupUser `gorm:"constraint:OnDelete:CASCADE;" json:"group_users"`
	GroupTeams []GroupTeam `gorm:"constraint:OnDelete:CASCADE;" json:"group_teams"`
	Groups     []Group     `gorm:"foreignKey:PhaseID" json:"groups"`
}
