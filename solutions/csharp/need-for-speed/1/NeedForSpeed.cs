class RemoteControlCar
{
    private int speed;
    private int batteryDrain;
    public int batteryPercentage;
    private int distanceDriven;

    public RemoteControlCar(int speed, int batteryDrain)
    {
        this.speed = speed;
        this.batteryDrain = batteryDrain;
        batteryPercentage = 100;
        distanceDriven = 0;
    }

    public bool BatteryDrained() => batteryPercentage < batteryDrain;

    public int DistanceDriven() => distanceDriven;

    public void Drive()
    {
        if (BatteryDrained()) 
        {
            return;
        }

        distanceDriven += speed;
        batteryPercentage -= batteryDrain;
    }

    public static RemoteControlCar Nitro() => new RemoteControlCar(50, 4);
}

class RaceTrack
{
    private int distance;

    public RaceTrack(int distance)
    {
        this.distance = distance;
    }

    public bool TryFinishTrack(RemoteControlCar car)
    {
        while (!car.BatteryDrained())
        {
            car.Drive();
        }

        return car.DistanceDriven() >= distance;
    }
}
