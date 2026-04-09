# Data Processing System - Java & Go Implementations

## 📋 Project Overview

A concurrent data processing system that simulates multiple worker threads/goroutines processing tasks in parallel. The system demonstrates proper concurrency management, thread-safe operations, and robust error handling in **both Java and Go**, highlighting the different concurrency models each language offers.

### Key Features (Common to Both Implementations)
- **Multiple workers** processing tasks concurrently
- **Thread-safe/goroutine-safe shared queue** with proper synchronization
- **Simulated processing delays** to represent real computational work
- **Comprehensive error handling** with graceful recovery
- **Detailed logging** of all system activities
- **Result persistence** to file output
## Repository Structure
```
/
├── javaapp
├── goapp
└── README.md

```
---

## ☕ Java Implementation

### Architecture

### Java Project Structure
```
javaapp/
├── Task.java
├── TaskQueue.java
├── Worker.java
├── ResultStore.java
├── DataProcessor.java
├── LoggerUtil.java
└── run.sh
```

### Java Requirements

- **Java Version**: Java 8 or higher
- **Memory**: Minimum 256MB RAM

### Java Setup & running

#### 1. Verify Java Installation
```bash
java -version
```


#### 2. Navigate to javaapp directory
```bash
cd javaapp
```
#### 3. Make run.sh Executable (Linux/macOS)
```bash
chmod +x run.sh
```
### Running the application
```bash
./run.sh
```
### Sample Output
```bash
$ ./run.sh 
2026-04-08 21:58:43.682 [INFO] Starting Data Processing System with 4 workers
2026-04-08 21:58:43.748 [INFO] Worker-3 started
2026-04-08 21:58:43.750 [INFO] Worker-2 started
2026-04-08 21:58:43.748 [INFO] Worker-1 started
2026-04-08 21:58:43.753 [INFO] Worker-4 started
2026-04-08 21:58:43.861 [INFO] Adding 15 tasks to queue
2026-04-08 21:58:43.862 [INFO] Task added to queue: b9806fd0-95e8-4aa5-b94f-8ab58fdb0251
2026-04-08 21:58:43.863 [INFO] Task added to queue: 0982c0da-8987-4354-8015-f9624f42c93f
2026-04-08 21:58:43.863 [INFO] Task added to queue: 1bf6394a-584f-4bfa-8e22-75f218383cac
2026-04-08 21:58:43.864 [INFO] Task added to queue: 29c07c9e-6304-407b-8602-8b6711ca858b
2026-04-08 21:58:43.864 [INFO] Task added to queue: e6c7f412-8af9-4280-9cfc-0c894609fbae
2026-04-08 21:58:43.864 [INFO] Task added to queue: fa901570-e3a5-4aa1-897c-6b5b363402d0
2026-04-08 21:58:43.864 [INFO] Task added to queue: 95737378-ddea-4bd1-a433-f02ece5befa4
2026-04-08 21:58:43.865 [INFO] Task added to queue: ca3a691a-9d8e-4155-9e04-dd8350026354
2026-04-08 21:58:43.865 [INFO] Task added to queue: 98d58685-b77b-4a7a-8285-9fdbdab8a316
2026-04-08 21:58:43.865 [INFO] Task added to queue: a3673a1d-d124-431a-8a5e-a38c51e29b11
2026-04-08 21:58:43.865 [INFO] Task added to queue: 74bee387-1d15-4fda-814b-d14289344a80
2026-04-08 21:58:43.866 [INFO] Task added to queue: cafd0827-a640-4281-9d0c-1cbc428a8196
2026-04-08 21:58:43.866 [INFO] Task added to queue: a8a5957e-e4ea-4ddb-adca-4bbf8b21266d
2026-04-08 21:58:43.866 [INFO] Task added to queue: fca7b3b5-f551-4654-9875-19c315325f6e
2026-04-08 21:58:43.867 [INFO] Task added to queue: cebc8cfb-cfba-4bd0-ab62-67431ef11cb1
2026-04-08 21:58:44.166 [INFO] Result stored for task: b9806fd0-95e8-4aa5-b94f-8ab58fdb0251
2026-04-08 21:58:44.167 [INFO] Result stored for task: 29c07c9e-6304-407b-8602-8b6711ca858b
2026-04-08 21:58:44.183 [INFO] Result stored for task: 0982c0da-8987-4354-8015-f9624f42c93f
2026-04-08 21:58:44.189 [INFO] Worker-4 completed task 29c07c9e-6304-407b-8602-8b6711ca858b (Processed: 1)
2026-04-08 21:58:44.191 [INFO] Worker-3 completed task b9806fd0-95e8-4aa5-b94f-8ab58fdb0251 (Processed: 1)
2026-04-08 21:58:44.198 [INFO] Worker-2 completed task 0982c0da-8987-4354-8015-f9624f42c93f (Processed: 1)
2026-04-08 21:58:44.293 [INFO] Result stored for task: 1bf6394a-584f-4bfa-8e22-75f218383cac
2026-04-08 21:58:44.293 [INFO] Worker-1 completed task 1bf6394a-584f-4bfa-8e22-75f218383cac (Processed: 1)
2026-04-08 21:58:44.306 [INFO] Result stored for task: e6c7f412-8af9-4280-9cfc-0c894609fbae
2026-04-08 21:58:44.306 [INFO] Worker-4 completed task e6c7f412-8af9-4280-9cfc-0c894609fbae (Processed: 2)
2026-04-08 21:58:44.495 [INFO] Result stored for task: 95737378-ddea-4bd1-a433-f02ece5befa4
2026-04-08 21:58:44.496 [INFO] Worker-2 completed task 95737378-ddea-4bd1-a433-f02ece5befa4 (Processed: 2)
2026-04-08 21:58:44.568 [ERROR] Worker-4 encountered error: Simulated processing error for task: 98d58685-b77b-4a7a-8285-9fdbdab8a316
2026-04-08 21:58:44.603 [INFO] Result stored for task: fa901570-e3a5-4aa1-897c-6b5b363402d0
2026-04-08 21:58:44.603 [INFO] Worker-3 completed task fa901570-e3a5-4aa1-897c-6b5b363402d0 (Processed: 2)
2026-04-08 21:58:44.695 [INFO] Result stored for task: ca3a691a-9d8e-4155-9e04-dd8350026354
2026-04-08 21:58:44.698 [INFO] Worker-1 completed task ca3a691a-9d8e-4155-9e04-dd8350026354 (Processed: 2)
2026-04-08 21:58:44.845 [INFO] Result stored for task: cafd0827-a640-4281-9d0c-1cbc428a8196
2026-04-08 21:58:44.845 [INFO] Worker-3 completed task cafd0827-a640-4281-9d0c-1cbc428a8196 (Processed: 3)
2026-04-08 21:58:44.868 [INFO] Initiating system shutdown...
2026-04-08 21:58:44.920 [INFO] Result stored for task: 74bee387-1d15-4fda-814b-d14289344a80
2026-04-08 21:58:44.924 [INFO] Worker-4 completed task 74bee387-1d15-4fda-814b-d14289344a80 (Processed: 3)
2026-04-08 21:58:44.932 [INFO] Result stored for task: a3673a1d-d124-431a-8a5e-a38c51e29b11
2026-04-08 21:58:44.934 [INFO] Worker-2 completed task a3673a1d-d124-431a-8a5e-a38c51e29b11 (Processed: 3)
2026-04-08 21:58:44.935 [INFO] Worker-2 received shutdown signal
2026-04-08 21:58:44.970 [INFO] Worker-2 finished. Total tasks processed: 3
2026-04-08 21:58:45.041 [INFO] Result stored for task: cebc8cfb-cfba-4bd0-ab62-67431ef11cb1
2026-04-08 21:58:45.042 [INFO] Worker-4 completed task cebc8cfb-cfba-4bd0-ab62-67431ef11cb1 (Processed: 4)
2026-04-08 21:58:45.045 [INFO] Worker-4 received shutdown signal
2026-04-08 21:58:45.049 [INFO] Worker-4 finished. Total tasks processed: 4
2026-04-08 21:58:45.093 [INFO] Result stored for task: a8a5957e-e4ea-4ddb-adca-4bbf8b21266d
2026-04-08 21:58:45.095 [INFO] Worker-1 completed task a8a5957e-e4ea-4ddb-adca-4bbf8b21266d (Processed: 3)
2026-04-08 21:58:45.099 [INFO] Worker-1 received shutdown signal
2026-04-08 21:58:45.100 [INFO] Worker-1 finished. Total tasks processed: 3
2026-04-08 21:58:45.306 [INFO] Result stored for task: fca7b3b5-f551-4654-9875-19c315325f6e
2026-04-08 21:58:45.306 [INFO] Worker-3 completed task fca7b3b5-f551-4654-9875-19c315325f6e (Processed: 4)
2026-04-08 21:58:45.307 [INFO] Worker-3 received shutdown signal
2026-04-08 21:58:45.307 [INFO] Worker-3 finished. Total tasks processed: 4
2026-04-08 21:58:45.339 [INFO] Results saved to file: processing_results.txt

=== PROCESSING RESULTS ===
[2026-04-08T21:58:44.156740455] Worker: Worker-4, Task ID: 29c07c9e-6304-407b-8602-8b6711ca858b, Result: PROCESSED_DATA-4
[2026-04-08T21:58:44.159398316] Worker: Worker-3, Task ID: b9806fd0-95e8-4aa5-b94f-8ab58fdb0251, Result: PROCESSED_DATA-1
[2026-04-08T21:58:44.182857745] Worker: Worker-2, Task ID: 0982c0da-8987-4354-8015-f9624f42c93f, Result: PROCESSED_DATA-2
[2026-04-08T21:58:44.292349027] Worker: Worker-1, Task ID: 1bf6394a-584f-4bfa-8e22-75f218383cac, Result: PROCESSED_DATA-3
[2026-04-08T21:58:44.306246659] Worker: Worker-4, Task ID: e6c7f412-8af9-4280-9cfc-0c894609fbae, Result: PROCESSED_DATA-5
[2026-04-08T21:58:44.495305468] Worker: Worker-2, Task ID: 95737378-ddea-4bd1-a433-f02ece5befa4, Result: PROCESSED_DATA-7
[2026-04-08T21:58:44.603296826] Worker: Worker-3, Task ID: fa901570-e3a5-4aa1-897c-6b5b363402d0, Result: PROCESSED_DATA-6
[2026-04-08T21:58:44.694504027] Worker: Worker-1, Task ID: ca3a691a-9d8e-4155-9e04-dd8350026354, Result: PROCESSED_DATA-8
[2026-04-08T21:58:44.844978328] Worker: Worker-3, Task ID: cafd0827-a640-4281-9d0c-1cbc428a8196, Result: PROCESSED_DATA-12
[2026-04-08T21:58:44.916748429] Worker: Worker-4, Task ID: 74bee387-1d15-4fda-814b-d14289344a80, Result: PROCESSED_DATA-11
[2026-04-08T21:58:44.932032633] Worker: Worker-2, Task ID: a3673a1d-d124-431a-8a5e-a38c51e29b11, Result: PROCESSED_DATA-10
[2026-04-08T21:58:45.040960006] Worker: Worker-4, Task ID: cebc8cfb-cfba-4bd0-ab62-67431ef11cb1, Result: PROCESSED_DATA-15
[2026-04-08T21:58:45.092672535] Worker: Worker-1, Task ID: a8a5957e-e4ea-4ddb-adca-4bbf8b21266d, Result: PROCESSED_DATA-13
[2026-04-08T21:58:45.306271068] Worker: Worker-3, Task ID: fca7b3b5-f551-4654-9875-19c315325f6e, Result: PROCESSED_DATA-14
Total tasks processed: 14
2026-04-08 21:58:45.355 [INFO] System shutdown complete

=== PROCESSING SUMMARY ===
Tasks created: 15
Tasks succeeded: 14
Tasks failed: 1
Success rate: 93.3%
```

## ⚙️ Configuration Options

| Parameter | Default | Location | Description |
|:----------|--------:|:---------|:------------|
| `numWorkers` | `4` | `DataProcessor.java` | Number of concurrent worker threads |
| `numTasks` | `15` | `DataProcessor.java` | Total number of tasks to process |
| `Error Rate` | `10%` | `Worker.java` | Simulated failure rate for testing |
| `Delay Range` | `100-500ms` | `Worker.java` | Processing time simulation range |
| `Output File` | `results.txt` | `DataProcessor.java` | File path for result storage |
| `Shutdown Timeout` | `30s` | `DataProcessor.java` | Graceful shutdown timeout |

## Customizing Behavior
#### Change Error Rate
In Worker.java, modify:

```java
// Change 10 to desired percentage (e.g., 30 for 30% error rate)
if (random.nextInt(100) < 10) {  // 10% error rate
    throw new RuntimeException("Simulated processing error");
}
```
#### Change Processing Speed
In Worker.java, modify:

```java
// Change delay range (currently 100-500ms)
int delay = 100 + random.nextInt(400);  // Min: 100ms, Max: 500ms
```
---

## Go Implementation
### Go Project Structure
```
goapp/
├── main.go          
├── task.go           
├── worker.go         
├── queue.go         
├── resultstore.go  
├── logger.go         
├── go.mod            
└── run.sh            
```
### Go Requirements
### Requirement	Version
- **Go Version**:	1.16 or higher
- **Memory**:	Minimum 128MB

### Go Setup & Running
### 1. Navigate to Go Directory
```bash
cd goapp
```
### 2. Initialize Go Module
```bash
go mod init dataprocessor
```
### 3. Run the Application
```
go run *.go
```
### Go Sample Output
```
$ go run *.go
2026-04-08 23:34:29.899 [INFO] Worker-1 started
2026-04-08 23:34:29.899 [INFO] Worker-2 started
2026-04-08 23:34:29.899 [INFO] Worker-3 started
2026-04-08 23:34:29.899 [INFO] Worker-4 started
2026-04-08 23:34:29.899 [INFO] Starting Data Processing System with 4 workers
2026-04-08 23:34:29.899 [INFO] Adding 15 tasks to queue
2026-04-08 23:34:29.899 [INFO] Task added to queue: task-1775709269899551866
2026-04-08 23:34:29.899 [INFO] Task added to queue: task-1775709269899570543
2026-04-08 23:34:29.899 [INFO] Task added to queue: task-1775709269899575341
2026-04-08 23:34:29.899 [INFO] Task added to queue: task-1775709269899580674
2026-04-08 23:34:29.899 [INFO] Task added to queue: task-1775709269899583835
2026-04-08 23:34:29.899 [INFO] Task added to queue: task-1775709269899586969
2026-04-08 23:34:29.899 [INFO] Task added to queue: task-1775709269899589948
2026-04-08 23:34:29.899 [INFO] Task added to queue: task-1775709269899592919
2026-04-08 23:34:29.899 [INFO] Task added to queue: task-1775709269899596015
2026-04-08 23:34:29.899 [INFO] Task added to queue: task-1775709269899599016
2026-04-08 23:34:29.899 [INFO] Task added to queue: task-1775709269899601976
2026-04-08 23:34:29.899 [INFO] Task added to queue: task-1775709269899604927
2026-04-08 23:34:29.899 [INFO] Task added to queue: task-1775709269899607860
2026-04-08 23:34:29.899 [INFO] Task added to queue: task-1775709269899610808
2026-04-08 23:34:29.899 [INFO] Task added to queue: task-1775709269899614026
2026-04-08 23:34:30.141 [INFO] Result stored for task: task-1775709269899570543
2026-04-08 23:34:30.141 [INFO] Worker-2 completed task task-1775709269899570543 (Processed: 1)
2026-04-08 23:34:30.172 [INFO] Result stored for task: task-1775709269899551866
2026-04-08 23:34:30.172 [INFO] Worker-1 completed task task-1775709269899551866 (Processed: 1)
2026-04-08 23:34:30.260 [INFO] Result stored for task: task-1775709269899575341
2026-04-08 23:34:30.260 [INFO] Worker-3 completed task task-1775709269899575341 (Processed: 1)
2026-04-08 23:34:30.381 [INFO] Result stored for task: task-1775709269899580674
2026-04-08 23:34:30.381 [INFO] Worker-4 completed task task-1775709269899580674 (Processed: 1)
2026-04-08 23:34:30.451 [INFO] Result stored for task: task-1775709269899586969
2026-04-08 23:34:30.451 [INFO] Worker-1 completed task task-1775709269899586969 (Processed: 2)
2026-04-08 23:34:30.505 [INFO] Result stored for task: task-1775709269899583835
2026-04-08 23:34:30.505 [INFO] Worker-2 completed task task-1775709269899583835 (Processed: 2)
2026-04-08 23:34:30.590 [INFO] Result stored for task: task-1775709269899589948
2026-04-08 23:34:30.590 [INFO] Worker-3 completed task task-1775709269899589948 (Processed: 2)
2026-04-08 23:34:30.628 [INFO] Result stored for task: task-1775709269899596015
2026-04-08 23:34:30.628 [INFO] Worker-1 completed task task-1775709269899596015 (Processed: 3)
2026-04-08 23:34:30.812 [INFO] Result stored for task: task-1775709269899601976
2026-04-08 23:34:30.812 [INFO] Worker-3 completed task task-1775709269899601976 (Processed: 3)
2026-04-08 23:34:30.869 [INFO] Result stored for task: task-1775709269899604927
2026-04-08 23:34:30.869 [INFO] Worker-1 completed task task-1775709269899604927 (Processed: 4)
2026-04-08 23:34:30.874 [INFO] Result stored for task: task-1775709269899592919
2026-04-08 23:34:30.874 [INFO] Worker-4 completed task task-1775709269899592919 (Processed: 2)
2026-04-08 23:34:30.933 [INFO] Result stored for task: task-1775709269899607860
2026-04-08 23:34:30.933 [INFO] Worker-3 completed task task-1775709269899607860 (Processed: 4)
2026-04-08 23:34:30.960 [INFO] Result stored for task: task-1775709269899599016
2026-04-08 23:34:30.960 [INFO] Worker-2 completed task task-1775709269899599016 (Processed: 3)
2026-04-08 23:34:31.281 [INFO] Result stored for task: task-1775709269899610808
2026-04-08 23:34:31.281 [INFO] Worker-1 completed task task-1775709269899610808 (Processed: 5)
2026-04-08 23:34:31.370 [INFO] Result stored for task: task-1775709269899614026
2026-04-08 23:34:31.370 [INFO] Worker-4 completed task task-1775709269899614026 (Processed: 3)
^C2026-04-08 23:34:46.250 [INFO] Initiating system shutdown...
2026-04-08 23:34:46.250 [INFO] Worker-4 received shutdown signal
2026-04-08 23:34:46.250 [INFO] Worker-4 finished. Total tasks processed: 3
2026-04-08 23:34:46.250 [INFO] Worker-1 received shutdown signal
2026-04-08 23:34:46.250 [INFO] Worker-1 finished. Total tasks processed: 5
2026-04-08 23:34:46.250 [INFO] Worker-2 received shutdown signal
2026-04-08 23:34:46.250 [INFO] Worker-2 finished. Total tasks processed: 3
2026-04-08 23:34:46.250 [INFO] Worker-3 received shutdown signal
2026-04-08 23:34:46.250 [INFO] Worker-3 finished. Total tasks processed: 4
2026-04-08 23:34:46.250 [INFO] All workers finished gracefully
2026-04-08 23:34:46.250 [INFO] Results saved to file: processing_results.txt

=== PROCESSING RESULTS ===
[2026-04-08T23:34:30.141106257-05:00] Worker: Worker-2, Task ID: task-1775709269899570543, Result: PROCESSED_Data-2
[2026-04-08T23:34:30.172166027-05:00] Worker: Worker-1, Task ID: task-1775709269899551866, Result: PROCESSED_Data-1
[2026-04-08T23:34:30.260155921-05:00] Worker: Worker-3, Task ID: task-1775709269899575341, Result: PROCESSED_Data-3
[2026-04-08T23:34:30.381509236-05:00] Worker: Worker-4, Task ID: task-1775709269899580674, Result: PROCESSED_Data-4
[2026-04-08T23:34:30.451757012-05:00] Worker: Worker-1, Task ID: task-1775709269899586969, Result: PROCESSED_Data-6
[2026-04-08T23:34:30.505046246-05:00] Worker: Worker-2, Task ID: task-1775709269899583835, Result: PROCESSED_Data-5
[2026-04-08T23:34:30.590376863-05:00] Worker: Worker-3, Task ID: task-1775709269899589948, Result: PROCESSED_Data-7
[2026-04-08T23:34:30.628663693-05:00] Worker: Worker-1, Task ID: task-1775709269899596015, Result: PROCESSED_Data-9
[2026-04-08T23:34:30.812110593-05:00] Worker: Worker-3, Task ID: task-1775709269899601976, Result: PROCESSED_Data-11
[2026-04-08T23:34:30.869425398-05:00] Worker: Worker-1, Task ID: task-1775709269899604927, Result: PROCESSED_Data-12
[2026-04-08T23:34:30.874637888-05:00] Worker: Worker-4, Task ID: task-1775709269899592919, Result: PROCESSED_Data-8
[2026-04-08T23:34:30.93303787-05:00] Worker: Worker-3, Task ID: task-1775709269899607860, Result: PROCESSED_Data-13
[2026-04-08T23:34:30.960340421-05:00] Worker: Worker-2, Task ID: task-1775709269899599016, Result: PROCESSED_Data-10
[2026-04-08T23:34:31.281834671-05:00] Worker: Worker-1, Task ID: task-1775709269899610808, Result: PROCESSED_Data-14
[2026-04-08T23:34:31.3701481-05:00] Worker: Worker-4, Task ID: task-1775709269899614026, Result: PROCESSED_Data-15
Total tasks processed: 15

=== PROCESSING SUMMARY ===
Tasks created: 15
Tasks succeeded: 15
Tasks failed: 0
Success rate: 100.0%
```
### Customizing Go Behavior
### Change Error Rate
In worker.go, modify:

``` go
// Change 10 to desired percentage (e.g., 30 for 30% error rate)
if rand.Intn(100) < 10 {  // 10% error rate
    return "", fmtError("simulated processing error")
}
```
### Change Processing Speed
In worker.go, modify:

```go
// Change delay range (currently 100-500ms)
delay := time.Duration(100+rand.Intn(400)) * time.Millisecond
```
