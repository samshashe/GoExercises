package store

import "time"

type TaskStore interface {
	GetTask(bool) ([]Task, error)
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
