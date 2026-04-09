package main

import (
	"math/rand"
	"time"
)

// Worker represents a worker goroutine
type Worker struct {
	name        string
	taskQueue   *TaskQueue
	resultStore *ResultStore
	processed   int
	stopChan    chan bool
}

// NewWorker creates a new worker
func NewWorker(name string, taskQueue *TaskQueue, resultStore *ResultStore) *Worker {
	return &Worker{
		name:        name,
		taskQueue:   taskQueue,
		resultStore: resultStore,
		stopChan:    make(chan bool),
	}
}

// Start begins the worker's processing loop
func (w *Worker) Start() {
	logInfo(w.name + " started")
	
	for {
		select {
		case <-w.stopChan:
			logInfo(w.name + " received shutdown signal")
			logInfo(w.name + " finished. Total tasks processed: " + itoa(w.processed))
			return
		default:
			// Get task from queue (blocks until task available or queue closed)
			task, ok := w.taskQueue.GetTask()
			if !ok {
				// Queue is closed and empty
				logInfo(w.name + " received shutdown signal")
				logInfo(w.name + " finished. Total tasks processed: " + itoa(w.processed))
				return
			}
			
			// Process the task
			processedData, err := w.processTask(task)
			if err != nil {
				logError(w.name + " encountered error: " + err.Error())
				// Failed tasks are not retried (could be added later)
				continue
			}
			
			// Store the result
			w.resultStore.AddResult(w.name, task, processedData)
			w.processed++
			logInfo(w.name + " completed task " + task.ID + " (Processed: " + itoa(w.processed) + ")")
		}
	}
}

// processTask simulates task processing with delay and random errors
func (w *Worker) processTask(task *Task) (string, error) {
	// Simulate processing delay (100-500ms)
	delay := time.Duration(100+rand.Intn(400)) * time.Millisecond
	time.Sleep(delay)
	
	// Simulate random errors (10% chance)
	if rand.Intn(100) < 10 {
		return "", fmtError("simulated processing error for task: " + task.ID)
	}
	
	// Simulate data transformation
	processed := "PROCESSED_" + task.Data
	
	return processed, nil
}

// Stop signals the worker to stop
func (w *Worker) Stop() {
	close(w.stopChan)
}

// Helper functions
func itoa(n int) string {
	return string(rune(n + '0'))
}