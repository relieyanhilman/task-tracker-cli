package main

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
)

// ValidateDescription validates the task description
func ValidateDescription(description string) error {
	trimmed := strings.TrimSpace(description)
	if len(trimmed) == 0 {
		return errors.New("task description cannot be empty or just spaces")
	}
	return nil
}

// LoadTasks reads tasks from tasks.json
func LoadTasks() ([]Task, error) {
	file, err := os.ReadFile("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			// Return an empty slice if file doesn't exist
			return []Task{}, nil
		}
		return nil, err
	}

	var tasks []Task
	if err := json.Unmarshal(file, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

// SaveTasks writes tasks to tasks.json
func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("tasks.json", data, 0644)
}
