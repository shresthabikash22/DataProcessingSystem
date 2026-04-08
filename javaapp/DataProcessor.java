import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.TimeUnit;

public class DataProcessor {
    private final TaskQueue taskQueue;
    private final ResultStore resultStore;
    private final ExecutorService executorService;
    private final List<Worker> workers;

    public DataProcessor(int numWorkers, String outputFile) {
        this.taskQueue = new TaskQueue();
        this.resultStore = new ResultStore(outputFile);
        this.executorService = Executors.newFixedThreadPool(numWorkers);
        this.workers = new ArrayList<>();

        // Create workers
        for (int i = 1; i <= numWorkers; i++) {
            workers.add(new Worker("Worker-" + i, taskQueue, resultStore));
        }
    }

    public void start() {
        LoggerUtil.log("Starting Data Processing System with " + workers.size() + " workers");

        // Submit all workers to executor
        for (Worker worker : workers) {
            executorService.submit(worker);
        }
    }

    public void addTasks(List<Task> tasks) {
        LoggerUtil.log("Adding " + tasks.size() + " tasks to queue");
        for (Task task : tasks) {
            taskQueue.addTask(task);
        }
    }

    public void shutdown() {
        LoggerUtil.log("Initiating system shutdown...");

        // Stop accepting new tasks
        taskQueue.stopAcceptingTasks();

        // Shutdown executor service gracefully
        executorService.shutdown();

        try {
            // Wait for workers to finish (max 30 seconds)
            if (!executorService.awaitTermination(30, TimeUnit.SECONDS)) {
                LoggerUtil.logError("Forcing shutdown - some workers didn't finish");
                executorService.shutdownNow();
            }
        } catch (InterruptedException e) {
            LoggerUtil.logError("Shutdown interrupted", e);
            executorService.shutdownNow();
        }

        // Save results to file
        resultStore.saveToFile();
        resultStore.printResults();

        LoggerUtil.log("System shutdown complete");
    }

    public int getTotalProcessed() {
        int total = 0;
        for (Worker worker : workers) {
            total += worker.getTasksProcessed();
        }
        return total;
    }

    public static void main(String[] args) {
        // Configuration
        int numWorkers = 4;
        int numTasks = 15;
        String outputFile = "processing_results.txt";

        // Create the processor
        DataProcessor processor = new DataProcessor(numWorkers, outputFile);

        // Start the workers
        processor.start();

        // Create and add tasks
        List<Task> tasks = new ArrayList<>();
        for (int i = 1; i <= numTasks; i++) {
            tasks.add(new Task("Data-" + i));
        }
        processor.addTasks(tasks);

        // Give a little time for processing to complete
        try {
            Thread.sleep(1000);
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
        }

        // Shutdown the system
        processor.shutdown();
        System.out.println("\n=== PROCESSING SUMMARY ===");
        System.out.println("Tasks created: " + numTasks);
        System.out.println("Tasks succeeded: " + processor.getTotalProcessed());
        System.out.println("Tasks failed: " + (numTasks - processor.getTotalProcessed()));
        System.out.println("Success rate: " +
                String.format("%.1f%%", (processor.getTotalProcessed() * 100.0 / numTasks)));

    }
}