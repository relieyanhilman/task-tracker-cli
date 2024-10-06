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
//   - mark-done: Marks the task with the given ID as done.
//
//     Usage: task-cli mark-done <task_id>
//
//   - list: Lists all tasks with the given status.
//
//     Usage: task-cli list || task-cli list todo || task-cli list in-progress || task-cli list done
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
		// Add a new task with the given description.
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
		// Update the task with the given ID and description.
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
		// Delete the task with the given ID.
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
		// Mark the task with the given ID as in-progress.
		if len(os.Args) != 3 {
			fmt.Println("Usage: task-cli mark-in-progress <task_id>")
			return
		}
		taskID := os.Args[2]
		err := MarkTask(taskID, "in-progress")
		if err != nil {
			fmt.Println("Error marking task as in-progress:", err)
		}

	case "mark-done":
		// Mark the task with the given ID as done.
		if len(os.Args) != 3 {
			fmt.Println("Usage: task-cli mark-done <task_id>")
			return
		}
		taskID := os.Args[2]
		err := MarkTask(taskID, "done")
		if err != nil {
			fmt.Println("Error marking task as done:", err)
		}

	case "list":
		// List all tasks with the given status.
		if len(os.Args) < 2 || len(os.Args) > 3 {
			fmt.Println("Usage: task-cli list || task-cli list todo || task-cli list in-progress || task-cli list done")
			return
		}

		if len(os.Args) == 2 {
			err := ListTasks("all")
			if err != nil {
				fmt.Println("Error listing all tasks:", err)
			}
		} else {
			if os.Args[2] == "todo" {
				err := ListTasks("todo")
				if err != nil {
					fmt.Println("Error listing todo tasks:", err)
				}
			} else if os.Args[2] == "in-progress" {
				err := ListTasks("in-progress")
				if err != nil {
					fmt.Println("Error listing in-progress tasks:", err)
				}
			} else if os.Args[2] == "done" {
				err := ListTasks("done")
				if err != nil {
					fmt.Println("Error listing done tasks:", err)
				}
			} else {
				fmt.Println("Usage: task-cli list || task-cli list todo || task-cli list in-progress || task-cli list done")
			}
		}

	default:
		fmt.Println("Unknown command:", command)
	}
}
