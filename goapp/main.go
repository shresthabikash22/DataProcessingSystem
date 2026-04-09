package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	numWorkers = 4
	numTasks   = 15
	outputFile = "processing_results.txt"
	queueSize  = 100
)

func main() {
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())
	
	// Create components
	taskQueue := NewTaskQueue(queueSize)
	resultStore := NewResultStore(outputFile)
	
	// Create and start workers
	workers := make([]*Worker, numWorkers)
	var wg sync.WaitGroup
	
	for i := 1; i <= numWorkers; i++ {
		workerName := fmt.Sprintf("Worker-%d", i)
		workers[i-1] = NewWorker(workerName, taskQueue, resultStore)
		
		wg.Add(1)
		go func(w *Worker) {
			defer wg.Done()
			w.Start()
		}(workers[i-1])
	}
	
	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	
	go func() {
		<-sigChan
		logInfo("Initiating system shutdown...")
		
		// Close queue to stop accepting new tasks
		taskQueue.Close()
		
		// Wait for workers to finish with timeout
		done := make(chan struct{})
		go func() {
			wg.Wait()
			close(done)
		}()
		
		select {
		case <-done:
			logInfo("All workers finished gracefully")
		case <-time.After(30 * time.Second):
			logError("Shutdown timeout - forcing exit")
		}
		
		// Save results
		if err := resultStore.SaveToFile(); err != nil {
			logError("Failed to save results: " + err.Error())
		}
		resultStore.PrintResults()
		
		// Print summary
		printSummary(resultStore)
		
		os.Exit(0)
	}()
	
	// Create and add tasks
	logInfo(fmt.Sprintf("Starting Data Processing System with %d workers", numWorkers))
	logInfo(fmt.Sprintf("Adding %d tasks to queue", numTasks))
	
	for i := 1; i <= numTasks; i++ {
		task := NewTask(fmt.Sprintf("Data-%d", i))
		taskQueue.AddTask(task)
		logInfo(fmt.Sprintf("Task added to queue: %s", task.ID))
	}
	
	// Wait for shutdown signal
	<-make(chan struct{})
}

func printSummary(resultStore *ResultStore) {
	totalProcessed := resultStore.GetCount()
	
	fmt.Println("\n=== PROCESSING SUMMARY ===")
	fmt.Printf("Tasks created: %d\n", numTasks)
	fmt.Printf("Tasks succeeded: %d\n", totalProcessed)
	fmt.Printf("Tasks failed: %d\n", numTasks-totalProcessed)
	fmt.Printf("Success rate: %.1f%%\n", float64(totalProcessed)*100.0/float64(numTasks))
}