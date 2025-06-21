package models

import "time"

type TaskDoc struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	Title       string    `gorm:"not null"`
	Description string    `gorm:"default: ' '"`
	Status      string    `gorm:"default:'created'"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP; autoUpdateTime"`
	UserID      uint      `gorm:"not null"`

	User UserDoc `gorm:"foreignKey:UserID; constraint:onCascade:DELETE"`
}
