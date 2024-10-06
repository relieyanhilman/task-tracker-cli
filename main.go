package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task-cli [command] [arguments]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli add <task-name>")
			return
		}
		taskDescription := os.Args[2]
		err := AddTask(taskDescription)
		if err != nil {
			fmt.Println("Error adding task:", err)
		}
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: task-cli update <id_task> <description>")
			return
		}
		taskID := os.Args[2]
		taskDescription := os.Args[3]
		err := UpdateTask(taskID, taskDescription)
		if err != nil {
			fmt.Println("Error updating task:", err)
		}

	case "delete":
        if len(os.Args) != 3 {
            fmt.Println("Usage: task-cli delete <id_task>")
            return
        }
        taskID := os.Args[2]
        err := DeleteTask(taskID)
        if err != nil {
            fmt.Println("Error deleting task:", err)
        }

	default:
		fmt.Println("Unknown command:", command)
	}
}
