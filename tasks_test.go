package main

import (
	"os"
	"testing"
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
