public class LogLevels {
    
    public static String message(String logLine) {
        return logLine.substring(logLine.indexOf(":") + 1).strip();
    }

    public static String logLevel(String logLine) {
        String[] splitLogLine = logLine.split(":");

        return splitLogLine[0].substring(1, splitLogLine[0].length() - 1).toLowerCase();
    }

    public static String reformat(String logLine) {
        return String.format("%s (%s)", message(logLine), logLevel(logLine));
    }
}
