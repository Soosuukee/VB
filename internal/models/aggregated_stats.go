package models

import (
	"gorm.io/gorm"
)

type AggregatedStats struct {
	gorm.Model

	UserID uint `gorm:"not null" json:"user_id"`
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`

	GameID uint `gorm:"not null" json:"game_id"`
	Game   Game `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"game"`

	Season string `gorm:"size:50;not null" json:"season"` // Exemple : "2025S1"

	TotalMatches int     `gorm:"default:0" json:"total_matches"`
	TotalKills   int     `gorm:"default:0" json:"total_kills"`
	TotalDamage  int     `gorm:"default:0" json:"total_damage"`
	AverageScore float64 `gorm:"default:0" json:"average_score"`
	WinRate      float64 `gorm:"default:0" json:"win_rate"`   // En pourcentage
	Top1Count    int     `gorm:"default:0" json:"top1_count"` // Nombre de Top 1 (ex: en BR)
}
