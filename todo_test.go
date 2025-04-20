package main

import (
	"os"
	"testing"
)

func TestSaveAndLoadTasks(t *testing.T) {
	filename := "test_tasks.json"
	defer os.Remove(filename)

	original := []Task{
		{Text: "Test Go", Done: false},
		{Text: "Write tests", Done: true},
	}

	err := SaveTasks(filename, original)
	if err != nil {
		t.Fatal("Failed to save tasks:", err)
	}

	loaded, err := LoadTasks(filename)
	if err != nil {
		t.Fatal("Failed to load tasks:", err)
	}

	if len(loaded) != len(original) {
		t.Fatal("Loaded tasks do not match")
	}
}
