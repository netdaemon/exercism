class Lasagna
{
    public int ExpectedMinutesInOven() => 40;

    public int RemainingMinutesInOven(int minutesInOven) => ExpectedMinutesInOven() - minutesInOven;

    public int PreparationTimeInMinutes(int numberOfLayers) => 2 * numberOfLayers;

    public int ElapsedTimeInMinutes(int numberOfLayers, int minutesInOven) => PreparationTimeInMinutes(numberOfLayers) + minutesInOven;
}
