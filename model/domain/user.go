package domain

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	RoleCostumer Role = "customer"
	RoleAdmin Role = "admin"
)

type User struct{
	ID			uint `gorm:"primaryKey"`
	Name		string `gorm:"not null"`
	Email 		*string `gorm:"uniqueIndex"`
	Password 	*string 
	Role 		Role
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	gorm.DeletedAt `gorm:"index"`
}