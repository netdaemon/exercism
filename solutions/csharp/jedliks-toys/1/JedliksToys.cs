class RemoteControlCar
{
    private int _metersDriven = 0;
    private int _batteryPercentage = 100;


    public static RemoteControlCar Buy() => new RemoteControlCar();

    public string DistanceDisplay() => $"Driven {_metersDriven} meters";

    public string BatteryDisplay() => _batteryPercentage == 0 ? "Battery empty" : $"Battery at {_batteryPercentage}%";

    public void Drive()
    {
        if (_batteryPercentage == 0)
        {
            return;
        }

        _metersDriven += 20;
        _batteryPercentage--;
    }
}
