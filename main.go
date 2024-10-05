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

	default:
		fmt.Println("Unknown command:", command)
	}
}
