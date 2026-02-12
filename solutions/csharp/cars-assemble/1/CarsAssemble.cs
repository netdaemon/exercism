static class AssemblyLine
{
    private const int CarsPerHour = 221;
    
    public static double SuccessRate(int speed) => speed switch
    {
        10 => 0.77,
        9 => 0.80,
        <= 8 and >= 5 => 0.90,
        <= 4 and >= 1 => 1.00,
        _ => 0.00
    };
    
    public static double ProductionRatePerHour(int speed) => CarsPerHour * speed * SuccessRate(speed);

    public static int WorkingItemsPerMinute(int speed) => (int)ProductionRatePerHour(speed) / 60;
}
