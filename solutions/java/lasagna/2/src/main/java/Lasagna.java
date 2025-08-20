public class Lasagna {
    private static final int EXPECTED_MINUTES_IN_OVEN = 40;
    private static final int PREPARATION_TIME_IN_MINUTES = 2;
    
    public int expectedMinutesInOven() {
        return EXPECTED_MINUTES_IN_OVEN;
    }

    public int remainingMinutesInOven(int actualMinutes) {
        return expectedMinutesInOven() - actualMinutes;        
    }

    public int preparationTimeInMinutes(int numberOfLayers) {
        return numberOfLayers * PREPARATION_TIME_IN_MINUTES;
    }

    public int totalTimeInMinutes(int numberOfLayers, int numberOfMinutes) {
        return preparationTimeInMinutes(numberOfLayers) + numberOfMinutes;
    }
}
