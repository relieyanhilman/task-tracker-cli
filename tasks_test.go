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

func TestDeleteTask(t *testing.T) {
	// Setup: Buat file tasks.json dengan beberapa task dummy untuk testing
	os.Remove("tasks.json")
	task1 := Task{
		ID:          1,
		Description: "Task 1",
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	task2 := Task{
		ID:          2,
		Description: "Task 2",
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	SaveTasks([]Task{task1, task2})

	// Case 1: Menghapus task dengan ID valid
	err := DeleteTask("1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	tasks, err := LoadTasks()
	if err != nil {
		t.Fatalf("Failed to load tasks: %v", err)
	}

	if len(tasks) != 1 {
		t.Fatalf("Expected 1 task left, got %d", len(tasks))
	}

	if tasks[0].ID != 2 {
		t.Errorf("Expected task ID 2, got %d", tasks[0].ID)
	}

	// Case 2: Menghapus task dengan ID yang tidak ada
	err = DeleteTask("99")
	if err == nil || err.Error() != "task ID not found" {
		t.Fatalf("Expected 'task ID not found' error, got %v", err)
	}

	// Case 3: Menghapus task dengan ID tidak valid (0 atau negatif)
	err = DeleteTask("0")
	if err == nil || err.Error() != "invalid task ID" {
		t.Fatalf("Expected 'invalid task ID' error, got %v", err)
	}

	err = DeleteTask("-1")
	if err == nil || err.Error() != "invalid task ID" {
		t.Fatalf("Expected 'invalid task ID' error, got %v", err)
	}

	// Case 4: Tidak memberikan ID
	err = DeleteTask("")
	if err == nil || err.Error() != "invalid task ID" {
		t.Fatalf("Expected 'invalid task ID' error, got %v", err)
	}

	// Cleanup: Hapus file tasks.json setelah test
	os.Remove("tasks.json")
}

func TestMarkTaskInProgress(t *testing.T) {
	// Setup: Buat file tasks.json dengan beberapa task dummy untuk testing
	os.Remove("tasks.json")
	task1 := Task{
		ID:          1,
		Description: "Task 1",
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	task2 := Task{
		ID:          2,
		Description: "Task 2",
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	SaveTasks([]Task{task1, task2})

	// Case 1: Menandai task dengan ID valid
	err := MarkTask("1", "in-progress")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	tasks, err := LoadTasks()
	if err != nil {
		t.Fatalf("Failed to load tasks: %v", err)
	}

	if tasks[0].Status != "in-progress" {
		t.Errorf("Expected status 'in-progress', got '%s'", tasks[0].Status)
	}

	if tasks[0].UpdatedAt.Before(tasks[0].CreatedAt) {
		t.Errorf("UpdatedAt should be later than CreatedAt")
	}

	// Case 2: Menandai task dengan ID yang tidak ada
	err = MarkTask("99", "in-progress")
	if err == nil || err.Error() != "task ID not found" {
		t.Fatalf("Expected 'task ID not found' error, got %v", err)
	}

	// Case 3: Menandai task dengan ID tidak valid (0 atau negatif)
	err = MarkTask("0", "in-progress")
	if err == nil || err.Error() != "invalid task ID" {
		t.Fatalf("Expected 'invalid task ID' error, got %v", err)
	}

	err = MarkTask("-1", "in-progress")
	if err == nil || err.Error() != "invalid task ID" {
		t.Fatalf("Expected 'invalid task ID' error, got %v", err)
	}

	// Case 4: Tidak memberikan ID
	err = MarkTask("", "in-progress")
	if err == nil || err.Error() != "invalid task ID" {
		t.Fatalf("Expected 'invalid task ID' error, got %v", err)
	}

	// Cleanup: Hapus file tasks.json setelah test
	os.Remove("tasks.json")
}

func TestMarkTaskDone(t *testing.T) {
	// Setup: Buat file tasks.json dengan beberapa task dummy untuk testing
	os.Remove("tasks.json")
	task1 := Task{
		ID:          1,
		Description: "Task 1",
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	task2 := Task{
		ID:          2,
		Description: "Task 2",
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	SaveTasks([]Task{task1, task2})

	// Case 1: Menandai task dengan ID valid
	err := MarkTask("1", "done")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	tasks, err := LoadTasks()
	if err != nil {
		t.Fatalf("Failed to load tasks: %v", err)
	}

	if tasks[0].Status != "done" {
		t.Errorf("Expected status 'done', got '%s'", tasks[0].Status)
	}

	if tasks[0].UpdatedAt.Before(tasks[0].CreatedAt) {
		t.Errorf("UpdatedAt should be later than CreatedAt")
	}

	// Case 2: Menandai task dengan ID yang tidak ada
	err = MarkTask("99", "done")
	if err == nil || err.Error() != "task ID not found" {
		t.Fatalf("Expected 'task ID not found' error, got %v", err)
	}

	// Case 3: Menandai task dengan ID tidak valid (0 atau negatif)
	err = MarkTask("0", "done")
	if err == nil || err.Error() != "invalid task ID" {
		t.Fatalf("Expected 'invalid task ID' error, got %v", err)
	}

	err = MarkTask("-1", "done")
	if err == nil || err.Error() != "invalid task ID" {
		t.Fatalf("Expected 'invalid task ID' error, got %v", err)
	}

	// Case 4: Tidak memberikan ID
	err = MarkTask("", "done")
	if err == nil || err.Error() != "invalid task ID" {
		t.Fatalf("Expected 'invalid task ID' error, got %v", err)
	}

	// Cleanup: Hapus file tasks.json setelah test
	os.Remove("tasks.json")
}
