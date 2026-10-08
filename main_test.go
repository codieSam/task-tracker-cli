package main

import "testing"

func TestTaskStatus(t *testing.T) {
	task := Task{
		ID:          1,
		Description: "Learning Go testing.",
		Status:      "todo",
	}
	if task.Status != "todo" {
		t.Errorf("Expected status %q, got %q", "todo", task.Status)
	}
}
