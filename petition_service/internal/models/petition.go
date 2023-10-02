package models

import "gorm.io/gorm"

type Petition struct {
	gorm.Model
	Title       string `gorm:"not null;" json:"title"`
	Category    string `gorm:"not null;" json:"category"`
	Description string `gorm:"not null;" json:"description"`
	Image       string `gorm:"not null;" json:"image"`
	StatusID    uint   `gorm:"not null;" json:"status_id"`
	Status      Status `gorm:"foreignKey:StatusID" json:"status"`
	UserID      uint   `gorm:"not null;" json:"user_id"`
	VoteGoal    uint   `gorm:"not null;default:1000" json:"vote_goal"`
	CurrVotes   uint   `gorm:"not null;default:0" json:"curr_votes"`
}
