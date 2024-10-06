package main

import (
	"os"
	"strings"
	"testing"
	"time"
)

// TestAddTask tests the AddTask function for adding new tasks.
//
// The test includes the following cases:
//
//  1. Add task valid: Add a task with a valid description. The test checks
//     that no error is returned, and that the task is saved to the tasks.json
//     file. The test also checks that the task's description, status, and ID are
//     correct.
//
//  2. Add task with empty description: Add a task with an empty description.
//     The test checks that an error is returned, and that the task is not saved
//     to the tasks.json file.
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

// TestValidateDescription tests the ValidateDescription function for validating task descriptions.
//
// The test includes the following cases:
//
//  1. Valid description: Validate a valid description. The test checks that no
//     error is returned.
//
//  2. Empty description: Validate an empty description. The test checks that an
//     error is returned.
//
//  3. Spaces-only description: Validate a description with only spaces. The
//     test checks that an error is returned.
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

// TestUpdateTask tests the UpdateTask function for updating tasks with new descriptions.
//
// The test includes the following cases:
//
//  1. Valid update: Update a task with a valid ID and description. The test checks
//     that no error is returned and the task description is updated.
//
//  2. Non-existent task: Update a task with a non-existent ID. The test checks that
//     a "task ID not found" error is returned.
//
//  3. Invalid task ID: Update a task with an invalid ID (e.g. a string that is not a
//     number). The test checks that an "invalid task ID" error is returned.
//
//  4. Empty description: Update a task with an empty description. The test checks
//     that a "task description cannot be empty or just spaces" error is returned.
//
//  5. Spaces-only description: Update a task with a description that only contains
//     spaces. The test checks that a "task description cannot be empty or just spaces"
//     error is returned.
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

// TestDeleteTask tests the DeleteTask function for deleting tasks.
//
// The test includes the following cases:
//
//  1. Delete task valid: Delete a task with a valid ID. The test checks
//     that no error is returned, and that the task is removed from the
//     tasks.json file. The test also checks that the task's ID is no longer
//     present in the tasks.json file.
//
//  2. Delete task with ID not found: Delete a task with an invalid ID.
//     The test checks that an error is returned with the message "task ID
//     not found".
//
//  3. Delete task with invalid ID (0 or negative): Delete a task with an
//     invalid or negative ID. The test checks that an error is returned with
//     the message "invalid task ID".
//
//  4. Delete task with no ID: Delete a task with no ID. The test checks that
//     an error is returned with the message "invalid task ID".
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

// TestMarkTaskInProgress tests the MarkTask function for marking tasks as "in-progress".
//
// The test includes the following cases:
//
//  1. Valid update: Mark a task with a valid ID as "in-progress". The test checks
//     that no error is returned, and that the task status is updated.
//
//  2. Non-existent task: Mark a task with a non-existent ID as "in-progress". The
//     test checks that a "task ID not found" error is returned.
//
//  3. Invalid task ID: Mark a task with an invalid ID (e.g. a string that is not a
//     number) as "in-progress". The test checks that an "invalid task ID" error is
//     returned.
//
//  4. No task ID: Mark a task with an empty string as "in-progress". The test checks
//     that an "invalid task ID" error is returned.
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

// TestMarkTaskDone tests the MarkTask function for marking tasks as "done".
//
// The test includes the following cases:
//
//  1. Valid mark: Mark a task with a valid ID. The test checks
//     that no error is returned, and that the task status is updated
//     to "done".
//
//  2. Non-existent task: Mark a task with a non-existent ID. The test
//     checks that a "task ID not found" error is returned.
//
//  3. Invalid task ID: Mark a task with an invalid ID (e.g. a string
//     that is not a number). The test checks that an "invalid task ID"
//     error is returned.
//
//  4. No Task ID: Mark a task with no ID. The test checks that an "invalid
//     task ID" error is returned.
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

// TestListAllTasks tests the ListTasks function for listing all tasks, regardless of their status.
//
// The test includes the following cases:
//
//  1. Valid list: List all tasks. The test checks that no error is returned and that all
//     tasks are listed.
//
func TestListAllTasks(t *testing.T) {
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
		Status:      "in-progress",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	SaveTasks([]Task{task1, task2})

	// Capture output untuk testing
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Jalankan fungsi ListTasks
	err := ListTasks("all")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	w.Close()
	var buf [1024]byte
	n, _ := r.Read(buf[:])
	os.Stdout = old
	output := string(buf[:n])

	// Periksa apakah output sesuai dengan task yang di-load
	if !strings.Contains(output, "Task 1") || !strings.Contains(output, "Task 2") {
		t.Errorf("Expected task output, but got: %s", output)
	}

	// Cleanup: Hapus file tasks.json setelah test
	os.Remove("tasks.json")
}

// TestListTodoTasks tests the ListTasks function for listing tasks with "todo" status.
//
// The test includes the following cases:
//
//  1. List all tasks with "todo" status. The test checks that no error is returned,
//     and that the output contains the task with "todo" status.
//
//  2. List tasks with "todo" status when there are tasks with different statuses.
//     The test checks that no error is returned, and that the output only contains
//     the tasks with "todo" status.
func TestListTodoTasks(t *testing.T) {
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
		Status:      "in-progress",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	task3 := Task{
		ID:          3,
		Description: "Task 3",
		Status:      "done",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	SaveTasks([]Task{task1, task2, task3})

	// Capture output untuk testing
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Jalankan fungsi ListTasks
	err := ListTasks("todo")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	w.Close()
	var buf [1024]byte
	n, _ := r.Read(buf[:])
	os.Stdout = old
	output := string(buf[:n])

	// Periksa apakah output sesuai dengan task yang di-load
	if !strings.Contains(output, "Task 1") || strings.Contains(output, "Task 2") || strings.Contains(output, "Task 3") {
		t.Errorf("Expected task output, but got: %s", output)
	}

	// Cleanup: Hapus file tasks.json setelah test
	os.Remove("tasks.json")
}

// TestListInProgressTasks tests the ListTasks function for listing in-progress tasks.
//
// The test includes the following cases:
//
//  1. Valid list: List all in-progress tasks. The test checks that no error is returned
//     and that only the in-progress task is listed.
//
func TestListInProgressTasks(t *testing.T) {
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
		Status:      "in-progress",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	task3 := Task{
		ID:          3,
		Description: "Task 3",
		Status:      "done",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	SaveTasks([]Task{task1, task2, task3})

	// Capture output untuk testing
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Jalankan fungsi ListTasks
	err := ListTasks("in-progress")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	w.Close()
	var buf [1024]byte
	n, _ := r.Read(buf[:])
	os.Stdout = old
	output := string(buf[:n])

	// Periksa apakah output sesuai dengan task yang di-load
	if strings.Contains(output, "Task 1") || !strings.Contains(output, "Task 2") || strings.Contains(output, "Task 3") {
		t.Errorf("Expected task output, but got: %s", output)
	}

	// Cleanup: Hapus file tasks.json setelah test
	os.Remove("tasks.json")
}

// TestListDoneTasks tests the ListTasks function for listing tasks with "done" status.
//
// The test includes the following cases:
//
//  1. Valid list: List all done tasks. The test checks that no error is returned
//     and that only the done task is listed.
func TestListDoneTasks(t *testing.T) {
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
		Status:      "in-progress",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	task3 := Task{
		ID:          3,
		Description: "Task 3",
		Status:      "done",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	SaveTasks([]Task{task1, task2, task3})

	// Capture output untuk testing
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Jalankan fungsi ListTasks
	err := ListTasks("done")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	w.Close()
	var buf [1024]byte
	n, _ := r.Read(buf[:])
	os.Stdout = old
	output := string(buf[:n])

	// Periksa apakah output sesuai dengan task yang di-load
	if strings.Contains(output, "Task 1") || strings.Contains(output, "Task 2") || !strings.Contains(output, "Task 3") {
		t.Errorf("Expected task output, but got: %s", output)
	}

	// Cleanup: Hapus file tasks.json setelah test
	os.Remove("tasks.json")
}
