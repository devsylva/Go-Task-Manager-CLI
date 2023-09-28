package main

import (
	"fmt"
	"task-manager-cli/commands"
	"task-manager-cli/storage"
	"task-manager-cli/task"
)

func main() {
	var tasks []task.Task

	// Load existing tasks from storage (if available)
	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	for {
		fmt.Println("\nTask Manager CLI")
		fmt.Println("1. List tasks")
		fmt.Println("2. Add a task")
		fmt.Println("3. Mark task as complete")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			commands.ListTasks(tasks)
		case 2:
			commands.AddTask(&tasks)
		case 3:
			commands.MarkTaskComplete(&tasks)
		case 4:
			err := storage.SaveTasks(tasks)
			if err != nil {
				fmt.Println("Error saving tasks:", err)
			}
			fmt.Println("Exiting Task Manager CLI.")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}
