package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username string `gorm:"size:100;not null;unique" json:"username"`
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"not null" json:"-"` // ⚠️ jamais exposé dans les réponses

	Role string `gorm:"size:50;default:'user'" json:"role"` // "user", "admin", etc.

	// Relations
	CreatedEvents     []Event            `gorm:"foreignKey:CreatedByID" json:"created_events,omitempty"`
	EventUsers        []EventUser        `gorm:"foreignKey:UserID" json:"event_users,omitempty"`
	TeamMembers       []TeamMember       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"team_members,omitempty"`
	MatchParticipants []MatchParticipant `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"match_participants,omitempty"`
	UserStats         []UserStats        `gorm:"foreignKey:UserID" json:"user_stats,omitempty"`
}
