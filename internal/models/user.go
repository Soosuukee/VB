package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username string `gorm:"size:100;not null;unique"`
	Email    string `gorm:"size:255;not null;unique"`
	Password string `gorm:"not null"` // stocke le hash

	Role string `gorm:"size:50;default:'user'"` // ex: "user", "admin"

	// Relations
	CreatedEvents     []Event            `gorm:"foreignKey:CreatedByID"`
	EventUsers        []EventUser        `gorm:"foreignKey:UserID"`
	TeamMembers       []TeamMember       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	MatchParticipants []MatchParticipant `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	UserStats         []UserStats        `gorm:"foreignKey:UserID"` // anciennement PlayerStats
}
