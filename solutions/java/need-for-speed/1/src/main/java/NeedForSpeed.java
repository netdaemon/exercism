class NeedForSpeed {
    public static final int MAX_BATTERY = 100;
    
    private int speed;
    private int batteryDrain;
    private int batteryLevel = MAX_BATTERY;
    private int metersDriven;
    
    NeedForSpeed(int speed, int batteryDrain) {
        this.speed = speed;
        this.batteryDrain = batteryDrain;
    }

    public boolean batteryDrained() {
        return batteryDrain > batteryLevel;
    }

    public int distanceDriven() {
        return metersDriven;
    }

    public void drive() {
        if (batteryDrained()) {
            return;
        }
        
        metersDriven += speed;
        batteryLevel -= batteryDrain;
    }

    public int getBatteryDrain() {
        return batteryDrain;
    }

    public int getSpeed() {
        return speed;
    }

    public static NeedForSpeed nitro() {
        return new NeedForSpeed(50, 4);
    }
}

class RaceTrack {
    private int distance;
    
    RaceTrack(int distance) {
        this.distance = distance;
    }

    public boolean canFinishRace(NeedForSpeed car) {
        int distanceCovered = (NeedForSpeed.MAX_BATTERY / car.getBatteryDrain()) * car.getSpeed();
        
        return distanceCovered >= distance;
    }
}
