package store

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sort"
)

type FileTaskStore struct {
	fileName string
	Tasks    map[int]Task
}

func NewFileTaskStore(fileName string) (*FileTaskStore, error) {
	store := &FileTaskStore{
		Tasks:    map[int]Task{},
		fileName: fileName,
	}

	contents, err := ioutil.ReadFile(fileName)

	if err != nil {
		// If it's a matter of the file not existing, that's ok
		if os.IsNotExist(err) {
			return store, nil
		}
		return nil, err
	}
	err = json.Unmarshal(contents, store)
	if err != nil {
		return nil, err
	}
	return store, nil
}

func (store *FileTaskStore) Add(task Task) error {
	task.ID = len(store.Tasks)
	store.Tasks[task.ID] = task
	contents, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(store.fileName, contents, 0660)
}

func (store *FileTaskStore) GetTasks(state bool) ([]Task, error) {
	tasks := []Task{}
	for _, task := range store.Tasks {
		if task.Completed == state {
			tasks = append(tasks, task)
		}
	}
	sort.Sort(ByTask(tasks))

	return tasks, nil
}

func (store *FileTaskStore) ToggleCompleted(name string) error {
	for i, task := range store.Tasks {
		if store.Tasks[i].Name == name {
			task.Completed = true
			store.Tasks[i] = task
		}
	}
	contents, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(store.fileName, contents, 0660)
}

func (store *FileTaskStore) Delete(name string) error {

	for key, task := range store.Tasks {
		if task.Name == name {
			delete(store.Tasks, key)
		}
	}
	return nil
}
