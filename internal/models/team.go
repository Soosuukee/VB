package models

import "gorm.io/gorm"

type Team struct {
	gorm.Model

	Name  string `gorm:"size:255;not null" json:"name"`
	Image string `gorm:"size:255;not null" json:"image"`

	EventParticipations []EventTeam        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"event_participations"`
	Phases              []Phase            `gorm:"many2many:phase_teams;" json:"phases"`
	MatchParticipants   []MatchParticipant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"match_game_participants"`
	TeamMembers         []TeamMember       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"team_members"`
	Groups              []Group            `gorm:"many2many:group_teams;" json:"groups"`
}

// Méthode pour récupérer le capitaine
func (t *Team) GetCaptain() *User {
	for _, member := range t.TeamMembers {
		if member.IsCaptain != nil && *member.IsCaptain {
			return &member.User
		}
	}
	return nil
}
