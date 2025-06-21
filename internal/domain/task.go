package domain

import "time"

type Task struct {
	title       string
	description string
	status      string
	createdAt   time.Time
	updatedAt   time.Time
}

type TaskRepo interface {
	Create(task Task) error
	Update(task Task) error
	Delete(task Task) error
}
