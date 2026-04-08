import java.util.concurrent.ConcurrentLinkedQueue;
import java.util.List;
import java.util.ArrayList;
import java.io.*;
import java.nio.file.*;

public class ResultStore {
    private final ConcurrentLinkedQueue<String> results = new ConcurrentLinkedQueue<>();
    private final String outputFile;

    public ResultStore(String outputFile) {
        this.outputFile = outputFile;
    }

    public void addResult(String workerName, Task task, String processedData) {
        String result = String.format("[%s] Worker: %s, Task ID: %s, Result: %s",
                java.time.LocalDateTime.now(),
                workerName,
                task.getId(),
                processedData);
        results.add(result);
        LoggerUtil.log("Result stored for task: " + task.getId());
    }

    public void saveToFile() {
        try (BufferedWriter writer = Files.newBufferedWriter(Paths.get(outputFile))) {
            writer.write("=== DATA PROCESSING RESULTS ===\n");
            writer.write("Generated: " + java.time.LocalDateTime.now() + "\n\n");

            for (String result : results) {
                writer.write(result + "\n");
            }

            writer.write("\n=== END OF RESULTS ===\n");
            LoggerUtil.log("Results saved to file: " + outputFile);

        } catch (IOException e) {
            LoggerUtil.logError("Failed to save results to file: " + e.getMessage());
        }
    }

    public void printResults() {
        System.out.println("\n=== PROCESSING RESULTS ===");
        for (String result : results) {
            System.out.println(result);
        }
        System.out.println("Total tasks processed: " + results.size());
    }
}