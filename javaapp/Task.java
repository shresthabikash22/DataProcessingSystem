
import java.time.LocalDateTime;
import java.util.UUID;

public class Task {
    private final String id;
    private final String data;
    private final LocalDateTime createdAt;

    public Task(String data) {
        this.id = UUID.randomUUID().toString();
        this.data = data;
        this.createdAt = LocalDateTime.now();
    }

    public String getId() { return id; }
    public String getData() { return data; }
    public LocalDateTime getCreatedAt() { return createdAt; }

    @Override
    public String toString() {
        return String.format("Task{id='%s', data='%s'}", id, data);
    }
}