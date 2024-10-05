package main

import (
	"fmt"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// AddTask adds a new task to the tasks.json
func AddTask(description string) error {
	// Validate task description
	if err := ValidateDescription(description); err != nil {
		return err
	}

	// Load existing tasks
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	// Create new task
	newTask := Task{
		ID:          len(tasks) + 1, // Incremental ID based on task count
		Description: description,
		Status:      "todo", // Default status is "todo"
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Append new task to the list
	tasks = append(tasks, newTask)

	// Save tasks back to tasks.json
	if err := SaveTasks(tasks); err != nil {
		return err
	}

	fmt.Printf("Task added successfully (ID: %d)\n", newTask.ID)
	return nil
}
