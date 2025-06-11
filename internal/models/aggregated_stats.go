package models

import (
	"gorm.io/gorm"
)

type AggregatedStats struct {
	gorm.Model

	UserID uint
	User   User `gorm:"constraint:OnDelete:CASCADE;"`

	GameID uint
	Game   Game `gorm:"constraint:OnDelete:CASCADE;"`

	Season string `gorm:"size:50;not null"` // ex: "2025S1"

	TotalMatches int     `json:"total_matches"`
	TotalKills   int     `json:"total_kills"`
	TotalDamage  int     `json:"total_damage"`
	AverageScore float64 `json:"average_score"`
	WinRate      float64 `json:"win_rate"`   // en %
	Top1Count    int     `json:"top1_count"` // ex: BR
}
