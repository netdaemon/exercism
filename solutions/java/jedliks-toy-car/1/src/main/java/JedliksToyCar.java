public class JedliksToyCar {
    private int distanceDriven;
    private int batteryLevel = 100;
    
    public static JedliksToyCar buy() {
        return new JedliksToyCar();
    }

    public String distanceDisplay() {
        return "Driven " + distanceDriven + " meters";
    }

    public String batteryDisplay() {
        return batteryLevel == 0 
            ? "Battery empty" 
            : "Battery at " + batteryLevel + "%";
    }

    public void drive() {
        if (batteryLevel == 0) {
            return;
        }
        
        distanceDriven += 20;
        batteryLevel--;
    }
}
