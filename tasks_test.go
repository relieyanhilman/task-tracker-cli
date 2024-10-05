package main

import (
	"os"
	"testing"
	"time"
)

func TestAddTask(t *testing.T) {
	// Setup: Hapus file tasks.json sebelum memulai test
	os.Remove("tasks.json")

	// Case 1: Tambah task valid
	err := AddTask("Write unit tests")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	tasks, err := LoadTasks()
	if err != nil {
		t.Fatalf("Failed to load tasks: %v", err)
	}

	if len(tasks) != 1 {
		t.Fatalf("Expected 1 task, got %d", len(tasks))
	}

	if tasks[0].Description != "Write unit tests" {
		t.Errorf("Expected task description 'Write unit tests', got '%s'", tasks[0].Description)
	}

	if tasks[0].Status != "todo" {
		t.Errorf("Expected status 'todo', got '%s'", tasks[0].Status)
	}

	// Case 2: Tambah task dengan deskripsi kosong
	err = AddTask("")
	if err == nil {
		t.Fatal("Expected error for empty description, got none")
	}

	// Cleanup: Hapus file tasks.json setelah test
	os.Remove("tasks.json")
}

func TestValidateDescription(t *testing.T) {
	// Case 1: Deskripsi valid
	err := ValidateDescription("Learn Go")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Case 2: Deskripsi kosong
	err = ValidateDescription("")
	if err == nil {
		t.Fatal("Expected error for empty description, got none")
	}

	// Case 3: Deskripsi dengan hanya spasi
	err = ValidateDescription("   ")
	if err == nil {
		t.Fatal("Expected error for spaces-only description, got none")
	}
}

func TestUpdateTask(t *testing.T) {
	// Setup: Buat file tasks.json dengan task dummy untuk testing
	os.Remove("tasks.json")
	task := Task{
		ID:          1,
		Description: "Original description",
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	SaveTasks([]Task{task})

	// Case 1: Update task dengan ID valid dan deskripsi baru
	err := UpdateTask("1", "Updated description")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	tasks, err := LoadTasks()
	if err != nil {
		t.Fatalf("Failed to load tasks: %v", err)
	}

	if tasks[0].Description != "Updated description" {
		t.Errorf("Expected description 'Updated description', got '%s'", tasks[0].Description)
	}

	if tasks[0].UpdatedAt.Before(tasks[0].CreatedAt) {
		t.Errorf("UpdatedAt should be later than CreatedAt")
	}

	// Case 2: Update task dengan ID yang tidak ada
	err = UpdateTask("99", "Non-existent task")
	if err == nil || err.Error() != "task ID not found" {
		t.Fatalf("Expected 'task ID not found' error, got %v", err)
	}

	// Case 3: Update task dengan ID tidak valid
	err = UpdateTask("invalid_id", "Invalid ID test")
	if err == nil || err.Error() != "invalid task ID" {
		t.Fatalf("Expected 'invalid task ID' error, got %v", err)
	}

	// Case 4: Update task dengan deskripsi kosong
	err = UpdateTask("1", "")
	if err == nil || err.Error() != "task description cannot be empty or just spaces" {
		t.Fatalf("Expected 'task description cannot be empty or just spaces' error, got %v", err)
	}

	// Case 5: Update task dengan deskripsi hanya spasi
	err = UpdateTask("1", "   ")
	if err == nil || err.Error() != "task description cannot be empty or just spaces" {
		t.Fatalf("Expected 'task description cannot be empty or just spaces' error, got %v", err)
	}

	// Cleanup: Hapus file tasks.json setelah test
	os.Remove("tasks.json")
}
