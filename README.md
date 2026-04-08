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

### Java Installation & Setup

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


