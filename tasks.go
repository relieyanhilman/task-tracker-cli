package main

import (
	"errors"
	"fmt"
	"strconv"
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

// Fungsi untuk mengupdate task berdasarkan ID dan deskripsi baru
func UpdateTask(id string, newDescription string) error {
	// Validasi ID task yang diberikan
	taskID, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("invalid task ID")
	}

	// Validasi deskripsi task
	if err := ValidateDescription(newDescription); err != nil {
		return err
	}

	// Muat task dari file JSON
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	// Cari task dengan ID yang cocok
	taskFound := false
	for i, task := range tasks {
		if task.ID == taskID {
			// Task ditemukan, lakukan update deskripsi dan updatedAt
			tasks[i].Description = newDescription
			tasks[i].UpdatedAt = time.Now()
			taskFound = true
			break
		}
	}

	if !taskFound {
		return errors.New("task ID not found")
	}

	// Simpan task yang sudah diupdate ke file JSON
	if err := SaveTasks(tasks); err != nil {
		return err
	}

	fmt.Printf("Task (ID: %d) updated successfully\n", taskID)
	return nil
}

// Fungsi untuk menghapus task berdasarkan ID
func DeleteTask(id string) error {
	// Validasi ID task yang diberikan
	taskID, err := strconv.Atoi(id)
	if err != nil || taskID <= 0 {
		return errors.New("invalid task ID")
	}

	// Muat task dari file JSON
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	// Cari task dengan ID yang cocok
	taskFound := false
	newTasks := make([]Task, 0)
	for _, task := range tasks {
		if task.ID == taskID {
			taskFound = true
			continue // Task dengan ID ini akan di-skip (dihapus)
		}
		newTasks = append(newTasks, task)
	}

	if !taskFound {
		return errors.New("task ID not found")
	}

	// Simpan task yang sudah diupdate (tanpa task yang dihapus) ke file JSON
	if err := SaveTasks(newTasks); err != nil {
		return err
	}

	fmt.Printf("Task (ID: %d) deleted successfully\n", taskID)
	return nil
}

// Fungsi untuk menandai task sebagai "in-progress" atau "done" berdasarkan ID
func MarkTask(id string, newStatus string) error {
	// Validasi ID task yang diberikan
	taskID, err := strconv.Atoi(id)
	if err != nil || taskID <= 0 {
		return errors.New("invalid task ID")
	}

	// Muat task dari file JSON
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	// Cari task dengan ID yang cocok
	taskFound := false
	for i, task := range tasks {
		if task.ID == taskID {
			tasks[i].Status = newStatus     // Update status menjadi in-progress
			tasks[i].UpdatedAt = time.Now() // Update waktu
			taskFound = true
			break
		}
	}

	if !taskFound {
		return errors.New("task ID not found")
	}

	// Simpan task yang sudah diupdate ke file JSON
	if err := SaveTasks(tasks); err != nil {
		return err
	}

	fmt.Printf("Task (ID: %d) marked as %s successfully\n", taskID, newStatus)
	return nil
}

// Fungsi untuk menampilkan task berdasarkan argumen status yang diberikan.
func ListTasks(status string) error {
	// Muat semua task dari file JSON
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	var processedTask []Task
	// Looping dan print setiap task
	if status == "all" {
		for _, task := range tasks {
			processedTask = append(processedTask, task)
			fmt.Printf("ID: %d, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n",
				task.ID, task.Description, task.Status, task.CreatedAt.Format("2006-01-02 15:04:05"), task.UpdatedAt.Format("2006-01-02 15:04:05"))
		}
	} else {
		for _, task := range tasks {
			if task.Status == status {
				processedTask = append(processedTask, task)
				fmt.Printf("ID: %d, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n",
					task.ID, task.Description, task.Status, task.CreatedAt.Format("2006-01-02 15:04:05"), task.UpdatedAt.Format("2006-01-02 15:04:05"))
			}
		}
	}

	// Cek jika tidak ada task
	if len(processedTask) == 0 {
		fmt.Println("No tasks found.")
		return nil
	}

	return nil
}
