package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"GoExercises/gophercise/clitaskmanager/store"
)

func init() {
	db, err := store.NewMySQLDB("root:iis6!dfu@tcp(127.0.0.1:3306)/", "Task")
	if err != nil {
		panic(err)
	}
	store.GlobalMySQLDB = db
	store.GlobalTaskStore = store.NewDBTaskStore()
}

var help = `task is a CLI for managing your TODOs.

Usage:
  task [command]

Available Commands:
  add         Add a new task to your TODO list
  do          Mark a task on your TODO list as complete
  list        List all of your incomplete tasks

Use "task [command] --help" for more information about a command.`

func main() {
	if len(os.Args) == 1 {
		fmt.Println(help)
	}
	if len(os.Args) == 2 && os.Args[1] == "list" {
		tasks, err := store.GlobalTaskStore.GetTask(false)
		if err != nil {
			log.Fatal(err)
		}
		for i, task := range tasks {
			fmt.Println(i+1, ".", task.Name)
		}
		if len(tasks) < 1 {
			fmt.Println("you don't have any todos. use 'Task add taskName' to add a new one")
		}
	}
	if len(os.Args) == 3 {
		if cmd := os.Args[1]; cmd == "add" || cmd == "do" {
			nameorid := os.Args[2]

			switch cmd {
			case "add":
				task := store.Task{ID: 0, Name: nameorid, Completed: false, CreatedDate: time.Now()}
				store.GlobalTaskStore.Add(task)
			case "do":
				intValue, err := strconv.Atoi(nameorid)
				if err != nil {
					log.Fatal(err)
				}
				if intValue < 1 {
					log.Fatal("input should be greater than 0.")
				}
				tasks, err := store.GlobalTaskStore.GetTask(false)
				if err != nil {
					log.Fatal(err)
				}
				if length := len(tasks); length < intValue {
					log.Fatal("input should be less than or equal to: ", length)
				}
				taskName := tasks[intValue-1].Name
				store.GlobalTaskStore.ToggleCompleted(taskName)
			}
		}
	}

}
