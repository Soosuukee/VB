package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username string `gorm:"size:100;not null;unique"`
	Email    string `gorm:"size:255;not null;unique"`
	Password string `gorm:"not null"` // hashé

	// Optionnel : rôle simple en string ou JSON/array plus tard
	Role string `gorm:"size:50;default:'user'"`

	// Relations
	// CreatedEvents []Event      `gorm:"foreignKey:CreatedByID"`
	// EventUsers    []EventUser  `gorm:"foreignKey:UserID"`
	// TeamMembers   []TeamMember `gorm:"foreignKey:UserID"`
}
