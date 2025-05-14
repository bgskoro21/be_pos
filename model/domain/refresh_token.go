package domain

import (
	"time"
)

type RefreshToken struct{
	ID 		uint	`gorm:"primaryKey"`
	UserID	uint	`gorm:"index"`
	Token	string	`gorm:"unique"`
	ExpiresAt	time.Time
	CreatedAt	time.Time
	UserAgent	string
	IPAddress	string

	User User `gorm:"foreignKey:UserID"`
}