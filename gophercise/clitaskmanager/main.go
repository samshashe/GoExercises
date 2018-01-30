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
		fmt.Println("list called with", os.Args[0], os.Args[1])
		tasks, err := store.GlobalTaskStore.GetAll()
		if err != nil {
			log.Fatal(err)
		}
		var count int
		for _, task := range tasks {
			if task.Completed == false {
				count++
				fmt.Println(count, ".", task.Name)
			}
		}
	}
	if len(os.Args) == 3 {
		if cmd := os.Args[1]; cmd == "add" || cmd == "do" {
			nameorid := os.Args[2]

			switch cmd {
			case "add":
				fmt.Println("add called", nameorid)
				task := store.Task{ID: 0, Name: nameorid, Completed: false, CreatedDate: time.Now()}
				store.GlobalTaskStore.Add(task)
			case "do":
				fmt.Println("do called", nameorid)
				intValue, err := strconv.Atoi(nameorid)
				if err != nil {
					log.Fatal(err)
				}
				taskName := GetName(intValue)
				store.GlobalTaskStore.ToggleCompleted(taskName)
			}
		}
	}

}

func GetName(id int) string {
	if id < 1 {
		log.Fatal("Bad Argument. value has to be greater than 0")
	}
	tasks, err := store.GlobalTaskStore.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	count, i := -1, 0
	for ; count < id; i++ {
		if tasks[i].Completed == false {
			count++
		}
	}
	return tasks[i].Name
}
