package store

import "time"

type TaskStore interface {
	GetAll() ([]Task, error)
	Add(Task) error
	ToggleCompleted(string) error
}

type Task struct {
	ID          int
	Name        string
	Completed   bool
	CreatedDate time.Time
}

var GlobalTaskStore TaskStore
