import java.util.LinkedList;
import java.util.Queue;
import java.util.concurrent.locks.ReentrantLock;
import java.util.concurrent.locks.Condition;

public class TaskQueue {
    private final Queue<Task> queue = new LinkedList<>();
    private final ReentrantLock lock = new ReentrantLock();
    private final Condition notEmpty = lock.newCondition();
    private boolean acceptingTasks = true;

    public void addTask(Task task) {
        lock.lock();
        try {
            queue.add(task);
            notEmpty.signal(); // Wake up a waiting worker
            LoggerUtil.log("Task added to queue: " + task.getId());
        } finally {
            lock.unlock();
        }
    }

    public Task getTask() throws InterruptedException {
        lock.lock();
        try {
            while (queue.isEmpty() && acceptingTasks) {
                notEmpty.await(); // Wait for task to become available
            }
            if (queue.isEmpty() && !acceptingTasks) {
                return null; // No more tasks coming
            }
            return queue.poll();
        } finally {
            lock.unlock();
        }
    }

    public void stopAcceptingTasks() {
        lock.lock();
        try {
            acceptingTasks = false;
            notEmpty.signalAll(); // Wake up all waiting workers
        } finally {
            lock.unlock();
        }
    }

    public boolean isEmpty() {
        lock.lock();
        try {
            return queue.isEmpty();
        } finally {
            lock.unlock();
        }
    }
}