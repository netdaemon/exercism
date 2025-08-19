public class CarsAssemble {
    private static final int CARS_PER_HOUR = 221;

    public double productionRatePerHour(int speed) {
        double carsProduced = speed * CARS_PER_HOUR;

        if (speed == 10) {
            return 0.77 * carsProduced;
        } else if (speed == 9) {
            return 0.80 * carsProduced;
        } else if ( speed >= 5 && speed <= 8) {
            return 0.90 * carsProduced;
        } else {
            return carsProduced;
        }
    }

    public int workingItemsPerMinute(int speed) {
        double ratePerHour = productionRatePerHour(speed);

        return (int)ratePerHour / 60;
    }
}
