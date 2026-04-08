
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;

public class LoggerUtil {
    private static final DateTimeFormatter FORMATTER =
            DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss.SSS");

    public static void log(String message) {
        String timestamp = LocalDateTime.now().format(FORMATTER);
        System.out.println(timestamp + " [INFO] " + message);
    }

    public static void logError(String message) {
        String timestamp = LocalDateTime.now().format(FORMATTER);
        System.err.println(timestamp + " [ERROR] " + message);
    }

    public static void logError(String message, Throwable throwable) {
        String timestamp = LocalDateTime.now().format(FORMATTER);
        System.err.println(timestamp + " [ERROR] " + message);
        throwable.printStackTrace(System.err);
    }
}