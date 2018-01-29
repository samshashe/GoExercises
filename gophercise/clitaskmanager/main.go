package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/boltdb/bolt"
)

func init() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	var err error
	db, err = setupDB()
	if err != nil {
		log.Fatal(err)
	}
}

var help = `task is a CLI for managing your TODOs.

Usage:
  task [command]

Available Commands:
  add         Add a new task to your TODO list
  do          Mark a task on your TODO list as complete
  list        List all of your incomplete tasks

Use "task [command] --help" for more information about a command.`
var db *bolt.DB

func main() {
	if len(os.Args) == 1 {
		fmt.Println(help)
	}
	if len(os.Args) == 2 && os.Args[1] == "list" {
		//fmt.Println(os.Args[0], os.Args[1])
		tasks := GetTasks(db, "Task")
		var count int
		for _, task := range tasks {
			if task.Completed {
				count++
				fmt.Println(count, ".", task.Name)
			}
		}
	}
	if len(os.Args) == 3 {
		if cmd := os.Args[1]; cmd == "add" || cmd == "do" {
			taskName := os.Args[2]

			switch cmd {
			case "add":
				fmt.Println("add called", taskName)
				task := Task{0, taskName, true}
				AddTask(db, &task)
			case "do":
				fmt.Println("do called", taskName)
				task := Task{0, taskName, true}
				MarkTaskDone(db, &task)
			}
		}
	}

}

type Task struct {
	ID        int
	Name      string
	Completed bool
}

func setupDB() (*bolt.DB, error) {
	db, err := bolt.Open("task.db", 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("could not open db, %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		root, err := tx.CreateBucketIfNotExists([]byte("DB"))
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists([]byte("Task"))
		if err != nil {
			return fmt.Errorf("could not create task bucket: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not set up buckets, %v", err)
	}
	fmt.Println("DB Setup Done")
	return db, nil
}

// CreateTak saves t to the db. The new task ID is set on t once the data is persisted.
func AddTask(db *bolt.DB, t *Task) error {
	return db.Update(func(tx *bolt.Tx) error {
		// Retrieve the tasks bucket.
		// This should be created when the DB is first opened.
		b := tx.Bucket([]byte("Task"))

		// Generate ID for the task.
		// This returns an error only if the Tx is closed or not writeable.
		// That can't happen in an Update() call so I ignore the error check.
		id, _ := b.NextSequence()
		t.ID = int(id)

		// Marshal user data into bytes.
		buf, err := json.Marshal(t)
		if err != nil {
			return err
		}

		// Persist bytes to users bucket.
		return b.Put(itob(t.ID), buf)
	})
}

func MarkTaskDone(db *bolt.DB, t *Task) error {
	return db.Update(func(tx *bolt.Tx) error {
		// Retrieve the users bucket.
		// This should be created when the DB is first opened.
		b := tx.Bucket([]byte("Task"))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			value, err := decode(v)
			if err != nil {
				return err
			}
			if t.Name == value.Name {
				t.Completed = true
				break
			}

			// fmt.Printf("key=%s, value=%s\n", k, v)
		}

		// Marshal user data into bytes.
		buf, err := json.Marshal(t)
		if err != nil {
			return err
		}

		// Persist bytes to users bucket.
		return b.Put(itob(t.ID), buf)
	})
}

func GetTasks(db *bolt.DB, bucket string) []Task {
	var tasks []Task
	db.View(func(tx *bolt.Tx) error {

		c := tx.Bucket([]byte(bucket)).Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			value, err := decode(v)
			if err != nil {
				return err
			}
			tasks = append(tasks, *value)

			// fmt.Printf("key=%s, value=%s\n", k, v)
		}
		return nil
	})
	return tasks
}
func decode(data []byte) (*Task, error) {
	var t *Task
	err := json.Unmarshal(data, &t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
