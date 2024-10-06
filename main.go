package main

import (
	"fmt"
	"os"
)

// main is the entry point of the command-line interface of the task tracker.
//
// The supported commands are:
//
//   - add: Adds a new task with the given description.
//
//     Usage: task-cli add <task-name>
//
//   - update: Updates the task with the given ID and description.
//
//     Usage: task-cli update <id_task> <description>
//
//   - delete: Deletes the task with the given ID.
//
//     Usage: task-cli delete <id_task>
//
//   - mark-in-progress: Marks the task with the given ID as in-progress.
//
//     Usage: task-cli mark-in-progress <task_id>
//
// The program prints an error message and returns if any of the commands is
// called with the wrong number of arguments.
//
// The program exits with an error code if any of the commands fails.
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

	case "mark-in-progress":
		if len(os.Args) != 3 {
			fmt.Println("Usage: task-cli mark-in-progress <task_id>")
			return
		}
		taskID := os.Args[2]
		err := MarkTask(taskID, "in-progress")
		if err != nil {
			fmt.Println("Error marking task as in-progress:", err)
		}

	default:
		fmt.Println("Unknown command:", command)
	}
}
