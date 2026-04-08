

import java.util.Random;

public class Worker implements Runnable {
    private final String name;
    private final TaskQueue taskQueue;
    private final ResultStore resultStore;
    private final Random random = new Random();
    private int tasksProcessed = 0;
    private volatile boolean running = true;

    public Worker(String name, TaskQueue taskQueue, ResultStore resultStore) {
        this.name = name;
        this.taskQueue = taskQueue;
        this.resultStore = resultStore;
    }

    public int getTasksProcessed() {
        return tasksProcessed;
    }

    @Override
    public void run() {
        LoggerUtil.log(name + " started");

        while (running) {
            try {
                Task task = taskQueue.getTask();

                if (task == null) {
                    // No more tasks, exit gracefully
                    LoggerUtil.log(name + " received shutdown signal");
                    break;
                }

                // Process the task
                String processedData = processTask(task);

                // Store the result
                resultStore.addResult(name, task, processedData);
                tasksProcessed++;

                LoggerUtil.log(name + " completed task " + task.getId() +
                        " (Processed: " + tasksProcessed + ")");

            } catch (InterruptedException e) {
                LoggerUtil.logError(name + " was interrupted: " + e.getMessage());
                Thread.currentThread().interrupt();
                break;
            } catch (Exception e) {
                LoggerUtil.logError(name + " encountered error: " + e.getMessage());
            }
        }

        LoggerUtil.log(name + " finished. Total tasks processed: " + tasksProcessed);
    }

    private String processTask(Task task) {
        // Simulate computational work with random delay (100-500ms)
        try {
            int delay = 100 + random.nextInt(400);
            Thread.sleep(delay);

            // Simulate data transformation
            String processed = "PROCESSED_" + task.getData().toUpperCase();

            // Occasionally simulate processing error (10% chance)
            if (random.nextInt(100) < 10) {
                throw new RuntimeException("Simulated processing error for task: " + task.getId());
            }

            return processed;

        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
            throw new RuntimeException("Processing interrupted", e);
        }
    }

    public void stop() {
        running = false;
    }
}