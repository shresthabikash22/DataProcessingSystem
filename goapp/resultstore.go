package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

// Result represents a processed task result
type Result struct {
	Timestamp   time.Time
	WorkerName  string
	TaskID      string
	ProcessedData string
}

// ResultStore manages results collection
type ResultStore struct {
	results    []Result
	mu         sync.Mutex
	outputFile string
}

// NewResultStore creates a new result store
func NewResultStore(outputFile string) *ResultStore {
	return &ResultStore{
		results:    make([]Result, 0),
		outputFile: outputFile,
	}
}

// AddResult adds a result to the store
func (rs *ResultStore) AddResult(workerName string, task *Task, processedData string) {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	
	result := Result{
		Timestamp:     time.Now(),
		WorkerName:    workerName,
		TaskID:        task.ID,
		ProcessedData: processedData,
	}
	
	rs.results = append(rs.results, result)
	logInfo("Result stored for task: " + task.ID)
}

// GetCount returns number of results
func (rs *ResultStore) GetCount() int {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	return len(rs.results)
}

// SaveToFile saves all results to file
func (rs *ResultStore) SaveToFile() error {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	
	file, err := os.Create(rs.outputFile)
	if err != nil {
		return err
	}
	defer file.Close()
	
	// Write header
	_, err = file.WriteString("=== DATA PROCESSING RESULTS ===\n")
	if err != nil {
		return err
	}
	
	_, err = file.WriteString(fmt.Sprintf("Generated: %s\n\n", time.Now().Format(time.RFC3339)))
	if err != nil {
		return err
	}
	
	// Write results
	for _, result := range rs.results {
		line := fmt.Sprintf("[%s] Worker: %s, Task ID: %s, Result: %s\n",
			result.Timestamp.Format(time.RFC3339Nano),
			result.WorkerName,
			result.TaskID,
			result.ProcessedData)
		
		_, err = file.WriteString(line)
		if err != nil {
			return err
		}
	}
	
	_, err = file.WriteString("\n=== END OF RESULTS ===\n")
	if err != nil {
		return err
	}
	
	logInfo("Results saved to file: " + rs.outputFile)
	return nil
}

// PrintResults prints results to console
func (rs *ResultStore) PrintResults() {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	
	fmt.Println("\n=== PROCESSING RESULTS ===")
	for _, result := range rs.results {
		fmt.Printf("[%s] Worker: %s, Task ID: %s, Result: %s\n",
			result.Timestamp.Format(time.RFC3339Nano),
			result.WorkerName,
			result.TaskID,
			result.ProcessedData)
	}
	fmt.Printf("Total tasks processed: %d\n", len(rs.results))
}