package models

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {
	ID			int			`gorm:"primaryKey,autoIncrement" json:"id"`
	Name		string		`json:"name" gorm:"unique"`
	CreatedAt	time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt	time.Time	`gorm:"defualt:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt	gorm.DeletedAt	`gorm:"index"`
}