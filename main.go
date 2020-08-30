package main

import (
	"fmt"
)

func listCommand() {
	fmt.Println("========================")
	fmt.Println("List command available:")
	fmt.Println("1. task-todo : to find out todo task that still not be done")
	fmt.Println("2. task-done : to find out task that had been done")
	fmt.Println("3. check-task --id (id of task) --rev (rev of task) : move task from todo to done")
	fmt.Println("========================")
}

func main() {
	initTasks()
	// if len(os.Args) < 2 {
	// 	fmt.Println("Error: subcommand is required")
	// 	listCommand()
	// 	os.Exit(1)
	// }

	// switch os.Args[1] {
	// case "task-todo":
	// 	resp, err := http.Get(baseURL + "/efishery_task")
	// case "task-done":
	// 	fmt.Println("task-done")
	// case "check-task":
	// 	fmt.Println("check-task")
	// default:
	// 	fmt.Println("Error: subcommand is wrong")
	// 	listCommand()
	// 	os.Exit(1)
	// }
}
