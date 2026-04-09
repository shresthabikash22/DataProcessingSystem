package main

import (
	"sync"
)

// TaskQueue is a thread-safe queue using channels
type TaskQueue struct {
	tasks    chan *Task
	closed   bool
	closeMu  sync.Mutex
}

// NewTaskQueue creates a new task queue with buffer
func NewTaskQueue(bufferSize int) *TaskQueue {
	return &TaskQueue{
		tasks: make(chan *Task, bufferSize),
	}
}

// AddTask adds a task to the queue
func (q *TaskQueue) AddTask(task *Task) bool {
	q.closeMu.Lock()
	defer q.closeMu.Unlock()
	
	if q.closed {
		return false
	}
	
	q.tasks <- task
	return true
}

// GetTask retrieves a task from the queue (blocks if empty)
func (q *TaskQueue) GetTask() (*Task, bool) {
	task, ok := <-q.tasks
	return task, ok
}

// Close stops accepting new tasks
func (q *TaskQueue) Close() {
	q.closeMu.Lock()
	defer q.closeMu.Unlock()
	
	if !q.closed {
		q.closed = true
		close(q.tasks)
	}
}

// Size returns the current queue size
func (q *TaskQueue) Size() int {
	return len(q.tasks)
}