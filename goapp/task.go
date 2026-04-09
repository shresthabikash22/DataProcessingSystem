package main

import (
	"fmt"
	"time"
)

// Task represents a unit of work
type Task struct {
	ID        string
	Data      string
	CreatedAt time.Time
}

// NewTask creates a new task with unique ID
func NewTask(data string) *Task {
	return &Task{
		ID:        generateID(),
		Data:      data,
		CreatedAt: time.Now(),
	}
}

func generateID() string {
	return fmt.Sprintf("task-%d", time.Now().UnixNano())
}

func (t *Task) String() string {
	return fmt.Sprintf("Task{ID='%s', Data='%s'}", t.ID, t.Data)
}