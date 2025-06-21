package domain

import "time"

type User struct {
	username     string
	email        string
	passwordHash string
	createdAt    time.Time
	updatedAt    time.Time
}

type UserRepo interface {
	Create(user User) error
	FindByUsername(username string) (User, error)
	FindByEmail(email string) (User, error)
	Update(user User) error
	Delete(user User) error
}
