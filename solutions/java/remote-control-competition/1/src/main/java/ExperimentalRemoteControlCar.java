public class ExperimentalRemoteControlCar implements RemoteControlCar {
    private static final int SPEED = 20;
    private int distanceTravelled;
    
    public void drive() {
        distanceTravelled += SPEED;
    }

    public int getDistanceTravelled() {
        return distanceTravelled;
    }
}
