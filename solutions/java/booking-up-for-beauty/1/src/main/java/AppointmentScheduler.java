import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;

class AppointmentScheduler {
    public LocalDateTime schedule(String appointmentDateDescription) {
        DateTimeFormatter parser = DateTimeFormatter.ofPattern("MM/dd/yyyy HH:mm:ss");

        return LocalDateTime.parse(appointmentDateDescription, parser);
    }

    public boolean hasPassed(LocalDateTime appointmentDate) {
        return appointmentDate.isBefore(LocalDateTime.now());
    }

    public boolean isAfternoonAppointment(LocalDateTime appointmentDate) {
        int hour = appointmentDate.getHour();

        return hour >= 12 && hour < 18;
    }

    public String getDescription(LocalDateTime appointmentDate) {
        DateTimeFormatter dateFormatter = DateTimeFormatter.ofPattern("EEEE, MMMM d, Y");
        DateTimeFormatter timeFormatter = DateTimeFormatter.ofPattern("h:mm a");
        
        String date = appointmentDate.format(dateFormatter);
        String time = appointmentDate.format(timeFormatter);

        return "You have an appointment on " + date + ", at " + time + ".";
    }

    public LocalDate getAnniversaryDate() {
        return LocalDate.of(LocalDateTime.now().getYear(), 9, 15);
    }
}
