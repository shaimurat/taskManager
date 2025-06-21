package models

import "time"

type UserDoc struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	Username     string    `gorm:"not null"`
	Email        string    `gorm:"not null"`
	PasswordHash string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP; autoUpdateTime"`
}
