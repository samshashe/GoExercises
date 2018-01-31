package store

import "time"

type TaskStore interface {
	Add(Task) error
	GetTasks(bool) ([]Task, error)
	ToggleCompleted(string) error
}

type Task struct {
	ID          int
	Name        string
	Completed   bool
	CreatedDate time.Time
}

var GlobalTaskStore TaskStore

type ByTask []Task

func (t ByTask) Len() int {
	return len(t)
}
func (t ByTask) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func (t ByTask) Less(i, j int) bool {
	return t[i].ID < t[j].ID
}
