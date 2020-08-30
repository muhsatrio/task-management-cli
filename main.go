package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func listCommand() {
	fmt.Println("========================")
	fmt.Println("List command available:")
	fmt.Println("1. task-todo : to find out todo task that still not be done")
	fmt.Println("2. task-done : to find out task that had been done")
	fmt.Println("3. check-task : move task from todo to done")
	fmt.Println("========================")
}

func main() {
	initTasks()

	scanner := bufio.NewScanner(os.Stdin)

	if len(os.Args) < 2 {
		fmt.Println("Error: subcommand is required")
		listCommand()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "task-todo":
		initTasks()
		filteredTask := filterTask(false)
		if len(filteredTask) > 0 {
			for i, task := range filteredTask {
				fmt.Printf("%d. %s\n", i+1, task.Name)
				fmt.Printf("ID: %s\n", task.ID)
				fmt.Printf("Rev: %s\n", task.Rev)
				fmt.Printf("Description: %s\n", task.Description)
			}
		} else {
			fmt.Println("Task is empty")
			os.Exit(0)
		}
		// fmt.Println(foundTask)
	case "task-done":
		initTasks()
		filteredTask := filterTask(true)
		if len(filteredTask) > 0 {
			for i, task := range filteredTask {
				fmt.Printf("%d. %s\n", i+1, task.Name)
				fmt.Printf("ID: %s\n", task.ID)
				fmt.Printf("Rev: %s\n", task.Rev)
				fmt.Printf("Description: %s\n", task.Description)
			}
		} else {
			fmt.Println("Task is empty")
			os.Exit(0)
		}
	case "check-task":
		fmt.Print("Enter ID: ")
		scanner.Scan()
		idInputed := scanner.Text()
		fmt.Print("Enter Rev: ")
		scanner.Scan()
		revInputed := scanner.Text()
		if len(idInputed) == 0 || len(revInputed) == 0 {
			log.Fatalln("ID or Rev is required")
		}
		initTasks()
		foundTask := getTask(idInputed)
		foundTask.Completed = true
		changeTask(idInputed, revInputed, foundTask)
	default:
		fmt.Println("Error: subcommand is wrong")
		listCommand()
		os.Exit(1)
	}
}
